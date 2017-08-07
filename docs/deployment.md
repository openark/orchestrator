# Orchestrator deployment in production

What does an `orchestrator` deployment look like? What do you need to set in `puppet/chef`? What services need to run and where?

### Deploying the service & clients

You will first decide whether you want to run `orchestrator` on a shared backend DB or with a `raft` setup. See [High availability](high-availability.md) for some options, and [orchestrator/raft vs. synchronous replication setup](raft-vs-sync-repl.md) for comparison & discussion.

Follow these deployment guides:

- Deploying `orchestrator` on [shared backend DB](deployment-shared-backend.md)
- Deploying `orchestrator` via [raft consensus](deployment-raft.md)

### Ongoing deployment hints and tips

##### Discovery

##### Promotion rule

##### Downtime

##### Pseudo-GTID
