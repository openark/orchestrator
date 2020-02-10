package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/openark/golib/log"
)

var SeededAgents chan *Agent = make(chan *Agent)

type SeedSide int

const (
	Target SeedSide = iota
	Source
)

var toSeedSide = map[string]SeedSide{
	"Target": Target,
	"Source": Source,
}

func (s SeedSide) String() string {
	return [...]string{"Target", "Source"}[s]
}

func (s SeedSide) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *SeedSide) UnmarshalJSON(b []byte) error {
	var value string
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}
	*s = toSeedSide[value]
	return nil
}

type Engine int

const (
	ROCKSDB Engine = iota
	MRG_MYISAM
	CSV
	BLACKHOLE
	InnoDB
	MEMORY
	ARCHIVE
	MyISAM
	FEDERATED
	TokuDB
)

func (e Engine) String() string {
	return [...]string{"ROCKSDB", "MRG_MYISAM", "CSV", "BLACKHOLE", "InnoDB", "MEMORY", "ARCHIVE", "MyISAM", "FEDERATED", "TokuDB"}[e]
}

// MarshalJSON marshals the enum as a quoted json string
func (e Engine) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(e.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

var toEngine = map[string]Engine{
	"ROCKSDB":    ROCKSDB,
	"MRG_MYISAM": MRG_MYISAM,
	"CSV":        CSV,
	"BLACKHOLE":  BLACKHOLE,
	"InnoDB":     InnoDB,
	"MEMORY":     MEMORY,
	"ARCHIVE":    ARCHIVE,
	"MyISAM":     MyISAM,
	"FEDERATED":  FEDERATED,
	"TokuDB":     TokuDB,
}

func (e *Engine) UnmarshalJSON(b []byte) error {
	var value string
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}
	*e = toEngine[value]
	return nil
}

type SeedMethod int

const (
	ClonePlugin SeedMethod = iota
	LVM
	Mydumper
	Mysqldump
	Xtrabackup
)

var toSeedMethod = map[string]SeedMethod{
	"ClonePlugin": ClonePlugin,
	"LVM":         LVM,
	"Mydumper":    Mydumper,
	"Mysqldump":   Mysqldump,
	"Xtrabackup":  Xtrabackup,
}

func (m SeedMethod) String() string {
	return [...]string{"ClonePlugin", "LVM", "Mydumper", "Mysqldump", "Xtrabackup"}[m]
}

func (m *SeedMethod) UnmarshalText(b []byte) error {
	*m = toSeedMethod[string(b)]
	return nil
}

func (m SeedMethod) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

type SeedStatus int

const (
	Started SeedStatus = iota
	Running
	Completed
	Error
	Failed
)

func (s SeedStatus) String() string {
	return [...]string{"Started", "Running", "Completed", "Error"}[s]
}

