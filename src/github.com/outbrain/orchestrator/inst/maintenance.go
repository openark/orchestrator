package inst

import (

)

// Maintenance indicates a maintenance entry (also in the database)
type Maintenance struct {
	MaintenanceId		uint
	Key					InstanceKey
	BeginTimestamp		string
	SecondsElapsed		uint
	IsActive			bool
	Owner				string
	Reason				string
}
