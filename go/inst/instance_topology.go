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
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/math"
	"github.com/outbrain/orchestrator/go/config"
)

// getASCIITopologyEntry will get an ascii topology tree rooted at given instance. Ir recursively
// draws the tree
func getASCIITopologyEntry(depth int, instance *Instance, replicationMap map[*Instance]([]*Instance), extendedOutput bool) []string {
	if instance == nil {
		return []string{}
	}
	if instance.IsCoMaster && depth > 1 {
		return []string{}
	}
	prefix := ""
	if depth > 0 {
		prefix = strings.Repeat(" ", (depth-1)*2)
		if instance.SlaveRunning() && instance.IsLastCheckValid && instance.IsRecentlyChecked {
			prefix += "+ "
		} else {
			prefix += "- "
		}
	}
	entry := fmt.Sprintf("%s%s", prefix, instance.Key.DisplayString())
	if extendedOutput {
		entry = fmt.Sprintf("%s %s", entry, instance.HumanReadableDescription())
	}
	result := []string{entry}
	for _, slave := range replicationMap[instance] {
		slavesResult := getASCIITopologyEntry(depth+1, slave, replicationMap, extendedOutput)
		result = append(result, slavesResult...)
	}
	return result
}

// ASCIITopology returns a string representation of the topology of given cluster.
func ASCIITopology(clusterName string, historyTimestampPattern string) (result string, err error) {
	var instances [](*Instance)
	if historyTimestampPattern == "" {
		instances, err = ReadClusterInstances(clusterName)
	} else {
		instances, err = ReadHistoryClusterInstances(clusterName, historyTimestampPattern)
	}
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
	// Get entries:
	var entries []string
	if masterInstance != nil {
		// Single master
		entries = getASCIITopologyEntry(0, masterInstance, replicationMap, historyTimestampPattern == "")
	} else {
		// Co-masters? For visualization we put each in its own branch while ignoring its other co-masters.
		for _, instance := range instances {
			if instance.IsCoMaster {
				entries = append(entries, getASCIITopologyEntry(1, instance, replicationMap, historyTimestampPattern == "")...)
			}
		}
	}
	// Beautify: make sure the "[...]" part is nicely aligned for all instances.
	{
		maxIndent := 0
		for _, entry := range entries {
			maxIndent = math.MaxInt(maxIndent, strings.Index(entry, "["))
		}
		for i, entry := range entries {
			entryIndent := strings.Index(entry, "[")
			if maxIndent > entryIndent {
				tokens := strings.Split(entry, "[")
				newEntry := fmt.Sprintf("%s%s[%s", tokens[0], strings.Repeat(" ", maxIndent-entryIndent), tokens[1])
				entries[i] = newEntry
			}
		}
	}
	// Turn into string
	result = strings.Join(entries, "\n")
	return result, nil
}

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
func InstanceIsMasterOf(allegedMaster, allegedSlave *Instance) bool {
	if !allegedSlave.IsSlave() {
		return false
	}
	if allegedMaster.Key.Equals(&allegedSlave.Key) {
		// same instance...
		return false
	}
	return allegedMaster.Key.Equals(&allegedSlave.MasterKey)
}

// MoveEquivalent will attempt moving instance indicated by instanceKey below another instance,
// based on known master coordinates equivalence
func MoveEquivalent(instanceKey, otherKey *InstanceKey) (*Instance, error) {
	instance, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return instance, err
	}
	if instance.Key.Equals(otherKey) {
		return instance, fmt.Errorf("MoveEquivalent: attempt to move an instance below itself %+v", instance.Key)
	}

	// Are there equivalent coordinates to this instance?
	instanceCoordinates := &InstanceBinlogCoordinates{Key: instance.MasterKey, Coordinates: instance.ExecBinlogCoordinates}
	binlogCoordinates, err := GetEquivalentBinlogCoordinatesFor(instanceCoordinates, otherKey)
	if err != nil {
		return instance, err
	}
	if binlogCoordinates == nil {
		return instance, fmt.Errorf("No equivalent coordinates found for %+v replicating from %+v at %+v", instance.Key, instance.MasterKey, instance.ExecBinlogCoordinates)
	}
	// For performance reasons, we did all the above before even checking the slave is stopped or stopping it at all.
	// This allows us to quickly skip the entire operation should there NOT be coordinates.
	// To elaborate: if the slave is actually running AND making progress, it is unlikely/impossible for it to have
	// equivalent coordinates, as the current coordinates are like to have never been seen.
	// This excludes the case, for example, that the master is itself not replicating.
	// Now if we DO get to happen on equivalent coordinates, we need to double check. For CHANGE MASTER to happen we must
	// stop the slave anyhow. But then let's verify the position hasn't changed.
	knownExecBinlogCoordinates := instance.ExecBinlogCoordinates
	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}
	if !instance.ExecBinlogCoordinates.Equals(&knownExecBinlogCoordinates) {
		// Seems like things were still running... We don't have an equivalence point
		err = fmt.Errorf("MoveEquivalent(): ExecBinlogCoordinates changed after stopping replication on %+v; aborting", instance.Key)
		goto Cleanup
	}
	instance, err = ChangeMasterTo(instanceKey, otherKey, binlogCoordinates, false, GTIDHintNeutral)

Cleanup:
	instance, _ = StartSlave(instanceKey)

	if err == nil {
		message := fmt.Sprintf("moved %+v via equivalence coordinates below %+v", *instanceKey, *otherKey)
		log.Debugf(message)
		AuditOperation("move-equivalent", instanceKey, message)
	}
	return instance, err
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
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}
	rinstance, _, _ := ReadInstance(&instance.Key)
	if canMove, merr := rinstance.CanMove(); !canMove {
		return instance, merr
	}
	master, err := GetInstanceMaster(instance)
	if err != nil {
		return instance, log.Errorf("Cannot GetInstanceMaster() for %+v. error=%+v", instance.Key, err)
	}

	if !master.IsSlave() {
		return instance, fmt.Errorf("master is not a slave itself: %+v", master.Key)
	}

	if canReplicate, err := instance.CanReplicateFrom(master); canReplicate == false {
		return instance, err
	}
	if master.IsBinlogServer() {
		// Quick solution via binlog servers
		return Repoint(instanceKey, &master.MasterKey, GTIDHintDeny)
	}

	log.Infof("Will move %+v up the topology", *instanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "move up"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(&master.Key, GetMaintenanceOwner(), fmt.Sprintf("child %+v moves up", *instanceKey)); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", master.Key)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	if !instance.UsingMariaDBGTID {
		master, err = StopSlave(&master.Key)
		if err != nil {
			goto Cleanup
		}
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	if !instance.UsingMariaDBGTID {
		instance, err = StartSlaveUntilMasterCoordinates(instanceKey, &master.SelfBinlogCoordinates)
		if err != nil {
			goto Cleanup
		}
	}

	// We can skip hostname unresolve; we just copy+paste whatever our master thinks of its master.
	instance, err = ChangeMasterTo(instanceKey, &master.MasterKey, &master.ExecBinlogCoordinates, true, GTIDHintDeny)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)
	if !instance.UsingMariaDBGTID {
		master, _ = StartSlave(&master.Key)
	}
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("move-up", instanceKey, fmt.Sprintf("moved up %+v. Previous master: %+v", *instanceKey, master.Key))

	return instance, err
}

