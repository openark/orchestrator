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
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"

	"github.com/openark/golib/log"
	"github.com/openark/golib/util"

	"github.com/github/orchestrator/go/agent"
	"github.com/github/orchestrator/go/collection"
	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/discovery"
	"github.com/github/orchestrator/go/inst"
	"github.com/github/orchestrator/go/logic"
	"github.com/github/orchestrator/go/metrics/query"
	"github.com/github/orchestrator/go/process"
	"github.com/github/orchestrator/go/raft"
)

// APIResponseCode is an OK/ERROR response code
type APIResponseCode int

const (
	ERROR APIResponseCode = iota
	OK
)

var apiSynonyms = map[string]string{
	"relocate-slaves":            "relocate-replicas",
	"regroup-slaves":             "regroup-replicas",
	"move-up-slaves":             "move-up-replicas",
	"repoint-slaves":             "repoint-replicas",
	"enslave-siblings":           "take-siblings",
	"enslave-master":             "take-master",
	"regroup-slaves-bls":         "regroup-replicas-bls",
	"move-slaves-gtid":           "move-replicas-gtid",
	"regroup-slaves-gtid":        "regroup-replicas-gtid",
	"match-slaves":               "match-replicas",
	"match-up-slaves":            "match-up-replicas",
	"regroup-slaves-pgtid":       "regroup-replicas-pgtid",
	"detach-slave":               "detach-replica",
	"reattach-slave":             "reattach-replica",
	"reattach-slave-master-host": "reattach-replica-master-host",
	"cluster-osc-slaves":         "cluster-osc-replicas",
}

var registeredPaths = []string{}
var emptyInstanceKey inst.InstanceKey

func (this *APIResponseCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.String())
}

func (this *APIResponseCode) String() string {
	switch *this {
	case ERROR:
		return "ERROR"
	case OK:
		return "OK"
	}
	return "unknown"
}

// HttpStatus returns the respective HTTP status for this response
func (this *APIResponseCode) HttpStatus() int {
	switch *this {
	case ERROR:
		return http.StatusInternalServerError
	case OK:
		return http.StatusOK
	}
	return http.StatusNotImplemented
}

// APIResponse is a response returned as JSON to various requests.
type APIResponse struct {
	Code    APIResponseCode
	Message string
	Details interface{}
}

func Respond(r render.Render, apiResponse *APIResponse) {
	r.JSON(apiResponse.Code.HttpStatus(), apiResponse)
}

type HttpAPI struct {
	URLPrefix string
}

var API HttpAPI = HttpAPI{}
var discoveryMetrics = collection.CreateOrReturnCollection("DISCOVERY_METRICS")
var queryMetrics = collection.CreateOrReturnCollection("BACKEND_WRITES")

func (this *HttpAPI) getInstanceKey(host string, port string) (inst.InstanceKey, error) {
	instanceKey, err := inst.NewInstanceKeyFromStrings(host, port)
	if err != nil {
		return emptyInstanceKey, err
	}
	instanceKey, err = inst.FigureInstanceKey(instanceKey, nil)
	if err != nil {
		return emptyInstanceKey, err
	}
	return *instanceKey, nil
}

func (this *HttpAPI) getBinlogCoordinates(logFile string, logPos string) (inst.BinlogCoordinates, error) {
	coordinates := inst.BinlogCoordinates{LogFile: logFile}
	var err error
	if coordinates.LogPos, err = strconv.ParseInt(logPos, 10, 0); err != nil {
		return coordinates, fmt.Errorf("Invalid logPos: %s", logPos)
	}

	return coordinates, err
}

// InstanceReplicas lists all replicas of given instance
func (this *HttpAPI) InstanceReplicas(params martini.Params, r render.Render, req *http.Request) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	replicas, err := inst.ReadReplicaInstances(&instanceKey)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot read instance: %+v", instanceKey)})
		return
	}
	r.JSON(http.StatusOK, replicas)
}

// Instance reads and returns an instance's details.
func (this *HttpAPI) Instance(params martini.Params, r render.Render, req *http.Request) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, found, err := inst.ReadInstance(&instanceKey)
	if (!found) || (err != nil) {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot read instance: %+v", instanceKey)})
		return
	}
	r.JSON(http.StatusOK, instance)
}

// AsyncDiscover issues an asynchronous read on an instance. This is
// useful for bulk loads of a new set of instances and will not block
// if the instance is slow to respond or not reachable.
func (this *HttpAPI) AsyncDiscover(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	go this.Discover(params, r, req, user)

	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Asynchronous discovery initiated for Instance: %+v", instanceKey)})
}

// Discover issues a synchronous read on an instance
func (this *HttpAPI) Discover(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.ReadTopologyInstance(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	go orcraft.PublishCommand("discover", instanceKey)

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance discovered: %+v", instance.Key), Details: instance})
}

// Refresh synchronuously re-reads a topology instance
func (this *HttpAPI) Refresh(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	_, err = inst.RefreshTopologyInstance(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance refreshed: %+v", instanceKey), Details: instanceKey})
}

// Forget removes an instance entry fro backend database
func (this *HttpAPI) Forget(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	// We ignore errors: we're looking to do a destructive operation anyhow.
	rawInstanceKey, _ := inst.NewRawInstanceKey(fmt.Sprintf("%s:%s", params["host"], params["port"]))

	if orcraft.IsRaftEnabled() {
		orcraft.PublishCommand("forget", rawInstanceKey)
	} else {
		inst.ForgetInstance(rawInstanceKey)
	}
	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance forgotten: %+v", *rawInstanceKey)})
}

// Resolve tries to resolve hostname and then checks to see if port is open on that host.
func (this *HttpAPI) Resolve(params martini.Params, r render.Render, req *http.Request) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	if conn, err := net.Dial("tcp", instanceKey.DisplayString()); err == nil {
		conn.Close()
	} else {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "Instance resolved", Details: instanceKey})
}

// BeginMaintenance begins maintenance mode for given instance
func (this *HttpAPI) BeginMaintenance(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	key, err := inst.BeginBoundedMaintenance(&instanceKey, params["owner"], params["reason"], 0, true)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error(), Details: key})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Maintenance begun: %+v", instanceKey), Details: instanceKey})
}

// EndMaintenance terminates maintenance mode
func (this *HttpAPI) EndMaintenance(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	maintenanceKey, err := strconv.ParseInt(params["maintenanceKey"], 10, 0)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	_, err = inst.EndMaintenance(maintenanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Maintenance ended: %+v", maintenanceKey), Details: maintenanceKey})
}

// EndMaintenanceByInstanceKey terminates maintenance mode for given instance
func (this *HttpAPI) EndMaintenanceByInstanceKey(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	_, err = inst.EndMaintenanceByInstanceKey(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Maintenance ended: %+v", instanceKey), Details: instanceKey})
}

// Maintenance provides list of instance under active maintenance
func (this *HttpAPI) Maintenance(params martini.Params, r render.Render, req *http.Request) {
	instanceKeys, err := inst.ReadActiveMaintenance()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instanceKeys)
}

// BeginDowntime sets a downtime flag with default duration
func (this *HttpAPI) BeginDowntime(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	var durationSeconds int = 0
	if params["duration"] != "" {
		durationSeconds, err = util.SimpleTimeToSeconds(params["duration"])
		if durationSeconds < 0 {
			err = fmt.Errorf("Duration value must be non-negative. Given value: %d", durationSeconds)
		}
		if err != nil {
			Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
			return
		}
	}
	duration := time.Duration(durationSeconds) * time.Second
	downtime := inst.NewDowntime(&instanceKey, params["owner"], params["reason"], duration)
	if orcraft.IsRaftEnabled() {
		_, err = orcraft.PublishCommand("begin-downtime", downtime)
	} else {
		err = inst.BeginDowntime(downtime)
	}

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error(), Details: instanceKey})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Downtime begun: %+v", instanceKey), Details: instanceKey})
}

// EndDowntime terminates downtime (removes downtime flag) for an instance
func (this *HttpAPI) EndDowntime(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	if orcraft.IsRaftEnabled() {
		_, err = orcraft.PublishCommand("end-downtime", instanceKey)
	} else {
		_, err = inst.EndDowntime(&instanceKey)
	}
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Downtime ended: %+v", instanceKey), Details: instanceKey})
}

