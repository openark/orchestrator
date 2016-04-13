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
	"strings"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
	"github.com/outbrain/orchestrator/go/inst"
	"github.com/outbrain/orchestrator/go/process"
)

// AttemptFailureDetectionRegistration tries to add a failure-detection entry; if this fails that means the problem has already been detected
func AttemptFailureDetectionRegistration(analysisEntry *inst.ReplicationAnalysis) (bool, error) {
	sqlResult, err := db.ExecOrchestrator(`
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
			`, analysisEntry.AnalyzedInstanceKey.Hostname, analysisEntry.AnalyzedInstanceKey.Port, process.ThisHostname, process.ProcessToken.Hash,
		string(analysisEntry.Analysis), analysisEntry.ClusterDetails.ClusterName, analysisEntry.ClusterDetails.ClusterAlias, analysisEntry.CountSlaves, analysisEntry.SlaveHosts.ToCommaDelimitedList(),
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

// clearAcknowledgedFailureDetections clears the "in_active_period" flag for detections
// that were acknowledged
func clearAcknowledgedFailureDetections(whereClause string, args []interface{}) error {
	query := fmt.Sprintf(`
			update topology_failure_detection set
				in_active_period = 0,
				end_active_period_unixtime = UNIX_TIMESTAMP()
			where
				in_active_period = 1
				and %s
			`, whereClause)
	_, err := db.ExecOrchestrator(query, args...)
	return log.Errore(err)
}

// AttemptRecoveryRegistration tries to add a recovery entry; if this fails that means recovery is already in place.
func AttemptRecoveryRegistration(analysisEntry *inst.ReplicationAnalysis, failIfFailedInstanceInActiveRecovery bool, failIfClusterInActiveRecovery bool) (*TopologyRecovery, error) {
	if failIfFailedInstanceInActiveRecovery {
		// Let's check if this instance has just been promoted recently and is still in active period.
		// If so, we reject recovery registration to avoid flapping.
		recoveries, err := ReadInActivePeriodSuccessorInstanceRecovery(&analysisEntry.AnalyzedInstanceKey)
		if err != nil {
			return nil, log.Errore(err)
		}
		if len(recoveries) > 0 {
			RegisterBlockedRecoveries(analysisEntry, recoveries)
			return nil, log.Errorf("AttemptRecoveryRegistration: instance %+v has recently been promoted (by failover of %+v) and is in active period. It will not be failed over. You may acknowledge the failure on %+v (-c ack-instance-recoveries) to remove this blockage", analysisEntry.AnalyzedInstanceKey, recoveries[0].AnalysisEntry.AnalyzedInstanceKey, recoveries[0].AnalysisEntry.AnalyzedInstanceKey)
		}
	}
	if failIfClusterInActiveRecovery {
		// Let's check if this cluster has just experienced a failover and is still in active period.
		// If so, we reject recovery registration to avoid flapping.
		recoveries, err := ReadInActivePeriodClusterRecovery(analysisEntry.ClusterDetails.ClusterName)
		if err != nil {
			return nil, log.Errore(err)
		}
		if len(recoveries) > 0 {
			RegisterBlockedRecoveries(analysisEntry, recoveries)
			return nil, log.Errorf("AttemptRecoveryRegistration: cluster %+v has recently experienced a failover (of %+v) and is in active period. It will not be failed over again. You may acknowledge the failure on this cluster (-c ack-cluster-recoveries) or on %+v (-c ack-instance-recoveries) to remove this blockage", analysisEntry.ClusterDetails.ClusterName, recoveries[0].AnalysisEntry.AnalyzedInstanceKey, recoveries[0].AnalysisEntry.AnalyzedInstanceKey)
		}
	}
	if !failIfFailedInstanceInActiveRecovery {
		// Implicitly acknowledge this instance's possibly existing active recovery, provided they are completed.
		AcknowledgeInstanceCompletedRecoveries(&analysisEntry.AnalyzedInstanceKey, "orchestrator", fmt.Sprintf("implicit acknowledge due to user invocation of recovery on same instance: %+v", analysisEntry.AnalyzedInstanceKey))
		// The fact we only acknowledge a completed recovery solves the possible case of two DBAs simultaneously
		// trying to recover the same instance at the same time
	}

	sqlResult, err := db.ExecOrchestrator(`
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
					slave_hosts,
					last_detection_id
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
					?,
					(select ifnull(max(detection_id), 0) from topology_failure_detection where hostname=? and port=?)
				)
			`, analysisEntry.AnalyzedInstanceKey.Hostname, analysisEntry.AnalyzedInstanceKey.Port, process.ThisHostname, process.ProcessToken.Hash,
		string(analysisEntry.Analysis), analysisEntry.ClusterDetails.ClusterName, analysisEntry.ClusterDetails.ClusterAlias, analysisEntry.CountSlaves, analysisEntry.SlaveHosts.ToCommaDelimitedList(),
		analysisEntry.AnalyzedInstanceKey.Hostname, analysisEntry.AnalyzedInstanceKey.Port,
	)
	if err != nil {
		return nil, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	if err != nil {
		return nil, log.Errore(err)
	}
	if rows == 0 {
		return nil, nil
	}
	// Success
	topologyRecovery := NewTopologyRecovery(*analysisEntry)
	topologyRecovery.Id, _ = sqlResult.LastInsertId()
	return topologyRecovery, nil
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
				AND start_active_period < NOW() - INTERVAL ? SECOND
			`,
		config.Config.RecoveryPeriodBlockSeconds,
	)
	return log.Errore(err)
}

// RegisterBlockedRecoveries writes down currently blocked recoveries, and indicates what recovery they are blocked on.
// Recoveries are blocked thru the in_active_period flag, which comes to avoid flapping.
func RegisterBlockedRecoveries(analysisEntry *inst.ReplicationAnalysis, blockingRecoveries []TopologyRecovery) error {
	for _, recovery := range blockingRecoveries {
		_, err := db.ExecOrchestrator(`
			insert
				into blocked_topology_recovery (
					hostname,
					port,
					cluster_name,
					analysis,
					last_blocked_timestamp,
					blocking_recovery_id
				) values (
					?,
					?,
					?,
					?,
					NOW(),
					?
				)
				on duplicate key update
					cluster_name=values(cluster_name),
					analysis=values(analysis),
					last_blocked_timestamp=values(last_blocked_timestamp),
					blocking_recovery_id=values(blocking_recovery_id)
			`, analysisEntry.AnalyzedInstanceKey.Hostname,
			analysisEntry.AnalyzedInstanceKey.Port,
			analysisEntry.ClusterDetails.ClusterName,
			string(analysisEntry.Analysis),
			recovery.Id,
		)
		if err != nil {
			log.Errore(err)
		}
	}
	return nil
}

// ExpireBlockedRecoveries clears listing of blocked recoveries that are no longer actually blocked.
func ExpireBlockedRecoveries() error {
	// Older recovery is acknowledged by now, hence blocked recovery should be released.
	// Do NOTE that the data in blocked_topology_recovery is only used for auditing: it is NOT the data
	// based on which we make automated decisions.
	_, err := db.ExecOrchestrator(`
			delete
				from blocked_topology_recovery
				using
					blocked_topology_recovery
					left join topology_recovery on (blocking_recovery_id = topology_recovery.recovery_id and acknowledged = 0)
				where
					acknowledged is null
			`,
	)
	if err != nil {
		return log.Errore(err)
	}
	// Some oversampling, if a problem has not been noticed for some time (e.g. the server came up alive
	// before action was taken), expire it.
	// Recall that RegisterBlockedRecoveries continuously updates the last_blocked_timestamp column.
	_, err = db.ExecOrchestrator(`
			delete
				from blocked_topology_recovery
				where
					last_blocked_timestamp < NOW() - interval ? second
			`, (config.Config.RecoveryPollSeconds * 2),
	)
	if err != nil {
		return log.Errore(err)
	}
	return nil
}

// acknowledgeRecoveries sets acknowledged* details and clears the in_active_period flags from a set of entries
func acknowledgeRecoveries(owner string, comment string, markEndRecovery bool, whereClause string, args []interface{}) (countAcknowledgedEntries int64, err error) {
	additionalSet := ``
	if markEndRecovery {
		additionalSet = `
				end_recovery=IFNULL(end_recovery, NOW()),
			`
	}
	query := fmt.Sprintf(`
			update topology_recovery set
				in_active_period = 0,
				end_active_period_unixtime = IF(end_active_period_unixtime = 0, UNIX_TIMESTAMP(), end_active_period_unixtime),
				%s
				acknowledged = 1,
				acknowledged_at = NOW(),
				acknowledged_by = ?,
				acknowledge_comment = ?
			where
				acknowledged = 0
				and
				%s
		`, additionalSet, whereClause)
	args = append(sqlutils.Args(owner, comment), args...)
	sqlResult, err := db.ExecOrchestrator(query, args...)
	if err != nil {
		return 0, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	return rows, log.Errore(err)
}

// AcknowledgeRecovery acknowledges a particular recovery.
// This also implied clearing their active period, which in turn enables further recoveries on those topologies
func AcknowledgeRecovery(recoveryId int64, owner string, comment string) (countAcknowledgedEntries int64, err error) {
	whereClause := `recovery_id = ?`
	return acknowledgeRecoveries(owner, comment, false, whereClause, sqlutils.Args(recoveryId))
}

// AcknowledgeClusterRecoveries marks active recoveries for given cluster as acknowledged.
// This also implied clearing their active period, which in turn enables further recoveries on those topologies
func AcknowledgeClusterRecoveries(clusterName string, owner string, comment string) (countAcknowledgedEntries int64, err error) {
	whereClause := `cluster_name = ?`
	args := sqlutils.Args(clusterName)
	clearAcknowledgedFailureDetections(whereClause, args)
	return acknowledgeRecoveries(owner, comment, false, whereClause, args)
}

// AcknowledgeInstanceRecoveries marks active recoveries for given instane as acknowledged.
// This also implied clearing their active period, which in turn enables further recoveries on those topologies
func AcknowledgeInstanceRecoveries(instanceKey *inst.InstanceKey, owner string, comment string) (countAcknowledgedEntries int64, err error) {
	whereClause := `
			hostname = ?
			and port = ?
		`
	args := sqlutils.Args(instanceKey.Hostname, instanceKey.Port)
	clearAcknowledgedFailureDetections(whereClause, args)
	return acknowledgeRecoveries(owner, comment, false, whereClause, args)
}

// AcknowledgeInstanceCompletedRecoveries marks active and COMPLETED recoveries for given instane as acknowledged.
// This also implied clearing their active period, which in turn enables further recoveries on those topologies
func AcknowledgeInstanceCompletedRecoveries(instanceKey *inst.InstanceKey, owner string, comment string) (countAcknowledgedEntries int64, err error) {
	whereClause := `
			hostname = ?
			and port = ?
			and end_recovery is not null
		`
	return acknowledgeRecoveries(owner, comment, false, whereClause, sqlutils.Args(instanceKey.Hostname, instanceKey.Port))
}

// AcknowledgeCrashedRecoveries marks recoveries whose processing nodes has crashed as acknowledged.
func AcknowledgeCrashedRecoveries() (countAcknowledgedEntries int64, err error) {
	whereClause := `
			in_active_period = 1
			and end_recovery is null
			and (processing_node_hostname, processcing_node_token) not in (
				select hostname, token from node_health
			)
		`
	return acknowledgeRecoveries("orchestrator", "detected crashed recovery", true, whereClause, sqlutils.Args())
}

// ResolveRecovery is called on completion of a recovery process and updates the recovery status.
// It does not clear the "active period" as this still takes place in order to avoid flapping.
func ResolveRecovery(topologyRecovery *TopologyRecovery, successorInstance *inst.Instance) error {

	isSuccessful := false
	var successorKeyToWrite inst.InstanceKey
	var successorAliasToWrite string
	if successorInstance != nil {
		topologyRecovery.SuccessorKey = &successorInstance.Key
		topologyRecovery.SuccessorAlias = successorInstance.InstanceAlias
		isSuccessful = true
		successorKeyToWrite = successorInstance.Key
	}
	_, err := db.ExecOrchestrator(`
			update topology_recovery set
				is_successful = ?,
				successor_hostname = ?,
				successor_port = ?,
				successor_alias = ?,
				lost_slaves = ?,
				participating_instances = ?,
				all_errors = ?,
				end_recovery = NOW()
			where
				recovery_id = ?
				AND in_active_period = 1
				AND processing_node_hostname = ?
				AND processcing_node_token = ?
			`, isSuccessful, successorKeyToWrite.Hostname, successorKeyToWrite.Port,
		successorAliasToWrite, topologyRecovery.LostSlaves.ToCommaDelimitedList(),
		topologyRecovery.ParticipatingInstanceKeys.ToCommaDelimitedList(),
		strings.Join(topologyRecovery.AllErrors, "\n"),
		topologyRecovery.Id, process.ThisHostname, process.ProcessToken.Hash,
	)
	return log.Errore(err)
}

// readRecoveries reads recovery entry/audit entires from topology_recovery
func readRecoveries(whereCondition string, limit string, args []interface{}) ([]TopologyRecovery, error) {
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
            is_successful,
            processing_node_hostname,
            processcing_node_token,
            ifnull(successor_hostname, '') as successor_hostname,
            ifnull(successor_port, 0) as successor_port,
            ifnull(successor_alias, '') as successor_alias,
            analysis,
            cluster_name,
            cluster_alias,
            count_affected_slaves,
            slave_hosts,
            participating_instances,
            lost_slaves,
            all_errors,
            acknowledged,
            acknowledged_at,
            acknowledged_by,
            acknowledge_comment,
            last_detection_id
		from
			topology_recovery
		%s
		order by
			recovery_id desc
		%s
		`, whereCondition, limit)
	err := db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		topologyRecovery := *NewTopologyRecovery(inst.ReplicationAnalysis{})
		topologyRecovery.Id = m.GetInt64("recovery_id")

		topologyRecovery.IsActive = m.GetBool("is_active")
		topologyRecovery.RecoveryStartTimestamp = m.GetString("start_active_period")
		topologyRecovery.RecoveryEndTimestamp = m.GetString("end_recovery")
		topologyRecovery.IsSuccessful = m.GetBool("is_successful")
		topologyRecovery.ProcessingNodeHostname = m.GetString("processing_node_hostname")
		topologyRecovery.ProcessingNodeToken = m.GetString("processcing_node_token")

		topologyRecovery.AnalysisEntry.AnalyzedInstanceKey.Hostname = m.GetString("hostname")
		topologyRecovery.AnalysisEntry.AnalyzedInstanceKey.Port = m.GetInt("port")
		topologyRecovery.AnalysisEntry.Analysis = inst.AnalysisCode(m.GetString("analysis"))
		topologyRecovery.AnalysisEntry.ClusterDetails.ClusterName = m.GetString("cluster_name")
		topologyRecovery.AnalysisEntry.ClusterDetails.ClusterAlias = m.GetString("cluster_alias")
		topologyRecovery.AnalysisEntry.CountSlaves = m.GetUint("count_affected_slaves")
		topologyRecovery.AnalysisEntry.ReadSlaveHostsFromString(m.GetString("slave_hosts"))

		topologyRecovery.SuccessorKey = &inst.InstanceKey{}
		topologyRecovery.SuccessorKey.Hostname = m.GetString("successor_hostname")
		topologyRecovery.SuccessorKey.Port = m.GetInt("successor_port")
		topologyRecovery.SuccessorAlias = m.GetString("successor_alias")

		topologyRecovery.AnalysisEntry.ClusterDetails.ReadRecoveryInfo()

		topologyRecovery.AllErrors = strings.Split(m.GetString("all_errors"), "\n")
		topologyRecovery.LostSlaves.ReadCommaDelimitedList(m.GetString("lost_slaves"))
		topologyRecovery.ParticipatingInstanceKeys.ReadCommaDelimitedList(m.GetString("participating_instances"))

		topologyRecovery.Acknowledged = m.GetBool("acknowledged")
		topologyRecovery.AcknowledgedAt = m.GetString("acknowledged_at")
		topologyRecovery.AcknowledgedBy = m.GetString("acknowledged_by")
		topologyRecovery.AcknowledgedComment = m.GetString("acknowledge_comment")

		topologyRecovery.LastDetectionId = m.GetInt64("last_detection_id")

		res = append(res, topologyRecovery)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadActiveRecoveries reads active recovery entry/audit entires from topology_recovery
func ReadActiveClusterRecovery(clusterName string) ([]TopologyRecovery, error) {
	whereClause := `
		where
			in_active_period=1
			and end_recovery is null
			and cluster_name=?`
	return readRecoveries(whereClause, ``, sqlutils.Args(clusterName))
}

// ReadInActivePeriodClusterRecovery reads recoveries (possibly complete!) that are in active period.
// (may be used to block further recoveries on this cluster)
func ReadInActivePeriodClusterRecovery(clusterName string) ([]TopologyRecovery, error) {
	whereClause := `
		where
			in_active_period=1
			and cluster_name=?`
	return readRecoveries(whereClause, ``, sqlutils.Args(clusterName))
}

// ReadRecentlyActiveClusterRecovery reads recently completed entries for a given cluster
func ReadRecentlyActiveClusterRecovery(clusterName string) ([]TopologyRecovery, error) {
	whereClause := `
		where
			end_recovery > now() - interval 5 minute
			and cluster_name=?`
	return readRecoveries(whereClause, ``, sqlutils.Args(clusterName))
}

// ReadInActivePeriodSuccessorInstanceRecovery reads completed recoveries for a given instance, where said instance
// was promoted as result, still in active period (may be used to block further recoveries should this instance die)
func ReadInActivePeriodSuccessorInstanceRecovery(instanceKey *inst.InstanceKey) ([]TopologyRecovery, error) {
	whereClause := `
		where
			in_active_period=1
			and
				successor_hostname=? and successor_port=?`
	return readRecoveries(whereClause, ``, sqlutils.Args(instanceKey.Hostname, instanceKey.Port))
}

// ReadRecentlyActiveInstanceRecovery reads recently completed entries for a given instance
func ReadRecentlyActiveInstanceRecovery(instanceKey *inst.InstanceKey) ([]TopologyRecovery, error) {
	whereClause := `
		where
			end_recovery > now() - interval 5 minute
			and
				successor_hostname=? and successor_port=?`
	return readRecoveries(whereClause, ``, sqlutils.Args(instanceKey.Hostname, instanceKey.Port))
}

// ReadActiveRecoveries reads active recovery entry/audit entires from topology_recovery
func ReadActiveRecoveries() ([]TopologyRecovery, error) {
	return readRecoveries(`
		where
			in_active_period=1
			and end_recovery is null`,
		``, sqlutils.Args())
}

// ReadCompletedRecoveries reads completed recovery entry/audit entires from topology_recovery
func ReadCompletedRecoveries(page int) ([]TopologyRecovery, error) {
	limit := `
		limit ?
		offset ?`
	return readRecoveries(`where end_recovery is not null`, limit, sqlutils.Args(config.Config.AuditPageSize, page*config.Config.AuditPageSize))
}

// ReadRecovery reads completed recovery entry/audit entires from topology_recovery
func ReadRecovery(recoveryId int64) ([]TopologyRecovery, error) {
	whereClause := `where recovery_id = ?`
	return readRecoveries(whereClause, ``, sqlutils.Args(recoveryId))
}

// ReadCRecoveries reads latest recovery entries from topology_recovery
func ReadRecentRecoveries(clusterName string, unacknowledgedOnly bool, page int) ([]TopologyRecovery, error) {
	whereConditions := []string{}
	whereClause := ""
	args := sqlutils.Args()
	if unacknowledgedOnly {
		whereConditions = append(whereConditions, `acknowledged=0`)
	}
	if clusterName != "" {
		whereConditions = append(whereConditions, `cluster_name=?`)
		args = append(args, clusterName)
	}
	if len(whereConditions) > 0 {
		whereClause = fmt.Sprintf("where %s", strings.Join(whereConditions, " and "))
	}
	limit := `
		limit ?
		offset ?`
	args = append(args, config.Config.AuditPageSize, page*config.Config.AuditPageSize)
	return readRecoveries(whereClause, limit, args)
}

// readRecoveries reads recovery entry/audit entires from topology_recovery
func readFailureDetections(whereCondition string, limit string, args []interface{}) ([]TopologyRecovery, error) {
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
            slave_hosts,
            (select max(recovery_id) from topology_recovery where topology_recovery.last_detection_id = detection_id) as related_recovery_id
		from
			topology_failure_detection
		%s
		order by
			detection_id desc
		%s
		`, whereCondition, limit)
	err := db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
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

		failureDetection.RelatedRecoveryId = m.GetInt64("related_recovery_id")

		failureDetection.AnalysisEntry.ClusterDetails.ReadRecoveryInfo()

		res = append(res, failureDetection)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadRecentFailureDetections
func ReadRecentFailureDetections(page int) ([]TopologyRecovery, error) {
	limit := `
		limit ?
		offset ?`
	return readFailureDetections(``, limit, sqlutils.Args(config.Config.AuditPageSize, page*config.Config.AuditPageSize))
}

// ReadFailureDetection
func ReadFailureDetection(detectionId int64) ([]TopologyRecovery, error) {
	whereClause := `where detection_id = ?`
	return readFailureDetections(whereClause, ``, sqlutils.Args(detectionId))
}

// ReadBlockedRecoveries reads blocked recovery entries, potentially filtered by cluster name (empty to unfilter)
func ReadBlockedRecoveries(clusterName string) ([]BlockedTopologyRecovery, error) {
	res := []BlockedTopologyRecovery{}
	whereClause := ""
	args := sqlutils.Args()
	if clusterName != "" {
		whereClause = `where cluster_name = ?`
		args = append(args, clusterName)
	}
	query := fmt.Sprintf(`
		select
				hostname,
				port,
				cluster_name,
				analysis,
				last_blocked_timestamp,
				blocking_recovery_id
			from
				blocked_topology_recovery
			%s
			order by
				last_blocked_timestamp desc
		`, whereClause)
	err := db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		blockedTopologyRecovery := BlockedTopologyRecovery{}
		blockedTopologyRecovery.FailedInstanceKey.Hostname = m.GetString("hostname")
		blockedTopologyRecovery.FailedInstanceKey.Port = m.GetInt("port")
		blockedTopologyRecovery.ClusterName = m.GetString("cluster_name")
		blockedTopologyRecovery.Analysis = inst.AnalysisCode(m.GetString("analysis"))
		blockedTopologyRecovery.LastBlockedTimestamp = m.GetString("last_blocked_timestamp")
		blockedTopologyRecovery.BlockingRecoveryId = m.GetInt64("blocking_recovery_id")

		res = append(res, blockedTopologyRecovery)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}
