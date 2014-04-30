package http

import (
	"strconv"
	"fmt"	
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/outbrain/inst"
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
}


func (this *HttpAPI) Instance(params martini.Params, r render.Render) {
	instanceKey := inst.InstanceKey{Hostname: params["host"]}
	var err error

	if instanceKey.Port, err = strconv.Atoi(params["port"]); err != nil {
		r.JSON(500, &APIResponse{Code:ERROR, Message: fmt.Sprintf("Invalid port: %s", params["port"]),})
		return
	}
	instance, found, err := inst.ReadInstance(&instanceKey)
	if instanceKey.Port, err = strconv.Atoi(params["port"]); !found || err != nil {
		r.JSON(500, &APIResponse{Code:ERROR, Message: fmt.Sprintf("Cannot read instance: %+v", instanceKey),})
		return
	}
	r.JSON(200, instance)
}


func (this *HttpAPI) RegisterRequests(m *martini.ClassicMartini) {
	m.Get("/api/instance/:host/:port", this.Instance) 
}
