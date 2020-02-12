package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/inst"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
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
	SeedID         int64
	TargetHostname string
	SourceHostname string
	SeedMethod     SeedMethod
	BackupSide     SeedSide
	Status         SeedStatus
	Stage          SeedStage
	Retries        int
	StartTimestamp time.Time
	EndTimestamp   time.Time
	UpdatedAt      time.Time
}

// SeedStageState represents a single stage in a seed operation
type SeedStageState struct {
	SeedStageID int64
	SeedID      int64
	Stage       SeedStage
	Hostname    string
	Timestamp   time.Time
	Status      SeedStatus
	Details     string
}

// SeedStageState represents a metadata information for seed with binlog coordinates or GtidExecuted
type SeedMetadata struct {
	LogFile      string
	LogPos       int64
	GtidExecuted string
}

// NewSeed is the entry point for making a seed. It's registers seed in db, so it can be later processed asynchronously
// CHECK THAT AGENT DOESN'T HAVE ALREADY ACTIVE SEEDS
func NewSeed(seedMethodName string, targetAgent *Agent, sourceAgent *Agent) (int64, error) {
	var seedMethod SeedMethod
	var ok bool
	if seedMethod, ok = toSeedMethod["seedMethodName"]; !ok {
		return 0, log.Errorf("SeedMethod %s not found", seedMethodName)
	}
	if err := targetAgent.GetAgentData(); err != nil {
		return 0, log.Errorf("Unable to get information from agent on taget host %s", targetAgent.Info.Hostname)
	}
	if err := sourceAgent.GetAgentData(); err != nil {
		return 0, log.Errorf("Unable to get information from agent on source host %s", sourceAgent.Info.Hostname)
	}
	if targetAgent.Info.Hostname == sourceAgent.Info.Hostname {
		return 0, log.Errorf("Cannot seed %s onto itself", targetAgent.Info.Hostname)
	}
	if _, ok = targetAgent.Data.AvailiableSeedMethods[seedMethod]; !ok {
		return 0, log.Errorf("Seed method %s not supported on target host %s", seedMethod.String(), targetAgent.Info.Hostname)
	}
	if _, ok = sourceAgent.Data.AvailiableSeedMethods[seedMethod]; !ok {
		return 0, log.Errorf("Seed method %s not supported on source host %s", seedMethod.String(), sourceAgent.Info.Hostname)
	}
	for db, opts := range sourceAgent.Data.MySQLDatabases {
		if !enginesSupported(opts.Engines, sourceAgent.Data.AvailiableSeedMethods[seedMethod].SupportedEngines) {
			return 0, log.Errorf("Database %s had a table with engine unsupported by seed method %s", db, seedMethod.String())
		}
	}
	if sourceAgent.Data.MySQLDatadirDiskUsed > targetAgent.Data.MySQLDatadirDiskFree {
		return 0, log.Errorf("Not enough disk space on target host %s. Required: %d, available: %d", targetAgent.Info.Hostname, sourceAgent.Data.MySQLDatadirDiskUsed, targetAgent.Data.MySQLDatadirDiskFree)
	}
	// do we need to check this? logical backup usually less than actual data size, so let's use 0.6 multiplier
	if !sourceAgent.Data.AvailiableSeedMethods[seedMethod].BackupToDatadir {
		if int64(float64(sourceAgent.Data.MySQLDatadirDiskUsed)*0.6) > targetAgent.Data.BackupDirDiskFree {
			return 0, log.Errorf("Not enough disk space on target host backup directory %s. Database size: %d, available: %d", targetAgent.Info.Hostname, sourceAgent.Data.MySQLDatadirDiskUsed, targetAgent.Data.BackupDirDiskFree)
		}
	}
	seed := &Seed{
		TargetHostname: targetAgent.Info.Hostname,
		SourceHostname: sourceAgent.Info.Hostname,
		SeedMethod:     seedMethod,
		BackupSide:     sourceAgent.Data.AvailiableSeedMethods[seedMethod].BackupSide,
		Status:         Started,
		Retries:        0,
		StartTimestamp: time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := seed.registerSeedEntry(); err != nil {
		return 0, log.Errore(err)
	}
	SeededAgents <- targetAgent

	log.Infof("Registered new seed with seedID %d. Target host: %s, source host: %s, seed method: %s", seed.SeedID, seed.TargetHostname, seed.SourceHostname, seed.SeedMethod.String())
	auditAgentOperation("agent-seed", sourceAgent, fmt.Sprintf("starting seed to host %s using seed method %s", seed.TargetHostname, seed.SeedMethod.String()))
	auditAgentOperation("agent-seed", targetAgent, fmt.Sprintf("starting seed from host %s using seed method %s", seed.SourceHostname, seed.SeedMethod.String()))
	return seed.SeedID, nil
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

// ReadActiveSeeds reads active seeds (where status != Completed or Failed)
func ReadActiveSeeds() ([]*Seed, error) {
	whereCondition := fmt.Sprintf(`
		where
			status not in (%s, %s)
		`, Completed.String(), Failed.String())
	return readSeeds(whereCondition, sqlutils.Args(), "")
}

func ProcessSeeds() {
	log.Info("Processing active seeds")
	activeSeeds, err := ReadActiveSeeds()
	if err != nil {
		log.Errore(fmt.Errorf("Unable to read active seeds: %+v", err))
		return
	}
	var wg sync.WaitGroup
	for _, seed := range activeSeeds {
		wg.Add(1)
		switch seed.Status {
		case Started:
			go seed.processStarted(&wg)
		case Running:
			go seed.processRunning(&wg)
		case Error:
			go seed.processErrored(&wg)
		}
	}
	wg.Wait()
	log.Info("All active seeds processed")
	return
}

func (s *Seed) getSeedAgents() (targetAgent *Agent, sourceAgent *Agent, err error) {
	targetAgent, err = ReadAgentInfo(s.TargetHostname)
	if err != nil {
		if err = s.setSeedStatusFailed(); err != nil {
			log.Errore(err)
		}
		return nil, nil, err
	}
	sourceAgent, err = ReadAgentInfo(s.SourceHostname)
	if err != nil {
		if err = s.setSeedStatusFailed(); err != nil {
			log.Errore(err)
		}
		return nil, nil, err
	}
	return targetAgent, sourceAgent, nil
}

// add additional log and debug
func (s *Seed) processStarted(wg *sync.WaitGroup) {
	defer wg.Done()
	targetAgent, sourceAgent, err := s.getSeedAgents()
	if err != nil {
		log.Errore(err)
		return
	}
	switch s.Stage {
	case Prepare:
		if err := targetAgent.prepare(s.SeedID, s.SeedMethod, Target); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if err := s.updateSeedState(Target, "Started prepare stage"); err != nil {
			log.Errore(err)
		}
		if err := sourceAgent.prepare(s.SeedID, s.SeedMethod, Source); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if err := s.updateSeedState(Source, "Started prepare stage"); err != nil {
			log.Errore(err)
		}
	case Backup:
		agent := targetAgent
		seedSide := Target
		if s.BackupSide == Source {
			agent = sourceAgent
			seedSide = Source
		}
		if err := agent.backup(s.SeedID, s.SeedMethod, sourceAgent.Info.Hostname, sourceAgent.Info.Port); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if err := s.updateSeedState(seedSide, "Started backup stage"); err != nil {
			log.Errore(err)
		}
	case Restore:
		if err := targetAgent.restore(s.SeedID, s.SeedMethod); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if err := s.updateSeedState(Target, "Started restore stage"); err != nil {
			log.Errore(err)
		}
	case Cleanup:
		if err := targetAgent.cleanup(s.SeedID, s.SeedMethod, Target); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if err := s.updateSeedState(Target, "Started cleanup stage"); err != nil {
			log.Errore(err)
		}
		if err := sourceAgent.cleanup(s.SeedID, s.SeedMethod, Source); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if err := s.updateSeedState(Source, "Started cleanup stage"); err != nil {
			log.Errore(err)
		}
	// also will connect targetHost as slave
	case GetMetadata:
		var gtidHint inst.OperationGTIDHint = inst.GTIDHintNeutral
		seedMetadata, err := targetAgent.getMetdata(s.SeedID, s.SeedMethod)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		slaveInstanceKey := &inst.InstanceKey{Hostname: targetAgent.Info.Hostname, Port: targetAgent.Info.Port}
		masterInstanceKey := &inst.InstanceKey{Hostname: sourceAgent.Info.Hostname, Port: sourceAgent.Info.Port}
		slave, err := inst.ReadTopologyInstance(slaveInstanceKey)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		master, err := inst.ReadTopologyInstance(masterInstanceKey)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if s.Retries > 0 && slave.IsReplicaOf(master) && slave.Slave_SQL_Running {
			//everything is already good, seems like we do not mark seed as completed so just exit switch
			break
		}
		slave, err = inst.ResetSlave(&slave.Key)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		slave, err = inst.ResetMaster(&slave.Key)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if len(seedMetadata.GtidExecuted) > 0 {
			if err = inst.SetGTIDPurged(slave, seedMetadata.GtidExecuted); err != nil {
				log.Errore(err)
				if err = s.setSeedStatusFailed(); err != nil {
					log.Errore(err)
				}
				return
			}
			gtidHint = inst.GTIDHintForce
		}
		binlogCoordinates := &inst.BinlogCoordinates{
			LogFile: seedMetadata.LogFile,
			LogPos:  seedMetadata.LogPos,
			Type:    inst.BinaryLog,
		}
		slave, err = inst.ChangeMasterTo(&slave.Key, &master.Key, binlogCoordinates, false, gtidHint)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		replicationUser, replicationPassword, err := inst.ReadReplicationCredentials(&slave.Key)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		slave, err = inst.ChangeMasterCredentials(&slave.Key, replicationUser, replicationPassword)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if slave.AllowTLS {
			if _, err := inst.EnableMasterSSL(&slave.Key); err != nil {
				log.Errore(err)
				if err = s.setSeedStatusFailed(); err != nil {
					log.Errore(err)
				}
				return
			}
		}
		slave, err = inst.StartSlave(&slave.Key)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
	}
	if s.Stage != GetMetadata {
		if err := s.updateSeedStageAndStatus(s.Stage, Running); err != nil {
			log.Errore(err)
			return
		}
	} else {
		if err := s.finishSeed(); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
	}
	return
}

func (s *Seed) processRunning(wg *sync.WaitGroup) {
	defer wg.Done()
	targetAgent, sourceAgent, err := s.getSeedAgents()
	if err != nil {
		log.Errore(err)
		return
	}
	switch s.Stage {
	case Prepare, Cleanup:
		agents := make(map[*Agent]bool)
		agents[targetAgent] = false
		agents[sourceAgent] = false
		for agent := range agents {
			agentSeedStageState, err := agent.SeedStageState(s.SeedID, s.Stage)
			if err != nil {
				log.Errore(err)
				if err = s.setSeedStatusFailed(); err != nil {
					log.Errore(err)
				}
				return
			}
			stageAlreadyCompleted, err := s.isSeedStageCompletedForAgent(agent)
			if err != nil {
				log.Errore(err)
				if err = s.setSeedStatusFailed(); err != nil {
					log.Errore(err)
				}
				return
			}
			if !stageAlreadyCompleted {
				if err := submitSeedStageState(agentSeedStageState); err != nil {
					log.Errore(err)
					if err = s.setSeedStatusFailed(); err != nil {
						log.Errore(err)
					}
					return
				}
				if agentSeedStageState.Status == Completed {
					agents[agent] = true
				}
			} else {
				agents[agent] = true
			}
		}
		if agents[targetAgent] && agents[sourceAgent] {
			if s.Stage == Prepare {
				if err := s.startNewStage(Backup); err != nil {
					log.Errore(err)
					return
				}
			}
			if s.Stage == Cleanup {
				if err := s.startNewStage(GetMetadata); err != nil {
					log.Errore(err)
					return
				}
			}
		}
	case Backup:
		agent := targetAgent
		if s.BackupSide == Source {
			agent = sourceAgent
		}
		agentSeedStageState, err := agent.SeedStageState(s.SeedID, s.Stage)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if err := submitSeedStageState(agentSeedStageState); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if agentSeedStageState.Status == Completed {
			if err := s.startNewStage(Restore); err != nil {
				log.Errore(err)
				return
			}
		}
	case Restore:
		agentSeedStageState, err := targetAgent.SeedStageState(s.SeedID, s.Stage)
		if err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if err := submitSeedStageState(agentSeedStageState); err != nil {
			log.Errore(err)
			if err = s.setSeedStatusFailed(); err != nil {
				log.Errore(err)
			}
			return
		}
		if agentSeedStageState.Status == Completed {
			if err := s.startNewStage(Cleanup); err != nil {
				log.Errore(err)
				return
			}
		}
	}
	return
}

func (s *Seed) processErrored(wg *sync.WaitGroup) {
	defer wg.Done()
	if s.Retries >= config.Config.MaxRetriesForSeedStage {
		if err := s.failSeed(); err != nil {
			log.Errore(err)
		}
		return
	}
	// prepare and cleanup stages are executed on both agents, so we need to check them and if it's running - stop it.
	if s.Stage == Prepare || s.Stage == Cleanup {
		targetAgent, sourceAgent, err := s.getSeedAgents()
		if err != nil {
			log.Errore(err)
			return
		}
		agents := make(map[*Agent]SeedSide)
		agents[targetAgent] = Target
		agents[sourceAgent] = Source
		for agent, seedSide := range agents {
			agentSeedStageState, err := agent.SeedStageState(s.SeedID, s.Stage)
			if err != nil {
				log.Errore(err)
				s.setSeedStatusFailed()
				return
			}
			if agentSeedStageState.Status == Running {
				if err := agent.AbortSeed(s.SeedID); err != nil {
					log.Errore(err)
					s.setSeedStatusFailed()
					return
				}
				if err := s.updateSeedState(seedSide, "aborted seed stage due to error on another agent"); err != nil {
					log.Errore(err)
				}
				return
			}
		}
	}
	// if stage is backup we need to restart from prepare stage (because we can have operations like stating netcat\socat on target in prepare stage)
	if s.Stage == Backup {
		if err := s.updateSeedStageAndStatus(Prepare, Started); err != nil {
			log.Errore(err)
			return
		}
	} else {
		if err := s.updateSeedStageAndStatus(s.Stage, Started); err != nil {
			log.Errore(err)
			return
		}
	}
	return
}

func (s *Seed) updateSeedState(seedSide SeedSide, details string) error {
	var hostname string
	switch seedSide {
	case Target:
		hostname = s.TargetHostname
	case Source:
		hostname = s.SourceHostname
	}
	seedStageState := &SeedStageState{
		SeedID:    s.SeedID,
		Stage:     s.Stage,
		Hostname:  hostname,
		Timestamp: time.Now(),
		Status:    s.Status,
		Details:   details,
	}
	return submitSeedStageState(seedStageState)
}
