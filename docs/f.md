# Frequently Asked Questions

### Who should use orchestrator?

DBAs and ops who have more than a mere single-master-single-slave replication topology.

### What can orchestrator do for me?

`Orchestrator` analyzes your replication topologies and provides information and actions: you will be able to visualize & manipulate these topologies (refactoring replication paths). Orchestrator can monitor and recover topology failures, including master death.

### Is this yet another monitoring tool?

No. `Orchestrator` is strictly _not_ a monitoring tool. There is no intention to make it so; no alerts or emails. It does provide with online visualization of your topology status though, and requires some thresholds of its own in order to manage the topology.

### What kind of replication does orchestrator support?

`Orchestrator` supports "plain-old-MySQL-replication", the one that uses binary log files and positions.
If you don't know what you're using, this is probably the one. It is the only type of replication up to and including MySQL 5.5.

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

Yes and no. `Orchestrator` is unaware of Galera replication. If you have three Galera masters and different slave topologies under each master,
then `orchestrator` sees these as three different topologies.

### Does orchestrator support GTID Replication?

Yes. Both Oracle GTID and MariaDB GTID are supported.

### Does orchestrator support 5.6 Parallel Replication (thread per schema)?

No. This is because `START SLAVE UNTIL` is not supported in parallel replication, and output of `SHOW SLAVE STATUS` is incomplete.
There is no expected work on this.

### Does orchestrator support 5.7 Parallel Replication?

Yes, when in-order-replication is enabled (see [slave_preserve_commit_order](http://dev.mysql.com/doc/refman/5.7/en/replication-options-slave.html#sysvar_slave_preserve_commit_order)). The same applies to MariaDB.

### Does orchestrator support Multi-Master Replication?

No. Multi Master Replication (e.g. as in MariaDB 10.0) is not supported.

### Does orchestrator support Tungsten Replication?

No.

### Does orchestrator support Yet Another Type of Replication?

No.

### Does orchestrator support...

No.

### Is orchestrator open source?

Yes. `Orchestrator` is released as open source under the MIT license and is available at: https://github.com/github/orchestrator

### Who develops orchestrator and why?

`Orchestrator` is developed by [Shlomi Noach](https://github.com/shlomi-noach) at [GitHub](http://github.com) (previously at [Booking.com](http://booking.com) and [Outbrain](http://outbrain.com)) to assist in managing multiple large replication topologies; time and human errors saved so far are almost priceless.