func (s SeedStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

var toSeedStatus = map[string]SeedStatus{
	"Started":   Started,
	"Running":   Running,
	"Completed": Completed,
	"Error":     Error,
	"Failed":    Failed,
}

type SeedStage int

const (
	Prepare SeedStage = iota
	Backup
	Restore
	GetMetadata
	Cleanup
)

func (s SeedStage) String() string {
	return [...]string{"Prepare", "Backup", "Restore", "GetMetadata", "Cleanup"}[s]
}

func (s SeedStage) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// Seed makes for the high level data & state of a seed operation
type Seed struct {
	SeedId         int64
	TargetHostname string
	SourceHostname string
	SeedMethod     SeedMethod
	BackupSide     SeedSide
	Status         SeedStatus
	Retries        int
	StartTimestamp time.Time
	EndTimestamp   time.Time
}

// SeedStageState represents a single stage in a seed operation
type SeedStageState struct {
	SeedStageId int64
	SeedId      int64
	Stage       SeedStage
	Hostname    string
	StartedAt   time.Time
	Status      SeedStatus
	Details     string
}

// NewSeed is the entry point for making a seed. It's registers seed in db, so it can be later processed asynchronously
func NewSeed(targetHostname string, sourceHostname string, seedMethodName string) (int64, error) {
	var seedMethod SeedMethod
	var ok bool
	if seedMethod, ok = toSeedMethod["seedMethodName"]; !ok {
		return 0, log.Errorf("SeedMethod %s not found", seedMethodName)
	}
	if targetHostname == sourceHostname {
		return 0, log.Errorf("Cannot seed %s onto itself", targetHostname)
	}
	targetAgent, err := ReadAgent(targetHostname)
	if err != nil {
		return 0, log.Errorf("Unable to get information from agent on taget host %s", targetHostname)
	}
	sourceAgent, err := ReadAgent(sourceHostname)
	if err != nil {
		return 0, log.Errorf("Unable to get information from agent on source host %s", sourceHostname)
	}
	if _, ok = targetAgent.Data.AvailiableSeedMethods[seedMethod]; !ok {
		return 0, log.Errorf("Seed method %s not supported on target host %s", seedMethod.String(), targetHostname)
	}
	if _, ok = sourceAgent.Data.AvailiableSeedMethods[seedMethod]; !ok {
		return 0, log.Errorf("Seed method %s not supported on source host %s", seedMethod.String(), sourceHostname)
	}
	for db, opts := range sourceAgent.Data.MySQLDatabases {
		if !enginesSupported(opts.Engines, sourceAgent.Data.AvailiableSeedMethods[seedMethod].SupportedEngines) {
			return 0, log.Errorf("Database %s had a table with engine unsupported by seed method %s", db, seedMethod.String())
		}
	}
	if sourceAgent.Data.MySQLDatadirDiskUsed > targetAgent.Data.MySQLDatadirDiskFree {
		return 0, log.Errorf("Not enough disk space on target host %s. Required: %d, available: %d", targetHostname, sourceAgent.Data.MySQLDatadirDiskUsed, targetAgent.Data.MySQLDatadirDiskFree)
	}
	// do we need to check this? logical backup usually less than actual data size, so let's use 0.6 multiplier
	if !sourceAgent.Data.AvailiableSeedMethods[seedMethod].BackupToDatadir {
		if int64(float64(sourceAgent.Data.MySQLDatadirDiskUsed)*0.6) > targetAgent.Data.BackupDirDiskFree {
			return 0, log.Errorf("Not enough disk space on target host backup directory %s. Database size: %d, available: %d", targetHostname, sourceAgent.Data.MySQLDatadirDiskUsed, targetAgent.Data.BackupDirDiskFree)
		}
	}
	seed := &Seed{
		TargetHostname: targetHostname,
		SourceHostname: sourceHostname,
		SeedMethod:     seedMethod,
		BackupSide:     sourceAgent.Data.AvailiableSeedMethods[seedMethod].BackupSide,
		Status:         Started,
		Retries:        0,
		StartTimestamp: time.Now(),
	}
	registeredSeed, err := registerSeedEntry(seed)
	if err != nil {
		return 0, log.Errore(err)
	}
	SeededAgents <- targetAgent

	log.Infof("Registered new seed with seedID %d. Target host: %s, source host: %s, seed method: %s", registeredSeed.SeedId, registeredSeed.TargetHostname, registeredSeed.SourceHostname, registeredSeed.SeedMethod.String())
	auditAgentOperation("agent-seed", sourceAgent, fmt.Sprintf("starting seed to host %s using seed method %s", registeredSeed.TargetHostname, registeredSeed.SeedMethod.String()))
	auditAgentOperation("agent-seed", targetAgent, fmt.Sprintf("starting seed from host %s using seed method %s", registeredSeed.SourceHostname, registeredSeed.SeedMethod.String()))
	return registeredSeed.SeedId, nil
}

func enginesSupported(first, second []Engine) bool {
	hashmap := make(map[Engine]int)
	for _, value := range first {
		hashmap[value]++
	}

	for _, value := range second {
		if count, found := hashmap[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			hashmap[value] = count - 1
		}
	}
	return true
}

func (s *Seed) ProcessStarted(wg *sync.WaitGroup) {
	defer wg.Done()

}

func (s *Seed) ProcessRunning(wg *sync.WaitGroup) {
	defer wg.Done()

}

func (s *Seed) ProcessErrored(wg *sync.WaitGroup) {
	defer wg.Done()

}