// MoveUp attempts to move an instance up the topology
func (this *HttpAPI) MoveUp(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.MoveUp(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v moved up", instanceKey), Details: instance})
}

// MoveUpReplicas attempts to move up all replicas of an instance
func (this *HttpAPI) MoveUpReplicas(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	replicas, newMaster, err, errs := inst.MoveUpReplicas(&instanceKey, req.URL.Query().Get("pattern"))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Moved up %d replicas of %+v below %+v; %d errors: %+v", len(replicas), instanceKey, newMaster.Key, len(errs), errs), Details: replicas})
}

// Repoint positiones a replica under another (or same) master with exact same coordinates.
// Useful for binlog servers
func (this *HttpAPI) Repoint(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.Repoint(&instanceKey, &belowKey, inst.GTIDHintNeutral)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v repointed below %+v", instanceKey, belowKey), Details: instance})
}

// MoveUpReplicas attempts to move up all replicas of an instance
func (this *HttpAPI) RepointReplicas(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	replicas, err, _ := inst.RepointReplicas(&instanceKey, req.URL.Query().Get("pattern"))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Repointed %d replicas of %+v", len(replicas), instanceKey), Details: replicas})
}

// MakeCoMaster attempts to make an instance co-master with its own master
func (this *HttpAPI) MakeCoMaster(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.MakeCoMaster(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance made co-master: %+v", instance.Key), Details: instance})
}

// ResetSlave makes a replica forget about its master, effectively breaking the replication
func (this *HttpAPI) ResetSlave(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.ResetSlaveOperation(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Replica reset on %+v", instance.Key), Details: instance})
}

// DetachReplica corrupts a replica's binlog corrdinates (though encodes it in such way
// that is reversible), effectively breaking replication
func (this *HttpAPI) DetachReplica(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.DetachReplicaOperation(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Replica detached: %+v", instance.Key), Details: instance})
}

// ReattachReplica reverts a DetachReplica commands by reassigning the correct
// binlog coordinates to an instance
func (this *HttpAPI) ReattachReplica(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.ReattachReplicaOperation(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Replica reattached: %+v", instance.Key), Details: instance})
}

// ReattachReplicaMasterHost reverts a achReplicaMasterHost command
// by resoting the original master hostname in CHANGE MASTER TO
func (this *HttpAPI) ReattachReplicaMasterHost(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.ReattachReplicaMasterHost(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Replica reattached: %+v", instance.Key), Details: instance})
}

// EnableGTID attempts to enable GTID on a replica
func (this *HttpAPI) EnableGTID(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.EnableGTID(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Enabled GTID on %+v", instance.Key), Details: instance})
}

// DisableGTID attempts to disable GTID on a replica, and revert to binlog file:pos
func (this *HttpAPI) DisableGTID(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.DisableGTID(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Disabled GTID on %+v", instance.Key), Details: instance})
}

// MoveBelow attempts to move an instance below its supposed sibling
func (this *HttpAPI) MoveBelow(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	siblingKey, err := this.getInstanceKey(params["siblingHost"], params["siblingPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.MoveBelow(&instanceKey, &siblingKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v moved below %+v", instanceKey, siblingKey), Details: instance})
}

// MoveBelowGTID attempts to move an instance below another, via GTID
func (this *HttpAPI) MoveBelowGTID(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.MoveBelowGTID(&instanceKey, &belowKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v moved below %+v via GTID", instanceKey, belowKey), Details: instance})
}

// MoveReplicasGTID attempts to move an instance below another, via GTID
func (this *HttpAPI) MoveReplicasGTID(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	movedReplicas, _, err, errs := inst.MoveReplicasGTID(&instanceKey, &belowKey, req.URL.Query().Get("pattern"))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Moved %d replicas of %+v below %+v via GTID; %d errors: %+v", len(movedReplicas), instanceKey, belowKey, len(errs), errs), Details: belowKey})
}

// TakeSiblings
func (this *HttpAPI) TakeSiblings(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, count, err := inst.TakeSiblings(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Took %d siblings of %+v", count, instanceKey), Details: instance})
}

// TakeMaster
func (this *HttpAPI) TakeMaster(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.TakeMaster(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("%+v took its master", instanceKey), Details: instance})
}

// RelocateBelow attempts to move an instance below another, orchestrator choosing the best (potentially multi-step)
// relocation method
func (this *HttpAPI) RelocateBelow(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.RelocateBelow(&instanceKey, &belowKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v relocated below %+v", instanceKey, belowKey), Details: instance})
}

// Relocates attempts to smartly relocate replicas of a given instance below another
func (this *HttpAPI) RelocateReplicas(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	replicas, _, err, errs := inst.RelocateReplicas(&instanceKey, &belowKey, req.URL.Query().Get("pattern"))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Relocated %d replicas of %+v below %+v; %d errors: %+v", len(replicas), instanceKey, belowKey, len(errs), errs), Details: replicas})
}

// MoveEquivalent attempts to move an instance below another, baseed on known equivalence master coordinates
func (this *HttpAPI) MoveEquivalent(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.MoveEquivalent(&instanceKey, &belowKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v relocated via equivalence coordinates below %+v", instanceKey, belowKey), Details: instance})
}

// LastPseudoGTID attempts to find the last pseugo-gtid entry in an instance
func (this *HttpAPI) LastPseudoGTID(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.ReadTopologyInstance(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	if instance == nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Instance not found: %+v", instanceKey)})
		return
	}
	coordinates, text, err := inst.FindLastPseudoGTIDEntry(instance, instance.RelaylogCoordinates, nil, false, nil)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("%+v", *coordinates), Details: text})
}

// MatchBelow attempts to move an instance below another via pseudo GTID matching of binlog entries
func (this *HttpAPI) MatchBelow(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, matchedCoordinates, err := inst.MatchBelow(&instanceKey, &belowKey, true)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v matched below %+v at %+v", instanceKey, belowKey, *matchedCoordinates), Details: instance})
}

// MatchBelow attempts to move an instance below another via pseudo GTID matching of binlog entries
func (this *HttpAPI) MatchUp(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, matchedCoordinates, err := inst.MatchUp(&instanceKey, true)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v matched up at %+v", instanceKey, *matchedCoordinates), Details: instance})
}

// MultiMatchReplicas attempts to match all replicas of a given instance below another, efficiently
func (this *HttpAPI) MultiMatchReplicas(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	replicas, newMaster, err, errs := inst.MultiMatchReplicas(&instanceKey, &belowKey, req.URL.Query().Get("pattern"))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Matched %d replicas of %+v below %+v; %d errors: %+v", len(replicas), instanceKey, newMaster.Key, len(errs), errs), Details: newMaster.Key})
}

// MatchUpReplicas attempts to match up all replicas of an instance
func (this *HttpAPI) MatchUpReplicas(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	replicas, newMaster, err, errs := inst.MatchUpReplicas(&instanceKey, req.URL.Query().Get("pattern"))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Matched up %d replicas of %+v below %+v; %d errors: %+v", len(replicas), instanceKey, newMaster.Key, len(errs), errs), Details: newMaster.Key})
}

// RegroupReplicas attempts to pick a replica of a given instance and make it take its siblings, using any
// method possible (GTID, Pseudo-GTID, binlog servers)
func (this *HttpAPI) RegroupReplicas(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	lostReplicas, equalReplicas, aheadReplicas, cannotReplicateReplicas, promotedReplica, err := inst.RegroupReplicas(&instanceKey, false, nil, nil)
	lostReplicas = append(lostReplicas, cannotReplicateReplicas...)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("promoted replica: %s, lost: %d, trivial: %d, pseudo-gtid: %d",
		promotedReplica.Key.DisplayString(), len(lostReplicas), len(equalReplicas), len(aheadReplicas)), Details: promotedReplica.Key})
}

// RegroupReplicas attempts to pick a replica of a given instance and make it take its siblings, efficiently,
// using pseudo-gtid if necessary
func (this *HttpAPI) RegroupReplicasPseudoGTID(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	lostReplicas, equalReplicas, aheadReplicas, cannotReplicateReplicas, promotedReplica, err := inst.RegroupReplicasPseudoGTID(&instanceKey, false, nil, nil)
	lostReplicas = append(lostReplicas, cannotReplicateReplicas...)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("promoted replica: %s, lost: %d, trivial: %d, pseudo-gtid: %d",
		promotedReplica.Key.DisplayString(), len(lostReplicas), len(equalReplicas), len(aheadReplicas)), Details: promotedReplica.Key})
}

// RegroupReplicasGTID attempts to pick a replica of a given instance and make it take its siblings, efficiently, using GTID
func (this *HttpAPI) RegroupReplicasGTID(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	lostReplicas, movedReplicas, cannotReplicateReplicas, promotedReplica, err := inst.RegroupReplicasGTID(&instanceKey, false, nil)
	lostReplicas = append(lostReplicas, cannotReplicateReplicas...)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("promoted replica: %s, lost: %d, moved: %d",
		promotedReplica.Key.DisplayString(), len(lostReplicas), len(movedReplicas)), Details: promotedReplica.Key})
}

// RegroupReplicasBinlogServers attempts to pick a replica of a given instance and make it take its siblings, efficiently, using GTID
func (this *HttpAPI) RegroupReplicasBinlogServers(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	_, promotedBinlogServer, err := inst.RegroupReplicasBinlogServers(&instanceKey, false)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("promoted binlog server: %s",
		promotedBinlogServer.Key.DisplayString()), Details: promotedBinlogServer.Key})
}

