# Topology recovery

`orchestrator` detects failure scenarios and optionally recovers in such cases.

At this time recovery requires either GTID, [Pseudo GTID](#pseudo-gtid) or Binlog Servers.

### Failure/recovery scenarios


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

This makes for a potential recovery process

### What are the current failure/recovery scenarios?

Some of the analysis above lead to recovery processes (depending on configuration) and some do not.
You may see up to date analysis via:

- Command line: `orchestrator -c replication-analysis`
- Web API: `/api/replication-analysis`
- Web: `/web/clusters-analysis/` page (`Clusters`->`Failure analysis`).
  This presents an incomplete list of problems, only highlighting actionable ones.

Note that recovery is not concerned with the death of a single replica machine, as it implies
no required changes to topology.

### What's in a recovery?

A "simple" recovery case is that of a `DeadIntermediateMaster`. Its replicas are orphaned, but when
using GTID or Pseudo-GTID they can still be re-connected to the topology. We might choose to:

- Find a sibling of the dead intermediate master, and move orphaned replicas below said sibling
- Promote a replica from among the orphaned replicas, make it intermediate master of its siblings, then
  connect promoted replica up the topology
- relocate all orphaned replicas
- Combine parts of the above

The exact implementation greatly depends on the topology setup (which instances have `log-slave-updates`? Are instances lagging? Do they
have replication filters? Which versions of MySQL? etc.). It is very (very) likely your topology will support at least one of the above
(in particular, matching-up the replicas is a trivial solution, unless replication filters are in place).

A "really-not-simple" recovery case is that of a `DeadMaster`. `orchestrator` will try to:

- Find the most advanced replica
- Promote it, enslaving its siblings
- Bring siblings up to date
- Check if there was a pre-defined "candidate replica"
- Promote that over the previously chosen replica
- And... call upon hooks (read further)

Success of the above depends on how many replicas have `log-slave-updates`, and whether those that are not configured as such
are more up-to-date or less up-to-date than others. When all your instances have `log-slave-updates` the problem is greatly simplified.
Of course, all other limitations apply (versions, binlog format, replication filters) - and orchestrator will attempt to find a good solution.

The `DeadMaster` case is not simple because your application must be made aware that a master is dead and another takes its place.
This is strictly outside `orchestrator`'s scope; however `orchestrator` supports hooks: lists of external processes to invoke
before/after recovery. These would do whatever changes needed: remove "read-only" flag, change DNS records, etc. It's on you.

Hooks are described in detail further on.

As with all operations, major steps & decisions are audited (see `/api/audit`) and of course logged. The backend `topology_recovery`
holds the state for recovery operations, if you like to SQL your way for information.

### Manual recovery

You may choose to ask `orchestrator` to recover a failure by providing a specific instance that is failed.
This instance must be recognized as having a failure scenario (see above). It is possible to request
recovery for an instance that is downtimed (as this is manual recovery it overrides automatic assumptions).
Recover via:

* Command line: `orchestrator -c recover -i dead.instance.com:3306 --debug`
* Web API: `/api/recover/dead.instance.com/:3306`
* Web: instance is colored black; click the `Recover` button

Manual recoveries don't block on `RecoveryPeriodBlockSeconds` (read more in next section). They also override
`RecoverMasterClusterFilters` and `RecoverIntermediateMasterClusterFilters`. A manual recovery will only block on
an already running (and incomplete) recovery on the very same instance the manual recovery wishes to operate on.

### Automated recovery

By default turned off, automatic recovery may be applied for specific clusters. For greater resolution, different configuration
applies for master recovery and for intermediate-master recovery. Detailed breakdown of recovery-related configuration follows.

The analysis mechanism runs at all times, and checks periodically for failure/recovery scenarios. It will make an
automated recovery for:

- An actionable type of scenario (duh)
- For an instance that is not downtimed
- For an instance belonging to a cluster for which recovery is explicitly enabled via configuration
- For an instance in a cluster that has not recently been recovered, unless such recent recoveries were acknowledged

#### Blocking, acknowledgements, anti-flapping

The last bullet is made so as to avoid flapping of recovery processes. Each recovery process is audited and recorded, and
two recoveries cannot run on the same instance unless a minimal amount of time has passed (indicated by `RecoveryPeriodBlockSeconds`).

Moreover, no two automated recoveries will be executed for the same _cluster_ in an interval shorter than `RecoveryPeriodBlockSeconds` (this of course a stronger condition than the previous one). The first recovery to be detected wins and the others block.

There is nothing to prevent concurrent recoveries running on _different clusters_.

Pending recoveries are unblocked either once `RecoveryPeriodBlockSeconds` has passed or such a recovery has been _acknowledged_.
Acknowledging a recovery is possible either via web API/interface (see audit/recovery page) or via command line interface (see `-c ack-instance-recoveries` or `-c ack-cluster-recoveries`).

> The observant reader may be confused a little as for the reason for two different commands: `ack-instance-recoveries` and `ack-cluster-recoveries`, or for the different discussion on "same instance" vs "same cluster".
> A recovery can be unacknowledged. While it is unacknowledged and still _recent_, there will be no automated recovery taking place. However if sufficient time has passed, it is perfectly valid that the recovery remains _unacknowledged_ even as `orchestrator` proceeds to recover a new failure scenario.
> It is therefore possible to have multiple unacknowledged recoveries on same instance or on same cluster.
> In the future, it may happen that we will change the behavior (based on configuration) such that constraints on recoveries of same cluster are relaxed. We therefore keep separate code and analysis for "same instance" vs. "same cluster">

As with all operations, a recovery process puts "maintenance" locks on instances, and will not be able to refactor an instance
already under maintenance. Furthermore, it will place a recovery lock on the instance. This protects against multiple clients
all trying to recover the same failure scenario.

### Downtime

All failure/recovery scenarios are analyzed. However also taken into consideration is the downtime status of
an instance. If an instance is downtimed (via `orchestrator -c begin-downtime`) this is noted in the
analysis summary. When considering automated recovery, downtimed servers are ignored.
Downtime is explicitly created for this purpose: to allow the DBA a way to suppress automated failover.


### Recovery hooks

`orchestrator` supports hooks -- external scripts invoked through the recovery process. These are arrays of commands invoked through
shell, in particular `bash`. Hooks are:

- `OnFailureDetectionProcesses`: called when a failure/recovery known scenario is detected. These scripts are called befroe even
  deciding whether action should be taken.
- `PreFailoverProcesses`: called after deciding to take action on a scenario. Order of execution is sequential. A failure
  (non-zero exit status) of any process *aborts the recovery operation*. This is your chance to decide whether to go on with
  the recovery or not.
- `PostIntermediateMasterFailoverProcesses`: specific commands to run after recovery of an intermediate-master failure.
  Order of execution is sequential. Failures are ignored (the recovery is done in terms of `orchestrator`).
- `PostMasterFailoverProcesses`: specific commands to run after recovery of a master failure.
  Order of execution is sequential. Failures are ignored (the recovery is done in terms of `orchestrator`).
- `PostFailoverProcesses`: commands to run after recovery of any type (and following the specific `PostIntermediateMasterFailoverProcesses`
  or `PostMasterFailoverProcesses` commands). Failures are ignored.
- `PostUnsuccessfulFailoverProcesses`: commands to run when recovery operation resulted with error, such that there is no known successor instance

#### Hooks arguments and environment

`orchestrator` provides all hooks with failure/recovery related information, such as the identity of the failed instance, identity of promoted instance, affecetd replicas, type of failure, name of cluster, etc.

This information is passed independently in two ways, and you may choose to use one or both:

1. Environment variables: `orchestrator` will set the following, which can be retrieved by your hooks:

- `ORC_FAILURE_TYPE`
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

And, in the event a recovery was successful:

- `ORC_SUCCESSOR_HOST`
- `ORC_SUCCESSOR_PORT`
- `ORC_SUCCESSOR_ALIAS`

2. Command line text replacement. `orchestrator` replaces the following magic tokens in your `*Proccesses` commands:

- `{failureType}`
- `{failureDescription}`
- `{failedHost}`
- `{failedPort}`
- `{failureCluster}`
- `{failureClusterAlias}`
- `{failureClusterDomain}`
- `{countReplicas}` aka `{countSlaves}`
- `{isDowntimed}`
- `{autoMasterRecovery}`
- `{autoIntermediateMasterRecovery}`
- `{orchestratorHost}`
- `{lostReplicas}` aka `{lostSlaves}`
- `{replicaHosts}` aka `{slaveHosts}`
- `{isSuccessful}`

And, in the event a recovery was successful:

- `{successorHost}`
- `{successorPort}`
- `{successorAlias}`


### Recovery configuration

Elaborating on recovery-related configuration:

- `RecoveryPeriodBlockSeconds`: minimal amount of seconds between two recoveries on same instance or same cluster (default: `3600`)
- `RecoveryIgnoreHostnameFilters`: Recovery analysis will completely ignore hosts matching given patterns (these could be, for example, test servers, dev machines that are in the topologies)
- `RecoverMasterClusterFilters`: list of cluster names, aliases or patterns that are included in automatic recovery for master failover. As an example:
```
  "RecoverMasterClusterFilters": [
    "myoltp", // cluster name includes this string
    "meta[0-9]+", // cluster name matches this regex
    "alias=olap", // cluster alias is exactly "olap"
    "alias~=shard[0-9]+" // cluster alias matches this pattern
  ]
```

- `RecoverIntermediateMasterClusterFilters`: list of cluster names, aliases or patterns that are included in automatic recovery for intermediate-master failover. Format is as above.
  Note that the `".*"` pattern matches everything.
- `PromotionIgnoreHostnameFilters`: instances matching given regex patterns will not be picked by orchestrator for promotion (these could be, for example, test servers, dev machines that are in the topologies)

- `FailureDetectionPeriodBlockMinutes`: a detection does not necessarily lead to a recovery (for example, the instance may be downtimed). This variable indicates the minimal time interval between invocation of `OnFailureDetectionProcesses`.

- `RecoveryPollSeconds`: poll interval at which orchestrator re-checks for crash scenarios (default: 10s)

- `DetachLostReplicasAfterMasterFailover`: in the case of master promotion and assuming that some replicas could not make it into
the refactored topology, should orchestrator forcibly issue a `detach-replica` command to make sure they don't accidentally resume
replication in the future.

- `PostponeSlaveRecoveryOnLagMinutes`: some recovery operations can be pushed to be the very last steps; so that more urgent
operations (e.g. change DNS entries) could be applied faster. Fixing replicas that are lagging at time of recovery (either because of `MASTER_DELAY` configuration or just because they were busy) could take a substantial time due to binary log exhaustive search (GTID & Pseudo-GTID). This variable defines the threshold above which a lagging replica's rewiring is pushed till the last moment.

- `ApplyMySQLPromotionAfterMasterFailover`: after master promotion, should orchestrator take it upon itself to clear the `read_only` flag & forcibly detach replication? (default: `false`)
