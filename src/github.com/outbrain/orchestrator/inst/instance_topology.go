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

package inst

import (
	"fmt"
	"errors"
	"strings"
	"github.com/outbrain/log"
)

// GetInstanceMaster synchronously reaches into the replication topology
// and retrieves master's data
func GetInstanceMaster(instance *Instance) (*Instance, error) {
	master, err := ReadTopologyInstance(&instance.MasterKey)
	return master, err
}


// InstancesAreSiblings checks whether both instances are replicating from same master
func InstancesAreSiblings(instance0, instance1 *Instance) bool {
	if !instance0.IsSlave() {
		return false
	}
	if !instance1.IsSlave() {
		return false
	}
	if instance0.Key.Equals(&instance1.Key) {
		// same instance...
		return false
	}
	return instance0.MasterKey.Equals(&instance1.MasterKey)
}

// MoveUp will attempt moving instance indicated by instanceKey up the topology hierarchy.
// It will perform all safety and sanity checks and will tamper with this instance's replication 
// as well as its master.
func MoveUp(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {	return instance, err}
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	rinstance, _, _ := ReadInstance(&instance.Key)
	if canMove, merr := rinstance.CanMove(); !canMove {
		return instance, merr
	}
	master, err := GetInstanceMaster(instance)
	if err != nil {	return instance, log.Errorf("Cannot GetInstanceMaster() for %+v. error=%+v", instance, err)}
	
	if !master.IsSlave() {
		return instance, errors.New(fmt.Sprintf("master is not a slave itself: %+v", master.Key))
	}
	
	if canReplicate, err := instance.CanReplicateFrom(master); canReplicate == false {
		return instance, err
	}
	
	log.Infof("Will move %+v up the topology", *instanceKey) 

	if maintenanceToken, merr := BeginMaintenance(instanceKey, "orchestrator", "move up"); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(&master.Key, "orchestrator", fmt.Sprintf("child %+v moves up", *instanceKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", master.Key))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	master, err = StopSlave(&master.Key)
	if	err	!=	nil	{goto Cleanup} 
	
	instance, err = StopSlave(instanceKey)
	if	err	!=	nil	{goto Cleanup} 
	
	instance, err = StartSlaveUntilMasterCoordinates(instanceKey, &master.SelfBinlogCoordinates)
	if	err	!=	nil	{goto Cleanup} 
	
	instance, err = ChangeMasterTo(instanceKey, &master.MasterKey, &master.ExecBinlogCoordinates)
	if	err	!=	nil	{goto Cleanup} 
	
	Cleanup:
	instance, _ = StartSlave(instanceKey)
	master, _ = StartSlave(&master.Key)
	if err != nil {	return instance, log.Errore(err)}
	// and we're done (pending deferred functions)
	AuditOperation("move-up", instanceKey, fmt.Sprintf("moved up %+v. Previous master: %+v", *instanceKey, master.Key))
	 
	return instance, err
}


// MoveUp will attempt moving instance indicated by instanceKey below its supposed sibling indicated by sinblingKey.
// It will perform all safety and sanity checks and will tamper with this instance's replication 
// as well as its sibling.
func MoveBelow(instanceKey, siblingKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {	return instance, err}
	sibling, err := ReadTopologyInstance(siblingKey)
	if err != nil {	return instance, err}

	rinstance, _, _ := ReadInstance(&instance.Key)
	if canMove, merr := rinstance.CanMove(); !canMove {
		return instance, merr
	}
	rinstance, _, _ = ReadInstance(&sibling.Key)
	if canMove, merr := rinstance.CanMove(); !canMove {
		return instance, merr
	}
	if !InstancesAreSiblings(instance, sibling) {
		return instance, errors.New(fmt.Sprintf("instances are not siblings: %+v, %+v", *instanceKey, *siblingKey))
	}
	
	if canReplicate, err := instance.CanReplicateFrom(sibling); !canReplicate {
		return instance, err
	}
	log.Infof("Will move %+v below its sibling %+v", instanceKey, siblingKey)
	
	if maintenanceToken, merr := BeginMaintenance(instanceKey, "orchestrator", fmt.Sprintf("move below %+v", *siblingKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(siblingKey, "orchestrator", fmt.Sprintf("%+v moves below this", *instanceKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *siblingKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	instance, err = StopSlave(instanceKey)
	if	err	!=	nil	{goto Cleanup} 
	
	sibling, err = StopSlave(siblingKey)
	if	err	!=	nil	{goto Cleanup} 
	
	if instance.ExecBinlogCoordinates.SmallerThan(&sibling.ExecBinlogCoordinates) {
		instance, err = StartSlaveUntilMasterCoordinates(instanceKey, &sibling.ExecBinlogCoordinates)
		if	err	!=	nil	{goto Cleanup} 
	} else if sibling.ExecBinlogCoordinates.SmallerThan(&instance.ExecBinlogCoordinates) {
		sibling, err = StartSlaveUntilMasterCoordinates(siblingKey, &instance.ExecBinlogCoordinates)
		if	err	!=	nil	{goto Cleanup} 
	}  
	// At this point both siblings have executed exact same statements and are identical
	 
	instance, err = ChangeMasterTo(instanceKey, &sibling.Key, &sibling.SelfBinlogCoordinates)
	if	err	!=	nil	{goto Cleanup} 
	
	
	Cleanup:
	instance, _ = StartSlave(instanceKey)
	sibling, _ = StartSlave(siblingKey)
	if err != nil {	return instance, log.Errore(err)}
	// and we're done (pending deferred functions)
	AuditOperation("move-below", instanceKey, fmt.Sprintf("moved %+v below %+v", *instanceKey, *siblingKey))	
	 
	return instance, err
}


// getAsciiTopologyEntry will get an ascii topology tree rooted at given instance. Ir recursively
// draws the tree 
func getAsciiTopologyEntry(depth int, instance *Instance, replicationMap map[*Instance]([]*Instance)) []string {
	prefix := ""
	if depth > 0 {
		prefix = strings.Repeat(" ", (depth - 1) * 2)
		if instance.SlaveRunning() {
			prefix += "+ "
		} else {
			prefix += "- "
		}
	}
	entry := fmt.Sprintf("%s%s", prefix, instance.Key.DisplayString()) 
	result := []string{entry}
	for _, slave := range replicationMap[instance] {
		slavesResult := getAsciiTopologyEntry(depth + 1, slave, replicationMap)
		result = append(result, slavesResult...)
	}
	return result
}

// AsciiTopology returns a string representation of the topology of given clusterName.
func AsciiTopology(clusterName string) (string, error) {
	instances, err := ReadClusterInstances(clusterName)
	if err != nil {return "", err} 

	instancesMap := make(map[InstanceKey](*Instance))
	for _, instance := range instances {
		log.Debugf("instanceKey: %+v", instance.Key)
		instancesMap[instance.Key] = instance
	}
	
	replicationMap := make(map[*Instance]([]*Instance))
	var masterInstance *Instance
	// Investigate slaves:
	for _, instance := range instances {
		master, ok := instancesMap[instance.MasterKey]
		if ok {
			if _, ok := replicationMap[master]; !ok {
				replicationMap[master] = [](*Instance){}
			}
			replicationMap[master] = append(replicationMap[master], instance)
		} else {
			masterInstance = instance
		} 
	}
	resultArray := getAsciiTopologyEntry(0, masterInstance, replicationMap)
	result := strings.Join(resultArray, "\n")
	return result, nil
}
