# Configuration: basic discovery

Let orchestrator know how to query the MySQL topologies, what information to extract.

```json
{
  "MySQLTopologyCredentialsConfigFile": "/etc/mysql/orchestrator-topology.cnf",
  "InstancePollSeconds": 5,
  "DiscoverByShowSlaveHosts": false,
}
```

`MySQLTopologyCredentialsConfigFile` follows similar rules as `MySQLOrchestratorCredentialsConfigFile`. You may choose to use plaintext credentials:

```json
{
  "MySQLTopologyUser": "orchestrator",
  "MySQLTopologyPassword": "orc_topology_password",
}
```

`orchestrator` will probe each server once per `InstancePollSeconds` seconds.

On all your MySQL topologies, grant the following:

```
CREATE USER 'orchestrator'@'orc_host' IDENTIFIED BY 'orc_topology_password';
GRANT SUPER, PROCESS, REPLICATION SLAVE, REPLICATION CLIENT, RELOAD ON *.* TO 'orchestrator'@'orc_host';
GRANT SELECT ON meta.* TO 'orchestrator'@'orc_host';
GRANT SELECT ON ndbinfo.processes TO 'orchestrator'@'orc_host'; -- Only for NDB Cluster
```

You may optionally let `orchestrator` know what your topologies' replication credentials are. This assumes same credentials on all topologies. `orchestrator` will utilize those credentials when setting up co-masters or when gracefully promoting a new master.

```json
{
  "MySQLReplicationUser": "replication_user",
  "MySQLTopologyPassword": "some_password",
}
```

Likewise, you may use `MySQLReplicationCredentialsConfigFile` instead of the above two variables.