// MoveUpSlaves will attempt moving up all slaves of a given instance, at the same time.
// Clock-time, this is fater than moving one at a time. However this means all slaves of the given instance, and the instance itself,
// will all stop replicating together.
func MoveUpSlaves(instanceKey *InstanceKey, pattern string) ([](*Instance), *Instance, error, []error) {
	res := [](*Instance){}
	errs := []error{}
	slaveMutex := make(chan bool, 1)
	var barrier chan *InstanceKey

	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return res, nil, err, errs
	}
	if !instance.IsSlave() {
		return res, instance, fmt.Errorf("instance is not a slave: %+v", instanceKey), errs
	}
	_, err = GetInstanceMaster(instance)
	if err != nil {
		return res, instance, log.Errorf("Cannot GetInstanceMaster() for %+v. error=%+v", instance.Key, err), errs
	}

	if instance.IsBinlogServer() {
		slaves, err, errors := RepointSlavesTo(instanceKey, pattern, &instance.MasterKey)
		// Bail out!
		return slaves, instance, err, errors
	}

	slaves, err := ReadSlaveInstances(instanceKey)
	if err != nil {
		return res, instance, err, errs
	}
	slaves = filterInstancesByPattern(slaves, pattern)
	if len(slaves) == 0 {
		return res, instance, nil, errs
	}
	log.Infof("Will move slaves of %+v up the topology", *instanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "move up slaves"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	for _, slave := range slaves {
		if maintenanceToken, merr := BeginMaintenance(&slave.Key, GetMaintenanceOwner(), fmt.Sprintf("%+v moves up", slave.Key)); merr != nil {
			err = fmt.Errorf("Cannot begin maintenance on %+v", slave.Key)
			goto Cleanup
		} else {
			defer EndMaintenance(maintenanceToken)
		}
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	barrier = make(chan *InstanceKey)
	for _, slave := range slaves {
		slave := slave
		go func() {
			defer func() {
				defer func() { barrier <- &slave.Key }()
				StartSlave(&slave.Key)
			}()

			var slaveErr error
			ExecuteOnTopology(func() {
				if canReplicate, err := slave.CanReplicateFrom(instance); canReplicate == false || err != nil {
					slaveErr = err
					return
				}
				if instance.IsBinlogServer() {
					// Special case. Just repoint
					slave, err = Repoint(&slave.Key, instanceKey, GTIDHintDeny)
					if err != nil {
						slaveErr = err
						return
					}
				} else {
					// Normal case. Do the math.
					slave, err = StopSlave(&slave.Key)
					if err != nil {
						slaveErr = err
						return
					}
					slave, err = StartSlaveUntilMasterCoordinates(&slave.Key, &instance.SelfBinlogCoordinates)
					if err != nil {
						slaveErr = err
						return
					}

					slave, err = ChangeMasterTo(&slave.Key, &instance.MasterKey, &instance.ExecBinlogCoordinates, false, GTIDHintDeny)
					if err != nil {
						slaveErr = err
						return
					}
				}
			})

			func() {
				slaveMutex <- true
				defer func() { <-slaveMutex }()
				if slaveErr == nil {
					res = append(res, slave)
				} else {
					errs = append(errs, slaveErr)
				}
			}()
		}()
	}
	for range slaves {
		<-barrier
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)
	if err != nil {
		return res, instance, log.Errore(err), errs
	}
	if len(errs) == len(slaves) {
		// All returned with error
		return res, instance, log.Error("Error on all operations"), errs
	}
	AuditOperation("move-up-slaves", instanceKey, fmt.Sprintf("moved up %d/%d slaves of %+v. New master: %+v", len(res), len(slaves), *instanceKey, instance.MasterKey))

	return res, instance, err, errs
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

	if sibling.IsBinlogServer() {
		// Binlog server has same coordinates as master
		// Easy solution!
		return Repoint(instanceKey, &sibling.Key, GTIDHintDeny)
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
		return instance, fmt.Errorf("instances are not siblings: %+v, %+v", *instanceKey, *siblingKey)
	}

	if canReplicate, err := instance.CanReplicateFrom(sibling); !canReplicate {
		return instance, err
	}
	log.Infof("Will move %+v below %+v", instanceKey, siblingKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("move below %+v", *siblingKey)); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(siblingKey, GetMaintenanceOwner(), fmt.Sprintf("%+v moves below this", *instanceKey)); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *siblingKey)
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

	instance, err = ChangeMasterTo(instanceKey, &sibling.Key, &sibling.SelfBinlogCoordinates, false, GTIDHintDeny)
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

func canMoveViaGTID(instance, otherInstance *Instance) (isOracleGTID bool, isMariaDBGTID, canMove bool) {
	isOracleGTID = (instance.UsingOracleGTID && otherInstance.SupportsOracleGTID)
	isMariaDBGTID = (instance.UsingMariaDBGTID && otherInstance.IsMariaDB())

	return isOracleGTID, isMariaDBGTID, isOracleGTID || isMariaDBGTID
}

// moveInstanceBelowViaGTID will attempt moving given instance below another instance using either Oracle GTID or MariaDB GTID.
func moveInstanceBelowViaGTID(instance, otherInstance *Instance) (*Instance, error) {
	_, _, canMove := canMoveViaGTID(instance, otherInstance)

	instanceKey := &instance.Key
	otherInstanceKey := &otherInstance.Key
	if !canMove {
		return instance, fmt.Errorf("Cannot move via GTID as not both instances use GTID: %+v, %+v", *instanceKey, *otherInstanceKey)
	}

	var err error

	rinstance, _, _ := ReadInstance(&instance.Key)
	if canMove, merr := rinstance.CanMoveViaMatch(); !canMove {
		return instance, merr
	}

	if canReplicate, err := instance.CanReplicateFrom(otherInstance); !canReplicate {
		return instance, err
	}
	log.Infof("Will move %+v below %+v via GTID", instanceKey, otherInstanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("move below %+v", *otherInstanceKey)); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	instance, err = ChangeMasterTo(instanceKey, &otherInstance.Key, &otherInstance.SelfBinlogCoordinates, false, GTIDHintForce)
	if err != nil {
		goto Cleanup
	}
Cleanup:
	instance, _ = StartSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("move-below-gtid", instanceKey, fmt.Sprintf("moved %+v below %+v", *instanceKey, *otherInstanceKey))

	return instance, err
}

// MoveBelowGTID will attempt moving instance indicated by instanceKey below another instance using either Oracle GTID or MariaDB GTID.
func MoveBelowGTID(instanceKey, otherKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	other, err := ReadTopologyInstance(otherKey)
	if err != nil {
		return instance, err
	}
	return moveInstanceBelowViaGTID(instance, other)
}

// moveSlavesViaGTID moves a list of slaves under another instance via GTID, returning those slaves
// that could not be moved (do not use GTID)
func moveSlavesViaGTID(slaves [](*Instance), other *Instance) (movedSlaves [](*Instance), unmovedSlaves [](*Instance), err error, errs []error) {
	slaves = RemoveInstance(slaves, &other.Key)
	if len(slaves) == 0 {
		// Nothing to do
		return movedSlaves, unmovedSlaves, nil, errs
	}

	log.Infof("Will move %+v slaves below %+v via GTID", len(slaves), other.Key)

	barrier := make(chan *InstanceKey)
	slaveMutex := make(chan bool, 1)
	for _, slave := range slaves {
		slave := slave

		// Parallelize repoints
		go func() {
			defer func() { barrier <- &slave.Key }()
			ExecuteOnTopology(func() {
				var slaveErr error
				if _, _, canMove := canMoveViaGTID(slave, other); canMove {
					slave, slaveErr = moveInstanceBelowViaGTID(slave, other)
				} else {
					slaveErr = fmt.Errorf("%+v cannot move below %+v via GTID", slave.Key, other.Key)
				}
				func() {
					// Instantaneous mutex.
					slaveMutex <- true
					defer func() { <-slaveMutex }()
					if slaveErr == nil {
						movedSlaves = append(movedSlaves, slave)
					} else {
						unmovedSlaves = append(unmovedSlaves, slave)
						errs = append(errs, slaveErr)
					}
				}()
			})
		}()
	}
	for range slaves {
		<-barrier
	}
	if len(errs) == len(slaves) {
		// All returned with error
		return movedSlaves, unmovedSlaves, fmt.Errorf("moveSlavesViaGTID: Error on all %+v operations", len(errs)), errs
	}
	AuditOperation("move-slaves-gtid", &other.Key, fmt.Sprintf("moved %d/%d slaves below %+v via GTID", len(movedSlaves), len(slaves), other.Key))

	return movedSlaves, unmovedSlaves, err, errs
}

// MoveSlavesGTID will (attempt to) move all slaves of given master below given instance.
func MoveSlavesGTID(masterKey *InstanceKey, belowKey *InstanceKey, pattern string) (movedSlaves [](*Instance), unmovedSlaves [](*Instance), err error, errs []error) {
	belowInstance, err := ReadTopologyInstance(belowKey)
	if err != nil {
		// Can't access "below" ==> can't move slaves beneath it
		return movedSlaves, unmovedSlaves, err, errs
	}

	// slaves involved
	slaves, err := ReadSlaveInstancesIncludingBinlogServerSubSlaves(masterKey)
	if err != nil {
		return movedSlaves, unmovedSlaves, err, errs
	}
	slaves = filterInstancesByPattern(slaves, pattern)
	movedSlaves, unmovedSlaves, err, errs = moveSlavesViaGTID(slaves, belowInstance)
	if err != nil {
		log.Errore(err)
	}

	if len(unmovedSlaves) > 0 {
		err = fmt.Errorf("MoveSlavesGTID: only moved %d out of %d slaves of %+v; error is: %+v", len(movedSlaves), len(slaves), *masterKey, err)
	}

	return movedSlaves, unmovedSlaves, err, errs
}

// Repoint connects a slave to a master using its exact same executing coordinates.
// The given masterKey can be null, in which case the existing master is used.
// Two use cases:
// - masterKey is nil: use case is corrupted relay logs on slave
// - masterKey is not nil: using Binlog servers (coordinates remain the same)
func Repoint(instanceKey *InstanceKey, masterKey *InstanceKey, gtidHint OperationGTIDHint) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", *instanceKey)
	}

	if masterKey == nil {
		masterKey = &instance.MasterKey
	}
	// With repoint we *prefer* the master to be alive, but we don't strictly require it.
	// The use case for the master being alive is with hostname-resolve or hostname-unresolve: asking the slave
	// to reconnect to its same master while changing the MASTER_HOST in CHANGE MASTER TO due to DNS changes etc.
	master, err := ReadTopologyInstance(masterKey)
	masterIsAccessible := (err == nil)
	if !masterIsAccessible {
		master, _, err = ReadInstance(masterKey)
		if err != nil {
			return instance, err
		}
	}
	if canReplicate, err := instance.CanReplicateFrom(master); !canReplicate {
		return instance, err
	}

	// if a binlog server check it is sufficiently up to date
	if master.IsBinlogServer() {
		// "Repoint" operation trusts the user. But only so much. Repoiting to a binlog server which is not yet there is strictly wrong.
		if !instance.ExecBinlogCoordinates.SmallerThanOrEquals(&master.SelfBinlogCoordinates) {
			return instance, fmt.Errorf("repoint: binlog server %+v is not sufficiently up to date to repoint %+v below it", *masterKey, *instanceKey)
		}
	}

	log.Infof("Will repoint %+v to master %+v", *instanceKey, *masterKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "repoint"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	// See above, we are relaxed about the master being accessible/inaccessible.
	// If accessible, we wish to do hostname-unresolve. If inaccessible, we can skip the test and not fail the
	// ChangeMasterTo operation. This is why we pass "!masterIsAccessible" below.
	if instance.ExecBinlogCoordinates.IsEmpty() {
		instance.ExecBinlogCoordinates.LogFile = "orchestrator-unknown-log-file"
	}
	instance, err = ChangeMasterTo(instanceKey, masterKey, &instance.ExecBinlogCoordinates, !masterIsAccessible, gtidHint)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("repoint", instanceKey, fmt.Sprintf("slave %+v repointed to master: %+v", *instanceKey, *masterKey))

	return instance, err

}

// RepointTo repoints list of slaves onto another master.
// Binlog Server is the major use case
func RepointTo(slaves [](*Instance), belowKey *InstanceKey) ([](*Instance), error, []error) {
	res := [](*Instance){}
	errs := []error{}

	slaves = RemoveInstance(slaves, belowKey)
	if len(slaves) == 0 {
		// Nothing to do
		return res, nil, errs
	}
	if belowKey == nil {
		return res, log.Errorf("RepointTo received nil belowKey"), errs
	}

	log.Infof("Will repoint %+v slaves below %+v", len(slaves), *belowKey)
	barrier := make(chan *InstanceKey)
	slaveMutex := make(chan bool, 1)
	for _, slave := range slaves {
		slave := slave

		// Parallelize repoints
		go func() {
			defer func() { barrier <- &slave.Key }()
			ExecuteOnTopology(func() {
				slave, slaveErr := Repoint(&slave.Key, belowKey, GTIDHintNeutral)

				func() {
					// Instantaneous mutex.
					slaveMutex <- true
					defer func() { <-slaveMutex }()
					if slaveErr == nil {
						res = append(res, slave)
					} else {
						errs = append(errs, slaveErr)
					}
				}()
			})
		}()
	}
	for range slaves {
		<-barrier
	}

	if len(errs) == len(slaves) {
		// All returned with error
		return res, log.Error("Error on all operations"), errs
	}
	AuditOperation("repoint-to", belowKey, fmt.Sprintf("repointed %d/%d slaves to %+v", len(res), len(slaves), *belowKey))

	return res, nil, errs
}

// RepointSlavesTo repoints slaves of a given instance (possibly filtered) onto another master.
// Binlog Server is the major use case
func RepointSlavesTo(instanceKey *InstanceKey, pattern string, belowKey *InstanceKey) ([](*Instance), error, []error) {
	res := [](*Instance){}
	errs := []error{}

	slaves, err := ReadSlaveInstances(instanceKey)
	if err != nil {
		return res, err, errs
	}
	slaves = RemoveInstance(slaves, belowKey)
	slaves = filterInstancesByPattern(slaves, pattern)
	if len(slaves) == 0 {
		// Nothing to do
		return res, nil, errs
	}
	if belowKey == nil {
		// Default to existing master. All slaves are of the same master, hence just pick one.
		belowKey = &slaves[0].MasterKey
	}
	log.Infof("Will repoint slaves of %+v to %+v", *instanceKey, *belowKey)
	return RepointTo(slaves, belowKey)
}

// RepointSlaves repoints all slaves of a given instance onto its existing master.
func RepointSlaves(instanceKey *InstanceKey, pattern string) ([](*Instance), error, []error) {
	return RepointSlavesTo(instanceKey, pattern, nil)
}