// MakeMaster attempts to make the given instance a master, and match its siblings to be its replicas
func (this *HttpAPI) MakeMaster(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.MakeMaster(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v now made master", instanceKey), Details: instance})
}

// MakeLocalMaster attempts to make the given instance a local master: take over its master by
// enslaving its siblings and replicating from its grandparent.
func (this *HttpAPI) MakeLocalMaster(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.MakeLocalMaster(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v now made local master", instanceKey), Details: instance})
}

// SkipQuery skips a single query on a failed replication instance
func (this *HttpAPI) SkipQuery(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.SkipQuery(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Query skipped on %+v", instance.Key), Details: instance})
}

// StartSlave starts replication on given instance
func (this *HttpAPI) StartSlave(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.StartSlave(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Replica started: %+v", instance.Key), Details: instance})
}

// RestartSlave stops & starts replication on given instance
func (this *HttpAPI) RestartSlave(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.RestartSlave(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Replica restarted: %+v", instance.Key), Details: instance})
}

// StopSlave stops replication on given instance
func (this *HttpAPI) StopSlave(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.StopSlave(&instanceKey)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Replica stopped: %+v", instance.Key), Details: instance})
}

// StopSlaveNicely stops replication on given instance, such that sql thead is aligned with IO thread
func (this *HttpAPI) StopSlaveNicely(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.StopSlaveNicely(&instanceKey, 0)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Replica stopped nicely: %+v", instance.Key), Details: instance})
}

// FlushBinaryLogs runs a single FLUSH BINARY LOGS
func (this *HttpAPI) FlushBinaryLogs(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.FlushBinaryLogs(&instanceKey, 1)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Binary logs flushed on: %+v", instance.Key), Details: instance})
}

// MasterEquivalent provides (possibly empty) list of master coordinates equivalent to the given ones
func (this *HttpAPI) MasterEquivalent(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	coordinates, err := this.getBinlogCoordinates(params["logFile"], params["logPos"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instanceCoordinates := &inst.InstanceBinlogCoordinates{Key: instanceKey, Coordinates: coordinates}

	equivalentCoordinates, err := inst.GetEquivalentMasterCoordinates(instanceCoordinates)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Found %+v equivalent coordinates", len(equivalentCoordinates)), Details: equivalentCoordinates})
}

// SetReadOnly sets the global read_only variable
func (this *HttpAPI) SetReadOnly(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.SetReadOnly(&instanceKey, true)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "Server set as read-only", Details: instance})
}

// SetWriteable clear the global read_only variable
func (this *HttpAPI) SetWriteable(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.SetReadOnly(&instanceKey, false)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "Server set as writeable", Details: instance})
}

// KillQuery kills a query running on a server
func (this *HttpAPI) KillQuery(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	processId, err := strconv.ParseInt(params["process"], 10, 0)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.KillQuery(&instanceKey, processId)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Query killed on : %+v", instance.Key), Details: instance})
}

// AsciiTopology returns an ascii graph of cluster's instances
func (this *HttpAPI) AsciiTopology(params martini.Params, r render.Render, req *http.Request) {
	clusterName, err := figureClusterName(params["clusterHint"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	asciiOutput, err := inst.ASCIITopology(clusterName, "")
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Topology for cluster %s", clusterName), Details: asciiOutput})
}

// Cluster provides list of instances in given cluster
func (this *HttpAPI) Cluster(params martini.Params, r render.Render, req *http.Request) {
	clusterName, err := figureClusterName(params["clusterHint"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	instances, err := inst.ReadClusterInstances(clusterName)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instances)
}

// ClusterByAlias provides list of instances in given cluster
func (this *HttpAPI) ClusterByAlias(params martini.Params, r render.Render, req *http.Request) {
	clusterName, err := inst.GetClusterByAlias(params["clusterAlias"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	params["clusterName"] = clusterName
	this.Cluster(params, r, req)
}

// ClusterByInstance provides list of instances in cluster an instance belongs to
func (this *HttpAPI) ClusterByInstance(params martini.Params, r render.Render, req *http.Request) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, found, err := inst.ReadInstance(&instanceKey)
	if (!found) || (err != nil) {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot read instance: %+v", instanceKey)})
		return
	}

	params["clusterName"] = instance.ClusterName
	this.Cluster(params, r, req)
}

// ClusterInfo provides details of a given cluster
func (this *HttpAPI) ClusterInfo(params martini.Params, r render.Render, req *http.Request) {
	clusterName, err := figureClusterName(params["clusterHint"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}
	clusterInfo, err := inst.ReadClusterInfo(clusterName)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, clusterInfo)
}

// Cluster provides list of instances in given cluster
func (this *HttpAPI) ClusterInfoByAlias(params martini.Params, r render.Render, req *http.Request) {
	clusterName, err := inst.GetClusterByAlias(params["clusterAlias"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	params["clusterName"] = clusterName
	this.ClusterInfo(params, r, req)
}

// ClusterOSCReplicas returns heuristic list of OSC replicas
func (this *HttpAPI) ClusterOSCReplicas(params martini.Params, r render.Render, req *http.Request) {
	clusterName, err := figureClusterName(params["clusterHint"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	instances, err := inst.GetClusterOSCReplicas(clusterName)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instances)
}

// SetClusterAlias will change an alias for a given clustername
func (this *HttpAPI) SetClusterAliasManualOverride(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	clusterName := params["clusterName"]
	alias := req.URL.Query().Get("alias")

	err := inst.SetClusterAliasManualOverride(clusterName, alias)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}
	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Cluster %s now has alias '%s'", clusterName, alias)})
}

// Clusters provides list of known clusters
func (this *HttpAPI) Clusters(params martini.Params, r render.Render, req *http.Request) {
	clusterNames, err := inst.ReadClusters()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, clusterNames)
}

// ClustersInfo provides list of known clusters, along with some added metadata per cluster
func (this *HttpAPI) ClustersInfo(params martini.Params, r render.Render, req *http.Request) {
	clustersInfo, err := inst.ReadClustersInfo("")

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, clustersInfo)
}

// Clusters provides list of known masters
func (this *HttpAPI) Masters(params martini.Params, r render.Render, req *http.Request) {
	instances, err := inst.ReadWriteableClustersMasters()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instances)
}

// Downtimed lists downtimed instances, potentially filtered by cluster
func (this *HttpAPI) Downtimed(params martini.Params, r render.Render, req *http.Request) {
	clusterName := ""
	if clusterHint := getClusterHint(params); clusterHint != "" {
		var err error
		clusterName, err = figureClusterName(clusterHint)
		if err != nil {
			Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
			return
		}
	}

	instances, err := inst.ReadDowntimedInstances(clusterName)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instances)
}

// AllInstances lists all known instances
func (this *HttpAPI) AllInstances(params martini.Params, r render.Render, req *http.Request) {
	instances, err := inst.SearchInstances("")

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instances)
}

// Search provides list of instances matching given search param via various criteria.
func (this *HttpAPI) Search(params martini.Params, r render.Render, req *http.Request) {
	searchString := params["searchString"]
	if searchString == "" {
		searchString = req.URL.Query().Get("s")
	}
	instances, err := inst.SearchInstances(searchString)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instances)
}

// Problems provides list of instances with known problems
func (this *HttpAPI) Problems(params martini.Params, r render.Render, req *http.Request) {
	clusterName := params["clusterName"]
	instances, err := inst.ReadProblemInstances(clusterName)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instances)
}

// Audit provides list of audit entries by given page number
func (this *HttpAPI) Audit(params martini.Params, r render.Render, req *http.Request) {
	page, err := strconv.Atoi(params["page"])
	if err != nil || page < 0 {
		page = 0
	}
	var auditedInstanceKey *inst.InstanceKey
	if instanceKey, err := this.getInstanceKey(params["host"], params["port"]); err == nil {
		auditedInstanceKey = &instanceKey
	}

	audits, err := inst.ReadRecentAudit(auditedInstanceKey, page)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, audits)
}

// LongQueries lists queries running for a long time, on all instances, optionally filtered by
// arbitrary text
func (this *HttpAPI) LongQueries(params martini.Params, r render.Render, req *http.Request) {
	longQueries, err := inst.ReadLongRunningProcesses(params["filter"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, longQueries)
}

// HostnameResolveCache shows content of in-memory hostname cache
func (this *HttpAPI) HostnameResolveCache(params martini.Params, r render.Render, req *http.Request) {
	content, err := inst.HostnameResolveCache()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "Cache retrieved", Details: content})
}

