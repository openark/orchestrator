/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package logic

import (
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
	"github.com/outbrain/orchestrator/go/inst"
)

// AttemptFailureDetectionRegistration tries to add a failure-detection entry; if this fails that means the problem has already been detected
func AttemptFailureDetectionRegistration(analysisEntry *inst.ReplicationAnalysis) (bool, error) {

	db, err := db.OpenOrchestrator()
	if err != nil {
		return false, log.Errore(err)
	}

	sqlResult, err := sqlutils.Exec(db, `
			insert ignore 
				into topology_failure_detection (
					hostname, 
					port, 
					in_active_period, 
					start_active_period, 
					end_active_period_unixtime, 
					processing_node_hostname, 
					processcing_node_token,
					analysis,
					cluster_name,
					cluster_alias,
					count_affected_slaves,
					slave_hosts
				) values (
					?,
					?,
					1,
					NOW(),
					0,
					?,
					?,
					?,
					?,
					?,
					?,
					?
				)
			`, analysisEntry.AnalyzedInstanceKey.Hostname, analysisEntry.AnalyzedInstanceKey.Port, ThisHostname, ProcessToken.Hash,
		string(analysisEntry.Analysis), analysisEntry.ClusterDetails.ClusterName, analysisEntry.ClusterDetails.ClusterAlias, analysisEntry.CountSlaves, analysisEntry.GetSlaveHostsAsString(),
	)
	if err != nil {
		return false, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	return (err == nil && rows > 0), err
}

// ClearActiveFailureDetections clears the "in_active_period" flag for old-enough detections, thereby allowing for
// further detections on cleared instances.
func ClearActiveFailureDetections() error {
	_, err := db.ExecOrchestrator(`
			update topology_failure_detection set 
				in_active_period = 0,
				end_active_period_unixtime = UNIX_TIMESTAMP()
			where
				in_active_period = 1
				AND start_active_period < NOW() - INTERVAL ? MINUTE
			`,
		config.Config.FailureDetectionPeriodBlockMinutes,
	)
	return log.Errore(err)
}

// AttemptRecoveryRegistration tries to add a recovery entry; if this fails that means recovery is already in place.
func AttemptRecoveryRegistration(analysisEntry *inst.ReplicationAnalysis) (bool, error) {

	db, err := db.OpenOrchestrator()
	if err != nil {
		return false, log.Errore(err)
	}

	sqlResult, err := sqlutils.Exec(db, `
			insert ignore 
				into topology_recovery (
					hostname, 
					port, 
					in_active_period, 
					start_active_period, 
					end_active_period_unixtime, 
					processing_node_hostname, 
					processcing_node_token,
					analysis,
					cluster_name,
					cluster_alias,
					count_affected_slaves,
					slave_hosts
				) values (
					?,
					?,
					1,
					NOW(),
					0,
					?,
					?,
					?,
					?,
					?,
					?,
					?
				)
			`, analysisEntry.AnalyzedInstanceKey.Hostname, analysisEntry.AnalyzedInstanceKey.Port, ThisHostname, ProcessToken.Hash,
		string(analysisEntry.Analysis), analysisEntry.ClusterDetails.ClusterName, analysisEntry.ClusterDetails.ClusterAlias, analysisEntry.CountSlaves, analysisEntry.GetSlaveHostsAsString(),
	)
	if err != nil {
		return false, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	return (err == nil && rows > 0), err
}

// ClearActiveRecoveries clears the "in_active_period" flag for old-enough recoveries, thereby allowing for
// further recoveries on cleared instances.
func ClearActiveRecoveries() error {
	_, err := db.ExecOrchestrator(`
			update topology_recovery set 
				in_active_period = 0,
				end_active_period_unixtime = UNIX_TIMESTAMP()
			where
				in_active_period = 1
				AND start_active_period < NOW() - INTERVAL ? MINUTE
			`,
		config.Config.FailureDetectionPeriodBlockMinutes,
	)
	return log.Errore(err)
}

// ResolveRecovery is called on completion of a recovery process and updates the recovery status.
// It does not clear the "active period" as this still takes place in order to avoid flapping.
func ResolveRecovery(failedKey *inst.InstanceKey, successorKey *inst.InstanceKey) error {

	if successorKey == nil {
		successorKey = &inst.InstanceKey{}
	}
	_, err := db.ExecOrchestrator(`
			update topology_recovery set 
				successor_hostname = ?,
				successor_port = ?,
				end_recovery = NOW()
			where
				hostname = ?
				AND port = ?
				AND in_active_period = 1
				AND processing_node_hostname = ?
				AND processcing_node_token = ?
			`, successorKey.Hostname, successorKey.Port,
		failedKey.Hostname, failedKey.Port, ThisHostname, ProcessToken.Hash,
	)
	return log.Errore(err)
}

// readRecoveries reads recovery entry/audit entires from topology_recovery
func readRecoveries(whereCondition string, limit string) ([]TopologyRecovery, error) {
	res := []TopologyRecovery{}
	query := fmt.Sprintf(`
		select 
            recovery_id,
            hostname,
            port,
            (IFNULL(end_active_period_unixtime, 0) = 0) as is_active,
            start_active_period,
            IFNULL(end_active_period_unixtime, 0) as end_active_period_unixtime,
            IFNULL(end_recovery, '') AS end_recovery,
            processing_node_hostname,
            processcing_node_token,
            ifnull(successor_hostname, '') as successor_hostname,
            ifnull(successor_port, 0) as successor_port,
            analysis,
            cluster_name,
            cluster_alias,
            count_affected_slaves,
            slave_hosts		
		from 
			topology_recovery
		%s
		order by
			recovery_id desc
		%s
		`, whereCondition, limit)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		topologyRecovery := TopologyRecovery{}
		topologyRecovery.Id = m.GetInt64("recovery_id")

		topologyRecovery.IsActive = m.GetBool("is_active")
		topologyRecovery.RecoveryStartTimestamp = m.GetString("start_active_period")
		topologyRecovery.RecoveryEndTimestamp = m.GetString("end_recovery")
		topologyRecovery.ProcessingNodeHostname = m.GetString("processing_node_hostname")
		topologyRecovery.ProcessingNodeToken = m.GetString("processcing_node_token")

		topologyRecovery.AnalysisEntry.AnalyzedInstanceKey.Hostname = m.GetString("hostname")
		topologyRecovery.AnalysisEntry.AnalyzedInstanceKey.Port = m.GetInt("port")
		topologyRecovery.AnalysisEntry.Analysis = inst.AnalysisCode(m.GetString("analysis"))
		topologyRecovery.AnalysisEntry.ClusterDetails.ClusterName = m.GetString("cluster_name")
		topologyRecovery.AnalysisEntry.ClusterDetails.ClusterAlias = m.GetString("cluster_alias")
		topologyRecovery.AnalysisEntry.CountSlaves = m.GetUint("count_affected_slaves")
		topologyRecovery.AnalysisEntry.ReadSlaveHostsFromString(m.GetString("slave_hosts"))

		topologyRecovery.SuccessorKey.Hostname = m.GetString("successor_hostname")
		topologyRecovery.SuccessorKey.Port = m.GetInt("successor_port")

		topologyRecovery.AnalysisEntry.ClusterDetails.ReadRecoveryInfo()

		res = append(res, topologyRecovery)
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadActiveRecoveries reads active recovery entry/audit entires from topology_recovery
func ReadActiveClusterRecovery(clusterName string) ([]TopologyRecovery, error) {
	whereClause := fmt.Sprintf(`
		where 
			in_active_period=1 
			and end_recovery is null
			and cluster_name='%s'`,
		clusterName)
	return readRecoveries(whereClause, ``)
}

// ReadRecentlyActiveInstanceRecovery reads recently completed entries for a given cluster
func ReadRecentlyActiveClusterRecovery(clusterName string) ([]TopologyRecovery, error) {
	whereClause := fmt.Sprintf(`
		where 
			end_recovery > now() - interval 5 minute 
			and cluster_name='%s'`,
		clusterName)
	return readRecoveries(whereClause, ``)
}

// ReadRecentlyActiveInstanceRecovery reads recently completed entries for a given instance
func ReadRecentlyActiveInstanceRecovery(instanceKey *inst.InstanceKey) ([]TopologyRecovery, error) {
	whereClause := fmt.Sprintf(`
		where 
			end_recovery > now() - interval 5 minute 
			and 
				successor_hostname='%s' and successor_port=%d`,
		instanceKey.Hostname, instanceKey.Port)
	return readRecoveries(whereClause, ``)
}

// ReadActiveRecoveries reads active recovery entry/audit entires from topology_recovery
func ReadActiveRecoveries() ([]TopologyRecovery, error) {
	return readRecoveries(`
		where 
			in_active_period=1 
			and end_recovery is null`,
		``)
}

// ReadCompletedRecoveries reads completed recovery entry/audit entires from topology_recovery
func ReadCompletedRecoveries(page int) ([]TopologyRecovery, error) {
	limit := fmt.Sprintf(`
		limit %d
		offset %d`,
		config.Config.AuditPageSize, page*config.Config.AuditPageSize)
	return readRecoveries(`where end_recovery is not null`, limit)
}

// ReadCRecoveries reads latest recovery entreis from topology_recovery
func ReadRecentRecoveries(page int) ([]TopologyRecovery, error) {
	limit := fmt.Sprintf(`
		limit %d
		offset %d`,
		config.Config.AuditPageSize, page*config.Config.AuditPageSize)
	return readRecoveries(``, limit)
}

// readRecoveries reads recovery entry/audit entires from topology_recovery
func readFailureDetections(whereCondition string, limit string) ([]TopologyRecovery, error) {
	res := []TopologyRecovery{}
	query := fmt.Sprintf(`
		select 
            detection_id,
            hostname,
            port,
            in_active_period as is_active,
            start_active_period,
            end_active_period_unixtime,
            processing_node_hostname,
            processcing_node_token,
            analysis,
            cluster_name,
            cluster_alias,
            count_affected_slaves,
            slave_hosts		
		from 
			topology_failure_detection
		%s
		order by
			detection_id desc
		%s
		`, whereCondition, limit)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		failureDetection := TopologyRecovery{}
		failureDetection.Id = m.GetInt64("detection_id")

		failureDetection.IsActive = m.GetBool("is_active")
		failureDetection.RecoveryStartTimestamp = m.GetString("start_active_period")
		failureDetection.ProcessingNodeHostname = m.GetString("processing_node_hostname")
		failureDetection.ProcessingNodeToken = m.GetString("processcing_node_token")

		failureDetection.AnalysisEntry.AnalyzedInstanceKey.Hostname = m.GetString("hostname")
		failureDetection.AnalysisEntry.AnalyzedInstanceKey.Port = m.GetInt("port")
		failureDetection.AnalysisEntry.Analysis = inst.AnalysisCode(m.GetString("analysis"))
		failureDetection.AnalysisEntry.ClusterDetails.ClusterName = m.GetString("cluster_name")
		failureDetection.AnalysisEntry.ClusterDetails.ClusterAlias = m.GetString("cluster_alias")
		failureDetection.AnalysisEntry.CountSlaves = m.GetUint("count_affected_slaves")
		failureDetection.AnalysisEntry.ReadSlaveHostsFromString(m.GetString("slave_hosts"))

		failureDetection.AnalysisEntry.ClusterDetails.ReadRecoveryInfo()

		res = append(res, failureDetection)
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadCRecoveries reads latest recovery entreis from topology_recovery
func ReadRecentFailureDetections(page int) ([]TopologyRecovery, error) {
	limit := fmt.Sprintf(`
		limit %d
		offset %d`,
		config.Config.AuditPageSize, page*config.Config.AuditPageSize)
	return readFailureDetections(``, limit)
}
