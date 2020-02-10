package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/github/orchestrator/go/inst"
	"github.com/openark/golib/log"
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

// httpPost is a convenience method for posting text data
func httpPost(url string, bodyType string, content string) (resp *http.Response, err error) {
	return httpClient.Post(url, bodyType, strings.NewReader(content))
}

func RelaylogContentsTail(hostname string, startCoordinates *inst.BinlogCoordinates, onResponse *func([]byte)) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("mysql-relaylog-contents-tail/%s/%d", startCoordinates.LogFile, startCoordinates.LogPos), onResponse)
}

func ApplyRelaylogContents(hostname string, content string) (Agent, error) {
	return executeAgentPostCommand(hostname, "apply-relaylog-contents", content, nil)
}

// Unmount unmounts the designated snapshot mount point
func Unmount(hostname string) (Agent, error) {
	return executeAgentCommand(hostname, "umount", nil)
}

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
func PostCopy(hostname, sourceHostname string) (Agent, error) {
	return executeAgentCommand(hostname, fmt.Sprintf("post-copy/?sourceHost=%s", sourceHostname), nil)
}

// executeAgentPostCommand requests an agent to execute a command via HTTP POST
func executeAgentPostCommand(hostname string, command string, content string, onResponse *func([]byte)) (Agent, error) {
	httpFunc := func(uri string) (resp *http.Response, err error) {
		return httpPost(uri, "text/plain", content)
	}
	return executeAgentCommandWithMethodFunc(hostname, command, httpFunc, onResponse)
}
