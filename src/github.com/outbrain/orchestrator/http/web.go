package http

import (
	"strconv"
	"fmt"	
	"errors"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/outbrain/orchestrator/inst"
)

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


func (this *HttpWeb) Cluster(params martini.Params, r render.Render) {
	r.HTML(200, "templates/cluster", map[string]interface{}{
		"title": "cluster", 
		"clusterName": params["clusterName"],
		})
}


func (this *HttpWeb) Search(params martini.Params, r render.Render, req *http.Request) {
	searchString := params["searchString"]
	if searchString == "" {
		searchString = req.URL.Query().Get("s");
	}
	r.HTML(200, "templates/search", map[string]interface{}{
		"title": "search", 
		"searchString": searchString, 
		})
}


func (this *HttpWeb) Discover(params martini.Params, r render.Render) {

	r.HTML(200, "templates/discover", map[string]interface{}{
		"title": "discover", 
		})
}


func (this *HttpWeb) Audit(params martini.Params, r render.Render) {

	r.HTML(200, "templates/audit", map[string]interface{}{
		"title": "audit", 
		})
}


func (this *HttpWeb) Home(params martini.Params, r render.Render) {

	r.HTML(200, "templates/home", map[string]interface{}{
		"title": "home", 
		})
}


func (this *HttpWeb) RegisterRequests(m *martini.ClassicMartini) {
	m.Get("/", this.Home) 
	m.Get("/web", this.Home) 
	m.Get("/web/cluster/:clusterName", this.Cluster) 
	m.Get("/web/search/:searchString", this.Search) 
	m.Get("/web/search", this.Search) 
	m.Get("/web/discover", this.Discover) 
	m.Get("/web/audit", this.Audit) 
}
