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
* DeadMasterAndSlaves
* DeadMasterAndSomeSlaves
* DeadMasterWithoutSlaves
* UnreachableMaster
* AllMasterSlavesNotReplicating
* AllMasterSlavesNotReplicatingOrDead
* DeadCoMaster
* DeadCoMasterAndSomeSlaves
* DeadIntermediateMaster
* DeadIntermediateMasterWithSingleSlaveFailingToConnect
* DeadIntermediateMasterWithSingleSlave
* DeadIntermediateMasterAndSomeSlaves
* DeadIntermediateMasterAndSlaves
* AllIntermediateMasterSlavesFailingToConnectOrDead
* AllIntermediateMasterSlavesNotReplicating
* UnreachableIntermediateMaster
* BinlogServerFailingToConnectToMaster

Briefly looking at some examples, here is how `orchestrator` reaches failure conclusions:

#### `DeadMaster`:

1. Master MySQL access failure
2. All of master's replicas are failing replication

This makes for a potential recovery process

#### `DeadMasterAndSomeSlaves`:

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

### Failures of no interest

The following scenarios are of no interest to `orchestrator`, and while the information and state are available to `orchestrator`, it does not recognize such scenarios as _failures_ per se; there's no detection hooks invoked and obviously no recoveries attempted:

- Failure of simple replicas (_leaves_ on the replication topology graph)
- Replication lags, even severe.

### Visibility

An up-to-date analysis is available via:

- Command line: `orchestrator-client -c replication-analysis`
  or `orchestrator -c replication-analysis`
- Web API: `/api/replication-analysis`
- Web: `/web/clusters-analysis/` page (`Clusters`->`Failure analysis`).
  This presents an incomplete list of problems, only highlighting actionable ones.

Read next: [Topology recovery](topology-recovery.md)