// ResetHostnameResolveCache clears in-memory hostname resovle cache
func (this *HttpAPI) ResetHostnameResolveCache(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	err := inst.ResetHostnameResolveCache()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "Hostname cache cleared"})
}

// DeregisterHostnameUnresolve deregisters the unresolve name used previously
func (this *HttpAPI) DeregisterHostnameUnresolve(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}

	var instanceKey *inst.InstanceKey
	if instKey, err := this.getInstanceKey(params["host"], params["port"]); err == nil {
		instanceKey = &instKey
	}

	var err error
	registration := inst.NewHostnameDeregistration(instanceKey)
	if orcraft.IsRaftEnabled() {
		_, err = orcraft.PublishCommand("register-hostname-unresolve", registration)
	} else {
		err = inst.RegisterHostnameUnresolve(registration)
	}
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}
	Respond(r, &APIResponse{Code: OK, Message: "Hostname deregister unresolve completed", Details: instanceKey})
}

// RegisterHostnameUnresolve registers the unresolve name to use
func (this *HttpAPI) RegisterHostnameUnresolve(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}

	var instanceKey *inst.InstanceKey
	if instKey, err := this.getInstanceKey(params["host"], params["port"]); err == nil {
		instanceKey = &instKey
	}

	hostname := params["virtualname"]
	var err error
	registration := inst.NewHostnameRegistration(instanceKey, hostname)
	if orcraft.IsRaftEnabled() {
		_, err = orcraft.PublishCommand("register-hostname-unresolve", registration)
	} else {
		err = inst.RegisterHostnameUnresolve(registration)
	}
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}
	Respond(r, &APIResponse{Code: OK, Message: "Hostname register unresolve completed", Details: instanceKey})
}

// SubmitPoolInstances (re-)applies the list of hostnames for a given pool
func (this *HttpAPI) SubmitPoolInstances(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	pool := params["pool"]
	instances := req.URL.Query().Get("instances")

	var err error
	submission := inst.NewPoolInstancesSubmission(pool, instances)
	if orcraft.IsRaftEnabled() {
		_, err = orcraft.PublishCommand("submit-pool-instances", submission)
	} else {
		err = inst.ApplyPoolInstances(submission)
	}
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Applied %s pool instances", pool), Details: pool})
}

// SubmitPoolHostnames (re-)applies the list of hostnames for a given pool
func (this *HttpAPI) ReadClusterPoolInstancesMap(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	clusterName := params["clusterName"]
	pool := params["pool"]

	poolInstancesMap, err := inst.ReadClusterPoolInstancesMap(clusterName, pool)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Read pool instances for cluster %s", clusterName), Details: poolInstancesMap})
}

// GetHeuristicClusterPoolInstances returns instances belonging to a cluster's pool
func (this *HttpAPI) GetHeuristicClusterPoolInstances(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	clusterName, err := figureClusterName(getClusterHint(params))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}
	pool := params["pool"]

	instances, err := inst.GetHeuristicClusterPoolInstances(clusterName, pool)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Heuristic pool instances for cluster %s", clusterName), Details: instances})
}

// GetHeuristicClusterPoolInstances returns instances belonging to a cluster's pool
func (this *HttpAPI) GetHeuristicClusterPoolInstancesLag(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	clusterName, err := inst.ReadClusterNameByAlias(params["clusterName"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}
	pool := params["pool"]

	lag, err := inst.GetHeuristicClusterPoolInstancesLag(clusterName, pool)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Heuristic pool lag for cluster %s", clusterName), Details: lag})
}

// ReloadClusterAlias clears in-memory hostname resovle cache
func (this *HttpAPI) ReloadClusterAlias(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}

	Respond(r, &APIResponse{Code: ERROR, Message: "This API call has been retired"})
}

// BulkPromotionRules returns a list of the known promotion rules for each instance
func (this *HttpAPI) BulkPromotionRules(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}

	promotionRules, err := inst.BulkReadCandidateDatabaseInstance()
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, promotionRules)
}

// BulkInstances returns a list of all known instances
func (this *HttpAPI) BulkInstances(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}

	instances, err := inst.BulkReadInstance()
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, instances)
}

// DiscoveryMetricsRaw will return the last X seconds worth of discovery information in time based order as a JSON array
func (this *HttpAPI) DiscoveryMetricsRaw(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	seconds, err := strconv.Atoi(params["seconds"])
	if err != nil || seconds <= 0 {
		Respond(r, &APIResponse{Code: ERROR, Message: "Invalid value provided for seconds"})
		return
	}

	refTime := time.Now().Add(-time.Duration(seconds) * time.Second)
	json, err := discovery.JSONSince(discoveryMetrics, refTime)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unable to determine start time. Perhaps seconds value is wrong?"})
		return
	}
	log.Debugf("DiscoveryMetricsRaw data: retrieved %d entries from discovery.MC", len(json))

	r.JSON(http.StatusOK, json)
}

// DiscoveryMetricsAggregated will return a single set of aggregated metrics for raw values collected since the
// specified time.
func (this *HttpAPI) DiscoveryMetricsAggregated(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	seconds, err := strconv.Atoi(params["seconds"])

	refTime := time.Now().Add(-time.Duration(seconds) * time.Second)
	aggregated, err := discovery.AggregatedSince(discoveryMetrics, refTime)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unable to generate aggregated discovery metrics"})
		return
	}
	// log.Debugf("DiscoveryMetricsAggregated data: %+v", aggregated)
	r.JSON(http.StatusOK, aggregated)
}

// DiscoveryQueueMetricsRaw returns the raw queue metrics (active and
// queued values), data taken secondly for the last N seconds.
func (this *HttpAPI) DiscoveryQueueMetricsRaw(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	seconds, err := strconv.Atoi(params["seconds"])
	log.Debugf("DiscoveryQueueMetricsRaw: seconds: %d", seconds)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unable to generate discovery queue  aggregated metrics"})
		return
	}

	queue := discovery.CreateOrReturnQueue("DEFAULT")
	metrics := queue.DiscoveryQueueMetrics(seconds)
	log.Debugf("DiscoveryQueueMetricsRaw data: %+v", metrics)

	r.JSON(http.StatusOK, metrics)
}

// DiscoveryQueueMetricsAggregated returns a single value showing the metrics of the discovery queue over the last N seconds.
// This is expected to be called every 60 seconds (?) and the config setting of the retention period is currently hard-coded.
// See go/discovery/ for more information.
func (this *HttpAPI) DiscoveryQueueMetricsAggregated(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	seconds, err := strconv.Atoi(params["seconds"])
	log.Debugf("DiscoveryQueueMetricsAggregated: seconds: %d", seconds)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unable to generate discovery queue aggregated metrics"})
		return
	}

	queue := discovery.CreateOrReturnQueue("DEFAULT")
	aggregated := queue.AggregatedDiscoveryQueueMetrics(seconds)
	log.Debugf("DiscoveryQueueMetricsAggregated data: %+v", aggregated)

	r.JSON(http.StatusOK, aggregated)
}

// BackendQueryMetricsRaw returns the raw backend query metrics
func (this *HttpAPI) BackendQueryMetricsRaw(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	seconds, err := strconv.Atoi(params["seconds"])
	log.Debugf("BackendQueryMetricsRaw: seconds: %d", seconds)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unable to generate raw backend query metrics"})
		return
	}

	refTime := time.Now().Add(-time.Duration(seconds) * time.Second)
	m, err := queryMetrics.Since(refTime)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unable to return backend query metrics"})
		return
	}

	log.Debugf("BackendQueryMetricsRaw data: %+v", m)

	r.JSON(http.StatusOK, m)
}

func (this *HttpAPI) BackendQueryMetricsAggregated(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	seconds, err := strconv.Atoi(params["seconds"])
	log.Debugf("BackendQueryMetricsAggregated: seconds: %d", seconds)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unable to aggregated generate backend query metrics"})
		return
	}

	refTime := time.Now().Add(-time.Duration(seconds) * time.Second)
	aggregated := query.AggregatedSince(queryMetrics, refTime)
	log.Debugf("BackendQueryMetricsAggregated data: %+v", aggregated)

	r.JSON(http.StatusOK, aggregated)
}

// Agents provides complete list of registered agents (See https://github.com/github/orchestrator-agent)
func (this *HttpAPI) Agents(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	agents, err := agent.ReadAgents()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, agents)
}

