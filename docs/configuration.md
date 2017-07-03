# Configuration

Documenting and explaining all configuration variables turns to be a daunting task, as goes as deep as explain-the-code-in-words. There is an ongoing effort in pruning and simplifying the configuration.

The de-facto configuration list is located in [config.go](https://github.com/github/orchestrator/blob/master/go/config/config.go).

You are undoubtedly interested in configuring some basic components: the backend database, hosts discoveries. You may choose to use Pseudo-GTID. You may want `orchestrator` to notify upon failure, or you may wish to run full blown automated recovery.

Use the following small steps to configure `orchestrator`:

- 


#### Credentials: must have

```
"MySQLTopologyUser": "orchestrator",
"MySQLTopologyPassword": "orch_topology_password",
```

or, alternatively, use:
```
"MySQLTopologyCredentialsConfigFile": "/path/to/.my-orchestrator.cnf",
```

`/path/to/.my-orchestrator.cnf` format expected to be:

```
[client]
user=orchestrator
password=orch_topology_password
```

Also, must-have credentials for backend database:

```
"MySQLOrchestratorHost": "backend.orchestrator.db.mycompany.com",
"MySQLOrchestratorPort": 3306,
"MySQLOrchestratorDatabase": "orchestrator",
"MySQLOrchestratorUser": "orchestrator_server",
"MySQLOrchestratorPassword": "thepassword",
```

or, for user & password, use:
```
"MySQLOrchestratorCredentialsConfigFile": "/path/to/.my-orchestrator-srv.cnf",
```


#### Security: want to have

See [security](#security) section.


#### Better information: want to have

Use a heartbeat mechanism (as with [pt-heartbeat](http://www.percona.com/doc/percona-toolkit/2.1/pt-heartbeat.html)), and configure:
```
  "ReplicationLagQuery": "select replication_lag_seconds from heartbeat_table",
```

If you have multiple instances on same host, you must configure your MySQL servers with `report_host` and `report_port` and add:

```
  "DiscoverByShowSlaveHosts": true,
```

Audit operations to log file in addition to backend database table:
```
  "AuditLogFile": "/var/log/orchestrator-audit.log",
```

If your hostnames follow a strict convention, and you are able to detect data center from FQDN, or you are able to detect enviroment settings (prod? dev?) from FQDN, use (and modify):

```
  "DataCenterPattern": "[.]([^.]+)[.][^.]+[.]mycompany[.]com",
  "PhysicalEnvironmentPattern": "[.][^.]+[.]([^.]+)[.]mycompany[.]com",
```

`orchestrator` recognizes a cluster by its master's hostname & port. However you may also assign an alias to a cluster. This makes a couple CLI commadns simpler and some web pages nicer. If the alias can be queried via SQL, configure (and modify):

```
  "DetectClusterAliasQuery": "SELECT SUBSTRING_INDEX(@@hostname, '-', 1)",
```

Most your servers' FQDN are likely to end with `.mycomany.com:3306`. This wastes a lot of pixels on web pages. You may omit these via:
```
  "RemoveTextFromHostnameDisplay": ".mycompany.com:3306",
```


#### Pseudo GTID: want to have

Well, I'm suggesting you want to have _Pseudo GTID_. If you agree, then you _must_ inject Pseudo GTISD queries, and _must_ configure (and modify):

```
  "PseudoGTIDPattern": "drop view if exists .*?`_pseudo_gtid_hint__",
```

See [Pseudo GTID](#pseudo-gtid) discussion.

It is best if you can also query for Pseudo-GTID existence via SQL. For this, configure (and modify):

```
  "DetectPseudoGTIDQuery": "select count(*) as pseudo_gtid_exists from meta.pseudo_gtid_status where anchor = 1 and time_generated > now() - interval 2 day",
```

While you're at it, make your Pseudo-GTID entries monotonicly increasing, and provide a hint such as (modify value):

```
   "PseudoGTIDMonotonicHint": "asc:",
```

#### Topology recovery: want to have if you want to own your database

When PseudoGTID is enabled, `orchestrator` can do automated recovery from dead intermediate master (reconnects orphaned replicas to the topology)
or from dead masters (auto-promotes best candidate replica).

By default this is disabled. You can specify patterns of clusters for which to enable both. Of course, `.*` matches everything:

```
"RecoverMasterClusterFilters": [
      "myoltp"
],
"RecoverIntermediateMasterClusterFilters": [
      "myoltp",
      "myolap",
],
```

`orchestrator` recovers the topology structure, but as a generic tool it does not understand the ouer context of your MySQL topologies
management, such as DNS, proxies etc. It allows for hooks to invoke upon failover detection, before taking action and after taking action.
You might want to configure the following:

```
"OnFailureDetectionProcesses": [
  "echo 'Detected {failureType} on {failureCluster}. Affected replicas: {countSlaves}'
],
"PreFailoverProcesses": [
  "echo 'Will recover from {failureType} on {failureCluster}'
],
"PostFailoverProcesses": [
  "echo 'Recovered from {failureType} on {failureCluster}. Failed: {failedHost}:{failedPort}; Successor: {successorHost}:{successorPort}'
],
"PostUnsuccessfulFailoverProcesses": [
  "echo 'There was a problem recovering from {failureType} on {failureCluster}. Failed: {failedHost}:{failedPort}'
],
"PostMasterFailoverProcesses": [
  "echo 'Recovered from {failureType} on {failureCluster}. Failed: {failedHost}:{failedPort}; Promoted: {successorHost}:{successorPort}'
],
"PostIntermediateMasterFailoverProcesses": [
  "echo 'Recovered from {failureType} on {failureCluster}. Failed: {failedHost}:{failedPort}; Successor: {successorHost}:{successorPort}'
]
```