// MakeCoMaster will attempt to make an instance co-master with its master, by making its master a slave of its own.
// This only works out if the master is not replicating; the master does not have a known master (it may have an unknown master).
func MakeCoMaster(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if canMove, merr := instance.CanMove(); !canMove {
		return instance, merr
	}
	master, err := GetInstanceMaster(instance)
	if err != nil {
		return instance, err
	}
	log.Debugf("Will check whether %+v's master (%+v) can become its co-master", instance.Key, master.Key)
	if canMove, merr := master.CanMoveAsCoMaster(); !canMove {
		return instance, merr
	}
	if instanceKey.Equals(&master.MasterKey) {
		return instance, fmt.Errorf("instance %+v is already co master of %+v", instance.Key, master.Key)
	}
	if !instance.ReadOnly {
		return instance, fmt.Errorf("instance %+v is not read-only; first make it read-only before making it co-master", instance.Key)
	}
	if master.IsCoMaster {
		// We allow breaking of an existing co-master replication. Here's the breakdown:
		// Ideally, this would not eb allowed, and we would first require the user to RESET SLAVE on 'master'
		// prior to making it participate as co-master with our 'instance'.
		// However there's the problem that upon RESET SLAVE we lose the replication's user/password info.
		// Thus, we come up with the following rule:
		// If S replicates from M1, and M1<->M2 are co masters, we allow S to become co-master of M1 (S<->M1) if:
		// - M1 is writeable
		// - M2 is read-only or is unreachable/invalid
		// - S  is read-only
		// And so we will be replacing one read-only co-master with another.
		otherCoMaster, found, _ := ReadInstance(&master.MasterKey)
		if found && otherCoMaster.IsLastCheckValid && !otherCoMaster.ReadOnly {
			return instance, fmt.Errorf("master %+v is already co-master with %+v, and %+v is alive, and not read-only; cowardly refusing to demote it. Please set it as read-only beforehand", master.Key, otherCoMaster.Key, otherCoMaster.Key)
		}
		// OK, good to go.
	} else if _, found, _ := ReadInstance(&master.MasterKey); found {
		return instance, fmt.Errorf("%+v is not a real master; it replicates from: %+v", master.Key, master.MasterKey)
	}
	if canReplicate, err := master.CanReplicateFrom(instance); !canReplicate {
		return instance, err
	}
	log.Infof("Will make %+v co-master of %+v", instanceKey, master.Key)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("make co-master of %+v", master.Key)); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(&master.Key, GetMaintenanceOwner(), fmt.Sprintf("%+v turns into co-master of this", *instanceKey)); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", master.Key)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	// the coMaster used to be merely a slave. Just point master into *some* position
	// within coMaster...
	if master.IsSlave() {
		// this is the case of a co-master. For masters, the StopSlave operation throws an error, and
		// there's really no point in doing it.
		master, err = StopSlave(&master.Key)
		if err != nil {
			goto Cleanup
		}
	}
	if instance.ReplicationCredentialsAvailable && !master.HasReplicationCredentials {
		// Yay! We can get credentials from the slave!
		replicationUser, replicationPassword, err := ReadReplicationCredentials(&instance.Key)
		if err != nil {
			goto Cleanup
		}
		log.Debugf("Got credentials from a replica. will now apply")
		_, err = ChangeMasterCredentials(&master.Key, replicationUser, replicationPassword)
		if err != nil {
			goto Cleanup
		}
	}
	master, err = ChangeMasterTo(&master.Key, instanceKey, &instance.SelfBinlogCoordinates, false, GTIDHintNeutral)
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

	log.Infof("Will reset slave on %+v", instanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "reset slave"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
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
	AuditOperation("reset-slave", instanceKey, fmt.Sprintf("%+v replication reset", *instanceKey))

	return instance, err
}

// DetachSlaveOperation will detach a slave from its master by forcibly corrupting its replication coordinates
func DetachSlaveOperation(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}

	log.Infof("Will detach %+v", instanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "detach slave"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
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

	instance, err = DetachSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)

	if err != nil {
		return instance, log.Errore(err)
	}

	// and we're done (pending deferred functions)
	AuditOperation("detach-slave", instanceKey, fmt.Sprintf("%+v replication detached", *instanceKey))

	return instance, err
}

// ReattachSlaveOperation will detach a slave from its master by forcibly corrupting its replication coordinates
func ReattachSlaveOperation(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}

	log.Infof("Will reattach %+v", instanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "detach slave"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
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

	instance, err = ReattachSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)

	if err != nil {
		return instance, log.Errore(err)
	}

	// and we're done (pending deferred functions)
	AuditOperation("reattach-slave", instanceKey, fmt.Sprintf("%+v replication reattached", *instanceKey))

	return instance, err
}

// DetachSlaveMasterHost detaches a slave from its master by corrupting the Master_Host (in such way that is reversible)
func DetachSlaveMasterHost(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", *instanceKey)
	}
	if instance.MasterKey.IsDetached() {
		return instance, fmt.Errorf("instance already detached: %+v", *instanceKey)
	}
	detachedMasterKey := instance.MasterKey.DetachedKey()

	log.Infof("Will detach master host on %+v. Detached key is %+v", *instanceKey, *detachedMasterKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "detach-slave-master-host"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	instance, err = ChangeMasterTo(instanceKey, detachedMasterKey, &instance.ExecBinlogCoordinates, true, GTIDHintNeutral)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("repoint", instanceKey, fmt.Sprintf("slave %+v detached from master into %+v", *instanceKey, *detachedMasterKey))

	return instance, err
}

// ReattachSlaveMasterHost reattaches a slave back onto its master by undoing a DetachSlaveMasterHost operation
func ReattachSlaveMasterHost(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", *instanceKey)
	}
	if !instance.MasterKey.IsDetached() {
		return instance, fmt.Errorf("instance does not seem to be detached: %+v", *instanceKey)
	}

	reattachedMasterKey := instance.MasterKey.ReattachedKey()

	log.Infof("Will reattach master host on %+v. Reattached key is %+v", *instanceKey, *reattachedMasterKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "reattach-slave-master-host"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	instance, err = ChangeMasterTo(instanceKey, reattachedMasterKey, &instance.ExecBinlogCoordinates, true, GTIDHintNeutral)
	if err != nil {
		goto Cleanup
	}
	// Just in case this instance used to be a master:
	ReplaceAliasClusterName(instanceKey.StringCode(), reattachedMasterKey.StringCode())

Cleanup:
	instance, _ = StartSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("repoint", instanceKey, fmt.Sprintf("slave %+v reattached to master %+v", *instanceKey, *reattachedMasterKey))

	return instance, err
}

// EnableGTID will attempt to enable GTID-mode (either Oracle or MariaDB)
func EnableGTID(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if instance.UsingGTID() {
		return instance, fmt.Errorf("%+v already uses GTID", *instanceKey)
	}

	log.Infof("Will attempt to enable GTID on %+v", *instanceKey)

	instance, err = Repoint(instanceKey, nil, GTIDHintForce)
	if err != nil {
		return instance, err
	}
	if !instance.UsingGTID() {
		return instance, fmt.Errorf("Cannot enable GTID on %+v", *instanceKey)
	}

	AuditOperation("enable-gtid", instanceKey, fmt.Sprintf("enabled GTID on %+v", *instanceKey))

	return instance, err
}

// DisableGTID will attempt to disable GTID-mode (either Oracle or MariaDB) and revert to binlog file:pos replication
func DisableGTID(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if !instance.UsingGTID() {
		return instance, fmt.Errorf("%+v is not using GTID", *instanceKey)
	}

	log.Infof("Will attempt to disable GTID on %+v", *instanceKey)

	instance, err = Repoint(instanceKey, nil, GTIDHintDeny)
	if err != nil {
		return instance, err
	}
	if instance.UsingGTID() {
		return instance, fmt.Errorf("Cannot disable GTID on %+v", *instanceKey)
	}

	AuditOperation("disable-gtid", instanceKey, fmt.Sprintf("disabled GTID on %+v", *instanceKey))

	return instance, err
}

// ResetMasterGTIDOperation will issue a safe RESET MASTER on a slave that replicates via GTID:
// It will make sure the gtid_purged set matches the executed set value as read just before the RESET.
// this will enable new slaves to be attached to given instance without complaints about missing/purged entries.
// This function requires that the instance does not have slaves.
func ResetMasterGTIDOperation(instanceKey *InstanceKey, removeSelfUUID bool, uuidToRemove string) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if !instance.UsingOracleGTID {
		return instance, log.Errorf("reset-master-gtid requested for %+v but it is not using oracle-gtid", *instanceKey)
	}
	if len(instance.SlaveHosts) > 0 {
		return instance, log.Errorf("reset-master-gtid will not operate on %+v because it has %+v slaves. Expecting no slaves", *instanceKey, len(instance.SlaveHosts))
	}

	log.Infof("Will reset master on %+v", instanceKey)

	var oracleGtidSet *OracleGtidSet
	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "reset-master-gtid"); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
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

	oracleGtidSet, err = ParseGtidSet(instance.ExecutedGtidSet)
	if err != nil {
		goto Cleanup
	}
	if removeSelfUUID {
		uuidToRemove = instance.ServerUUID
	}
	if uuidToRemove != "" {
		removed := oracleGtidSet.RemoveUUID(uuidToRemove)
		if removed {
			log.Debugf("Will remove UUID %s", uuidToRemove)
		} else {
			log.Debugf("UUID %s not found", uuidToRemove)
		}
	}

	instance, err = ResetMaster(instanceKey)
	if err != nil {
		goto Cleanup
	}
	err = setGTIDPurged(instance, oracleGtidSet.String())
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)

	if err != nil {
		return instance, log.Errore(err)
	}

	// and we're done (pending deferred functions)
	AuditOperation("reset-master-gtid", instanceKey, fmt.Sprintf("%+v master reset", *instanceKey))

	return instance, err
}

// FindLastPseudoGTIDEntry will search an instance's binary logs or relay logs for the last pseudo-GTID entry,
// and return found coordinates as well as entry text
func FindLastPseudoGTIDEntry(instance *Instance, recordedInstanceRelayLogCoordinates BinlogCoordinates, maxBinlogCoordinates *BinlogCoordinates, exhaustiveSearch bool, expectedBinlogFormat *string) (instancePseudoGtidCoordinates *BinlogCoordinates, instancePseudoGtidText string, err error) {

	if config.Config.PseudoGTIDPattern == "" {
		return instancePseudoGtidCoordinates, instancePseudoGtidText, fmt.Errorf("PseudoGTIDPattern not configured; cannot use Pseudo-GTID")
	}

	minBinlogCoordinates, minRelaylogCoordinates, err := GetHeuristiclyRecentCoordinatesForInstance(&instance.Key)
	if instance.LogBinEnabled && instance.LogSlaveUpdatesEnabled && (expectedBinlogFormat == nil || instance.Binlog_format == *expectedBinlogFormat) {
		// Well no need to search this instance's binary logs if it doesn't have any...
		// With regard log-slave-updates, some edge cases are possible, like having this instance's log-slave-updates
		// enabled/disabled (of course having restarted it)
		// The approach is not to take chances. If log-slave-updates is disabled, fail and go for relay-logs.
		// If log-slave-updates was just enabled then possibly no pseudo-gtid is found, and so again we will go
		// for relay logs.
		// Also, if master has STATEMENT binlog format, and the slave has ROW binlog format, then comparing binlog entries would urely fail if based on the slave's binary logs.
		// Instead, we revert to the relay logs.
		instancePseudoGtidCoordinates, instancePseudoGtidText, err = getLastPseudoGTIDEntryInInstance(instance, minBinlogCoordinates, maxBinlogCoordinates, exhaustiveSearch)
	}
	if err != nil || instancePseudoGtidCoordinates == nil {
		// Unable to find pseudo GTID in binary logs.
		// Then MAYBE we are lucky enough (chances are we are, if this slave did not crash) that we can
		// extract the Pseudo GTID entry from the last (current) relay log file.
		instancePseudoGtidCoordinates, instancePseudoGtidText, err = getLastPseudoGTIDEntryInRelayLogs(instance, minRelaylogCoordinates, recordedInstanceRelayLogCoordinates, exhaustiveSearch)
	}
	return instancePseudoGtidCoordinates, instancePseudoGtidText, err
}

