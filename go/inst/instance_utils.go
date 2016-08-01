/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

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
	"regexp"
	"strconv"
	"strings"
)

var (
	DowntimeLostInRecoveryMessage = "lost-in-recovery"
)

// InstancesByExecBinlogCoordinates is a sortabel type for BinlogCoordinates
type InstancesByExecBinlogCoordinates [](*Instance)

func (this InstancesByExecBinlogCoordinates) Len() int      { return len(this) }
func (this InstancesByExecBinlogCoordinates) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func (this InstancesByExecBinlogCoordinates) Less(i, j int) bool {
	// Returning "true" in this function means [i] is "smaller" than [j],
	// which will lead to [j] be a better candidate for promotion

	// Sh*t happens. We just might get nil while attempting to discover/recover
	if this[i] == nil {
		return false
	}
	if this[j] == nil {
		return true
	}
	if this[i].ExecBinlogCoordinates.Equals(&this[j].ExecBinlogCoordinates) {
		// Secondary sorting: "smaller" if not logging slave updates
		if this[j].LogSlaveUpdatesEnabled && !this[i].LogSlaveUpdatesEnabled {
			return true
		}
		// Next sorting: "smaller" if of higher version (this will be reversed eventually)
		// Idea is that given 5.6 a& 5.7 both of the exact position, we will want to promote
		// the 5.6 on top of 5.7, as the other way around is invalid
		if this[j].IsSmallerMajorVersion(this[i]) {
			return true
		}
		// Next sorting: "smaller" if of larger binlog-format (this will be reversed eventually)
		// Idea is that given ROW & STATEMENT both of the exact position, we will want to promote
		// the STATEMENT on top of ROW, as the other way around is invalid
		if this[j].IsSmallerBinlogFormat(this[i]) {
			return true
		}
	}
	return this[i].ExecBinlogCoordinates.SmallerThan(&this[j].ExecBinlogCoordinates)
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

// removeInstance will remove an instance from a list of instances
func RemoveInstance(instances [](*Instance), instanceKey *InstanceKey) [](*Instance) {
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
func RemoveBinlogServerInstances(instances [](*Instance)) [](*Instance) {
	for i := len(instances) - 1; i >= 0; i-- {
		if instances[i].IsBinlogServer() {
			instances = append(instances[:i], instances[i+1:]...)
		}
	}
	return instances
}

// removeNilInstances
func RemoveNilInstances(instances [](*Instance)) [](*Instance) {
	for i := len(instances) - 1; i >= 0; i-- {
		if instances[i] == nil {
			instances = append(instances[:i], instances[i+1:]...)
		}
	}
	return instances
}

// SemicolonTerminated is a utility function that makes sure a statement is terminated with
// a semicolon, if it isn't already
func SemicolonTerminated(statement string) string {
	statement = strings.TrimSpace(statement)
	statement = strings.TrimRight(statement, ";")
	statement = statement + ";"
	return statement
}

// MajorVersion returns a MySQL major version number (e.g. given "5.5.36" it returns "5.5")
func MajorVersion(version string) []string {
	tokens := strings.Split(version, ".")
	if len(tokens) < 2 {
		return []string{"0", "0"}
	}
	return tokens[:2]
}

// IsSmallerMajorVersion tests two versions against another and returns true if
// the former is a smaller "major" varsion than the latter.
// e.g. 5.5.36 is NOT a smaller major version as comapred to 5.5.40, but IS as compared to 5.6.9
func IsSmallerMajorVersion(version string, otherVersion string) bool {
	thisMajorVersion := MajorVersion(version)
	otherMajorVersion := MajorVersion(otherVersion)
	for i := 0; i < len(thisMajorVersion); i++ {
		thisToken, _ := strconv.Atoi(thisMajorVersion[i])
		otherToken, _ := strconv.Atoi(otherMajorVersion[i])
		if thisToken < otherToken {
			return true
		}
		if thisToken > otherToken {
			return false
		}
	}
	return false
}

// IsSmallerBinlogFormat tests two binlog formats and sees if one is "smaller" than the other.
// "smaller" binlog format means you can replicate from the smaller to the larger.
func IsSmallerBinlogFormat(binlogFormat string, otherBinlogFormat string) bool {
	if binlogFormat == "STATEMENT" {
		return (otherBinlogFormat == "ROW" || otherBinlogFormat == "MIXED")
	}
	if binlogFormat == "MIXED" {
		return otherBinlogFormat == "ROW"
	}
	return false
}
