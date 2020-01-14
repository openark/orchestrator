# Configuration: Key-Value stores

`orchestrator` supports these key-value stores:

- An internal store based on a relational table
- [Consul](https://github.com/hashicorp/consul)
- [ZooKeeper](https://zookeeper.apache.org/)

`orchestrator` supports master discovery by storing clusters' masters in KV.

```json
  "KVClusterMasterPrefix": "mysql/master",
  "ConsulAddress": "127.0.0.1:8500",
  "ZkAddress": "srv-a,srv-b:12181,srv-c",
  "ConsulCrossDataCenterDistribution": true,
  "RemoveForgottenClustersFromKV": false,
```

`KVClusterMasterPrefix` is the prefix to use for master discovery entries. As example, your cluster alias is `mycluster` and the master host is `some.host-17.com` then you will expect an entry where:

- The Key is `mysql/master/mycluster`
- The Value is `some.host-17.com:3306`

Note: on ZooKeeper the key will automatically prefix with a `/` if not already so.

#### Breakdown entries

In addition to the above, `orchestrator` also breaks down the master entries and adds the follows (illustrating via example above):

- `mysql/master/mycluster/hostname`, value is `some.host-17.com`
- `mysql/master/mycluster/port`, value is `3306`
- `mysql/master/mycluster/ipv4`, value is `192.168.0.1`
- `mysql/master/mycluster/ipv6`, value is `<whatever>`

The `/hostname`, `/port`, `/ipv4` and `/ipv6` extensions are automatically added for any master entry.

### Stores

If specified, `ConsulAddress` indicates an address where a Consul HTTP service is available. If unspecified, no Consul access is attempted.

If specified, `ZkAddress` indicates one or more ZooKeeper servers to connect to. Default port per server is `2181`. All the following are equivalent:

- `srv-a,srv-b:12181,srv-c`
- `srv-a,srv-b:12181,srv-c:2181`
- `srv-a:2181,srv-b:12181,srv-c:2181`

### Consul specific

See [kv](kv.md) documentation for Consul specific settings.
