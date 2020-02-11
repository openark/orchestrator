package agent

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"
)

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

// httpPost is a convenience method for posting text data
func httpPost(url string, bodyType string, content string) (resp *http.Response, err error) {
	return httpClient.Post(url, bodyType, strings.NewReader(content))
}

// executeAgentCommandWithMethodFunc requests an agent to execute a command via HTTP api, either GET or POST,
// with specific http method implementation by the caller
func executeCommandWithMethodFunc(hostname string, port int, token string, command string, methodFunc httpMethodFunc, onResponse *func([]byte)) error {
	uri := baseAgentURI(hostname, port)

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
		log.Errore(err)
	}
	if onResponse != nil {
		(*onResponse)(body)
	}

	return err
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
func baseAgentURI(agentHostname string, agentPort int) string {
	protocol := "http"
	if config.Config.AgentsUseSSL {
		protocol = "https"
	}
	uri := fmt.Sprintf("%s://%s:%d/api", protocol, agentHostname, agentPort)
	log.Debugf("orchestrator-agent uri: %s", uri)
	return uri
}
