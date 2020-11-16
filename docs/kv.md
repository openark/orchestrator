# Key-Value stores

`orchestrator` supports these key value stores:

- An internal store based on a relational table
- [Consul](https://github.com/hashicorp/consul)
- [ZooKeeper](https://zookeeper.apache.org/)

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

- Encountering a new cluster, or encountering a master for which there is no existing KV entry. This check runs automatically and periodically.
  - The periodic check first consults with `orchestrator`'s internal KV store. It will only attempt to populate external stores (`Consul`, `Zookeeper`) if the internal store does not already have the master entries.
  It follows that the periodic checks will only inject external KV _once_.
- An actual failover: `orchestrator` overwrites existing entry with identity of new master
- A manual request for entry population:
  - `orchestrator-client -c submit-masters-to-kv-stores` to submit all clusters' masters to KV, or
  - `orchestrator-client -c submit-masters-to-kv-stores -alias mycluster` to submit the master of `mycluster` to KV

    See [orchestrator-client.md](orchestrator-client) documentation. You may use the `orchestrator`
    command line invocation as well.

  Or you may directly accessing the API via:

  - `/api/submit-masters-to-kv-stores`, or
  - `/api/submit-masters-to-kv-stores/:alias`, or

  respectively.

Both actual failover and manual request will override any existing KV entries, internal and external.

### KV and orchestrator/raft

On an [orchestrator/raft](raft.md) setup, all KV writes go through the `raft` protocol. Thus, once the leader determines a write needs to be made to KV stores, it publishes the request to all `raft` nodes. Each of the nodes will apply the write independently, based on its own configuration.

#### Implications

By way of example, let's say you are running `orchestrator/raft` in a 3-data-center setup, one node per DC.
Also, let's assume you have Consul set up on each of these DCs. Consul setups are typically inter-DC, with possibly cross-DC async replication.

Upon master failover, each and every `orchestrator` node will update Consul with new master's identity.

If your Consul runs cross-DC replication, then it is possible that the same KV update runs twice: once by means of Consul replication, once by the local `orchestrator` node. The two updates are identical and consistent, and are therefore safe to run.

If your Consul setups do not replicate from each other, `orchestrator` is the _only_ means by which your master discovery is made consistent across your Consul clusters. You get all the nice traits that come with `raft`: if one DC is network partitioned, the `orchestrator` node in that DC will not receive the KV update event, and for a while, neither will the Consul cluster. However, once network access is regained, `orchestartor` will catch up with event log and apply the KV update to the local Consul cluster. The setup is eventual-consistent.

Shortly following a master failover, `orchestrator` generates a `raft` snapshot. This isn't strictly required but is a useful operation: in the event the `orchestrator` node restarts, the snapshot prevents `orchestrator` from replaying the KV write. This is in particular interesting in an event of failover-and-failback, where a remote KV like consul might get two updates for the same cluster. The snapshot mitigates such incidents.

### Consul specific

Optionally, you may configure:

```json
  "ConsulCrossDataCenterDistribution": true,
```

...which can (and will) take place in addition to the flow illustrated above.

With `ConsulCrossDataCenterDistribution`, `orchestrator` runs an additional, periodic update to an extended list of Consul clusters.

Once per minute, `orchestrator` leader node queries its configured Consul server for the list of [known datacenters](https://www.consul.io/api/catalog.html#list-datacenters). It then iterates throught those data center clusters, and updates each and every one with the current identities of masters.

This functionality is required in case one has more Consul datacenters than just one-local-consul-per-orchestrator-node. We illustrated above how in a `orchestrator/raft` setup, each node updates its local Consul cluster. However, Consul clusters that are not local to any `orchestrator` node are unaffected by that approach. `ConsulCrossDataCenterDistribution` is the way to include all those other DCs.
