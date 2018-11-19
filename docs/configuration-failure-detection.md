# Configuration: failure detection

`orchestrator` will [detect failures](failure-detection.md) to your topology, always. As a matter of configuration you may set the polling frequency and specific ways for `orchestrator` to notify you on such detection.

Recovery is discussed in [configuration: recovery](configuration-recovery.md)

```json
{
  "FailureDetectionPeriodBlockMinutes": 60,
}
```

`orchestrator` runs detection every second.

`FailureDetectionPeriodBlockMinutes` is an anti-spam mechanism that blocks `orchestrator` from notifying the same detection again and again and again.

### Hooks

Configure `orchestrator` to take action on discovery:

```json
{
  "OnFailureDetectionProcesses": [
    "echo 'Detected {failureType} on {failureCluster}. Affected replicas: {countReplicas}' >> /tmp/recovery.log"
  ],
}
```

There are many magic variables (as `{failureCluster}`, above) that you can send to your external hooks. See full list in [Topology recovery](topology-recovery.md)

### MySQL configuration

Since failure detection uses the MySQL topology itself as a source of information, it is advisable that you setup your MySQL replication such that errors will be clearly indicated or quickly mitigated.

- `set global slave_net_timeout = 4`, see [docuemntation](https://dev.mysql.com/doc/refman/5.7/en/replication-options-slave.html#sysvar_slave_net_timeout). This sets a short (`2sec`) heartbeat interval between a replica and its master, and will make the replica recognize failure quickly. Without this setting, some scenarios may take up to a minute to detect.
- `CHANGE MASTER TO MASTER_CONNECT_RETRY=1, MASTER_RETRY_COUNT=86400`. In the event of replication failure, make the replica attempt reconnection every `1sec` (default is `60sec`). With brief network issues this setting attempts a quick replication recovery and, if successful, will avoid a general failure/recovery operation by `orchestrator`.
