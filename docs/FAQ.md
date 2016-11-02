### Who should use orchestrator?

DBAs and ops who have more than a mere single-master-single-slave replication topology.

### What can orchestrator do for me?

_Orchestrator_ analyzes your replication topologies. It can visualize those topologies, and it allows you to
move slaves around the topology easily and safely. It provides full audit to operations making for a
topology changelog. It can serve as a command line tool or it can provide with JSON API for all operations.

### Is this yet another monitoring tool?

No. _Orchestrator_ is strictly _not_ a monitoring tool. There is no intention to make it so; no alerts or emails. It does provide with online visualization of your topology status though, and requires some thresholds of its own in order to manage the topology.

### What kind of replication does orchestrator support?

_Orchestrator_ supports "plain-old-MySQL-replication", the one that uses binary log files and positions.
If you don't know what you're using, this is probably the one. It is the only type of replication up to and including MySQL 5.5.

### Does orchestrator support Row Based Replication?

Yes. Statement Based Replication and Row Based Replication are both supported (and the distinction
is in fact irrelevant to _orchestrator_)

### Does orchestrator support Semi Sync Replication?

Yes. And _orchestrator_ is in fact ignorant about it.

### Does orchestrator support Master-Master (ring) Replication?

Yes, for a ring of two masters (active-active, active-passive).

Master-Master-Master[-Master...] topologies, where the ring is composed of 3 or more masters are not supported and not tested.
And are discouraged. And are an abomination.

### Does orchestrator support Galera Replication?

Yes and no. _Orchestrator_ is unaware of Galera replication. If you have three Galera masters and different slave topologies under each master,
then _orchestrator_ sees these as three different topologies.

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

Yes. _Orchestrator_ is released as open source under the Apache 2.0 license and is available at: https://github.com/github/orchestrator

### Who develops orchestrator and why?

_Orchestrator_ is developed by [Shlomi Noach](https://github.com/shlomi-noach) at [GitHub](http://github.com) (previously at [Booking.com](http://booking.com) and [Outbrain](http://outbrain.com)) to assist in managing multiple large replication topologies; time and human errors saved so far are almost priceless.
