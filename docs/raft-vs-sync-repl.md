# orchestrator/raft vs. synchronous replication setup

This compares deployment, behavioral, limitations and benefits of two high availability deployment approaches: `orchestrator/raft` vs `orchestrator/[galera|xtradb cluster|innodb cluster]`

We will assume and compare:

- `3` data-centers setup (an _availability zone_ may count as a data-center)
- `3` node `orchestrator/raft` setup
- `3` node `orchestrator` on multi-writer `galera|xtradb cluster|innodb cluster` (each MySQL in cluster may accept writes)
- A proxy able to run `HTTP` or `mysql` health checks
- `MySQL`, `MariaDB`, `Percona Server` all considered under the term `MySQL`.

![orchestrator HA via raft](images/orchestrator-ha-raft-vs-sync-repl.png)

| Compare | orchestrator/raft | synchronous replication backend |
| --- | --- | --- |
General wiring | Each `orchestrator` node has a private backend DB; `orchestrator` nodes communicate by `raft` protocol | Each `orchestrator` node connects to a different `MySQL` member in a synchronous replication group. `orchestrator` nodes do not communicate with each other.
Backend DB | `MySQL` or `SQLite` | `MySQL`
DB data | Independent across DB backends. May vary, but on a stable system converges to same overall picture | Single dataset, synchronously replicated across DB backends.
Leader and actions | Single leader. Only the leader runs recoveries. All nodes run discoveries (probing) and self-analysis | Single leader. Only the leader runs discoveries (probing), analysis and recoveries.
HTTP Access | Must only access the leader (should be enforced by proxy) | May access any healthy node (should be enforced by proxy). For read consistency always best to speak to leader only (can be enforced by proxy)

Proxy
Access
comamnd line
cross DC communication
Backend DB
probing nodes
HTTP API
Writes to backend DB


Here are considerations for choosing between the two approaches:

- You only have a single data center (DC): pick shared DB or even a [simpler setup](high-availability.md)
- You are comfortable with Galera/XtraDB Cluster/InnoDB Cluster and have the automation to set them up and maintain them: pick shared DB backend.
- You have high-latency cross DC network: choose `orchestrator/raft`.
- You don't want to allocate MySQL servers for the `orchestrator` backend: choose `orchestrator/raft` with `SQLite` backend
- You have thousands of MySQL boxes: choose either, but choose `MySQL` backend which is more write performant than `SQLite`.
