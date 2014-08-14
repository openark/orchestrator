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

package agent

import (

)




// LogicalVolume describes an LVM volume
type LogicalVolume struct {
	Name			string
	GroupName		string
	Path			string
	IsSnapshot		bool
	SnapshotPercent	float64
}


// Mount describes a file system mount point
type Mount struct {
	Path			string
	Device			string
	LVPath			string
	FileSystem		string
	IsMounted		bool
}


// Audit presents a single audit entry (namely in the database)
type Agent struct {
	Hostname			string
	Port				int
	Token				string
	LastSubmitted		string
	AvailableLocalSnapshots	[]string
	AvailableSnapshots		[]string
	LogicalVolumes			[]LogicalVolume
	MountPoint				Mount
}
