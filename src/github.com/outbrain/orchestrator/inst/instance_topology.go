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
	"github.com/outbrain/orchestrator/config"
	"regexp"
	"sort"
	"strings"
	"time"
)

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
	entry := fmt.Sprintf("%s%s %s", prefix, instance.Key.DisplayString(), instance.HumanReadableDescription())
	result := []string{entry}
	for _, slave := range replicationMap[instance] {
		slavesResult := getAsciiTopologyEntry(depth+1, slave, replicationMap)
		result = append(result, slavesResult...)
	}
	return result
}

// AsciiTopology returns a string representation of the topology of given instance.
func AsciiTopology(instanceKey *InstanceKey) (string, error) {
	instance, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return "", err
	}
	instances, err := ReadClusterInstances(instance.ClusterName)
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
	resultArray := getAsciiTopologyEntry(0, masterInstance, replicationMap)
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
		return instance, log.Errorf("Cannot GetInstanceMaster() for %+v. error=%+v", instance.Key, err)
	}

	if !master.IsSlave() {
		return instance, errors.New(fmt.Sprintf("master is not a slave itself: %+v", master.Key))
	}

	if canReplicate, err := instance.CanReplicateFrom(master); canReplicate == false {
		return instance, err
	}

	log.Infof("Will move %+v up the topology", *instanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "move up"); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(&master.Key, GetMaintenanceOwner(), fmt.Sprintf("child %+v moves up", *instanceKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", master.Key))
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

	instance, err = ChangeMasterTo(instanceKey, &master.MasterKey, &master.ExecBinlogCoordinates)
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
		return res, instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey)), errs
	}
	_, err = GetInstanceMaster(instance)
	if err != nil {
		return res, instance, log.Errorf("Cannot GetInstanceMaster() for %+v. error=%+v", instance.Key, err), errs
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
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	for _, slave := range slaves {
		if maintenanceToken, merr := BeginMaintenance(&slave.Key, GetMaintenanceOwner(), fmt.Sprintf("%+v moves up", slave.Key)); merr != nil {
			err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", slave.Key))
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

			var slaveErr error = nil
			ExecuteOnTopology(func() {
				if canReplicate, err := slave.CanReplicateFrom(instance); canReplicate == false || err != nil {
					slaveErr = err
					return
				}
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

				slave, err = ChangeMasterTo(&slave.Key, &instance.MasterKey, &instance.ExecBinlogCoordinates)
				if err != nil {
					slaveErr = err
					return
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
	for _ = range slaves {
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

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("move below %+v", *siblingKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(siblingKey, GetMaintenanceOwner(), fmt.Sprintf("%+v moves below this", *instanceKey)); merr != nil {
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

// Repoint connects a slave to a master using its exact same executing coordinates.
// The given masterKey can be null, in which case the existing master is used.
// Two use cases:
// - masterKey is nil: use case is corrupted relay logs on slave
// - masterKey is not nil: using MaxScale and Binlog servers (coordinates remain the same)
func Repoint(instanceKey *InstanceKey, masterKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", *instanceKey))
	}

	if masterKey == nil {
		masterKey = &instance.MasterKey
	}
	master, err := ReadTopologyInstance(masterKey)
	if err != nil {
		return instance, err
	}
	if canReplicate, err := instance.CanReplicateFrom(master); !canReplicate {
		return instance, err
	}

	log.Infof("Will repoint %+v to master %+v", *instanceKey, *masterKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "repoint"); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		goto Cleanup
	}

	instance, err = ChangeMasterTo(instanceKey, masterKey, &instance.ExecBinlogCoordinates)
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

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("make co-master of %+v", master.Key)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
		goto Cleanup
	} else {
		defer EndMaintenance(maintenanceToken)
	}
	if maintenanceToken, merr := BeginMaintenance(&master.Key, GetMaintenanceOwner(), fmt.Sprintf("%+v turns into co-master of this", *instanceKey)); merr != nil {
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

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "reset slave"); merr != nil {
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

// DetachSlaveOperation will detach a slave from its master by forcibly corrupting its replication coordinates
func DetachSlaveOperation(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}

	log.Infof("Will detach %+v", instanceKey)

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), "detach slave"); merr != nil {
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

// FindLastPseudoGTIDEntry will search an instance's binary logs or relay logs for the last pseudo-GTID entry,
// and return found coordinates as well as entry text
func FindLastPseudoGTIDEntry(instance *Instance, recordedInstanceRelayLogCoordinates BinlogCoordinates, exhaustiveSearch bool) (*BinlogCoordinates, string, error) {
	var instancePseudoGtidText string
	var instancePseudoGtidCoordinates *BinlogCoordinates
	var err error = nil

	if instance.LogBinEnabled && instance.LogSlaveUpdatesEnabled {
		// Well no need to search this instance's binary logs if it doesn't have any...
		// With regard log-slave-updates, some edge cases are possible, like having this instance's log-slave-updates
		// enabled/disabled (of course having restarted it)
		// The approach is not to take chances. If log-slave-updates is disabled, fail and go for relay-logs.
		// If log-slave-updates was just enabled then possibly no pseudo-gtid is found, and so again we will go
		// for relay logs.
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
// The "other instance" could be the sibling of the moving instance any of its ancestors. It may actuall be
// a cousin of some sort (though unlikely). The only important thing is that the "other instance" is more
// advanced in replication than given instance.
func MatchBelow(instanceKey, otherKey *InstanceKey, requireInstanceMaintenance bool, requireOtherMaintenance bool) (*Instance, *BinlogCoordinates, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, nil, err
	}
	if instanceKey.Equals(otherKey) {
		return instance, nil, errors.New(fmt.Sprintf("MatchBelow: attempt to match an instance below itself %+v", *instanceKey))
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
	log.Infof("Will match %+v below %+v", *instanceKey, *otherKey)

	var instancePseudoGtidText string
	var instancePseudoGtidCoordinates *BinlogCoordinates
	var otherInstancePseudoGtidCoordinates *BinlogCoordinates
	var nextBinlogCoordinatesToMatch *BinlogCoordinates
	var recordedInstanceRelayLogCoordinates BinlogCoordinates
	var countMatchedEvents int = 0

	if requireInstanceMaintenance {
		if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("match below %+v", *otherKey)); merr != nil {
			err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
			goto Cleanup
		} else {
			defer EndMaintenance(maintenanceToken)
		}
	}
	if requireOtherMaintenance {
		if maintenanceToken, merr := BeginMaintenance(otherKey, GetMaintenanceOwner(), fmt.Sprintf("%+v matches below this", *instanceKey)); merr != nil {
			err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *otherKey))
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
	instancePseudoGtidCoordinates, instancePseudoGtidText, err = FindLastPseudoGTIDEntry(instance, recordedInstanceRelayLogCoordinates, true)

	if err != nil {
		goto Cleanup
	}
	otherInstancePseudoGtidCoordinates, err = SearchPseudoGTIDEntryInInstance(otherInstance, instancePseudoGtidText)
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
		err = errors.New(fmt.Sprintf("Unexpected: 0 events processed while iterating logs. Something went wrong; aborting. nextBinlogCoordinatesToMatch: %+v", nextBinlogCoordinatesToMatch))
		goto Cleanup
	}
	log.Debugf("%+v will match below %+v at %+v; validated events: %d", *instanceKey, *otherKey, *nextBinlogCoordinatesToMatch, countMatchedEvents)

	// Drum roll......
	instance, err = ChangeMasterTo(instanceKey, otherKey, nextBinlogCoordinatesToMatch)
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
func RematchSlave(instanceKey *InstanceKey, requireInstanceMaintenance bool, requireOtherMaintenance bool) (*Instance, *BinlogCoordinates, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, nil, err
	}
	masterInstance, found, err := ReadInstance(&instance.MasterKey)
	if err != nil || !found {
		return instance, nil, err
	}
	return MatchBelow(instanceKey, &masterInstance.Key, requireInstanceMaintenance, requireOtherMaintenance)
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
			return instance, errors.New(fmt.Sprintf("MakeMaster: instance's master %+v seems to be replicating", masterInstance.Key))
		}
		if masterInstance.IsLastCheckValid {
			return instance, errors.New(fmt.Sprintf("MakeMaster: instance's master %+v seems to be accessible", masterInstance.Key))
		}
	}
	// If err == nil this is "good": that means the master is inaccessible... So it's OK to do the promotion
	if !instance.SQLThreadUpToDate() {
		return instance, errors.New(fmt.Sprintf("MakeMaster: instance's SQL thread must be up-to-date with I/O thread for %+v", *instanceKey))
	}
	siblings, err := ReadSlaveInstances(&masterInstance.Key)
	if err != nil {
		return instance, err
	}
	for _, sibling := range siblings {
		if instance.ExecBinlogCoordinates.SmallerThan(&sibling.ExecBinlogCoordinates) {
			return instance, errors.New(fmt.Sprintf("MakeMaster: instance %+v has more advanced sibling: %+v", *instanceKey, sibling.Key))
		}
	}

	if maintenanceToken, merr := BeginMaintenance(instanceKey, GetMaintenanceOwner(), fmt.Sprintf("siblings match below this", *instanceKey)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", *instanceKey))
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
			return instance, errors.New(fmt.Sprintf("MakeMaster: instance %+v has more advanced sibling: %+v", *instanceKey, sibling.Key))
		}
	}

	instance, err = StopSlaveNicely(instanceKey, 0)
	if err != nil {
		goto Cleanup
	}

	_, _, err = MatchBelow(instanceKey, &grandparentInstance.Key, true, true)
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
	for i := len(instances) - 1; i >= 0; i-- {
		if instances[i].Key.Equals(instanceKey) {
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

	slaves = removeInstance(slaves, belowKey)

	belowInstance, err := ReadTopologyInstance(belowKey)
	if err != nil {
		// Can't access the server below which we need to match ==> can't move slaves
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
	knownCoordinatesMap := make(map[BinlogCoordinates](*BinlogCoordinates))
	for _, slave := range slaves {
		slave := slave
		slaveBuckets[slave.ExecBinlogCoordinates] = append(slaveBuckets[slave.ExecBinlogCoordinates], slave)
	}
	log.Debugf("MultiMatchBelow: %d slaves merged into %d buckets", len(slaves), len(slaveBuckets))
	for bucket, bucketSlaves := range slaveBuckets {
		log.Debugf("+- bucket: %+v, %d slaves", bucket, len(bucketSlaves))
	}
	matchedSlaves := make(map[InstanceKey]bool)
	barrier := make(chan *BinlogCoordinates)
	// Now go over the buckets, and try a single slave from each bucket
	// (though if one results with an error, synchronuously-for-that-bucket continue to the next slave in bucket)

	if maintenanceToken, merr := BeginMaintenance(&belowInstance.Key, GetMaintenanceOwner(), fmt.Sprintf("slaves multi match below this: %+v", belowInstance.Key)); merr != nil {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance on %+v", belowInstance.Key))
		return res, belowInstance, err, errs
	} else {
		defer EndMaintenance(maintenanceToken)
	}

	for execCoordinates, bucketSlaves := range slaveBuckets {
		execCoordinates := execCoordinates
		bucketSlaves := bucketSlaves
		// Buckets concurrent
		go func() {
			defer func() { barrier <- &execCoordinates }()
			func() {
				for _, slave := range bucketSlaves {
					slave := slave
					var slaveErr error = nil
					var matchedCoordinates *BinlogCoordinates
					log.Debugf("MultiMatchBelow: attempting slave %+v in bucket %+v", slave.Key, execCoordinates)
					ExecuteOnTopology(func() {
						_, matchedCoordinates, slaveErr = MatchBelow(&slave.Key, &belowInstance.Key, true, false)
					})
					log.Debugf("MultiMatchBelow: match result: %+v, %+v", matchedCoordinates, slaveErr)

					if slaveErr == nil {
						// Success! We matched a slave of this bucket
						knownCoordinatesMap[execCoordinates] = matchedCoordinates
						matchedSlaves[slave.Key] = true
						log.Debugf("MultiMatchBelow: matched slave %+v in bucket %+v", slave.Key, execCoordinates)
						return
					}

					func() {
						slaveMutex <- true
						defer func() { <-slaveMutex }()
						errs = append(errs, slaveErr)
					}()
					log.Errore(slaveErr)
					// Failure: some unknown problem with bucket slave. Let's try the next one (continue loop)
				}
			}()
		}()
	}
	for _ = range slaveBuckets {
		<-barrier
	}
	// Now that we've handled the representative slaves-per-bucket, let's go over all other slaves
	for _, slave := range slaves {
		slave := slave
		if _, found := matchedSlaves[slave.Key]; found {
			// Already matched this slave
			continue
		}
		matchedCoordinates, found := knownCoordinatesMap[slave.ExecBinlogCoordinates]
		if !found {
			log.Errorf("MultiMatchBelow: Cannot match up %+v since cannot lookup matched coordinates for %+v", slave.Key, slave.ExecBinlogCoordinates)
			continue
		}
		if matchedCoordinates == nil {
			log.Errorf("MultiMatchBelow: found nil matchedCoordinates in bucket %+v", slave.ExecBinlogCoordinates)
			continue
		}
		log.Debugf("MultiMatchBelow: Will match up %+v to previously matched master coordinates %+v", slave.Key, matchedCoordinates)
		if _, err := ChangeMasterTo(&slave.Key, &belowInstance.Key, matchedCoordinates); err == nil {
			StartSlave(&slave.Key)
			matchedSlaves[slave.Key] = true
		} else {
			func() {
				slaveMutex <- true
				defer func() { <-slaveMutex }()
				errs = append(errs, err)
			}()
			log.Errorf("MultiMatchBelow: Cannot match up %+v: error is %+v", slave.Key, err)
			continue
		}
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

	// slaves involved
	slaves, err := ReadSlaveInstances(masterKey)
	if err != nil {
		return res, belowInstance, err, errs
	}
	slaves = filterInstancesByPattern(slaves, pattern)
	matchedSlaves, belowInstance, err, errs := MultiMatchBelow(slaves, &belowInstance.Key, false)

	if len(matchedSlaves) != len(slaves) {
		err = errors.New(fmt.Sprintf("MultiMatchSlaves: only matched %d out of %d slaves of %+v", len(matchedSlaves), len(slaves), *masterKey))
	}

	return matchedSlaves, belowInstance, err, errs
}

// MatchUp will move a slave up the replication chain, so that it becomes sibling of its master, via Pseudo-GTID
func MatchUp(instanceKey *InstanceKey, requireInstanceMaintenance bool, requireOtherMaintenance bool) (*Instance, *BinlogCoordinates, error) {
	instance, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return nil, nil, err
	}
	if !instance.IsSlave() {
		return instance, nil, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	master, found, err := ReadInstance(&instance.MasterKey)
	if err != nil || !found {
		return instance, nil, log.Errorf("Cannot get master for %+v. error=%+v", instance.Key, err)
	}

	if !master.IsSlave() {
		return instance, nil, errors.New(fmt.Sprintf("master is not a slave itself: %+v", master.Key))
	}

	return MatchBelow(instanceKey, &master.MasterKey, requireInstanceMaintenance, requireOtherMaintenance)
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
	if !slave.LogBinEnabled {
		return false
	}
	if !slave.LogSlaveUpdatesEnabled {
		return false
	}
	if config.Config.DenyAutoPromotionHostnamePattern != "" {
		if matched, _ := regexp.MatchString(config.Config.DenyAutoPromotionHostnamePattern, slave.Key.Hostname); matched {
			return false
		}
	}
	return true
}

// GetCandidateSlave chooses the best slave to promote given a (possibly dead) master
func GetCandidateSlave(masterKey *InstanceKey, forRematchPurposes bool) (*Instance, [](*Instance), [](*Instance), [](*Instance), error) {
	var candidateSlave *Instance = nil
	aheadSlaves := [](*Instance){}
	equalSlaves := [](*Instance){}
	laterSlaves := [](*Instance){}

	slaves, err := sortedSlaves(masterKey, forRematchPurposes)
	if err != nil {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, err
	}
	if len(slaves) == 0 {
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, errors.New(fmt.Sprintf("No slaves found for %+v", *masterKey))
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
		return candidateSlave, aheadSlaves, equalSlaves, laterSlaves, errors.New(fmt.Sprintf("No slaves found with log_slave_updates for %+v", *masterKey))
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
func RegroupSlaves(masterKey *InstanceKey, onCandidateSlaveChosen func(*Instance)) ([](*Instance), [](*Instance), [](*Instance), *Instance, error) {
	candidateSlave, aheadSlaves, equalSlaves, laterSlaves, err := GetCandidateSlave(masterKey, true)
	if err != nil {
		return aheadSlaves, equalSlaves, laterSlaves, nil, err
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
				ChangeMasterTo(&slave.Key, &candidateSlave.Key, &candidateSlave.SelfBinlogCoordinates)
			})
		}()
	}
	for _, _ = range equalSlaves {
		<-barrier
	}

	log.Debugf("RegroupSlaves: multi matching %d later slaves", len(laterSlaves))
	// As for the laterSlaves, we'll have to apply pseudo GTID
	laterSlaves, instance, err, _ := MultiMatchBelow(laterSlaves, &candidateSlave.Key, true)

	barrier = make(chan *InstanceKey)
	operatedSlaves := append(equalSlaves, candidateSlave)
	operatedSlaves = append(operatedSlaves, laterSlaves...)
	log.Debugf("RegroupSlaves: starting %d slaves", len(operatedSlaves))
	for _, slave := range operatedSlaves {
		slave := slave
		// This slave has the exact same executing coordinates as the candidate slave. This slave
		// is *extremely* easy to attach below the candidate slave!
		go func() {
			defer func() { barrier <- &candidateSlave.Key }()
			ExecuteOnTopology(func() {
				StartSlave(&slave.Key)
			})
		}()
	}
	for _, _ = range operatedSlaves {
		<-barrier
	}

	log.Debugf("RegroupSlaves: done")
	return aheadSlaves, equalSlaves, laterSlaves, instance, err
}
