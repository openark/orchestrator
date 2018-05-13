# Topology recovery

`orchestrator` is able to recover from a set of [failure scenarios](failure-detection.md). Notably, it can recover a failed master or a failed intermediate master.

## Automated and manual

`orchestrator` supports:

- [Automated recovery](#automated-recovery) (takes action on unexpected failures).
- [Graceful, planned, master promotion](#graceful-master-promotion).
- [Manual recovery](#manual-recovery).
- [Manual, forced/panic failovers](#manual-forced-failover).

## Requirements 

To run any kind of failovers, your topologies must support either:

- Oracle GTID
- MariaDB GTID
- [Pseudo GTID](#pseudo-gtid)
- Binlog Servers

Automated recovery is _opt in_. Please consider [recovery configuration](configuration-recovery.md).

### What's in a recovery?

Based on a [failure detection](failure-detection.md), a sequence of events compose a recovery:

- Pre-recovery hooks (external processes execution)
- Healing of topology
- Post-recovery hooks

Where:

- Pre-recovery hooks are configured by the user. 
  - Executed sequentially
  - Failure of any of these hooks (nonzero exit code) will abort the failover.
- Topology healing is managed by `orchestrator` and is state-based rather than configuration-based. `orchestrator` tries to make the best out of a bad situation, taking into consideration the existing topology, versions, server configurations, etc.
- Post-recovery hooks are, again, configured by the user.

### Discussion: recovering a dead intermediate master

The following highlights some of the complexity of a recovery.

A "simple" recovery case is that of a `DeadIntermediateMaster`. Its replicas are orphaned, but when
using GTID or Pseudo-GTID they can still be re-connected to the topology. We might choose to:

- Find a sibling of the dead intermediate master, and move orphaned replicas below said sibling
- Promote a replica from among the orphaned replicas, make it intermediate master of its siblings, then
  connect promoted replica up the topology
- relocate all orphaned replicas
- Combine parts of the above

The exact implementation greatly depends on the topology setup (which instances have `log-slave-updates`? Are instances lagging? Do they have replication filters? Which versions of MySQL? etc.). It is very (very) likely your topology will support at least one of the above (in particular, matching-up the replicas is a trivial solution, unless replication filters are in place).

### Discussion: recovering a dead master

Recovering from a dead master is a much more complex operation, for various reasons:

- There is outage implied and recovery is expected to be as fast as possible.
- Some servers may be lost in the process. `orchestrator` must determine which, if any.
- State of the topology may be such that the user wishes to prevent the recovery.
- Master service discovery must take place: the app must be able to communicate with the new master (potentially be _told_ that the master has changed).

- Find the best replica to promote.
  - A naive approach would be to pick the most up-to-date replica, but that may not always be the right choice.
  - It may so happen that the most up-to-date replica will not have the necessary configuration to act as master to other replicas (e.g. binlog format, MySQL versioning, replication filters and more). By blindly promoting the most up-to-date replica one may lose replica capacity.
  - `orchestrator` attempts to promote a replica that will retain the most serving capacity.
- Promote said replica, taking over its siblings.
- Bring siblings up to date
- Possibly, do a 2nd phase promotion; the user may have tagged specific servers to be promoted if possible (see `register-candidate` command).
- Call upon hooks (read further)

Master service discovery is largely the user's responsibility to implement. Common solutions are:
- DNS based discovery; `orchestrator` will need to invoke a hook that modifies DNS entries.
- ZooKeeper/Consul KV/etcd/other key-value based discovery; `orchestrator` has built-in support for Consul KV, otherwise an external hook must update KV stores
- Proxy based discovery; `orchestrator` will call upon external hook that modifies proxy config, or will update Consul/Zk/etcd as above, which will itself trigger an update to proxy configuration.
- Other solutions...

`orchestrator` attempts to be a generic solution hence takes no stance on your service discovery method.


### Automated recovery

Opt-in. Automated recovery may be applied to all (`"*"`) clusters or specific ones.

Recovery follows detection, and assuming nothing blocks recovery (read further below)

For greater resolution, different configuration applies for master recovery and for intermediate-master recovery. Detailed breakdown of recovery-related configuration follows.

The analysis mechanism runs at all times, and checks periodically for failure/recovery scenarios. It will make an automated recovery for:

- An actionable type of scenario
- For an instance that is not downtimed
- For an instance belonging to a cluster for which recovery is explicitly enabled via configuration
- For an instance in a cluster that has not recently been recovered, unless such recent recoveries were acknowledged
- Where global recoveries are enabled

### Graceful master promotion

TL;DR use this for replacing a master in an orderly, planned operation.

Typically for purposes of upgrade, host maintenance, etc., you may wish to replace an existing master with another one. This is a gracefull promotion (aka _graceful master takeover_).

In a graceful takeover:

- You indicate a server to promote.
- `orchestrator` turns your master to be `read-only`.
- `orchestrator` makes sure your designated server is caught up with replication.
- `orchestrator` promotes your designated server as the new master.
- `orchestrator` turns promoted server to be writable.

The operation can take a few seconds, during which time your app is expected to complain, seeing that the master is `read-only`.

In addition to standard hooks, `orchestrator` provides you with specialized hooks to run a graceful takeover:

- `PreGracefulTakeoverProcesses`
- `PostGracefulTakeoverProcesses`

For example, you may want to disable the pager for the duration of a planned failover. Advanced usage may include stalling traffic at proxy layer.

In a graceful promotion you must either:

- Indicate the designated master (must be a direct replica of the existing master)
- Set your topology such that there is exactly one direct replica under the master (at which case the identity of the designated replica is trivial and needs not be mentioned).

Invoke graceful takeover via:

* Command line: `orchestrator-client -c graceful-master-takeover -alias mycluster -s designated.master.to.promote:3306`

* Web API: 

  - `/api/graceful-master-takeover/:clusterHint/:designatedHost/:designatedPort`: gracefully promote a new master (planned failover), indicating the designated master to promote.
  - `/api/graceful-master-takeover/:clusterHint`: gracefully promote a new master (planned failover). Designated server not indicated, works when the master has exactly one direct replica.

* Web interface: drag a direct master's replica onto the left half of the master's box.

### Manual recovery

TL;DR use this when an instance is recognized as failed but where auto-recovery is disabled or blocked.

You may choose to ask `orchestrator` to recover a failure by providing a specific instance that is failed.
This instance _must be recognized as having a failure_. It is possible to request recovery for an instance that is downtimed (as this is manual recovery it overrides automatic assumptions).
Recover via:

* Command line: `orchestrator-client -c recover -i dead.instance.com:3306 --debug`
* Web API: `/api/recover/dead.instance.com/:3306`
* Web: instance is colored black; click the `Recover` button

Manual recoveries don't block on `RecoveryPeriodBlockSeconds` (read more in next section). They also override `RecoverMasterClusterFilters` and `RecoverIntermediateMasterClusterFilters`. Thus, a human can always invoke a recovery by demand. A recovery may only block on yet another recovery running at that time on the same database instance.

### Manual, forced failover

TL;DR force master failover _right now_ regardless of what `orchestrator` thinks.

Perhaps `orchestrator` doesn't see that the instance is failed, or you have some app-logic that requires the master must change _right now_, or perhaps the type of failure is such that `orchestrator` is unsure about. You wish to kick a master failover _right now_. You will run:

* Command line: `orchestrator-client -c force-master-failover --alias mycluster`

  or `orchestrator-client -c force-master-failover -i instance.in.that.cluster`
* Web API: `/api/force-master-failover/mycluster`

  or `/api/force-master-failover/instance.in.that.cluster/3306`


### Web, API, command line

Recoveries are audited via:

- `/web/audit-recovery`
- `/api/audit-recovery`
- `/api/audit-recovery-steps/:uid`

Nuance auditing and control available via:
- `/api/blocked-recoveries`: see blocked recoveries
- `/api/ack-recovery/cluster/:clusterHint`: acknowledge a recovery on a given cluster
- `/api/ack-all-recoveries`: acknowledge all recoveries
- `/api/disable-global-recoveries`: global switch to disable `orchestrator` from running any recoveries
- `/api/enable-global-recoveries`: re-enable recoveries
- `/api/check-global-recoveries`: check is global recoveries are enabled

Running manual recoveries (see next sections):

- `/api/recover/:host/:port`: recover specific host, assuming `orchestrator` agrees there is failure.
- `/api/recover-lite/:host/:port`: same, do not invoke external hooks (can be useful for testing)
- `/api/graceful-master-takeover/:clusterHint/:designatedHost/:designatedPort`: gracefully promote a new master (planned failover), indicating the designated master to promote.
- `/api/graceful-master-takeover/:clusterHint`: gracefully promote a new master (planned failover). Designated server not indicated, works when the master has exactly one direct replica.
- `/api/force-master-failover/:clusterHint`: panic, force master failover for given cluster

Some corresponding command line invocations:

- `orchestrator-client -c recover -i some.instance:3306`
- `orchestrator-client -c graceful-master-takeover -i some.instance.in.somecluster:3306`
- `orchestrator-client -c graceful-master-takeover -alias somecluster`
- `orchestrator-client -c force-master-takeover -alias somecluster`
- `orchestrator-client -c ack-cluster-recoveries -alias somecluster`
- `orchestrator-client -c ack-all-recoveries`
- `orchestrator-client -c disable-global-recoveries`
- `orchestrator-client -c enable-global-recoveries`
- `orchestrator-client -c check-global-recoveries`

#### Blocking, acknowledgements, anti-flapping

`orchestrator` avoid flapping (cascading failures causing continuous outage and elimination of resources) by introducing a block period, where on any given cluster, `orchesrartor` will not kick in automated recovery on an interval smaller than said period, unless cleared to do so by a human.

The block period is indicated by `RecoveryPeriodBlockSeconds`. It only applies to recoveries on _same cluster_. There is nothing to prevent concurrent recoveries running on _different clusters_.

Pending recoveries are unblocked either once `RecoveryPeriodBlockSeconds` has passed or such a recovery has been _acknowledged_.

Acknowledging a recovery is possible either via web API/interface (see audit/recovery page) or via command line interface (`orchestrator-client -c ack-cluster-recoveries -alias somealias`).

Note that manual recovery (e.g. `orchestrator-client -c recover` or `orchstrator-client -c force-master-failover`) ignores the blocking period.

### Downtime

All failure/recovery scenarios are analyzed. However also taken into consideration is the downtime status of
an instance. An instance can be downtimed (via `orchestrator-client -c begin-downtime`) and this is noted in the analysis summary. When considering automated recovery, downtimed servers are skipped.

Downtime was, in fact, explicitly created for this very purpose, and allows the DBA a way to suppress automated failover and a specific server.

Note that manual recovery (e.g. `orchestrator-client -c recover`) overrides downtime.

### Recovery hooks

`orchestrator` supports hooks -- external scripts invoked through the recovery process. These are arrays of commands invoked via shell, in particular `bash`. See hook configuration details in [recovery configuration](configuration-recovery.md#hooks)

- `OnFailureDetectionProcesses`: described in [failure detection](failure-detection.md).
- `PreGracefulTakeoverProcesses`: invoked on `graceful-master-takeover` command, before master goes `read-only`.
- `PreFailoverProcesses`
- `PostMasterFailoverProcesses`
- `PostIntermediateMasterFailoverProcesses`
- `PostFailoverProcesses`
- `PostUnsuccessfulFailoverProcesses`
- `PostGracefulTakeoverProcesses`: executed on planned, graceful master takeover, after the old master is positioned under the newly promoted master.
