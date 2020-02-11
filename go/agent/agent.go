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
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/inst"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

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

// AuditAgentOperation creates and writes a new audit entry by given agent
func auditAgentOperation(auditType string, agent *Agent, message string) error {
	instanceKey := &inst.InstanceKey{}
	if agent != nil {
		instanceKey = &inst.InstanceKey{Hostname: agent.Info.Hostname, Port: agent.Info.MySQLPort}
	}
	return inst.AuditOperation(auditType, instanceKey, message)
}

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
	if len(res) == 0 {
		return nil, fmt.Errorf("Agent %s not found", hostname)
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
	if len(res) == 0 {
		return nil, fmt.Errorf("Agent %s not found", hostname)
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

// executeAgentCommand requests an agent to execute a command via HTTP api
func (agent *Agent) executeAgentCommand(command string, onResponse *func([]byte)) error {
	httpFunc := func(uri string) (resp *http.Response, err error) {
		return httpGet(uri)
	}
	auditAgentOperation("agent-command", agent, command)
	return executeCommandWithMethodFunc(agent.Info.Hostname, agent.Info.Port, agent.Info.Token, command, httpFunc, onResponse)
}

// executeAgentPostCommand requests an agent to execute a command via HTTP POST
func (agent *Agent) executeAgentPostCommand(hostname string, command string, content string, onResponse *func([]byte)) error {
	httpFunc := func(uri string) (resp *http.Response, err error) {
		return httpPost(uri, "text/plain", content)
	}
	auditAgentOperation("agent-command", agent, command)
	return executeCommandWithMethodFunc(agent.Info.Hostname, agent.Info.Port, agent.Info.Token, command, httpFunc, onResponse)
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

// GetAgentData gets information about MySQL\LVM from agent
func (agent *Agent) GetAgentData() (*Data, error) {
	agentData := &Data{}
	onResponse := func(body []byte) {
		json.Unmarshal(body, &agentData)
	}
	err := agent.executeAgentCommand("get-agent-data", &onResponse)
	return agentData, err
}

// Unmount unmounts the designated snapshot mount point
func (agent *Agent) Unmount() error {
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, agent.Data)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand("umount", &onResponse); err != nil {
		return err
	}
	return agent.updateAgentData()
}

// UpdateAgent reads information from agent API and updates orchestrator database
func (agent *Agent) UpdateAgent() error {
	log.Debugf("Updating information for agent %+v", agent.Info.Hostname)
	if err := agent.updateAgentLastChecked(); err != nil {
		return fmt.Errorf("Unable to update last_checked field for agent %s: %+v", agent.Info.Hostname, err)
	}
	agentData, err := agent.GetAgentData()
	if err != nil {
		statusUpdateErr := agent.updateAgentStatus(Inactive)
		if statusUpdateErr != nil {
			return fmt.Errorf("Unable to update status for agent %s: %+v", agent.Info.Hostname, statusUpdateErr)
		}
		return err
	}
	agent.Data = agentData
	return agent.updateAgentData()
}
