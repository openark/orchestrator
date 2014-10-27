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
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/outbrain/orchestrator/agent"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/inst"
	"github.com/outbrain/orchestrator/logic"
)

// APIResponseCode is an OK/ERROR response code
type APIResponseCode int

const (
	ERROR APIResponseCode = iota
	OK
)

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

// APIResponse is a response returned as JSON to various requests.
type APIResponse struct {
	Code    APIResponseCode
	Message string
	Details interface{}
}

type HttpAPI struct{}

var API HttpAPI = HttpAPI{}

func (this *HttpAPI) getAuthUser(req *http.Request) string {
	for _, user := range req.Header[config.Config.AuthUserHeader] {
		return user
	}
	return ""
}

// isAuthorizedForAction checks req to see whether authenticated user has write-privileges.
// This depends on configured authentication method, and at this stage only applies in "proxy" method.
func (this *HttpAPI) isAuthorizedForAction(req *http.Request) bool {
	if strings.ToLower(config.Config.AuthenticationMethod) != "proxy" {
		return true
	}
	authUser := this.getAuthUser(req)
	for _, user := range config.Config.PowerAuthUsers {
		if user == "*" || user == authUser {
			return true
		}
	}
	return false
}

func (this *HttpAPI) getInstanceKey(host string, port string) (inst.InstanceKey, error) {
	instanceKey, err := inst.NewInstanceKeyFromStrings(host, port)
	return *instanceKey, err
}

// Instance reads and returns an instance's details.
func (this *HttpAPI) Instance(params martini.Params, r render.Render, req *http.Request) {
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
	r.JSON(200, instance)
}

// Discover starts an asynchronuous discovery for an instance
func (this *HttpAPI) Discover(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	go orchestrator.StartDiscovery(instanceKey)

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance submitted for discovery: %+v", instanceKey)})
}

// Refresh synchronuously re-reads a topology instance
func (this *HttpAPI) Refresh(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	_, err = inst.RefreshTopologyInstance(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance refreshedh: %+v", instanceKey)})
}

// Forget removes an instance entry fro backend database
func (this *HttpAPI) Forget(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	// We ignore errors: we're looking to do a destructive operation anyhow.
	instanceKey, _ := this.getInstanceKey(params["host"], params["port"])

	inst.ForgetInstance(&instanceKey)

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance forgotten: %+v", instanceKey)})
}

// Resolve tries to resolve hostname and then checks to see if port is open on that host.
func (this *HttpAPI) Resolve(params martini.Params, r render.Render, req *http.Request) {
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	if conn, err := net.Dial("tcp", instanceKey.DisplayString()); err == nil {
		conn.Close()
	} else {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Instance resolved", Details: instanceKey})
}

// BeginMaintenance begins maintenance mode for given instance
func (this *HttpAPI) BeginMaintenance(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	key, err := inst.BeginMaintenance(&instanceKey, params["owner"], params["reason"])
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error(), Details: key})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Maintenance begun: %+v", instanceKey)})
}

