## Orchestrator High Availability

_Orchestrator_ makes your MySQL topologies available, but what makes _orchestrator_ highly available?

Before drilling down into the details, we should first observe that orchestrator is a service that runs with a MySQL backend. Thus, we need to substantiate the HA of both these components, as well as the continued correctness in the failover process of either of the two or of both.

### High Availability of the Orchestrator service

### High Availability of the Orchestrator backend database

At this time _Orchestrator_ relies on a MySQL backend. The state of the clusters is persisted to tables and is queried via SQL. It is worth considering the following:

- The backend database is very small, and is linear with your number of servers. For most setups it's a matter of a few MB and well below 1GB, depending on history configuration, rate of polling etc. This easily allows for a fully in-memory database even on simplest machines.

- Write rate is low

To that extent you may use one of the following solutions in order to make the backend database highly available:

- 2-node MySQL Cluster
  This is a synchronous solution; anything you write on one node is guaranteed to exist on the second. Data is available and up to date even in the face of a death of one server.
  Suggestion: abstract via HAProxy with `first` load-balancing algorithm.
  > NOTE: right now table creation explicitly creates tables using InnoDB engine; you may `ALTER TABLE ... ENGINE=NDB`

- 3-node Galera/XtraDB cluster
  This is a synchronous solution; anything you write on one node is guaranteed to exist on both other servers.
  Galera is eventually consistent.
  Data is available and up to date even in the face of a death of one server.
  Suggestion: abstract via HAProxy with `first` load-balancing algorithm.

- 2-node active-passive master-master configuration

> NOTE: there has been an initial discussion on supporting Consul/etcd as backend datastore; there is no pending work on that at this time.
