package inst

import ()

type MinimalInstance struct {
	Key         InstanceKey
	MasterKey   InstanceKey
	ClusterName string
}

func (this *MinimalInstance) ToInstance() *Instance {
	return &Instance{
		Key:         this.Key,
		MasterKey:   this.MasterKey,
		ClusterName: this.ClusterName,
	}
}
