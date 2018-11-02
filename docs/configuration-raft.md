# Configuration: raft

Set up a [orchestrator/raft](raft.md) cluster for high availability.

Assuming you will run `orchestrator/raft` on a `3` node setup, you will configure this on each node:

```json
  "RaftEnabled": true,
  "RaftDataDir": "<path.to.orchestrator.data.directory>",
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
- `RaftDataDir` must be set to a directory writable to `orchestrator.` `orchestrator` will attempt to create the directory if not exists.
- `RaftBind` must be set, use the IP address or full hostname of local host. This IP or hostname will also be listed as one of the `RaftNodes` variable.
- `DefaultRaftPort` can be set to any port, but must be consistent across all deployments.
- `RaftNodes` should list all nodes of the raft cluster. This list will consist of either IP addresses or host names, and will include the value of this host itself as presented in `RaftBind`.

As example, the following might be a working setup:

```json
  "RaftEnabled": true,
  "RaftDataDir": "/var/lib/orchestrator",
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
  "RaftDataDir": "/var/lib/orchestrator",
  "RaftBind": "node-full-hostname-2.here.com",
  "DefaultRaftPort": 10008,
  "RaftNodes": [
    "node-full-hostname-1.here.com",
    "node-full-hostname-2.here.com",
    "node-full-hostname-3.here.com"
  ],
```

### NAT, firewalls, routing

If your orchestrator/raft nodes need to communicate via NAT gateways, you can additionally set:

- `"RaftAdvertise": "<ip.or.fqdn.visible.to.other.nodes>"`

to IP or hostname which other nodes should contact. Otherwise other nodes would try to talk to the "RaftBind" address and fail.

Raft nodes will reverse proxy HTTP requests to the leader. `orchestrator` will attempt to heuristically compute the leader's URL to which redirect requests. If behind NAT, rerouting ports etc., `orchestrator` may not be able to compute that URL. You may configure:

- `"HTTPAdvertise": "scheme://hostname:port"`

to explicitly specify where a node, assuming it were the leader, would be accessed through HTTP API. As example, you would: `"HTTPAdvertise": "http://my.public.hostname:3000"`

### Backend DB

A `raft` setup supports either `MySQL` or `SQLite` backend DB. See [backend](configuration-backend.md) configuration for either. Read [high-availability](high-availability.md) page for scenarios, possibilities and reasons to using either.

### Single raft node setups

In production, you will want using multiple `raft` nodes, such as `3` or `5`.

In a testing environment you may run a `orchestrator/raft` setup composed of a single node. This node will implicitly be the leader, and will advertise raft messages to itself.

To run a single-node `orchestrator/raft`, configure an empty `RaftNodes`:

```json
  "RaftNodes": [],
```

or, alternatively, specify a single node, which is identical to `RaftBind` or `RaftAdvertise`:

```json
  "RaftEnabled": true,
  "RaftBind": "127.0.0.1",
  "DefaultRaftPort": 10008,
  "RaftNodes": [
    "127.0.0.1"
  ],
```
