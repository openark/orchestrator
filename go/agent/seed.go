package agent

import (
	"bytes"
	"encoding/json"
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
	DatabaseSelection bool
	BackupSide        SeedSide
}

type BackupMetadata struct {
	LogFile      string
	LogPos       int64
	GtidExecuted string
}
