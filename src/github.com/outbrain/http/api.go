package http

import (
	"strconv"
	
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/outbrain/inst"
)

type HttpAPI struct{}

var API HttpAPI = HttpAPI{}

func (this *HttpAPI) Instance(params martini.Params, r render.Render) {
	instanceKey := inst.InstanceKey{Hostname: params["host"]}
	var err error
	if instanceKey.Port, err = strconv.Atoi(params["port"]); err != nil {
		r.HTML(500, "invalid port: %s", params["port"])
	}
	instance, found, err := inst.ReadInstance(&instanceKey)
	if instanceKey.Port, err = strconv.Atoi(params["port"]); !found {
		r.HTML(500, "Cannot read instance: %+v", instanceKey)
	}
	//r.HTML(200, "instance", map[string]interface{}{"name": params["host"], "PI":params["port"], "complex": instance})
	r.JSON(200, instance)
}