// CorrelateBinlogCoordinates find out, if possible, the binlog coordinates of given otherInstance that correlate
// with given coordinates of given instance.
func CorrelateBinlogCoordinates(instance *Instance, binlogCoordinates *BinlogCoordinates, otherInstance *Instance) (*BinlogCoordinates, int, error) {
	// We record the relay log coordinates just after the instance stopped since the coordinates can change upon
	// a FLUSH LOGS/FLUSH RELAY LOGS (or a START SLAVE, though that's an altogether different problem) etc.
	// We want to be on the safe side; we don't utterly trust that we are the only ones playing with the instance.
	recordedInstanceRelayLogCoordinates := instance.RelaylogCoordinates
	instancePseudoGtidCoordinates, instancePseudoGtidText, err := FindLastPseudoGTIDEntry(instance, recordedInstanceRelayLogCoordinates, binlogCoordinates, true, &otherInstance.Binlog_format)

	if err != nil {
		return nil, 0, err
	}
	entriesMonotonic := (config.Config.PseudoGTIDMonotonicHint != "") && strings.Contains(instancePseudoGtidText, config.Config.PseudoGTIDMonotonicHint)
	minBinlogCoordinates, _, err := GetHeuristiclyRecentCoordinatesForInstance(&otherInstance.Key)
	otherInstancePseudoGtidCoordinates, err := SearchEntryInInstanceBinlogs(otherInstance, instancePseudoGtidText, entriesMonotonic, minBinlogCoordinates)
	if err != nil {
		return nil, 0, err
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
	nextBinlogCoordinatesToMatch, countMatchedEvents, err := GetNextBinlogCoordinatesToMatch(instance, *instancePseudoGtidCoordinates,
		recordedInstanceRelayLogCoordinates, binlogCoordinates, otherInstance, *otherInstancePseudoGtidCoordinates)
	if err != nil {
		return nil, 0, err
	}
	if countMatchedEvents == 0 {
		err = fmt.Errorf("Unexpected: 0 events processed while iterating logs. Something went wrong; aborting. nextBinlogCoordinatesToMatch: %+v", nextBinlogCoordinatesToMatch)
		return nil, 0, err
	}
	return nextBinlogCoordinatesToMatch, countMatchedEvents, nil
}

// MatchBelow will attempt moving instance indicated by instanceKey below its the one indicated by otherKey.
// The refactoring is based on matching binlog entries, not on "classic" positions comparisons.
// The "other instance" could be the sibling of the moving instance any of its ancestors. It may actually be
// a cousin of some sort (though unlikely). The only important thing is that the "other instance" is more
// advanced in replication than given instance.
func MatchBelow(instanceKey, otherKey *InstanceKey, requireInstanceMaintenance bool) (*Instance, *BinlogCoordinates, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, nil, err
	}
	if config.Config.PseudoGTIDPattern == "" {
		return instance, nil, fmt.Errorf("PseudoGTIDPattern not configured; cannot use Pseudo-GTID")
	}
	if instanceKey.Equals(otherKey) {
		return instance, nil, fmt.Errorf("MatchBelow: attempt to match an instance below itself %+v", *instanceKey)
	}
	otherInstance, err := ReadTopologyInstance(otherKey)
	if err != nil {
		return instance, nil, err
	}

	rinstance, _, _ := ReadInstance(&instance.Key)
	if canMove, merr := rinstance.CanMoveViaMatch(); !canMove {
		return instance, nil, merr
	}

	if canReplicate, err := instance.CanReplicateFrom(otherInstance); !canReplicate {
		return instance, nil, err
	}
	var nextBinlogCoordinatesToMatch *BinlogCoordinates
	var countMatchedEvents int

	if otherInstance.IsBinlogServer() {
		// A Binlog Server does not do all the SHOW BINLOG EVENTS stuff
		err = fmt.Errorf("Cannot use PseudoGTID with Binlog Server %+v", otherInstance.Key)
		goto Cleanup
	}

	log.Infof("Will match %+v below %+v", *instanceKey, *otherKey)

	if requireInstanceMaintenance {
		if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("match below %+v", *otherKey)); merr != nil {
			err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
			goto Cleanup
		} else {
			defer EndMaintenance(maintenanceToken)
		}
	}

	log.Debugf("Stopping slave on %+v", *instanceKey)
	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	nextBinlogCoordinatesToMatch, countMatchedEvents, err = CorrelateBinlogCoordinates(instance, nil, otherInstance)

	if countMatchedEvents == 0 {
		err = fmt.Errorf("Unexpected: 0 events processed while iterating logs. Something went wrong; aborting. nextBinlogCoordinatesToMatch: %+v", nextBinlogCoordinatesToMatch)
		goto Cleanup
	}
	log.Debugf("%+v will match below %+v at %+v; validated events: %d", *instanceKey, *otherKey, *nextBinlogCoordinatesToMatch, countMatchedEvents)

	// Drum roll......
	instance, err = ChangeMasterTo(instanceKey, otherKey, nextBinlogCoordinatesToMatch, false, GTIDHintDeny)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	instance, _ = StartSlave(instanceKey)
	if err != nil {
		return instance, nextBinlogCoordinatesToMatch, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("match-below", instanceKey, fmt.Sprintf("matched %+v below %+v", *instanceKey, *otherKey))

	return instance, nextBinlogCoordinatesToMatch, err
}

// RematchSlave will re-match a slave to its master, using pseudo-gtid
func RematchSlave(instanceKey *InstanceKey, requireInstanceMaintenance bool) (*Instance, *BinlogCoordinates, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, nil, err
	}
	masterInstance, found, err := ReadInstance(&instance.MasterKey)
	if err != nil || !found {
		return instance, nil, err
	}
	return MatchBelow(instanceKey, &masterInstance.Key, requireInstanceMaintenance)
}

// MakeMaster will take an instance, make all its siblings its slaves (via pseudo-GTID) and make it master
// (stop its replicaiton, make writeable).
func MakeMaster(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	masterInstance, err := ReadTopologyInstance(&instance.MasterKey)
	if err == nil {
		// If the read succeeded, check the master status.
		if masterInstance.IsSlave() {
			return instance, fmt.Errorf("MakeMaster: instance's master %+v seems to be replicating", masterInstance.Key)
		}
		if masterInstance.IsLastCheckValid {
			return instance, fmt.Errorf("MakeMaster: instance's master %+v seems to be accessible", masterInstance.Key)
		}
	}
	// Continue anyway if the read failed, because that means the master is
	// inaccessible... So it's OK to do the promotion.
	if !instance.SQLThreadUpToDate() {
		return instance, fmt.Errorf("MakeMaster: instance's SQL thread must be up-to-date with I/O thread for %+v", *instanceKey)
	}
	siblings, err := ReadSlaveInstances(&masterInstance.Key)
	if err != nil {
		return instance, err
	}
	for _, sibling := range siblings {
		if instance.ExecBinlogCoordinates.SmallerThan(&sibling.ExecBinlogCoordinates) {
			return instance, fmt.Errorf("MakeMaster: instance %+v has more advanced sibling: %+v", *instanceKey, sibling.Key)
		}
	}

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("siblings match below this: %+v", *instanceKey)); merr != nil {
		err = fmt.Errorf("Cannot begin maintenance on %+v", *instanceKey)
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	_, _, err, _ = MultiMatchBelow(siblings, instanceKey, false, nil)
	if err != nil {
		goto Cleanup
	}

	SetReadOnly(instanceKey, false)

Cleanup:
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("make-master", instanceKey, fmt.Sprintf("made master of %+v", *instanceKey))

	return instance, err
}

// EnslaveSiblings is a convenience method for turning sublings of a slave to be its subordinates.
// This uses normal connected replication (does not utilize Pseudo-GTID)
func EnslaveSiblings(instanceKey *InstanceKey) (*Instance, int, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, 0, err
	}
	masterInstance, found, err := ReadInstance(&instance.MasterKey)
	if err != nil || !found {
		return instance, 0, err
	}
	siblings, err := ReadSlaveInstances(&masterInstance.Key)
	if err != nil {
		return instance, 0, err
	}
	enslavedSiblings := 0
	for _, sibling := range siblings {
		if _, err := MoveBelow(&sibling.Key, &instance.Key); err == nil {
			enslavedSiblings++
		}
	}

	return instance, enslavedSiblings, err
}

// EnslaveMaster will move an instance up the chain and cause its master to become its slave.
// It's almost a role change, just that other slaves of either 'instance' or its master are currently unaffected
// (they continue replicate without change)
// Note that the master must itself be a slave; however the grandparent does not necessarily have to be reachable
// and can in fact be dead.
func EnslaveMaster(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	masterInstance, found, err := ReadInstance(&instance.MasterKey)
	if err != nil || !found {
		return instance, err
	}
	log.Debugf("EnslaveMaster: will attempt making %+v enslave its master %+v, now resolved as %+v", *instanceKey, instance.MasterKey, masterInstance.Key)

	if canReplicate, err := masterInstance.CanReplicateFrom(instance); canReplicate == false {
		return instance, err
	}
	// We begin
	masterInstance, err = StopSlave(&masterInstance.Key)
	if err != nil {
		goto Cleanup
	}
	instance, err = StopSlave(&instance.Key)
	if err != nil {
		goto Cleanup
	}

	instance, err = StartSlaveUntilMasterCoordinates(&instance.Key, &masterInstance.SelfBinlogCoordinates)
	if err != nil {
		goto Cleanup
	}

	// instance and masterInstance are equal
	// We skip name unresolve. It is OK if the master's master is dead, unreachable, does not resolve properly.
	// We just copy+paste info from the master.
	// In particular, this is commonly calledin DeadMaster recovery
	instance, err = ChangeMasterTo(&instance.Key, &masterInstance.MasterKey, &masterInstance.ExecBinlogCoordinates, true, GTIDHintNeutral)
	if err != nil {
		goto Cleanup
	}
	// instance is now sibling of master
	masterInstance, err = ChangeMasterTo(&masterInstance.Key, &instance.Key, &instance.SelfBinlogCoordinates, false, GTIDHintNeutral)
	if err != nil {
		goto Cleanup
	}
	// swap is done!

Cleanup:
	instance, _ = StartSlave(&instance.Key)
	masterInstance, _ = StartSlave(&masterInstance.Key)
	if err != nil {
		return instance, err
	}
	AuditOperation("enslave-master", instanceKey, fmt.Sprintf("enslaved master: %+v", masterInstance.Key))

	return instance, err
}

