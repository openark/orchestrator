# Configuration sample file

The following is a production configuration file, with some details redacted.

```json
{
  "Debug": true,
  "EnableSyslog": false,
  "ListenAddress": ":3000",
  "MySQLTopologyCredentialsConfigFile": "/etc/mysql/orchestrator.cnf",
  "MySQLTopologySSLPrivateKeyFile": "",
  "MySQLTopologySSLCertFile": "",
  "MySQLTopologySSLCAFile": "",
  "MySQLTopologySSLSkipVerify": true,
  "MySQLTopologyUseMutualTLS": false,
  "MySQLTopologyMaxPoolConnections": 3,
  "MySQLOrchestratorHost": "127.0.0.1",
  "MySQLOrchestratorPort": 3306,
  "MySQLOrchestratorDatabase": "orchestrator",
  "MySQLOrchestratorCredentialsConfigFile": "/etc/mysql/orchestrator_srv.cnf",
  "MySQLOrchestratorSSLPrivateKeyFile": "",
  "MySQLOrchestratorSSLCertFile": "",
  "MySQLOrchestratorSSLCAFile": "",
  "MySQLOrchestratorSSLSkipVerify": true,
  "MySQLOrchestratorUseMutualTLS": false,
  "MySQLConnectTimeoutSeconds": 1,
  "DefaultInstancePort": 3306,
  "ReplicationLagQuery": "select round(absolute_lag) from meta.heartbeat_view",
  "SlaveStartPostWaitMilliseconds": 1000,
  "DiscoverByShowSlaveHosts": false,
  "InstancePollSeconds": 5,
  "DiscoveryIgnoreReplicaHostnameFilters": [
    "a_host_i_want_to_ignore[.]example[.]com",
    ".*[.]ignore_all_hosts_from_this_domain[.]example[.]com",
    "a_host_with_extra_port_i_want_to_ignore[.]example[.]com:3307"
  ],
  "ReadLongRunningQueries": false,
  "SkipMaxScaleCheck": true,
  "BinlogFileHistoryDays": 10,
  "UnseenInstanceForgetHours": 240,
  "SnapshotTopologiesIntervalHours": 0,
  "InstanceBulkOperationsWaitTimeoutSeconds": 10,
  "ActiveNodeExpireSeconds": 5,
  "HostnameResolveMethod": "default",
  "MySQLHostnameResolveMethod": "@@hostname",
  "SkipBinlogServerUnresolveCheck": true,
  "ExpiryHostnameResolvesMinutes": 60,
  "RejectHostnameResolvePattern": "",
  "ReasonableReplicationLagSeconds": 10,
  "ProblemIgnoreHostnameFilters": [

  ],
  "VerifyReplicationFilters": false,
  "MaintenanceOwner": "orchestrator",
  "ReasonableMaintenanceReplicationLagSeconds": 20,
  "MaintenanceExpireMinutes": 10,
  "MaintenancePurgeDays": 365,
  "CandidateInstanceExpireMinutes": 60,
  "AuditLogFile": "",
  "AuditToSyslog": false,
  "AuditPageSize": 20,
  "AuditPurgeDays": 365,
  "RemoveTextFromHostnameDisplay": ":3306",
  "ReadOnly": false,
  "AuthenticationMethod": "",
  "HTTPAuthUser": "",
  "HTTPAuthPassword": "",
  "AuthUserHeader": "",
  "PowerAuthUsers": [
    "*"
  ],
  "ClusterNameToAlias": {
    "127.0.0.1": "test suite"
  },
  "AccessTokenUseExpirySeconds": 60,
  "AccessTokenExpiryMinutes": 1440,
  "DetectClusterAliasQuery": "select ifnull(max(cluster_name), '') as cluster_alias from meta.cluster where anchor=1",
  "DetectClusterDomainQuery": "",
  "DataCenterPattern": "",
  "DetectDataCenterQuery": "select 'redacted'",
  "PhysicalEnvironmentPattern": "",
  "PromotionIgnoreHostnameFilters": [

  ],
  "ServeAgentsHttp": false,
  "UseSSL": false,
  "UseMutualTLS": false,
  "SSLSkipVerify": false,
  "SSLPrivateKeyFile": "",
  "SSLCertFile": "",
  "SSLCAFile": "",
  "SSLValidOUs": [

  ],
  "StatusEndpoint": "/api/status",
  "StatusSimpleHealth": true,
  "StatusOUVerify": false,
  "HttpTimeoutSeconds": 60,
  "StaleSeedFailMinutes": 60,
  "SeedAcceptableBytesDiff": 8192,
  "SeedWaitSecondsBeforeSend": 2,
  "PseudoGTIDPattern": "drop view if exists `meta`.`_pseudo_gtid_hint__asc:",
  "PseudoGTIDPatternIsFixedSubstring": true,
  "PseudoGTIDMonotonicHint": "asc:",
  "DetectPseudoGTIDQuery": "select count(*) as pseudo_gtid_exists from meta.pseudo_gtid_status where anchor = 1 and time_generated > now() - interval 1 day",
  "BinlogEventsChunkSize": 10000,
  "BufferBinlogEvents": true,
  "SkipBinlogEventsContaining": [
    "@@SESSION.GTID_NEXT= 'ANONYMOUS'"
  ],
  "ReduceReplicationAnalysisCount": false,
  "FailureDetectionPeriodBlockMinutes": 60,
  "RecoveryPeriodBlockSeconds": 600,
  "RecoveryIgnoreHostnameFilters": [

  ],
  "RecoverMasterClusterFilters": [
    "*"
  ],
  "RecoverIntermediateMasterClusterFilters": [
    "*"
  ],
  "OnFailureDetectionProcesses": [
    "/redacted/our-orchestrator-recovery-handler -t 'detection' -f '{failureType}' -h '{failedHost}' -C '{failureCluster}' -A '{failureClusterAlias}' -n '{countReplicas}'"
  ],
  "PreGracefulTakeoverProcesses": [
    "echo 'Planned takeover about to take place on {failureCluster}. Master will switch to read_only' >> /tmp/recovery.log"
  ],
  "PreFailoverProcesses": [
    "/redacted/our-orchestrator-recovery-handler -t 'pre-failover' -f '{failureType}' -h '{failedHost}' -C '{failureCluster}' -A '{failureClusterAlias}' -n '{countReplicas}'"
  ],
  "PostFailoverProcesses": [
    "/redacted/our-orchestrator-recovery-handler -t 'post-failover' -f '{failureType}' -h '{failedHost}' -H '{successorHost}' -C '{failureCluster}' -A '{failureClusterAlias}' -n '{countReplicas}' -u '{recoveryUID}'"
  ],
  "PostUnsuccessfulFailoverProcesses": [
    "/redacted/our-orchestrator-recovery-handler -t 'post-unsuccessful-failover' -f '{failureType}' -h '{failedHost}' -C '{failureCluster}' -A '{failureClusterAlias}' -n '{countReplicas}' -u '{recoveryUID}'"
  ],
  "PostMasterFailoverProcesses": [
    "/redacted/do-something # e.g. kick pt-heartbeat on promoted master"
  ],
  "PostIntermediateMasterFailoverProcesses": [
  ],
  "PostGracefulTakeoverProcesses": [
    "echo 'Planned takeover complete' >> /tmp/recovery.log"
  ],
  "CoMasterRecoveryMustPromoteOtherCoMaster": true,
  "DetachLostSlavesAfterMasterFailover": true,
  "ApplyMySQLPromotionAfterMasterFailover": true,
  "PreventCrossDataCenterMasterFailover": false,
  "PreventCrossRegionMasterFailover": false,
  "MasterFailoverLostInstancesDowntimeMinutes": 60,
  "PostponeReplicaRecoveryOnLagMinutes": 10,
  "OSCIgnoreHostnameFilters": [

  ],
  "GraphitePollSeconds": 60,
  "GraphiteAddr": "",
  "GraphitePath": "",
  "GraphiteConvertHostnameDotsToUnderscores": true,
  "BackendDB": "mysql",
  "MySQLTopologyReadTimeoutSeconds": 3,
  "MySQLDiscoveryReadTimeoutSeconds": 3,
  "SQLite3DataFile": "/var/lib/orchestrator/orchestrator-sqlite.db",
  "RaftEnabled": false,
  "RaftBind": "redacted",
  "RaftDataDir": "/var/lib/orchestrator",
  "DefaultRaftPort": 10008,
  "ConsulAddress": "redacted:8500",
  "RaftNodes": [
    "redacted",
    "redacted",
    "redacted"
  ]
}
```