// Agent returns complete information of a given agent
func (this *HttpAPI) Agent(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	agent, err := agent.GetAgent(params["host"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, agent)
}

// AgentUnmount instructs an agent to unmount the designated mount point
func (this *HttpAPI) AgentUnmount(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.Unmount(params["host"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentMountLV instructs an agent to mount a given volume on the designated mount point
func (this *HttpAPI) AgentMountLV(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.MountLV(params["host"], req.URL.Query().Get("lv"))

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentCreateSnapshot instructs an agent to create a new snapshot. Agent's DIY implementation.
func (this *HttpAPI) AgentCreateSnapshot(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.CreateSnapshot(params["host"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentRemoveLV instructs an agent to remove a logical volume
func (this *HttpAPI) AgentRemoveLV(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.RemoveLV(params["host"], req.URL.Query().Get("lv"))

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentMySQLStop stops MySQL service on agent
func (this *HttpAPI) AgentMySQLStop(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.MySQLStop(params["host"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentMySQLStart starts MySQL service on agent
func (this *HttpAPI) AgentMySQLStart(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.MySQLStart(params["host"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

func (this *HttpAPI) AgentCustomCommand(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.CustomCommand(params["host"], params["command"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentSeed completely seeds a host with another host's snapshots. This is a complex operation
// governed by orchestrator and executed by the two agents involved.
func (this *HttpAPI) AgentSeed(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.Seed(params["targetHost"], params["sourceHost"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentActiveSeeds lists active seeds and their state
func (this *HttpAPI) AgentActiveSeeds(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.ReadActiveSeedsForHost(params["host"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentRecentSeeds lists recent seeds of a given agent
func (this *HttpAPI) AgentRecentSeeds(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.ReadRecentCompletedSeedsForHost(params["host"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentSeedDetails provides details of a given seed
func (this *HttpAPI) AgentSeedDetails(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	seedId, err := strconv.ParseInt(params["seedId"], 10, 0)
	output, err := agent.AgentSeedDetails(seedId)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AgentSeedStates returns the breakdown of states (steps) of a given seed
func (this *HttpAPI) AgentSeedStates(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	seedId, err := strconv.ParseInt(params["seedId"], 10, 0)
	output, err := agent.ReadSeedStates(seedId)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// Seeds retruns all recent seeds
func (this *HttpAPI) Seeds(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.ReadRecentSeeds()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, output)
}

// AbortSeed instructs agents to abort an active seed
func (this *HttpAPI) AbortSeed(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		Respond(r, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	seedId, err := strconv.ParseInt(params["seedId"], 10, 0)
	err = agent.AbortSeed(seedId)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, err == nil)
}

// Headers is a self-test call which returns HTTP headers
func (this *HttpAPI) Headers(params martini.Params, r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, req.Header)
}

// Health performs a self test
func (this *HttpAPI) Health(params martini.Params, r render.Render, req *http.Request) {
	health, err := process.HealthTest()
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Application node is unhealthy %+v", err), Details: health})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Application node is healthy"), Details: health})

}

// LBCheck returns a constant respnse, and this can be used by load balancers that expect a given string.
func (this *HttpAPI) LBCheck(params martini.Params, r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, "OK")
}

// LBCheck returns a constant respnse, and this can be used by load balancers that expect a given string.
func (this *HttpAPI) LeaderCheck(params martini.Params, r render.Render, req *http.Request) {
	if logic.IsLeader() {
		r.JSON(http.StatusOK, "OK")
	} else {
		r.JSON(http.StatusNotFound, "Not leader")
	}
}

// A configurable endpoint that can be for regular status checks or whatever.  While similar to
// Health() this returns 500 on failure.  This will prevent issues for those that have come to
// expect a 200
// It might be a good idea to deprecate the current Health() behavior and roll this in at some
// point
func (this *HttpAPI) StatusCheck(params martini.Params, r render.Render, req *http.Request) {
	// SimpleHealthTest just checks to see if we can connect to the database.  Lighter weight if you intend to call it a lot
	var health *process.HealthStatus
	var err error
	if config.Config.StatusSimpleHealth {
		health, err = process.SimpleHealthTest()
	} else {
		health, err = process.HealthTest()
	}
	if err != nil {
		r.JSON(500, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Application node is unhealthy %+v", err), Details: health})
		return
	}
	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Application node is healthy"), Details: health})
}

// GrabElection forcibly grabs leadership. Use with care!!
func (this *HttpAPI) GrabElection(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	err := process.GrabElection()
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Unable to grab election: %+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Node elected as leader")})
}

// Reelect causes re-elections for an active node
func (this *HttpAPI) Reelect(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	err := process.Reelect()
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Unable to re-elect: %+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Set re-elections")})

}

// Reelect causes re-elections for an active node
func (this *HttpAPI) RaftYield(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !orcraft.IsRaftEnabled() {
		Respond(r, &APIResponse{Code: ERROR, Message: "raft-yield: not running with raft setup"})
		return
	}

	orcraft.PublishYield(params["node"])

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Asynchronously yielded")})
}

// RaftPeers returns the list of peers in a raft setup
func (this *HttpAPI) RaftPeers(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !orcraft.IsRaftEnabled() {
		Respond(r, &APIResponse{Code: ERROR, Message: "raft-nodes: not running with raft setup"})
		return
	}

	peers, err := orcraft.GetPeers()
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot get raft peers: %+v", err)})
		return
	}

	r.JSON(http.StatusOK, peers)
}

// ReloadConfiguration reloads confiug settings (not all of which will apply after change)
func (this *HttpAPI) ReloadConfiguration(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	config.Reload()
	inst.AuditOperation("reload-configuration", nil, "Triggered via API")

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Config reloaded")})

}

// ReplicationAnalysis retuens list of issues
func (this *HttpAPI) replicationAnalysis(clusterName string, instanceKey *inst.InstanceKey, params martini.Params, r render.Render, req *http.Request) {
	analysis, err := inst.GetReplicationAnalysis(clusterName, true, false)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot get analysis: %+v", err)})
		return
	}
	// Possibly filter single instance
	if instanceKey != nil {
		filtered := analysis[:0]
		for _, analysisEntry := range analysis {
			if instanceKey.Equals(&analysisEntry.AnalyzedInstanceKey) {
				filtered = append(filtered, analysisEntry)
			}
		}
		analysis = filtered
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Analysis"), Details: analysis})
}

// ReplicationAnalysis retuens list of issues
func (this *HttpAPI) ReplicationAnalysis(params martini.Params, r render.Render, req *http.Request) {
	this.replicationAnalysis("", nil, params, r, req)
}

// ReplicationAnalysis retuens list of issues
func (this *HttpAPI) ReplicationAnalysisForCluster(params martini.Params, r render.Render, req *http.Request) {
	clusterName := params["clusterName"]

	var err error
	if clusterName, err = inst.DeduceClusterName(params["clusterName"]); err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot get analysis: %+v", err)})
		return
	}
	if clusterName == "" {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot get cluster name: %+v", params["clusterName"])})
		return
	}
	this.replicationAnalysis(clusterName, nil, params, r, req)
}

// ReplicationAnalysis retuens list of issues
func (this *HttpAPI) ReplicationAnalysisForKey(params martini.Params, r render.Render, req *http.Request) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot get analysis: %+v", err)})
		return
	}
	if !instanceKey.IsValid() {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("Cannot get analysis: invalid key %+v", instanceKey)})
		return
	}
	this.replicationAnalysis("", &instanceKey, params, r, req)
}

// RecoverLite attempts recovery on a given instance, without executing external processes
func (this *HttpAPI) RecoverLite(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	params["skipProcesses"] = "true"
	this.Recover(params, r, req, user)
}

// Recover attempts recovery on a given instance
func (this *HttpAPI) Recover(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	var candidateKey *inst.InstanceKey
	if key, err := this.getInstanceKey(params["candidateHost"], params["candidatePort"]); err == nil {
		candidateKey = &key
	}

	skipProcesses := (req.URL.Query().Get("skipProcesses") == "true") || (params["skipProcesses"] == "true")
	recoveryAttempted, promotedInstanceKey, err := logic.CheckAndRecover(&instanceKey, candidateKey, skipProcesses)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error(), Details: instanceKey})
		return
	}
	if !recoveryAttempted {
		Respond(r, &APIResponse{Code: ERROR, Message: "Recovery not attempted", Details: instanceKey})
		return
	}
	if promotedInstanceKey == nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "Recovery attempted but no instance promoted", Details: instanceKey})
		return
	}
	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Recovery executed on %+v", instanceKey), Details: *promotedInstanceKey})
}

// GracefulMasterTakeover gracefully fails over a master onto its single replica.
func (this *HttpAPI) GracefulMasterTakeover(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	clusterName, err := figureClusterName(getClusterHint(params))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	topologyRecovery, _, err := logic.GracefulMasterTakeover(clusterName)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error(), Details: topologyRecovery})
		return
	}
	if topologyRecovery.SuccessorKey == nil {
		Respond(r, &APIResponse{Code: ERROR, Message: "graceful-master-takeover: no successor promoted", Details: topologyRecovery})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "graceful-master-takeover: successor promoted", Details: topologyRecovery})
}

// ForceMasterFailover fails over a master (even if there's no particular problem with the master)
func (this *HttpAPI) ForceMasterFailover(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	clusterName, err := figureClusterName(getClusterHint(params))
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	topologyRecovery, err := logic.ForceMasterFailover(clusterName)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	if topologyRecovery.SuccessorKey != nil {
		Respond(r, &APIResponse{Code: OK, Message: "Master failed over", Details: topologyRecovery})
	} else {
		Respond(r, &APIResponse{Code: ERROR, Message: "Master not failed over", Details: topologyRecovery})
	}
}

// Registers promotion preference for given instance
func (this *HttpAPI) RegisterCandidate(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	promotionRule, err := inst.ParseCandidatePromotionRule(params["promotionRule"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	candidate := inst.NewCandidateDatabaseInstance(&instanceKey, promotionRule)

	if orcraft.IsRaftEnabled() {
		_, err = orcraft.PublishCommand("register-candidate", candidate)
	} else {
		err = inst.RegisterCandidateInstance(candidate)
	}

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "Registered candidate", Details: instanceKey})
}

// AutomatedRecoveryFilters retuens list of clusters which are configured with automated recovery
func (this *HttpAPI) AutomatedRecoveryFilters(params martini.Params, r render.Render, req *http.Request) {
	automatedRecoveryMap := make(map[string]interface{})
	automatedRecoveryMap["RecoverMasterClusterFilters"] = config.Config.RecoverMasterClusterFilters
	automatedRecoveryMap["RecoverIntermediateMasterClusterFilters"] = config.Config.RecoverIntermediateMasterClusterFilters
	automatedRecoveryMap["RecoveryIgnoreHostnameFilters"] = config.Config.RecoveryIgnoreHostnameFilters

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Automated recovery configuration details"), Details: automatedRecoveryMap})
}

// AuditFailureDetection provides list of topology_failure_detection entries
func (this *HttpAPI) AuditFailureDetection(params martini.Params, r render.Render, req *http.Request) {

	var audits []logic.TopologyRecovery
	var err error

	if detectionId, derr := strconv.ParseInt(params["id"], 10, 0); derr == nil && detectionId > 0 {
		audits, err = logic.ReadFailureDetection(detectionId)
	} else {
		page, derr := strconv.Atoi(params["page"])
		if derr != nil || page < 0 {
			page = 0
		}
		audits, err = logic.ReadRecentFailureDetections(page)
	}

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, audits)
}

// AuditRecoverySteps returns audited steps of a given recovery
func (this *HttpAPI) AuditRecoverySteps(params martini.Params, r render.Render, req *http.Request) {
	recoveryUID := params["uid"]
	audits, err := logic.ReadTopologyRecoverySteps(recoveryUID)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, audits)
}

// ReadReplicationAnalysisChangelog lists instances and their analysis changelog
func (this *HttpAPI) ReadReplicationAnalysisChangelog(params martini.Params, r render.Render, req *http.Request) {
	changelogs, err := inst.ReadReplicationAnalysisChangelog()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, changelogs)
}

// AuditRecovery provides list of topology-recovery entries
func (this *HttpAPI) AuditRecovery(params martini.Params, r render.Render, req *http.Request) {
	var audits []logic.TopologyRecovery
	var err error

	if recoveryUID := params["uid"]; recoveryUID != "" {
		audits, err = logic.ReadRecoveryByUID(recoveryUID)
	} else if recoveryId, derr := strconv.ParseInt(params["id"], 10, 0); derr == nil && recoveryId > 0 {
		audits, err = logic.ReadRecovery(recoveryId)
	} else {
		page, derr := strconv.Atoi(params["page"])
		if derr != nil || page < 0 {
			page = 0
		}
		unacknowledgedOnly := (req.URL.Query().Get("unacknowledged") == "true")
		audits, err = logic.ReadRecentRecoveries(params["clusterName"], unacknowledgedOnly, page)
	}

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, audits)
}

// ActiveClusterRecovery returns recoveries in-progress for a given cluster
func (this *HttpAPI) ActiveClusterRecovery(params martini.Params, r render.Render, req *http.Request) {
	recoveries, err := logic.ReadActiveClusterRecovery(params["clusterName"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, recoveries)
}

// RecentlyActiveClusterRecovery returns recoveries in-progress for a given cluster
func (this *HttpAPI) RecentlyActiveClusterRecovery(params martini.Params, r render.Render, req *http.Request) {
	recoveries, err := logic.ReadRecentlyActiveClusterRecovery(params["clusterName"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, recoveries)
}

// RecentlyActiveClusterRecovery returns recoveries in-progress for a given cluster
func (this *HttpAPI) RecentlyActiveInstanceRecovery(params martini.Params, r render.Render, req *http.Request) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	recoveries, err := logic.ReadRecentlyActiveInstanceRecovery(&instanceKey)

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, recoveries)
}

// ClusterInfo provides details of a given cluster
func (this *HttpAPI) AcknowledgeClusterRecoveries(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	var err error

	clusterName := params["clusterName"]
	if params["clusterAlias"] != "" {
		clusterName, err = inst.GetClusterByAlias(params["clusterAlias"])
		if err != nil {
			Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
			return
		}
	}

	comment := req.URL.Query().Get("comment")
	if comment == "" {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("No acknowledge comment given")})
		return
	}
	userId := getUserId(req, user)
	if userId == "" {
		userId = inst.GetMaintenanceOwner()
	}
	var countAcknowledgedRecoveries int64
	if orcraft.IsRaftEnabled() {
		ack := logic.NewRecoveryAcknowledgement(userId, comment)
		ack.ClusterName = clusterName
		orcraft.PublishCommand("ack-recovery", ack)
	} else {
		countAcknowledgedRecoveries, err = logic.AcknowledgeClusterRecoveries(clusterName, userId, comment)
	}
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Acknowledged %s recoveries on %+v", countAcknowledgedRecoveries, clusterName), Details: countAcknowledgedRecoveries})
}

// ClusterInfo provides details of a given cluster
func (this *HttpAPI) AcknowledgeInstanceRecoveries(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}

	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	comment := req.URL.Query().Get("comment")
	if comment == "" {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("No acknowledge comment given")})
		return
	}
	userId := getUserId(req, user)
	if userId == "" {
		userId = inst.GetMaintenanceOwner()
	}
	var countAcknowledgedRecoveries int64
	if orcraft.IsRaftEnabled() {
		ack := logic.NewRecoveryAcknowledgement(userId, comment)
		ack.Key = instanceKey
		orcraft.PublishCommand("ack-recovery", ack)
	} else {
		countAcknowledgedRecoveries, err = logic.AcknowledgeInstanceRecoveries(&instanceKey, userId, comment)
	}
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, countAcknowledgedRecoveries)
}

