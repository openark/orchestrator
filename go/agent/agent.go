/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package agent

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/inst"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
	"github.com/patrickmn/go-cache"
)

var agentsCache = cache.New(time.Duration(config.Config.AgentCacheTTLSeconds)*time.Second, time.Duration(config.Config.AgentCacheTTLSeconds*2)*time.Second)

type Agent struct {
	Info        *Info
	Data        *Data
	LastSeen    time.Time
	ClusterName string
}

type Info struct {
	Hostname  string
	Port      int
	Token     string
	MySQLPort int
}

type Data struct {
	LocalSnapshotsHosts   []string         // AvailableLocalSnapshots in Orchestrator
	SnaphostHosts         []string         // AvailableSnapshots in Orchestrator
	LogicalVolumes        []*LogicalVolume // pass by reference ??
	MountPoint            *Mount           // pass by reference ??
	BackupDir             string
	BackupDirDiskFree     int64
	MySQLRunning          bool
	MySQLDatadir          string
	MySQLDatadirDiskUsed  int64
	MySQLDatadirDiskFree  int64
	MySQLVersion          string
	MySQLDatabases        map[string]*MySQLDatabase
	MySQLErrorLogTail     []string
	AvailiableSeedMethods map[SeedMethod]SeedMethodOpts
}

// MySQLDatabase describes a MySQL database
type MySQLDatabase struct {
	Engines []Engine
	Size    int64
}

// LogicalVolume describes an LVM volume
type LogicalVolume struct {
	Name            string
	GroupName       string
	Path            string
	IsSnapshot      bool
	SnapshotPercent float64
}

// Mount describes a file system mount point
type Mount struct {
	Path       string
	Device     string
	LVPath     string
	FileSystem string
	IsMounted  bool
	DiskUsage  int64
}

type SeedMethodOpts struct {
	BackupSide       SeedSide
	SupportedEngines []Engine
	BackupToDatadir  bool
}

type BackupMetadata struct {
	LogFile      string
	LogPos       int64
	GtidExecuted string
}

type AgentStatus int

const (
	Active AgentStatus = iota
	Inactive
)

func (a AgentStatus) String() string {
	return [...]string{"Active", "Inactive"}[a]
}

func (a AgentStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(a.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

var toAgentStatus = map[string]AgentStatus{
	"Active":   Active,
	"Inactive": Inactive,
}

type httpMethodFunc func(uri string) (resp *http.Response, err error)

var httpClient *http.Client
var httpClientMutex = &sync.Mutex{}

// InitHttpClient gets called once, and initializes httpClient according to config.Config
func InitHttpClient() {
	httpClientMutex.Lock()
	defer httpClientMutex.Unlock()

	if httpClient != nil {
		return
	}

	httpTimeout := time.Duration(time.Duration(config.AgentHttpTimeoutSeconds) * time.Second)
	dialTimeout := func(network, addr string) (net.Conn, error) {
		return net.DialTimeout(network, addr, httpTimeout)
	}
	httpTransport := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: config.Config.AgentSSLSkipVerify},
		Dial:                  dialTimeout,
		ResponseHeaderTimeout: httpTimeout,
	}
	httpClient = &http.Client{Transport: httpTransport}
}

// httpGet is a convenience method for getting http response from URL, optionaly skipping SSL cert verification
func httpGet(url string) (resp *http.Response, err error) {
	return httpClient.Get(url)
}

// executeAgentCommand requests an agent to execute a command via HTTP api
func executeAgentCommand(hostname string, command string, onResponse *func([]byte)) (Agent, error) {
	httpFunc := func(uri string) (resp *http.Response, err error) {
		return httpGet(uri)
	}
	return executeAgentCommandWithMethodFunc(hostname, command, httpFunc, onResponse)
}

// executeAgentCommandWithMethodFunc requests an agent to execute a command via HTTP api, either GET or POST,
// with specific http method implementation by the caller
func executeAgentCommandWithMethodFunc(hostname string, command string, methodFunc httpMethodFunc, onResponse *func([]byte)) (Agent, error) {
	agent, token, err := readAgentBasicInfo(hostname)
	if err != nil {
		return agent, err
	}

	// All seems to be in order. Now make some inquiries from orchestrator-agent service:
	uri := baseAgentUri(agent.Info.Hostname, agent.Info.Port)

	var fullCommand string
	if strings.Contains(command, "?") {
		fullCommand = fmt.Sprintf("%s&token=%s", command, token)
	} else {
		fullCommand = fmt.Sprintf("%s?token=%s", command, token)
	}
	log.Debugf("orchestrator-agent command: %s", fullCommand)
	agentCommandURI := fmt.Sprintf("%s/%s", uri, fullCommand)

	body, err := readResponse(methodFunc(agentCommandURI))
	if err != nil {
		return agent, log.Errore(err)
	}
	if onResponse != nil {
		(*onResponse)(body)
	}
	auditAgentOperation("agent-command", &agent, command)

	return agent, err
}

// readResponse returns the body of an HTTP response
func readResponse(res *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.Status == "500" {
		return body, errors.New("Response Status 500")
	}

	return body, nil
}

// baseAgentUri returns the base URI for accessing an agent
func baseAgentUri(agentHostname string, agentPort int) string {
	protocol := "http"
	if config.Config.AgentsUseSSL {
		protocol = "https"
	}
	uri := fmt.Sprintf("%s://%s:%d/api", protocol, agentHostname, agentPort)
	log.Debugf("orchestrator-agent uri: %s", uri)
	return uri
}

