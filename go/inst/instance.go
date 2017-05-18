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
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/math"
)

// CandidatePromotionRule describe the promotion preference/rule for an instance.
// It maps to promotion_rule column in candidate_database_instance
type CandidatePromotionRule string

const (
	MustPromoteRule      CandidatePromotionRule = "must"
	PreferPromoteRule                           = "prefer"
	NeutralPromoteRule                          = "neutral"
	PreferNotPromoteRule                        = "prefer_not"
	MustNotPromoteRule                          = "must_not"
)

const ReasonableDiscoveryLatency = 500 * time.Millisecond

// ParseCandidatePromotionRule returns a CandidatePromotionRule by name.
// It returns an error if there is no known rule by the given name.
func ParseCandidatePromotionRule(ruleName string) (CandidatePromotionRule, error) {
	switch ruleName {
	case "prefer", "neutral", "must_not":
		return CandidatePromotionRule(ruleName), nil
	case "must", "prefer_not":
		return CandidatePromotionRule(""), fmt.Errorf("CandidatePromotionRule: %v not supported yet", ruleName)
	default:
		return CandidatePromotionRule(""), fmt.Errorf("Invalid CandidatePromotionRule: %v", ruleName)
	}
}

// Instance represents a database instance, including its current configuration & status.
// It presents important replication configuration and detailed replication status.
type Instance struct {
	Key                    InstanceKey
	InstanceAlias          string
	Uptime                 uint
	ServerID               uint
	ServerUUID             string
	Version                string
	VersionComment         string
	FlavorName             string
	ReadOnly               bool
	Binlog_format          string
	BinlogRowImage         string
	LogBinEnabled          bool
	LogSlaveUpdatesEnabled bool
	SelfBinlogCoordinates  BinlogCoordinates
	MasterKey              InstanceKey
	IsDetachedMaster       bool
	Slave_SQL_Running      bool
	Slave_IO_Running       bool
	HasReplicationFilters  bool
	SupportsOracleGTID     bool
	UsingOracleGTID        bool
	UsingMariaDBGTID       bool
	UsingPseudoGTID        bool
	ReadBinlogCoordinates  BinlogCoordinates
	ExecBinlogCoordinates  BinlogCoordinates
	IsDetached             bool
	RelaylogCoordinates    BinlogCoordinates
	LastSQLError           string
	LastIOError            string
	SecondsBehindMaster    sql.NullInt64
	SQLDelay               uint
	ExecutedGtidSet        string
	GtidPurged             string

	SlaveLagSeconds                 sql.NullInt64
	SlaveHosts                      InstanceKeyMap
	ClusterName                     string
	SuggestedClusterAlias           string
	DataCenter                      string
	PhysicalEnvironment             string
	ReplicationDepth                uint
	IsCoMaster                      bool
	HasReplicationCredentials       bool
	ReplicationCredentialsAvailable bool
	SemiSyncEnforced                bool

	LastSeenTimestamp    string
	IsLastCheckValid     bool
	IsUpToDate           bool
	IsRecentlyChecked    bool
	SecondsSinceLastSeen sql.NullInt64
	CountMySQLSnapshots  int

	// Careful. IsCandidate and PromotionRule are used together
	// and probably need to be merged. IsCandidate's value may
	// be picked up from daabase_candidate_instance's value when
	// reading an instance from the db.
	IsCandidate          bool
	PromotionRule        CandidatePromotionRule
	IsDowntimed          bool
	DowntimeReason       string
	DowntimeOwner        string
	DowntimeEndTimestamp string
	ElapsedDowntime      time.Duration
	UnresolvedHostname   string
	AllowTLS             bool

	LastDiscoveryLatency time.Duration
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
	return MajorVersion(this.Version)
}

// MajorVersion returns this instance's major version number (e.g. for 5.5.36 it returns "5.5")
func (this *Instance) MajorVersionString() string {
	return strings.Join(this.MajorVersion(), ".")
}

func (this *Instance) IsMySQL51() bool {
	return this.MajorVersionString() == "5.1"
}

func (this *Instance) IsMySQL55() bool {
	return this.MajorVersionString() == "5.5"
}

func (this *Instance) IsMySQL56() bool {
	return this.MajorVersionString() == "5.6"
}

func (this *Instance) IsMySQL57() bool {
	return this.MajorVersionString() == "5.7"
}

