# Configuration: raft

Set up a [orchestrator/raft](raft.md) cluster for high availability.

Assuming you will run `orchestrator/raft` on a `3` node setup, you will configure this on each node:

```json
  "RaftEnabled": true,
  "RaftBind": "<ip.of.this.orchestrator.node>",
  "DefaultRaftPort": 10008,
  "RaftNodes": [
    "<ip.of.orchestrator.node1>",
    "<ip.of.orchestrator.node2>",
    "<ip.of.orchestrator.node3>",
  ],
```

Some breakdown:

- `RaftEnabled` must be set to `true`, otherwise `orchestrator` runs in shared-backend mode.
- `RaftBind` must be set, use the IP address of local host. This IP will also be listed as one of the `RaftNodes` variable.
- `DefaultRaftPort` can be set to any port, but must be consistent across all deployments.
- `RaftNodes` should list all nodes of the raft cluster. This list will consist of IP addresses (not host names) and will include the value of this host itself as presented in `RaftBind`.
