package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

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

/***************************************************************************************************************************************/
/***************************************************************************************************************************************/
/************************************************************** OLD FUNCS **************************************************************/
/***************************************************************************************************************************************/
/***************************************************************************************************************************************/

func RelaylogContentsTail(hostname string, startCoordinates *inst.BinlogCoordinates, onResponse *func([]byte)) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("mysql-relaylog-contents-tail/%s/%d", startCoordinates.LogFile, startCoordinates.LogPos), onResponse)
}

func ApplyRelaylogContents(hostname string, content string) (Agent, error) {
	return executeAgentPostCommand(hostname, "apply-relaylog-contents", content, nil)
}

/*
// Unmount unmounts the designated snapshot mount point
func Unmount(hostname string) (Agent, error) {
	return executeAgentCommand(hostname, "umount", nil)
}
*/

// MountLV requests an agent to mount the given volume on the designated mount point
func MountLV(hostname string, lv string) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("mountlv?lv=%s", lv), nil)
}

// RemoveLV requests an agent to remove a snapshot
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

func CustomCommand(hostname string, cmd string) (output string, err error) {
	onResponse := func(body []byte) {
		output = string(body)
		log.Debugf("output: %v", output)
	}

	_, err = executeAgentCommand(hostname, fmt.Sprintf("custom-commands/%s", cmd), &onResponse)
	return output, err
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

// seedCommandCompleted checks an agent to see if it thinks a seed was completed.
func seedCommandCompleted(hostname string, seedId int64) (Agent, bool, error) {
	result := false
	onResponse := func(body []byte) {
		json.Unmarshal(body, &result)
	}
	agent, err := executeAgentCommand(hostname, fmt.Sprintf("seed-command-completed/%d", seedId), &onResponse)
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
func PostCopy(hostname, sourceHostname string) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("post-copy/?sourceHost=%s", sourceHostname), nil)
}

// executeAgentCommand requests an agent to execute a command via HTTP api
func executeAgentCommand(hostname string, command string, onResponse *func([]byte)) (Agent, error) {
	httpFunc := func(uri string) (resp *http.Response, err error) {
		return httpGet(uri)
	}
	return executeAgentCommandWithMethodFunc(hostname, command, httpFunc, onResponse)
}

// executeAgentPostCommand requests an agent to execute a command via HTTP POST
func executeAgentPostCommand(hostname string, command string, content string, onResponse *func([]byte)) (Agent, error) {
	httpFunc := func(uri string) (resp *http.Response, err error) {
		return httpPost(uri, "text/plain", content)
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
	uri := baseAgentURI(agent.Info.Hostname, agent.Info.Port)

	var fullCommand string
	if strings.Contains(command, "?") {
		fullCommand = fmt.Sprintf("%s&token=%s", command, token)
	} else {
		fullCommand = fmt.Sprintf("%s?token=%s", command, token)
	}
	log.Debugf("orchestrator-agent command: %s", fullCommand)
	agentCommandUri := fmt.Sprintf("%s/%s", uri, fullCommand)

	body, err := readResponse(methodFunc(agentCommandUri))
	if err != nil {
		return agent, log.Errore(err)
	}
	if onResponse != nil {
		(*onResponse)(body)
	}
	auditAgentOperation("agent-command", &agent, command)

	return agent, err
}

// readAgentBasicInfo returns the basic data for an agent directly from backend table (no agent access)
func readAgentBasicInfo(hostname string) (Agent, string, error) {
	agent := Agent{Info: &Info{}}
	token := ""
	query := `
		select
			hostname,
			port,
			token,
			last_seen,
			mysql_port
		from
			host_agent
		where
			hostname = ?
		`
	err := db.QueryOrchestrator(query, sqlutils.Args(hostname), func(m sqlutils.RowMap) error {
		agent.Info.Hostname = m.GetString("hostname")
		agent.Info.Port = m.GetInt("port")
		agent.LastSeen = m.GetTime("last_seen")
		agent.Info.MySQLPort = m.GetInt("mysql_port")
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