// AuditAgentOperation creates and writes a new audit entry by given agent
func auditAgentOperation(auditType string, agent *Agent, message string) error {
	instanceKey := &inst.InstanceKey{}
	if agent != nil {
		instanceKey = &inst.InstanceKey{Hostname: agent.Info.Hostname, Port: agent.Info.MySQLPort}
	}
	return inst.AuditOperation(auditType, instanceKey, message)
}

// GetAgent gets a single agent status from the agent service

/*
func GetAgent(hostname string) (Agent, error) {
	agent, token, err := readAgentBasicInfo(hostname)
	if err != nil {
		return agent, log.Errore(err)
	}
	uri := fmt.Sprintf("%s/get-agent?token=%s", baseAgentUri(agent.Info.Hostname, agent.Info.Port), token)
	log.Debugf("orchestrator-agent uri: %s", uri)
	body, err := readResponse(httpGet(uri))
	if err == nil {
		err = json.Unmarshal(body, &agent.Info)
	}
	if err != nil {
		log.Errore(err)
	}
	return agent, err
}
*/

// RegisterAgent registers a new agent
func RegisterAgent(agentInfo *Info) (string, error) {
	agent := &Agent{Info: agentInfo}
	agentData, err := agent.GetAgentData()
	if err != nil {
		return "", log.Errore(fmt.Errorf("Unable to get agent data: %+v", err))
	}
	agent.Data = agentData
	err = agent.registerAgent()
	if err != nil {
		return "", log.Errore(fmt.Errorf("Unable to save agent to database: %+v", err))
	}

	// Try to discover topology instances when an agent submits
	go agent.discoverAgentInstance()

	return agentInfo.Hostname, err
}

// ReadAgents returns a list of all known agents with their data from database
func ReadAgents() ([]*Agent, error) {
	return readAgents(``, sqlutils.Args(), "")
}

// ReadAgentsInfo returns a list of all known agents without data from database
func ReadAgentsInfo() ([]*Agent, error) {
	return readAgentsInfo(``, sqlutils.Args(), "")
}

// ReadAgent returns an information about an agent and it's data from database
func ReadAgent(hostname string) (*Agent, error) {
	whereCondition := `
		WHERE
			ha.hostname = ?
		`
	res, err := readAgents(whereCondition, sqlutils.Args(hostname), "")
	if err != nil {
		return nil, err
	}
	return res[0], nil
}

// ReadAgentInfo returns an information about an agent without data from database
func ReadAgentInfo(hostname string) (*Agent, error) {
	whereCondition := `
		WHERE
			hostname = ?
		`
	res, err := readAgentsInfo(whereCondition, sqlutils.Args(hostname), "")
	if err != nil {
		return nil, err
	}
	return res[0], nil
}

// ReadOutdatedAgents returns agents that need to be updated
func ReadOutdatedAgents() ([]*Agent, error) {
	whereCondition := `
		WHERE
			IFNULL(last_checked < now() - interval ? minute, 1)
	`
	return readAgentsInfo(whereCondition, sqlutils.Args(config.Config.AgentPollMinutes), "")
}

func (agent *Agent) GetAgentData() (*Data, error) {
	agentData := &Data{}
	uri := fmt.Sprintf("%s/get-agent?token=%s", baseAgentUri(agent.Info.Hostname, agent.Info.Port), agent.Info.Token)
	log.Debugf("orchestrator-agent uri: %s", uri)
	body, err := readResponse(httpGet(uri))
	if err == nil {
		err = json.Unmarshal(body, agentData)
	}
	return agentData, err
}

// If a mysql port is available, try to discover against it
func (agent *Agent) discoverAgentInstance() error {
	instanceKey := &inst.InstanceKey{Hostname: agent.Info.Hostname, Port: agent.Info.Port}
	instance, err := inst.ReadTopologyInstance(instanceKey)
	if err != nil {
		log.Errorf("Failed to read topology for %v. err=%+v", instanceKey, err)
		return err
	}
	if instance == nil {
		log.Errorf("Failed to read topology for %v", instanceKey)
		return err
	}
	log.Infof("Discovered Agent Instance: %v", instance.Key)
	return nil
}

// UpdateAgent reads information from agent API and updates orchestrator database
func (agent *Agent) UpdateAgent() error {
	log.Debugf("Updating information for agent %s", agent.Info.Hostname)
	err := agent.updateAgentLastChecked()
	if err != nil {
		return log.Errore(fmt.Errorf("Unable to update last_checked field for agent %s: %+v", agent.Info.Hostname, err))
	}
	agentData, err := agent.GetAgentData()
	if err != nil {
		statusUpdateErr := agent.updateAgentStatus(Inactive)
		if err != nil {
			log.Errore(fmt.Errorf("Unable to update status for agent %s: %+v", agent.Info.Hostname, statusUpdateErr))
		}
		return log.Errore(fmt.Errorf("Unable to get agent data for agent %s: %+v", agent.Info.Hostname, err))
	}
	agent.Data = agentData
	err = agent.updateAgentData()
	return log.Errore(fmt.Errorf("Unable to update agent data for agent %s: %+v", agent.Info.Hostname, err))
}
