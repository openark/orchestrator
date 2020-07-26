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
GRANT SELECT ON performance_schema.replication_group_members TO 'orchestrator'@'orc_host'; -- Only for Group Replication / InnoDB cluster
```
