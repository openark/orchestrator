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
	"errors"
	"fmt"
	"github.com/outbrain/golib/log"
	"strings"
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

// InstanceIsMasterOf checks whether an instance is the master of another
func InstanceIsMasterOf(instance0, instance1 *Instance) bool {
	if !instance1.IsSlave() {
		return false
	}
	if instance0.Key.Equals(&instance1.Key) {
		// same instance...
		return false
	}
	return instance0.Key.Equals(&instance1.MasterKey)
}

// MoveUp will attempt moving instance indicated by instanceKey up the topology hierarchy.
// It will perform all safety and sanity checks and will tamper with this instance's replication
// as well as its master.
func MoveUp(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	rinstance, _, _ := ReadInstance(&instance.Key)
	if canMove, merr := rinstance.CanMove(); !canMove {
		return instance, merr
	}
	master, err := GetInstanceMaster(instance)
	if err != nil {
		return instance, log.Errorf("Cannot GetInstanceMaster() for %+v. error=%+v", instance, err)
	}

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
	if err != nil {
		goto Cleanup
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	instance, err = StartSlaveUntilMasterCoordinates(instanceKey, &master.SelfBinlogCoordinates)
	if err != nil {
		goto Cleanup
	}

	instance, err = ChangeMasterTo(instanceKey, &master.MasterKey, &master.ExecBinlogCoordinates)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)
	master, _ = StartSlave(&master.Key)
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("move-up", instanceKey, fmt.Sprintf("moved up %+v. Previous master: %+v", *instanceKey, master.Key))

	return instance, err
}

// MoveBelow will attempt moving instance indicated by instanceKey below its supposed sibling indicated by sinblingKey.
// It will perform all safety and sanity checks and will tamper with this instance's replication
// as well as its sibling.
func MoveBelow(instanceKey, siblingKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	sibling, err := ReadTopologyInstance(siblingKey)
	if err != nil {
		return instance, err
	}

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
	if err != nil {
		goto Cleanup
	}

	sibling, err = StopSlave(siblingKey)
	if err != nil {
		goto Cleanup
	}

	if instance.ExecBinlogCoordinates.SmallerThan(&sibling.ExecBinlogCoordinates) {
		instance, err = StartSlaveUntilMasterCoordinates(instanceKey, &sibling.ExecBinlogCoordinates)
		if err != nil {
			goto Cleanup
		}
	} else if sibling.ExecBinlogCoordinates.SmallerThan(&instance.ExecBinlogCoordinates) {
		sibling, err = StartSlaveUntilMasterCoordinates(siblingKey, &instance.ExecBinlogCoordinates)
		if err != nil {
			goto Cleanup
		}
	}
	// At this point both siblings have executed exact same statements and are identical

	instance, err = ChangeMasterTo(instanceKey, &sibling.Key, &sibling.SelfBinlogCoordinates)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)
	sibling, _ = StartSlave(siblingKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("move-below", instanceKey, fmt.Sprintf("moved %+v below %+v", *instanceKey, *siblingKey))

	return instance, err
}

// MakeCoMaster will attempt to make an instance co-master with its master, by making its master a slave of its own.
// This only works out if the master is not replicating; the master does not have a known master (it may have an unknown master).
func MakeCoMaster(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	master, err := GetInstanceMaster(instance)
	if err != nil {
		return instance, err
	}

	rinstance, _, _ := ReadInstance(&master.Key)
	if canMove, merr := rinstance.CanMoveAsCoMaster(); !canMove {
		return instance, merr
	}
	rinstance, _, _ = ReadInstance(instanceKey)
	if canMove, merr := rinstance.CanMove(); !canMove {
		return instance, merr
	}

	if instanceKey.Equals(&master.MasterKey) {
		return instance, errors.New(fmt.Sprintf("instance  %+v is already co master of %+v", instanceKey, master.Key))
	}
	if _, found, _ := ReadInstance(&master.MasterKey); found {
		return instance, errors.New(fmt.Sprintf("master %+v already has known master: %+v", master.Key, master.MasterKey))
	}
	if canReplicate, err := master.CanReplicateFrom(instance); !canReplicate {
		return instance, err
	}
	log.Infof("Will make %+v co-master of %+v", instanceKey, master.Key)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, "orchestrator", fmt.Sprintf("make co-master of %+v", master.Key)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(&master.Key, "orchestrator", fmt.Sprintf("%+v turns into co-master of this", *instanceKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", master.Key))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	// the coMaster used to be merely a slave. Just point master into *some* position
	// within coMaster...
	master, err = ChangeMasterTo(&master.Key, instanceKey, &instance.SelfBinlogCoordinates)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	master, _ = StartSlave(&master.Key)
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("make-co-master", instanceKey, fmt.Sprintf("%+v made co-master of %+v", *instanceKey, master.Key))

	return instance, err
}

// ResetSlaveOperation will reset a slave
func ResetSlaveOperation(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}

	log.Infof("Will reset %+v", instanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, "orchestrator", "reset slave"); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	if instance.IsSlave() {
		instance, err = StopSlave(instanceKey)
		if err != nil {
			goto Cleanup
		}
	}

	instance, err = ResetSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)

	if err != nil {
		return instance, log.Errore(err)
	}

	// and we're done (pending deferred functions)
	AuditOperation("reset slave", instanceKey, fmt.Sprintf("%+v replication reset", *instanceKey))

	return instance, err
}

