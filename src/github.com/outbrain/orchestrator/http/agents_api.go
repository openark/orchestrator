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
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	
	"github.com/outbrain/orchestrator/inst"	
)


type HttpAgentsAPI struct{}

var AgentsAPI HttpAgentsAPI = HttpAgentsAPI{}



// SubmitAgent registeres an agent on a given host
func (this *HttpAgentsAPI) SubmitAgent(params martini.Params, r render.Render) {
	output, err := inst.SubmitAgent(params["host"], params["token"])
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	r.JSON(200, output)
}

// RegisterRequests makes for the de-facto list of known API calls
func (this *HttpAgentsAPI) RegisterRequests(m *martini.ClassicMartini) {
	m.Get("/api/submit-agent/:host/:token", this.SubmitAgent) 
}
