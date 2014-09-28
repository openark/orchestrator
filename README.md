orchestrator
============

_Orchestrator_ is a MySQL replication topology management and visualization tool, allowing for:

* Detection and interrogation of replication clusters
* Safe topology refactoring: moving slaves around the topology
* Sleek topology visualization
* Replication problems visualization
* Topology changes via intuitive drag & drop
* Maintenance mode declaration and enforcement
* Reviewing and killing long running queries
* Auditing of operations
* When working with [orchestrator-agent](https://github.com/outbrain/orchestrator-agent), seed new/corrupt instances
* More...

Read the [Orchestrator Manual](https://github.com/outbrain/orchestrator/wiki/Orchestrator-Manual) for comprehensive documentation.

![Orcehstrator screenshot](https://github.com/outbrain/orchestrator/wiki/images/orchestrator-simple.png)

Refactoring the topology is a matter of a simple drag & drop. _Orchestrator_ will keep you safe by disallowing invalid replication topologies 
(e.g. replicating from ROW based to STATEMENT based, from 5.5 to 5.1 etc.)

_Orchestrator_ is developed at [Outbrain](http://www.outbrain.com/) to answer for the difficulty in managing replication topologies. 
At the time of _orchestrator_ creation GTID is available via MySQL 5.6 and MariaDB 10.0. And yet GTID is still not as mature as one would hope for. 
The majority of users still use plain-old binlog file:position based MySQL replication, and apparently this will hold for some time.

Authored by Shlomi Noach
 

