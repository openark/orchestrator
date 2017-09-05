# orchestrator [[Documentation]](https://github.com/github/orchestrator/tree/master/docs)

[![build status](https://travis-ci.org/github/orchestrator.svg)](https://travis-ci.org/github/orchestrator) [![downloads](https://img.shields.io/github/downloads/github/orchestrator/total.svg)](https://github.com/github/orchestrator/releases) [![release](https://img.shields.io/github/release/github/orchestrator.svg)](https://github.com/github/orchestrator/releases)

![Orchestrator logo](https://github.com/github/orchestrator/raw/master/docs/images/orchestrator-logo-wide.png)

`orchestrator` is a MySQL high availability and replication management tool, allowing for:

#### Discovery

`orchestrator` actively crawls through your topologies and maps them. It reads basic MySQL info such as replication status and configuration.

It provides with slick visualization of your topologies, including replication problems, even in the face of failures.

#### Refactoring

`orchestrator` understands replication rules. It knows about binlog file:position, GTID, Pseudo GTID, Binlog Servers.

Refactoring replication topologies can be a matter of drag & drop a replica under another master. Moving replicas around is safe: `orchestrator` will reject an illegal refactoring attempt.

Fine-grained control is achieved by various command line options.

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

#### Additional perks

- Highly available
- Controlled master takeovers
- Manual failovers
- Failover auditing
- Audited operations
- Pseudo-GTID
- Datacenter/physical location awareness
- MySQL-Pool association
- HTTP security/authentication methods
- Coupled with [orchestrator-agent](https://github.com/github/orchestrator-agent), seed new/corrupt instances
- There is also an [orchestrator-mysql](https://groups.google.com/forum/#!forum/orchestrator-mysql) Google groups forum to discuss topics related to orchestrator
- More...

Read the [Orchestrator documentation](https://github.com/github/orchestrator/tree/master/docs)

Authored by [Shlomi Noach](https://github.com/shlomi-noach) at [GitHub](http://github.com). Previously at [Booking.com](http://booking.com) and [Outbrain](http://outbrain.com)

#### Related projects

- Orchestrator Puppet module: https://github.com/github/puppet-orchestrator-for-mysql
- Nagios / Icinga check based on Orchestrator API: https://github.com/mcrauwel/go-check-orchestrator

#### License

`orchestrator` is free and open sourced under the [Apache 2.0 license](LICENSE).
