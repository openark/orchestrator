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

Some servers are better candidate for promotion in the event of failovers. Some servers aren't good picks. Examples:

- A server has weaker hardware configuration. You prefer to not promote it.
- A server is in a remote data center and you don't want to promote it.
- A server is used as your backup source and has LVM snapshots open at all times. You don't want to promote it.
- A server has a good setup and is ideal as candidate. You prefer to promote it.
- A server is OK, and you don't have any particular opinion.

You will announce your preference for a given server to `orchestrator` in the following way:

```
orchestrator -c register-candidate -i ${::fqdn} --promotion-rule ${promotion_rule}
```

Supported promotion rules are:

- `prefer`
- `neutral`
- `prefer_not`
- `must_not`

Promotion rules expire after an hour. That's the dynamic nature of `orchestrator`. You will want to setup a cron job that will announce the promotion rule for a server:

```
*/2 * * * * root "/usr/bin/perl -le 'sleep rand 10' && /usr/bin/orchestrator-client -c register-candidate -i this.hostname.com --promotion-rule prefer"
```

This setup comes from production environments. The cron entries get updated by `puppet` to reflect the appropriate `promotion_rule`. A server may have `prefer` at this time, and `prefer_not` in 5 minutes from now. Integrate your own service discovery method, your own scripting, to provide with your up-to-date `promotion-rule`.

##### Downtime



##### Pseudo-GTID

##### Populating meta data