func (this *Instance) IsMySQL80() bool {
	return this.MajorVersionString() == "8.0"
}

// IsSmallerBinlogFormat returns true when this instance's binlgo format is
// "smaller" than the other's, i.e. binary logs cannot flow from the other instance to this one
func (this *Instance) IsSmallerBinlogFormat(other *Instance) bool {
	return IsSmallerBinlogFormat(this.Binlog_format, other.Binlog_format)
}

// IsSmallerMajorVersion tests this instance against another and returns true if this instance is of a smaller "major" varsion.
// e.g. 5.5.36 is NOT a smaller major version as comapred to 5.5.36, but IS as compared to 5.6.9
func (this *Instance) IsSmallerMajorVersion(other *Instance) bool {
	return IsSmallerMajorVersion(this.Version, other.Version)
}

// IsSmallerMajorVersionByString cehcks if this instance has a smaller major version number than given one
func (this *Instance) IsSmallerMajorVersionByString(otherVersion string) bool {
	return IsSmallerMajorVersion(this.Version, otherVersion)
}

// IsMariaDB checks whether this is any version of MariaDB
func (this *Instance) IsMariaDB() bool {
	return strings.Contains(this.Version, "MariaDB")
}

// IsPercona checks whether this is any version of Percona Server
func (this *Instance) IsPercona() bool {
	return strings.Contains(this.VersionComment, "Percona")
}

// isMaxScale checks whether this is any version of MaxScale
func (this *Instance) isMaxScale() bool {
	return strings.Contains(this.Version, "maxscale")
}

// IsBinlogServer checks whether this is any type of a binlog server (currently only maxscale)
func (this *Instance) IsBinlogServer() bool {
	if this.isMaxScale() {
		return true
	}
	return false
}

// IsOracleMySQL checks whether this is an Oracle MySQL distribution
func (this *Instance) IsOracleMySQL() bool {
	if this.IsMariaDB() {
		return false
	}
	if this.IsPercona() {
		return false
	}
	if this.isMaxScale() {
		return false
	}
	if this.IsBinlogServer() {
		return false
	}
	return true
}

// applyFlavorName
func (this *Instance) applyFlavorName() {
	if this == nil {
		return
	}
	if this.IsOracleMySQL() {
		this.FlavorName = "MySQL"
	} else if this.IsMariaDB() {
		this.FlavorName = "MariaDB"
	} else if this.IsPercona() {
		this.FlavorName = "Percona"
	} else if this.isMaxScale() {
		this.FlavorName = "MaxScale"
	} else {
		this.FlavorName = "unknown"
	}
}

// FlavorNameAndMajorVersion returns a string of the combined
// flavor and major version which is useful in some checks.
func (this *Instance) FlavorNameAndMajorVersion() string {
	if this.FlavorName == "" {
		this.applyFlavorName()
	}

	return this.FlavorName + "-" + this.MajorVersionString()
}

// IsReplica makes simple heuristics to decide whether this insatnce is a replica of another instance
func (this *Instance) IsReplica() bool {
	return this.MasterKey.Hostname != "" && this.MasterKey.Hostname != "_" && this.MasterKey.Port != 0 && (this.ReadBinlogCoordinates.LogFile != "" || this.UsingGTID())
}

// ReplicaRunning returns true when this instance's status is of a replicating replica.
func (this *Instance) ReplicaRunning() bool {
	return this.IsReplica() && this.Slave_SQL_Running && this.Slave_IO_Running
}

// SQLThreadUpToDate returns true when the instance had consumed all relay logs.
func (this *Instance) SQLThreadUpToDate() bool {
	return this.ReadBinlogCoordinates.Equals(&this.ExecBinlogCoordinates)
}

// UsingGTID returns true when this replica is currently replicating via GTID (either Oracle or MariaDB)
func (this *Instance) UsingGTID() bool {
	return this.UsingOracleGTID || this.UsingMariaDBGTID
}

