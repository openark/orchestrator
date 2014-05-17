package inst

import (
	"strconv"
	"strings"
	"net"
	"fmt"
	"errors"
	"database/sql"
	"encoding/json"
	"github.com/outbrain/log"
)


func GetCNAME(hostName string) (string, error) {
	res, err := net.LookupCNAME(hostName);
	if err != nil {
		return hostName, err
	}
	res = strings.TrimRight(res, ".")
	return res, nil
}

type InstanceKey struct {
	Hostname 			string
	Port	 			int
}

func NewInstanceKeyFromStrings(hostname string, port string) (*InstanceKey, error) {
	instanceKey := &InstanceKey{}
	var err error
	if instanceKey.Hostname, err = GetCNAME(hostname); err != nil {return instanceKey, err}

	if instanceKey.Port, err = strconv.Atoi(port); err != nil {
		return instanceKey, errors.New(fmt.Sprintf("Invalid port: %s", port))
	}
	return instanceKey, nil
}

func ParseInstanceKey(hostPort string) (*InstanceKey, error) {
	tokens := strings.SplitN(hostPort, ":", 2)
	if len(tokens) != 2 {
		return nil, errors.New(fmt.Sprintf("Cannpt parse InstanceKey from %s. Expected format is host:port", hostPort))
	}
	return NewInstanceKeyFromStrings(tokens[0], tokens[1])
}

func (this *InstanceKey) Formalize() *InstanceKey {
	this.Hostname, _ = GetCNAME(this.Hostname) 
	return this
}

func (this *InstanceKey) Equals(other *InstanceKey) bool {
	return this.Hostname == other.Hostname && this.Port == other.Port
}

func (this *InstanceKey) IsValid() bool {
	return len(this.Hostname) > 0 && this.Port > 0
}



//
type BinlogCoordinates struct {
	LogFile	string
	LogPos	int64
}


func (this *BinlogCoordinates) Equals(other *BinlogCoordinates) bool {
	return this.LogFile == other.LogFile && this.LogPos == other.LogPos
}


func (this *BinlogCoordinates) SmallerThan(other *BinlogCoordinates) bool {
	if this.LogFile < other.LogFile {
		return true
	}
	if this.LogFile == other.LogFile && this.LogPos < other.LogPos {
		return true
	}
	return false
}

//
type InstanceKeyMap map[InstanceKey]bool


func (this *InstanceKeyMap) GetInstanceKeys() []InstanceKey {
	res := []InstanceKey{}
	for key, _ := range *this {
    	res = append(res, key)
	}
	return res
}

func (this *InstanceKeyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.GetInstanceKeys())
}


type Instance struct {
	Key					InstanceKey
	ServerID			uint
	Version				string
	Binlog_format		string
	LogBinEnabled		bool
	LogSlaveUpdatesEnabled	bool
	SelfBinlogCoordinates	BinlogCoordinates
	MasterKey			InstanceKey
	Slave_SQL_Running	bool
	Slave_IO_Running	bool
	ReadBinlogCoordinates	BinlogCoordinates
	ExecBinlogCoordinates	BinlogCoordinates
	SecondsBehindMaster		sql.NullInt64
	SlaveHosts			InstanceKeyMap
	ClusterName			string
	
	IsLastCheckValid		bool
	IsUpToDate			bool
	SecondsSinceLastSeen	sql.NullInt64
}

func NewInstance() *Instance {
    return &Instance{
    	SlaveHosts: make(map[InstanceKey]bool),
    }
}

func (this *Instance) Equals(other *Instance) bool {
	return this.Key == other.Key
}

func (this *Instance) MajorVersion() []string {
	return strings.Split(this.Version, ".")[:2]
}

func (this *Instance) IsSmallerMajorVersion(other *Instance) bool {
	thisMajorVersion := this.MajorVersion()
	otherMajorVersion := other.MajorVersion()
	for i := 0 ; i < len(thisMajorVersion); i++ {
		this_token, _ := strconv.Atoi(thisMajorVersion[i])
		other_token, _ := strconv.Atoi(otherMajorVersion[i])
		if this_token < other_token {
			return true
		} 
	}
	return false
}

func (this *Instance) IsSlave() bool {
	return this.MasterKey.Hostname != "" && this.MasterKey.Port != 0 && this.ReadBinlogCoordinates.LogFile != ""
}

func (this *Instance) SlaveRunning() bool {
	return this.IsSlave() && this.Slave_SQL_Running && this.Slave_IO_Running
}

func (this *Instance) SQLThreadUpToDate() bool {
	return this.ReadBinlogCoordinates.Equals(&this.ExecBinlogCoordinates)
}


func (this *Instance) AddSlaveKey(slaveKey *InstanceKey) {
	this.SlaveHosts[*slaveKey] = true
}


func (this *Instance) GetSlaveHostsAsJson() string {
	blob, _ := this.SlaveHosts.MarshalJSON()
	return string(blob)
}


func (this *Instance) ReadSlaveHostsFromJson(jsonString string) error {
	var keys []InstanceKey;
	err := json.Unmarshal([]byte(jsonString), &keys)
	if err != nil {return log.Errore(err)}
	
	this.SlaveHosts = make(map[InstanceKey]bool)
	for _, key := range keys {
    	this.AddSlaveKey(&key)
	}
	return err
}


func (this *Instance) IsSlaveOf(master *Instance) bool {
	return this.MasterKey.Equals(&master.Key)
}

func (this *Instance) IsMasterOf(slave *Instance) bool {
	return slave.IsSlaveOf(this)
}

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

