# Scripting samples

This document presents scripting usage and ideas around `orchestrator`.

### show all clusters with aliases

```shell
$ orchestrator-client -c clusters-alias
mysql-9766.dc1.domain.net:3306,cl1
mysql-0909.dc1.domain.net:3306,olap
mysql-0246.dc1.domain.net:3306,mycluster
mysql-1111.dc1.domain.net:3306,oltp1
mysql-9002.dc1.domain.net:3306,oltp2
mysql-3972.dc1.domain.net:3306,oltp3
mysql-0019.dc1.domain.net:3306,oltp4
```
### Show only aliases

```shell
$ orchestrator-client -c clusters-alias | cut -d"," -f2 | sort
cl1
mycluster
olap
oltp1
oltp2
oltp3
oltp4
```

#### master of cluster

```shell
$ orchestrator-client -c which-cluster-master -alias mycluster
mysql-0246.dc1.domain.net:3306
```

#### All instances of cluster

```shell
$ orchestrator-client -c which-cluster-instances -alias mycluster
mysql-0246.dc1.domain.net:3306
mysql-1357.dc2.domain.net:3306
mysql-bb00.dc1.domain.net:3306
mysql-00ff.dc1.domain.net:3306
mysql-8181.dc2.domain.net:3306
mysql-2222.dc1.domain.net:3306
mysql-ecec.dc2.domain.net:3306
```

The above indicates what `orchestrator` knows about the replication graph. The list includes servers that may be offline/broken.

#### Shell loop over instances

```shell
$ orchestrator-client -c which-cluster-instances -alias mycluster | cut -d":" -f 1 | while read h ; do echo "Host is $h" ; done
Host is mysql-0246.dc1.domain.net
Host is mysql-1357.dc2.domain.net
Host is mysql-bb00.dc1.domain.net
Host is mysql-00ff.dc1.domain.net
Host is mysql-8181.dc2.domain.net
Host is mysql-2222.dc1.domain.net
Host is mysql-ecec.dc2.domain.net
```

#### disable semi sync on cluster

```shell
$ orchestrator-client -c which-cluster-instances -alias mycluster | while read i ; do
  orchestrator-client -c disable-semi-sync-master -i $i
done
mysql-0246.dc1.domain.net:3306
mysql-1357.dc2.domain.net:3306
mysql-bb00.dc1.domain.net:3306
mysql-00ff.dc1.domain.net:3306
mysql-8181.dc2.domain.net:3306
mysql-2222.dc1.domain.net:3306
mysql-ecec.dc2.domain.net:3306
```

#### enable semi sync on cluster master

```shell
$ orchestrator-client -c which-cluster-master -alias mycluster | while read i ; do
  orchestrator-client -c enable-semi-sync-master -i $i
done
mysql-0246.dc1.domain.net:3306
```

#### Let's try again. This time disable semi sync on all instances _except master_

```shell
$ master=$(orchestrator-client -c which-cluster-master -alias mycluster)
$ orchestrator-client -c which-cluster-instances -alias mycluster | grep -v $master | while read i ; do
  orchestrator-client -c disable-semi-sync-master -i $i
done
```

#### Likewise, set read-only on all replicas

```shell
$ orchestrator-client -c which-cluster-instances -alias mycluster | grep -v $master | while read i ; do
  orchestrator-client -c set-read-only -i $i
done
```

#### We don't really need to loop. We can use ccql

