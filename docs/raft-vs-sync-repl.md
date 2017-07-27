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
Backend DB dependency | Service panics if cannot access its own private backend DB | Service _unhealthy_ if cannot access its own private backend DB
DB data | Independent across DB backends. May vary, but on a stable system converges to same overall picture | Single dataset, synchronously replicated across DB backends.
DB access | Never access directly. Only `raft` nodes access the backend DB while coordinating/cooperating. Or else inconsistencies can be introduced. | Possible to access directly; all `orchestrator` nodes/clients see exact same picture.
Leader and actions | Single leader. Only the leader runs recoveries. All nodes run discoveries (probing) and self-analysis | Single leader. Only the leader runs discoveries (probing), analysis and recoveries.
HTTP Access | Must only access the leader (should be enforced by proxy) | May access any healthy node (should be enforced by proxy). For read consistency always best to speak to leader only (can be enforced by proxy)
Command line | HTTP/API access (e.g. `curl`, `jq`) or `orchestrator-client` script which wraps common HTTP /API calls with familiar command line interface | HTTP/API, and/or `orchestrator-client` script, or `orchestrator ...` command line invocation.
Install | `orchestrator` service on service nodes only. `orchestrator-client` script wherever (requires access to HTTP/API). | `orchestrator` service on service nodes. `orchestrator-client` script wherever (requires access to HTTP/API). `orchestrator` client wherever (requires access to backend DBs)
HTTP Proxy | Must only direct traffic to the leader (`/api/leader-check`) | Must only direct traffic to healthy nodes (`/api/status`) ; best to only direct traffic to leader node (`/api/leader-check`)
Cross DC | Each `orchestrator` node (along with private backend) can run on a different DC. Nodes do not communicate much, low traffic. | Each `orchestrator` node (along with associated backend) can run on a different DC. `orchestrator` nodes do not communicate directly. `MySQL` group replication is chatty. Amount of traffic mostly linear by size of topologies and by polling rate. Write latencies.
Probing | Each topology server probed by all `orchestrator` nodes | Each topology server probed by single active node
Failure analysis | Performed independently by all nodes | Performed by leader only (DB is shared so all nodes see exact same picture anyhow)
Failover | Performed by leader node only | Performed by leader node only
Resiliency to failure | `1` node may go down (`2` on a `5` node cluster) | `1` node may go down (`2` on a `5` node cluster)
Node back from short failure | Node rejoins cluster, gets updated with changes. | DB node rejoins cluster, gets updated with changes.
Node back from long outage | DB must be cloned from healthy node. | Depends on your MySQL backend implementation. Potentially SST/restore from backup.


Here are considerations for choosing between the two approaches:

- You only have a single data center (DC): pick shared DB or even a [simpler setup](high-availability.md)
- You are comfortable with Galera/XtraDB Cluster/InnoDB Cluster and have the automation to set them up and maintain them: pick shared DB backend.
- You have high-latency cross DC network: choose `orchestrator/raft`.
- You don't want to allocate MySQL servers for the `orchestrator` backend: choose `orchestrator/raft` with `SQLite` backend
- You have thousands of MySQL boxes: choose either, but choose `MySQL` backend which is more write performant than `SQLite`.
