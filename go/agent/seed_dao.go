package agent

import (
	"fmt"
	"time"

	"github.com/github/orchestrator/go/db"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

// registerSeedEntry registers a new seed operation entry, returning seed with it's unique id
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

// updateSeedData updates an information about seed
func (seed *Seed) updateSeedData() error {
	_, err := db.ExecOrchestrator(`
			update
				agent_seed
					set
					stage = ?,
					status = ?,
					retries = ?,
					updated_at = NOW(),
					end_timestamp = ?
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

// isSeedStageCompletedForAgent checks if current seed stage is already completed for current agent
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
				AND agent_hostname = ?
				AND status = ?
				AND agent_seed_state_id > (
					SELECT 
						MAX(agent_seed_state_id)
					FROM
						agent_seed_state
					WHERE
					agent_seed_id = ?
					AND stage = ?
					AND agent_hostname = ?
					AND status = ?
				)
			`
	if err := db.QueryOrchestrator(query, sqlutils.Args(seed.SeedID, seed.Stage.String(), agent.Info.Hostname, Completed.String(), seed.SeedID, seed.Stage.String(), agent.Info.Hostname, Scheduled.String()), func(m sqlutils.RowMap) error {
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
					agent_seed_id, stage, agent_hostname, state_timestamp, status, details
				) VALUES (
					?, ?, ?, ?, ?, ?
				)
			`,
		seedStageState.SeedID,
		seedStageState.Stage.String(),
		seedStageState.Hostname,
		time.Now(),
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
		seed.Status = ToSeedStatus[m.GetString("status")]
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

// ReadSeedStageStates reads states for a given seed operation
func (seed *Seed) ReadSeedStageStates() ([]SeedStageState, error) {
	res := []SeedStageState{}
	query := `
		select
			agent_seed_state_id,
			agent_seed_id,
			stage,
			agent_hostname,
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
		seedState.Hostname = m.GetString("agent_hostname")
		seedState.Timestamp = m.GetTime("state_timestamp")
		seedState.Status = ToSeedStatus[m.GetString("status")]
		seedState.Details = m.GetString("details")

		res = append(res, seedState)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}
