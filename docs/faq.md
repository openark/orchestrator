# Frequently Asked Questions

### Who should use orchestrator?

DBAs and ops who have more than a mere single-master-single-replica replication topology.

### What can orchestrator do for me?

`Orchestrator` analyzes your replication topologies and provides information and actions: you will be able to visualize & manipulate these topologies (refactoring replication paths). Orchestrator can monitor and recover topology failures, including master death.

### Is this yet another monitoring tool?

No. `Orchestrator` is strictly _not_ a monitoring tool. There is no intention to make it so; no alerts or emails. It does provide with online visualization of your topology status though, and requires some thresholds of its own in order to manage the topology.

### What kind of replication does orchestrator support?

`Orchestrator` supports "plain-old-MySQL-replication", the one that uses binary log files and positions.
If you don't know what you're using, this is probably the one. 

### Does orchestrator support Row Based Replication?

Yes. Statement Based Replication and Row Based Replication are both supported (and the distinction
is in fact irrelevant to `orchestrator`)

### Does orchestrator support Semi Sync Replication?

Yes.

### Does orchestrator support Master-Master (ring) Replication?

Yes, for a ring of two masters (active-active, active-passive).

Master-Master-Master[-Master...] topologies, where the ring is composed of 3 or more masters are not supported and not tested.
And are discouraged. And are an abomination.

### Does orchestrator support Galera Replication?

Yes and no. `Orchestrator` is unaware of Galera replication. If you have three Galera masters and different replica topologies under each master,
then `orchestrator` sees these as three different topologies.

### Does orchestrator support GTID Replication?

Yes. Both Oracle GTID and MariaDB GTID are supported.

### Does orchestrator support 5.6 Parallel Replication (thread per schema)?

No. This is because `START SLAVE UNTIL` is not supported in parallel replication, and output of `SHOW SLAVE STATUS` is incomplete.
There is no expected work on this.

### Does orchestrator support 5.7 Parallel Replication?

Yes. When using GTID, you're all good.
When using Pseudo-GTID you must have in-order-replication is enabled (set [slave_preserve_commit_order](http://dev.mysql.com/doc/refman/5.7/en/replication-options-slave.html#sysvar_slave_preserve_commit_order)). 

### Does orchestrator support Multi-Master Replication?

No. Multi Master Replication (e.g. as in MariaDB 10.0) is not supported.

### Does orchestrator support Tungsten Replication?

No.

### Does orchestrator support MySQL Group Replication?

Partially. Replication groups in single primary mode are somewhat supported under MySQL 8.0. The extent of the support so far is:

* Orchestrator understands that all group members are part of the same cluster, retrieves replication group information
  as part of instance discovery, stores it in its database, and exposes it via the API.
* The orchestrator web UI displays single primary group members. They are shown like this:
    * All group secondary members as replicating from the primary.
    * All group members have an icon that shows they are group members (as opposed to traditional async/semi-sync replicas).
    * Hovering over the icon mentioned above provides information about the state and role of the DB instance in the
      group.
* Some relocation operations are forbidden for group members. In particular, orchestrator will refuse to relocate a secondary group member, as it, by definition, replicates always from the group primary. It will also reject an attempt to relocate a group primary under a secondary of the same group.

No support has been added (yet) to handling group member failure. If all you have is a single replication group, this is fine, because you don't need it; the group will handle all failures as long as it can secure a majority.

If, however, you have the primary of a group as a replica to another instance; or you have replicas under your group
members, know that this has not been tested and results are, therefore, unpredictable at the moment. It *might* work, but it might also create a singularity and suck your database under the event horizon.

### Does orchestrator support Yet Another Type of Replication?

No.

### Does orchestrator support...

No.

### Is orchestrator open source?

Yes. `Orchestrator` is released as open source under the Apache License 2.0 and is available at: https://github.com/openark/orchestrator

### Who develops orchestrator and why?

`Orchestrator` is developed by [Shlomi Noach](https://github.com/shlomi-noach) at [GitHub](http://github.com) (previously at [Booking.com](http://booking.com) and [Outbrain](http://outbrain.com)) to assist in managing multiple large replication topologies; time and human errors saved so far are almost priceless.

### Does orchestrator work with a cluster containing multiple major versions of MySQL?

Partially. This often arises when upgrading a cluster which can not
take downtime. Each replica will be taken offline and upgraded to
a new major version and added back to the cluster until all replicas
have been upgraded. Orchestrator is aware of MySQL versions and
will allow a replica with a higher major version to be moved under
a master or intermediate master of a lower version but not vice-versa
as this is generally not supported by the upstream vendors even if
it may actually work.  In most circumstances orchestrator will do
the right thing and it will allow for safe movement of such replicas
within the topology where this is possible. This has been used
extensively with MySQL 5.5/5.6 and also between with 5.6/5.7 but
not so much with MariaDB 10.  If you see issues which may be related
to this please report them.
