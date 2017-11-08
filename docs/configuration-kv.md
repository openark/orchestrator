# Configuration: Key-Value stores

`orchestrator` supports these key-value stores:

- An internal store based on a relational table
- [Consul](https://github.com/hashicorp/consul)
- [ZooKeeper](https://zookeeper.apache.org/) - TO BE IMPLEMENTED

`orchestrator` supports master discovery by storing clusters' masters in KV.

```json
  "KVClusterMasterPrefix": "mysql/master",
  "ConsulAddress": "127.0.0.1:8500",
  "ZkAddress": "srv-a,srv-b:12181,srv-c",
```

`KVClusterMasterPrefix` is the prefix to use for master discovery entries. As example, your cluster alias is `mycluster` and the master host is `some.host-17.com` then you will expect an entry where:

- The Key is `mysql/master/mycluster`
- The Value is `some.host-17.com`

If specified, `ConsulAddress` indicates an address where a Consul HTTP service is available. If unspecified, no Consul access is attempted.

**UNIMPLEMENTED YET** If specified, `ZkAddress` indicates one or more ZooKeeper servers to connect to. Default port per server is `2181`. All the following are equivalent:

- `srv-a,srv-b:12181,srv-c`
- `srv-a,srv-b:12181,srv-c:2181`
- `srv-a:2181,srv-b:12181,srv-c:2181`
