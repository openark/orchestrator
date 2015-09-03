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
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/config"
	"regexp"
	"sort"
	"strings"
	"time"
)

// InstancesByExecBinlogCoordinates is a sortabel type for BinlogCoordinates
type InstancesByExecBinlogCoordinates [](*Instance)

func (this InstancesByExecBinlogCoordinates) Len() int      { return len(this) }
func (this InstancesByExecBinlogCoordinates) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func (this InstancesByExecBinlogCoordinates) Less(i, j int) bool {
	if this[i].ExecBinlogCoordinates.Equals(&this[j].ExecBinlogCoordinates) {
		// Secondary sorting: "smaller" if not logging slave updates
		if this[j].LogSlaveUpdatesEnabled && !this[i].LogSlaveUpdatesEnabled {
			return true
		}
	}
	return this[i].ExecBinlogCoordinates.SmallerThan(&this[j].ExecBinlogCoordinates)
}

// getASCIITopologyEntry will get an ascii topology tree rooted at given instance. Ir recursively
// draws the tree
func getASCIITopologyEntry(depth int, instance *Instance, replicationMap map[*Instance]([]*Instance), extendedOutput bool) []string {
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

// ASCIITopology returns a string representation of the topology of given instance.
func ASCIITopology(instanceKey *InstanceKey, historyTimestampPattern string) (string, error) {
	instance, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return "", err
	}
	var instances [](*Instance)
	if historyTimestampPattern == "" {
		instances, err = ReadClusterInstances(instance.ClusterName)
	} else {
		instances, err = ReadHistoryClusterInstances(instance.ClusterName, historyTimestampPattern)
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
	if masterInstance == nil {
		return "", nil
	}
	resultArray := getASCIITopologyEntry(0, masterInstance, replicationMap, historyTimestampPattern == "")
	result := strings.Join(resultArray, "\n")
	return result, nil
}

// filterInstancesByPattern will filter given array of instances according to regular expression pattern
func filterInstancesByPattern(instances [](*Instance), pattern string) [](*Instance) {
	if pattern == "" {
		return instances
	}
	filtered := [](*Instance){}
	for _, instance := range instances {
		if matched, _ := regexp.MatchString(pattern, instance.Key.DisplayString()); matched {
			filtered = append(filtered, instance)
		}
	}
	return filtered
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
	var barrier chan *Instance

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

	barrier = make(chan *Instance)
	for _, slave := range slaves {
		slave := slave
		go func() {
			defer func() {
				slave, _ := StartSlave(&slave.Key)
				barrier <- slave
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
	slaves = removeInstance(slaves, &other.Key)
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
		return movedSlaves, unmovedSlaves, log.Error("Error on all operations"), errs
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
func Repoint(instanceKey *InstanceKey, masterKey *InstanceKey, gitdHint OperationGTIDHint) (*Instance, error) {
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
	// With repoint we *prefer* the master to be alive, but we don't structly require it.
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
	instance, err = ChangeMasterTo(instanceKey, masterKey, &instance.ExecBinlogCoordinates, !masterIsAccessible, gitdHint)
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

// RepointSlaves repoints slaves of a given instance (possibly filtered) onto another master.
// Binlog Server is the major use case
func RepointSlavesTo(instanceKey *InstanceKey, pattern string, belowKey *InstanceKey) ([](*Instance), error, []error) {
	res := [](*Instance){}
	errs := []error{}

	slaves, err := ReadSlaveInstances(instanceKey)
	if err != nil {
		return res, err, errs
	}
	slaves = removeInstance(slaves, belowKey)
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
		return instance, fmt.Errorf("instance  %+v is already co master of %+v", instanceKey, master.Key)
	}
	if _, found, _ := ReadInstance(&master.MasterKey); found {
		return instance, fmt.Errorf("master %+v already has known master: %+v", master.Key, master.MasterKey)
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

	log.Infof("Will reset %+v", instanceKey)

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
	AuditOperation("reset slave", instanceKey, fmt.Sprintf("%+v replication reset", *instanceKey))

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
	AuditOperation("detach slave", instanceKey, fmt.Sprintf("%+v replication detached", *instanceKey))

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
	AuditOperation("reattach slave", instanceKey, fmt.Sprintf("%+v replication reattached", *instanceKey))

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

// FindLastPseudoGTIDEntry will search an instance's binary logs or relay logs for the last pseudo-GTID entry,
// and return found coordinates as well as entry text
func FindLastPseudoGTIDEntry(instance *Instance, recordedInstanceRelayLogCoordinates BinlogCoordinates, exhaustiveSearch bool, expectedBinlogFormat *string) (*BinlogCoordinates, string, error) {
	var instancePseudoGtidText string
	var instancePseudoGtidCoordinates *BinlogCoordinates
	var err error

	if instance.LogBinEnabled && instance.LogSlaveUpdatesEnabled && (expectedBinlogFormat == nil || instance.Binlog_format == *expectedBinlogFormat) {
		// Well no need to search this instance's binary logs if it doesn't have any...
		// With regard log-slave-updates, some edge cases are possible, like having this instance's log-slave-updates
		// enabled/disabled (of course having restarted it)
		// The approach is not to take chances. If log-slave-updates is disabled, fail and go for relay-logs.
		// If log-slave-updates was just enabled then possibly no pseudo-gtid is found, and so again we will go
		// for relay logs.
		// Also, if master has STATEMENT binlog format, and the slave has ROW binlog format, then comparing binlog entries would urely fail if based on the slave's binary logs.
		// Instead, we revert to the relay logs.
		instancePseudoGtidCoordinates, instancePseudoGtidText, err = getLastPseudoGTIDEntryInInstance(instance, exhaustiveSearch)
	}
	if err != nil || instancePseudoGtidCoordinates == nil {
		// Unable to find pseudo GTID in binary logs.
		// Then MAYBE we are lucky enough (chances are we are, if this slave did not crash) that we can
		// extract the Pseudo GTID entry from the last (current) relay log file.
		instancePseudoGtidCoordinates, instancePseudoGtidText, err = getLastPseudoGTIDEntryInRelayLogs(instance, recordedInstanceRelayLogCoordinates, exhaustiveSearch)
	}
	return instancePseudoGtidCoordinates, instancePseudoGtidText, err
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
	var instancePseudoGtidText string
	var instancePseudoGtidCoordinates *BinlogCoordinates
	var otherInstancePseudoGtidCoordinates *BinlogCoordinates
	var nextBinlogCoordinatesToMatch *BinlogCoordinates
	var recordedInstanceRelayLogCoordinates BinlogCoordinates
	var countMatchedEvents int
	var entriesMonotonic bool

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
	// We record the relay log coordinates just after the instance stopped since the coordinates can change upon
	// a FLUSH LOGS/FLUSH RELAY LOGS (or a START SLAVE, though that's an altogether different problem) etc.
	// We want to be on the safe side; we don't utterly trust that we are the only ones playing with the instance.
	recordedInstanceRelayLogCoordinates = instance.RelaylogCoordinates
	instancePseudoGtidCoordinates, instancePseudoGtidText, err = FindLastPseudoGTIDEntry(instance, recordedInstanceRelayLogCoordinates, true, &otherInstance.Binlog_format)

	if err != nil {
		goto Cleanup
	}
	entriesMonotonic = (config.Config.PseudoGTIDMonotonicHint != "") && strings.Contains(instancePseudoGtidText, config.Config.PseudoGTIDMonotonicHint)
	otherInstancePseudoGtidCoordinates, err = SearchPseudoGTIDEntryInInstance(otherInstance, instancePseudoGtidText, entriesMonotonic)
	if err != nil {
		goto Cleanup
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
	nextBinlogCoordinatesToMatch, countMatchedEvents, err = GetNextBinlogCoordinatesToMatch(instance, *instancePseudoGtidCoordinates,
		recordedInstanceRelayLogCoordinates, otherInstance, *otherInstancePseudoGtidCoordinates)
	if err != nil {
		goto Cleanup
	}
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
	if err != nil {
		if masterInstance.IsSlave() {
			return instance, fmt.Errorf("MakeMaster: instance's master %+v seems to be replicating", masterInstance.Key)
		}
		if masterInstance.IsLastCheckValid {
			return instance, fmt.Errorf("MakeMaster: instance's master %+v seems to be accessible", masterInstance.Key)
		}
	}
	// If err == nil this is "good": that means the master is inaccessible... So it's OK to do the promotion
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

	_, _, err, _ = MultiMatchBelow(siblings, instanceKey, false)
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

	_, _, err, _ = MultiMatchBelow(siblings, instanceKey, false)
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

// sortedSlaves returns the list of slaves of a given master, sorted by exec coordinates
// (most up-to-date slave first)
func sortedSlaves(masterKey *InstanceKey, shouldStopSlaves bool) ([](*Instance), error) {
	slaves, err := ReadSlaveInstances(masterKey)
	if err != nil {
		return slaves, err
	}
	if len(slaves) == 0 {
		return slaves, nil
	}
	if shouldStopSlaves {
		log.Debugf("sortedSlaves: stopping %d slaves nicely", len(slaves))
		slaves = StopSlavesNicely(slaves, time.Duration(config.Config.InstanceBulkOperationsWaitTimeoutSeconds)*time.Second)
	}

	sort.Sort(sort.Reverse(InstancesByExecBinlogCoordinates(slaves)))
	for _, slave := range slaves {
		log.Debugf("- sorted slave: %+v %+v", slave.Key, slave.ExecBinlogCoordinates)
	}

	return slaves, err
}

// removeInstance will remove an instance from a list of instances
func removeInstance(instances [](*Instance), instanceKey *InstanceKey) [](*Instance) {
	if instanceKey == nil {
		return instances
	}
	for i := len(instances) - 1; i >= 0; i-- {
		if instances[i].Key.Equals(instanceKey) {
			instances = append(instances[:i], instances[i+1:]...)
		}
	}
	return instances
}

// removeBinlogServerInstances will remove all binlog servers from given lsit
func removeBinlogServerInstances(instances [](*Instance)) [](*Instance) {
	for i := len(instances) - 1; i >= 0; i-- {
		if instances[i].IsBinlogServer() {
			instances = append(instances[:i], instances[i+1:]...)
		}
	}
	return instances
}

// MultiMatchBelow will efficiently match multiple slaves below a given instance.
// It is assumed that all given slaves are siblings
func MultiMatchBelow(slaves [](*Instance), belowKey *InstanceKey, slavesAlreadyStopped bool) ([](*Instance), *Instance, error, []error) {
	res := [](*Instance){}
	errs := []error{}
	slaveMutex := make(chan bool, 1)

	if config.Config.PseudoGTIDPattern == "" {
		return res, nil, fmt.Errorf("PseudoGTIDPattern not configured; cannot use Pseudo-GTID"), errs
	}

	slaves = removeInstance(slaves, belowKey)
	slaves = removeBinlogServerInstances(slaves)

	for _, slave := range slaves {
		if maintenanceToken, merr := BeginMaintenance(&slave.Key, GetMaintenanceOwner(), fmt.Sprintf("%+v match below %+v as part of MultiMatchBelow", slave.Key, *belowKey)); merr != nil {
			errs = append(errs, fmt.Errorf("Cannot begin maintenance on %+v", slave.Key))
			slaves = removeInstance(slaves, &slave.Key)
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
					ExecuteOnTopology(func() {
						_, matchedCoordinates, slaveErr = MatchBelow(&slave.Key, &belowInstance.Key, false)
					})
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
			// At this point our bucket is complete.
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
	matchedSlaves, belowInstance, err, errs := MultiMatchBelow(slaves, &belowInstance.Key, false)

	if len(matchedSlaves) != len(slaves) {
		err = fmt.Errorf("MultiMatchSlaves: only matched %d out of %d slaves of %+v; error is: %+v", len(matchedSlaves), len(slaves), *masterKey, err)
	}

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

func isGenerallyValidAsCandidateSlave(slave *Instance) bool {
	if !slave.IsLastCheckValid {
		// something wrong with this slave irhgt now. We shouldn't hope to be able to promote it
		return false
	}
	if !slave.LogBinEnabled {
		return false
	}
	if !slave.LogSlaveUpdatesEnabled {
		return false
	}
	if slave.IsBinlogServer() {
		// Can't regroup under a binlog server because it does not support pseudo-gtid related queries such as SHOW BINLOG EVENTS
		return false
	}
	for _, filter := range config.Config.PromotionIgnoreHostnameFilters {
		if matched, _ := regexp.MatchString(filter, slave.Key.Hostname); matched {
			return false
		}
	}

	return true
}

// GetCandidateSlave chooses the best slave to promote given a (possibly dead) master
func GetCandidateSlave(masterKey *InstanceKey, forRematchPurposes bool) (*Instance, [](*Instance), [](*Instance), [](*Instance), error) {
	var candidateSlave *Instance
	aheadSlaves := [](*Instance){}
	equalSlaves := [](*Instance){}
	laterSlaves := [](*Instance){}

	slaves, err := sortedSlaves(masterKey, forRematchPurposes)
	if err != nil {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, err
	}
	if len(slaves) == 0 {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, fmt.Errorf("No slaves found for %+v", *masterKey)
	}
	for _, slave := range slaves {
		slave := slave
		if isGenerallyValidAsCandidateSlave(slave) {
			// this is the one
			candidateSlave = slave
			break
		}
	}
	if candidateSlave == nil {
		return slaves[0], slaves[1:], equalSlaves, laterSlaves, fmt.Errorf("No slaves found with log_slave_updates for %+v", *masterKey)
	}
	slaves = removeInstance(slaves, &candidateSlave.Key)
	for _, slave := range slaves {
		slave := slave
		if slave.ExecBinlogCoordinates.SmallerThan(&candidateSlave.ExecBinlogCoordinates) {
			laterSlaves = append(laterSlaves, slave)
		} else if slave.ExecBinlogCoordinates.Equals(&candidateSlave.ExecBinlogCoordinates) {
			equalSlaves = append(equalSlaves, slave)
		} else {
			aheadSlaves = append(aheadSlaves, slave)
		}
	}
	log.Debugf("sortedSlaves: candidate: %+v, ahead: %d, equal: %d, late: %d", candidateSlave.Key, len(aheadSlaves), len(equalSlaves), len(laterSlaves))
	return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, nil
}

// RegroupSlaves will choose a candidate slave of a given instance, and enslave its siblings using
// either simple CHANGE MASTER TO, where possible, or pseudo-gtid
func RegroupSlaves(masterKey *InstanceKey, returnSlaveEvenOnFailureToRegroup bool, onCandidateSlaveChosen func(*Instance)) ([](*Instance), [](*Instance), [](*Instance), *Instance, error) {
	candidateSlave, aheadSlaves, equalSlaves, laterSlaves, err := GetCandidateSlave(masterKey, true)
	if err != nil {
		if !returnSlaveEvenOnFailureToRegroup {
			candidateSlave = nil
		}
		return aheadSlaves, equalSlaves, laterSlaves, candidateSlave, err
	}

	if config.Config.PseudoGTIDPattern == "" {
		return aheadSlaves, equalSlaves, laterSlaves, candidateSlave, fmt.Errorf("PseudoGTIDPattern not configured; cannot use Pseudo-GTID")
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
	laterSlaves, instance, err, _ := MultiMatchBelow(laterSlaves, &candidateSlave.Key, true)

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
	return aheadSlaves, equalSlaves, laterSlaves, instance, err
}

func getMostUpToDateActiveBinlogServer(masterKey *InstanceKey) (mostAdvancedBinlogServer *Instance, err error) {
	var binlogServerSlaves [](*Instance)
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
	return mostAdvancedBinlogServer, err
}

// RegroupSlaves will choose a candidate slave of a given instance, and enslave its siblings using
// either simple CHANGE MASTER TO, where possible, or pseudo-gtid
func RegroupSlavesIncludingSubSlavesOfBinlogServers(masterKey *InstanceKey, returnSlaveEvenOnFailureToRegroup bool, onCandidateSlaveChosen func(*Instance)) ([](*Instance), [](*Instance), [](*Instance), *Instance, error) {

	if mostAdvancedBinlogServer, _ := getMostUpToDateActiveBinlogServer(masterKey); mostAdvancedBinlogServer != nil {
		log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: mostAdvancedBinlogServer is %+v", mostAdvancedBinlogServer.Key)
		if allSlaves, err := ReadSlaveInstancesIncludingBinlogServerSubSlaves(masterKey); err == nil {
			slavesBehindMostAdvancedBinlogServer := [](*Instance){}
			foundPotentialPromotedSlaveAheadOfBinlogServer := false
			for _, slave := range allSlaves {
				slave := slave
				if slave.ExecBinlogCoordinates.SmallerThan(&mostAdvancedBinlogServer.ExecBinlogCoordinates) {
					slavesBehindMostAdvancedBinlogServer = append(slavesBehindMostAdvancedBinlogServer, slave)
				} else if isGenerallyValidAsCandidateSlave(slave) && !foundPotentialPromotedSlaveAheadOfBinlogServer {
					// We have a slave with log-slave-updates that is *ahead* of most-up-to-date binlog server.
					// This means all slaves behind binlog servers are able to match below said slave via pseudo-gtid.
					// This keeps us on safe grounds. We don't need to make further precautionary checks or waits
					// on binlog server slaves.
					foundPotentialPromotedSlaveAheadOfBinlogServer = true
					log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: Found %+v to be ahead of most up to date binlog server %+v", slave.Key, mostAdvancedBinlogServer.Key)
				}
			}
			// Everything that stopped replicating on an earlier point than the mostAdvancedBinlogServer will be
			// directed at mostAdvancedBinlogServer to continue replicating as much as possible
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: will repoint %+v slaves that are behind replication as compared to %+v below it", len(slavesBehindMostAdvancedBinlogServer), mostAdvancedBinlogServer.Key)
			RepointTo(slavesBehindMostAdvancedBinlogServer, &mostAdvancedBinlogServer.Key)

			if foundPotentialPromotedSlaveAheadOfBinlogServer {
				// We on safe grounds. But let's spend a very short time to allow slaves to align up to the binlog server
				if config.Config.SlaveStartPostWaitMilliseconds > 0 {
					time.Sleep(time.Duration(config.Config.SlaveStartPostWaitMilliseconds) * time.Millisecond)
					// on one hand, we just wasted time. On the other hand, we hope all slaves are now aligned, so are in same exec binlog positions, which means
					// less computation for pseudo-gtid
				}
			} else {
				// Make sure a good candidate slave is aligned with the binlog server. We spend time on this now.
				// This slave would be able to enslave all its siblings (ie all slaves that were behind binlog server)
				log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: Haven't found a slave to promote that is ahead of most up to date binlog server %+v. Will find a suitable slave behind the binlog server.", mostAdvancedBinlogServer.Key)
				if candidateSlaveOfBinlogServer, _, _, _, err := GetCandidateSlave(&mostAdvancedBinlogServer.Key, false); err == nil && candidateSlaveOfBinlogServer != nil {
					log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: will start %+v until most-up-to-date binlog server %+v coordinates", candidateSlaveOfBinlogServer.Key, mostAdvancedBinlogServer.Key)
					StartSlaveUntilMasterCoordinates(&candidateSlaveOfBinlogServer.Key, &mostAdvancedBinlogServer.ExecBinlogCoordinates)
				}
			}

			// Remember all slaves involved were either direct slaves of masterKey or slaves of binlog server slaves
			// of masterKey. It is thus safe to move everything (ie normal slaves) around. We will now repoint everything
			// to masterKey, thus all the slaves are now aligned as siblings. We can now continue with a normal regroup.
			log.Debugf("RegroupSlavesIncludingSubSlavesOfBinlogServers: will align up all slaves and subslaves of binlog servers of %+v ", *masterKey)
			RepointTo(allSlaves, masterKey)
		}
	}
	return RegroupSlaves(masterKey, returnSlaveEvenOnFailureToRegroup, onCandidateSlaveChosen)
}

// RegroupSlavesGTID will choose a candidate slave of a given instance, and enslave its siblings using GTID
func RegroupSlavesGTID(masterKey *InstanceKey, returnSlaveEvenOnFailureToRegroup bool, onCandidateSlaveChosen func(*Instance)) ([](*Instance), [](*Instance), *Instance, error) {
	var emptySlaves [](*Instance)
	candidateSlave, aheadSlaves, equalSlaves, laterSlaves, err := GetCandidateSlave(masterKey, true)
	if err != nil {
		if !returnSlaveEvenOnFailureToRegroup {
			candidateSlave = nil
		}
		return emptySlaves, emptySlaves, candidateSlave, err
	}

	if onCandidateSlaveChosen != nil {
		onCandidateSlaveChosen(candidateSlave)
	}

	slavesToMove := append(equalSlaves, laterSlaves...)
	log.Debugf("RegroupSlavesGTID: working on %d slaves", len(slavesToMove))

	movedSlaves, unmovedSlaves, err, _ := moveSlavesViaGTID(slavesToMove, candidateSlave)
	unmovedSlaves = append(unmovedSlaves, aheadSlaves...)
	StartSlave(&candidateSlave.Key)

	log.Debugf("RegroupSlavesGTID: done")
	return unmovedSlaves, movedSlaves, candidateSlave, err
}

// relocateBelowInternal is a protentially recursive function which chooses how to relocate an instance below another.
// It may choose to use Pseudo-GTID, or normal binlog positions, or take advantage of binlog servers,
// or it may combine any of the above in a multi-step operation.
func relocateBelowInternal(instance, other *Instance) (*Instance, error) {
	if canReplicate, _ := instance.CanReplicateFrom(other); !canReplicate {
		return instance, log.Errorf("%+v cannot replicate from %+v", instance.Key, other.Key)
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
		return MoveUp(&instance.Key)
	}
	if other.IsBinlogServer() {
		// Relocate to its master, then repoint to the binlog server
		otherMaster, found, err := ReadInstance(&other.MasterKey)
		if err != nil || !found {
			return instance, err
		}
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
		return MoveBelow(&instance.Key, &other.Key)
	}
	// See if we need to MoveUp
	if instanceMaster.MasterKey.Equals(&other.Key) {
		// Moving to grandparent
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
		pseudoGTIDSlaves, _, err, errs = MultiMatchBelow(pseudoGTIDSlaves, &other.Key, false)
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
	slaves = removeInstance(slaves, otherKey)
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
