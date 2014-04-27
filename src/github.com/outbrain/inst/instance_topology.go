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


// InstancesAreBrothers checks whether both instances are replicating from same master
func InstancesAreBrothers(instance0, instance1 *Instance) bool {
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
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	instance, err = StopSlave(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	instance, err = StartSlaveUntilMasterCoordinates(instanceKey, &master.SelfBinlogCoordinates)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	instance, err = ChangeMasterTo(instanceKey, master.GetMasterInstanceKey(), &master.ExecBinlogCoordinates)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	instance, err = StartSlave(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	master, err = StartSlave(&master.Key)
	if	err	!=	nil	{log.Println(err); return instance, err} 

	return instance, err
}


func MoveBelow(instanceKey, brotherKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	brother, err := ReadTopologyInstance(brotherKey)
	if !InstancesAreBrothers(instance, brother) {
		return instance, errors.New(fmt.Sprintf("instances are not siblings: %+v, %+v", instanceKey, brotherKey))
	}
	
	if canReplicate, err := instance.CanReplicateFrom(brother); !canReplicate {
		return instance, err
	}
	log.Println(fmt.Sprintf("Will move %+v below its sibling %+v", instanceKey, brotherKey)) 

	instance, err = StopSlave(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	brother, err = StopSlave(brotherKey)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	if instance.ExecBinlogCoordinates.SmallerThan(&brother.ExecBinlogCoordinates) {
		instance, err = StartSlaveUntilMasterCoordinates(instanceKey, &brother.ExecBinlogCoordinates)
		if	err	!=	nil	{log.Println(err); return instance, err} 
	} else if brother.ExecBinlogCoordinates.SmallerThan(&instance.ExecBinlogCoordinates) {
		brother, err = StartSlaveUntilMasterCoordinates(brotherKey, &instance.ExecBinlogCoordinates)
		if	err	!=	nil	{log.Println(err); return instance, err} 
	}  
	// At this point both siblings have executed exact same statements and are identical
	 
	instance, err = ChangeMasterTo(instanceKey, &brother.Key, &brother.SelfBinlogCoordinates)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	
	instance, err = StartSlave(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	brother, err = StartSlave(brotherKey)
	if	err	!=	nil	{log.Println(err); return instance, err} 

	return instance, err
}

