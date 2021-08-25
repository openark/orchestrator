# Configuration: recovery

`orchestrator` will recover failures to your topologies. You will instruct `orchestrator` which clusters to auto-recover and which to expect a human to recover. You will configure hooks for `orchestrator` to move VIPs, update service discovery etc.

Recovery depends on detection, discussed in [configuration: failure detection](configuration-failure-detection.md)

See [Topology recovery](topology-recovery.md) for all things recoveries.

Also consider that your MySQL topologies themselves need to follow some rules, see [MySQL Configuration](#mysql-configuration)

```json
{
  "RecoveryPeriodBlockSeconds": 3600,
  "RecoveryIgnoreHostnameFilters": [],
  "RecoverMasterClusterFilters": [
    "thiscluster",
    "thatcluster"
  ],
  "RecoverIntermediateMasterClusterFilters": [
    "*"
  ],
}
```

In the above:

- `orchestrator` will auto-recover intermediate master failures for all clusters
- `orchestrator` will auto-recover master failures for two specified clusters; masters of other clusters will not auto recover. A human will be able to initiate recoveries.
- Once a cluster experienced a recovery, `orchestrator` will block auto-recoveries for `3600` seconds (`1` hour) following. This is an anti-flapping mechanism.

Note, again, that automated recovery is _opt in_.

### Promotion actions

Different environments require different actions taken on recovery/promotion

```json
{
  "ApplyMySQLPromotionAfterMasterFailover": true,
  "PreventCrossDataCenterMasterFailover": false,
  "PreventCrossRegionMasterFailover": false,
  "FailMasterPromotionOnLagMinutes": 0,
  "FailMasterPromotionIfSQLThreadNotUpToDate": true,
  "DelayMasterPromotionIfSQLThreadNotUpToDate": false,
  "MasterFailoverLostInstancesDowntimeMinutes": 10,
  "DetachLostReplicasAfterMasterFailover": true,
  "MasterFailoverDetachReplicaMasterHost": false,
  "MasterFailoverLostInstancesDowntimeMinutes": 0,
  "PostponeReplicaRecoveryOnLagMinutes": 0,
}
```

- `ApplyMySQLPromotionAfterMasterFailover`: when `true`, `orchestrator` will `reset slave all` and `set read_only=0` on promoted master. Default: `true`. When `true`, overrides `MasterFailoverDetachSlaveMasterHost`.
- `PreventCrossDataCenterMasterFailover`: defaults `false`. When `true`, `orchestrator` will only replace a failed master with a server from the same DC. It will do its best to find a replacement from same DC, and will abort (fail) the failover if it cannot find one. See also `DetectDataCenterQuery` and `DataCenterPattern` configuration variables.
- `PreventCrossRegionMasterFailover`: defaults `false`. When `true`, `orchestrator` will only replace a failed master with a server from the same region. It will do its best to find a replacement from same region, and will abort (fail) the failover if it cannot find one. See also `DetectRegionQuery` and `RegionPattern` configuration variables.
- `FailMasterPromotionOnLagMinutes`: defaults `0` (not failing promotion). Can be used to fail a promotion if the candidate replica is too far behind. Example: replicas were broken for 5 hours, and then master failed. One might want to prevent the failover in order to recover the binary logs / relay logs for those lost 5 hours.
  To use this flag, you must set `ReplicationLagQuery` and use a heartbeat mechanism such as `pt-heartbeat`. The MySQL built-in `Seconds_behind_master` output of `SHOW SLAVE STATUS` (pre 8.0) does not report replication lag when replication is broken.
- `FailMasterPromotionIfSQLThreadNotUpToDate`: if all replicas were lagging at time of failure, even the most up-to-date, promoted replica may yet have unapplied relay logs. Issuing `reset slave all` on such a server will lose the relay log data. Your choice.
- `DelayMasterPromotionIfSQLThreadNotUpToDate`: if all replicas were lagging at time of failure, even the most up-to-date, promoted replica may yet have unapplied relay logs. When `true`, 'orchestrator' will wait for the SQL thread to catch up before promoting a new master.
  `FailMasterPromotionIfSQLThreadNotUpToDate` and `DelayMasterPromotionIfSQLThreadNotUpToDate` are mutually exclusive.
- `DetachLostReplicasAfterMasterFailover`: some replicas may get lost during recovery. When `true`, `orchestrator` will forcibly break their replication via `detach-replica` command to make sure no one assumes they're at all functional.
- `MasterFailoverDetachReplicaMasterHost` : when `true`, `orchestrator` will issue a `detach-replica-master-host` on promoted master (this makes sure the new master will not attempt to replicate old master if that comes back to life). Default: `false`. Meaningless if `ApplyMySQLPromotionAfterMasterFailover` is `true`. `MasterFailoverDetachSlaveMasterHost` is an alias to this.
- `MasterFailoverLostInstancesDowntimeMinutes`: number of minutes to downtime any server that was lost after a master failover (including failed master & lost replicas). Set to 0 to disable. Default: 0.
- `PostponeReplicaRecoveryOnLagMinutes`: on crash recovery, replicas that are lagging more than given minutes are only resurrected late in the recovery process, after master/IM has been elected and processes executed. Value of 0 disables this feature. Default: 0. `PostponeSlaveRecoveryOnLagMinutes` is an alias to this.

### Hooks

These hooks are available for recoveries:

- `PreGracefulTakeoverProcesses`: executed on planned, graceful master takeover, immediately before the master goes `read-only`.
- `PreFailoverProcesses`: executed immediately before `orchestrator` takes recovery action. Failure (nonzero exit code) of any of these processes aborts the recovery.
  Hint: this gives you the opportunity to abort recovery based on some internal state of your system.
- `PostMasterFailoverProcesses`: executed at the end of a successful master recovery.
- `PostIntermediateMasterFailoverProcesses`: executed at the end of a successful intermediate master or replication 
  group member with replicas recovery.
- `PostFailoverProcesses`: executed at the end of any successful recovery (including and adding to the above two).
- `PostUnsuccessfulFailoverProcesses`: executed at the end of any unsuccessful recovery.
- `PostGracefulTakeoverProcesses`: executed on planned, graceful master takeover, after the old master is positioned under the newly promoted master.

Any process command that ends with `"&"` will be executed asynchronously, and a failure for such process is ignored.

All of the above are lists of commands which `orchestrator` executes sequentially, in order of definition.

A naive implementation might look like:

```json
{
  "PreGracefulTakeoverProcesses": [
    "echo 'Planned takeover about to take place on {failureCluster}. Master will switch to read_only' >> /tmp/recovery.log"
  ],
  "PreFailoverProcesses": [
    "echo 'Will recover from {failureType} on {failureCluster}' >> /tmp/recovery.log"
  ],
  "PostFailoverProcesses": [
    "echo '(for all types) Recovered from {failureType} on {failureCluster}. Failed: {failedHost}:{failedPort}; Successor: {successorHost}:{successorPort}' >> /tmp/recovery.log"
  ],
  "PostUnsuccessfulFailoverProcesses": [],
  "PostMasterFailoverProcesses": [
    "echo 'Recovered from {failureType} on {failureCluster}. Failed: {failedHost}:     {failedPort}; Promoted: {successorHost}:{successorPort}' >> /tmp/recovery.log"
  ],
  "PostIntermediateMasterFailoverProcesses": [],
  "PostGracefulTakeoverProcesses": [
    "echo 'Planned takeover complete' >> /tmp/recovery.log"
  ],
}
```

#### Hooks arguments and environment

`orchestrator` provides all hooks with failure/recovery related information, such as the identity of the failed instance, identity of promoted instance, affected replicas, type of failure, name of cluster, etc.

This information is passed independently in two ways, and you may choose to use one or both:

1. Environment variables: `orchestrator` will set the following, which can be retrieved by your hooks:

- `ORC_FAILURE_TYPE`
- `ORC_INSTANCE_TYPE` ("master", "co-master", "intermediate-master")
- `ORC_IS_MASTER` (true/false)
- `ORC_IS_CO_MASTER` (true/false)
- `ORC_FAILURE_DESCRIPTION`
- `ORC_FAILED_HOST`
- `ORC_FAILED_PORT`
- `ORC_FAILURE_CLUSTER`
- `ORC_FAILURE_CLUSTER_ALIAS`
- `ORC_FAILURE_CLUSTER_DOMAIN`
- `ORC_COUNT_REPLICAS`
- `ORC_IS_DOWNTIMED`
- `ORC_AUTO_MASTER_RECOVERY`
- `ORC_AUTO_INTERMEDIATE_MASTER_RECOVERY`
- `ORC_ORCHESTRATOR_HOST`
- `ORC_IS_SUCCESSFUL`
- `ORC_LOST_REPLICAS`
- `ORC_REPLICA_HOSTS`
- `ORC_COMMAND` (`"force-master-failover"`, `"force-master-takeover"`, `"graceful-master-takeover"` if applicable)

And, in the event a recovery was successful:

- `ORC_SUCCESSOR_HOST`
- `ORC_SUCCESSOR_PORT`
- `ORC_SUCCESSOR_BINLOG_COORDINATES`
- `ORC_SUCCESSOR_ALIAS`

2. Command line text replacement. `orchestrator` replaces the following magic tokens in your `*Proccesses` commands:

- `{failureType}`
- `{instanceType}` ("master", "co-master", "intermediate-master")
- `{isMaster}` (true/false)
- `{isCoMaster}` (true/false)
- `{failureDescription}`
- `{failedHost}`
- `{failedPort}`
- `{failureCluster}`
- `{failureClusterAlias}`
- `{failureClusterDomain}`
- `{countReplicas}` (replaces `{countSlaves}`)
- `{isDowntimed}`
- `{autoMasterRecovery}`
- `{autoIntermediateMasterRecovery}`
- `{orchestratorHost}`
- `{lostReplicas}` (replaces `{lostSlaves}`)
- `{countLostReplicas}`
- `{replicaHosts}` (replaces `{slaveHosts}`)
- `{isSuccessful}`
- `{command}` (`"force-master-failover"`, `"force-master-takeover"`, `"graceful-master-takeover"` if applicable)

And, in the event a recovery was successful:

- `{successorHost}`
- `{successorPort}`
- `{successorBinlogCoordinates}`
- `{successorAlias}`

### MySQL Configuration

Your MySQL topologies must fulfill some requirements in order to support failovers. Those requirements largely depends on the types of topologies/configuration you use.

- Oracle/Percona with GTID: promotable servers must have `log_bin` and `log_slave_updates` enabled. Replicas must be using `AUTO_POSITION=1` (via `CHANGE MASTER TO MASTER_AUTO_POSITION=1`).
- MariaDB GTID: promotable servers must have `log_bin` and `log_slave_updates` enabled.
- [Pseudo GTID](#pseudo-gtid): promotable servers must have `log_bin` and `log_slave_updates` enabled. If using `5.7/8.0` parallel replication, set `slave_preserve_commit_order=1`.
- BinlogServers: promotable servers must have `log_bin` enabled.


Also consider improving failure detection via [MySQL Configuration](configuration-failure-detection.md#mysql-configuration)