// ClusterInfo provides details of a given cluster
func (this *HttpAPI) AcknowledgeRecovery(params martini.Params, r render.Render, req *http.Request, user auth.User) {
	if !isAuthorizedForAction(req, user) {
		Respond(r, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}

	recoveryId, err := strconv.ParseInt(params["recoveryId"], 10, 0)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	comment := req.URL.Query().Get("comment")
	if comment == "" {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("No acknowledge comment given")})
		return
	}
	userId := getUserId(req, user)
	if userId == "" {
		userId = inst.GetMaintenanceOwner()
	}
	countAcknowledgedRecoveries, err := logic.AcknowledgeRecovery(recoveryId, userId, comment)
	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, countAcknowledgedRecoveries)
}

// BlockedRecoveries reads list of currently blocked recoveries, optionally filtered by cluster name
func (this *HttpAPI) BlockedRecoveries(params martini.Params, r render.Render, req *http.Request) {
	blockedRecoveries, err := logic.ReadBlockedRecoveries(params["clusterName"])

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(http.StatusOK, blockedRecoveries)
}

// DisableGlobalRecoveries globally disables recoveries
func (this *HttpAPI) DisableGlobalRecoveries(params martini.Params, r render.Render, req *http.Request) {
	err := logic.DisableRecovery()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "Globally disabled recoveries", Details: "disabled"})
}

// EnableGlobalRecoveries globally enables recoveries
func (this *HttpAPI) EnableGlobalRecoveries(params martini.Params, r render.Render, req *http.Request) {
	err := logic.EnableRecovery()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	Respond(r, &APIResponse{Code: OK, Message: "Globally enabled recoveries", Details: "enabled"})
}