// EndMaintenance terminates maintenance mode
func (this *HttpAPI) EndMaintenance(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	maintenanceKey, err := strconv.ParseInt(params["maintenanceKey"], 10, 0)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	err = inst.EndMaintenance(maintenanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Maintenance ended: %+v", maintenanceKey)})
}

// EndMaintenanceByInstanceKey terminates maintenance mode for given instance
func (this *HttpAPI) EndMaintenanceByInstanceKey(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	err = inst.EndMaintenanceByInstanceKey(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Maintenance ended: %+v", instanceKey)})
}

// Maintenance provides list of instance under active maintenance
func (this *HttpAPI) Maintenance(params martini.Params, r render.Render, req *http.Request) {
	instanceKeys, err := inst.ReadActiveMaintenance()

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, instanceKeys)
}

// MoveUp attempts to move an instance up the topology
func (this *HttpAPI) MoveUp(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.MoveUp(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Instance moved up", Details: instance})
}

// MakeCoMaster attempts to make an instance co-master with its own master
func (this *HttpAPI) MakeCoMaster(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.MakeCoMaster(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Instance made co-master", Details: instance})
}

// ResetSlave makes a slave forget about its master, effectively breaking the replication
func (this *HttpAPI) ResetSlave(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.ResetSlaveOperation(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Slave reset", Details: instance})
}

// MoveBelow attempts to move an instance below its supposed sibling
func (this *HttpAPI) MoveBelow(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	siblingKey, err := this.getInstanceKey(params["siblingHost"], params["siblingPort"])
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.MoveBelow(&instanceKey, &siblingKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v moved below %+v", instanceKey, siblingKey), Details: instance})
}

// MatchBelow attempts to move an instance below another via pseudo GTID matching of binlog entries
func (this *HttpAPI) MatchBelow(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	belowKey, err := this.getInstanceKey(params["belowHost"], params["belowPort"])
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.MatchBelow(&instanceKey, &belowKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v matched below %+v", instanceKey, belowKey), Details: instance})
}

// MakeMaster attempts to make the given instance a master, and match its siblings to be its slaves
func (this *HttpAPI) MakeMaster(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	instance, err := inst.MakeMaster(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: fmt.Sprintf("Instance %+v now made master", instanceKey), Details: instance})
}

// StartSlave starts replication on given instance
func (this *HttpAPI) StartSlave(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.StartSlave(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Slave started", Details: instance})
}

// StopSlave stops replication on given instance
func (this *HttpAPI) StopSlave(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.StopSlave(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Slave stopped", Details: instance})
}

// StopSlaveNicely stops replication on given instance, such that sql thead is aligned with IO thread
func (this *HttpAPI) StopSlaveNicely(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.StopSlaveNicely(&instanceKey)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Slave stopped nicely", Details: instance})
}

// SetReadOnly sets the global read_only variable
func (this *HttpAPI) SetReadOnly(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.SetReadOnly(&instanceKey, true)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Server set as read-only", Details: instance})
}

// SetWriteable clear the global read_only variable
func (this *HttpAPI) SetWriteable(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.SetReadOnly(&instanceKey, false)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Server set as writeable", Details: instance})
}

// KillQuery kills a query running on a server
func (this *HttpAPI) KillQuery(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	instanceKey, err := this.getInstanceKey(params["host"], params["port"])
	processId, err := strconv.ParseInt(params["process"], 10, 0)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}
	instance, err := inst.KillQuery(&instanceKey, processId)
	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: err.Error()})
		return
	}

	r.JSON(200, &APIResponse{Code: OK, Message: "Slave stopped", Details: instance})
}

// Cluster provides list of instances in given cluster
func (this *HttpAPI) Cluster(params martini.Params, r render.Render, req *http.Request) {
	instances, err := inst.ReadClusterInstances(params["clusterName"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, instances)
}

// Clusters provides list of known clusters
func (this *HttpAPI) Clusters(params martini.Params, r render.Render, req *http.Request) {
	clusterNames, err := inst.ReadClusters()

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, clusterNames)
}

// ClustersInfo provides list of known clusters, along with some added metadata per cluster
func (this *HttpAPI) ClustersInfo(params martini.Params, r render.Render, req *http.Request) {
	clustersInfo, err := inst.ReadClustersInfo()

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, clustersInfo)
}

// Search provides list of instances matching given search param via various criteria.
func (this *HttpAPI) Search(params martini.Params, r render.Render, req *http.Request) {
	searchString := params["searchString"]
	if searchString == "" {
		searchString = req.URL.Query().Get("s")
	}
	instances, err := inst.SearchInstances(searchString)

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, instances)
}

// Problems provides list of instances with known problems
func (this *HttpAPI) Problems(params martini.Params, r render.Render, req *http.Request) {
	instances, err := inst.ReadProblemInstances()

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, instances)
}

// Audit provides list of audit entries by given page number
func (this *HttpAPI) Audit(params martini.Params, r render.Render, req *http.Request) {
	page, err := strconv.Atoi(params["page"])
	if err != nil || page < 0 {
		page = 0
	}
	audits, err := inst.ReadRecentAudit(page)

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, audits)
}

// LongQueries lists queries running for a long time, on all instances, optionally filtered by
// arbitrary text
func (this *HttpAPI) LongQueries(params martini.Params, r render.Render, req *http.Request) {
	longQueries, err := inst.ReadLongRunningProcesses(params["filter"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, longQueries)
}

// Agents provides complete list of registered agents (See https://github.com/outbrain/orchestrator-agent)
func (this *HttpAPI) Agents(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	agents, err := agent.ReadAgents()

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, agents)
}

