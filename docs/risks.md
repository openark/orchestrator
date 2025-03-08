# Risks

Most of the time `orchestrator` only reads status from your topologies. Default configuration is to poll each instance once per minute.

`orchestrator` will connect to your topology servers, and will cap number of concurrent connections.

You may use `orchestrator` to refactor your topologies: move replicas around and change the replication tree. `orchestrator` will do its best to:

1. Make sure you only move an instance to a location where it is valid for it to replicate (e.g. that you don't put a 5.5 server below a 5.6 server)
2. Make sure you move an instance at the right time (ie the instance and whichever affected servers are not lagging badly, so that operation can completely in a timely manner).
3. Do the math correctly: stop the replica at the right time, roll it forward to the right position, `CHANGE MASTER` to the correct location & position.

The above is well tested.

You may use `orchestrator` to failover your topologies. You will be concerned that:

- `orchestrator` doesn't failover when there is no need.
- `orchestrator` does fail over when there is a need.
- A failover doesn't lose too many servers
- Failover ends with a consistent topology

Failovers are inherently tied to your deployments through hooks.

There's always risk with failovers. Make sure to test them.

Please make sure to read the [LICENSE](https://github.com/openark/orchestrator/blob/master/LICENSE), and especially the "WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND" part.