// MatchBelow will attempt moving instance indicated by instanceKey below its the one indicated by otherKey.
// The refactoring is based on matching binlog entries, not on "classic" positions comparisons.
// The "other instance" could be the sibling of the moving instance any of its ancestors. It may actuall be
// a cousing of some sort (though unlikely). The only important thing is that the "other instance" is more
// advanced in replication than given instance.
func MatchBelow(instanceKey, otherKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	otherInstance, err := ReadTopologyInstance(otherKey)
	if err != nil {
		return instance, err
	}

	rinstance, _, _ := ReadInstance(&instance.Key)
	if canMove, merr := rinstance.CanMoveViaMatch(); !canMove {
		return instance, merr
	}

	if canReplicate, err := instance.CanReplicateFrom(otherInstance); !canReplicate {
		return instance, err
	}
	log.Infof("Will match %+v below %+v", *instanceKey, *otherKey)

	var instanceBinlogs []string
	var otherInstanceBinlogs []string
	var nextBinlogCoordinatestoMatch BinlogCoordinates
	var instancePseudoGtidCoordinates BinlogCoordinates
	var otherInstancePseudoGtidCoordinates BinlogCoordinates
	var instancePseudoGtidText string

	if maintenanceToken, merr := BeginMaintenance(instanceKey, "orchestrator", fmt.Sprintf("match below %+v", *otherKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(otherKey, "orchestrator", fmt.Sprintf("%+v matches below this", *instanceKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *otherKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	instance, err = StopSlaveNicely(instanceKey)
	if err != nil {
		goto Cleanup
	}

	{
		// Look for last GTID in instance:
		instanceBinlogs, err = GetInstanceBinaryLogNames(instanceKey)
		if err != nil {
			goto Cleanup
		}
		for i := len(instanceBinlogs) - 1; i >= 0 && instancePseudoGtidCoordinates.LogPos == 0; i-- {
			instancePseudoGtidCoordinates, instancePseudoGtidText, err = GetLastPseudoGTIDEntryInBinlog(instanceKey, instanceBinlogs[i])
		}
		if instancePseudoGtidCoordinates.LogPos == 0 {
			err = log.Errorf("Cannot find pseudo GTID entry in binlogs of %+v", *instanceKey)
			goto Cleanup
		}
		log.Debugf("Found pseudo gtid entry in %+v: %+v", *instanceKey, instancePseudoGtidCoordinates)
	}
	{
		// Look for GTID entry in other-instance:
		otherInstanceBinlogs, err = GetInstanceBinaryLogNames(otherKey)
		if err != nil {
			goto Cleanup
		}
		for i := len(otherInstanceBinlogs) - 1; i >= 0 && otherInstancePseudoGtidCoordinates.LogPos == 0; i-- {
			otherInstancePseudoGtidCoordinates, err = SearchPseudoGTIDEntryInBinlog(otherKey, otherInstanceBinlogs[i], instancePseudoGtidText)
		}
		if otherInstancePseudoGtidCoordinates.LogPos == 0 {
			err = log.Errorf("Cannot match pseudo GTID entry in binlogs of %+v", *otherKey)
			goto Cleanup
		}
		log.Debugf("Matched entry in %+v: %+v", *otherKey, otherInstancePseudoGtidCoordinates)
	}
	// We've found a match: the latest Pseudo GTID position within instance and its identical twin in otherInstance
	// We now iterate the events in both, up to the completion of events in instance (recall that we looked for
	// the last entry in instance, hence, assuming pseudo GTID entries are frequent, the amount of entries to read
	// from instance is not long)
	// The result of the iteration will be either:
	// - bad conclusion that instance is actually more advanced than otherInstance (we find more entries in instance
	//   following the pseudo gtid than we can match in otherInstance), hence we cannot ask instance to replicate
	//   from otherInstance
	// - good result: both instances are exactly in same shape (have replicated the exact same number of events since
	//   the last pseudo gtid). Since they are identical, it is easy to point instance into otherInstance.
	// - good result: the first position within otherInstance where instance has not replicated yet. It is easy to point
	//   instance into otherInstance.
	nextBinlogCoordinatesToMatch, err = GetNextBinlogCoordinatesToMatch(instanceKey, instanceBinlogs, instancePseudoGtidCoordinates,
		otherKey, otherInstanceBinlogs, otherInstancePseudoGtidCoordinates)
	if err != nil {
		goto Cleanup
	}
	log.Debugf("%+v will match below %+v at %+v", *instanceKey, *otherKey, nextBinlogCoordinatesToMatch)

Cleanup:
	instance, _ = StartSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("match-below", instanceKey, fmt.Sprintf("matched %+v below %+v", *instanceKey, *otherKey))

	return instance, err
}

// getAsciiTopologyEntry will get an ascii topology tree rooted at given instance. Ir recursively
// draws the tree
func getAsciiTopologyEntry(depth int, instance *Instance, replicationMap map[*Instance]([]*Instance)) []string {
	prefix := ""
	if depth > 0 {
		prefix = strings.Repeat(" ", (depth-1)*2)
		if instance.SlaveRunning() {
			prefix += "+ "
		} else {
			prefix += "- "
		}
	}
	entry := fmt.Sprintf("%s%s", prefix, instance.Key.DisplayString())
	result := []string{entry}
	for _, slave := range replicationMap[instance] {
		slavesResult := getAsciiTopologyEntry(depth+1, slave, replicationMap)
		result = append(result, slavesResult...)
	}
	return result
}

// AsciiTopology returns a string representation of the topology of given clusterName.
func AsciiTopology(clusterName string) (string, error) {
	instances, err := ReadClusterInstances(clusterName)
	if err != nil {
		return "", err
	}

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
