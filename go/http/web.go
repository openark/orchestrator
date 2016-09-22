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
	"expvar"
	"fmt"
	"net/http"
	"net/http/pprof"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
	"github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/exp"

	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
)

// HttpWeb is the web requests server, mapping each request to a web page
type HttpWeb struct {
	URLPrefix string
}

var Web HttpWeb = HttpWeb{}

func (this *HttpWeb) getInstanceKey(host string, port string) (inst.InstanceKey, error) {
	instanceKey := inst.InstanceKey{Hostname: host}
	var err error

	if instanceKey.Port, err = strconv.Atoi(port); err != nil {
		return instanceKey, fmt.Errorf("Invalid port: %s", port)
	}
	return instanceKey, err
}

func (this *HttpWeb) AccessToken(params martini.Params, r render.Render, req *http.Request, resp http.ResponseWriter, user auth.User) {
	publicToken := req.URL.Query().Get("publicToken")
	err := authenticateToken(publicToken, resp)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}
	r.Redirect(this.URLPrefix + "/")
}

func (this *HttpWeb) Index(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	// Redirect index so that all web URLs begin with "/web/".
	// We also redirect /web/ to /web/clusters so that
	// the Clusters page has a single canonical URL.
	r.Redirect(this.URLPrefix + "/web/clusters")
}

func (this *HttpWeb) Clusters(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/clusters", map[string]interface{}{
		"agentsHttpActive":              config.Config.ServeAgentsHttp,
		"title":                         "clusters",
		"activePage":                    "cluster",
		"autoshow_problems":             false,
		"authorizedForAction":           isAuthorizedForAction(req, user),
		"userId":                        getUserId(req, user),
		"removeTextFromHostnameDisplay": config.Config.RemoveTextFromHostnameDisplay,
		"prefix":                        this.URLPrefix,
	})
}

func (this *HttpWeb) ClustersAnalysis(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/clusters_analysis", map[string]interface{}{
		"agentsHttpActive":              config.Config.ServeAgentsHttp,
		"title":                         "clusters",
		"activePage":                    "cluster",
		"autoshow_problems":             false,
		"authorizedForAction":           isAuthorizedForAction(req, user),
		"userId":                        getUserId(req, user),
		"removeTextFromHostnameDisplay": config.Config.RemoveTextFromHostnameDisplay,
		"prefix":                        this.URLPrefix,
	})
}

func (this *HttpWeb) Cluster(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/cluster", map[string]interface{}{
		"agentsHttpActive":              config.Config.ServeAgentsHttp,
		"title":                         "cluster",
		"activePage":                    "cluster",
		"clusterName":                   params["clusterName"],
		"autoshow_problems":             true,
		"contextMenuVisible":            true,
		"pseudoGTIDModeEnabled":         (config.Config.PseudoGTIDPattern != ""),
		"authorizedForAction":           isAuthorizedForAction(req, user),
		"userId":                        getUserId(req, user),
		"removeTextFromHostnameDisplay": config.Config.RemoveTextFromHostnameDisplay,
		"compactDisplay":                req.URL.Query().Get("compact"),
		"prefix":                        this.URLPrefix,
	})
}

func (this *HttpWeb) ClusterByAlias(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	clusterName, err := inst.GetClusterByAlias(params["clusterAlias"])
	// Willing to accept the case of multiple clusters; we just present one
	if clusterName == "" && err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	params["clusterName"] = clusterName
	this.Cluster(params, r, req, user)
}

func (this *HttpWeb) ClusterByInstance(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, found, err := inst.ReadInstance(&instanceKey)
	if (!found) || (err != nil) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot read instance: %+v", instanceKey)})
		return
	}

	// Willing to accept the case of multiple clusters; we just present one
	if instance.ClusterName == "" && err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	params["clusterName"] = instance.ClusterName
	this.Cluster(params, r, req, user)
}

