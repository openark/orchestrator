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

//Agent describes an agent
type Agent struct {
	Info         *Info
	Data         *Data
	LastSeen     time.Time
	Status       AgentStatus
	ClusterAlias string
	MySQLVersion string
}

// Info stores basic agent information
type Info struct {
	Hostname  string
	Port      int
	Token     string
	MySQLPort int
}

// Data stores additional agent information
type Data struct {
	LocalSnapshotsHosts   []string
	SnaphostHosts         []string
	LogicalVolumes        []*LogicalVolume
	MountPoint            *Mount
	BackupDir             string
	BackupDirDiskFree     int64
	BackupDirDiskUsed     int64
	MySQLRunning          bool
	MySQLDatadir          string
	MySQLDatadirDiskUsed  int64
	MySQLDatadirDiskFree  int64
	MySQLDatabases        map[string]*MySQLDatabase
	AvailiableSeedMethods map[SeedMethod]*SeedMethodOpts
	AgentCommands         []string
	NumCPU                int
	MemTotal              int64
	OsName                string
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
	CreatedAt       time.Time
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

// SeedMethodOpts stores configuration options for SeedMethod
type SeedMethodOpts struct {
	BackupSide       SeedSide
	SupportedEngines []Engine
	BackupToDatadir  bool
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

var ToAgentStatus = map[string]AgentStatus{
	"Active":   Active,
	"Inactive": Inactive,
}

var AgentStatuses = []AgentStatus{Active, Inactive}

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
	agent := &Agent{Info: agentInfo, Data: &Data{}}
	err := agent.getAgentData()
	if err != nil {
		return "", log.Errore(fmt.Errorf("Unable to get agent data: %+v", err))
	}
	agent.Status = Active
	agent.LastSeen = time.Now()
	err = agent.registerAgent()
	if err != nil {
		return "", log.Errore(fmt.Errorf("Unable to save agent to database: %+v", err))
	}

	// Try to discover topology instances when an agent submits
	go agent.discoverAgentInstance()

	return agentInfo.Hostname, err
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

// ReadAgents returns a list of all known agents with their data from database
func ReadAgents() ([]*Agent, error) {
	return readAgents(``, sqlutils.Args(), "")
}

// ReadAgentsInfo returns a list of all known agents without data from database
func ReadAgentsInfo() ([]*Agent, error) {
	return readAgentsInfo(``, sqlutils.Args(), "")
}

// ReadOutdatedAgents returns agents that need to be updated
func ReadOutdatedAgents() ([]*Agent, error) {
	whereCondition := `
		WHERE
			IFNULL(last_checked < now() - interval ? minute, 1)
	`
	return readAgentsInfo(whereCondition, sqlutils.Args(config.Config.AgentPollMinutes), "")
}

// ReadAgentsHosts returns a list of agent hostnames
func ReadAgentsHosts(hostname string) ([]string, error) {
	if len(hostname) == 0 {
		return readAgentsHosts(``, sqlutils.Args())
	}
	return readAgentsHosts(`WHERE hostname RLIKE ?`, sqlutils.Args(fmt.Sprintf("^"+hostname)))
}

// ReadAgentsPaged returns a list of all known agents with their data from database with pagination
func ReadAgentsPaged(page int) ([]*Agent, error) {
	limit := fmt.Sprintf("limit %d offset %d", config.AuditPageSize, page*config.AuditPageSize)
	return readAgents(``, sqlutils.Args(), limit)
}

// ReadAgentsForClusterPaged returns an agents for cluster with pagination
func ReadAgentsForClusterPaged(clusterAlias string, page int) ([]*Agent, error) {
	limit := fmt.Sprintf("limit %d offset %d", config.AuditPageSize, page*config.AuditPageSize)
	whereCondition := `
		WHERE
			di.suggested_cluster_alias = ?
	`
	return readAgents(whereCondition, sqlutils.Args(clusterAlias), limit)
}

// ReadAgentsInStatusPaged returns an agents in provided status with pagination
func ReadAgentsInStatusPaged(status AgentStatus, page int) ([]*Agent, error) {
	limit := fmt.Sprintf("limit %d offset %d", config.AuditPageSize, page*config.AuditPageSize)
	whereCondition := `
		WHERE
			ha.status = ?
	`
	return readAgents(whereCondition, sqlutils.Args(status.String()), limit)
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
func (agent *Agent) executeAgentPostCommand(command string, content string, onResponse *func([]byte)) error {
	httpFunc := func(uri string) (resp *http.Response, err error) {
		return httpPost(uri, "text/plain", content)
	}
	auditAgentOperation("agent-command", agent, command)
	return executeCommandWithMethodFunc(agent.Info.Hostname, agent.Info.Port, agent.Info.Token, command, httpFunc, onResponse)
}

// If a mysql port is available, try to discover against it
func (agent *Agent) discoverAgentInstance() error {
	instanceKey := &inst.InstanceKey{Hostname: agent.Info.Hostname, Port: agent.Info.MySQLPort}
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
	log.Debugf("Updating information for agent %+v", agent.Info.Hostname)
	if err := agent.updateAgentLastChecked(); err != nil {
		return fmt.Errorf("Unable to update last_checked field for agent %s: %+v", agent.Info.Hostname, err)
	}
	err := agent.getAgentData()
	if err != nil {
		agent.Status = Inactive
		if statusUpdateErr := agent.updateAgentStatus(); statusUpdateErr != nil {
			return fmt.Errorf("Unable to update status for agent %s: %+v", agent.Info.Hostname, statusUpdateErr)
		}
	}
	return err
}

// GetAgentData gets information about MySQL\LVM from agent
func (agent *Agent) getAgentData() error {
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, agent.Data)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand("get-agent-data", &onResponse); err != nil {
		return err
	}
	agent.Status = Active
	agent.LastSeen = time.Now()
	return agent.updateAgentData()
}

func (agent *Agent) relaylogContentsTail(startCoordinates *inst.BinlogCoordinates, onResponse *func([]byte)) error {
	return agent.executeAgentCommand(fmt.Sprintf("mysql-relaylog-contents-tail/%s/%d", startCoordinates.LogFile, startCoordinates.LogPos), onResponse)
}

func (agent *Agent) applyRelaylogContents(content string) error {
	return agent.executeAgentPostCommand("apply-relaylog-contents", content, nil)
}

// prepare starts prepare stage for seed on agent
func (agent *Agent) prepare(seedID int64, seedMethod SeedMethod, seedSide SeedSide) error {
	return agent.executeAgentCommand(fmt.Sprintf("prepare/%d/%s/%s", seedID, seedMethod.String(), seedSide.String()), nil)
}

// backup starts backup stage for seed on agent
func (agent *Agent) backup(seedID int64, seedMethod SeedMethod, seedHost string, mysqlPort int) error {
	return agent.executeAgentCommand(fmt.Sprintf("backup/%d/%s/%s/%d", seedID, seedMethod.String(), seedHost, mysqlPort), nil)
}

// restore starts restore stage for seed on agent
func (agent *Agent) restore(seedID int64, seedMethod SeedMethod) error {
	return agent.executeAgentCommand(fmt.Sprintf("restore/%d/%s", seedID, seedMethod.String()), nil)
}

// cleanup starts cleanup stage for seed on agent
func (agent *Agent) cleanup(seedID int64, seedMethod SeedMethod, seedSide SeedSide) error {
	return agent.executeAgentCommand(fmt.Sprintf("cleanup/%d/%s/%s", seedID, seedMethod.String(), seedSide.String()), nil)
}

// cleanup starts cleanup stage for seed on agent
func (agent *Agent) postSeedCmd(seedID int64) error {
	return agent.executeAgentCommand(fmt.Sprintf("post-seed-cmd/%d", seedID), nil)
}

// AbortSeed stops seed on agent
func (agent *Agent) AbortSeed(seedID int64, seedStage SeedStage) error {
	return agent.executeAgentCommand(fmt.Sprintf("abort-seed-stage/%d/%s", seedID, seedStage), nil)
}

// ErrorLogTail returns tail of MySQL error log from agent
func (agent *Agent) ErrorLogTail() (output string, err error) {
	onResponse := func(body []byte) {
		output = string(body)
		log.Debugf("output: %v", output)
	}
	if err := agent.executeAgentCommand("mysql-error-log-tail", &onResponse); err != nil {
		return "", err
	}
	return output, err
}

// CustomCommand executes specified command on agent
func (agent *Agent) CustomCommand(cmd string) (output string, err error) {
	onResponse := func(body []byte) {
		output = string(body)
		log.Debugf("output: %v", output)
	}
	if err := agent.executeAgentCommand(fmt.Sprintf("custom-commands/%s", cmd), &onResponse); err != nil {
		return "", err
	}
	return output, err
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

// MountLV requests an agent to mount the given volume on the designated mount point
func (agent *Agent) MountLV(lv string) error {
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, agent.Data)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand(fmt.Sprintf("mountlv?lv=%s", lv), &onResponse); err != nil {
		return err
	}
	return agent.updateAgentData()
}

// RemoveLV requests an agent to remove a snapshot
func (agent *Agent) RemoveLV(lv string) error {
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, agent.Data)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand(fmt.Sprintf("removelv?lv=%s", lv), &onResponse); err != nil {
		return err
	}
	return agent.updateAgentData()
}

// CreateSnapshot requests an agent to create a new snapshot -- a DIY implementation
func (agent *Agent) CreateSnapshot() error {
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, agent.Data)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand("create-snapshot", &onResponse); err != nil {
		return err
	}
	return agent.updateAgentData()
}

