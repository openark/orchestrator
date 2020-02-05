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
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/db"
	"github.com/github/orchestrator/go/inst"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

/***************************************************************************************************************************************/
/***************************************************************************************************************************************/
/************************************************************** NEW FUNCS **************************************************************/
/***************************************************************************************************************************************/
/***************************************************************************************************************************************/

/** TO DO **/
/** add ReadFailedSeedsForHost **/

var SeededAgents chan *Agent = make(chan *Agent)

// RegisterSeed is the entry point for making a seed. It's registers seed in db, so it can be later processed asynchronously
func RegisterSeed(seedMethodName string, targetHostname string, sourceHostname string) (int64, error) {
	var seedMethod SeedMethod
	var ok bool
	if seedMethod, ok = toSeedMethod["seedMethodName"]; !ok {
		return 0, log.Errorf("SeedMethod %s not found", seedMethodName)
	}
	if targetHostname == sourceHostname {
		return 0, log.Errorf("Cannot seed %s onto itself", targetHostname)
	}
	targetAgent, err := GetAgent(targetHostname)
	if err != nil {
		return 0, log.Errorf("Unable to get information from agent on taget host %s", targetHostname)
	}
	sourceAgent, err := GetAgent(sourceHostname)
	if err != nil {
		return 0, log.Errorf("Unable to get information from agent on source host %s", sourceHostname)
	}
	if _, ok = targetAgent.Params.AvailiableSeedMethods[seedMethod]; !ok {
		return 0, log.Errorf("Seed method %s not supported on target host %s", seedMethod.String(), targetHostname)
	}
	if _, ok = sourceAgent.Params.AvailiableSeedMethods[seedMethod]; !ok {
		return 0, log.Errorf("Seed method %s not supported on source host %s", seedMethod.String(), sourceHostname)
	}
	for db, opts := range sourceAgent.Info.MySQLDatabases {
		if !enginesSupported(opts.Engines, sourceAgent.Params.AvailiableSeedMethods[seedMethod].SupportedEngines) {
			return 0, log.Errorf("Database %s had a table with engine unsupported by seed method %s", db, seedMethod.String())
		}
	}
	if sourceAgent.Info.MySQLDatadirDiskUsed > targetAgent.Info.MySQLDatadirDiskFree {
		return 0, log.Errorf("Not enough disk space on target host %s. Required: %d, available: %d", targetHostname, sourceAgent.Info.MySQLDatadirDiskUsed, targetAgent.Info.MySQLDatadirDiskFree)
	}
	// do we need to check this? logical backup usually less than actual data size, so let's use 0.6 multiplier
	if !sourceAgent.Params.AvailiableSeedMethods[seedMethod].BackupToDatadir {
		if int64(float64(sourceAgent.Info.MySQLDatadirDiskUsed)*0.6) > targetAgent.Info.BackupDirDiskFree {
			return 0, log.Errorf("Not enough disk space on target host backup directory %s. Database size: %d, available: %d", targetHostname, sourceAgent.Info.MySQLDatadirDiskUsed, targetAgent.Info.BackupDirDiskFree)
		}
	}
	seedID, err := registerSeedEntry(seedMethod, targetHostname, sourceHostname)
	if err != nil {
		return 0, log.Errore(err)
	}
	SeededAgents <- &targetAgent

	log.Infof("Registered new seed with seedID %d. Target host: %s, source host: %s, seed method: %s", seedID, targetHostname, sourceHostname, seedMethod.String())
	auditAgentOperation("agent-seed", &sourceAgent, fmt.Sprintf("starting seed to host %s using seed method %s", targetHostname, seedMethod.String()))
	auditAgentOperation("agent-seed", &targetAgent, fmt.Sprintf("starting seed from host %s using seed method %s", sourceHostname, seedMethod.String()))
	return seedID, nil
}