func (this *HttpWeb) ClusterPools(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/cluster_pools", map[string]interface{}{
		"agentsHttpActive":              config.Config.ServeAgentsHttp,
		"title":                         "cluster pools",
		"activePage":                    "cluster_pools",
		"clusterName":                   params["clusterName"],
		"autoshow_problems":             false, // because pool screen by default expands all hosts
		"contextMenuVisible":            true,
		"pseudoGTIDModeEnabled":         (config.Config.PseudoGTIDPattern != ""),
		"authorizedForAction":           isAuthorizedForAction(req, user),
		"userId":                        getUserId(req, user),
		"removeTextFromHostnameDisplay": config.Config.RemoveTextFromHostnameDisplay,
		"compactDisplay":                req.URL.Query().Get("compact"),
		"prefix":                        this.URLPrefix,
	})
}

func (this *HttpWeb) Search(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	searchString := params["searchString"]
	if searchString == "" {
		searchString = req.URL.Query().Get("s")
	}
	r.HTML(200, "templates/search", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "search",
		"activePage":          "search",
		"searchString":        searchString,
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) Discover(params martini.Params, r render.Render, req *http.Request, user auth.User) {

	r.HTML(200, "templates/discover", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "discover",
		"activePage":          "discover",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) LongQueries(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	filter := params["filter"]
	if filter == "" {
		filter = req.URL.Query().Get("filter")
	}

	r.HTML(200, "templates/long_queries", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "long queries",
		"activePage":          "queries",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"filter":              filter,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) Audit(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		page = 0
	}

	r.HTML(200, "templates/audit", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "audit",
		"activePage":          "audit",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"page":                page,
		"auditHostname":       params["host"],
		"auditPort":           params["port"],
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) AuditRecovery(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		page = 0
	}
	recoveryId, err := strconv.ParseInt(params["id"], 10, 0)
	if err != nil {
		recoveryId = 0
	}

	r.HTML(200, "templates/audit_recovery", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "audit-recovery",
		"activePage":          "audit",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"page":                page,
		"clusterName":         params["clusterName"],
		"recoveryId":          recoveryId,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) AuditFailureDetection(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	page, err := strconv.Atoi(params["page"])
	if err != nil {
		page = 0
	}
	detectionId, err := strconv.ParseInt(params["id"], 10, 0)
	if err != nil {
		detectionId = 0
	}

	r.HTML(200, "templates/audit_failure_detection", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "audit-failure-detection",
		"activePage":          "audit",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"page":                page,
		"detectionId":         detectionId,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) Agents(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/agents", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "agents",
		"activePage":          "agents",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) Agent(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/agent", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "agent",
		"activePage":          "agents",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"agentHost":           params["host"],
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) AgentSeedDetails(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/agent_seed_details", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "agent seed details",
		"activePage":          "agents",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"seedId":              params["seedId"],
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) Seeds(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	r.HTML(200, "templates/seeds", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "seeds",
		"activePage":          "agents",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) Home(params martini.Params, r render.Render, req *http.Request, user auth.User) {

	r.HTML(200, "templates/home", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "home",
		"activePage":          "home",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) About(params martini.Params, r render.Render, req *http.Request, user auth.User) {

	r.HTML(200, "templates/about", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "about",
		"activePage":          "home",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) KeepCalm(params martini.Params, r render.Render, req *http.Request, user auth.User) {

	r.HTML(200, "templates/keep-calm", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "Keep Calm",
		"activePage":          "home",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) FAQ(params martini.Params, r render.Render, req *http.Request, user auth.User) {

	r.HTML(200, "templates/faq", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "FAQ",
		"activePage":          "home",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

func (this *HttpWeb) Status(params martini.Params, r render.Render, req *http.Request, user auth.User) {

	r.HTML(200, "templates/status", map[string]interface{}{
		"agentsHttpActive":    config.Config.ServeAgentsHttp,
		"title":               "status",
		"activePage":          "home",
		"authorizedForAction": isAuthorizedForAction(req, user),
		"userId":              getUserId(req, user),
		"autoshow_problems":   false,
		"prefix":              this.URLPrefix,
	})
}

// RegisterRequests makes for the de-facto list of known Web calls
func (this *HttpWeb) RegisterRequests(m *martini.ClassicMartini) {
	m.Get(this.URLPrefix+"/web/access-token", this.AccessToken)
	m.Get(this.URLPrefix+"/", this.Index)
	m.Get(this.URLPrefix+"/web", this.Index)
	m.Get(this.URLPrefix+"/web/home", this.About)
	m.Get(this.URLPrefix+"/web/about", this.About)
	m.Get(this.URLPrefix+"/web/keep-calm", this.KeepCalm)
	m.Get(this.URLPrefix+"/web/faq", this.FAQ)
	m.Get(this.URLPrefix+"/web/status", this.Status)
	m.Get(this.URLPrefix+"/web/clusters", this.Clusters)
	m.Get(this.URLPrefix+"/web/clusters-analysis", this.ClustersAnalysis)
	m.Get(this.URLPrefix+"/web/cluster/:clusterName", this.Cluster)
	m.Get(this.URLPrefix+"/web/cluster/alias/:clusterAlias", this.ClusterByAlias)
	m.Get(this.URLPrefix+"/web/cluster/instance/:host/:port", this.ClusterByInstance)
	m.Get(this.URLPrefix+"/web/cluster-pools/:clusterName", this.ClusterPools)
	m.Get(this.URLPrefix+"/web/search/:searchString", this.Search)
	m.Get(this.URLPrefix+"/web/search", this.Search)
	m.Get(this.URLPrefix+"/web/discover", this.Discover)
	m.Get(this.URLPrefix+"/web/long-queries", this.LongQueries)
	m.Get(this.URLPrefix+"/web/audit", this.Audit)
	m.Get(this.URLPrefix+"/web/audit/:page", this.Audit)
	m.Get(this.URLPrefix+"/web/audit/instance/:host/:port", this.Audit)
	m.Get(this.URLPrefix+"/web/audit/instance/:host/:port/:page", this.Audit)
	m.Get(this.URLPrefix+"/web/audit-recovery", this.AuditRecovery)
	m.Get(this.URLPrefix+"/web/audit-recovery/:page", this.AuditRecovery)
	m.Get(this.URLPrefix+"/web/audit-recovery/id/:id", this.AuditRecovery)
	m.Get(this.URLPrefix+"/web/audit-recovery/cluster/:clusterName", this.AuditRecovery)
	m.Get(this.URLPrefix+"/web/audit-recovery/cluster/:clusterName/:page", this.AuditRecovery)
	m.Get(this.URLPrefix+"/web/audit-failure-detection", this.AuditFailureDetection)
	m.Get(this.URLPrefix+"/web/audit-failure-detection/:page", this.AuditFailureDetection)
	m.Get(this.URLPrefix+"/web/audit-failure-detection/id/:id", this.AuditFailureDetection)
	m.Get(this.URLPrefix+"/web/agents", this.Agents)
	m.Get(this.URLPrefix+"/web/agent/:host", this.Agent)
	m.Get(this.URLPrefix+"/web/seed-details/:seedId", this.AgentSeedDetails)
	m.Get(this.URLPrefix+"/web/seeds", this.Seeds)

	this.RegisterDebug(m)
}

// RegisterDebug adds handlers for /debug/vars (expvar) and /debug/pprof (net/http/pprof) support
func (this *HttpWeb) RegisterDebug(m *martini.ClassicMartini) {
	m.Get(this.URLPrefix+"/debug/vars", func(w http.ResponseWriter, r *http.Request) {
		// from expvar.go, since the expvarHandler isn't exported :(
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, "{\n")
		first := true
		expvar.Do(func(kv expvar.KeyValue) {
			if !first {
				fmt.Fprintf(w, ",\n")
			}
			first = false
			fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
		})
		fmt.Fprintf(w, "\n}\n")
	})

	// list all the /debug/ endpoints we want
	m.Get(this.URLPrefix+"/debug/pprof", pprof.Index)
	m.Get(this.URLPrefix+"/debug/pprof/cmdline", pprof.Cmdline)
	m.Get(this.URLPrefix+"/debug/pprof/profile", pprof.Profile)
	m.Get(this.URLPrefix+"/debug/pprof/symbol", pprof.Symbol)
	m.Post(this.URLPrefix+"/debug/pprof/symbol", pprof.Symbol)
	m.Get(this.URLPrefix+"/debug/pprof/block", pprof.Handler("block").ServeHTTP)
	m.Get(this.URLPrefix+"/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	m.Get(this.URLPrefix+"/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	m.Get(this.URLPrefix+"/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)

	// go-metrics
	m.Get(this.URLPrefix+"/debug/metrics", exp.ExpHandler(metrics.DefaultRegistry))
}
