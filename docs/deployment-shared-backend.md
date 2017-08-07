# Orchestrator deployment: shared backend

This text describes deployments for shared backend DB. While we will depict synchronous replication backend this document applies to all backend DB setups. See [High availability](high-availability.md) for the various backend DB setups.

This follows general [deployment](deployment.md).

### Shared backend

You will need to set up a shared backend database. This could be synchronous replication (Galera/XtraDB Cluster/InnoDB Cluster) for high availability, or it could be a master-replicas setup etc.

The backend database has the _state_ of your topologies. `orchestrator` itself is almost stateless, and trusts the data in the backend database.

In a shared backend setup multiple `orchestrator` services will all speak to the same backend.

- For **synchronous replication**, the advise is:
  - Configure multi-writer mode (each node in the MySQL cluster is writable)
  - Have `1:1` mapping between `orchestrator` services and `MySQL` nodes: each `orchestrator` service to speak with its own node.
  ![orchestrator deployment, shared backend](images/orchestrator-deployment-shared-backend.png)
- For **master-replicas** (asynchronous & semi-synchronous), do:
  - Configure all `orchestrator` nodes to access the _same_ backend DB (the master)
  - Optionally you will have your own load balancer to direct traffic to said master, in which case configure all `orchestrator` nodes to access the proxy.

### MySQL backend setup and high availability

Setting up the backend DB is on you. Also, `orchestartor` doesn't eat its own dog food, and cannot recover a failure on its own backend DB.
You will need to handle, for example, the issue of adding a Galera node, or of managing your proxy health checks etc.

### What to deploy: service

- Deploy the `orchestrator` service onto service boxes.
  - In a synchronous replication shared backend setup, these may well be the very MySQL boxes, in a `1:1` mapping.
- Consider adding a proxy on top of the service boxes; the proxy would ideally redirect all traffic to the leader node. There is one and only one leader node, and the status check endpoint is `/api/leader-check`. It is OK to direct traffic to any healthy service. Since all `orchestrator` nodes speak to the same shared backend DB, it is OK to operate some actions from one service node, and other actions from another service nodes. Internal locks are placed to avoid running contradicting or interfering commands.


### What to deploy: client

To interact with orchestrator from shell/automation/scripts, you may choose to:

- Directly interact with the HTTP API
- use the [orchestrator-client](orchestrator-client.md) script.
  - Deploy `orchestrator-client` on any box from which you wish to interact with `orchestrator`.
  - Create and edit `/etc/profile.d/orchestrator-client.sh` on those boxes to read:
    ```
    ORCHESTRATOR_API="http://your.orchestrator.service.proxy:80/api"
    ```
    or
    ```
    ORCHESTRATOR_API="http://your.orchestrator.service.host1:80/api http://your.orchestrator.service.host2:80/api http://your.orchestrator.service.host3:80/api"
    ```
    In the latter case you will provide the list of all `orchestrator` nodes, and the `orchetsrator-client` script will automatically figure out which is the leader. With this setup your automation will not need a proxy (though you may still wish to use a proxy for web interface users).

    Make sure to chef/puppet/whatever the `ORCHESTRATOR_API` value such that it adapts to changes in your environment.

- The [orchestrator command line](executing-via-command-line.md).
  - Deploy the `orchestrator` binary (you may use the `orchestrator-cli` distributed package) on any box from which you wish to interact with `orchestrator`.
  - Create `/etc/orchestrator.conf.json` on those boxes, populate with credentials. This file should generally be the same as for the `orchestrator` service boxes. If you're unsure, use exact same file content.
  - The `orchestrator` binary will access the shared backend DB. Make sure to give it access. Typicaly this will be port `3306`.

It is OK to run `orchestrator` CLI even while the `orchestrator` service is operating, since they will all coordinate on the same backend DB.

###### Orchestrator service

- You _can_ and _should_ have multiple `orchestrator` services running on different machines, all running on the same backend database.
  - When running as a service, `orchestrator` repeatedly attempts to claim _leadership_. Should one `orchestrator` service
    fail, another one will pick up where it left.
  - Only one service will be the leader at any given time. If a leader steps down, another will take its place. Leader election is performed by means of interacting through the shared backend DB.
  - The leader is the one polling for servers; doing database cleanup; checking for crash recovery scenarios etc.
- You may choose to have all your `orchestrator` services load-balanced

###### Orchestrator CLI
- You _can_ and _should_ have as many deployments of orchestrator CLI as you like, on multiple servers, all configured to work
  with the _same backend MySQL_
- You _may_ concurrently issue commands from multiple CLIs, as well as working with the Web UI at the same time.
- The (single) MySQL backend has the necessary state for managing concurrent operations.
- `orchestrator` has "maintenance locks" which prevent destructive concurrent operations on the same instance. At worst an
  operation will be rejected due to not being able to acquire maintenance lock.

The author of `orchestrator` has it deployed on multiple machines
as a service, behind a load balancer. On the same setup, CLI is
deployed and can be executed from thousands of machines.

##### A visualized example

![Orchestrator deployment](images/orchestrator-deployment.png)

In the above four `orchestrator` services are behind an HTTP load balancer. Only one of them is the _leader_ at any given time; they compete for leadership between themselves and recognize if the leader is non doing its duty.

To the right: the many topologies polled by `orchestrator`. The leader polls each an every server in those topologies.

On top left: the `orchestrator` MySQL backend: a master & three replicas. All 4 services use the same backend database.

Not shown in this picture (for clarity purposes), but the `orchestrator` backend database and its replicas are themselves one of those topologies
polled by `orchestrator` It eats its own dogfood.

On left, bottom: `orchestrator` CLI is installed on any and all MySQL servers. All these CLI deployments use the very same
MySQL backend database.

Not shown in this picture (for clarity purposes), the MySQL servers on these hosts are being polled by `orchestrator` just as those
on the right.

You may choose to install `orchestrator` on non-MySQL hosts, of course. It has no MySQL dependencies on the deployed host.