// CheckGlobalRecoveries checks whether
func (this *HttpAPI) CheckGlobalRecoveries(params martini.Params, r render.Render, req *http.Request) {
	isDisabled, err := logic.IsRecoveryDisabled()

	if err != nil {
		Respond(r, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}
	details := "enabled"
	if isDisabled {
		details = "disabled"
	}
	Respond(r, &APIResponse{Code: OK, Message: fmt.Sprintf("Global recoveries %+v", details), Details: details})
}

func (this *HttpAPI) getSynonymPath(path string) (synonymPath string) {
	pathBase := strings.Split(path, "/")[0]
	if synonym, ok := apiSynonyms[pathBase]; ok {
		synonymPath = fmt.Sprintf("%s%s", synonym, path[len(pathBase):])
	}
	return synonymPath
}

func (this *HttpAPI) registerSingleRequest(m *martini.ClassicMartini, path string, handler martini.Handler) {
	registeredPaths = append(registeredPaths, path)
	fullPath := fmt.Sprintf("%s/api/%s", this.URLPrefix, path)
	m.Get(fullPath, handler)

}

func (this *HttpAPI) registerRequest(m *martini.ClassicMartini, path string, handler martini.Handler) {
	this.registerSingleRequest(m, path, handler)

	if synonym := this.getSynonymPath(path); synonym != "" {
		this.registerSingleRequest(m, synonym, handler)
	}
}

// RegisterRequests makes for the de-facto list of known API calls
func (this *HttpAPI) RegisterRequests(m *martini.ClassicMartini) {
	// Smart relocation:
	this.registerRequest(m, "relocate/:host/:port/:belowHost/:belowPort", this.RelocateBelow)
	this.registerRequest(m, "relocate-below/:host/:port/:belowHost/:belowPort", this.RelocateBelow)
	this.registerRequest(m, "relocate-slaves/:host/:port/:belowHost/:belowPort", this.RelocateReplicas)
	this.registerRequest(m, "regroup-slaves/:host/:port", this.RegroupReplicas)

	// Classic file:pos relocation:
	this.registerRequest(m, "move-up/:host/:port", this.MoveUp)
	this.registerRequest(m, "move-up-slaves/:host/:port", this.MoveUpReplicas)
	this.registerRequest(m, "move-below/:host/:port/:siblingHost/:siblingPort", this.MoveBelow)
	this.registerRequest(m, "move-equivalent/:host/:port/:belowHost/:belowPort", this.MoveEquivalent)
	this.registerRequest(m, "repoint/:host/:port/:belowHost/:belowPort", this.Repoint)
	this.registerRequest(m, "repoint-slaves/:host/:port", this.RepointReplicas)
	this.registerRequest(m, "make-co-master/:host/:port", this.MakeCoMaster)
	this.registerRequest(m, "enslave-siblings/:host/:port", this.TakeSiblings)
	this.registerRequest(m, "enslave-master/:host/:port", this.TakeMaster)
	this.registerRequest(m, "master-equivalent/:host/:port/:logFile/:logPos", this.MasterEquivalent)

	// Binlog server relocation:
	this.registerRequest(m, "regroup-slaves-bls/:host/:port", this.RegroupReplicasBinlogServers)

	// GTID relocation:
	this.registerRequest(m, "move-below-gtid/:host/:port/:belowHost/:belowPort", this.MoveBelowGTID)
	this.registerRequest(m, "move-slaves-gtid/:host/:port/:belowHost/:belowPort", this.MoveReplicasGTID)
	this.registerRequest(m, "regroup-slaves-gtid/:host/:port", this.RegroupReplicasGTID)

	// Pseudo-GTID relocation:
	this.registerRequest(m, "match/:host/:port/:belowHost/:belowPort", this.MatchBelow)
	this.registerRequest(m, "match-below/:host/:port/:belowHost/:belowPort", this.MatchBelow)
	this.registerRequest(m, "match-up/:host/:port", this.MatchUp)
	this.registerRequest(m, "match-slaves/:host/:port/:belowHost/:belowPort", this.MultiMatchReplicas)
	this.registerRequest(m, "match-up-slaves/:host/:port", this.MatchUpReplicas)
	this.registerRequest(m, "regroup-slaves-pgtid/:host/:port", this.RegroupReplicasPseudoGTID)
	// Legacy, need to revisit:
	this.registerRequest(m, "make-master/:host/:port", this.MakeMaster)
	this.registerRequest(m, "make-local-master/:host/:port", this.MakeLocalMaster)

	// Replication, general:
	this.registerRequest(m, "enable-gtid/:host/:port", this.EnableGTID)
	this.registerRequest(m, "disable-gtid/:host/:port", this.DisableGTID)
	this.registerRequest(m, "skip-query/:host/:port", this.SkipQuery)
	this.registerRequest(m, "start-slave/:host/:port", this.StartSlave)
	this.registerRequest(m, "restart-slave/:host/:port", this.RestartSlave)
	this.registerRequest(m, "stop-slave/:host/:port", this.StopSlave)
	this.registerRequest(m, "stop-slave-nice/:host/:port", this.StopSlaveNicely)
	this.registerRequest(m, "reset-slave/:host/:port", this.ResetSlave)
	this.registerRequest(m, "detach-slave/:host/:port", this.DetachReplica)
	this.registerRequest(m, "reattach-slave/:host/:port", this.ReattachReplica)
	this.registerRequest(m, "reattach-slave-master-host/:host/:port", this.ReattachReplicaMasterHost)
	this.registerRequest(m, "flush-binary-logs/:host/:port", this.FlushBinaryLogs)

	// Instance:
	this.registerRequest(m, "set-read-only/:host/:port", this.SetReadOnly)
	this.registerRequest(m, "set-writeable/:host/:port", this.SetWriteable)
	this.registerRequest(m, "kill-query/:host/:port/:process", this.KillQuery)

	// Binary logs:
	this.registerRequest(m, "last-pseudo-gtid/:host/:port", this.LastPseudoGTID)

	// Pools:
	this.registerRequest(m, "submit-pool-instances/:pool", this.SubmitPoolInstances)
	this.registerRequest(m, "cluster-pool-instances/:clusterName", this.ReadClusterPoolInstancesMap)
	this.registerRequest(m, "cluster-pool-instances/:clusterName/:pool", this.ReadClusterPoolInstancesMap)
	this.registerRequest(m, "heuristic-cluster-pool-instances/:clusterName", this.GetHeuristicClusterPoolInstances)
	this.registerRequest(m, "heuristic-cluster-pool-instances/:clusterName/:pool", this.GetHeuristicClusterPoolInstances)
	this.registerRequest(m, "heuristic-cluster-pool-lag/:clusterName", this.GetHeuristicClusterPoolInstancesLag)
	this.registerRequest(m, "heuristic-cluster-pool-lag/:clusterName/:pool", this.GetHeuristicClusterPoolInstancesLag)

	// Information:
	this.registerRequest(m, "search/:searchString", this.Search)
	this.registerRequest(m, "search", this.Search)

	// Cluster
	this.registerRequest(m, "cluster/:clusterHint", this.Cluster)
	this.registerRequest(m, "cluster/alias/:clusterAlias", this.ClusterByAlias)
	this.registerRequest(m, "cluster/instance/:host/:port", this.ClusterByInstance)
	this.registerRequest(m, "cluster-info/:clusterHint", this.ClusterInfo)
	this.registerRequest(m, "cluster-info/alias/:clusterAlias", this.ClusterInfoByAlias)
	this.registerRequest(m, "cluster-osc-slaves/:clusterHint", this.ClusterOSCReplicas)
	this.registerRequest(m, "set-cluster-alias/:clusterName", this.SetClusterAliasManualOverride)
	this.registerRequest(m, "clusters", this.Clusters)
	this.registerRequest(m, "clusters-info", this.ClustersInfo)
	this.registerRequest(m, "masters", this.Masters)
	this.registerRequest(m, "instance-replicas/:host/:port", this.InstanceReplicas)
	this.registerRequest(m, "all-instances", this.AllInstances)
	this.registerRequest(m, "downtimed", this.Downtimed)
	this.registerRequest(m, "downtimed/:clusterHint", this.Downtimed)
	this.registerRequest(m, "topology/:clusterHint", this.AsciiTopology)
	this.registerRequest(m, "topology/:host/:port", this.AsciiTopology)

	// Instance management:
	this.registerRequest(m, "instance/:host/:port", this.Instance)
	this.registerRequest(m, "discover/:host/:port", this.Discover)
	this.registerRequest(m, "async-discover/:host/:port", this.AsyncDiscover)
	this.registerRequest(m, "refresh/:host/:port", this.Refresh)
	this.registerRequest(m, "forget/:host/:port", this.Forget)
	this.registerRequest(m, "begin-maintenance/:host/:port/:owner/:reason", this.BeginMaintenance)
	this.registerRequest(m, "end-maintenance/:host/:port", this.EndMaintenanceByInstanceKey)
	this.registerRequest(m, "end-maintenance/:maintenanceKey", this.EndMaintenance)
	this.registerRequest(m, "begin-downtime/:host/:port/:owner/:reason", this.BeginDowntime)
	this.registerRequest(m, "begin-downtime/:host/:port/:owner/:reason/:duration", this.BeginDowntime)
	this.registerRequest(m, "end-downtime/:host/:port", this.EndDowntime)

	// Recovery:
	this.registerRequest(m, "replication-analysis", this.ReplicationAnalysis)
	this.registerRequest(m, "replication-analysis/:clusterName", this.ReplicationAnalysisForCluster)
	this.registerRequest(m, "replication-analysis/instance/:host/:port", this.ReplicationAnalysisForKey)
	this.registerRequest(m, "recover/:host/:port", this.Recover)
	this.registerRequest(m, "recover/:host/:port/:candidateHost/:candidatePort", this.Recover)
	this.registerRequest(m, "recover-lite/:host/:port", this.RecoverLite)
	this.registerRequest(m, "recover-lite/:host/:port/:candidateHost/:candidatePort", this.RecoverLite)
	this.registerRequest(m, "graceful-master-takeover/:host/:port", this.GracefulMasterTakeover)
	this.registerRequest(m, "graceful-master-takeover/:clusterHint", this.GracefulMasterTakeover)
	this.registerRequest(m, "force-master-failover/:host/:port", this.ForceMasterFailover)
	this.registerRequest(m, "force-master-failover/:clusterHint", this.ForceMasterFailover)
	this.registerRequest(m, "register-candidate/:host/:port/:promotionRule", this.RegisterCandidate)
	this.registerRequest(m, "automated-recovery-filters", this.AutomatedRecoveryFilters)
	this.registerRequest(m, "audit-failure-detection", this.AuditFailureDetection)
	this.registerRequest(m, "audit-failure-detection/:page", this.AuditFailureDetection)
	this.registerRequest(m, "audit-failure-detection/id/:id", this.AuditFailureDetection)
	this.registerRequest(m, "replication-analysis-changelog", this.ReadReplicationAnalysisChangelog)
	this.registerRequest(m, "audit-recovery", this.AuditRecovery)
	this.registerRequest(m, "audit-recovery/:page", this.AuditRecovery)
	this.registerRequest(m, "audit-recovery/id/:id", this.AuditRecovery)
	this.registerRequest(m, "audit-recovery/uid/:uid", this.AuditRecovery)
	this.registerRequest(m, "audit-recovery/cluster/:clusterName", this.AuditRecovery)
	this.registerRequest(m, "audit-recovery/cluster/:clusterName/:page", this.AuditRecovery)
	this.registerRequest(m, "audit-recovery-steps/:uid", this.AuditRecoverySteps)
	this.registerRequest(m, "active-cluster-recovery/:clusterName", this.ActiveClusterRecovery)
	this.registerRequest(m, "recently-active-cluster-recovery/:clusterName", this.RecentlyActiveClusterRecovery)
	this.registerRequest(m, "recently-active-instance-recovery/:host/:port", this.RecentlyActiveInstanceRecovery)
	this.registerRequest(m, "ack-recovery/cluster/:clusterName", this.AcknowledgeClusterRecoveries)
	this.registerRequest(m, "ack-recovery/cluster/alias/:clusterAlias", this.AcknowledgeClusterRecoveries)
	this.registerRequest(m, "ack-recovery/instance/:host/:port", this.AcknowledgeInstanceRecoveries)
	this.registerRequest(m, "ack-recovery/:recoveryId", this.AcknowledgeRecovery)
	this.registerRequest(m, "blocked-recoveries", this.BlockedRecoveries)
	this.registerRequest(m, "blocked-recoveries/cluster/:clusterName", this.BlockedRecoveries)
	this.registerRequest(m, "disable-global-recoveries", this.DisableGlobalRecoveries)
	this.registerRequest(m, "enable-global-recoveries", this.EnableGlobalRecoveries)
	this.registerRequest(m, "check-global-recoveries", this.CheckGlobalRecoveries)

	// General
	this.registerRequest(m, "problems", this.Problems)
	this.registerRequest(m, "problems/:clusterName", this.Problems)
	this.registerRequest(m, "long-queries", this.LongQueries)
	this.registerRequest(m, "long-queries/:filter", this.LongQueries)
	this.registerRequest(m, "audit", this.Audit)
	this.registerRequest(m, "audit/:page", this.Audit)
	this.registerRequest(m, "audit/instance/:host/:port", this.Audit)
	this.registerRequest(m, "audit/instance/:host/:port/:page", this.Audit)
	this.registerRequest(m, "resolve/:host/:port", this.Resolve)

	// Meta
	this.registerRequest(m, "maintenance", this.Maintenance)
	this.registerRequest(m, "headers", this.Headers)
	this.registerRequest(m, "health", this.Health)
	this.registerRequest(m, "lb-check", this.LBCheck)
	this.registerRequest(m, "_ping", this.LBCheck)
	this.registerRequest(m, "leader-check", this.LeaderCheck)
	this.registerRequest(m, "grab-election", this.GrabElection)
	this.registerRequest(m, "raft-yield/:node", this.RaftYield)
	this.registerRequest(m, "raft-peers", this.RaftPeers)
	this.registerRequest(m, "reelect", this.Reelect)
	this.registerRequest(m, "reload-configuration", this.ReloadConfiguration)
	this.registerRequest(m, "reload-cluster-alias", this.ReloadClusterAlias)
	this.registerRequest(m, "hostname-resolve-cache", this.HostnameResolveCache)
	this.registerRequest(m, "reset-hostname-resolve-cache", this.ResetHostnameResolveCache)
	this.registerRequest(m, "deregister-hostname-unresolve/:host/:port", this.DeregisterHostnameUnresolve)
	this.registerRequest(m, "register-hostname-unresolve/:host/:port/:virtualname", this.RegisterHostnameUnresolve)
	// Bulk access to information
	this.registerRequest(m, "bulk-instances", this.BulkInstances)
	this.registerRequest(m, "bulk-promotion-rules", this.BulkPromotionRules)

	// Monitoring
	this.registerRequest(m, "discovery-metrics-raw/:seconds", this.DiscoveryMetricsRaw)
	this.registerRequest(m, "discovery-metrics-aggregated/:seconds", this.DiscoveryMetricsAggregated)
	this.registerRequest(m, "discovery-queue-metrics-raw/:seconds", this.DiscoveryQueueMetricsRaw)
	this.registerRequest(m, "discovery-queue-metrics-aggregated/:seconds", this.DiscoveryQueueMetricsAggregated)
	this.registerRequest(m, "backend-query-metrics-raw/:seconds", this.BackendQueryMetricsRaw)
	this.registerRequest(m, "backend-query-metrics-aggregated/:seconds", this.BackendQueryMetricsAggregated)

	// Agents
	this.registerRequest(m, "agents", this.Agents)
	this.registerRequest(m, "agent/:host", this.Agent)
	this.registerRequest(m, "agent-umount/:host", this.AgentUnmount)
	this.registerRequest(m, "agent-mount/:host", this.AgentMountLV)
	this.registerRequest(m, "agent-create-snapshot/:host", this.AgentCreateSnapshot)
	this.registerRequest(m, "agent-removelv/:host", this.AgentRemoveLV)
	this.registerRequest(m, "agent-mysql-stop/:host", this.AgentMySQLStop)
	this.registerRequest(m, "agent-mysql-start/:host", this.AgentMySQLStart)
	this.registerRequest(m, "agent-seed/:targetHost/:sourceHost", this.AgentSeed)
	this.registerRequest(m, "agent-active-seeds/:host", this.AgentActiveSeeds)
	this.registerRequest(m, "agent-recent-seeds/:host", this.AgentRecentSeeds)
	this.registerRequest(m, "agent-seed-details/:seedId", this.AgentSeedDetails)
	this.registerRequest(m, "agent-seed-states/:seedId", this.AgentSeedStates)
	this.registerRequest(m, "agent-abort-seed/:seedId", this.AbortSeed)
	this.registerRequest(m, "agent-custom-command/:host/:command", this.AgentCustomCommand)
	this.registerRequest(m, "seeds", this.Seeds)

	// Configurable status check endpoint
	m.Get(config.Config.StatusEndpoint, this.StatusCheck)
}
