# Configuration

The following is an incomplete list of configuration parameters. "Complete" is always behind the latest code; you may also want to look at [config.go](https://github.com/github/orchestrator/blob/master/go/config/config.go)

* `Debug`                   (bool), set debug mode (similar to --debug option)
* `ListenAddress`           (string), host & port to listen on (default `":3000"`). You can limit connections to local machine via `"127.0.0.1:3000"`
* `MySQLTopologyUser`       (string), credentials for replication topology servers (masters & replicas)
* `MySQLTopologyPassword`   (string), credentials for replication topology servers (masters & replicas)
* `MySQLTopologyCredentialsConfigFile` (string), as an alternative to providing `MySQLTopologyUser`, `MySQLTopologyPassword`, name of file in `my.cnf`-like format where credentials are stored.
* `MySQLTopologyMaxPoolConnections` (int), Max concurrent connections on any topology instance
* `MySQLOrchestratorHost`   (string), hostname for backend MySQL server
* `MySQLOrchestratorPort`   (uint), port for backend MySQL server
* `MySQLOrchestratorDatabase`   (string), name of backend MySQL server schema
* `MySQLOrchestratorUser`       (string), credentials for backend MySQL server
* `MySQLOrchestratorPassword`   (string), credentials for backend MySQL server
* `MySQLOrchestratorCredentialsConfigFile`  (string), as an alternative to providing `MySQLOrchestratorUser`, `MySQLOrchestratorPassword`, name of file in `my.cnf`-like format where credentials are stored.
* `MySQLConnectTimeoutSeconds`  (int), Number of seconds before connection is aborted (driver-side)
* `MySQLReadTimeoutSeconds`     (int), Number of seconds to wait for a response from the MySQL server before connection is aborted (driver-side)
* `MySQLHostnameResolveMethod` (string), Method to resolve how to reach the MySQL instance. This is more powerful than `HostnameResolveMethod` and is ideal for complex setups like multiple instances on a host with a VIP per instance. Defaults to `none` but can be set to `@@report_host`
* `DefaultInstancePort` (int), In case port was not specified on command line (default value for this default is `3306`)
* `SkipOrchestratorDatabaseUpdate`  (bool), When false, orchestrator will attempt to create & update all tables in backend database; when true, this is skipped. It makes sense to skip on command-line invocations and to enable for http or occasional invocations, or just after upgrades
* `ReplicaLagQuery`               (string), custom query to check on replica lg (e.g. heartbeat table). If unprovided,
  replica's `Seconds_Behind_Master` is used.
* `SlaveStartPostWaitMilliseconds`  (int), Time to wait after `START SLAVE` before re-reading instance (give replica chance to connect to master)
* `DiscoverByShowSlaveHosts`    (bool), Attempt `SHOW SLAVE HOSTS` before `SHOW PROCESSLIST`
* `InstancePollSeconds`         (uint), Number of seconds between instance reads
* `UnseenInstanceForgetHours`   (uint), Number of hours after which an unseen instance is forgotten
* `DiscoveryPollSeconds`        (uint), Auto/continuous discovery of instances sleep time between polls
* `InstanceBulkOperationsWaitTimeoutSeconds`  (uint), Time to wait on a single instance when doing bulk (many instances) operation
* `ActiveNodeExpireSeconds` (uint), Maximum time to wait for active node to send keepalive before attempting to take over as active node.
* `HostnameResolveMethod`		(string), Type of hostname resolve method (either `"none"` or `"cname"`)
* `ExpiryHostnameResolvesMinutes`	(int), Number of minute after which a hostname resolve expires (hostname resolve are cached for up to this number of minutes)
* `RejectHostnameResolvePattern`  (string), Regexp pattern for resolved hostname that will not be accepted (not cached, not written to db). This is done to avoid storing wrong resolves due to network glitches.
* `ReasonableReplicationLagSeconds` (int), Above this value is considered a problem
* `VerifyReplicationFilters`  (bool), Include replication filters check before approving topology refactoring (e.g. `orchestrator` will not allow placing a non-filteres replica under a filtered one)
* `MaintenanceOwner`  (string), (Default) name of maintenance owner to use if none provided
* `ReasonableMaintenanceReplicationLagSeconds` (int), Above this value move-up and move-below are blocked
* `MaintenanceExpireMinutes`  (int), Minutes after which a maintenance flag is considered stale and is cleared
* `MaintenancePurgeDays`  (int), Days after which maintenance entries are purged from the database
* `AuditLogFile`  (string), Name of log file for audit operations. Disabled when empty.
* `AuditPageSize`       (int), Number of entries in an audit page
* `RemoveTextFromHostnameDisplay` (string), Text to strip off the hostname on cluster/clusters pages. Save pixels (e.g. `mycompany.com`)
* `ReadOnly`				(bool) When `"true"`, no write operations (e.g. stopping a replica, repointing replicas, discovering) are allowed
* `AuthenticationMethod`    (string), type of authentication. Either empty (no authentication, default), `"basic"`, `"multi"` or `"proxy"`. See [Security](#security) section.
* `AuthUserHeader`          (string), name of HTTP header which contains authenticated user when `AuthenticationMethod` is `"proxy"`
* `PowerAuthUsers`          (string list), users considered as *power users* (allowed to manipulate the topology); applies on `"proxy"` `AuthenticationMethod`.
* `HTTPAuthUser`        (string), Username for HTTP Basic authentication (blank disables authentication)
* `HTTPAuthPassword`    (string), Password for HTTP Basic authentication
* `ClusterNameToAlias`  (string-to-string map), Map between regex matching cluster name to a human friendly alias.
  The human friendly alias is then presented on the `Clusters` menu and in the `Clusters Dashboard` page.
* `DetectClusterAliasQuery` (string), Optional query (executed on topology instance) that returns the alias of a cluster. Query will only be executed on cluster master (though until the topology's master is resovled it may execute on other/all replicas). If provided, must return one row, one column. This overrides `ClusterNameToAlias`.
* `DataCenterPattern` (string), Regexp pattern with one group, extracting the datacenter name from the hostname
* `PhysicalEnvironmentPattern`  (string), Regexp pattern with one group, extracting physical environment info from hostname (e.g. combination of datacenter & prod/dev env)
* `DenyAutoPromotionHostnamePattern`  (string), Orchestrator will not auto-promote hosts with name matching patterb (via -c recovery; for example, avoid promoting dev-dedicated machines)
* `ServeAgentsHttp`     (bool), should *orchestrator* accept agent registrations and serve agent-related requests (see [Agents](#agents))
* `AgentsUseSSL` (bool), When `true` orchestrator will listen on agents port with SSL as well as connect to agents via SSL (see [SSL and TLS](#ssl-and-tls))
* `AgentsUseMutualTLS` (bool), When `true` Use mutual TLS for the server to agent communication
* `AgentSSLSkipVerify` (bool), When using SSL for the Agent, should we ignore SSL certification error
* `AgentSSLPrivateKeyFile` (string), Name of Agent SSL private key file, applies only when `AgentsUseSSL` = `true`
* `AgentSSLCertFile` (string), Name of Agent SSL certification file, applies only when `AgentsUseSSL` = `true`
* `AgentSSLCAFile` (string), Name of the Agent Certificate Authority file, applies only when `AgentsUseSSL` = `true`
* `AgentSSLValidOUs` ([]string), Valid organizational units when using mutual TLS to communicate with the agents
* `UseSSL` (bool), Use SSL on the server web port (see [SSL and TLS](#ssl-and-tls))
* `UseMutualTLS` (bool), When `true` Use mutual TLS for the server's web and API connections
* `SSLSkipVerify` (bool), When using SSL, should we ignore SSL certification error
* `SSLPrivateKeyFile` (string), Name of SSL private key file, applies only when `UseSSL` = `true`
* `SSLCertFile` (string), Name of SSL certification file, applies only when `UseSSL` = `true`
* `SSLCAFile` (string), Name of the Certificate Authority file, applies only when `UseSSL` = `true`
* `SSLValidOUs` ([]string), Valid organizational units when using mutual TLS
* `StatusEndpoint` (string), Override the status endpoint.  Defaults to `/api/status`
* `StatusSimpleHealth` (bool), If true, calling the status endpoint will use the simplified health check
* `StatusOUVerify` (bool), If true, try to verify OUs when Mutual TLS is on.  Defaults to false
* `HttpTimeoutSeconds`  (int),    HTTP GET request timeout (when connecting to _orchestrator-agent_)
* `AgentPollMinutes`     (uint), interval at which *orchestrator* contacts agents for brief status update
* `UnseenAgentForgetHours`     (uint), time without contact after which an agent is forgotten
* `StaleSeedFailMinutes`     (uint), time after which a seed with no state update is considered to be failed
* `PseudoGTIDPattern`   (string), Pattern to look for in binary logs that makes for a unique entry (pseudo GTID). When empty, Pseudo-GTID based refactoring is disabled.
* `PseudoGTIDMonotonicHint` (string), Optional, subtring in Pseudo-GTID entry which indicates Pseudo-GTID entries are expected to be monotonically increasing
* `DetectPseudoGTIDQuery` (string), Optional query which is used to authoritatively decide whether pseudo gtid is enabled on instance
* `BinlogEventsChunkSize` (int), Chunk size (X) for `SHOW BINLOG|RELAYLOG EVENTS LIMIT ?,X` statements. Smaller means less locking and more work to be done. Recommendation: keep `10000` or below, due to locking issues.
* `BufferBinlogEvents`  (bool), Should we used buffered read on `SHOW BINLOG|RELAYLOG EVENTS` -- releases the database lock sooner (recommended).
* `RecoveryPeriodBlockSeconds`  (int), The time for which an instance's recovery is kept "active", so as to avoid concurrent recoveries on smae instance as well as flapping
* `RecoveryIgnoreHostnameFilters` ([]string), Recovery analysis will completely ignore hosts matching given patterns
* `RecoverMasterClusterFilters` ([]string), Only do master recovery on clusters matching these regexp patterns (of course the ``.*`` pattern matches everything)
* `RecoverIntermediateMasterClusterFilters` ([]string), Only do intermediate-master recovery on clusters matching these regexp patterns (of course the ``.*`` pattern matches everything)

See [sample config file](https://github.com/github/orchestrator/blob/master/conf/orchestrator-simple.conf.json) in master branch.

#### Minimal config to work with

Most of the above configuration variables have good defaults, or may otherwise not be applicable to all use cases.
Here's a friendly breakdown of the stuff you _have_ to have and may _want_ to have.

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
