# Topology recovery

`orchestrator` is able to recover from a set of [failure scenarios](failure-detection.md). Notably, it can recover a failed master or a failed intermediate master.

At this time recovery requires either GTID (Oracle or MariaDB), [Pseudo GTID](#pseudo-gtid) or Binlog Servers.

Automated recovery is _opt in_. Please consider [recovery configuration](configuration-recovery.md).

### What's in a recovery?

Based on a [failure detection](failure-detection.md), a sequence of events compose a recovery:

- Pre-recovery hooks (external processes execution)
- Healing of topology
- Post-recovery hooks

Where:

- Pre-recovery hooks are configured by the user. Failure of any will abort the failover.
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

### Web, API, command line

Recoveries are audited via:

- `/web/audit-recovery`
- `/api/audit-recovery`
- `/api/audit-recovery-steps/:uid`

Nuance auditing and control available via:
- `/api/blocked-recoveries`: see blocked recoveries
- `/api/ack-recovery/cluster/:clusterHint`: acknowledge a recovery on a given cluster
- `/api/disable-global-recoveries`: global switch to disable `orchestrator` from running any recoveries
- `/api/enable-global-recoveries`: re-enable recoveries
- `/api/check-global-recoveries`: check is global recoveries are enabled

Running manual recoveries (see next sections):

- `/api/recover/:host/:port`: recover specific host, assuming `orchestrator` agrees there is failure.
- `/api/recover-lite/:host/:port`: same, do not invoke external hooks (can be useful for testing)
- `/api/graceful-master-takeover/:clusterHint`: gracefully promote a new master (planned failover)
- `/api/force-master-failover/:clusterHint`: panic, force master failover for given cluster

Some corresponding command line invocations:

- `orchestrator-client -c recover -i some.instance:3306`
- `orchestrator-client -c graceful-master-takeover -i some.instance.in.somecluster:3306`
- `orchestrator-client -c graceful-master-takeover -alias somecluster`
- `orchestrator-client -c force-master-takeover -alias somecluster`
- `orchestrator-client -c ack-cluster-recoveries -alias somecluster`
- `orchestrator-client -c disable-global-recoveries`
- `orchestrator-client -c enable-global-recoveries`
- `orchestrator-client -c check-global-recoveries`

### Automated recovery

Opt-in. Automated recovery may be applied to all (`"*"`) clusters or specific ones.

Recovery follows detection, and assuming nothing block recovery (read further below)

For greater resolution, different configuration applies for master recovery and for intermediate-master recovery. Detailed breakdown of recovery-related configuration follows.

The analysis mechanism runs at all times, and checks periodically for failure/recovery scenarios. It will make an automated recovery for:

- An actionable type of scenario
- For an instance that is not downtimed
- For an instance belonging to a cluster for which recovery is explicitly enabled via configuration
- For an instance in a cluster that has not recently been recovered, unless such recent recoveries were acknowledged
- Where global recoveries are enabled

### Manual recovery

TL;DR use this when an instance is recognized as failed but where auto-recovery is disabled or blocked.

You may choose to ask `orchestrator` to recover a failure by providing a specific instance that is failed.
This instance _must be recognized as having a failure_. It is possible to request recovery for an instance that is downtimed (as this is manual recovery it overrides automatic assumptions).
Recover via:

* Command line: `orchestrator-client -c recover -i dead.instance.com:3306 --debug`
* Web API: `/api/recover/dead.instance.com/:3306`
* Web: instance is colored black; click the `Recover` button

Manual recoveries don't block on `RecoveryPeriodBlockSeconds` (read more in next section). They also override `RecoverMasterClusterFilters` and `RecoverIntermediateMasterClusterFilters`. A manual recovery will only block on an already running (and incomplete) recovery on the very same instance the manual recovery wishes to operate on.

### Manual, forced failover

TL;DR force master failover _right now_ regardless of what `orchestrator` thinks.

Perhaps `orchestrator` doesn't see that the instance is failed, or you have some app-logic that requires the master must change _right now_, or perhaps the type of failure is such that `orchestrator` is unsure about. You wish to kick a master failover _right now_. You will run:

* Command line: `orchestrator-client -c force-master-failover --alias mycluster`

  or `orchestrator-client -c force-master-failover -i instance.in.that.cluster`
* Web API: `/api/force-master-failover/mycluster`

  or `/api/force-master-failover/instance.in.that.cluster/3306`


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
an instance. An instance can be downtimed (via `orchestrator-client -c begin-downtime`) and this is noted in the analysis summary. When considering automated recovery, downtimed servers are skipped.

Downtime was, in fact, explicitly created for this very purpose, and allows the DBA a way to suppress automated failover and a specific server.

Note that manual recovery (e.g. `orchestrator-client -c recover`) overrides downtime.

### Recovery hooks

`orchestrator` supports hooks -- external scripts invoked through the recovery process. These are arrays of commands invoked via shell, in particular `bash`. See hook configuration details in [recovery configuration](configuration-recovery.md#hooks)

- `OnFailureDetectionProcesses`: described in [failure detection](failure-detection.md).
- `PreFailoverProcesses`
- `PostMasterFailoverProcesses`
- `PostIntermediateMasterFailoverProcesses`
- `PostFailoverProcesses`
- `PostUnsuccessfulFailoverProcesses`
