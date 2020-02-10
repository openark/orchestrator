package agent

import (
	"fmt"

	"github.com/github/orchestrator/go/config"
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
func registerSeedEntry(seed *Seed) (*Seed, error) {
	res, err := db.ExecOrchestrator(`
			insert
				into agent_seed (
					seed_method, target_hostname, source_hostname, start_timestamp, status, backup_side, retries
				) VALUES (
					?, ?, ?, ?, ?, ?, ?
				)
			`,
		seed.SeedMethod.String(),
		seed.TargetHostname,
		seed.SourceHostname,
		seed.StartTimestamp,
		seed.Status.String(),
		seed.BackupSide.String(),
		seed.Retries,
	)
	if err != nil {
		return seed, log.Errore(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return seed, err
	}
	seed.SeedId = id
	return seed, nil
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
			status,
			retries,
			start_timestamp,
			end_timestamp
		from
			agent_seed
		%s
		order by
			agent_seed_id desc
		%s
		`, whereCondition, limit)
	err := db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		seed := Seed{}
		seed.SeedId = m.GetInt64("agent_seed_id")
		seed.TargetHostname = m.GetString("target_hostname")
		seed.SourceHostname = m.GetString("source_hostname")
		seed.SeedMethod = toSeedMethod[m.GetString("seed_method")]
		seed.BackupSide = toSeedSide[m.GetString("backup_side")]
		seed.Status = toSeedStatus[m.GetString("status")]
		seed.Retries = m.GetInt("retries")
		seed.StartTimestamp = m.GetTime("start_timestamp")
		seed.EndTimestamp = m.GetTime("end_timestamp")

		res = append(res, &seed)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadActiveSeeds reads active seeds (where status != Completed or Failed)
func ReadActiveSeeds([]Seed, error) ([]*Seed, error) {
	whereCondition := fmt.Sprintf(`
		where
			status not in (%s, %s)
		`, Completed.String(), Failed.String())
	return readSeeds(whereCondition, sqlutils.Args(), "")
}

/***************************************************************************************************************************************/
/***************************************************************************************************************************************/
/************************************************************** OLD FUNCS **************************************************************/
/***************************************************************************************************************************************/
/***************************************************************************************************************************************/

// ReadRecentSeeds reads seeds from backend table.
func ReadRecentSeeds() ([]*Seed, error) {
	return readSeeds(``, sqlutils.Args(), "limit 100")
}

// AgentSeedDetails reads details from backend table
func AgentSeedDetails(seedId int64) ([]*Seed, error) {
	whereCondition := `
		where
			agent_seed_id = ?
		`
	return readSeeds(whereCondition, sqlutils.Args(seedId), "")
}

// ReadRecentCompletedSeedsForHost reads active seeds where host participates either as source or target
func ReadRecentCompletedSeedsForHost(hostname string) ([]*Seed, error) {
	whereCondition := `
		where
			status = "Completed"
			and (
				target_hostname = ?
				or source_hostname = ?
			)
		`
	return readSeeds(whereCondition, sqlutils.Args(hostname, hostname), "limit 10")
}

// updateSeedComplete updates the seed entry, signing for completion
func updateSeedComplete(seedId int64, seedError error) error {
	_, err := db.ExecOrchestrator(`
			update
				agent_seed
					set end_timestamp = NOW(),
					is_complete = 1,
					is_successful = ?
				where
					agent_seed_id = ?
			`,
		(seedError == nil),
		seedId,
	)
	if err != nil {
		return log.Errore(err)
	}

	return nil
}

// SeedOperationState reads states for a given seed operation
func ReadSeedStates(seedId int64) ([]SeedStageState, error) {
	res := []SeedStageState{}
	query := `
		select
			agent_seed_state_id,
			agent_seed_id,
			state_timestamp,
			state_action,
			error_message
		from
			agent_seed_state
		where
			agent_seed_id = ?
		order by
			agent_seed_state_id desc
		`
	err := db.QueryOrchestrator(query, sqlutils.Args(seedId), func(m sqlutils.RowMap) error {
		seedState := SeedStageState{}
		seedState.SeedStageId = m.GetInt64("agent_seed_state_id")
		seedState.SeedId = m.GetInt64("agent_seed_id")
		seedState.StartedAt = m.GetTime("state_timestamp")
		//seedState.Action = m.GetString("state_action")
		//seedState.ErrorMessage = m.GetString("error_message")

		res = append(res, seedState)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// submitSeedStateEntry submits a seed state: a single step in the overall seed process
func submitSeedStateEntry(seedId int64, action string, errorMessage string) (int64, error) {
	res, err := db.ExecOrchestrator(`
			insert
				into agent_seed_state (
					agent_seed_id, state_timestamp, state_action, error_message
				) VALUES (
					?, NOW(), ?, ?
				)
			`,
		seedId,
		action,
		errorMessage,
	)
	if err != nil {
		return 0, log.Errore(err)
	}
	id, err := res.LastInsertId()

	return id, err
}

// updateSeedStateEntry updates seed step state
func updateSeedStateEntry(seedStateId int64, reason error) error {
	_, err := db.ExecOrchestrator(`
			update
				agent_seed_state
					set error_message = ?
				where
					agent_seed_state_id = ?
			`,
		reason.Error(),
		seedStateId,
	)
	if err != nil {
		return log.Errore(err)
	}

	return reason
}

// FailStaleSeeds marks as failed seeds where no progress have been seen recently
func FailStaleSeeds() error {
	_, err := db.ExecOrchestrator(`
				update
						agent_seed
					set
						is_complete=1,
						is_successful=0
					where
						is_complete=0
						and (
							select
									max(state_timestamp) as last_state_timestamp
								from
									agent_seed_state
								where
									agent_seed.agent_seed_id = agent_seed_state.agent_seed_id
						) < now() - interval ? minute`,
		config.Config.StaleSeedFailMinutes,
	)
	return err
}

// ReadActiveSeedsForHost reads active seeds where host participates either as source or target
func ReadActiveSeedsForHost(hostname string) ([]*Seed, error) {
	whereCondition := `
		where
			is_complete = 0
			and (
				target_hostname = ?
				or source_hostname = ?
			)
		`
	return readSeeds(whereCondition, sqlutils.Args(hostname, hostname), "")
}

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
