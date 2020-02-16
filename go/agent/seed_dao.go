package agent

import (
	"fmt"

	"github.com/github/orchestrator/go/db"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

/***************************************************************************************************************************************/
/***************************************************************************************************************************************/
/************************************************************** NEW FUNCS **************************************************************/
/***************************************************************************************************************************************/
/***************************************************************************************************************************************/

/** TO DO **/
/** add ReadFailedSeedsForHost **/

// registerSeedEntry register a new seed operation entry, returning seed with it's unique id
func (seed *Seed) registerSeedEntry() error {
	res, err := db.ExecOrchestrator(`
			insert
				into agent_seed (
					seed_method, target_hostname, source_hostname, start_timestamp, updated_at, status, backup_side, retries, stage
				) VALUES (
					?, ?, ?, ?, ?, ?, ?, 0, ?
				)
			`,
		seed.SeedMethod.String(),
		seed.TargetHostname,
		seed.SourceHostname,
		seed.StartTimestamp,
		seed.UpdatedAt,
		seed.Status.String(),
		seed.BackupSide.String(),
		seed.Stage.String(),
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	seed.SeedID = id
	return nil
}

func (seed *Seed) updateSeedData() error {
	_, err := db.ExecOrchestrator(`
			update
				agent_seed
					set
					stage = ?,
					status = ?,
					retries = ?
					updated_at = NOW(),
					EndTimestamp = ?
				where
					agent_seed_id = ?
			`,
		seed.Stage.String(),
		seed.Status.String(),
		seed.Retries,
		seed.EndTimestamp,
		seed.SeedID,
	)
	if err != nil {
		return log.Errore(fmt.Errorf("Unable to update seed: %+v", err))
	}
	return nil
}

// setSeedStatusFailed sets status Failed for seed and increase retries counter
func (seed *Seed) isSeedStageCompletedForAgent(agent *Agent) (bool, error) {
	var cnt int
	query := `
			SELECT
				count(*) as cnt
			FROM
				agent_seed_state
			WHERE
				agent_seed_id = ?
				AND stage = ?
				AND hostname = ?
				AND status = ?
			`
	if err := db.QueryOrchestrator(query, sqlutils.Args(seed.SeedID, seed.Stage, agent.Info.Hostname, Completed), func(m sqlutils.RowMap) error {
		cnt = m.GetInt("cnt")
		return nil
	}); err != nil {
		return false, log.Errore(fmt.Errorf("Unable to get information about completed stages for seed: %+v", err))
	}
	if cnt > 0 {
		return true, nil
	}
	return false, nil
}

// submitSeedStageState submits a seed stage state: a single step in the overall seed process
func submitSeedStageState(seedStageState *SeedStageState) error {
	_, err := db.ExecOrchestrator(`
			insert
				into agent_seed_state (
					agent_seed_id, stage, hostname, state_timestamp, status, details
				) VALUES (
					?, ?, ?, ?, ?, ?
				)
			`,
		seedStageState.SeedID,
		seedStageState.Stage.String(),
		seedStageState.Hostname,
		seedStageState.Timestamp,
		seedStageState.Status.String(),
		seedStageState.Details,
	)
	if err != nil {
		return log.Errore(fmt.Errorf("Unable to submit seed stage to database: %+v", err))
	}
	return nil
}

// readSeeds reads seed from the backend table
func readSeeds(whereCondition string, args []interface{}, limit string) ([]*Seed, error) {
	res := []*Seed{}
	query := fmt.Sprintf(`
		select
			agent_seed_id,
			target_hostname,
			source_hostname,
			seed_method,
			backup_side,
			stage,
			status,
			retries,
			start_timestamp,
			end_timestamp,
			updated_at
		from
			agent_seed
		%s
		order by
			agent_seed_id desc
		%s
		`, whereCondition, limit)
	err := db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		seed := Seed{}
		seed.SeedID = m.GetInt64("agent_seed_id")
		seed.TargetHostname = m.GetString("target_hostname")
		seed.SourceHostname = m.GetString("source_hostname")
		seed.SeedMethod = toSeedMethod[m.GetString("seed_method")]
		seed.BackupSide = toSeedSide[m.GetString("backup_side")]
		seed.Stage = toSeedStage[m.GetString("stage")]
		seed.Status = toSeedStatus[m.GetString("status")]
		seed.Retries = m.GetInt("retries")
		seed.StartTimestamp = m.GetTime("start_timestamp")
		seed.EndTimestamp = m.GetTime("end_timestamp")
		seed.UpdatedAt = m.GetTime("updated_at")

		res = append(res, &seed)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadSeedStates reads states for a given seed operation
func (seed *Seed) ReadSeedStates() ([]SeedStageState, error) {
	res := []SeedStageState{}
	query := `
		select
			agent_seed_state_id,
			agent_seed_id,
			stage,
			hostname,
			state_timestamp,
			status,
			details
		from
			agent_seed_state
		where
			agent_seed_id = ?
		order by
			agent_seed_state_id desc
		`
	err := db.QueryOrchestrator(query, sqlutils.Args(seed.SeedID), func(m sqlutils.RowMap) error {
		seedState := SeedStageState{}
		seedState.SeedStageID = m.GetInt64("agent_seed_state_id")
		seedState.SeedID = m.GetInt64("agent_seed_id")
		seedState.Stage = toSeedStage[m.GetString("stage")]
		seedState.Timestamp = m.GetTime("state_timestamp")
		seedState.Status = toSeedStatus[m.GetString("status")]
		seedState.Details = m.GetString("details")

		res = append(res, seedState)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

/***************************************************************************************************************************************/
/***************************************************************************************************************************************/
/************************************************************** OLD FUNCS **************************************************************/
/***************************************************************************************************************************************/
/***************************************************************************************************************************************/

/*

// executeSeed is *the* function for taking a seed. It is a complex operation of testing, preparing, re-testing
// agents on both sides, initiating data transfer, following up, awaiting completion, diagnosing errors, claning up.
func executeSeed(seedId int64, seedMethod SeedMethod, targetHostname string, sourceHostname string) error {

	var err error
	var seedStateId int64

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("getting target agent info for %s", targetHostname), "")
	targetAgent, err := GetAgent(targetHostname)
	SeededAgents <- &targetAgent
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("getting source agent info for %s", sourceHostname), "")
	sourceAgent, err := GetAgent(sourceHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Checking MySQL status on target %s", targetHostname), "")
	if targetAgent.Info.MySQLRunning {
		return updateSeedStateEntry(seedStateId, errors.New("MySQL is running on target host. Cowardly refusing to proceeed. Please stop the MySQL service"))
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Looking up available snapshots on source %s", sourceHostname), "")
	if len(sourceAgent.Info.LogicalVolumes) == 0 {
		return updateSeedStateEntry(seedStateId, errors.New("No logical volumes found on source host"))
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Checking mount point on source %s", sourceHostname), "")
	if sourceAgent.Info.MountPoint.IsMounted {
		return updateSeedStateEntry(seedStateId, errors.New("Volume already mounted on source host; please unmount"))
	}

	seedFromLogicalVolume := sourceAgent.Info.LogicalVolumes[0]
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s Mounting logical volume: %s", sourceHostname, seedFromLogicalVolume.Path), "")
	_, err = MountLV(sourceHostname, seedFromLogicalVolume.Path)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}
	sourceAgent, err = GetAgent(sourceHostname)
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("MySQL data volume on source host %s is %d bytes", sourceHostname, sourceAgent.Info.MySQLDatadirDiskUsed), "")

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Erasing MySQL data on %s", targetHostname), "")
	_, err = deleteMySQLDatadir(targetHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Aquiring target host datadir free space on %s", targetHostname), "")
	targetAgent, err = GetAgent(targetHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	if sourceAgent.Info.MySQLDatadirDiskUsed > targetAgent.Info.MySQLDatadirDiskFree {
		Unmount(sourceHostname)
		return updateSeedStateEntry(seedStateId, fmt.Errorf("Not enough disk space on target host %s. Required: %d, available: %d. Bailing out.", targetHostname, sourceAgent.Info.MySQLDatadirDiskUsed, targetAgent.Info.MySQLDatadirDiskFree))
	}

	// ...
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s will now receive data in background", targetHostname), "")
	ReceiveMySQLSeedData(targetHostname, seedId)

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Waiting %d seconds for %s to start listening for incoming data", config.Config.SeedWaitSecondsBeforeSend, targetHostname), "")
	time.Sleep(time.Duration(config.Config.SeedWaitSecondsBeforeSend) * time.Second)

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s will now send data to %s in background", sourceHostname, targetHostname), "")
	SendMySQLSeedData(sourceHostname, targetHostname, seedId)

	copyComplete := false
	numStaleIterations := 0
	var bytesCopied int64 = 0

	for !copyComplete {
		targetAgentPoll, err := GetAgent(targetHostname)
		if err != nil {
			return log.Errore(err)
		}

		if targetAgentPoll.Info.MySQLDatadirDiskUsed == bytesCopied {
			numStaleIterations++
		}
		bytesCopied = targetAgentPoll.Info.MySQLDatadirDiskUsed

		copyFailed := false
		if _, commandCompleted, _ := seedCommandCompleted(targetHostname, seedId); commandCompleted {
			copyComplete = true
			if _, commandSucceeded, _ := seedCommandSucceeded(targetHostname, seedId); !commandSucceeded {
				// failed.
				copyFailed = true
			}
		}
		if numStaleIterations > 10 {
			copyFailed = true
		}
		if copyFailed {
			AbortSeedCommand(sourceHostname, seedId)
			AbortSeedCommand(targetHostname, seedId)
			Unmount(sourceHostname)
			return updateSeedStateEntry(seedStateId, errors.New("10 iterations have passed without progress. Bailing out."))
		}

		var copyPct int64 = 0
		if sourceAgent.Info.MySQLDatadirDiskUsed > 0 {
			copyPct = 100 * bytesCopied / sourceAgent.Info.MySQLDatadirDiskUsed
		}
		seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Copied %d/%d bytes (%d%%)", bytesCopied, sourceAgent.Info.MySQLDatadirDiskUsed, copyPct), "")

		if !copyComplete {
			time.Sleep(30 * time.Second)
		}
	}

	// Cleanup:
	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Executing post-copy command on %s", targetHostname), "")
	_, err = PostCopy(targetHostname, sourceHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("%s Unmounting logical volume: %s", sourceHostname, seedFromLogicalVolume.Path), "")
	_, err = Unmount(sourceHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Starting MySQL on target: %s", targetHostname), "")
	_, err = MySQLStart(targetHostname)
	if err != nil {
		return updateSeedStateEntry(seedStateId, err)
	}

	seedStateId, _ = submitSeedStateEntry(seedId, fmt.Sprintf("Submitting MySQL instance for discovery: %s", targetHostname), "")
	SeededAgents <- &targetAgent

	seedStateId, _ = submitSeedStateEntry(seedId, "Done", "")

	return nil
}


*/
