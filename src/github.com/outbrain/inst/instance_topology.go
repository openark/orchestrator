package inst

import (
	"fmt"
	"log"
	"errors"
)


func RefreshInstance(instance *Instance) (*Instance, error) {
	instance, err := ReadInstance(&instance.Key)
	if err != nil {
		log.Println(err)
	}
	return instance, err
}

func RefreshTopologyInstance(instanceKey *InstanceKey) error {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {	log.Println(err); return err}
	
	err = WriteInstance(instance, err)
	if err != nil {	log.Println(err); return err}

	return err
}

// GetInstanceMaster synchronously reaches into the replication topology
// and retrieves master's data
func GetInstanceMaster(instance *Instance) (*Instance, error) {
	masterKey := instance.GetMasterInstanceKey()
	master, err := ReadTopologyInstance(masterKey)
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
	return instance0.GetMasterInstanceKey().Equals(instance1.GetMasterInstanceKey())
}


func MoveUp(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	master, err := GetInstanceMaster(instance)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	if !master.IsSlave() {
		return instance, errors.New(fmt.Sprintf("master is not a slave itself: %+v", master.Key))
	}
	
	if canReplicate, err := instance.CanReplicateFrom(master); canReplicate == false {
		return instance, err
	}
	
	log.Println(fmt.Sprintf("Will move %+v up the topology", instanceKey)) 

	master, err = StopSlave(&master.Key)
	if	err	!=	nil	{goto Cleanup} 
	
	instance, err = StopSlave(instanceKey)
	if	err	!=	nil	{goto Cleanup} 
	
	instance, err = StartSlaveUntilMasterCoordinates(instanceKey, &master.SelfBinlogCoordinates)
	if	err	!=	nil	{goto Cleanup} 
	
	instance, err = ChangeMasterTo(instanceKey, master.GetMasterInstanceKey(), &master.ExecBinlogCoordinates)
	if	err	!=	nil	{goto Cleanup} 
	
	instance, err = StartSlave(instanceKey)
	if	err	!=	nil	{goto Cleanup} 
	
	master, err = StartSlave(&master.Key)
	if	err	!=	nil	{goto Cleanup} 
	
	Cleanup:
	instance, _ = StartSlave(instanceKey)
	master, _ = StartSlave(&master.Key)
	if	err	!=	nil	{log.Println(err)} 
	 
	return instance, err
}


func MoveBelow(instanceKey, siblingKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	sibling, err := ReadTopologyInstance(siblingKey)
	if !InstancesAreSiblings(instance, sibling) {
		return instance, errors.New(fmt.Sprintf("instances are not siblings: %+v, %+v", instanceKey, siblingKey))
	}
	
	if canReplicate, err := instance.CanReplicateFrom(sibling); !canReplicate {
		return instance, err
	}
	log.Println(fmt.Sprintf("Will move %+v below its sibling %+v", instanceKey, siblingKey)) 

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
	if	err	!=	nil	{log.Println(err)} 
	 
	return instance, err
}

