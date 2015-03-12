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
	"errors"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"

	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/inst"
)

// HttpWeb is the web requests server, mapping each request to a web page
type HttpWeb struct{}

var Web HttpWeb = HttpWeb{}

func (this *HttpWeb) getInstanceKey(host string, port string) (inst.InstanceKey, error) {
	instanceKey := inst.InstanceKey{Hostname: host}
	var err error

	if instanceKey.Port, err = strconv.Atoi(port); err != nil {
		return instanceKey, errors.New(fmt.Sprintf("Invalid port: %s", port))
	}
	return instanceKey, err
}

func (this *HttpWeb) Clusters(params martini.Params, r render.Render) {
	r.HTML(200, "templates/clusters", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "clusters",
		"activePage":        "cluster",
		"autoshow_problems": false,
	})
}

func (this *HttpWeb) Cluster(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/cluster", map[string]interface{}{
		"agentsHttpActive":      config.Config.ServeAgentsHttp,
		"title":                 "cluster",
		"activePage":            "cluster",
		"clusterName":           params["clusterName"],
		"autoshow_problems":     true,
		"contextMenuVisible":    true,
		"pseudoGTIDModeEnabled": (config.Config.PseudoGTIDPattern != ""),
		"authorizedForAction":   isAuthorizedForAction(req, user),
	})
}

func (this *HttpWeb) Search(params martini.Params, r render.Render, req *http.Request) {
	searchString := params["searchString"]
	if searchString == "" {
		searchString = req.URL.Query().Get("s")
	}
	r.HTML(200, "templates/search", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "search",
		"activePage":        "search",
		"searchString":      searchString,
		"autoshow_problems": false,
	})
}

func (this *HttpWeb) Discover(params martini.Params, r render.Render) {

	r.HTML(200, "templates/discover", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "discover",
		"activePage":        "discover",
		"autoshow_problems": false,
	})
}

func (this *HttpWeb) LongQueries(params martini.Params, r render.Render, req *http.Request) {
	filter := params["filter"]
	if filter == "" {
		filter = req.URL.Query().Get("filter")
	}

	r.HTML(200, "templates/long_queries", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "long queries",
		"activePage":        "queries",
		"autoshow_problems": false,
		"filter":            filter,
	})
}

func (this *HttpWeb) Audit(params martini.Params, r render.Render) {
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		page = 0
	}

	r.HTML(200, "templates/audit", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "audit",
		"activePage":        "audit",
		"autoshow_problems": false,
		"page":              page,
	})
}

func (this *HttpWeb) Agents(params martini.Params, r render.Render) {
	r.HTML(200, "templates/agents", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "agents",
		"activePage":        "agents",
		"autoshow_problems": false,
	})
}

func (this *HttpWeb) Agent(params martini.Params, r render.Render) {
	r.HTML(200, "templates/agent", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "agent",
		"activePage":        "agents",
		"autoshow_problems": false,
		"agentHost":         params["host"],
	})
}

func (this *HttpWeb) AgentSeedDetails(params martini.Params, r render.Render) {
	r.HTML(200, "templates/agent_seed_details", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "agent seed details",
		"activePage":        "agents",
		"autoshow_problems": false,
		"seedId":            params["seedId"],
	})
}

func (this *HttpWeb) Seeds(params martini.Params, r render.Render) {
	r.HTML(200, "templates/seeds", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "seeds",
		"activePage":        "agents",
		"autoshow_problems": false,
	})
}

func (this *HttpWeb) Home(params martini.Params, r render.Render) {

	r.HTML(200, "templates/home", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "home",
		"activePage":        "home",
		"autoshow_problems": false,
	})
}

func (this *HttpWeb) About(params martini.Params, r render.Render) {

	r.HTML(200, "templates/about", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "about",
		"activePage":        "home",
		"autoshow_problems": false,
	})
}

func (this *HttpWeb) KeepCalm(params martini.Params, r render.Render) {

	r.HTML(200, "templates/keep-calm", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "Keep Calm",
		"activePage":        "home",
		"autoshow_problems": false,
	})
}

func (this *HttpWeb) FAQ(params martini.Params, r render.Render) {

	r.HTML(200, "templates/faq", map[string]interface{}{
		"agentsHttpActive":  config.Config.ServeAgentsHttp,
		"title":             "FAQ",
		"activePage":        "home",
		"autoshow_problems": false,
	})
}

// RegisterRequests makes for the de-facto list of known Web calls
func (this *HttpWeb) RegisterRequests(m *martini.ClassicMartini) {
	m.Get("/", this.About)
	m.Get("/web", this.About)
	m.Get("/web/home", this.About)
	m.Get("/web/about", this.About)
	m.Get("/web/keep-calm", this.KeepCalm)
	m.Get("/web/faq", this.FAQ)
	m.Get("/web/clusters", this.Clusters)
	m.Get("/web/cluster/:clusterName", this.Cluster)
	m.Get("/web/search/:searchString", this.Search)
	m.Get("/web/search", this.Search)
	m.Get("/web/discover", this.Discover)
	m.Get("/web/long-queries", this.LongQueries)
	m.Get("/web/audit", this.Audit)
	m.Get("/web/audit/:page", this.Audit)
	m.Get("/web/agents", this.Agents)
	m.Get("/web/agent/:host", this.Agent)
	m.Get("/web/seed-details/:seedId", this.AgentSeedDetails)
	m.Get("/web/seeds", this.Seeds)
}