// MakeLocalMaster promotes a slave above its master, making it slave of its grandparent, while also enslaving its siblings.
// This serves as a convenience method to recover replication when a local master fails; the instance promoted is one of its slaves,
// which is most advanced among its siblings.
// This method utilizes Pseudo GTID
func MakeLocalMaster(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	masterInstance, found, err := ReadInstance(&instance.MasterKey)
	if err != nil || !found {
		return instance, err
	}
	grandparentInstance, err := ReadTopologyInstance(&masterInstance.MasterKey)
	if err != nil {
		return instance, err
	}
	siblings, err := ReadSlaveInstances(&masterInstance.Key)
	if err != nil {
		return instance, err
	}
	for _, sibling := range siblings {
		if instance.ExecBinlogCoordinates.SmallerThan(&sibling.ExecBinlogCoordinates) {
			return instance, fmt.Errorf("MakeMaster: instance %+v has more advanced sibling: %+v", *instanceKey, sibling.Key)
		}
	}

	instance, err = StopSlaveNicely(instanceKey, 0)
	if err != nil {
		goto Cleanup
	}

	_, _, err = MatchBelow(instanceKey, &grandparentInstance.Key, true)
	if err != nil {
		goto Cleanup
	}

	_, _, err, _ = MultiMatchBelow(siblings, instanceKey, false, nil)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	AuditOperation("make-local-master", instanceKey, fmt.Sprintf("made master of %+v", *instanceKey))

	return instance, err
}

// sortInstances shuffles given list of instances according to some logic
func sortInstances(instances [](*Instance)) {
	sort.Sort(sort.Reverse(InstancesByExecBinlogCoordinates(instances)))
}

// getSlavesForSorting returns a list of slaves of a given master potentially for candidate choosing
func getSlavesForSorting(masterKey *InstanceKey, includeBinlogServerSubSlaves bool) (slaves [](*Instance), err error) {
	if includeBinlogServerSubSlaves {
		slaves, err = ReadSlaveInstancesIncludingBinlogServerSubSlaves(masterKey)
	} else {
		slaves, err = ReadSlaveInstances(masterKey)
	}
	return slaves, err
}

// sortedSlaves returns the list of slaves of some master, sorted by exec coordinates
// (most up-to-date slave first).
// This function assumes given `slaves` argument is indeed a list of instances all replicating
// from the same master (the result of `getSlavesForSorting()` is appropriate)
func sortedSlaves(slaves [](*Instance), shouldStopSlaves bool) [](*Instance) {
	if len(slaves) == 0 {
		return slaves
	}
	if shouldStopSlaves {
		log.Debugf("sortedSlaves: stopping %d slaves nicely", len(slaves))
		slaves = StopSlavesNicely(slaves, time.Duration(config.Config.InstanceBulkOperationsWaitTimeoutSeconds)*time.Second)
	}
	slaves = RemoveNilInstances(slaves)

	sortInstances(slaves)
	for _, slave := range slaves {
		log.Debugf("- sorted slave: %+v %+v", slave.Key, slave.ExecBinlogCoordinates)
	}

	return slaves
}

// MultiMatchBelow will efficiently match multiple slaves below a given instance.
// It is assumed that all given slaves are siblings
func MultiMatchBelow(slaves [](*Instance), belowKey *InstanceKey, slavesAlreadyStopped bool, postponedFunctionsContainer *PostponedFunctionsContainer) ([](*Instance), *Instance, error, []error) {
	res := [](*Instance){}
	errs := []error{}
	slaveMutex := make(chan bool, 1)

	if config.Config.PseudoGTIDPattern == "" {
		return res, nil, fmt.Errorf("PseudoGTIDPattern not configured; cannot use Pseudo-GTID"), errs
	}

	slaves = RemoveInstance(slaves, belowKey)
	slaves = RemoveBinlogServerInstances(slaves)

	for _, slave := range slaves {
		if maintenanceToken, merr := BeginMaintenance(&slave.Key, GetMaintenanceOwner(), fmt.Sprintf("%+v match below %+v as part of MultiMatchBelow", slave.Key, *belowKey)); merr != nil {
			errs = append(errs, fmt.Errorf("Cannot begin maintenance on %+v", slave.Key))
			slaves = RemoveInstance(slaves, &slave.Key)
		} else {
			defer EndMaintenance(maintenanceToken)
		}
	}

	belowInstance, err := ReadTopologyInstance(belowKey)
	if err != nil {
		// Can't access the server below which we need to match ==> can't move slaves
		return res, belowInstance, err, errs
	}
	if belowInstance.IsBinlogServer() {
		// A Binlog Server does not do all the SHOW BINLOG EVENTS stuff
		err = fmt.Errorf("Cannot use PseudoGTID with Binlog Server %+v", belowInstance.Key)
		return res, belowInstance, err, errs
	}

	// slaves involved
	if len(slaves) == 0 {
		return res, belowInstance, nil, errs
	}
	if !slavesAlreadyStopped {
		log.Debugf("MultiMatchBelow: stopping %d slaves nicely", len(slaves))
		// We want the slaves to have SQL thread up to date with IO thread.
		// We will wait for them (up to a timeout) to do so.
		slaves = StopSlavesNicely(slaves, time.Duration(config.Config.InstanceBulkOperationsWaitTimeoutSeconds)*time.Second)
	}
	slaves = RemoveNilInstances(slaves)
	sort.Sort(sort.Reverse(InstancesByExecBinlogCoordinates(slaves)))

	// Optimizations:
	// Slaves which broke on the same Exec-coordinates can be handled in the exact same way:
	// we only need to figure out one slave of each group/bucket of exec-coordinates; then apply the CHANGE MASTER TO
	// on all its fellow members using same coordinates.
	slaveBuckets := make(map[BinlogCoordinates][](*Instance))
	for _, slave := range slaves {
		slave := slave
		slaveBuckets[slave.ExecBinlogCoordinates] = append(slaveBuckets[slave.ExecBinlogCoordinates], slave)
	}
	log.Debugf("MultiMatchBelow: %d slaves merged into %d buckets", len(slaves), len(slaveBuckets))
	for bucket, bucketSlaves := range slaveBuckets {
		log.Debugf("+- bucket: %+v, %d slaves", bucket, len(bucketSlaves))
	}
	matchedSlaves := make(map[InstanceKey]bool)
	bucketsBarrier := make(chan *BinlogCoordinates)
	// Now go over the buckets, and try a single slave from each bucket
	// (though if one results with an error, synchronuously-for-that-bucket continue to the next slave in bucket)

	for execCoordinates, bucketSlaves := range slaveBuckets {
		execCoordinates := execCoordinates
		bucketSlaves := bucketSlaves
		var bucketMatchedCoordinates *BinlogCoordinates
		// Buckets concurrent
		go func() {
			// find coordinates for a single bucket based on a slave in said bucket
			defer func() { bucketsBarrier <- &execCoordinates }()
			func() {
				for _, slave := range bucketSlaves {
					slave := slave
					var slaveErr error
					var matchedCoordinates *BinlogCoordinates
					log.Debugf("MultiMatchBelow: attempting slave %+v in bucket %+v", slave.Key, execCoordinates)
					matchFunc := func() error {
						ExecuteOnTopology(func() {
							_, matchedCoordinates, slaveErr = MatchBelow(&slave.Key, &belowInstance.Key, false)
						})
						return nil
					}
					if postponedFunctionsContainer != nil &&
						config.Config.PostponeSlaveRecoveryOnLagMinutes > 0 &&
						slave.SQLDelay > config.Config.PostponeSlaveRecoveryOnLagMinutes*60 &&
						len(bucketSlaves) == 1 {
						// This slave is the only one in the bucket, AND it's lagging very much, AND
						// we're configured to postpone operation on this slave so as not to delay everyone else.
						(*postponedFunctionsContainer).AddPostponedFunction(matchFunc)
						return
						// We bail out and trust our invoker to later call upon this postponed function
					}
					matchFunc()
					log.Debugf("MultiMatchBelow: match result: %+v, %+v", matchedCoordinates, slaveErr)

					if slaveErr == nil {
						// Success! We matched a slave of this bucket
						func() {
							// Instantaneous mutex.
							slaveMutex <- true
							defer func() { <-slaveMutex }()
							bucketMatchedCoordinates = matchedCoordinates
							matchedSlaves[slave.Key] = true
						}()
						log.Debugf("MultiMatchBelow: matched slave %+v in bucket %+v", slave.Key, execCoordinates)
						return
					}

					// Got here? Error!
					func() {
						// Instantaneous mutex.
						slaveMutex <- true
						defer func() { <-slaveMutex }()
						errs = append(errs, slaveErr)
					}()
					log.Errore(slaveErr)
					// Failure: some unknown problem with bucket slave. Let's try the next one (continue loop)
				}
			}()
			if bucketMatchedCoordinates == nil {
				log.Errorf("MultiMatchBelow: Cannot match up %d slaves since their bucket %+v is failed", len(bucketSlaves), execCoordinates)
				return
			}
			log.Debugf("MultiMatchBelow: bucket %+v coordinates are: %+v. Proceeding to match all bucket slaves", execCoordinates, *bucketMatchedCoordinates)
			// At this point our bucket has a known salvaged slave.
			// We don't wait for the other buckets -- we immediately work out all the other slaves in this bucket.
			// (perhaps another bucket is busy matching a 24h delayed-replica; we definitely don't want to hold on that)
			func() {
				barrier := make(chan *InstanceKey)
				// We point all this bucket's slaves into the same coordinates, concurrently
				// We are already doing concurrent buckets; but for each bucket we also want to do concurrent slaves,
				// otherwise one large bucket would make for a sequential work...
				for _, slave := range bucketSlaves {
					slave := slave
					go func() {
						defer func() { barrier <- &slave.Key }()

						var err error
						if _, found := matchedSlaves[slave.Key]; found {
							// Already matched this slave
							return
						}
						log.Debugf("MultiMatchBelow: Will match up %+v to previously matched master coordinates %+v", slave.Key, *bucketMatchedCoordinates)
						slaveMatchSuccess := false
						ExecuteOnTopology(func() {
							if _, err = ChangeMasterTo(&slave.Key, &belowInstance.Key, bucketMatchedCoordinates, false, GTIDHintDeny); err == nil {
								StartSlave(&slave.Key)
								slaveMatchSuccess = true
							}
						})
						func() {
							// Quickly update lists; mutext is instantenous
							slaveMutex <- true
							defer func() { <-slaveMutex }()
							if slaveMatchSuccess {
								matchedSlaves[slave.Key] = true
							} else {
								errs = append(errs, err)
								log.Errorf("MultiMatchBelow: Cannot match up %+v: error is %+v", slave.Key, err)
							}
						}()
					}()
				}
				for range bucketSlaves {
					<-barrier
				}
			}()
		}()
	}
	for range slaveBuckets {
		<-bucketsBarrier
	}

	for _, slave := range slaves {
		slave := slave
		if _, found := matchedSlaves[slave.Key]; found {
			res = append(res, slave)
		}
	}
	return res, belowInstance, err, errs
}

