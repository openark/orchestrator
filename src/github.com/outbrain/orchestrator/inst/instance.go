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
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/config"
	"net"
	"strconv"
	"strings"
)

// GetCNAME resolves an IP or hostname into a normalized valid CNAME
func GetCNAME(hostName string) (string, error) {
	res, err := net.LookupCNAME(hostName)
	if err != nil {
		return hostName, err
	}
	res = strings.TrimRight(res, ".")
	return res, nil
}

const InvalidPort = 65535

// InstanceKey is an instance indicator, identifued by hostname and port
type InstanceKey struct {
	Hostname string
	Port     int
}

// NewInstanceKeyFromStrings creates a new InstanceKey by resolving hostname and port.
// hostname is normalized via GetCNAME. port is tested to be valid integer.
func NewInstanceKeyFromStrings(hostname string, port string) (*InstanceKey, error) {
	instanceKey := &InstanceKey{}
	var err error
	if instanceKey.Hostname, err = GetCNAME(hostname); err != nil {
		return instanceKey, err
	}

	if instanceKey.Port, err = strconv.Atoi(port); err != nil {
		return instanceKey, errors.New(fmt.Sprintf("Invalid port: %s", port))
	}
	return instanceKey, nil
}

// ParseInstanceKey will parse an InstanceKey from a string representation such as 127.0.0.1:3306
func ParseInstanceKey(hostPort string) (*InstanceKey, error) {
	tokens := strings.SplitN(hostPort, ":", 2)
	if len(tokens) != 2 {
		return nil, errors.New(fmt.Sprintf("Cannpt parse InstanceKey from %s. Expected format is host:port", hostPort))
	}
	return NewInstanceKeyFromStrings(tokens[0], tokens[1])
}

// Formalize this key by getting CNAME for hostname
func (this *InstanceKey) Formalize() *InstanceKey {
	this.Hostname, _ = GetCNAME(this.Hostname)
	return this
}

// Equals tests equality between this key and another key
func (this *InstanceKey) Equals(other *InstanceKey) bool {
	return this.Hostname == other.Hostname && this.Port == other.Port
}

// IsValid uses simple heuristics to see whether this key represents an actual instance
func (this *InstanceKey) IsValid() bool {
	return len(this.Hostname) > 0 && this.Port > 0
}

// DisplayString returns a user-friendly string representation of this key
func (this *InstanceKey) DisplayString() string {
	return fmt.Sprintf("%s:%d", this.Hostname, this.Port)
}

// BinlogCoordinates described binary log coordinates in the form of log file & log position.
type BinlogCoordinates struct {
	LogFile string
	LogPos  int64
}

// Equals tests equality of this corrdinate and another one.
func (this *BinlogCoordinates) Equals(other *BinlogCoordinates) bool {
	return this.LogFile == other.LogFile && this.LogPos == other.LogPos
}

// SmallerThan returns true if this coordinate is smaller than the other.
func (this *BinlogCoordinates) SmallerThan(other *BinlogCoordinates) bool {
	if this.LogFile < other.LogFile {
		return true
	}
	if this.LogFile == other.LogFile && this.LogPos < other.LogPos {
		return true
	}
	return false
}

// InstanceKeyMap is a convenience struct for listing InstanceKey-s
type InstanceKeyMap map[InstanceKey]bool

// GetInstanceKeys returns keys in this map in the form of an array
func (this *InstanceKeyMap) GetInstanceKeys() []InstanceKey {
	res := []InstanceKey{}
	for key, _ := range *this {
		res = append(res, key)
	}
	return res
}

// MarshalJSON will marshal this map as JSON
func (this *InstanceKeyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.GetInstanceKeys())
}

