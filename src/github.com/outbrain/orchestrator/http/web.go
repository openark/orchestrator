package http

import (
	"strconv"
	"fmt"	
	"errors"
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
	_, err := inst.ReadClusterInstances(params["clusterName"])

	if err != nil {
		return
	}

	r.HTML(200, "templates/cluster", map[string]interface{}{
		"title": "cluster", 
		"clusterName": params["clusterName"],
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
}