// MultiMatchSlaves will match (via pseudo-gtid) all slaves of given master below given instance.
func MultiMatchSlaves(masterKey *InstanceKey, belowKey *InstanceKey, pattern string) ([](*Instance), *Instance, error, []error) {
	res := [](*Instance){}
	errs := []error{}

	belowInstance, err := ReadTopologyInstance(belowKey)
	if err != nil {
		// Can't access "below" ==> can't match slaves beneath it
		return res, nil, err, errs
	}

	masterInstance, found, err := ReadInstance(masterKey)
	if err != nil || !found {
		return res, nil, err, errs
	}

	// See if we have a binlog server case (special handling):
	binlogCase := false
	if masterInstance.IsBinlogServer() && masterInstance.MasterKey.Equals(belowKey) {
		// repoint-up
		log.Debugf("MultiMatchSlaves: pointing slaves up from binlog server")
		binlogCase = true
	} else if belowInstance.IsBinlogServer() && belowInstance.MasterKey.Equals(masterKey) {
		// repoint-down
		log.Debugf("MultiMatchSlaves: pointing slaves down to binlog server")
		binlogCase = true
	} else if masterInstance.IsBinlogServer() && belowInstance.IsBinlogServer() && masterInstance.MasterKey.Equals(&belowInstance.MasterKey) {
		// Both BLS, siblings
		log.Debugf("MultiMatchSlaves: pointing slaves to binlong sibling")
		binlogCase = true
	}
	if binlogCase {
		slaves, err, errors := RepointSlavesTo(masterKey, pattern, belowKey)
		// Bail out!
		return slaves, masterInstance, err, errors
	}

	// Not binlog server

	// slaves involved
	slaves, err := ReadSlaveInstancesIncludingBinlogServerSubSlaves(masterKey)
	if err != nil {
		return res, belowInstance, err, errs
	}
	slaves = filterInstancesByPattern(slaves, pattern)
	matchedSlaves, belowInstance, err, errs := MultiMatchBelow(slaves, &belowInstance.Key, false, nil)

	if len(matchedSlaves) != len(slaves) {
		err = fmt.Errorf("MultiMatchSlaves: only matched %d out of %d slaves of %+v; error is: %+v", len(matchedSlaves), len(slaves), *masterKey, err)
	}
	AuditOperation("multi-match-slaves", masterKey, fmt.Sprintf("matched %d slaves under %+v", len(matchedSlaves), *belowKey))

	return matchedSlaves, belowInstance, err, errs
}

// MatchUp will move a slave up the replication chain, so that it becomes sibling of its master, via Pseudo-GTID
func MatchUp(instanceKey *InstanceKey, requireInstanceMaintenance bool) (*Instance, *BinlogCoordinates, error) {
	instance, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return nil, nil, err
	}
	if !instance.IsSlave() {
		return instance, nil, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}
	master, found, err := ReadInstance(&instance.MasterKey)
	if err != nil || !found {
		return instance, nil, log.Errorf("Cannot get master for %+v. error=%+v", instance.Key, err)
	}

	if !master.IsSlave() {
		return instance, nil, fmt.Errorf("master is not a slave itself: %+v", master.Key)
	}

	return MatchBelow(instanceKey, &master.MasterKey, requireInstanceMaintenance)
}

// MatchUpSlaves will move all slaves of given master up the replication chain,
// so that they become siblings of their master.
// This should be called when the local master dies, and all its slaves are to be resurrected via Pseudo-GTID
func MatchUpSlaves(masterKey *InstanceKey, pattern string) ([](*Instance), *Instance, error, []error) {
	res := [](*Instance){}
	errs := []error{}

	masterInstance, found, err := ReadInstance(masterKey)
	if err != nil || !found {
		return res, nil, err, errs
	}

	return MultiMatchSlaves(masterKey, &masterInstance.MasterKey, pattern)
}

func isGenerallyValidAsBinlogSource(slave *Instance) bool {
	if !slave.IsLastCheckValid {
		// something wrong with this slave right now. We shouldn't hope to be able to promote it
		return false
	}
	if !slave.LogBinEnabled {
		return false
	}
	if !slave.LogSlaveUpdatesEnabled {
		return false
	}

	return true
}

func isGenerallyValidAsCandidateSlave(slave *Instance) bool {
	if !isGenerallyValidAsBinlogSource(slave) {
		// does not have binary logs
		return false
	}
	if slave.IsBinlogServer() {
		// Can't regroup under a binlog server because it does not support pseudo-gtid related queries such as SHOW BINLOG EVENTS
		return false
	}

	return true
}

// isValidAsCandidateMasterInBinlogServerTopology let's us know whether a given slave is generally
// valid to promote to be master.
func isValidAsCandidateMasterInBinlogServerTopology(slave *Instance) bool {
	if !slave.IsLastCheckValid {
		// something wrong with this slave right now. We shouldn't hope to be able to promote it
		return false
	}
	if !slave.LogBinEnabled {
		return false
	}
	if slave.LogSlaveUpdatesEnabled {
		// That's right: we *disallow* log-slave-updates
		return false
	}
	if slave.IsBinlogServer() {
		return false
	}

	return true
}

func isBannedFromBeingCandidateSlave(slave *Instance) bool {
	if slave.PromotionRule == MustNotPromoteRule {
		log.Debugf("instance %+v is banned because of promotion rule", slave.Key)
		return true
	}
	for _, filter := range config.Config.PromotionIgnoreHostnameFilters {
		if matched, _ := regexp.MatchString(filter, slave.Key.Hostname); matched {
			return true
		}
	}
	return false
}

// getPriorityMajorVersionForCandidate returns the primary (most common) major version found
// among given instances. This will be used for choosing best candidate for promotion.
func getPriorityMajorVersionForCandidate(slaves [](*Instance)) (priorityMajorVersion string, err error) {
	if len(slaves) == 0 {
		return "", log.Errorf("empty slaves list in getPriorityMajorVersionForCandidate")
	}
	majorVersionsCount := make(map[string]int)
	for _, slave := range slaves {
		majorVersionsCount[slave.MajorVersionString()] = majorVersionsCount[slave.MajorVersionString()] + 1
	}
	if len(majorVersionsCount) == 1 {
		// all same version, simple case
		return slaves[0].MajorVersionString(), nil
	}

	currentMaxMajorVersionCount := 0
	for majorVersion, count := range majorVersionsCount {
		if count > currentMaxMajorVersionCount {
			currentMaxMajorVersionCount = count
			priorityMajorVersion = majorVersion
		}
	}
	return priorityMajorVersion, nil
}

// getPriorityBinlogFormatForCandidate returns the primary (most common) binlog format found
// among given instances. This will be used for choosing best candidate for promotion.
func getPriorityBinlogFormatForCandidate(slaves [](*Instance)) (priorityBinlogFormat string, err error) {
	if len(slaves) == 0 {
		return "", log.Errorf("empty slaves list in getPriorityBinlogFormatForCandidate")
	}
	binlogFormatsCount := make(map[string]int)
	for _, slave := range slaves {
		binlogFormatsCount[slave.Binlog_format] = binlogFormatsCount[slave.Binlog_format] + 1
	}
	if len(binlogFormatsCount) == 1 {
		// all same binlog format, simple case
		return slaves[0].Binlog_format, nil
	}

	currentMaxBinlogFormatCount := 0
	for binlogFormat, count := range binlogFormatsCount {
		if count > currentMaxBinlogFormatCount {
			currentMaxBinlogFormatCount = count
			priorityBinlogFormat = binlogFormat
		}
	}
	return priorityBinlogFormat, nil
}

// chooseCandidateSlave
func chooseCandidateSlave(slaves [](*Instance)) (candidateSlave *Instance, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves [](*Instance), err error) {
	if len(slaves) == 0 {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, fmt.Errorf("No slaves found given in chooseCandidateSlave")
	}
	priorityMajorVersion, _ := getPriorityMajorVersionForCandidate(slaves)
	priorityBinlogFormat, _ := getPriorityBinlogFormatForCandidate(slaves)

	for _, slave := range slaves {
		slave := slave
		if isGenerallyValidAsCandidateSlave(slave) &&
			!isBannedFromBeingCandidateSlave(slave) &&
			!IsSmallerMajorVersion(priorityMajorVersion, slave.MajorVersionString()) &&
			!IsSmallerBinlogFormat(priorityBinlogFormat, slave.Binlog_format) {
			// this is the one
			candidateSlave = slave
			break
		}
	}
	if candidateSlave == nil {
		// Unable to find a candidate that will master others.
		// Instead, pick a (single) slave which is not banned.
		for _, slave := range slaves {
			slave := slave
			if !isBannedFromBeingCandidateSlave(slave) {
				// this is the one
				candidateSlave = slave
				break
			}
		}
		if candidateSlave != nil {
			slaves = RemoveInstance(slaves, &candidateSlave.Key)
		}
		return candidateSlave, slaves, equalSlaves, laterSlaves, cannotReplicateSlaves, fmt.Errorf("chooseCandidateSlave: no candidate slave found")
	}
	slaves = RemoveInstance(slaves, &candidateSlave.Key)
	for _, slave := range slaves {
		slave := slave
		if canReplicate, _ := slave.CanReplicateFrom(candidateSlave); !canReplicate {
			cannotReplicateSlaves = append(cannotReplicateSlaves, slave)
		} else if slave.ExecBinlogCoordinates.SmallerThan(&candidateSlave.ExecBinlogCoordinates) {
			laterSlaves = append(laterSlaves, slave)
		} else if slave.ExecBinlogCoordinates.Equals(&candidateSlave.ExecBinlogCoordinates) {
			equalSlaves = append(equalSlaves, slave)
		} else {
			aheadSlaves = append(aheadSlaves, slave)
		}
	}
	return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err
}

// GetCandidateSlave chooses the best slave to promote given a (possibly dead) master
func GetCandidateSlave(masterKey *InstanceKey, forRematchPurposes bool) (*Instance, [](*Instance), [](*Instance), [](*Instance), [](*Instance), error) {
	var candidateSlave *Instance
	aheadSlaves := [](*Instance){}
	equalSlaves := [](*Instance){}
	laterSlaves := [](*Instance){}
	cannotReplicateSlaves := [](*Instance){}

	slaves, err := getSlavesForSorting(masterKey, false)
	if err != nil {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err
	}
	slaves = sortedSlaves(slaves, forRematchPurposes)
	if err != nil {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err
	}
	if len(slaves) == 0 {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, fmt.Errorf("No slaves found for %+v", *masterKey)
	}
	candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err = chooseCandidateSlave(slaves)
	if err != nil {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err
	}
	log.Debugf("GetCandidateSlave: candidate: %+v, ahead: %d, equal: %d, late: %d, break: %d", candidateSlave.Key, len(aheadSlaves), len(equalSlaves), len(laterSlaves), len(cannotReplicateSlaves))
	return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, nil
}

// GetCandidateSlaveOfBinlogServerTopology chooses the best slave to promote given a (possibly dead) master
func GetCandidateSlaveOfBinlogServerTopology(masterKey *InstanceKey) (candidateSlave *Instance, err error) {
	slaves, err := getSlavesForSorting(masterKey, true)
	if err != nil {
		return candidateSlave, err
	}
	slaves = sortedSlaves(slaves, false)
	if len(slaves) == 0 {
		return candidateSlave, fmt.Errorf("No slaves found for %+v", *masterKey)
	}
	for _, slave := range slaves {
		slave := slave
		if candidateSlave != nil {
			break
		}
		if isValidAsCandidateMasterInBinlogServerTopology(slave) && !isBannedFromBeingCandidateSlave(slave) {
			// this is the one
			candidateSlave = slave
		}
	}
	if candidateSlave != nil {
		log.Debugf("GetCandidateSlaveOfBinlogServerTopology: returning %+v as candidate slave for %+v", candidateSlave.Key, *masterKey)
	} else {
		log.Debugf("GetCandidateSlaveOfBinlogServerTopology: no candidate slave found for %+v", *masterKey)
	}
	return candidateSlave, err
}

