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

type seedProgress struct {
	bytesCopied   int64
	databasesSize int64
	updatedAt     time.Time
}

var seedOperationProgress = make(map[int64]*seedProgress)

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
	Scheduled SeedStatus = iota
	Running
	Completed
	Error
	Failed
	Aborting
)

var SeedStatuses = []SeedStatus{Scheduled, Running, Completed, Error, Failed, Aborting}

func (s SeedStatus) String() string {
	return [...]string{"Scheduled", "Running", "Completed", "Error", "Failed", "Aborting"}[s]
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
	*s = ToSeedStatus[value]
	return nil
}

var ToSeedStatus = map[string]SeedStatus{
	"Scheduled": Scheduled,
	"Running":   Running,
	"Completed": Completed,
	"Error":     Error,
	"Failed":    Failed,
	"Aborting":  Aborting,
}

type SeedStage int

const (
	Prepare SeedStage = iota
	Backup
	Restore
	ConnectSlave
	Cleanup
)

func (s SeedStage) String() string {
	return [...]string{"Prepare", "Backup", "Restore", "ConnectSlave", "Cleanup"}[s]
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
	"ConnectSlave": ConnectSlave,
	"Cleanup":      Cleanup,
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

func byteCount(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
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
		return 0, log.Errorf("Unable to get information from agent on target host %s", targetAgent.Info.Hostname)
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
	targetInstance, err := inst.ReadTopologyInstance(&inst.InstanceKey{Hostname: targetAgent.Info.Hostname, Port: targetAgent.Info.MySQLPort})
	if err != nil {
		return 0, log.Errorf("Unable to read target agent instance %s", targetAgent.Info.Hostname)
	}
	sourceInstance, err := inst.ReadTopologyInstance(&inst.InstanceKey{Hostname: sourceAgent.Info.Hostname, Port: sourceAgent.Info.MySQLPort})
	if err != nil {
		return 0, log.Errorf("Unable to read source agent instance %s", targetAgent.Info.Hostname)
	}
	if targetInstance.IsSmallerMajorVersion(sourceInstance) {
		return 0, log.Errorf("MySQL version on target agent is smaller that on source. Source agent MySQL version: %s, target agent MySQL version: %s", sourceInstance.Version, targetInstance.Version)
	}
	if !sourceInstance.LogBinEnabled {
		return 0, log.Errorf("Binlog is not enabled on MySQL source agent host %s", sourceAgent.Info.Hostname)
	}
	if sourceInstance.IsReplica() {
		if !sourceInstance.LogSlaveUpdatesEnabled {
			return 0, log.Errorf("log-slave-updates is not enabled on MySQL source agent host %s, but the host is a replica", sourceAgent.Info.Hostname)
		}
	}
	if targetInstance.IsReplica() {
		return 0, log.Errorf("Cannot seed on target agent host %s, because it is replica. Please reset slave before starting seed", targetAgent.Info.Hostname)
	}
	if len(targetInstance.SlaveHosts) > 0 {
		return 0, log.Errorf("Cannot seed on target agent host %s, because it has slave attached. Please disconnect all slaves before seed", targetAgent.Info.Hostname)
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
	// special checks for LVM seed method
	if seedMethod == LVM {
		if sourceAgent.Data.MountPoint.IsMounted {
			return 0, log.Errorf("Volume already mounted on source host %s. Please unmount", sourceAgent.Info.Hostname)
		}
		if len(sourceAgent.Data.LogicalVolumes) == 0 {
			return 0, log.Errorf("No logical volumes found on source host %s", sourceAgent.Info.Hostname)
		}
	}
	seed := &Seed{
		TargetHostname: targetAgent.Info.Hostname,
		SourceHostname: sourceAgent.Info.Hostname,
		SeedMethod:     seedMethod,
		BackupSide:     sourceAgent.Data.AvailiableSeedMethods[seedMethod].BackupSide,
		Status:         Scheduled,
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

// ReadActiveSeeds reads active seeds (where status != Completed or Failed)
func ReadActiveSeeds() ([]*Seed, error) {
	whereCondition := `
		where
			status not in (?, ?)
		`
	return readSeeds(whereCondition, sqlutils.Args(Completed.String(), Failed.String()), "")
}

// ReadSeeds reads seeds from backend table.
func ReadSeeds() ([]*Seed, error) {
	return readSeeds(``, sqlutils.Args(), "")
}

// ReadSeedsPaged returns a list of seeds with pagination
func ReadSeedsPaged(page int) ([]*Seed, error) {
	limit := fmt.Sprintf("limit %d offset %d", config.AuditPageSize, page*config.AuditPageSize)
	return readSeeds(``, sqlutils.Args(), limit)
}

// ReadSeedsForTargetAgentPaged returns a list of seeds for specified targetagent with pagination
func ReadSeedsForTargetAgentPaged(targetAgent *Agent, page int) ([]*Seed, error) {
	whereCondition := `
		where
			target_hostname = ?	
	`
	limit := fmt.Sprintf("limit %d offset %d", config.AuditPageSize, page*config.AuditPageSize)
	return readSeeds(whereCondition, sqlutils.Args(targetAgent.Info.Hostname), limit)
}

// ReadSeedsForSourceAgentPaged returns a list of seeds for specified sourceagent with pagination
func ReadSeedsForSourceAgentPaged(sourceAgent *Agent, page int) ([]*Seed, error) {
	whereCondition := `
		where
			source_hostname = ?	
	`
	limit := fmt.Sprintf("limit %d offset %d", config.AuditPageSize, page*config.AuditPageSize)
	return readSeeds(whereCondition, sqlutils.Args(sourceAgent.Info.Hostname), limit)
}

// ReadSeedsInStatusPaged returns seeds in specified status with pagination
func ReadSeedsInStatusPaged(status SeedStatus, page int) ([]*Seed, error) {
	limit := fmt.Sprintf("limit %d offset %d", config.AuditPageSize, page*config.AuditPageSize)
	whereCondition := `
		where
			status = ?
	`
	return readSeeds(whereCondition, sqlutils.Args(status.String()), limit)
}

// ReadSeedsInStatus returns seeds in specified status
func ReadSeedsInStatus(status SeedStatus) ([]*Seed, error) {
	whereCondition := `
		where
			status = ?
	`
	return readSeeds(whereCondition, sqlutils.Args(status.String()), ``)
}

// ReadSeedsForAgent reads seeds where agent participates either as source or target in given status
func ReadSeedsForAgent(agent *Agent, limit string) ([]*Seed, error) {
	whereCondition := `
		where
			target_hostname = ?
			or source_hostname = ?
			
	`
	return readSeeds(whereCondition, sqlutils.Args(agent.Info.Hostname, agent.Info.Hostname), limit)
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

// ProcessSeeds is the main function in seed process
func ProcessSeeds() []*Seed {
	activeSeeds, err := ReadActiveSeeds()
	if err != nil {
		log.Errore(fmt.Errorf("Unable to read active seeds: %+v", err))
		return nil
	}
	if len(activeSeeds) > 0 {
		inst.AuditOperation("process-seeds", nil, fmt.Sprintf("Will process %d active seeds", len(activeSeeds)))
		var wg sync.WaitGroup
		for _, seed := range activeSeeds {
			wg.Add(1)
			switch seed.Status {
			case Scheduled:
				go seed.processScheduled(&wg)
			case Running:
				go seed.processRunning(&wg)
			case Error:
				go seed.processErrored(&wg)
			case Aborting:
				go seed.processAborting(&wg)
			}
		}
		wg.Wait()
		inst.AuditOperation("process-seeds", nil, "All active seeds processed")
	}
	return activeSeeds
}

// readSeadAgents returns target and source agents associated with seed from db. If updateAgentData is specified, it will use agent API in order to update agent data
func (s *Seed) readSeadAgents(updateAgentsData bool) (targetAgent *Agent, sourceAgent *Agent, err error) {
	targetAgent, err = ReadAgentInfo(s.TargetHostname)
	if err != nil {
		s.Status = Error
		s.updateSeed(targetAgent, fmt.Sprintf("Unable to read target agent info: %+v", err))
		return nil, nil, log.Errore(err)
	}
	if updateAgentsData {
		if err = targetAgent.getAgentData(); err != nil {
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Unable to update agent data: %+v", err))
			return nil, nil, log.Errore(err)
		}
	}
	sourceAgent, err = ReadAgentInfo(s.SourceHostname)
	if err != nil {
		s.Status = Error
		s.updateSeed(sourceAgent, fmt.Sprintf("Unable to read source agent info: %+v", err))
		return nil, nil, log.Errore(err)
	}
	if updateAgentsData {
		if err = sourceAgent.getAgentData(); err != nil {
			s.Status = Error
			s.updateSeed(sourceAgent, fmt.Sprintf("Unable to update agent data: %+v", err))
			return nil, nil, log.Errore(err)
		}
	}
	return targetAgent, sourceAgent, nil
}

func (s *Seed) processAborting(wg *sync.WaitGroup) {
	defer wg.Done()
	s.AbortSeed()
}

func (s *Seed) processScheduled(wg *sync.WaitGroup) {
	defer wg.Done()
	switch s.Stage {
	case Prepare:
		targetAgent, sourceAgent, err := s.readSeadAgents(false)
		if err != nil {
			return
		}
		targetInstanceKey := &inst.InstanceKey{Hostname: targetAgent.Info.Hostname, Port: targetAgent.Info.MySQLPort}
		downtimeDuration := time.Duration(config.SeedDowntimeHours) * time.Hour
		downtime := inst.NewDowntime(targetInstanceKey, "orchestrator-agent", "seed-in-process", downtimeDuration)
		if err = inst.BeginDowntime(downtime); err != nil {
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error starting downtime on targetAgent: %+v", err))
			return
		}
		agents := make(map[*Agent]SeedSide)
		agents[targetAgent] = Target
		agents[sourceAgent] = Source
		for agent, seedSide := range agents {
			if err := agent.prepare(s.SeedID, s.SeedMethod, seedSide); err != nil {
				s.Status = Error
				s.updateSeed(agent, fmt.Sprintf("Error calling prepare API on agent: %+v", err))
				return
			}
			s.Status = Running
			s.updateSeed(agent, "Stage started")
		}
	case Backup:
		targetAgent, sourceAgent, err := s.readSeadAgents(false)
		if err != nil {
			return
		}
		agent := targetAgent
		seedHost := sourceAgent.Info.Hostname
		if s.BackupSide == Source {
			agent = sourceAgent
			seedHost = targetAgent.Info.Hostname
		}
		if err := agent.backup(s.SeedID, s.SeedMethod, seedHost, sourceAgent.Info.MySQLPort); err != nil {
			s.Status = Error
			s.updateSeed(agent, fmt.Sprintf("Error calling backup API on agent: %+v", err))
			return
		}
		s.Status = Running
		s.updateSeed(agent, "Stage started")
		if _, ok := seedOperationProgress[s.SeedID]; ok {
			delete(seedOperationProgress, s.SeedID)
		}
	case Restore:
		targetAgent, _, err := s.readSeadAgents(false)
		if err != nil {
			return
		}
		if err := targetAgent.restore(s.SeedID, s.SeedMethod); err != nil {
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error calling restore API on agent: %+v", err))
			return
		}
		s.Status = Running
		s.updateSeed(targetAgent, "Stage started")
		if _, ok := seedOperationProgress[s.SeedID]; ok {
			delete(seedOperationProgress, s.SeedID)
		}
	case Cleanup:
		targetAgent, sourceAgent, err := s.readSeadAgents(false)
		if err != nil {
			return
		}
		agents := make(map[*Agent]SeedSide)
		agents[targetAgent] = Target
		agents[sourceAgent] = Source
		for agent, seedSide := range agents {
			if err := agent.cleanup(s.SeedID, s.SeedMethod, seedSide); err != nil {
				s.Status = Error
				s.updateSeed(targetAgent, fmt.Sprintf("Error calling prepare API on agent: %+v", err))
				return
			}
			s.Status = Running
			s.updateSeed(agent, "Stage started")
		}
	case ConnectSlave:
		targetAgent, _, err := s.readSeadAgents(false)
		if err != nil {
			return
		}
		s.Status = Running
		s.updateSeed(targetAgent, "Stage started")
	}
}

func (s *Seed) processRunning(wg *sync.WaitGroup) {
	defer wg.Done()
	switch s.Stage {
	case Prepare, Cleanup:
		targetAgent, sourceAgent, err := s.readSeadAgents(false)
		if err != nil {
			return
		}
		agents := make(map[*Agent]bool)
		agents[targetAgent] = false
		agents[sourceAgent] = false
		for agent := range agents {
			stageAlreadyCompleted, err := s.isSeedStageCompletedForAgent(agent)
			if err != nil {
				s.Status = Error
				s.updateSeed(agent, fmt.Sprintf("Error getting information about completed stages from Orchestrator: %+v", err))
				return
			}
			if !stageAlreadyCompleted {
				agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
				if err != nil {
					s.Status = Error
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
					s.updateSeed(agent, "")
					return
				}
			} else {
				agents[agent] = true
			}
		}
		if agents[targetAgent] && agents[sourceAgent] {
			if s.Stage == Prepare {
				s.Stage = s.Stage + 1
				s.Status = Scheduled
				s.updateSeed(nil, "")
			} else {
				for _, agent := range []*Agent{targetAgent, sourceAgent} {
					if err := agent.postSeedCmd(s.SeedID); err != nil {
						s.Status = Error
						s.updateSeedState(agent.Info.Hostname, fmt.Sprintf("Failed to execute post seed command on agent: %+v", err))
					} else {
						s.Status = Completed
						s.updateSeedState(agent.Info.Hostname, "Executed post-seed command")
					}
				}
				targetInstanceKey := &inst.InstanceKey{Hostname: targetAgent.Info.Hostname, Port: targetAgent.Info.MySQLPort}
				if _, err = inst.EndDowntime(targetInstanceKey); err != nil {
					s.Status = Error
					s.updateSeed(targetAgent, fmt.Sprintf("Error ending downtime on targetAgent: %+v", err))
					return
				}
				s.Status = Completed
				s.EndTimestamp = time.Now()
				s.updateSeed(nil, "")
				for _, agent := range []*Agent{targetAgent, sourceAgent} {
					s.updateSeedState(agent.Info.Hostname, fmt.Sprintf("Seed completed"))
				}
			}
		}
	case Backup:
		targetAgent, sourceAgent, err := s.readSeadAgents(true)
		if err != nil {
			return
		}
		agent := targetAgent
		if s.BackupSide == Source {
			agent = sourceAgent
		}
		agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
		if err != nil {
			s.Status = Error
			s.updateSeed(agent, fmt.Sprintf("Error getting seed stage state information from agent: %+v", err))
			return
		}
		// Create struct to store seed progress in case if they not exists
		if _, ok := seedOperationProgress[s.SeedID]; !ok {
			var databasesSize int64
			for _, dbProps := range sourceAgent.Data.MySQLDatabases {
				databasesSize += dbProps.Size
			}
			seedOperationProgress[s.SeedID] = &seedProgress{
				bytesCopied:   0,
				databasesSize: databasesSize,
				updatedAt:     time.Now(),
			}
		}
		if agentSeedStageState.Status == Running {
			// check that seed is not stale
			bytesCopied := targetAgent.Data.BackupDirDiskUsed
			if agent.Data.AvailiableSeedMethods[s.SeedMethod].BackupToDatadir == true {
				bytesCopied = targetAgent.Data.MySQLDatadirDiskUsed
			}
			// if bytesCopied increased - just update info for this SeedID
			if bytesCopied > seedOperationProgress[s.SeedID].bytesCopied {
				seedOperationProgress[s.SeedID].bytesCopied = bytesCopied
				seedOperationProgress[s.SeedID].updatedAt = agentSeedStageState.Timestamp
			} else {
				// else check diff between agentSeedStageState.Timestamp and seedOperationProgress[s.SeedID].updatedAt. If it is more than config.Config.StaleSeedFailMinutes - abort seed and mark it as failed
				if agentSeedStageState.Timestamp.Sub(seedOperationProgress[s.SeedID].updatedAt).Minutes() >= float64(config.Config.SeedBackupStaleFailMinutes) {
					if err := agent.AbortSeed(s.SeedID, s.Stage); err != nil {
						s.Retries++
						s.updateSeed(agent, fmt.Sprintf("Error aborting stale seed on agent: %+v", err))
						return
					}
					s.Status = Error
					s.updateSeed(agent, fmt.Sprintf("No data copied in SeedBackupStaleFailMinutes interval (%d minutes). Aborting seed", config.Config.SeedBackupStaleFailMinutes))
					return
				}
			}
			agentSeedStageState.Details = fmt.Sprintf("Copied: %s. MySQL databases size: %s", byteCount(bytesCopied), byteCount(seedOperationProgress[s.SeedID].databasesSize))
			agentSeedStageState.Timestamp = time.Now()
			submitSeedStageState(agentSeedStageState)
		}
		if agentSeedStageState.Status == Completed {
			submitSeedStageState(agentSeedStageState)
			auditSeedOperation(agent, fmt.Sprintf("SeedID: %d, Stage: %s, Status: %s, Retries: %d", s.SeedID, s.Stage.String(), Completed.String(), s.Retries))
			s.Stage = s.Stage + 1
			s.Status = Scheduled
			s.Retries = 0
			s.updateSeed(agent, "")
			delete(seedOperationProgress, s.SeedID)
		}
		if agentSeedStageState.Status == Error {
			submitSeedStageState(agentSeedStageState)
			s.Status = Error
			s.updateSeed(agent, "")
			delete(seedOperationProgress, s.SeedID)
		}
	case Restore:
		targetAgent, sourceAgent, err := s.readSeadAgents(true)
		if err != nil {
			return
		}
		agentSeedStageState, err := targetAgent.getSeedStageState(s.SeedID, s.Stage)
		if err != nil {
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error getting seed stage state information from agent: %+v", err))
			return
		}
		if agentSeedStageState.Status == Running {
			if targetAgent.Data.AvailiableSeedMethods[s.SeedMethod].BackupToDatadir == false {
				// Create struct to store seed progress in case if they not exists
				if _, ok := seedOperationProgress[s.SeedID]; !ok {
					var databasesSizeSource int64
					for _, dbPropsSource := range sourceAgent.Data.MySQLDatabases {
						databasesSizeSource += dbPropsSource.Size
					}
					seedOperationProgress[s.SeedID] = &seedProgress{
						bytesCopied:   0,
						databasesSize: databasesSizeSource,
						updatedAt:     time.Now(),
					}
				}
				// check that seed is not stale
				var databasesSize int64
				for _, dbProps := range targetAgent.Data.MySQLDatabases {
					databasesSize += dbProps.Size
				}
				// if bytesCopied increased - just update info for this SeedID
				if databasesSize > seedOperationProgress[s.SeedID].bytesCopied {
					seedOperationProgress[s.SeedID].bytesCopied = databasesSize
					seedOperationProgress[s.SeedID].updatedAt = agentSeedStageState.Timestamp
				} else {
					// else check diff between agentSeedStageState.Timestamp and seedBackupProgress[s.SeedID].updatedAt. If it is more than config.Config.StaleSeedFailMinutes - abort seed and mark it as failed
					if agentSeedStageState.Timestamp.Sub(seedOperationProgress[s.SeedID].updatedAt).Minutes() >= float64(config.Config.SeedBackupStaleFailMinutes) {
						if err := targetAgent.AbortSeed(s.SeedID, s.Stage); err != nil {
							s.Retries++
							s.updateSeed(targetAgent, fmt.Sprintf("Error aborting stale seed on agent: %+v", err))
							return
						}
						s.Status = Error
						s.updateSeed(targetAgent, fmt.Sprintf("No data restored in SeedBackupStaleFailMinutes interval (%d minutes). Aborting seed", config.Config.SeedBackupStaleFailMinutes))
						return
					}
				}
				agentSeedStageState.Details = fmt.Sprintf("Restored: %s. MySQL databases size: %s", byteCount(databasesSize), byteCount(seedOperationProgress[s.SeedID].databasesSize))
				agentSeedStageState.Timestamp = time.Now()
			}
			submitSeedStageState(agentSeedStageState)
		}
		if agentSeedStageState.Status == Completed {
			submitSeedStageState(agentSeedStageState)
			auditSeedOperation(targetAgent, fmt.Sprintf("SeedID: %d, Stage: %s, Status: %s, Retries: %d", s.SeedID, s.Stage.String(), Completed.String(), s.Retries))
			s.Stage = s.Stage + 1
			s.Status = Scheduled
			s.Retries = 0
			if targetAgent.Data.AvailiableSeedMethods[s.SeedMethod].BackupToDatadir == false {
				delete(seedOperationProgress, s.SeedID)
			}
			s.updateSeed(targetAgent, "")
		}
		if agentSeedStageState.Status == Error {
			if targetAgent.Data.AvailiableSeedMethods[s.SeedMethod].BackupToDatadir == false {
				delete(seedOperationProgress, s.SeedID)
			}
			submitSeedStageState(agentSeedStageState)
			s.Status = Error
			s.updateSeed(targetAgent, "")
		}
	case ConnectSlave:
		targetAgent, sourceAgent, err := s.readSeadAgents(false)
		if err != nil {
			return
		}
		var gtidHint inst.OperationGTIDHint = inst.GTIDHintNeutral
		seedMetadata, err := targetAgent.getMetadata(s.SeedID, s.SeedMethod)
		if err != nil {
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error calling getMetadata API on agent: %+v", err))
			return
		}
		s.updateSeedState(targetAgent.Info.Hostname, fmt.Sprintf("Got seed metadata from agent. LogFile: %s, LogPos: %d, GtidExecuted: %s", seedMetadata.LogFile, seedMetadata.LogPos, seedMetadata.GtidExecuted))
		log.Debugf("Got seed metadata from agent. LogFile: %s, LogPos: %d, GtidExecuted: %s", seedMetadata.LogFile, seedMetadata.LogPos, seedMetadata.GtidExecuted)
		slaveInstanceKey := &inst.InstanceKey{Hostname: targetAgent.Info.Hostname, Port: targetAgent.Info.MySQLPort}
		masterInstanceKey := &inst.InstanceKey{Hostname: sourceAgent.Info.Hostname, Port: sourceAgent.Info.MySQLPort}
		slave, err := inst.ReadTopologyInstance(slaveInstanceKey)
		if err != nil {
			log.Errore(err)
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error reading topology instance: %+v", err))
			return
		}
		master, err := inst.ReadTopologyInstance(masterInstanceKey)
		if err != nil {
			log.Errore(err)
			s.Status = Error
			s.updateSeed(sourceAgent, fmt.Sprintf("Error reading topology instance: %+v", err))
			return
		}
		if s.Retries > 0 && slave.IsReplicaOf(master) && slave.Slave_SQL_Running && slave.Slave_IO_Running {
			//everything is already good, seems like we do not mark seed stage as completed so just do it
			auditSeedOperation(targetAgent, fmt.Sprintf("SeedID: %d, Stage: %s, Status: %s, Retries: %d", s.SeedID, s.Stage.String(), Completed.String(), s.Retries))
			s.Stage = s.Stage + 1
			s.Status = Scheduled
			s.Retries = 0
			s.updateSeed(targetAgent, "")
		}
		if slave.ReplicationThreadsExist() && !slave.ReplicationThreadsStopped() {
			slave, err = inst.StopSlave(slaveInstanceKey)
			if err != nil {
				log.Errore(err)
				s.Status = Error
				s.updateSeed(targetAgent, fmt.Sprintf("Error stopping slave threads: %+v", err))
				return
			}
		}
		slave, err = inst.ResetSlave(slaveInstanceKey)
		if err != nil {
			log.Errore(err)
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing RESET SLAVE: %+v", err))
			return
		}
		s.updateSeedState(targetAgent.Info.Hostname, "Executed RESET SLAVE")
		slave, err = inst.ResetMaster(slaveInstanceKey)
		if err != nil {
			log.Errore(err)
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing RESET MASTER: %+v", err))
			return
		}
		s.updateSeedState(targetAgent.Info.Hostname, "Executed RESET MASTER")
		if len(seedMetadata.GtidExecuted) > 0 {
			if err = inst.SetGTIDPurged(slave, seedMetadata.GtidExecuted); err != nil {
				log.Errore(err)
				s.Status = Error
				s.updateSeed(targetAgent, fmt.Sprintf("Error setting GTID_PURGED: %+v", err))
				return
			}
		}
		if master.UsingGTID() {
			gtidHint = inst.GTIDHintForce
		}
		binlogCoordinates := &inst.BinlogCoordinates{
			LogFile: seedMetadata.LogFile,
			LogPos:  seedMetadata.LogPos,
			Type:    inst.BinaryLog,
		}
		slave, err = inst.ChangeMasterTo(slaveInstanceKey, masterInstanceKey, binlogCoordinates, false, gtidHint)
		if err != nil {
			log.Errore(err)
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing CHANGE MASTER TO: %+v", err))
			return
		}
		s.updateSeedState(targetAgent.Info.Hostname, "Executed CHANGE MASTER TO binlog position")
		replicationUser, replicationPassword, err := inst.ReadReplicationCredentials(slaveInstanceKey)
		if err != nil {
			log.Errore(err)
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error getting replication credentials: %+v", err))
			return
		}
		slave, err = inst.ChangeMasterCredentials(slaveInstanceKey, replicationUser, replicationPassword)
		if err != nil {
			log.Errore(err)
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing CHANGE MASTER TO with replication credentials: %+v", err))
			return
		}
		s.updateSeedState(targetAgent.Info.Hostname, "Executed CHANGE MASTER TO master credentials")
		if slave.AllowTLS {
			if _, err := inst.EnableMasterSSL(slaveInstanceKey); err != nil {
				log.Errore(err)
				s.Status = Error
				s.updateSeed(targetAgent, fmt.Sprintf("Error executing CHANGE MASTER TO enable SSL: %+v", err))
				return
			}
		}
		slave, err = inst.StartSlave(slaveInstanceKey)
		if err != nil {
			log.Errore(err)
			s.Status = Error
			s.updateSeed(targetAgent, fmt.Sprintf("Error executing START SLAVE: %+v", err))
			return
		}
		s.updateSeedState(targetAgent.Info.Hostname, "Executed START SLAVE")
		s.Status = Completed
		auditSeedOperation(targetAgent, fmt.Sprintf("SeedID: %d, Stage: %s, Status: %s, Retries: %d", s.SeedID, s.Stage.String(), Completed.String(), s.Retries))
		s.updateSeedState(targetAgent.Info.Hostname, "Completed connect slave stage")
		s.Stage = s.Stage + 1
		s.Status = Scheduled
		s.Retries = 0
		s.updateSeed(targetAgent, "")
	}
}

func (s *Seed) processErrored(wg *sync.WaitGroup) {
	defer wg.Done()
	targetAgent, sourceAgent, err := s.readSeadAgents(false)
	if err != nil {
		return
	}
	if s.Retries >= config.Config.MaxRetriesForSeedStage {
		s.Status = Failed
		s.EndTimestamp = time.Now()
		if s.Stage == Prepare || s.Stage == Cleanup {
			for _, agent := range []*Agent{targetAgent, sourceAgent} {
				s.updateSeed(agent, fmt.Sprintf("Seed failed. Seed stage retries: %d, MaxRetriesForSeedStage: %d", s.Retries, config.Config.MaxRetriesForSeedStage))
			}
		}
		if s.Stage == Restore || s.Stage == ConnectSlave {
			s.updateSeed(targetAgent, fmt.Sprintf("Seed failed. Seed stage retries: %d, MaxRetriesForSeedStage: %d", s.Retries, config.Config.MaxRetriesForSeedStage))
		}
		if s.Stage == Backup {
			s.Stage = Prepare
			if s.BackupSide == Target {
				s.updateSeed(targetAgent, fmt.Sprintf("Seed failed. Seed stage retries: %d, MaxRetriesForSeedStage: %d", s.Retries, config.Config.MaxRetriesForSeedStage))
			} else {
				s.updateSeed(sourceAgent, fmt.Sprintf("Seed failed. Seed stage retries: %d, MaxRetriesForSeedStage: %d", s.Retries, config.Config.MaxRetriesForSeedStage))
			}
		}
		return
	}
	// update set SeedStatus Scheduled and update seed according to stage
	s.Status = Scheduled
	s.Retries++
	switch s.Stage {
	case Prepare, Cleanup:
		// prepare and cleanup stages are executed on both agents, so we need to check them and if it's running - stop it.
		for _, agent := range []*Agent{targetAgent, sourceAgent} {
			agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
			if err != nil {
				s.Retries++
				s.updateSeed(agent, fmt.Sprintf("Error getting seed stage state information from agent: %+v", err))
				return
			}
			if agentSeedStageState.Status == Running {
				if err := agent.AbortSeed(s.SeedID, s.Stage); err != nil {
					s.Retries++
					s.updateSeed(agent, fmt.Sprintf("Error aborting seed on agent: %+v", err))
					return
				}
				s.updateSeedState(agent.Info.Hostname, fmt.Sprintf("Aborted %s seed stage due to error on another agent", s.Stage.String()))
			}
		}
		for _, agent := range []*Agent{targetAgent, sourceAgent} {
			s.updateSeed(agent, fmt.Sprintf("Seed will be restarted from %s stage", s.Stage.String()))
		}
	case Restore, ConnectSlave:
		s.updateSeed(targetAgent, fmt.Sprintf("Seed will be restarted from %s stage", s.Stage.String()))
	// if stage is backup we need to restart from prepare stage (because we can have operations like stating netcat\socat on target in prepare stage)
	case Backup:
		s.Stage = Prepare
		if s.BackupSide == Target {
			s.updateSeed(targetAgent, fmt.Sprintf("Seed will be restarted from %s stage", s.Stage.String()))
		} else {
			s.updateSeed(sourceAgent, fmt.Sprintf("Seed will be restarted from %s stage", s.Stage.String()))
		}
	}
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

// AbortSeed aborts seed operation
func (s *Seed) AbortSeed() error {
	previousStatus := s.Status
	if previousStatus != Aborting {
		s.Status = Aborting
		s.updateSeedData()
	}
	targetAgent, sourceAgent, err := s.readSeadAgents(false)
	if err != nil {
		return err
	}
	if previousStatus == Scheduled || previousStatus == Error {
		for _, agent := range []*Agent{targetAgent, sourceAgent} {
			s.updateSeedState(agent.Info.Hostname, fmt.Sprintf("Aborted %s seed stage", s.Stage.String()))
		}
		s.Status = Failed
		s.EndTimestamp = time.Now()
		return s.updateSeedData()
	} else if previousStatus == Failed || previousStatus == Completed {
		return fmt.Errorf("Unable to abort seed in %s status", previousStatus.String())
	}
	switch s.Stage {
	case Prepare:
		for _, agent := range []*Agent{targetAgent, sourceAgent} {
			agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
			if err != nil {
				return err
			}
			if agentSeedStageState.Status == Running || (agent.Info.Hostname == targetAgent.Info.Hostname && s.BackupSide == Source) {
				if err := agent.AbortSeed(s.SeedID, s.Stage); err != nil {
					return err
				}
				s.updateSeedState(agent.Info.Hostname, fmt.Sprintf("Aborted %s seed stage", s.Stage.String()))
			}
		}
	case Cleanup:
		for _, agent := range []*Agent{targetAgent, sourceAgent} {
			agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
			if err != nil {
				return err
			}
			if agentSeedStageState.Status == Running {
				if err := agent.AbortSeed(s.SeedID, s.Stage); err != nil {
					return err
				}
				s.updateSeedState(agent.Info.Hostname, fmt.Sprintf("Aborted %s seed stage", s.Stage.String()))
			}
		}
	case Backup, Restore:
		agent := targetAgent
		if s.BackupSide == Source && s.Stage == Backup {
			agent = sourceAgent
		}
		agentSeedStageState, err := agent.getSeedStageState(s.SeedID, s.Stage)
		if err != nil {
			return err
		}
		if agentSeedStageState.Status == Running {
			if err := agent.AbortSeed(s.SeedID, s.Stage); err != nil {
				return err
			}
			s.updateSeedState(agent.Info.Hostname, fmt.Sprintf("Aborted %s seed stage", s.Stage.String()))
		}
	}
	s.Status = Failed
	s.EndTimestamp = time.Now()
	return s.updateSeedData()
}