// NextGTID returns the next (Oracle) GTID to be executed. Useful for skipping queries
func (this *Instance) NextGTID() (string, error) {
	if this.ExecutedGtidSet == "" {
		return "", fmt.Errorf("No value found in Executed_Gtid_Set; cannot compute NextGTID")
	}

	firstToken := func(s string, delimiter string) string {
		tokens := strings.Split(s, delimiter)
		return tokens[0]
	}
	lastToken := func(s string, delimiter string) string {
		tokens := strings.Split(s, delimiter)
		return tokens[len(tokens)-1]
	}
	// executed GTID set: 4f6d62ed-df65-11e3-b395-60672090eb04:1,b9b4712a-df64-11e3-b391-60672090eb04:1-6
	executedGTIDsFromMaster := lastToken(this.ExecutedGtidSet, ",")
	// executedGTIDsFromMaster: b9b4712a-df64-11e3-b391-60672090eb04:1-6
	executedRange := lastToken(executedGTIDsFromMaster, ":")
	// executedRange: 1-6
	lastExecutedNumberToken := lastToken(executedRange, "-")
	// lastExecutedNumber: 6
	lastExecutedNumber, err := strconv.Atoi(lastExecutedNumberToken)
	if err != nil {
		return "", err
	}
	nextNumber := lastExecutedNumber + 1
	nextGTID := fmt.Sprintf("%s:%d", firstToken(executedGTIDsFromMaster, ":"), nextNumber)
	return nextGTID, nil
}

// AddReplicaKey adds a replica to the list of this instance's replicas.
func (this *Instance) AddReplicaKey(replicaKey *InstanceKey) {
	this.SlaveHosts.AddKey(*replicaKey)
}

// GetNextBinaryLog returns the successive, if any, binary log file to the one given
func (this *Instance) GetNextBinaryLog(binlogCoordinates BinlogCoordinates) (BinlogCoordinates, error) {
	if binlogCoordinates.LogFile == this.SelfBinlogCoordinates.LogFile {
		return binlogCoordinates, fmt.Errorf("Cannot find next binary log for %+v", binlogCoordinates)
	}
	return binlogCoordinates.NextFileCoordinates()
}

// IsReplicaOf returns true if this instance claims to replicate from given master
func (this *Instance) IsReplicaOf(master *Instance) bool {
	return this.MasterKey.Equals(&master.Key)
}

// IsReplicaOf returns true if this i supposed master of given replica
func (this *Instance) IsMasterOf(replica *Instance) bool {
	return replica.IsReplicaOf(this)
}

// CanReplicateFrom uses heursitics to decide whether this instacne can practically replicate from other instance.
// Checks are made to binlog format, version number, binary logs etc.
func (this *Instance) CanReplicateFrom(other *Instance) (bool, error) {
	if this.Key.Equals(&other.Key) {
		return false, fmt.Errorf("instance cannot replicate from itself: %+v", this.Key)
	}
	if !other.LogBinEnabled {
		return false, fmt.Errorf("instance does not have binary logs enabled: %+v", other.Key)
	}
	if other.IsReplica() {
		if !other.LogSlaveUpdatesEnabled {
			return false, fmt.Errorf("instance does not have log_slave_updates enabled: %+v", other.Key)
		}
		// OK for a master to not have log_slave_updates
		// Not OK for a replica, for it has to relay the logs.
	}
	if this.IsSmallerMajorVersion(other) && !this.IsBinlogServer() {
		return false, fmt.Errorf("instance %+v has version %s, which is lower than %s on %+v ", this.Key, this.Version, other.Version, other.Key)
	}
	if this.LogBinEnabled && this.LogSlaveUpdatesEnabled {
		if this.IsSmallerBinlogFormat(other) {
			return false, fmt.Errorf("Cannot replicate from %+v binlog format on %+v to %+v on %+v", other.Binlog_format, other.Key, this.Binlog_format, this.Key)
		}
	}
	if config.Config.VerifyReplicationFilters {
		if other.HasReplicationFilters && !this.HasReplicationFilters {
			return false, fmt.Errorf("%+v has replication filters", other.Key)
		}
	}
	if this.ServerID == other.ServerID && !this.IsBinlogServer() {
		return false, fmt.Errorf("Identical server id: %+v, %+v both have %d", other.Key, this.Key, this.ServerID)
	}
	return true, nil
}

// HasReasonableMaintenanceReplicationLag returns true when the replica lag is reasonable, and maintenance operations should have a green light to go.
func (this *Instance) HasReasonableMaintenanceReplicationLag() bool {
	// replicas with SQLDelay are a special case
	if this.SQLDelay > 0 {
		return math.AbsInt64(this.SecondsBehindMaster.Int64-int64(this.SQLDelay)) <= int64(config.Config.ReasonableMaintenanceReplicationLagSeconds)
	}
	return this.SecondsBehindMaster.Int64 <= int64(config.Config.ReasonableMaintenanceReplicationLagSeconds)
}

