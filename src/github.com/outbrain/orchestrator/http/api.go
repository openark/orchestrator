package http

import (
	"net/http"	
	"fmt"
	"strconv"	
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/outbrain/orchestrator/inst"
	"github.com/outbrain/orchestrator/logic"
)

type HttpAPI struct{}

var API HttpAPI = HttpAPI{}


//
type APIResponseCode int


func (this *APIResponseCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.String())
}

func (this *APIResponseCode) String() string {
	switch *this {
		case ERROR: return "ERROR"
		case OK: return "OK"
	}
	return "unknown"
}

const (
	ERROR APIResponseCode = iota
	OK
)

//
type APIResponse struct {
	Code	APIResponseCode
	Message	string
	Details	interface{}
}


func (this *HttpAPI) getInstanceKey(host string, port string) (inst.InstanceKey, error) {
	instanceKey, err := inst.NewInstanceKeyFromStrings(host, port)
	return *instanceKey, err
}

func (this *HttpAPI) Instance(params martini.Params, r render.Render) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	instance, found, err := inst.ReadInstance(&instanceKey)
	if (!found) || (err != nil) {
		r.JSON(200, &APIResponse{Code:ERROR, Message: fmt.Sprintf("Cannot read instance: %+v", instanceKey),})
		return
	}
	r.JSON(200, instance)
}


func (this *HttpAPI) Discover(params martini.Params, r render.Render) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	go orchestrator.StartDiscovery(instanceKey)

	r.JSON(200, &APIResponse{Code:OK, Message: fmt.Sprintf("Instance submitted for discovery: %+v", instanceKey),})
}

func (this *HttpAPI) Forget(params martini.Params, r render.Render) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	inst.ForgetInstance(&instanceKey)

	r.JSON(200, &APIResponse{Code:OK, Message: fmt.Sprintf("Instance forgotten: %+v", instanceKey),})
}



func (this *HttpAPI) BeginMaintenance(params martini.Params, r render.Render) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	key, err := inst.BeginMaintenance(&instanceKey, params["owner"], params["reason"])
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(), Details: key,})
		return
	}

	r.JSON(200, &APIResponse{Code:OK, Message: fmt.Sprintf("Maintenance begun: %+v", instanceKey),})
}



func (this *HttpAPI) EndMaintenance(params martini.Params, r render.Render) {
	maintenanceKey, err := strconv.ParseInt(params["maintenanceKey"], 10, 0)
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	err = inst.EndMaintenance(maintenanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}

	r.JSON(200, &APIResponse{Code:OK, Message: fmt.Sprintf("Maintenance ended: %+v", maintenanceKey),})
}


func (this *HttpAPI) EndMaintenanceByInstanceKey(params martini.Params, r render.Render) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	err = inst.EndMaintenanceByInstanceKey(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}

	r.JSON(200, &APIResponse{Code:OK, Message: fmt.Sprintf("Maintenance ended: %+v", instanceKey),})
}


func (this *HttpAPI) Maintenance(params martini.Params, r render.Render) {
	instanceKeys, err := inst.ReadActiveMaintenance()
							  
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: fmt.Sprintf("%+v", err),})
		return
	}

	r.JSON(200, instanceKeys)
}


func (this *HttpAPI) MoveUp(params martini.Params, r render.Render) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	instance, err := inst.MoveUp(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}

	r.JSON(200, &APIResponse{Code:OK, Message: "Instance moved up", Details: instance})
}


func (this *HttpAPI) MoveBelow(params martini.Params, r render.Render) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	siblingKey, err := this.getInstanceKey(params["siblingHost"], params["siblingPort"])
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}
	
	instance, err := inst.MoveBelow(&instanceKey, &siblingKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: err.Error(),})
		return
	}

	r.JSON(200, &APIResponse{Code:OK, Message: fmt.Sprintf("Instance %+v moved below %+v", instanceKey, siblingKey), Details: instance})
}


func (this *HttpAPI) Cluster(params martini.Params, r render.Render) {
	instances, err := inst.ReadClusterInstances(params["clusterName"])

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: fmt.Sprintf("%+v", err),})
		return
	}

	r.JSON(200, instances)
}


func (this *HttpAPI) Clusters(params martini.Params, r render.Render) {
	clusterNames, err := inst.ReadClusters()

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: fmt.Sprintf("%+v", err),})
		return
	}

	r.JSON(200, clusterNames)
}

func (this *HttpAPI) Search(params martini.Params, r render.Render, req *http.Request) {
	searchString := params["searchString"]
	if searchString == "" {
		searchString = req.URL.Query().Get("s");
	}
	instances, err := inst.SearchInstances(searchString)

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: fmt.Sprintf("%+v", err),})
		return
	}

	r.JSON(200, instances)
}


func (this *HttpAPI) Problems(params martini.Params, r render.Render, req *http.Request) {
	instances, err := inst.ReadProblemInstances()

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: fmt.Sprintf("%+v", err),})
		return
	}

	r.JSON(200, instances)
}


func (this *HttpAPI) Audit(params martini.Params, r render.Render, req *http.Request) {
	audits, err := inst.ReadRecentAudit()

	if err != nil {
		r.JSON(200, &APIResponse{Code:ERROR, Message: fmt.Sprintf("%+v", err),})
		return
	}

	r.JSON(200, audits)
}


func (this *HttpAPI) RegisterRequests(m *martini.ClassicMartini) {
	m.Get("/api/instance/:host/:port", this.Instance) 
	m.Get("/api/discover/:host/:port", this.Discover) 
	m.Get("/api/forget/:host/:port", this.Forget) 
	m.Get("/api/move-up/:host/:port", this.MoveUp) 
	m.Get("/api/move-below/:host/:port/:siblingHost/:siblingPort", this.MoveBelow) 
	m.Get("/api/begin-maintenance/:host/:port/:owner/:reason", this.BeginMaintenance) 
	m.Get("/api/end-maintenance/:host/:port", this.EndMaintenanceByInstanceKey) 
	m.Get("/api/end-maintenance/:maintenanceKey", this.EndMaintenance)	
	m.Get("/api/maintenance", this.Maintenance) 
	m.Get("/api/cluster/:clusterName", this.Cluster) 
	m.Get("/api/clusters", this.Clusters) 
	m.Get("/api/search/:searchString", this.Search) 
	m.Get("/api/search", this.Search) 
	m.Get("/api/problems", this.Problems) 
	m.Get("/api/audit", this.Audit) 
}