// RegroupSlavesPseudoGTID will choose a candidate slave of a given instance, and enslave its siblings using pseudo-gtid
func RegroupSlavesPseudoGTID(masterKey *InstanceKey, returnSlaveEvenOnFailureToRegroup bool, onCandidateSlaveChosen func(*Instance), postponedFunctionsContainer *PostponedFunctionsContainer) ([](*Instance), [](*Instance), [](*Instance), [](*Instance), *Instance, error) {
	candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := GetCandidateSlave(masterKey, true)
	if err != nil {
		if !returnSlaveEvenOnFailureToRegroup {
			candidateSlave = nil
		}
		return aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, candidateSlave, err
	}

	if config.Config.PseudoGTIDPattern == "" {
		return aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, candidateSlave, fmt.Errorf("PseudoGTIDPattern not configured; cannot use Pseudo-GTID")
	}

	if onCandidateSlaveChosen != nil {
		onCandidateSlaveChosen(candidateSlave)
	}

	log.Debugf("RegroupSlaves: working on %d equals slaves", len(equalSlaves))
	barrier := make(chan *InstanceKey)
	for _, slave := range equalSlaves {
		slave := slave
		// This slave has the exact same executing coordinates as the candidate slave. This slave
		// is *extremely* easy to attach below the candidate slave!
		go func() {
			defer func() { barrier <- &candidateSlave.Key }()
			ExecuteOnTopology(func() {
				ChangeMasterTo(&slave.Key, &candidateSlave.Key, &candidateSlave.SelfBinlogCoordinates, false, GTIDHintDeny)
			})
		}()
	}
	for range equalSlaves {
		<-barrier
	}

	log.Debugf("RegroupSlaves: multi matching %d later slaves", len(laterSlaves))
	// As for the laterSlaves, we'll have to apply pseudo GTID
	laterSlaves, instance, err, _ := MultiMatchBelow(laterSlaves, &candidateSlave.Key, true, postponedFunctionsContainer)

	operatedSlaves := append(equalSlaves, candidateSlave)
	operatedSlaves = append(operatedSlaves, laterSlaves...)
	log.Debugf("RegroupSlaves: starting %d slaves", len(operatedSlaves))
	barrier = make(chan *InstanceKey)
	for _, slave := range operatedSlaves {
		slave := slave
		go func() {
			defer func() { barrier <- &candidateSlave.Key }()
			ExecuteOnTopology(func() {
				StartSlave(&slave.Key)
			})
		}()
	}
	for range operatedSlaves {
		<-barrier
	}

	log.Debugf("RegroupSlaves: done")
	AuditOperation("regroup-slaves", masterKey, fmt.Sprintf("regrouped %+v slaves below %+v", len(operatedSlaves), *masterKey))
	// aheadSlaves are lost (they were ahead in replication as compared to promoted slave)
	return aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, instance, err
}

func getMostUpToDateActiveBinlogServer(masterKey *InstanceKey) (mostAdvancedBinlogServer *Instance, binlogServerSlaves [](*Instance), err error) {
	if binlogServerSlaves, err = ReadBinlogServerSlaveInstances(masterKey); err == nil && len(binlogServerSlaves) > 0 {
		// Pick the most advanced binlog sever that is good to go
		for _, binlogServer := range binlogServerSlaves {
			if binlogServer.IsLastCheckValid {
				if mostAdvancedBinlogServer == nil {
					mostAdvancedBinlogServer = binlogServer
				}
				if mostAdvancedBinlogServer.ExecBinlogCoordinates.SmallerThan(&binlogServer.ExecBinlogCoordinates) {
					mostAdvancedBinlogServer = binlogServer
				}
			}
		}
	}
	return mostAdvancedBinlogServer, binlogServerSlaves, err
}

// RegroupSlavesPseudoGTIDIncludingSubSlavesOfBinlogServers uses Pseugo-GTID to regroup slaves
// of given instance. The function also drill in to slaves of binlog servers that are replicating from given instance,
// and other recursive binlog servers, as long as they're in the same binlog-server-family.
func RegroupSlavesPseudoGTIDIncludingSubSlavesOfBinlogServers(masterKey *InstanceKey, returnSlaveEvenOnFailureToRegroup bool, onCandidateSlaveChosen func(*Instance), postponedFunctionsContainer *PostponedFunctionsContainer) ([](*Instance), [](*Instance), [](*Instance), [](*Instance), *Instance, error) {
	// First, handle binlog server issues:
	func() error {
		log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: starting on slaves of %+v", *masterKey)
		// Find the most up to date binlog server:
		mostUpToDateBinlogServer, binlogServerSlaves, err := getMostUpToDateActiveBinlogServer(masterKey)
		if err != nil {
			return log.Errore(err)
		}
		if mostUpToDateBinlogServer == nil {
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: no binlog server replicates from %+v", *masterKey)
			// No binlog server; proceed as normal
			return nil
		}
		log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: most up to date binlog server of %+v: %+v", *masterKey, mostUpToDateBinlogServer.Key)

		// Find the most up to date candidate slave:
		candidateSlave, _, _, _, _, err := GetCandidateSlave(masterKey, true)
		if err != nil {
			return log.Errore(err)
		}
		if candidateSlave == nil {
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: no candidate slave for %+v", *masterKey)
			// Let the followup code handle that
			return nil
		}
		log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: candidate slave of %+v: %+v", *masterKey, candidateSlave.Key)

		if candidateSlave.ExecBinlogCoordinates.SmallerThan(&mostUpToDateBinlogServer.ExecBinlogCoordinates) {
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: candidate slave %+v coordinates smaller than binlog server %+v", candidateSlave.Key, mostUpToDateBinlogServer.Key)
			// Need to align under binlog server...
			candidateSlave, err = Repoint(&candidateSlave.Key, &mostUpToDateBinlogServer.Key, GTIDHintDeny)
			if err != nil {
				return log.Errore(err)
			}
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: repointed candidate slave %+v under binlog server %+v", candidateSlave.Key, mostUpToDateBinlogServer.Key)
			candidateSlave, err = StartSlaveUntilMasterCoordinates(&candidateSlave.Key, &mostUpToDateBinlogServer.ExecBinlogCoordinates)
			if err != nil {
				return log.Errore(err)
			}
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: aligned candidate slave %+v under binlog server %+v", candidateSlave.Key, mostUpToDateBinlogServer.Key)
			// and move back
			candidateSlave, err = Repoint(&candidateSlave.Key, masterKey, GTIDHintDeny)
			if err != nil {
				return log.Errore(err)
			}
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: repointed candidate slave %+v under master %+v", candidateSlave.Key, *masterKey)
			return nil
		}
		// Either because it _was_ like that, or we _made_ it so,
		// candidate slave is as/more up to date than all binlog servers
		for _, binlogServer := range binlogServerSlaves {
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: matching slaves of binlog server %+v below %+v", binlogServer.Key, candidateSlave.Key)
			// Right now sequentially.
			// At this point just do what you can, don't return an error
			MultiMatchSlaves(&binlogServer.Key, &candidateSlave.Key, "")
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: done matching slaves of binlog server %+v below %+v", binlogServer.Key, candidateSlave.Key)
		}
		log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: done handling binlog regrouping for %+v; will proceed with normal RegroupSlaves", *masterKey)
		AuditOperation("regroup-slaves-including-bls", masterKey, fmt.Sprintf("matched slaves of binlog server slaves of %+v under %+v", *masterKey, candidateSlave.Key))
		return nil
	}()
	// Proceed to normal regroup:
	return RegroupSlavesPseudoGTID(masterKey, returnSlaveEvenOnFailureToRegroup, onCandidateSlaveChosen, postponedFunctionsContainer)
}

// RegroupSlavesGTID will choose a candidate slave of a given instance, and enslave its siblings using GTID
func RegroupSlavesGTID(masterKey *InstanceKey, returnSlaveEvenOnFailureToRegroup bool, onCandidateSlaveChosen func(*Instance)) ([](*Instance), [](*Instance), [](*Instance), *Instance, error) {
	var emptySlaves [](*Instance)
	candidateSlave, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := GetCandidateSlave(masterKey, true)
	if err != nil {
		if !returnSlaveEvenOnFailureToRegroup {
			candidateSlave = nil
		}
		return emptySlaves, emptySlaves, emptySlaves, candidateSlave, err
	}

	if onCandidateSlaveChosen != nil {
		onCandidateSlaveChosen(candidateSlave)
	}

	slavesToMove := append(equalSlaves, laterSlaves...)
	log.Debugf("RegroupSlavesGTID: working on %d slaves", len(slavesToMove))

	movedSlaves, unmovedSlaves, err, _ := moveSlavesViaGTID(slavesToMove, candidateSlave)
	if err != nil {
		log.Errore(err)
	}
	unmovedSlaves = append(unmovedSlaves, aheadSlaves...)
	StartSlave(&candidateSlave.Key)

	log.Debugf("RegroupSlavesGTID: done")
	AuditOperation("regroup-slaves-gtid", masterKey, fmt.Sprintf("regrouped slaves of %+v via GTID; promoted %+v", *masterKey, candidateSlave.Key))
	return unmovedSlaves, movedSlaves, cannotReplicateSlaves, candidateSlave, err
}

// RegroupSlavesBinlogServers works on a binlog-servers topology. It picks the most up-to-date BLS and repoints all other
// BLS below it
func RegroupSlavesBinlogServers(masterKey *InstanceKey, returnSlaveEvenOnFailureToRegroup bool) (repointedBinlogServers [](*Instance), promotedBinlogServer *Instance, err error) {
	var binlogServerSlaves [](*Instance)
	promotedBinlogServer, binlogServerSlaves, err = getMostUpToDateActiveBinlogServer(masterKey)

	resultOnError := func(err error) ([](*Instance), *Instance, error) {
		if !returnSlaveEvenOnFailureToRegroup {
			promotedBinlogServer = nil
		}
		return repointedBinlogServers, promotedBinlogServer, err
	}

	if err != nil {
		return resultOnError(err)
	}

	repointedBinlogServers, err, _ = RepointTo(binlogServerSlaves, &promotedBinlogServer.Key)

	if err != nil {
		return resultOnError(err)
	}
	AuditOperation("regroup-slaves-bls", masterKey, fmt.Sprintf("regrouped binlog server slaves of %+v; promoted %+v", *masterKey, promotedBinlogServer.Key))
	return repointedBinlogServers, promotedBinlogServer, nil
}

