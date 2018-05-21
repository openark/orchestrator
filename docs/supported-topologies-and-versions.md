# Supported Topologies and Versions

The following setups are supported by `orchestrator`:

- Plain-old MySQL replication; the _classic_ one, based on log file + position
- GTID replication. Both Oracle GTID and MariaDB GTID are supported.
- Statement based replication (SBR)
- Row based replication (RBR)
- Semi-sync replication
- Single master (aka standard) replication
- Master-Master (two node in circle) replication
- 5.7 Parallel replication
  - When using GTID there's no further constraints.
  - When using Pseudo-GTID in-order-replication must be enabled (see [slave_preserve_commit_order](http://dev.mysql.com/doc/refman/5.7/en/replication-options-slave.html#sysvar_slave_preserve_commit_order)).

The following setups are _unsupported_:

- Master-master...-master (circular) replication with 3 or more nodes in ring.
- 5.6 Parallel (thread per schema) replication
- Multi master replication (one replica replicating from multiple masters)
- Tungsten replicator


Also note:

Master-master (ring) replication is supported for two master nodes. Topologies of three master nodes or more in a ring are unsupported.

Galera/XtraDB Cluster replication is not strictly supported: `orchestrator` will not recognize that co-masters
in a Galera topology are related. Each such master would appear to `orchestrator` to be the head of its own distinct
topology.

Replication topologies with multiple MySQL instances on the same host are supported. For example, the testing
environment for `orchestrator` is composed of four instances all running on the same machine, courtesy MySQLSandbox.
However, MySQL's lack of information sharing between replicas and masters make it impossible for `orchestrator` to
analyze the topology top-to-bottom, since a master does not know which ports its replicas are listening on.
The default assumption is that replicas are listening on same port as their master. With multiple instances on a single
machine (and on same network) this is impossible. In such case you must configure your MySQL instances'
`report_host` and `report_port` ([read more](http://code.openark.org/blog/mysql/the-importance-of-report_host-report_port))
parameters, and set `orchestrator`'s configuration parameter `DiscoverByShowSlaveHosts` to `true`.
