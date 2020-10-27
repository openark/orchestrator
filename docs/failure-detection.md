# Failure detection

`orchestrator` uses a holistic approach to detect master and intermediate master failures.

In a naive approach, a monitoring tool would probe the master, for example, and alert when is cannot contact or query the master server. Such approach is susceptible to false positives caused by network glitches. The naive approach further mitigates false positives by running `n` tests spaced by `t`-long intervals. This reduces, on some cases, the chance for false positive, and then increases the response time on real failure.

`orchestrator` harnesses the replication topology. It observes not only the server itself, but also its replicas. For example, to diagnose a dead master scenario, `orchestrator` must both:

- Fail to contact said master
- Be able to contact master's replicas, and confirm they, too, cannot see the master.

Instead of triaging the error by time, `orchestrator` triages by multiple observers, the replication topology servers themselves. In fact, when all master's replicas agree that they cannot contact their master, the replication topology is broken _de-facto_, and a failover is justified.

`orchestrator`'s holistic failure detection approach is known to be very reliable in production.

### Detection and recovery

Detection does not always lead to [recovery](topology-recovery.md). There are scenarios where a recovery is undesired:

- The cluster is not listed for auto-failovers.
- The admin user has indicated a recovery should not take place on the specific server.
- The admin user has globally disabled recoveries.
- A previous recovery on same cluster completed shortly before, and anti-flapping takes place.
- The failure type is not deemed worthy of recovery.

On desired scenarios, recovery immediately follows detection. On others, such as blocked recoveries, a recovery may follow a detection after long minutes.

Detection is independent of recovery, and is always enabled. `OnFailureDetectionProcesses` hooks are executed per detection, see [failure detection configuration](configuration-failure-detection.md).

### Failure detection scenarios

Observe the following list of potential failures:

* DeadMaster
* DeadMasterAndReplicas
* DeadMasterAndSomeReplicas
* DeadMasterWithoutReplicas
* UnreachableMasterWithLaggingReplicas
* UnreachableMaster
* LockedSemiSyncMaster
* AllMasterReplicasNotReplicating
* AllMasterReplicasNotReplicatingOrDead
* DeadCoMaster
* DeadCoMasterAndSomeReplicas
* DeadIntermediateMaster
* DeadIntermediateMasterWithSingleReplicaFailingToConnect
* DeadIntermediateMasterWithSingleReplica
* DeadIntermediateMasterAndSomeReplicas
* DeadIntermediateMasterAndReplicas
* AllIntermediateMasterReplicasFailingToConnectOrDead
* AllIntermediateMasterReplicasNotReplicating
* UnreachableIntermediateMasterWithLaggingReplicas
* UnreachableIntermediateMaster
* BinlogServerFailingToConnectToMaster

Briefly looking at some examples, here is how `orchestrator` reaches failure conclusions:

#### `DeadMaster`:

1. Master MySQL access failure
2. All of master's replicas are failing replication

This makes for a potential recovery process

#### `DeadMasterAndSomeReplicas`:

1. Master MySQL access failure
2. Some of its replicas are also unreachable
3. Rest of the replicas are failing replication

This makes for a potential recovery process

#### `UnreachableMaster`:

1. Master MySQL access failure
2. But it has replicating replicas.

This does not make for a recovery process. However, to improve analysis, `orchestrator` will
issue an emergent re-read of the replicas, to figure out whether they are really happy with the master
(in which case maybe `orchestrator` cannot see it due to a network glitch) or were actually taking
their time to figure out they were failing replication.

#### `DeadIntermediateMaster`:

1. An intermediate master (replica with replicas) cannot be reached
2. All of its replicas are failing replication

This makes for a potential recovery process.

#### `UnreachableMasterWithLaggingReplicas`:

1. Master cannot be reached
2. All of its immediate replicas (excluding SQL delayed) are lagging

This scenario can happen when the master is overloaded. Clients would see a "Too many connections", while the replicas, which have been connected since long ago, claim the master is fine. Similarly, if the master is locked due to some metadata operation, clients would be blocked on connection while replicas _may claim_ everything's fine. However, since apps cannot connect to the master, no actual data gets written, and when using a heartbeat mechanism such as `pt-heartbeat`, we can observe a growing lag on replicas.

`orchestrator` responds to this scenario by restarting replication on all of master's immediate replicas. This will close the old client connections on those replicas and attempt to initiate new ones. These may now fail to connect, leading to a complete replication failure on all replicas. This will next lead `orchestrator` to analyze a `DeadMaster`.

### LockedSemiSyncMaster

1. Master is running with semi-sync enabled
2. Number of connected semi-sync replicas falls short of expected `rpl_semi_sync_master_wait_for_slave_count`
3. `rpl_semi_sync_master_timeout` is high enough such that master locks writes and does not fall back to asynchronous replication

Remediation can be to disable semi-sync on the master, or to bring up (or enable) sufficient semi-sync replicas.

At this time `orchestrator` does not invoke processes for this type of analysis.

### Failures of no interest

The following scenarios are of no interest to `orchestrator`, and while the information and state are available to `orchestrator`, it does not recognize such scenarios as _failures_ per se; there's no detection hooks invoked and obviously no recoveries attempted:

- Failure of simple replicas (_leaves_ on the replication topology graph)
  Exception: semi sync replicas causing `LockedSemiSyncMaster`
- Replication lags, even severe.

### Visibility

An up-to-date analysis is available via:

- Command line: `orchestrator-client -c replication-analysis`
  or `orchestrator -c replication-analysis`
- Web API: `/api/replication-analysis`
- Web: `/web/clusters-analysis/` page (`Clusters`->`Failure analysis`).
  This presents an incomplete list of problems, only highlighting actionable ones.

Read next: [Topology recovery](topology-recovery.md)
