package inst

import (

)


type Maintenance struct {
	MaintenanceId		uint
	Key					InstanceKey
	BeginTimestamp		string
	SecondsElapsed		uint
	IsActive			bool
	Owner				string
	Reason				string
}