// Instance represents a database instance, including its current configuration & status.
// It presents important replication configuration and detailed replication status.
type Instance struct {
	Key                    InstanceKey
	ServerID               uint
	Version                string
	ReadOnly               bool
	Binlog_format          string
	LogBinEnabled          bool
	LogSlaveUpdatesEnabled bool
	SelfBinlogCoordinates  BinlogCoordinates
	MasterKey              InstanceKey
	Slave_SQL_Running      bool
	Slave_IO_Running       bool
	ReadBinlogCoordinates  BinlogCoordinates
	ExecBinlogCoordinates  BinlogCoordinates
	LastSQLError           string
	LastIOError            string
	SecondsBehindMaster    sql.NullInt64
	SlaveLagSeconds        sql.NullInt64
	SlaveHosts             InstanceKeyMap
	ClusterName            string

	IsLastCheckValid     bool
	IsUpToDate           bool
	IsRecentlyChecked    bool
	SecondsSinceLastSeen sql.NullInt64
	CountMySQLSnapshots  int
}

// NewInstance creates a new, empty instance
func NewInstance() *Instance {
	return &Instance{
		SlaveHosts: make(map[InstanceKey]bool),
	}
}

// Equals tests that this instance is the same instance as other. The function does not test
// configuration or status.
func (this *Instance) Equals(other *Instance) bool {
	return this.Key == other.Key
}

// MajorVersion returns this instance's major version number (e.g. for 5.5.36 it returns "5.5")
func (this *Instance) MajorVersion() []string {
	return strings.Split(this.Version, ".")[:2]
}

// MajorVersion tests this instance against another and returns true if this instance is of a smaller "major" varsion.
// e.g. 5.5.36 is NOT a smaller major version as comapred to 5.5.36, but IS as compared to 5.6.9
func (this *Instance) IsSmallerMajorVersion(other *Instance) bool {
	thisMajorVersion := this.MajorVersion()
	otherMajorVersion := other.MajorVersion()
	for i := 0; i < len(thisMajorVersion); i++ {
		this_token, _ := strconv.Atoi(thisMajorVersion[i])
		other_token, _ := strconv.Atoi(otherMajorVersion[i])
		if this_token < other_token {
			return true
		}
	}
	return false
}

// IsSlave makes simple heuristics to decide whether this insatnce is a slave of another instance
func (this *Instance) IsSlave() bool {
	return this.MasterKey.Hostname != "" && this.MasterKey.Port != 0 && this.MasterKey.Port != InvalidPort && this.ReadBinlogCoordinates.LogFile != ""
}

// SlaveRunning returns true when this instance's status is of a replicating slave.
func (this *Instance) SlaveRunning() bool {
	return this.IsSlave() && this.Slave_SQL_Running && this.Slave_IO_Running
}

// SQLThreadUpToDate returns true when the instance had consumed all relay logs.
func (this *Instance) SQLThreadUpToDate() bool {
	return this.ReadBinlogCoordinates.Equals(&this.ExecBinlogCoordinates)
}

// AddSlaveKey adds a slave to the list of this instance's slaves.
func (this *Instance) AddSlaveKey(slaveKey *InstanceKey) {
	this.SlaveHosts[*slaveKey] = true
}

// GetSlaveHostsAsJson Marshals slaves list a JSON
func (this *Instance) GetSlaveHostsAsJson() string {
	blob, _ := this.SlaveHosts.MarshalJSON()
	return string(blob)
}

// ReadSlaveHostsFromJson unmarshalls a json to read list of slaves
func (this *Instance) ReadSlaveHostsFromJson(jsonString string) error {
	var keys []InstanceKey
	err := json.Unmarshal([]byte(jsonString), &keys)
	if err != nil {
		return log.Errore(err)
	}

	this.SlaveHosts = make(map[InstanceKey]bool)
	for _, key := range keys {
		this.AddSlaveKey(&key)
	}
	return err
}

// IsSlaveOf returns true if this instance claims to replicate from given master
func (this *Instance) IsSlaveOf(master *Instance) bool {
	return this.MasterKey.Equals(&master.Key)
}

// IsSlaveOf returns true if this i supposed master of given slave
func (this *Instance) IsMasterOf(slave *Instance) bool {
	return slave.IsSlaveOf(this)
}

