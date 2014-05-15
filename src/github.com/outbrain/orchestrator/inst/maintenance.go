package inst

import (

)


type Maintenance struct {
	MaintenanceId		uint
	Key					InstanceKey
	BeginTimestamp		string
	SecondsSinceBegun	uint
	maintenance_active	bool
	Owner				string
	Reason				string
}
