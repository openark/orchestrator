package inst

import (
	"strconv"
	"strings"
	"net"
)


func GetCNAME(hostName string) string {
	res, _ := net.LookupCNAME(hostName);
	res = strings.TrimRight(res, ".")
	return res
}

type InstanceKey struct {
	Hostname 			string
	Port	 			int
}

func (this *InstanceKey) Equals(other *InstanceKey) bool {
	return this.Hostname == other.Hostname && this.Port == other.Port
}


func ParseInstanceKey(hostPort string) *InstanceKey {
	tokens := strings.SplitN(hostPort, ":", 2)
	if len(tokens) != 2 {
		return nil
	}
	port, err :=  strconv.Atoi(tokens[1])
	if err != nil {
		return nil
	}
	
    return &InstanceKey{
    	Hostname: tokens[0], Port: port,
    }
}


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

type Instance struct {
	Key					InstanceKey
	Version				string
	Binlog_format		string
	SelfBinlogCoordinates	BinlogCoordinates
	Master_Host			string
	Master_Port			int
	Slave_SQL_Running	bool
	Slave_IO_Running	bool
	ReadBinlogCoordinates	BinlogCoordinates
	ExecBinlogCoordinates	BinlogCoordinates
	SecondsBehindMaster	int
	SlaveHosts			map[string]string
}

func NewInstance() *Instance {
    return &Instance{
    	SlaveHosts: make(map[string]string),
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
	return this.Master_Host != "" && this.ReadBinlogCoordinates.LogFile != ""
}

func (this *Instance) SlaveRunning() bool {
	return this.IsSlave() && this.Slave_SQL_Running && this.Slave_IO_Running
}

func (this *Instance) SQLThreadUpToDate() bool {
	return this.ReadBinlogCoordinates.Equals(&this.ExecBinlogCoordinates)
}


func (this *Instance) AddSlaveHost(hostname string) {
	this.SlaveHosts[hostname] = hostname
}

func (this *Instance) GetMasterInstanceKey() *InstanceKey {
	return &InstanceKey{Hostname: this.Master_Host, Port: this.Master_Port}
}

func (this *Instance) GetSlaveInstanceKeys() [](*InstanceKey) {
	res := [](*InstanceKey){}
	for host, _ := range this.SlaveHosts {
    	res = append(res, &InstanceKey{Hostname: host, Port: this.Key.Port})
	}
	return res
}

func (this *Instance) IsSlaveOf(master *Instance) bool {
	return this.GetMasterInstanceKey().Equals(&master.Key)
}

func (this *Instance) IsMasterOf(slave *Instance) bool {
	return slave.IsSlaveOf(this)
}

