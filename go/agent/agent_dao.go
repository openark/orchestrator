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
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
	"github.com/outbrain/orchestrator/go/inst"
)

var SeededAgents chan *Agent = make(chan *Agent)

var httpTimeout = time.Duration(time.Duration(config.Config.HttpTimeoutSeconds) * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, httpTimeout)
}

var httpTransport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: config.Config.SSLSkipVerify},
	Dial:            dialTimeout,
	ResponseHeaderTimeout: httpTimeout,
}
var httpClient = &http.Client{Transport: httpTransport}

// httpGet is a convenience method for getting http response from URL, optionaly skipping SSL cert verification
func httpGet(url string) (resp *http.Response, err error) {
	return httpClient.Get(url)
}

// AuditAgentOperation creates and writes a new audit entry by given agent
func auditAgentOperation(auditType string, agent *Agent, message string) error {

	instanceKey := &inst.InstanceKey{}
	if agent != nil {
		instanceKey = &inst.InstanceKey{Hostname: agent.Hostname, Port: int(agent.MySQLPort)}
	}
	return inst.AuditOperation(auditType, instanceKey, message)
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

// SubmitAgent submits a new agent for listing
func SubmitAgent(hostname string, port int, token string) (string, error) {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return "", log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
			replace 
				into host_agent (
					hostname, port, token, last_submitted
				) VALUES (
					?, ?, ?, NOW()
				)
			`,
		hostname,
		port,
		token,
	)
	if err != nil {
		return "", log.Errore(err)
	}

	return hostname, err
}

// ForgetLongUnseenAgents will remove entries of all agents that have long since been last seen.
func ForgetLongUnseenAgents() error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
			delete 
				from host_agent 
			where 
				last_submitted < NOW() - interval ? hour`,
		config.Config.UnseenAgentForgetHours,
	)
	return err
}

// ReadOutdatedAgentsHosts returns agents that need to be updated
func ReadOutdatedAgentsHosts() ([]string, error) {
	res := []string{}
	query := fmt.Sprintf(`
		select 
			hostname 
		from 
			host_agent 
		where
			IFNULL(last_checked < now() - interval %d minute, true)
			`,
		config.Config.AgentPollMinutes)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		hostname := m.GetString("hostname")
		res = append(res, hostname)
		return nil
	})
Cleanup:

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
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		agent := Agent{}
		agent.Hostname = m.GetString("hostname")
		agent.Port = m.GetInt("port")
		agent.MySQLPort = m.GetInt64("mysql_port")
		agent.Token = ""
		agent.LastSubmitted = m.GetString("last_submitted")

		res = append(res, agent)
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err

}

// readAgentBasicInfo returns the basic data for an agent directly from backend table (no agent access)
func readAgentBasicInfo(hostname string) (Agent, string, error) {
	agent := Agent{}
	token := ""
	query := fmt.Sprintf(`
		select 
			hostname,
			port,
			token,
			last_submitted,
			mysql_port
		from 
			host_agent
		where
			hostname = '%s'
		`, hostname)
	db, err := db.OpenOrchestrator()
	if err != nil {
		return agent, "", err
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		agent.Hostname = m.GetString("hostname")
		agent.Port = m.GetInt("port")
		agent.LastSubmitted = m.GetString("last_submitted")
		agent.MySQLPort = m.GetInt64("mysql_port")
		token = m.GetString("token")

		return nil
	})

	if token == "" {
		return agent, "", log.Errorf("Cannot get agent/token: %s", hostname)
	}
	return agent, token, nil
}