[ccql](https://github.com/github/ccql) is a concurrent, multi server MySQL client. It plays well with scripting in general and with `orchestrator` in particular.

```shell
$ orchestrator-client -c which-cluster-instances -alias mycluster | grep -v $master | ccql -C ~/.my.cnf -q "set global read_only=1"
```

#### Extract master hostname (no ":3306")

```shell
$ master_host=$(orchestrator-client -c which-cluster-master -alias mycluster | cut -d":" -f1)
$ echo $master_host
mysql-0246.dc1.domain.net
```

We will use the `master_host` variable following.

#### Using the API to show all data of a specific host

```shell
$ orchestrator-client -c api -path instance/$master_host/3306 | jq .
{
  "Key": {
    "Hostname": "mysql-0246.dc1.domain.net",
    "Port": 3306
  },
  "InstanceAlias": "",
  "Uptime": 12203,
  "ServerID": 65884260,
  "ServerUUID": "3e87bd92-2be0-13e8-ac1b-008cda544064",
  "Version": "5.7.18-log",
  "VersionComment": "MySQL Community Server (GPL)",
  "FlavorName": "MySQL",
  "ReadOnly": false,
  "Binlog_format": "ROW",
  "BinlogRowImage": "FULL",
  "LogBinEnabled": true,
  "LogReplicationUpdatesEnabled": true,
  "SelfBinlogCoordinates": {
    "LogFile": "mysql-bin.000002",
    "LogPos": 333006336,
    "Type": 0
  },
  "MasterKey": {
    "Hostname": "",
    "Port": 0
  },
  "IsDetachedMaster": false,
  "Slave_SQL_Running": false,
  "Slave_IO_Running": false,
  "HasReplicationFilters": false,
  "GTIDMode": "OFF",
  "SupportsOracleGTID": false,
  "UsingOracleGTID": false,
  "UsingMariaDBGTID": false,
  "UsingPseudoGTID": true,
  "ReadBinlogCoordinates": {
    "LogFile": "",
    "LogPos": 0,
    "Type": 0
  },
  "ExecBinlogCoordinates": {
    "LogFile": "",
    "LogPos": 0,
    "Type": 0
  },
  "IsDetached": false,
  "RelaylogCoordinates": {
    "LogFile": "",
    "LogPos": 0,
    "Type": 1
  },
  "LastSQLError": "",
  "LastIOError": "",
  "SecondsBehindMaster": {
    "Int64": 0,
    "Valid": false
  },
  "SQLDelay": 0,
  "ExecutedGtidSet": "",
  "GtidPurged": "",
  "ReplicationLagSeconds": {
    "Int64": 0,
    "Valid": true
  },
  "SlaveHosts": [
    {
      "Hostname": "mysql-2222.dc1.domain.net",
      "Port": 3306
    },
    {
      "Hostname": "mysql-00ff.dc1.domain.net",
      "Port": 3306
    },
    {
      "Hostname": "mysql-1357.dc2.domain.net",
      "Port": 3306
    }
  ],
  "ClusterName": "mysql-0246.dc1.domain.net:3306",
  "SuggestedClusterAlias": "mycluster",
  "DataCenter": "dc1",
  "PhysicalEnvironment": "",
  "ReplicationDepth": 0,
  "IsCoMaster": false,
  "HasReplicationCredentials": false,
  "ReplicationCredentialsAvailable": false,
  "SemiSyncEnforced": false,
  "SemiSyncMasterEnabled": true,
  "SemiSyncReplicaEnabled": false,
  "LastSeenTimestamp": "2018-03-21 04:40:38",
  "IsLastCheckValid": true,
  "IsUpToDate": true,
  "IsRecentlyChecked": true,
  "SecondsSinceLastSeen": {
    "Int64": 2,
    "Valid": true
  },
  "CountMySQLSnapshots": 0,
  "IsCandidate": false,
  "PromotionRule": "neutral",
  "IsDowntimed": false,
  "DowntimeReason": "",
  "DowntimeOwner": "",
  "DowntimeEndTimestamp": "",
  "ElapsedDowntime": 0,
  "UnresolvedHostname": "",
  "AllowTLS": false,
  "LastDiscoveryLatency": 7233416
}
```

#### Extract the hostname from the JSON:

```shell
$ orchestrator-client -c api -path instance/$master_host/3306 | jq .Key.Hostname -r
mysql-0246.dc1.domain.net
```

#### Extract master's hostname from the JSON:

```shell
$ orchestrator-client -c api -path instance/$master_host/3306 | jq .MasterKey.Hostname -r

(empty, this is the master)
```

#### Another way of listing all hostnames in a cluster: using API and jq

```shell
$ orchestrator-client -c api -path cluster/alias/mycluster | jq .[].Key.Hostname -r
mysql-0246.dc1.domain.net
mysql-1357.dc2.domain.net
mysql-bb00.dc1.domain.net
mysql-00ff.dc1.domain.net
mysql-8181.dc2.domain.net
mysql-2222.dc1.domain.net
mysql-ecec.dc2.domain.net
```

#### Show the master host for each member in the cluster:
```shell
$ orchestrator-client -c api -path cluster/alias/mycluster | jq .[].MasterKey.Hostname -r

mysql-0246.dc1.domain.net
mysql-00ff.dc1.domain.net
mysql-0246.dc1.domain.net
mysql-bb00.dc1.domain.net
mysql-0246.dc1.domain.net
mysql-bb00.dc1.domain.net
```

#### What is the master hostname of a specific instance?

```shell
$ orchestrator-client -c api -path instance/mysql-bb00.dc1.domain.net/3306 | jq .MasterKey.Hostname -r
mysql-00ff.dc1.domain.net
```

#### How many replicas to a specific instance?

```shell
$ orchestrator-client -c api -path instance/$master_host/3306 | jq '.SlaveHosts | length'
3
```

#### How many replicas to each of a cluster's members?

```shell
$ orchestrator-client -c api -path cluster/alias/mycluster | jq '.[].SlaveHosts | length'
3
0
2
1
0
0
0
```

#### Another way of listing all replicas

We filter out those that don't have output for `show slave status`:

```shell
$ orchestrator-client -c which-cluster-instances -alias mycluster | ccql -C ~/.my.cnf -q "show slave status" | awk '{print $1}'
mysql-00ff.dc1.domain.net:3306
mysql-bb00.dc1.domain.net:3306
mysql-2222.dc1.domain.net:3306
mysql-ecec.dc2.domain.net:3306
mysql-1357.dc2.domain.net:3306
mysql-8181.dc2.domain.net:3306
```

#### Followup, restart replication on all cluster's instances

```shell
$ orchestrator-client -c which-cluster-instances -alias mycluster | ccql -C ~/.my.cnf -q "show slave status" | awk '{print $1}' | ccql -C ~/.my.cnf -q "stop slave; start slave;"
```

#### I'd like to apply changes to replication, without changing the replica's state (if it's running, I want it to keep running. If it's not running, I don't want to start replication)

```shell
$ orchestrator-client -c restart-replica-statements -i mysql-bb00.dc1.domain.net -query "change master to auto_position=1" | jq .[] -r
stop slave io_thread;
stop slave sql_thread;
change master to auto_position=1;
start slave sql_thread;
start slave io_thread;
```

Compare with:

```shell
$ orchestrator-client -c stop-replica -i mysql-bb00.dc1.domain.net
mysql-bb00.dc1.domain.net:3306

$ orchestrator-client -c restart-replica-statements -i mysql-bb00.dc1.domain.net -query "change master to auto_position=1" | jq .[] -r
change master to auto_position=1;
```

The above just outputs statements, we need to push them back to the server:

```shell
orchestrator-client -c restart-replica-statements -i mysql-bb00.dc1.domain.net -query "change master to auto_position=1" | jq .[] -r | mysql -h mysql-bb00.dc1.domain.net
```

#### In which DC (data center) is a specific instance?

This and the next questions assume either `DetectDataCenterQuery` or `DataCenterPattern` are configured.

```shell
$ orchestrator-client -c api -path instance/mysql-bb00.dc1.domain.net/3306 | jq '.DataCenter'
dc1
```

#### In which DCs is a cluster deployed, and how many hosts in each DC?

```shell
$ orchestrator-client -c api -path cluster/mycluster | jq '.[].DataCenter' -r | sort | uniq -c
  4 dc1
  3 dc2
```

#### Which replicas are replicating cross DC?

```shell
$ orchestrator-client -c api -path cluster/mycluster |
    jq '.[] | select(.MasterKey.Hostname != "") |
        (.Key.Hostname + ":" + (.Key.Port | tostring) + " " + .DataCenter + " " + .MasterKey.Hostname + "/" + (.MasterKey.Port | tostring))' -r |
    while read h dc m ; do
      orchestrator-client -c api -path "instance/$m" | jq '.DataCenter' -r |
        { read master_dc ; [ "$master_dc" != "$dc" ] && echo $h ; } ;
    done

mysql-bb00.dc1.domain.net:3306
mysql-8181.dc2.domain.net:3306
```
