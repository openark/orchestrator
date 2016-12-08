orchestrator [[Documentation]](https://github.com/github/orchestrator/tree/master/docs)
============

![Orcehstrator logo](https://github.com/github/orchestrator/raw/readme-logo/docs/images/orchestrator-logo-wide.png)

`orchestrator` is a MySQL replication topology management and visualization tool, allowing for:

#### Discovery

`orchestrator` actively crawls through your topologies and maps them. It reads basic MySQL info such as replication status and configuration.

It provides with slick visualization of your topologies, including replication problems, even in the face of failures.

#### Refactoring

`orchestrator` understands replication rules. It knows about binlog file:position, GTID, Pseudo GTID, Binlog Servers.

Refactoring replication topologies can be a matter of drag & drop a replica under another master. Moving replicas around is safe: `orchestrator` will reject an illegal refactoring attempt.

Find grained control is achieved by various command line options.

#### Recovery

`orchestrator` uses a holistic approach to detect master and intermediate master failures. Based on information gained from the topology itself, it recognizes a variety of failure scenarios.

Configurable, it may choose to perform automated recovery (or allow the user to choose type of manual recovery). Intermediate master recovery achieved internally to `orchestrator`. Master failover supported by pre/post failure hooks.

Recovery process utilizes _orchestrator's_ understanding of the topology and of its ability to perform refactoring. It is based on _state_ as opposed to _configuration_: `orchestrator` picks the best recovery method by investigating/evaluating the topology at the time of
recovery itself.

#### The interface

`orchestrator` supports:

- Command line interface (love your debug messages, take control of automated scripting)
- Web API (HTTP GET access)
- Web interface, a _slick_ one.

![Orcehstrator screenshot](https://github.com/github/orchestrator/raw/master/docs/images/orchestrator-simple-topology.png)

#### More

- Auditing
- Supports Pseudo-GTID
- Datacenter/physical location awareness
- Maintenenace/downtime server states
- MySQL-Pool association
- Run as a service; orchestrator multi-service HA
- HTTP security/authentication methods
- Coupled with [orchestrator-agent](https://github.com/github/orchestrator-agent), seed new/corrupt instances
- More...

Read the [Orchestrator documentation](https://github.com/github/orchestrator/tree/master/docs)

Authored by [Shlomi Noach](https://github.com/shlomi-noach) at [GitHub](http://github.com). Previously at [Booking.com](http://booking.com) and [Outbrain](http://outbrain.com)

#### License

`orchestrator` is free and open sourced under the [Apache 2.0 license](LICENSE).
