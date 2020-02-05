package agent

import (
	"bytes"
	"encoding/json"
	"time"
)

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

type SeedMethodOpts struct {
	BackupSide       SeedSide
	SupportedEngines []Engine
	BackupToDatadir  bool
}

type BackupMetadata struct {
	LogFile      string
	LogPos       int64
	GtidExecuted string
}

type SeedStatus int

const (
	Started SeedStatus = iota
	Running
	Completed
	Error
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

var toStatus = map[string]SeedStatus{
	"Started":   Started,
	"Running":   Running,
	"Completed": Completed,
	"Error":     Error,
}

// SeedOperation makes for the high level data & state of a seed operation
type SeedOperation struct {
	SeedId         int64
	TargetHostname string
	SourceHostname string
	StartTimestamp time.Time
	EndTimestamp   time.Time
	Status         SeedStatus
}