// UpdateAgentLastChecked updates the last_check timestamp in the orchestrator backed database
// for a given agent
func UpdateAgentLastChecked(hostname string) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
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
func UpdateAgentInfo(hostname string, agent Agent) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
        	update 
        		host_agent 
        	set
        		last_seen = NOW(),
        		mysql_port = ?,
        		count_mysql_snapshots = ?
			where 
				hostname = ?`,
		agent.MySQLPort,
		len(agent.LogicalVolumes),
		hostname,
	)
	if err != nil {
		return log.Errore(err)
	}

	return nil
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

// GetAgent gets a single agent status from the agent service. This involves multiple HTTP requests.
func GetAgent(hostname string) (Agent, error) {
	agent, token, err := readAgentBasicInfo(hostname)
	if err != nil {
		return agent, log.Errore(err)
	}

	// All seems to be in order. Now make some inquiries from orchestrator-agent service:
	{
		uri := baseAgentUri(agent.Hostname, agent.Port)
		log.Debugf("orchestrator-agent uri: %s", uri)

		{
			availableLocalSnapshotsUri := fmt.Sprintf("%s/available-snapshots-local?token=%s", uri, token)
			body, err := readResponse(httpGet(availableLocalSnapshotsUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.AvailableLocalSnapshots)
			}
			if err != nil {
				log.Errore(err)
			}
		}
		{
			availableSnapshotsUri := fmt.Sprintf("%s/available-snapshots?token=%s", uri, token)
			body, err := readResponse(httpGet(availableSnapshotsUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.AvailableSnapshots)
			}
			if err != nil {
				log.Errore(err)
			}
		}
		{
			lvSnapshotsUri := fmt.Sprintf("%s/lvs-snapshots?token=%s", uri, token)
			body, err := readResponse(httpGet(lvSnapshotsUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.LogicalVolumes)
			}
			if err != nil {
				log.Errore(err)
			}
		}
		{
			mountUri := fmt.Sprintf("%s/mount?token=%s", uri, token)
			body, err := readResponse(httpGet(mountUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.MountPoint)
			}
			if err != nil {
				log.Errore(err)
			}
		}
		{
			mySQLRunningUri := fmt.Sprintf("%s/mysql-status?token=%s", uri, token)
			body, err := readResponse(httpGet(mySQLRunningUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.MySQLRunning)
			}
			// Actually an error is OK here since "status" returns with non-zero exit code when MySQL not running
		}
		{
			mySQLRunningUri := fmt.Sprintf("%s/mysql-port?token=%s", uri, token)
			body, err := readResponse(httpGet(mySQLRunningUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.MySQLPort)
			}
			if err != nil {
				log.Errore(err)
			}
		}
		{
			mySQLDiskUsageUri := fmt.Sprintf("%s/mysql-du?token=%s", uri, token)
			body, err := readResponse(httpGet(mySQLDiskUsageUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.MySQLDiskUsage)
			}
			if err != nil {
				log.Errore(err)
			}
		}
		{
			mySQLDatadirDiskFreeUri := fmt.Sprintf("%s/mysql-datadir-available-space?token=%s", uri, token)
			body, err := readResponse(httpGet(mySQLDatadirDiskFreeUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.MySQLDatadirDiskFree)
			}
			if err != nil {
				log.Errore(err)
			}
		}
		{
			errorLogTailUri := fmt.Sprintf("%s/mysql-error-log-tail?token=%s", uri, token)
			body, err := readResponse(httpGet(errorLogTailUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.MySQLErrorLogTail)
			}
			if err != nil {
				log.Errore(err)
			}
		}
	}
	return agent, err
}

// executeAgentCommand requests an agent to execute a command via HTTP api
func executeAgentCommand(hostname string, command string, onResponse *func([]byte)) (Agent, error) {
	agent, token, err := readAgentBasicInfo(hostname)
	if err != nil {
		return agent, err
	}

	// All seems to be in order. Now make some inquiries from orchestrator-agent service:
	uri := baseAgentUri(agent.Hostname, agent.Port)

	var fullCommand string
	if strings.Contains(command, "?") {
		fullCommand = fmt.Sprintf("%s&token=%s", command, token)
	} else {
		fullCommand = fmt.Sprintf("%s?token=%s", command, token)
	}
	log.Debugf("orchestrator-agent command: %s", fullCommand)
	agentCommandUri := fmt.Sprintf("%s/%s", uri, fullCommand)

	body, err := readResponse(httpGet(agentCommandUri))
	if err != nil {
		return agent, log.Errore(err)
	}
	if onResponse != nil {
		(*onResponse)(body)
	}
	auditAgentOperation("agent-command", &agent, command)

	return agent, err
}

// Unmount unmounts the designated snapshot mount point
func Unmount(hostname string) (Agent, error) {
	return executeAgentCommand(hostname, "umount", nil)
}

// MountLV requests an agent to mount the given volume on the designated mount point
func MountLV(hostname string, lv string) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("mountlv?lv=%s", lv), nil)
}

// RemoveLV requests an agent to remvoe a snapshot
func RemoveLV(hostname string, lv string) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("removelv?lv=%s", lv), nil)
}

// CreateSnapshot requests an agent to create a new snapshot -- a DIY implementation
func CreateSnapshot(hostname string) (Agent, error) {
	return executeAgentCommand(hostname, "create-snapshot", nil)
}

// deleteMySQLDatadir requests an agent to purge the MySQL data directory (step before seed)
func deleteMySQLDatadir(hostname string) (Agent, error) {
	return executeAgentCommand(hostname, "delete-mysql-datadir", nil)
}

// MySQLStop requests an agent to stop MySQL service
func MySQLStop(hostname string) (Agent, error) {
	return executeAgentCommand(hostname, "mysql-stop", nil)
}

// MySQLStart requests an agent to start the MySQL service
func MySQLStart(hostname string) (Agent, error) {
	return executeAgentCommand(hostname, "mysql-start", nil)
}

// ReceiveMySQLSeedData requests an agent to start listening for incoming seed data
func ReceiveMySQLSeedData(hostname string, seedId int64) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("receive-mysql-seed-data/%d", seedId), nil)
}

// ReceiveMySQLSeedData requests an agent to start sending seed data
func SendMySQLSeedData(hostname string, targetHostname string, seedId int64) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("send-mysql-seed-data/%s/%d", targetHostname, seedId), nil)
}

// ReceiveMySQLSeedData requests an agent to abort seed send/receive (depending on the agent's role)
func AbortSeedCommand(hostname string, seedId int64) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("abort-seed/%d", seedId), nil)
}

// seedCommandCompleted checks an agent to see if it thinks a seed was completed.
func seedCommandCompleted(hostname string, seedId int64) (Agent, bool, error) {
	result := false
	onResponse := func(body []byte) {
		json.Unmarshal(body, &result)
	}
	agent, err := executeAgentCommand(hostname, fmt.Sprintf("seed-command-completed/%d", seedId), &onResponse)
	return agent, result, err
}

// seedCommandCompleted checks an agent to see if it thinks a seed was successful.
func seedCommandSucceeded(hostname string, seedId int64) (Agent, bool, error) {
	result := false
	onResponse := func(body []byte) {
		json.Unmarshal(body, &result)
	}
	agent, err := executeAgentCommand(hostname, fmt.Sprintf("seed-command-succeeded/%d", seedId), &onResponse)
	return agent, result, err
}

// AbortSeed will contact agents associated with a seed and request abort.
func AbortSeed(seedId int64) error {
	seedOperations, err := AgentSeedDetails(seedId)
	if err != nil {
		return log.Errore(err)
	}

	for _, seedOperation := range seedOperations {
		AbortSeedCommand(seedOperation.TargetHostname, seedId)
		AbortSeedCommand(seedOperation.SourceHostname, seedId)
	}
	updateSeedComplete(seedId, errors.New("Aborted"))
	return nil
}

// PostCopy will request an agent to invoke post-copy commands
func PostCopy(hostname string) (Agent, error) {
	return executeAgentCommand(hostname, "post-copy", nil)
}

// SubmitSeedEntry submits a new seed operation entry, returning its unique ID
func SubmitSeedEntry(targetHostname string, sourceHostname string) (int64, error) {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return 0, log.Errore(err)
	}

	res, err := sqlutils.Exec(db, `
			insert 
				into agent_seed (
					target_hostname, source_hostname, start_timestamp
				) VALUES (
					?, ?, NOW()
				)
			`,
		targetHostname,
		sourceHostname,
	)
	if err != nil {
		return 0, log.Errore(err)
	}
	id, err := res.LastInsertId()

	return id, err
}

// updateSeedComplete updates the seed entry, signing for completion
func updateSeedComplete(seedId int64, seedError error) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
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

// submitSeedStateEntry submits a seed state: a single step in the overall seed process
func submitSeedStateEntry(seedId int64, action string, errorMessage string) (int64, error) {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return 0, log.Errore(err)
	}

	res, err := sqlutils.Exec(db, `
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
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
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
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
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
func executeSeed(seedId int64, targetHostname string, sourceHostname string) error {

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
	if targetAgent.MySQLRunning {
		return updateSeedStateEntry(seedStateId, errors.New("MySQL is running on target host. Cowardly refusing to proceeed. Please stop the MySQL service"))
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Looking up available snapshots on source %s", sourceHostname), "")
	if len(sourceAgent.LogicalVolumes) == 0 {
		return updateSeedStateEntry(seedStateId, errors.New("No logical volumes found on source host"))
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Checking mount point on source %s", sourceHostname), "")
	if sourceAgent.MountPoint.IsMounted {
		return updateSeedStateEntry(seedStateId, errors.New("Volume already mounted on source host; please unmount"))
	}

	seedFromLogicalVolume := sourceAgent.LogicalVolumes[0]
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Mounting logical volume: %s", seedFromLogicalVolume.Path), "")
	_, err = MountLV(sourceHostname, seedFromLogicalVolume.Path)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}
	sourceAgent, err = GetAgent(sourceHostname)
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("MySQL data volume on source host %s is %d bytes", sourceHostname, sourceAgent.MountPoint.MySQLDiskUsage), "")

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

	if sourceAgent.MountPoint.MySQLDiskUsage > targetAgent.MySQLDatadirDiskFree {
		Unmount(sourceHostname)
		return updateSeedStateEntry(seedStateId, fmt.Errorf("Not enough disk space on target host %s. Required: %d, available: %d. Bailing out.", targetHostname, sourceAgent.MountPoint.MySQLDiskUsage, targetAgent.MySQLDatadirDiskFree))
	}

	// ...
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s will now receive data in background", targetHostname), "")
	ReceiveMySQLSeedData(targetHostname, seedId)

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Waiting some time for %s to start listening for incoming data", targetHostname), "")
	time.Sleep(2 * time.Second)

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

		if targetAgentPoll.MySQLDiskUsage == bytesCopied {
			numStaleIterations++
		}
		bytesCopied = targetAgentPoll.MySQLDiskUsage

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
		if sourceAgent.MountPoint.MySQLDiskUsage > 0 {
			copyPct = 100 * bytesCopied / sourceAgent.MountPoint.MySQLDiskUsage
		}
		seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Copied %d/%d bytes (%d%%)", bytesCopied, sourceAgent.MountPoint.MySQLDiskUsage, copyPct), "")

		if !copyComplete {
			time.Sleep(30 * time.Second)
		}
	}

	// Cleanup:
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Executing post-copy command on %s", targetHostname), "")
	_, err = PostCopy(targetHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Unmounting logical volume: %s", seedFromLogicalVolume.Path), "")
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

// Seed is the entry point for making a seed
func Seed(targetHostname string, sourceHostname string) (int64, error) {
	if targetHostname == sourceHostname {
		return 0, log.Errorf("Cannot seed %s onto itself", targetHostname)
	}
	seedId, err := SubmitSeedEntry(targetHostname, sourceHostname)
	if err != nil {
		return 0, log.Errore(err)
	}

	go func() {
		err := executeSeed(seedId, targetHostname, sourceHostname)
		updateSeedComplete(seedId, err)
	}()

	return seedId, nil
}

// readSeeds reads seed from the backend table
func readSeeds(whereCondition string, limit string) ([]SeedOperation, error) {
	res := []SeedOperation{}
	query := fmt.Sprintf(`
		select 
			agent_seed_id,
			target_hostname,
			source_hostname,
			start_timestamp,
			end_timestamp,
			is_complete,
			is_successful
		from 
			agent_seed
		%s
		order by
			agent_seed_id desc
		%s
		`, whereCondition, limit)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		seedOperation := SeedOperation{}
		seedOperation.SeedId = m.GetInt64("agent_seed_id")
		seedOperation.TargetHostname = m.GetString("target_hostname")
		seedOperation.SourceHostname = m.GetString("source_hostname")
		seedOperation.StartTimestamp = m.GetString("start_timestamp")
		seedOperation.EndTimestamp = m.GetString("end_timestamp")
		seedOperation.IsComplete = m.GetBool("is_complete")
		seedOperation.IsSuccessful = m.GetBool("is_successful")

		res = append(res, seedOperation)
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadActiveSeedsForHost reads active seeds where host participates either as source or target
func ReadActiveSeedsForHost(hostname string) ([]SeedOperation, error) {
	whereCondition := fmt.Sprintf(`
		where
			is_complete = 0
			and (
				target_hostname = '%s'
				or source_hostname = '%s'
			)
		`, hostname, hostname)
	return readSeeds(whereCondition, "")
}

// ReadRecentCompletedSeedsForHost reads active seeds where host participates either as source or target
func ReadRecentCompletedSeedsForHost(hostname string) ([]SeedOperation, error) {
	whereCondition := fmt.Sprintf(`
		where
			is_complete = 1
			and (
				target_hostname = '%s'
				or source_hostname = '%s'
			)
		`, hostname, hostname)
	return readSeeds(whereCondition, "limit 10")
}

// AgentSeedDetails reads details from backend table
func AgentSeedDetails(seedId int64) ([]SeedOperation, error) {
	whereCondition := fmt.Sprintf(`
		where
			agent_seed_id = %d
		`, seedId)
	return readSeeds(whereCondition, "")
}

// ReadRecentSeeds reads seeds from backend table.
func ReadRecentSeeds() ([]SeedOperation, error) {
	return readSeeds("", "limit 100")
}

// SeedOperationState reads states for a given seed operation
func ReadSeedStates(seedId int64) ([]SeedOperationState, error) {
	res := []SeedOperationState{}
	query := fmt.Sprintf(`
		select 
			agent_seed_state_id,
			agent_seed_id,
			state_timestamp,
			state_action,
			error_message
		from 
			agent_seed_state
		where
			agent_seed_id = %d
		order by
			agent_seed_state_id desc
		`, seedId)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		seedState := SeedOperationState{}
		seedState.SeedStateId = m.GetInt64("agent_seed_state_id")
		seedState.SeedId = m.GetInt64("agent_seed_id")
		seedState.StateTimestamp = m.GetString("state_timestamp")
		seedState.Action = m.GetString("state_action")
		seedState.ErrorMessage = m.GetString("error_message")

		res = append(res, seedState)
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err
}