// CanMove returns true if this instance's state allows it to be repositioned. For example,
// if this instance lags too much, it will not be moveable.
func (this *Instance) CanMove() (bool, error) {
	if !this.IsLastCheckValid {
		return false, fmt.Errorf("%+v: last check invalid", this.Key)
	}
	if !this.IsRecentlyChecked {
		return false, fmt.Errorf("%+v: not recently checked", this.Key)
	}
	if !this.Slave_SQL_Running {
		return false, fmt.Errorf("%+v: instance is not replicating", this.Key)
	}
	if !this.Slave_IO_Running {
		return false, fmt.Errorf("%+v: instance is not replicating", this.Key)
	}
	if !this.SecondsBehindMaster.Valid {
		return false, fmt.Errorf("%+v: cannot determine slave lag", this.Key)
	}
	if !this.HasReasonableMaintenanceReplicationLag() {
		return false, fmt.Errorf("%+v: lags too much", this.Key)
	}
	return true, nil
}

// CanMoveAsCoMaster returns true if this instance's state allows it to be repositioned.
func (this *Instance) CanMoveAsCoMaster() (bool, error) {
	if !this.IsLastCheckValid {
		return false, fmt.Errorf("%+v: last check invalid", this.Key)
	}
	if !this.IsRecentlyChecked {
		return false, fmt.Errorf("%+v: not recently checked", this.Key)
	}
	return true, nil
}

// CanMoveViaMatch returns true if this instance's state allows it to be repositioned via pseudo-GTID matching
func (this *Instance) CanMoveViaMatch() (bool, error) {
	if !this.IsLastCheckValid {
		return false, fmt.Errorf("%+v: last check invalid", this.Key)
	}
	if !this.IsRecentlyChecked {
		return false, fmt.Errorf("%+v: not recently checked", this.Key)
	}
	return true, nil
}

// StatusString returns a human readable description of this instance's status
func (this *Instance) StatusString() string {
	if !this.IsLastCheckValid {
		return "invalid"
	}
	if !this.IsRecentlyChecked {
		return "unchecked"
	}
	if this.IsReplica() && !(this.Slave_SQL_Running && this.Slave_IO_Running) {
		return "nonreplicating"
	}
	if this.IsReplica() && !this.HasReasonableMaintenanceReplicationLag() {
		return "lag"
	}
	return "ok"
}

// LagStatusString returns a human readable representation of current lag
func (this *Instance) LagStatusString() string {
	if this.IsDetached {
		return "detached"
	}
	if !this.IsLastCheckValid {
		return "unknown"
	}
	if !this.IsRecentlyChecked {
		return "unknown"
	}
	if this.IsReplica() && !(this.Slave_SQL_Running && this.Slave_IO_Running) {
		return "null"
	}
	if this.IsReplica() && !this.SecondsBehindMaster.Valid {
		return "null"
	}
	if this.IsReplica() && this.SlaveLagSeconds.Int64 > int64(config.Config.ReasonableMaintenanceReplicationLagSeconds) {
		return fmt.Sprintf("%+vs", this.SlaveLagSeconds.Int64)
	}
	return fmt.Sprintf("%+vs", this.SlaveLagSeconds.Int64)
}

// HumanReadableDescription returns a simple readable string describing the status, version,
// etc. properties of this instance
func (this *Instance) HumanReadableDescription() string {
	tokens := []string{}
	tokens = append(tokens, this.LagStatusString())
	tokens = append(tokens, this.StatusString())
	tokens = append(tokens, this.Version)
	if this.ReadOnly {
		tokens = append(tokens, "ro")
	} else {
		tokens = append(tokens, "rw")
	}
	if this.LogBinEnabled {
		tokens = append(tokens, this.Binlog_format)
	} else {
		tokens = append(tokens, "nobinlog")
	}
	if this.LogBinEnabled && this.LogSlaveUpdatesEnabled {
		tokens = append(tokens, ">>")
	}
	if this.UsingGTID() {
		tokens = append(tokens, "GTID")
	}
	if this.UsingPseudoGTID {
		tokens = append(tokens, "P-GTID")
	}
	if this.IsDowntimed {
		tokens = append(tokens, "downtimed")
	}
	description := fmt.Sprintf("[%s]", strings.Join(tokens, ","))
	return description
}
