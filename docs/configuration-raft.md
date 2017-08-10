# Configuration: raft

Set up a [orchestrator/raft](raft.md) cluster for high availability.

Assuming you will run `orchestrator/raft` on a `3` node setup, you will configure this on each node:

```json
  "RaftEnabled": true,
  "RaftBind": "<ip.or.fqdn.of.this.orchestrator.node>",
  "DefaultRaftPort": 10008,
  "RaftNodes": [
    "<ip.or.fqdn.of.orchestrator.node1>",
    "<ip.or.fqdn.of.orchestrator.node2>",
    "<ip.or.fqdn.of.orchestrator.node3>"
  ],
```

Some breakdown:

- `RaftEnabled` must be set to `true`, otherwise `orchestrator` runs in shared-backend mode.
- `RaftBind` must be set, use the IP address or full hostname of local host. This IP or hostname will also be listed as one of the `RaftNodes` variable.
- `DefaultRaftPort` can be set to any port, but must be consistent across all deployments.
- `RaftNodes` should list all nodes of the raft cluster. This list will consist of either IP addresses or host names, and will include the value of this host itself as presented in `RaftBind`.

As example, the following might be a working setup:

```json
  "RaftEnabled": true,
  "RaftBind": "10.0.0.2",
  "DefaultRaftPort": 10008,
  "RaftNodes": [
    "10.0.0.1",
    "10.0.0.2",
    "10.0.0.3"
  ],
```

as well as this:
```json
  "RaftEnabled": true,
  "RaftBind": "node-full-hostname-2.here.com",
  "DefaultRaftPort": 10008,
  "RaftNodes": [
    "node-full-hostname-1.here.com",
    "node-full-hostname-2.here.com",
    "node-full-hostname-3.here.com"
  ],
```

### Backend DB

A `raft` setup supports either `MySQL` or `SQLite` backend DB. See [backend](configuration-backend.md) configuration for either. Read [high-availability](high-availability.md) page for scenarios, possibilities and reasons to using either.
