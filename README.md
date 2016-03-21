orchestrator [[Manual]](https://github.com/outbrain/orchestrator/wiki/Orchestrator-Manual)
============

_Orchestrator_ is a MySQL replication topology management and visualization tool, allowing for:

#### Discovery

_orchestrator_ actively crawls through your topologies and maps them. It reads basic MySQL info such as replication status and configuration.

It provides with slick visualization of your topologies, including replication problems, even in the face of failures.

#### Refactoring

_orchestrator_ understands replication rules. It knows about binlog file:position, GTID, Pseudo GTID, Binlog Servers.

Refactoring replication topologies can be a matter of drag & drop a replica under another master. Moving slaves around becomes
safe: _orchestrator_ will reject an illegal refactoring attempt.

Find grained control is achieved by various command line options.

#### Recovery

_Orchestrator_ uses a holistic approach to detect master and intermediate master failures. Based on information gained from
the topology itself, it recognizes a variety of failure scenarios.

Configurable, it may choose to perform automated recovery (or allow the user to choose type of manual recovery). Intermediate master
recovery achieved internally to _orchestrator_. Master failover supported by pre/post failure hooks.

Recovery process utilizes _orchestrator's_ understanding of the topology and of its ability to perform refactoring. It is based on _state_ as opposed to _configuration_: _orchestrator_ picks the best recovery method by investigating/evaluating the topology at the time of
recovery itself.


#### The interface

_Orchestrator_ supports:

- Command line interface (love your debug messages, take control of automated scripting)
- Web API (HTTP GET access)
- Web interface, a _slick_ one.

![Orcehstrator screenshot](https://raw.githubusercontent.com/wiki/outbrain/orchestrator/images/orchestrator-simple.png)

#### More

- Auditing
- Supports Pseudo-GTID
- Datacenter/physical location awareness
- Maintenenace/downtime server states
- MySQL-Pool association
- Run as a service; orchestrator multi-service HA
- HTTP security/authentication methods
- When working with [orchestrator-agent](https://github.com/outbrain/orchestrator-agent), seed new/corrupt instances
- More...

Read the [Orchestrator Manual](https://github.com/outbrain/orchestrator/wiki/Orchestrator-Manual) for comprehensive documentation.

Authored by [Shlomi Noach](https://github.com/shlomi-noach) at [GitHub](http://github.com). Previously at [Booking.com](http://booking.com) and [Outbrain](http://outbrain.com)