// MySQLStart requests an agent to start the MySQL service
func (agent *Agent) MySQLStart() error {
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, agent.Data)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand("mysql-start", &onResponse); err != nil {
		return err
	}
	return agent.updateAgentData()
}

// MySQLStop requests an agent to stop MySQL service
func (agent *Agent) MySQLStop() error {
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, agent.Data)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand("mysql-stop", &onResponse); err != nil {
		return err
	}
	return agent.updateAgentData()
}

// getMetdata returns SeedMetadata for seed
func (agent *Agent) getMetadata(seedID int64, seedMethod SeedMethod) (*SeedMetadata, error) {
	seedMetadata := &SeedMetadata{}
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, seedMetadata)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand(fmt.Sprintf("get-metadata/%d/%s", seedID, seedMethod.String()), &onResponse); err != nil {
		return nil, err
	}
	return seedMetadata, nil
}

// SeedStageState gets current state for seed stage for seedID
func (agent *Agent) getSeedStageState(seedID int64, seedStage SeedStage) (*SeedStageState, error) {
	seedStageState := &SeedStageState{}
	onResponse := func(body []byte) {
		err := json.Unmarshal(body, seedStageState)
		if err != nil {
			log.Errore(err)
		}
	}
	if err := agent.executeAgentCommand(fmt.Sprintf("seed-stage-state/%d/%s", seedID, seedStage.String()), &onResponse); err != nil {
		return nil, err
	}
	return seedStageState, nil
}