// RegroupSlaves is a "smart" method of promoting one slave over the others ("promoting" it on top of its siblings)
// This method decides which strategy to use: GTID, Pseudo-GTID, Binlog Servers.
func RegroupSlaves(masterKey *InstanceKey, returnSlaveEvenOnFailureToRegroup bool,
	onCandidateSlaveChosen func(*Instance),
	postponedFunctionsContainer *PostponedFunctionsContainer) (
	aheadSlaves [](*Instance), equalSlaves [](*Instance), laterSlaves [](*Instance), cannotReplicateSlaves [](*Instance), instance *Instance, err error) {
	//
	var emptySlaves [](*Instance)

	slaves, err := ReadSlaveInstances(masterKey)
	if err != nil {
		return emptySlaves, emptySlaves, emptySlaves, emptySlaves, instance, err
	}
	if len(slaves) == 0 {
		return emptySlaves, emptySlaves, emptySlaves, emptySlaves, instance, err
	}
	if len(slaves) == 1 {
		return emptySlaves, emptySlaves, emptySlaves, emptySlaves, slaves[0], err
	}
	allGTID := true
	allBinlogServers := true
	allPseudoGTID := true
	for _, slave := range slaves {
		if !slave.UsingGTID() {
			allGTID = false
		}
		if !slave.IsBinlogServer() {
			allBinlogServers = false
		}
		if !slave.UsingPseudoGTID {
			allPseudoGTID = false
		}
	}
	if allGTID {
		log.Debugf("RegroupSlaves: using GTID to regroup slaves of %+v", *masterKey)
		unmovedSlaves, movedSlaves, cannotReplicateSlaves, candidateSlave, err := RegroupSlavesGTID(masterKey, returnSlaveEvenOnFailureToRegroup, onCandidateSlaveChosen)
		return unmovedSlaves, emptySlaves, movedSlaves, cannotReplicateSlaves, candidateSlave, err
	}
	if allBinlogServers {
		log.Debugf("RegroupSlaves: using binlog servers to regroup slaves of %+v", *masterKey)
		movedSlaves, candidateSlave, err := RegroupSlavesBinlogServers(masterKey, returnSlaveEvenOnFailureToRegroup)
		return emptySlaves, emptySlaves, movedSlaves, cannotReplicateSlaves, candidateSlave, err
	}
	if allPseudoGTID {
		log.Debugf("RegroupSlaves: using Pseudo-GTID to regroup slaves of %+v", *masterKey)
		return RegroupSlavesPseudoGTID(masterKey, returnSlaveEvenOnFailureToRegroup, onCandidateSlaveChosen, postponedFunctionsContainer)
	}
	// And, as last resort, we do PseudoGTID & binlog servers
	log.Warningf("RegroupSlaves: unsure what method to invoke for %+v; trying Pseudo-GTID+Binlog Servers", *masterKey)
	return RegroupSlavesPseudoGTIDIncludingSubSlavesOfBinlogServers(masterKey, returnSlaveEvenOnFailureToRegroup, onCandidateSlaveChosen, postponedFunctionsContainer)
}

// relocateBelowInternal is a protentially recursive function which chooses how to relocate an instance below another.
// It may choose to use Pseudo-GTID, or normal binlog positions, or take advantage of binlog servers,
// or it may combine any of the above in a multi-step operation.
func relocateBelowInternal(instance, other *Instance) (*Instance, error) {
	if canReplicate, err := instance.CanReplicateFrom(other); !canReplicate {
		return instance, log.Errorf("%+v cannot replicate from %+v. Reason: %+v", instance.Key, other.Key, err)
	}
	// simplest:
	if InstanceIsMasterOf(other, instance) {
		// already the desired setup.
		return Repoint(&instance.Key, &other.Key, GTIDHintNeutral)
	}
	// Do we have record of equivalent coordinates?
	if !instance.IsBinlogServer() {
		if movedInstance, err := MoveEquivalent(&instance.Key, &other.Key); err == nil {
			return movedInstance, nil
		}
	}
	// Try and take advantage of binlog servers:
	if InstancesAreSiblings(instance, other) && other.IsBinlogServer() {
		return MoveBelow(&instance.Key, &other.Key)
	}
	instanceMaster, found, err := ReadInstance(&instance.MasterKey)
	if err != nil || !found {
		return instance, err
	}
	if instanceMaster.MasterKey.Equals(&other.Key) && instanceMaster.IsBinlogServer() {
		// Moving to grandparent via binlog server
		return Repoint(&instance.Key, &instanceMaster.MasterKey, GTIDHintDeny)
	}
	if other.IsBinlogServer() {
		if instanceMaster.IsBinlogServer() && InstancesAreSiblings(instanceMaster, other) {
			// Special case: this is a binlog server family; we move under the uncle, in one single step
			return Repoint(&instance.Key, &other.Key, GTIDHintDeny)
		}

		// Relocate to its master, then repoint to the binlog server
		otherMaster, found, err := ReadInstance(&other.MasterKey)
		if err != nil || !found {
			return instance, err
		}
		if !other.IsLastCheckValid {
			return instance, log.Errorf("Binlog server %+v is not reachable. It would take two steps to relocate %+v below it, and I won't even do the first step.", other.Key, instance.Key)
		}

		log.Debugf("Relocating to a binlog server; will first attempt to relocate to the binlog server's master: %+v, and then repoint down", otherMaster.Key)
		if _, err := relocateBelowInternal(instance, otherMaster); err != nil {
			return instance, err
		}
		return Repoint(&instance.Key, &other.Key, GTIDHintDeny)
	}
	if instance.IsBinlogServer() {
		// Can only move within the binlog-server family tree
		// And these have been covered just now: move up from a master binlog server, move below a binling binlog server.
		// sure, the family can be more complex, but we keep these operations atomic
		return nil, log.Errorf("Relocating binlog server %+v below %+v turns to be too complex; please do it manually", instance.Key, other.Key)
	}
	// Next, try GTID
	if _, _, canMove := canMoveViaGTID(instance, other); canMove {
		return moveInstanceBelowViaGTID(instance, other)
	}

	// Next, try Pseudo-GTID
	if instance.UsingPseudoGTID && other.UsingPseudoGTID {
		// We prefer PseudoGTID to anything else because, while it takes longer to run, it does not issue
		// a STOP SLAVE on any server other than "instance" itself.
		instance, _, err := MatchBelow(&instance.Key, &other.Key, true)
		return instance, err
	}
	// No Pseudo-GTID; cehck simple binlog file/pos operations:
	if InstancesAreSiblings(instance, other) {
		// If comastering, only move below if it's read-only
		if !other.IsCoMaster || other.ReadOnly {
			return MoveBelow(&instance.Key, &other.Key)
		}
	}
	// See if we need to MoveUp
	if instanceMaster.MasterKey.Equals(&other.Key) {
		// Moving to grandparent--handles co-mastering writable case
		return MoveUp(&instance.Key)
	}
	if instanceMaster.IsBinlogServer() {
		// Break operation into two: move (repoint) up, then continue
		if _, err := MoveUp(&instance.Key); err != nil {
			return instance, err
		}
		return relocateBelowInternal(instance, other)
	}
	// Too complex
	return nil, log.Errorf("Relocating %+v below %+v turns to be too complex; please do it manually", instance.Key, other.Key)
}

// RelocateBelow will attempt moving instance indicated by instanceKey below another instance.
// Orchestrator will try and figure out the best way to relocate the server. This could span normal
// binlog-position, pseudo-gtid, repointing, binlog servers...
func RelocateBelow(instanceKey, otherKey *InstanceKey) (*Instance, error) {
	instance, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return instance, log.Errorf("Error reading %+v", *instanceKey)
	}
	other, found, err := ReadInstance(otherKey)
	if err != nil || !found {
		return instance, log.Errorf("Error reading %+v", *otherKey)
	}
	instance, err = relocateBelowInternal(instance, other)
	if err == nil {
		AuditOperation("relocate-below", instanceKey, fmt.Sprintf("relocated %+v below %+v", *instanceKey, *otherKey))
	}
	return instance, err
}

// relocateSlavesInternal is a protentially recursive function which chooses how to relocate
// slaves of an instance below another.
// It may choose to use Pseudo-GTID, or normal binlog positions, or take advantage of binlog servers,
// or it may combine any of the above in a multi-step operation.
func relocateSlavesInternal(slaves [](*Instance), instance, other *Instance) ([](*Instance), error, []error) {
	errs := []error{}
	var err error
	// simplest:
	if instance.Key.Equals(&other.Key) {
		// already the desired setup.
		return RepointTo(slaves, &other.Key)
	}
	// Try and take advantage of binlog servers:
	if InstanceIsMasterOf(other, instance) && instance.IsBinlogServer() {
		// Up from a binlog server
		return RepointTo(slaves, &other.Key)
	}
	if InstanceIsMasterOf(instance, other) && other.IsBinlogServer() {
		// Down under a binlog server
		return RepointTo(slaves, &other.Key)
	}
	if InstancesAreSiblings(instance, other) && instance.IsBinlogServer() && other.IsBinlogServer() {
		// Between siblings
		return RepointTo(slaves, &other.Key)
	}
	if other.IsBinlogServer() {
		// Relocate to binlog server's parent (recursive call), then repoint down
		otherMaster, found, err := ReadInstance(&other.MasterKey)
		if err != nil || !found {
			return nil, err, errs
		}
		slaves, err, errs = relocateSlavesInternal(slaves, instance, otherMaster)
		if err != nil {
			return slaves, err, errs
		}

		return RepointTo(slaves, &other.Key)
	}
	// GTID
	{
		movedSlaves, unmovedSlaves, err, errs := moveSlavesViaGTID(slaves, other)

		if len(movedSlaves) == len(slaves) {
			// Moved (or tried moving) everything via GTID
			return movedSlaves, err, errs
		} else if len(movedSlaves) > 0 {
			// something was moved via GTID; let's try further on
			return relocateSlavesInternal(unmovedSlaves, instance, other)
		}
		// Otherwise nothing was moved via GTID. Maybe we don't have any GTIDs, we continue.
	}

	// Pseudo GTID
	if other.UsingPseudoGTID {
		// Which slaves are using Pseudo GTID?
		var pseudoGTIDSlaves [](*Instance)
		for _, slave := range slaves {
			if slave.UsingPseudoGTID {
				pseudoGTIDSlaves = append(pseudoGTIDSlaves, slave)
			}
		}
		pseudoGTIDSlaves, _, err, errs = MultiMatchBelow(pseudoGTIDSlaves, &other.Key, false, nil)
		return pseudoGTIDSlaves, err, errs
	}

	// Normal binlog file:pos
	if InstanceIsMasterOf(other, instance) {
		// moveUpSlaves -- but not supporting "slaves" argument at this time.
	}

	// Too complex
	return nil, log.Errorf("Relocating %+v slaves of %+v below %+v turns to be too complex; please do it manually", len(slaves), instance.Key, other.Key), errs
}

// RelocateSlaves will attempt moving slaves of an instance indicated by instanceKey below another instance.
// Orchestrator will try and figure out the best way to relocate the servers. This could span normal
// binlog-position, pseudo-gtid, repointing, binlog servers...
func RelocateSlaves(instanceKey, otherKey *InstanceKey, pattern string) (slaves [](*Instance), other *Instance, err error, errs []error) {

	instance, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return slaves, other, log.Errorf("Error reading %+v", *instanceKey), errs
	}
	other, found, err = ReadInstance(otherKey)
	if err != nil || !found {
		return slaves, other, log.Errorf("Error reading %+v", *otherKey), errs
	}

	slaves, err = ReadSlaveInstances(instanceKey)
	if err != nil {
		return slaves, other, err, errs
	}
	slaves = RemoveInstance(slaves, otherKey)
	slaves = filterInstancesByPattern(slaves, pattern)
	if len(slaves) == 0 {
		// Nothing to do
		return slaves, other, nil, errs
	}
	slaves, err, errs = relocateSlavesInternal(slaves, instance, other)

	if err == nil {
		AuditOperation("relocate-slaves", instanceKey, fmt.Sprintf("relocated %+v slaves of %+v below %+v", len(slaves), *instanceKey, *otherKey))
	}
	return slaves, other, err, errs
}
