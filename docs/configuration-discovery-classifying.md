# Configuration: discovery, classifying servers

`orchestrator` will figure out the name of the cluster, data center, region, environment, and more.

```json
{
  "ReplicationLagQuery": "select absolute_lag from meta.heartbeat_view",
  "DetectClusterAliasQuery": "select ifnull(max(cluster_name), '') as cluster_alias from meta.cluster where anchor=1",
  "DetectClusterDomainQuery": "select ifnull(max(cluster_domain), '') as cluster_domain from meta.cluster where anchor=1",
  "DataCenterPattern": "",
  "DetectDataCenterQuery": "select substring_index(substring_index(@@hostname, '-',3), '-', -1) as dc",
  "PhysicalEnvironmentPattern": "",
}
```

### Replication lag

By default, `orchestrator` uses `SHOW SLAVE STATUS` and takes a 1-second granularity value for lag. However this lag does not take into account cascading lags in the event of chained replication. Many use custom heartbeat mechanisms such as `pt-heartbeat`. This provides with "absolute" lag from master, as well as sub-second resolution.

`ReplicationLagQuery` allows you to setup your own query.

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

You will configure data center awareness using one of three methods:

- `DataCenterPattern`: a regular expression to be used on the fqdn. e.g.: `"db-.*?-.*?[.](.*?)[.].myservice[.]com"`
- `DataCenterMap`: a mapping to be used on the hostname (with optional port), IP address (with optional port, or CIDR network. e.g.:
  ```json
  {
    "db-foo-bar.myservice.com": "dc1",
    "192.168.123.123:3307": "dc2",
    "192.168.0.0/24": "dc1",
    "192.168.1.0/24": "dc2",
    "192.168.3.0/24": "dc3"
  }
  ```
- `DetectDataCenterQuery`: a query that returns the data center name

### Region

`orchestrator` can maintain metadata about each server's region using the `RegionPattern`, `RegionMap`, and `DetectRegionQuery` configuration settings that correspond to those described under [Data center](#data-center) above.

### Environment

`orchestrator` can maintain metadata about each server's environment using the `PhysicalEnvironmentPattern`, `PhysicalEnvironmentMap`, and `DetectPhysicalEnvironmentQuery` configuration settings that correspond to those described under [Data center](#data-center) above.

### Cluster domain

To a lesser importance, and mostly for visibility, `DetectClusterDomainQuery` should return the VIP or CNAME or otherwise the address of the cluster's master
