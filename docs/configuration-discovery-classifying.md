# Configuration: discovery, classifying servers

`orchestrator` will figure out the name of the cluster, data center, and more.

```json
{
  "ReplicationLagQuery": "select absolute_lag from meta.heartbeat_view",
  "DetectClusterAliasQuery": "select ifnull(max(cluster_name), '') as cluster_alias from meta.cluster where anchor=1",
  "DetectClusterDomainQuery": "select ifnull(max(cluster_domain), '') as cluster_domain from meta.cluster where anchor=1",
  "DataCenterPattern": "",
  "DetectDataCenterQuery": "select substring_index(substring_index(@@hostname, '-',3), '-', -1) as dc",
  "PhysicalEnvironmentPattern": "",
  "DetectSemiSyncEnforcedQuery": ""
}
```

### Replication lag

By default, `orchestrator` uses `SHOW SLAVE STATUS` and takes a 1-second granularity value for lag. However this lag does not take into account cascading lags in the event of chained replication. Many use custom heartbeat mechanisms such as `pt-heartbeat`. This provides with "absolute" lag from master, as well as sub-second resolution.

`ReplicationLagQuery` allows you to setup your own query.``~~~~``

### Cluster alias

At your company the different clusters have common names. "Main", "Analytics", "Shard031" etc. However the MySQL clusters themselves are unaware of such names.

`DetectClusterAliasQuery` is a query by which you will let `orchestrator` know the name of your cluster.

The name is important. You will likely use it to tell `orchestrator` stuff like: "please auto recover this cluster", or "what are all the participating instances in this cluster".

To make such data accessible by a query, one trick is to create a table in some `meta` schema:

```sql
CREATE TABLE IF NOT EXISTS cluster (
  anchor TINYINT NOT NULL,
  cluster_name VARCHAR(128) CHARSET ascii NOT NULL DEFAULT '',
  cluster_domain VARCHAR(128) CHARSET ascii NOT NULL DEFAULT '',
  PRIMARY KEY (anchor)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

... and populate it as follows (e.g. via puppet/cron):

```shell
mysql meta -e "INSERT INTO cluster (anchor, cluster_name, cluster_domain) \
  VALUES (1, '${cluster_name}', '${cluster_domain}') \
  ON DUPLICATE KEY UPDATE \     cluster_name=VALUES(cluster_name), cluster_domain=VALUES(cluster_domain)"
```

Perhaps your host naming conventions will disclose the cluster name and you only need a simple query on `@@hostname`.

### Data center

`orchestrator` is data-center aware. Not only will it color them nicely on the web interface; but it will take DC into consideration when running failovers.

You will configure data center awareness in one of two methods:

- `DataCenterPattern`: a regular expression to be used on the fqdn. e.g.: `"db-.*?-.*?[.](.*?)[.].myservice[.]com"`
- `DetectDataCenterQuery`: a query that returns the data center name

### Cluster domain

To a lesser importance, and mostly for visibility, `DetectClusterDomainQuery` should return the VIP or CNAME or otherwise the address of the cluster's master

### Semi-sync topology 

In some environments, it is important to control the not only the number of semi-sync replicas, but also if a replica is a semi-sync or an async replica. 
`orchestrator` can detect an undesired semi-sync configuration and toggle the semi-sync flags 
`rpl_semi_sync_slave_enabled` and `rpl_semi_sync_master_enabled` to correct the situation.

#### Semi-sync master (`rpl_semi_sync_master_enabled`)

`orchestrator` enables the semi-sync master flag during a master failover (e.g. `DeadMaster`) if `DetectSemiSyncEnforcedQuery` returns a value > 0
for the new master. `orchestrator` does not trigger any recoveries if the master flag is otherwise changed or incorrectly set.

A semi-sync master can enter two failure scenarios: [`LockedSemiSyncMaster`](failure-detection.md#lockedsemisyncmaster) and 
[`MasterWithTooManySemiSyncReplicas`](failure-detection.md#masterwithtoomanysemisyncreplicas). `orchestrator` disables the 
semi-sync master flag on semi-sync replicas during a recovery of either of these two conditions.

#### Semi-sync replicas (`rpl_semi_sync_slave_enabled`)

`orchestrator` can detect if there is an incorrect number of semi-sync replicas in the topology ([`LockedSemiSyncMaster`](failure-detection.md#lockedsemisyncmaster) and
[`MasterWithTooManySemiSyncReplicas`](failure-detection.md#masterwithtoomanysemisyncreplicas)), and can then correct the situation by enabling/disabling
the semi-sync replica flags accordingly.

This behavior can be controlled by the following options:

- `DetectSemiSyncEnforcedQuery`: query that returns the semi-sync priority (zero means async replica; higher number means higher priority)
- `EnforceExactSemiSyncReplicas`: flag that decides whether to enforce a _strict_ semi-sync replica topology. If enabled, the recovery of `LockedSemiSyncMaster` 
   and `MasterWithTooManyReplicas` will enable _and disable_ semi-sync on the replicas to match the desired topology exactly based on the priority order.
- `RecoverLockedSemiSyncMaster`: flag that decides whether to recover from a `LockedSemiSyncMaster` scenario. If enabled, the recovery of `LockedSemiSyncMaster`
  will enable _(but never disable)_ semi-sync on the replicas in the priority order to match the master wait count. This option has no effect if 
  `EnforceExactSemiSyncReplicas` is set. It is only useful if you'd like to only handle a situation which which there are too few semi-sync replicas, 
  but not if there are too many.
- `ReasonableLockedSemiSyncMasterSeconds`: number of seconds after which the `LockedSemiSyncMaster` condition is triggered; if not set, falls back to `ReasonableReplicationLagSeconds`

The priority order is defined by `DetectSemiSyncEnforcedQuery` (zero means async replica; higher number is higher priority), the promotion rule (`DetectPromotionRuleQuery`)
and the hostname (fallback). 

**Example 1**: Enforcing a strict semi-sync replica topology with two replicas and `rpl_semi_sync_master_wait_for_slave_count=1`:

```
  "DetectSemiSyncEnforcedQuery": "select priority from meta.semi_sync where cluster_member = @@hostname",
  "EnforceExactSemiSyncReplicas": true
```

Assuming this topology:

```
         ,- replica1 (priority = 10, rpl_semi_sync_slave_enabled = 1)
  master 
         `- replica2 (priority = 20, rpl_semi_sync_slave_enabled = 1)
```

`orchestrator` would detect a [`MasterWithTooManySemiSyncReplicas`](failure-detection.md#masterwithtoomanysemisyncreplicas) scenario
and disable semi-sync on replica2.

**Example 2**: Enforcing a weak semi-sync replica toplogy with two replicas and `rpl_semi_sync_master_wait_for_slave_count=1`:

```
  "DetectSemiSyncEnforcedQuery": "select 2586",
  "DetectPromotionRuleQuery": "select promotion_rule from meta.promotion_rules where cluster_member = @@hostname",
  "RecoverLockedSemiSyncMaster": true
```

Assuming this topology:

```
         ,- replica1 (priority = 2586, promotion rule = prefer, rpl_semi_sync_slave_enabled = 0)
  master 
         `- replica2 (priority = 2586, promotion rule = neutral, rpl_semi_sync_slave_enabled = 0)
```

`orchestrator` would detect a [`LockedSemiSyncMaster`](failure-detection.md#lockedsemisyncmaster) scenario
and enable semi-sync on replica1.
