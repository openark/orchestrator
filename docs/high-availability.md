# Orchestrator High Availability

`Orchestrator` makes your MySQL topologies available, but what makes `orchestrator` highly available?

Before drilling down into the details, we should first observe that orchestrator is a service that runs with a MySQL backend. Thus, we need to substantiate the HA of both these components, as well as the continued correctness in the failover process of either of the two or of both.

### High Availability of the Orchestrator service

`Orchestrator` runs as a service and its configuration needs to
reference a MySQL backend. You can quite easily add more orchestrator
applications probably running on different hosts to provide redundancy.
These servers would have an identical configuration.  `Orchestrator`
uses the database to record the different applications which are
running and through it will allow an election process to choose one
of the processes to be the active node.  If that process fails the
remaining processes will notice and shortly afterwards choose a new
active node. The active node is the one which periodically checks
each of the MySQL servers being monitored to determine if they are
healthy.  If it detects a failure it will recover the topology if
so configured.

If you use the web interface to look at the topology information
or to relocate replicas within a cluster and you have multiple
orchestrator processes running then you need to use a load balancer
to provide the redundant web service through a common URL.  Requests
through the load balancer may not hit the active node but that is
not an issue as any of the running processes can serve the web
requests.

### High Availability of the Orchestrator backend database

At this time `Orchestrator` relies on a MySQL backend. The state of the clusters is persisted to tables and is queried via SQL. It is worth considering the following:

- The backend database is very small, and is linear with your number of servers. For most setups it's a matter of a few MB and well below 1GB, depending on history configuration, rate of polling etc. This easily allows for a fully in-memory database even on simplest machines.

- Write rate is dependent on the frequency the MySQL hosts are polled and the number of servers involved.  For most orchestrator installations the write rate is low.

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

- MySQL Group Replication
  This is similar to the MySQL Cluster (but uses the InnoDB engine) or the Galera/XtraDB cluster and is available in MySQL 5.7.17 (December 2016) and later.
  Similar considerations apply as for the previous two options.

- 2-node active-passive master-master configuration

> NOTE: there has been an initial discussion on supporting Consul/etcd as backend datastore; there is no pending work on that at this time.
