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
	return [...]string{"Started", "Running", "Completed", "Error", "Failed"}[s]
}

func (s SeedStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *SeedStatus) UnmarshalJSON(b []byte) error {
	var value string
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}
	*s = toSeedStatus[value]
	return nil
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
	Cleanup
	ConnectSlave
)

func (s SeedStage) String() string {
	return [...]string{"Prepare", "Backup", "Restore", "Cleanup", "ConnectSlave"}[s]
}

func (s SeedStage) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *SeedStage) UnmarshalJSON(b []byte) error {
	var value string
	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}
	*s = toSeedStage[value]
	return nil
}

var toSeedStage = map[string]SeedStage{
	"Prepare":      Prepare,
	"Backup":       Backup,
	"Restore":      Restore,
	"Cleanup":      Cleanup,
	"ConnectSlave": ConnectSlave,
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

func auditSeedOperation(agent *Agent, message string) error {
	instanceKey := &inst.InstanceKey{}
	if agent != nil {
		instanceKey = &inst.InstanceKey{Hostname: agent.Info.Hostname, Port: agent.Info.MySQLPort}
	}
	return inst.AuditOperation("seed", instanceKey, message)
}

// NewSeed is the entry point for making a seed. It's registers seed in db, so it can be later processed asynchronously
func NewSeed(seedMethodName string, targetAgent *Agent, sourceAgent *Agent) (int64, error) {
	var seedMethod SeedMethod
	var ok bool
	if seedMethod, ok = toSeedMethod[seedMethodName]; !ok {
		return 0, log.Errorf("SeedMethod %s not found", seedMethodName)
	}
	if err := targetAgent.getAgentData(); err != nil {
		return 0, log.Errorf("Unable to get information from agent on taget host %s", targetAgent.Info.Hostname)
	}
	if err := sourceAgent.getAgentData(); err != nil {
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
		if !enginesSupported(sourceAgent.Data.AvailiableSeedMethods[seedMethod].SupportedEngines, opts.Engines) {
			return 0, log.Errorf("Database %s had a table with engine unsupported by seed method %s", db, seedMethod.String())
		}
	}
	if inst.IsSmallerMajorVersion(targetAgent.Data.MySQLVersion, sourceAgent.Data.MySQLVersion) {
		return 0, log.Errorf("MySQL version on target agent is smaller that on source. Source agent MySQL version: %s, target agent MySQL version: %s", sourceAgent.Data.MySQLVersion, targetAgent.Data.MySQLVersion)
	}
	activeSourceSeeds, _ := ReadActiveSeedsForAgent(sourceAgent)
	activeTargetSeeds, _ := ReadActiveSeedsForAgent(targetAgent)
	if len(activeSourceSeeds) > 0 {
		return 0, log.Errorf("Source agent has active seeds")
	}
	if len(activeTargetSeeds) > 0 {
		return 0, log.Errorf("Target agent has active seeds")
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
		Stage:          Prepare,
		Retries:        0,
		StartTimestamp: time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := seed.registerSeedEntry(); err != nil {
		return 0, log.Errore(err)
	}
	SeededAgents <- targetAgent

	auditSeedOperation(nil, fmt.Sprintf("Registered new seed. SeedID: %d, source host: %s, target host: %s, seed method: %s", seed.SeedID, seed.SourceHostname, seed.TargetHostname, seed.SeedMethod.String()))
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
	whereCondition := `
		where
			status not in (?, ?)
		`
	return readSeeds(whereCondition, sqlutils.Args(Completed.String(), Failed.String()), "")
}

// ReadRecentSeeds reads recent seeds from backend table.
func ReadRecentSeeds() ([]*Seed, error) {
	return readSeeds(``, sqlutils.Args(), "limit 100")
}

// ReadSeed returns an information about seed from database
func ReadSeed(seedID int64) (*Seed, error) {
	whereCondition := `
		where
			agent_seed_id = ?
		`
	res, err := readSeeds(whereCondition, sqlutils.Args(seedID), "")
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("Seed %d not found", seedID)
	}
	return res[0], nil
}

// ReadRecentSeedsForAgentInStatus reads active seeds where agent participates either as source or target in given status
func ReadRecentSeedsForAgentInStatus(agent *Agent, status SeedStatus, limit string) ([]*Seed, error) {
	whereCondition := `
		where
			status = ?
			and (
				target_hostname = ?
				or source_hostname = ?
			)
		`
	return readSeeds(whereCondition, sqlutils.Args(status.String(), agent.Info.Hostname, agent.Info.Hostname), "limit 10")
}

// ReadActiveSeedsForAgent reads active seeds where agent participates either as source or target
func ReadActiveSeedsForAgent(agent *Agent) ([]*Seed, error) {
	whereCondition := `
		where
			status not in (?, ?)
			and (
				target_hostname = ?
				or source_hostname = ?
			)
		`
	return readSeeds(whereCondition, sqlutils.Args(Completed.String(), Failed.String(), agent.Info.Hostname, agent.Info.Hostname), "")
}

func ProcessSeeds() []*Seed {
	activeSeeds, err := ReadActiveSeeds()
	if err != nil {
		log.Errore(fmt.Errorf("Unable to read active seeds: %+v", err))
		return nil
	}
	inst.AuditOperation("process-seeds", nil, fmt.Sprintf("Will process %d active seeds", len(activeSeeds)))
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
	inst.AuditOperation("process-seeds", nil, "All active seeds processed")
	return activeSeeds
}

func (s *Seed) GetSeedAgents() (targetAgent *Agent, sourceAgent *Agent, err error) {
	targetAgent, err = ReadAgentInfo(s.TargetHostname)
	if err != nil {
		s.Status = Failed
		s.Retries++
		s.updateSeed(targetAgent, fmt.Sprintf("Unable to read target agent info: %+v", err))
		return nil, nil, log.Errore(err)
	}
	sourceAgent, err = ReadAgentInfo(s.SourceHostname)
	if err != nil {
		s.Status = Failed
		s.Retries++
		s.updateSeed(sourceAgent, fmt.Sprintf("Unable to read source agent info: %+v", err))
		return nil, nil, log.Errore(err)
	}
	return targetAgent, sourceAgent, nil
}

func (s *Seed) processStarted(wg *sync.WaitGroup) {
	defer wg.Done()
	targetAgent, sourceAgent, err := s.GetSeedAgents()
	if err != nil {
		return
	}
	switch s.Stage {
	case Prepare:
		agents := make(map[*Agent]SeedSide)
		agents[targetAgent] = Target
		agents[sourceAgent] = Source
		for agent, seedSide := range agents {
			if err := agent.prepare(s.SeedID, s.SeedMethod, seedSide); err != nil {
				s.Status = Failed
				s.Retries++
				s.updateSeed(agent, fmt.Sprintf("Error calling prepare API on agent: %+v", err))
				return
			}
			s.Status = Running
			s.updateSeed(agent, "Started prepare stage")
		}
	case Backup:
		agent := targetAgent
		if s.BackupSide == Source {
			agent = sourceAgent
		}
		if err := agent.backup(s.SeedID, s.SeedMethod, sourceAgent.Info.Hostname, sourceAgent.Info.MySQLPort); err != nil {
			s.Status = Failed
			s.Retries++
			s.updateSeed(agent, fmt.Sprintf("Error calling backup API on agent: %+v", err))
			return
		}
		s.Status = Running
		s.updateSeed(agent, "Started backup stage")
	case Restore:
		if err := targetAgent.restore(s.SeedID, s.SeedMethod); err != nil {
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error calling restore API on agent: %+v", err))
			return
		}
		s.Status = Running
		s.updateSeed(targetAgent, "Started restore stage")
	case Cleanup:
		agents := make(map[*Agent]SeedSide)
		agents[targetAgent] = Target
		agents[sourceAgent] = Source
		for agent, seedSide := range agents {
			if err := agent.cleanup(s.SeedID, s.SeedMethod, seedSide); err != nil {
				s.Status = Failed
				s.Retries++
				s.updateSeed(targetAgent, fmt.Sprintf("Error calling prepare API on agent: %+v", err))
				return
			}
			s.Status = Running
			s.updateSeed(agent, "Started prepare stage")
		}
	// also will connect targetHost as slave
	case ConnectSlave:
		var gtidHint inst.OperationGTIDHint = inst.GTIDHintNeutral
		s.updateSeedState(targetAgent.Info.Hostname, "Started connect slave stage")
		seedMetadata, err := targetAgent.getMetadata(s.SeedID, s.SeedMethod)
		if err != nil {
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error calling getMetadata API on agent: %+v", err))
			return
		}
		slaveInstanceKey := &inst.InstanceKey{Hostname: targetAgent.Info.Hostname, Port: targetAgent.Info.MySQLPort}
		masterInstanceKey := &inst.InstanceKey{Hostname: sourceAgent.Info.Hostname, Port: sourceAgent.Info.MySQLPort}
		slave, err := inst.ReadTopologyInstance(slaveInstanceKey)
		if err != nil {
			log.Errore(err)
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error reading topology instance: %+v", err))
			return
		}
		master, err := inst.ReadTopologyInstance(masterInstanceKey)
		if err != nil {
			log.Errore(err)
			s.Status = Failed
			s.Retries++
			s.updateSeed(sourceAgent, fmt.Sprintf("Error reading topology instance: %+v", err))
			return
		}
		if s.Retries > 0 && slave.IsReplicaOf(master) && slave.Slave_SQL_Running {
			//everything is already good, seems like we do not mark seed as completed so just exit switch
			break
		}
		slave, err = inst.ResetSlave(&slave.Key)
		if err != nil {
			log.Errore(err)
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing RESET SLAVE: %+v", err))
			return
		}
		slave, err = inst.ResetMaster(&slave.Key)
		if err != nil {
			log.Errore(err)
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing RESET MASTER: %+v", err))
			return
		}
		if len(seedMetadata.GtidExecuted) > 0 {
			if err = inst.SetGTIDPurged(slave, seedMetadata.GtidExecuted); err != nil {
				log.Errore(err)
				s.Status = Failed
				s.Retries++
				s.updateSeed(targetAgent, fmt.Sprintf("Error setting GTID_PURGED: %+v", err))
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
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing CHANGE MASTER TO: %+v", err))
			return
		}
		replicationUser, replicationPassword, err := inst.ReadReplicationCredentials(&slave.Key)
		if err != nil {
			log.Errore(err)
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error getting replication credentials: %+v", err))
			return
		}
		slave, err = inst.ChangeMasterCredentials(&slave.Key, replicationUser, replicationPassword)
		if err != nil {
			log.Errore(err)
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing CHANGE MASTER TO with replication credentials: %+v", err))
			return
		}
		if slave.AllowTLS {
			if _, err := inst.EnableMasterSSL(&slave.Key); err != nil {
				log.Errore(err)
				s.Status = Failed
				s.Retries++
				s.updateSeed(targetAgent, fmt.Sprintf("Error executing CHANGE MASTER TO enable SSL: %+v", err))
				return
			}
		}
		slave, err = inst.StartSlave(&slave.Key)
		if err != nil {
			log.Errore(err)
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing START SLAVE: %+v", err))
			return
		}
		s.Status = Completed
		s.EndTimestamp = time.Now()
		s.updateSeed(targetAgent, fmt.Sprintf("Seed completed"))
	}
}

func (s *Seed) processRunning(wg *sync.WaitGroup) {
	defer wg.Done()
	targetAgent, sourceAgent, err := s.GetSeedAgents()
	if err != nil {
		return
	}
	switch s.Stage {
	case Prepare, Cleanup:
		agents := make(map[*Agent]bool)
		agents[targetAgent] = false
		agents[sourceAgent] = false
		for agent := range agents {
			stageAlreadyCompleted, err := s.isSeedStageCompletedForAgent(agent)
			if err != nil {
				s.Status = Failed
				s.Retries++
				s.updateSeed(agent, fmt.Sprintf("Error getting information about completed stages from Orchestrator: %+v", err))
				return
			}
			if !stageAlreadyCompleted {
				agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
				if err != nil {
					s.Status = Failed
					s.Retries++
					s.updateSeed(targetAgent, fmt.Sprintf("Error getting seed stage state information from agent: %+v", err))
					return
				}
				submitSeedStageState(agentSeedStageState)
				if agentSeedStageState.Status == Completed {
					agents[agent] = true
					auditSeedOperation(agent, fmt.Sprintf("SeedID: %d, Stage: %s, Status: %s, Retries: %d", s.SeedID, s.Stage.String(), Completed.String(), s.Retries))
				}
				if agentSeedStageState.Status == Error {
					s.Status = Error
					s.Retries++
					s.updateSeed(agent, "")
					return
				}
			} else {
				agents[agent] = true
			}
		}
		if agents[targetAgent] && agents[sourceAgent] {
			s.Stage = s.Stage + 1
			s.Status = Started
			s.Retries = 0
			s.updateSeed(nil, "")
		}
	case Backup, Restore:
		agent := targetAgent
		if s.Stage == Backup && s.BackupSide == Source {
			agent = sourceAgent
		}
		agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
		if err != nil {
			s.Status = Failed
			s.Retries++
			s.updateSeed(targetAgent, fmt.Sprintf("Error getting seed stage state information from agent: %+v", err))
			return
		}
		submitSeedStageState(agentSeedStageState)
		if agentSeedStageState.Status == Completed {
			auditSeedOperation(agent, fmt.Sprintf("SeedID: %d, Stage: %s, Status: %s, Retries: %d", s.SeedID, s.Stage.String(), Completed.String(), s.Retries))
			s.Stage = s.Stage + 1
			s.Status = Started
			s.Retries = 0
			s.updateSeed(agent, "")
		}
		if agentSeedStageState.Status == Error {
			s.Status = Error
			s.Retries++
			s.updateSeed(agent, "")
		}
	}
}

func (s *Seed) processErrored(wg *sync.WaitGroup) {
	defer wg.Done()
	if s.Retries >= config.Config.MaxRetriesForSeedStage {
		s.Status = Failed
		s.EndTimestamp = time.Now()
		s.updateSeed(nil, fmt.Sprintf("Seed failed. Seed stage retries: %d, MaxRetriesForSeedStage: %d", s.Retries, config.Config.MaxRetriesForSeedStage))
		return
	}
	// prepare and cleanup stages are executed on both agents, so we need to check them and if it's running - stop it.
	if s.Stage == Prepare || s.Stage == Cleanup {
		targetAgent, sourceAgent, err := s.GetSeedAgents()
		if err != nil {
			return
		}
		agents := []*Agent{}
		agents = append(agents, targetAgent, sourceAgent)

		for _, agent := range agents {
			agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
			if err != nil {
				s.Retries++
				s.updateSeed(agent, fmt.Sprintf("Error getting seed stage state information from agent: %+v", err))
				return
			}
			if agentSeedStageState.Status == Running {
				if err := agent.AbortSeed(s.SeedID); err != nil {
					s.Retries++
					s.updateSeed(agent, fmt.Sprintf("Error aborting seed on agent: %+v", err))
					return
				}
				s.updateSeedState(agent.Info.Hostname, fmt.Sprintf("Aborted %s seed stage due to error on another agent", s.Stage.String()))
			}
		}
	}
	// if stage is backup we need to restart from prepare stage (because we can have operations like stating netcat\socat on target in prepare stage)
	s.Status = Started
	if s.Stage == Backup {
		s.Stage = Prepare
	}
	s.updateSeed(nil, fmt.Sprintf("Seed will be restarted from % stage", s.Stage.String()))
}

func (s *Seed) updateSeed(agent *Agent, seedStateDetails string) {
	if err := s.updateSeedData(); err == nil {
		auditSeedOperation(agent, fmt.Sprintf("SeedID: %d, Stage: %s, Status: %s, Retries: %d", s.SeedID, s.Stage.String(), s.Status.String(), s.Retries))
	}
	if seedStateDetails != "" {
		s.updateSeedState(agent.Info.Hostname, seedStateDetails)
	}
}

func (s *Seed) updateSeedState(hostname string, details string) error {
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
