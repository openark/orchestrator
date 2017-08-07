# Orchestrator deployment in production

What does an `orchestrator` deployment look like? What do you need to set in `puppet/chef`? What services need to run and where?

### Deploying the service & clients

You will first decide whether you want to run `orchestrator` on a shared backend DB or with a `raft` setup. See [High availability](high-availability.md) for some options, and [orchestrator/raft vs. synchronous replication setup](raft-vs-sync-repl.md) for comparison & discussion.

Follow these deployment guides:

- Deploying `orchestrator` on [shared backend DB](deployment-shared-backend.md)
- Deploying `orchestrator` via [raft consensus](deployment-raft.md)

### Deployment hints and tips

`orchestrator` works well in dynamic environments and adapts to changes in inventory, configuration and topology. Its dynamic nature suggests that the environment should interact with it in dynamic nature, as well. Instead of hard-coded configuration, `orchestrator` is happy to accept dynamic hints and requests that change its perspective on the topologies.

##### Discovery

`orchestrator` auto-discovers boxes joining a topology. If a new replica joins an existing cluster that is monitored by `orchestrator`, it is discovered when `orchestrator` next probes its master.

However how does `orchestrator` discover completely new topologies?

- You may ask `orchestrator` to _discover_ (probe) any single server in such topology, and from there on it will crawl its way across the entire topology.
- Or you may choose to just let `orchestrator` know about any single production server you have, routinely. Set up a cronjob on any production `MySQL` server to read:

  ```
  0 0 * * * root "/usr/bin/perl -le 'sleep rand 600' && /usr/bin/orchestrator-client -c discover -i this.hostname.com"
  ```

  In the above, each host lets `orchestrator` know about itself once per day; newly bootstrapped hosts are discovered the next midnight. The `sleep` in introduced to avoid storming `orchestrator` by all servers at the same time.

  The above uses [orchestrator-client](orchestrator-client.md), but you may use the [orchestrator cli](executing-via-command-line.md) if running on a shared-backend setup.

##### Promotion rule

##### Downtime

##### Pseudo-GTID

##### Populating meta data
