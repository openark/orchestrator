# Configuration: recovery

`orchestrator` will recover failures to your topologies. You will instruct `orchestrator` which clusters to auto-recover and which to expect a human to recover. You will configure hooks for `orchestrator` to move VIPs, update service discovery etc.

Recovery depends on detection, discussed in [configuration: failure detection](configuration-failure-detection.md)

See [Topology recovery](topology-recovery.md) for all things recoveries.

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
- Once a cluster experienced a recovery, `orchestrator` will block auto-recoveries for `1` hour following. This is an anti-flapping mechanism.

Note, again, that automated recovery is _opt in_.

### Promotion actions

Different environments require different actions taken on recovery/promotion

```json
{
  "ApplyMySQLPromotionAfterMasterFailover": true,
  "MasterFailoverLostInstancesDowntimeMinutes": 10,
  "FailMasterPromotionIfSQLThreadNotUpToDate": true,
  "DetachLostReplicasAfterMasterFailover": true,
}
```

- `ApplyMySQLPromotionAfterMasterFailover`: when `true`, `orchestrator` will `reset slave all` and `set read_only=0` on promoted master.
- `FailMasterPromotionIfSQLThreadNotUpToDate`: if all replicas were lagging at time of failure, even the most up-to-date, promoted replica may yet have unapplied relay logs. Issuing `reset slave all` on such a server will lose the relay log data. Your choice.
- `DetachLostReplicasAfterMasterFailover`: some replicas may get lost during recovery. When `true`, `orchestrator` will forcibly break their replication via `detach-replica` command to make sure no one assumes they're at all functional.

### Hooks

These hooks are available for recoveries:

- `PreFailoverProcesses`: executed immediately before `orchestrator` takes recovery action. Failure (nonzero exit code) of any of these processes aborts the recovery.
  Hint: this gives you the opportunity to abort recovery based on some internal state of your system.
- `PostMasterFailoverProcesses`: executed at the end of a successful master recovery.
- `PostIntermediateMasterFailoverProcesses`: executed at the end of a successful intermediate master recovery.
- `PostFailoverProcesses`: executed at the end of any successful recovery (including and adding to the above two).
- `PostUnsuccessfulFailoverProcesses`: executed at the end of any unsuccessful recovery.

All of the above are lists of commands which `orchestrator` executes sequentially, in order of definition.

A naive implementation might look like:

```json
{
  "PreFailoverProcesses": [
    "echo 'Will recover from {failureType} on {failureCluster}' >> /tmp/recovery.log"
  ],
  "PostFailoverProcesses": [
    "echo '(for all types) Recovered from {failureType} on {failureCluster}. Failed:      {failedHost}:{failedPort}; Successor: {successorHost}:{successorPort}' >> /tmp/     recovery.log"
  ],
  "PostUnsuccessfulFailoverProcesses": [],
  "PostMasterFailoverProcesses": [
    "echo 'Recovered from {failureType} on {failureCluster}. Failed: {failedHost}:     {failedPort}; Promoted: {successorHost}:{successorPort}' >> /tmp/recovery.log"
  ],
  "PostIntermediateMasterFailoverProcesses": [],
}
```
