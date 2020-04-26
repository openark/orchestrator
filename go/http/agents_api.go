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

package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/github/orchestrator/go/agent"
	"github.com/github/orchestrator/go/attributes"
)

type HttpAgentsAPI struct {
	URLPrefix string
}

var AgentsAPI HttpAgentsAPI = HttpAgentsAPI{}

// RegisterAgent registeres an agent. It is initiated by an agent to register itself.
func (this *HttpAgentsAPI) RegisterAgent(params martini.Params, r render.Render, req *http.Request) {
	var agentInfo agent.Info
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&agentInfo)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	output, err := agent.RegisterAgent(&agentInfo)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	r.JSON(200, output)
}

// SetHostAttribute is a utility method that allows per-host key-value store.
func (this *HttpAgentsAPI) SetHostAttribute(params martini.Params, r render.Render, req *http.Request) {
	err := attributes.SetHostAttributes(params["host"], params["attrVame"], params["attrValue"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, (err == nil))
}

// GetHostAttributeByAttributeName returns a host attribute
func (this *HttpAgentsAPI) GetHostAttributeByAttributeName(params martini.Params, r render.Render, req *http.Request) {

	output, err := attributes.GetHostAttributesByAttribute(params["attr"], req.URL.Query().Get("valueMatch"))

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentsInstances provides list of assumed MySQL instances (host:port)
func (this *HttpAgentsAPI) AgentsInstances(params martini.Params, r render.Render, req *http.Request) string {
	agents, err := agent.ReadAgentsInfo()
	hostnames := []string{}
	for _, agent := range agents {
		hostnames = append(hostnames, fmt.Sprintf("%s:%d", agent.Info.Hostname, agent.Info.MySQLPort))
	}

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return ""
	}

	if req.URL.Query().Get("format") == "txt" {
		return strings.Join(hostnames, "\n")
	} else {
		r.JSON(200, hostnames)
	}
	return ""
}

func (this *HttpAgentsAPI) AgentPing(params martini.Params, r render.Render, req *http.Request) {
	r.JSON(200, "OK")
}

// RegisterRequests makes for the de-facto list of known API calls
func (this *HttpAgentsAPI) RegisterRequests(m *martini.ClassicMartini) {
	m.Post(this.URLPrefix+"/api/register-agent", this.RegisterAgent)
	m.Get(this.URLPrefix+"/api/host-attribute/:host/:attrVame/:attrValue", this.SetHostAttribute)
	m.Get(this.URLPrefix+"/api/host-attribute/attr/:attr/", this.GetHostAttributeByAttributeName)
	m.Get(this.URLPrefix+"/api/agents-instances", this.AgentsInstances)
	m.Get(this.URLPrefix+"/api/agent-ping", this.AgentPing)
}