// Agent returns complete information of a given agent
func (this *HttpAPI) Agent(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	agent, err := agent.GetAgent(params["host"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, agent)
}

// AgentUnmount instructs an agent to unmount the designated mount point
func (this *HttpAPI) AgentUnmount(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.Unmount(params["host"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentMountLV instructs an agent to mount a given volume on the designated mount point
func (this *HttpAPI) AgentMountLV(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.MountLV(params["host"], req.URL.Query().Get("lv"))

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentCreateSnapshot instructs an agent to create a new snapshot. Agent's DIY implementation.
func (this *HttpAPI) AgentCreateSnapshot(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.CreateSnapshot(params["host"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentRemoveLV instructs an agent to remove a logical volume
func (this *HttpAPI) AgentRemoveLV(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.RemoveLV(params["host"], req.URL.Query().Get("lv"))

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentMySQLStop stops MySQL service on agent
func (this *HttpAPI) AgentMySQLStop(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.MySQLStop(params["host"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentMySQLStart starts MySQL service on agent
func (this *HttpAPI) AgentMySQLStart(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.MySQLStart(params["host"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentSeed completely seeds a host with another host's snapshots. This is a complex operation
// governed by orchestrator and executed by the two agents involved.
func (this *HttpAPI) AgentSeed(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.Seed(params["targetHost"], params["sourceHost"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentActiveSeeds lists active seeds and their state
func (this *HttpAPI) AgentActiveSeeds(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.ReadActiveSeedsForHost(params["host"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentRecentSeeds lists recent seeds of a given agent
func (this *HttpAPI) AgentRecentSeeds(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.ReadRecentCompletedSeedsForHost(params["host"])

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentSeedDetails provides details of a given seed
func (this *HttpAPI) AgentSeedDetails(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	seedId, err := strconv.ParseInt(params["seedId"], 10, 0)
	output, err := agent.AgentSeedDetails(seedId)

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AgentSeedStates returns the breakdown of states (steps) of a given seed
func (this *HttpAPI) AgentSeedStates(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	seedId, err := strconv.ParseInt(params["seedId"], 10, 0)
	output, err := agent.ReadSeedStates(seedId)

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// Seeds retruns all recent seeds
func (this *HttpAPI) Seeds(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	output, err := agent.ReadRecentSeeds()

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, output)
}

// AbortSeed instructs agents to abort an active seed
func (this *HttpAPI) AbortSeed(params martini.Params, r render.Render, req *http.Request) {
	if !this.isAuthorizedForAction(req) {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Unauthorized"})
		return
	}
	if !config.Config.ServeAgentsHttp {
		r.JSON(200, &APIResponse{Code: ERROR, Message: "Agents not served"})
		return
	}

	seedId, err := strconv.ParseInt(params["seedId"], 10, 0)
	err = agent.AbortSeed(seedId)

	if err != nil {
		r.JSON(200, &APIResponse{Code: ERROR, Message: fmt.Sprintf("%+v", err)})
		return
	}

	r.JSON(200, err == nil)
}

// Headers is a self-test call which returns HTTP headers
func (this *HttpAPI) Headers(params martini.Params, r render.Render, req *http.Request) {
	r.JSON(200, req.Header)
}

// RegisterRequests makes for the de-facto list of known API calls
func (this *HttpAPI) RegisterRequests(m *martini.ClassicMartini) {
	m.Get("/api/instance/:host/:port", this.Instance)
	m.Get("/api/discover/:host/:port", this.Discover)
	m.Get("/api/refresh/:host/:port", this.Refresh)
	m.Get("/api/forget/:host/:port", this.Forget)
	m.Get("/api/resolve/:host/:port", this.Resolve)
	m.Get("/api/move-up/:host/:port", this.MoveUp)
	m.Get("/api/make-co-master/:host/:port", this.MakeCoMaster)
	m.Get("/api/reset-slave/:host/:port", this.ResetSlave)
	m.Get("/api/move-below/:host/:port/:siblingHost/:siblingPort", this.MoveBelow)
	m.Get("/api/match-below/:host/:port/:belowHost/:belowPort", this.MatchBelow)
	m.Get("/api/make-master/:host/:port", this.MakeMaster)
	m.Get("/api/begin-maintenance/:host/:port/:owner/:reason", this.BeginMaintenance)
	m.Get("/api/end-maintenance/:host/:port", this.EndMaintenanceByInstanceKey)
	m.Get("/api/end-maintenance/:maintenanceKey", this.EndMaintenance)
	m.Get("/api/start-slave/:host/:port", this.StartSlave)
	m.Get("/api/stop-slave/:host/:port", this.StopSlave)
	m.Get("/api/stop-slave-nice/:host/:port", this.StopSlaveNicely)
	m.Get("/api/set-read-only/:host/:port", this.SetReadOnly)
	m.Get("/api/set-writeable/:host/:port", this.SetWriteable)
	m.Get("/api/kill-query/:host/:port/:process", this.KillQuery)
	m.Get("/api/maintenance", this.Maintenance)
	m.Get("/api/cluster/:clusterName", this.Cluster)
	m.Get("/api/clusters", this.Clusters)
	m.Get("/api/clusters-info", this.ClustersInfo)
	m.Get("/api/search/:searchString", this.Search)
	m.Get("/api/search", this.Search)
	m.Get("/api/problems", this.Problems)
	m.Get("/api/long-queries", this.LongQueries)
	m.Get("/api/long-queries/:filter", this.LongQueries)
	m.Get("/api/audit", this.Audit)
	m.Get("/api/audit/:page", this.Audit)
	m.Get("/api/agents", this.Agents)
	m.Get("/api/agent/:host", this.Agent)
	m.Get("/api/agent-umount/:host", this.AgentUnmount)
	m.Get("/api/agent-mount/:host", this.AgentMountLV)
	m.Get("/api/agent-create-snapshot/:host", this.AgentCreateSnapshot)
	m.Get("/api/agent-removelv/:host", this.AgentRemoveLV)
	m.Get("/api/agent-mysql-stop/:host", this.AgentMySQLStop)
	m.Get("/api/agent-mysql-start/:host", this.AgentMySQLStart)
	m.Get("/api/agent-seed/:targetHost/:sourceHost", this.AgentSeed)
	m.Get("/api/agent-active-seeds/:host", this.AgentActiveSeeds)
	m.Get("/api/agent-recent-seeds/:host", this.AgentRecentSeeds)
	m.Get("/api/agent-seed-details/:seedId", this.AgentSeedDetails)
	m.Get("/api/agent-seed-states/:seedId", this.AgentSeedStates)
	m.Get("/api/agent-abort-seed/:seedId", this.AbortSeed)
	m.Get("/api/seeds", this.Seeds)
	m.Get("/api/headers", this.Headers)
}