// registerSeedEntry register a new seed operation entry, returning its unique ID
func registerSeedEntry(seedMethod SeedMethod, targetHostname string, sourceHostname string) (int64, error) {
	res, err := db.ExecOrchestrator(`
			insert
				into agent_seed (
					seed_method, target_hostname, source_hostname, start_timestamp, status
				) VALUES (
					?, ?, ?, NOW(), ?
				)
			`,
		seedMethod.String(),
		targetHostname,
		sourceHostname,
		Started.String(),
	)
	if err != nil {
		return 0, log.Errore(err)
	}
	id, err := res.LastInsertId()

	return id, err
}

// AuditAgentOperation creates and writes a new audit entry by given agent
func auditAgentOperation(auditType string, agent *Agent, message string) error {
	instanceKey := &inst.InstanceKey{}
	if agent != nil {
		instanceKey = &inst.InstanceKey{Hostname: agent.Params.Hostname, Port: agent.Params.MySQLPort}
	}
	return inst.AuditOperation(auditType, instanceKey, message)
}

// SubmitAgent submits a new agent for listing
func SubmitAgent(agentParams *AgentParams) (string, error) {
	seedMethods, err := json.Marshal(agentParams.AvailiableSeedMethods)
	if err != nil {
		return "", log.Errore(err)
	}
	_, err = db.ExecOrchestrator(`
			replace
				into host_agent (
					hostname, port, token, last_submitted, mysql_port, count_mysql_snapshots, seed_methods
				) VALUES (
					?, ?, ?, NOW(), ?, 0, ?
				)
			`,
		agentParams.Hostname,
		agentParams.Port,
		agentParams.Token,
		agentParams.MySQLPort,
		seedMethods,
	)
	if err != nil {
		return "", log.Errore(err)
	}

	// Try to discover topology instances when an agent submits
	go discoverAgentInstance(agentParams.Hostname, agentParams.Port)

	return agentParams.Hostname, err
}