// CanReplicateFrom uses heursitics to decide whether this instacne can practically replicate from other instance.
// Checks are made to binlog format, version number, binary logs etc.
func (this *Instance) CanReplicateFrom(other *Instance) (bool, error) {
	if !other.LogBinEnabled {
		return false, errors.New(fmt.Sprintf("instance does not have binary logs enabled: %+v", other.Key))
	}
	if !other.LogSlaveUpdatesEnabled {
		return false, errors.New(fmt.Sprintf("instance does not have log_slave_updates enabled: %+v", other.Key))
	}
	if this.IsSmallerMajorVersion(other) {
		return false, errors.New(fmt.Sprintf("instance %+v has version %s, which is lower than %s on %+v ", this.Key, this.Version, other.Version, other.Key))
	}
	if this.LogBinEnabled && this.LogSlaveUpdatesEnabled {
		if this.Binlog_format == "STATEMENT" && (other.Binlog_format == "ROW" || other.Binlog_format == "MIXED") {
			return false, errors.New(fmt.Sprintf("Cannot replicate from ROW/MIXED binlog format on %+v to STATEMENT on %+v", other.Key, this.Key))
		}
		if this.Binlog_format == "MIXED" && other.Binlog_format == "ROW" {
			return false, errors.New(fmt.Sprintf("Cannot replicate from ROW binlog format on %+v to MIXED on %+v", other.Key, this.Key))
		}
	}
	if this.ServerID == other.ServerID {
		return false, errors.New(fmt.Sprintf("Identical server id: %+v, %+v both have %d", other.Key, this.Key, this.ServerID))
	}
	return true, nil
}

// CanMove returns true if this instance's state allows it to be repositioned. For example,
// if this instance lags too much, it will not be moveable.
func (this *Instance) CanMove() (bool, error) {
	if !this.IsLastCheckValid {
		return false, errors.New(fmt.Sprintf("%+v: last check invalid", this.Key))
	}
	if !this.IsRecentlyChecked {
		return false, errors.New(fmt.Sprintf("%+v: not recently checked", this.Key))
	}
	if !this.Slave_SQL_Running {
		return false, errors.New(fmt.Sprintf("%+v: instance is not replicating", this.Key))
	}
	if !this.Slave_IO_Running {
		return false, errors.New(fmt.Sprintf("%+v: instance is not replicating", this.Key))
	}
	if !this.SecondsBehindMaster.Valid {
		return false, errors.New(fmt.Sprintf("%+v: cannot determine slave lag", this.Key))
	}
	if this.SecondsBehindMaster.Int64 > int64(config.Config.ReasonableMaintenanceReplicationLagSeconds) {
		return false, errors.New(fmt.Sprintf("%+v: lags too much", this.Key))
	}
	return true, nil
}

// CanMoveAsCoMaster returns true if this instance's state allows it to be repositioned.
func (this *Instance) CanMoveAsCoMaster() (bool, error) {
	if !this.IsLastCheckValid {
		return false, errors.New(fmt.Sprintf("%+v: last check invalid", this.Key))
	}
	if !this.IsRecentlyChecked {
		return false, errors.New(fmt.Sprintf("%+v: not recently checked", this.Key))
	}
	if this.Slave_SQL_Running {
		return false, errors.New(fmt.Sprintf("%+v: instance is replicating", this.Key))
	}
	if this.Slave_IO_Running {
		return false, errors.New(fmt.Sprintf("%+v: instance is replicating", this.Key))
	}
	return true, nil
}

// CanMoveViaMatch returns true if this instance's state allows it to be repositioned via pseudo-GTID matching
func (this *Instance) CanMoveViaMatch() (bool, error) {
	if !this.IsLastCheckValid {
		return false, errors.New(fmt.Sprintf("%+v: last check invalid", this.Key))
	}
	if !this.IsRecentlyChecked {
		return false, errors.New(fmt.Sprintf("%+v: not recently checked", this.Key))
	}
	return true, nil
}
