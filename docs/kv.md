# Key-Value stores

`orchestrator` supports these key value stores:

- An internal store based on a relational table
- [Consul](https://github.com/hashicorp/consul)
- [ZooKeeper](https://zookeeper.apache.org/) - **TO BE IMPLEMENTED**

See also [Key-Value configuration](configuration-kv.md).

### Key-Value usage

At this time Key-Value (aka KV) stores are used for:

- Master discoveries

### Master discoveries, key-values and failovers

The objective is that service discoveries such as `Consul` or `ZookKeeper`-based would be able serve master discovery and/or take action based on cluster's master identity and change.

The most common scenario is to update a Proxy to direct cluster's write traffic to a specific master. As an example, one may set up `HAProxy` via `consul-template`, such that `consul-template` populates a single-host master pool based on the key-value store authored by `orchestrator`.

`orchestrator` updates all KV stores upon master failover.

#### Populating master entries

Clusters' master entries are populated on:

- An actual failover: `orchestrator` overwrites existing entry with identity of new master
- A manual request for entry population

  For each cluster, you will want to make one manual request for entry population. KV stores have no initial knowledge of your setup. `orchestrator` does, but does not routinely update the stores. Use:

  - `orchestrator-client -c submit-masters-to-kv-stores` to submit all clusters' masters to KV, or
  - `orchestrator-client -c submit-masters-to-kv-stores -alias mycluster` to submit the master of `mycluster` to KV

    See [orchestrator-client.md](orchestrator-client) documentation. You may use the `orchestrator`
    command line invocation as well.

  Or you may directly accessing the API via:

  - `/api/submit-masters-to-kv-stores`, or
  - `/api/submit-masters-to-kv-stores/:alias`, or

  respectively.

### KV and raft setups