// If a mysql port is available, try to discover against it
func discoverAgentInstance(hostname string, port int) error {
	agent, err := GetAgent(hostname)
	if err != nil {
		log.Errorf("Couldn't get agent for %s: %v", hostname, err)
		return err
	}

	instanceKey := agent.GetInstance()
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

// readSeeds reads seed from the backend table
func readSeeds(whereCondition string, args []interface{}, limit string) ([]SeedOperation, error) {
	res := []SeedOperation{}
	query := fmt.Sprintf(`
		select
			agent_seed_id,
			target_hostname,
			source_hostname,
			start_timestamp,
			end_timestamp,
			status
		from
			agent_seed
		%s
		order by
			agent_seed_id desc
		%s
		`, whereCondition, limit)
	err := db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		seedOperation := SeedOperation{}
		seedOperation.SeedId = m.GetInt64("agent_seed_id")
		seedOperation.TargetHostname = m.GetString("target_hostname")
		seedOperation.SourceHostname = m.GetString("source_hostname")
		seedOperation.StartTimestamp = m.GetTime("start_timestamp")
		seedOperation.EndTimestamp = m.GetTime("end_timestamp")
		seedOperation.Status = toStatus[m.GetString("status")]

		res = append(res, seedOperation)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadActiveSeedsForHost reads active seeds where host participates either as source or target
func ReadActiveSeedsForHost(hostname string) ([]SeedOperation, error) {
	whereCondition := `
		where
			status in ("Started", "Running")
			and (
				target_hostname = ?
				or source_hostname = ?
			)
		`
	return readSeeds(whereCondition, sqlutils.Args(hostname, hostname), "")
}

// ReadRecentCompletedSeedsForHost reads active seeds where host participates either as source or target
func ReadRecentCompletedSeedsForHost(hostname string) ([]SeedOperation, error) {
	whereCondition := `
		where
			status = "Completed"
			and (
				target_hostname = ?
				or source_hostname = ?
			)
		`
	return readSeeds(whereCondition, sqlutils.Args(hostname, hostname), "limit 10")
}

// AgentSeedDetails reads details from backend table
func AgentSeedDetails(seedId int64) ([]SeedOperation, error) {
	whereCondition := `
		where
			agent_seed_id = ?
		`
	return readSeeds(whereCondition, sqlutils.Args(seedId), "")
}

// ReadRecentSeeds reads seeds from backend table.
func ReadRecentSeeds() ([]SeedOperation, error) {
	return readSeeds(``, sqlutils.Args(), "limit 100")
}

// ForgetLongUnseenAgents will remove entries of all agents that have long since been last seen.
func ForgetLongUnseenAgents() error {
	_, err := db.ExecOrchestrator(`
			delete
				from host_agent
			where
				last_submitted < NOW() - interval ? hour`,
		config.Config.UnseenAgentForgetHours,
	)
	return err
}

// UpdateAgentLastChecked updates the last_check timestamp in the orchestrator backed database
// for a given agent
func UpdateAgentLastChecked(hostname string) error {
	_, err := db.ExecOrchestrator(`
        	update
        		host_agent
        	set
        		last_checked = NOW()
			where
				hostname = ?`,
		hostname,
	)
	if err != nil {
		return log.Errore(err)
	}

	return nil
}

// UpdateAgentInfo  updates some agent state in backend table
func UpdateAgentInfo(agent Agent) error {
	_, err := db.ExecOrchestrator(`
        	update
        		host_agent
        	set
        		last_seen = NOW(),
        		mysql_port = ?,
        		count_mysql_snapshots = ?
			where
				hostname = ?`,
		agent.Params.MySQLPort,
		len(agent.Info.LogicalVolumes),
		agent.Params.Hostname,
	)
	if err != nil {
		return log.Errore(err)
	}

	return nil
}

// readAgentBasicInfo returns the basic data for an agent directly from backend table (no agent access)
func readAgentBasicInfo(hostname string) (Agent, string, error) {
	agent := Agent{Params: &AgentParams{}, Info: &AgentInfo{}}
	token := ""
	query := `
		select
			hostname,
			port,
			token,
			last_submitted,
			mysql_port
		from
			host_agent
		where
			hostname = ?
		`
	err := db.QueryOrchestrator(query, sqlutils.Args(hostname), func(m sqlutils.RowMap) error {
		agent.Params.Hostname = m.GetString("hostname")
		agent.Params.Port = m.GetInt("port")
		agent.LastSubmitted = m.GetString("last_submitted")
		agent.Params.MySQLPort = m.GetInt("mysql_port")
		token = m.GetString("token")

		return nil
	})
	if err != nil {
		return agent, "", err
	}

	if token == "" {
		return agent, "", log.Errorf("Cannot get agent/token: %s", hostname)
	}
	return agent, token, nil
}

// ReadOutdatedAgentsHosts returns agents that need to be updated
func ReadOutdatedAgentsHosts() ([]string, error) {
	res := []string{}
	query := `
		select
			hostname
		from
			host_agent
		where
			IFNULL(last_checked < now() - interval ? minute, 1)
			`
	err := db.QueryOrchestrator(query, sqlutils.Args(config.Config.AgentPollMinutes), func(m sqlutils.RowMap) error {
		hostname := m.GetString("hostname")
		res = append(res, hostname)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadAgents returns a list of all known agents
func ReadAgents() ([]Agent, error) {
	res := []Agent{}
	query := `
		select
			hostname,
			port,
			token,
			last_submitted,
			mysql_port
		from
			host_agent
		order by
			hostname
		`
	err := db.QueryOrchestratorRowsMap(query, func(m sqlutils.RowMap) error {
		agent := Agent{Params: &AgentParams{}}
		agent.Params.Hostname = m.GetString("hostname")
		agent.Params.Port = m.GetInt("port")
		agent.Params.MySQLPort = m.GetInt("mysql_port")
		agent.Params.Token = ""
		agent.LastSubmitted = m.GetString("last_submitted")

		res = append(res, agent)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

/***************************************************************************************************************************************/
/***************************************************************************************************************************************/
/************************************************************** OLD FUNCS **************************************************************/
/***************************************************************************************************************************************/
/***************************************************************************************************************************************/

// updateSeedComplete updates the seed entry, signing for completion
func updateSeedComplete(seedId int64, seedError error) error {
	_, err := db.ExecOrchestrator(`
			update
				agent_seed
					set end_timestamp = NOW(),
					is_complete = 1,
					is_successful = ?
				where
					agent_seed_id = ?
			`,
		(seedError == nil),
		seedId,
	)
	if err != nil {
		return log.Errore(err)
	}

	return nil
}

// SeedOperationState reads states for a given seed operation
func ReadSeedStates(seedId int64) ([]SeedOperationState, error) {
	res := []SeedOperationState{}
	query := `
		select
			agent_seed_state_id,
			agent_seed_id,
			state_timestamp,
			state_action,
			error_message
		from
			agent_seed_state
		where
			agent_seed_id = ?
		order by
			agent_seed_state_id desc
		`
	err := db.QueryOrchestrator(query, sqlutils.Args(seedId), func(m sqlutils.RowMap) error {
		seedState := SeedOperationState{}
		seedState.SeedStateId = m.GetInt64("agent_seed_state_id")
		seedState.SeedId = m.GetInt64("agent_seed_id")
		seedState.StateTimestamp = m.GetString("state_timestamp")
		seedState.Action = m.GetString("state_action")
		seedState.ErrorMessage = m.GetString("error_message")

		res = append(res, seedState)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// submitSeedStateEntry submits a seed state: a single step in the overall seed process
func submitSeedStateEntry(seedId int64, action string, errorMessage string) (int64, error) {
	res, err := db.ExecOrchestrator(`
			insert
				into agent_seed_state (
					agent_seed_id, state_timestamp, state_action, error_message
				) VALUES (
					?, NOW(), ?, ?
				)
			`,
		seedId,
		action,
		errorMessage,
	)
	if err != nil {
		return 0, log.Errore(err)
	}
	id, err := res.LastInsertId()

	return id, err
}

// updateSeedStateEntry updates seed step state
func updateSeedStateEntry(seedStateId int64, reason error) error {
	_, err := db.ExecOrchestrator(`
			update
				agent_seed_state
					set error_message = ?
				where
					agent_seed_state_id = ?
			`,
		reason.Error(),
		seedStateId,
	)
	if err != nil {
		return log.Errore(err)
	}

	return reason
}

// FailStaleSeeds marks as failed seeds where no progress have been seen recently
func FailStaleSeeds() error {
	_, err := db.ExecOrchestrator(`
				update
						agent_seed
					set
						is_complete=1,
						is_successful=0
					where
						is_complete=0
						and (
							select
									max(state_timestamp) as last_state_timestamp
								from
									agent_seed_state
								where
									agent_seed.agent_seed_id = agent_seed_state.agent_seed_id
						) < now() - interval ? minute`,
		config.Config.StaleSeedFailMinutes,
	)
	return err
}

// executeSeed is *the* function for taking a seed. It is a complex operation of testing, preparing, re-testing
// agents on both sides, initiating data transfer, following up, awaiting completion, diagnosing errors, claning up.
func executeSeed(seedId int64, seedMethod SeedMethod, targetHostname string, sourceHostname string) error {

	var err error
	var seedStateId int64

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("getting target agent info for %s", targetHostname), "")
	targetAgent, err := GetAgent(targetHostname)
	SeededAgents <- &targetAgent
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("getting source agent info for %s", sourceHostname), "")
	sourceAgent, err := GetAgent(sourceHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Checking MySQL status on target %s", targetHostname), "")
	if targetAgent.Info.MySQLRunning {
		return updateSeedStateEntry(seedStateId, errors.New("MySQL is running on target host. Cowardly refusing to proceeed. Please stop the MySQL service"))
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Looking up available snapshots on source %s", sourceHostname), "")
	if len(sourceAgent.Info.LogicalVolumes) == 0 {
		return updateSeedStateEntry(seedStateId, errors.New("No logical volumes found on source host"))
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Checking mount point on source %s", sourceHostname), "")
	if sourceAgent.Info.MountPoint.IsMounted {
		return updateSeedStateEntry(seedStateId, errors.New("Volume already mounted on source host; please unmount"))
	}

	seedFromLogicalVolume := sourceAgent.Info.LogicalVolumes[0]
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s Mounting logical volume: %s", sourceHostname, seedFromLogicalVolume.Path), "")
	_, err = MountLV(sourceHostname, seedFromLogicalVolume.Path)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}
	sourceAgent, err = GetAgent(sourceHostname)
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("MySQL data volume on source host %s is %d bytes", sourceHostname, sourceAgent.Info.MySQLDatadirDiskUsed), "")

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Erasing MySQL data on %s", targetHostname), "")
	_, err = deleteMySQLDatadir(targetHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Aquiring target host datadir free space on %s", targetHostname), "")
	targetAgent, err = GetAgent(targetHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	if sourceAgent.Info.MySQLDatadirDiskUsed > targetAgent.Info.MySQLDatadirDiskFree {
		Unmount(sourceHostname)
		return updateSeedStateEntry(seedStateId, fmt.Errorf("Not enough disk space on target host %s. Required: %d, available: %d. Bailing out.", targetHostname, sourceAgent.Info.MySQLDatadirDiskUsed, targetAgent.Info.MySQLDatadirDiskFree))
	}

	// ...
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s will now receive data in background", targetHostname), "")
	ReceiveMySQLSeedData(targetHostname, seedId)

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Waiting %d seconds for %s to start listening for incoming data", config.Config.SeedWaitSecondsBeforeSend, targetHostname), "")
	time.Sleep(time.Duration(config.Config.SeedWaitSecondsBeforeSend) * time.Second)

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s will now send data to %s in background", sourceHostname, targetHostname), "")
	SendMySQLSeedData(sourceHostname, targetHostname, seedId)

	copyComplete := false
	numStaleIterations := 0
	var bytesCopied int64 = 0

	for !copyComplete {
		targetAgentPoll, err := GetAgent(targetHostname)
		if err != nil {
			return log.Errore(err)
		}

		if targetAgentPoll.Info.MySQLDatadirDiskUsed == bytesCopied {
			numStaleIterations++
		}
		bytesCopied = targetAgentPoll.Info.MySQLDatadirDiskUsed

		copyFailed := false
		if _, commandCompleted, _ := seedCommandCompleted(targetHostname, seedId); commandCompleted {
			copyComplete = true
			if _, commandSucceeded, _ := seedCommandSucceeded(targetHostname, seedId); !commandSucceeded {
				// failed.
				copyFailed = true
			}
		}
		if numStaleIterations > 10 {
			copyFailed = true
		}
		if copyFailed {
			AbortSeedCommand(sourceHostname, seedId)
			AbortSeedCommand(targetHostname, seedId)
			Unmount(sourceHostname)
			return updateSeedStateEntry(seedStateId, errors.New("10 iterations have passed without progress. Bailing out."))
		}

		var copyPct int64 = 0
		if sourceAgent.Info.MySQLDatadirDiskUsed > 0 {
			copyPct = 100 * bytesCopied / sourceAgent.Info.MySQLDatadirDiskUsed
		}
		seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Copied %d/%d bytes (%d%%)", bytesCopied, sourceAgent.Info.MySQLDatadirDiskUsed, copyPct), "")

		if !copyComplete {
			time.Sleep(30 * time.Second)
		}
	}

	// Cleanup:
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Executing post-copy command on %s", targetHostname), "")
	_, err = PostCopy(targetHostname, sourceHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s Unmounting logical volume: %s", sourceHostname, seedFromLogicalVolume.Path), "")
	_, err = Unmount(sourceHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Starting MySQL on target: %s", targetHostname), "")
	_, err = MySQLStart(targetHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Submitting MySQL instance for discovery: %s", targetHostname), "")
	SeededAgents <- &targetAgent

	seedStateId, _ = submitSeedStateEntry(seedId, "Done", "")

	return nil
}
