## Using the web API

A keen web developer would notice (via Firebug or Developer Tools) how the web interface
completely relies on JSON API requests.

The JSON API provides with all the maintenance functionality you can find in the web interface or the
command line mode.

> Most users will not be interested in accessing the API. If you're unsure: you don't need it.
> For creators of frameworks and maintenance tools, it may provide with great powers (and great responsibility).

The following is a brief listing of the web API exposed by `orchestrator`. Documentation tends to fall behind the code; see the
latest [API source code](https://github.com/github/orchestrator/blob/master/go/http/api.go) for the de-facto lsiting (scroll to end of file).

* `/api/instance/:host/:port`: reads and returns an instance's details (example `/api/instance/mysql10/3306`)
* `/api/discover/:host/:port`: discover given instance (a running `orchestrator` service will pick it up from there and
recursively scan the entire topology)
* `/api/refresh/:host/:port`: synchronously re-read instance status
* `/api/forget/:host/:port`: remove records of this instance. It may be automatically rediscovered by
  following up on its master or one of its replicas.
* `/api/resolve/:host/:port`: check if hostname resolves and whether TCP connection can be established (example: `/api/resolve/myhost.mydomain/3306`)  
* `/api/relocate/:host/:port/:belowHost/:belowPort` (attempt to) move an instance below another instance.
`Orchestrator` picks best course of action.
* `/api/relocate-replicas/:host/:port/:belowHost/:belowPort` (attempt to) move replicas of an instance below another instance.
`Orchestrator` picks best course of action.
* `/api/move-up/:host/:port` (attempt to) move this instance up the topology (make it child of its grandparent)
* `/api/move-below/:host/:port/:siblingHost/:siblingPort` (attempt to) move an instance below its sibling.
  the two provided instances must be siblings: replicas of the same master. (example `/api/move-below/mysql10/3306/mysql24/3306`)
* `/api/make-co-master/:host/:port` (attempt to) make this instance co-master with its own master, creating a
  circular master-master topology.
* `/api/reset-slave/:host/:port` reset a replica, breaking replication (destructive operation)
* `/api/begin-maintenance/:host/:port/:owner/:reason`: declares and begins maintenance mode for an instance.
  While in maintenance mode, `orchestrator` will not allow moving this instance.
  (example `/api/begin-maintenance/mysql10/3306/gromit/upgrading+mysql+version`)
* `/api/end-maintenance/:host/:port`: end maintenance on instance
* `/api/end-maintenance/:maintenanceKey` end maintenance on instance, based on maintenance key given on `begin-maintenance`
* `/api/start-slave/:host/:port`: issue a `START SLAVE` on an instance
* `/api/stop-slave/:host/:port`: issue a `STOP SLAVE` on an instance
* `/api/stop-slave-nice/:host/:port`: stop a replica such that the SQL thread is aligned with IO thread
* `/api/set-read-only/:host/:port`: issue a `SET GLOBAL read_only := 1` on an instance
* `/api/set-writeable/:host/:port`: issue a `SET GLOBAL read_only := 0` on an instance
* `/api/kill-query/:host/:port/:process`: kill a query (denoted by process id) on given instance. Synchronous call.
* `/api/maintenance`: list instances in active maintenance mode
* `/api/cluster/:clusterName`: list instances in a topology cluster. Each topology is automatically given a unique
  name. At this point the name is set by the topology's master (and if there's a master-master setup, then one of the masters).
  For example, a topology's name might be `mysql10:3306`, based on the understanding the server `mysql10` on port `3306`
  is the master of the topology.  
* `/api/clusters`: list names of known topologies.
* `/api/clusters-info`: list known clusters (topologies) and basic info
* `/api/cluster-pool-instances/:clusterName`: get pool information
* `/api/search/:searchString`: list instances matching search string
* `/api/problems`: list instances who have known problems (e.g. not replicating, lagging etc.)
* `/api/long-queries`: list of long running queries on all topologies (queries running for over 60 seconds, excluding replication and event-scheduler queries)
* `/api/long-queries/:filter`: list of long running queries on all topologies, filtered by text match
* `/api/audit`: show most recent audit entries
* `/api/audit/:page`: show latest audit entries, paginated (example: `/api/audit/3` for 3rd page)
* `/api/deregister-hostname-unresolve/:host/:port`:  unregister the given mapping for the given host
* `/api/register-hostname-unresolve/:host/:port/:virtualname`: register the host which should be used when unresolving the given virtual name

The following bulk retrieval rules are intended for allowing the information inside orchestrator to be used by external systems.
* `/api/bulk-instance`: provide a json list of instances in the form of Hostname Port
* `/api/bulk-promotion-rules`: provide a json list of instance promotion rules in the form of Hostname Port PromotionRule

The following urls allow access to the discovery metrics
* `/api/discovery-metrics-raw/:seconds`: list the raw discovery metrics for the last specified number of seconds. Raw discovery metric data is kept for DiscoveryCollectionRetentionSeconds seconds (default: 120)

#### Instance JSON breakdown

Many API calls return _instance objects_, describing a single MySQL server.
This sample is followed by a field breakdown:

```json
{

    "Key": {
        "Hostname": "mysql.02.instance.com",
        "Port": 3306
    },
    "Uptime": 45,
    "ServerID": 101,
    "Version": "5.6.22-log",
    "ReadOnly": false,
    "Binlog_format": "ROW",
    "LogBinEnabled": true,
    "LogSlaveUpdatesEnabled": true,
    "SelfBinlogCoordinates": {
        "LogFile": "mysql-bin.015656",
        "LogPos": 15082,
        "Type": 0
    },
    "MasterKey": {
        "Hostname": "mysql.01.instance.com",
        "Port": 3306
    },
    "Slave_SQL_Running": true,
    "Slave_IO_Running": true,
    "HasReplicationFilters": false,
    "SupportsOracleGTID": true,
    "UsingOracleGTID": true,
    "UsingMariaDBGTID": false,
    "UsingPseudoGTID": false,
    "ReadBinlogCoordinates": {
        "LogFile": "mysql-bin.015993",
        "LogPos": 20146,
        "Type": 0
    },
    "ExecBinlogCoordinates": {
        "LogFile": "mysql-bin.015993",
        "LogPos": 20146,
        "Type": 0
    },
    "RelaylogCoordinates": {
        "LogFile": "mysql_sandbox21088-relay-bin.000051",
        "LogPos": 16769,
        "Type": 1
    },
    "LastSQLError": "",
    "LastIOError": "",
    "SecondsBehindMaster": {
        "Int64": 0,
        "Valid": true
    },
    "SQLDelay": 0,
    "ExecutedGtidSet": "230ea8ea-81e3-11e4-972a-e25ec4bd140a:1-49",
    "SlaveLagSeconds": {
        "Int64": 0,
        "Valid": true
    },
    "SlaveHosts": [ ],
    "ClusterName": "mysql.01.instance.com:3306",
    "DataCenter": "",
    "PhysicalEnvironment": "",
    "ReplicationDepth": 1,
    "IsCoMaster": false,
    "IsLastCheckValid": true,
    "IsUpToDate": true,
    "IsRecentlyChecked": true,
    "SecondsSinceLastSeen": {
        "Int64": 9,
        "Valid": true
    },
    "CountMySQLSnapshots": 0,
    "IsCandidate": false,
    "UnresolvedHostname": ""
}
```

* `Key`: unique indicator for the instance: a combination of host & port
* `ServerID`: the MySQL `server_id` param
* `Version`: MySQL version
* `ReadOnly`: the global `read_only` boolean value
* `Binlog_format`: the global `binlog_format` MySQL param
* `LogBinEnabled`: whether binary logs are enabled
* `LogSlaveUpdatesEnabled`:  whether `log_slave_updates` MySQL param is enabled
* `SelfBinlogCoordinates`: binary log file & position this instance write to (as in `SHOW MASTER STATUS`)
* `MasterKey`: hostname & port of master, if any
* `Slave_SQL_Running`: direct mapping from `SHOW SLAVE STATUS`'s `Slave_SQL_Running`
* `Slave_IO_Running`: direct mapping from `SHOW SLAVE STATUS`'s `Slave_IO_Running`
* `HasReplicationFilters`: true if there's any replication filter
* `SupportsOracleGTID`: true if cnfigured with `gtid_mode` (Oracle MySQL >= 5.6)
* `UsingOracleGTID`: true if replica replicates via Oracle GTID
* `UsingMariaDBGTID`:  true if replica replicates via MariaDB GTID
* `UsingPseudoGTID`: true if replica known to have Pseudo-GTID coordinates (see related `DetectPseudoGTIDQuery` config)
* `ReadBinlogCoordinates`: (when replicating) the coordinates being read from the master (what `IO_THREAD` polls)
* `ExecBinlogCoordinates`: (when replicating) the master's coordinates that are being executed right now (what `SQL_THREAD` executed)
* `RelaylogCoordinates`: (when replicating) the relay log's coordinates that are being executed right now
* `LastSQLError`: copied from `SHOW SLAVE STATUS`
* `LastIOError`: copied from `SHOW SLAVE STATUS`
* `SecondsBehindMaster`: direct mapping from `SHOW SLAVE STATUS`'s `Seconds_Behind_Master`
    `"Valid": false` indicates a `NULL`
* `SQLDelay`: the configured `MASTER_DELAY`
* `ExecutedGtidSet`: if using Oracle GTID, the executed GTID set
* `SlaveLagSeconds`: when `SlaveLagQuery` provided, the computed replica lag; otherwise same as `SecondsBehindMaster`
* `SlaveHosts`: list of MySQL replicas _hostname & port_)
* `ClusterName`: name of cluster this instance is associated with; uniquely identifies cluster
* `DataCenter`: (metadata) name of data center, infered by `DataCenterPattern` config variable
* `PhysicalEnvironment`: (metadata) name of environment, infered by `PhysicalEnvironmentPattern` config variable
* `ReplicationDepth`: distance from the master (master is `0`, direct replica is `1` and so on)
* `IsCoMaster`: true when this instanceis part of a master-master pair
* `IsLastCheckValid`: whether last attempt at reading this instane succeeeded
* `IsUpToDate`: whether this data is up to date
* `IsRecentlyChecked`: whether a read attempt on this instance has been recently made
* `SecondsSinceLastSeen`: time elapsed since last successfully accessed this instance
* `CountMySQLSnapshots`: number of known snapshots (data provided by `orchestrator-agent`)
* `IsCandidate`: (metadata) `true` when this instance has been marked as _candidate_ via the `register-candidate` CLI command. Can be used in crash recovery for prioritizing failover options
* `UnresolvedHostname`: name this host _unresolves_ to, as indicated by the `register-hostname-unresolve` CLI command
