# About Orchestrator

`Orchestrator` is a MySQL replication topology HA, management and visualization tool, allowing for:

#### Discovery

`orchestrator` actively crawls through your topologies and maps them. It reads basic MySQL info such as replication status and configuration.

It provides with slick visualization of your topologies, including replication problems, even in the face of failures.

#### Refactoring

`orchestrator` understands replication rules. It knows about binlog file:position, GTID, Pseudo GTID, Binlog Servers.

Refactoring replication topologies can be a matter of drag & drop a replica under another master. Moving replicas around becomes
safe: `orchestrator` will reject an illegal refactoring attempt.

Find grained control is achieved by various command line options.

#### Recovery

`Orchestrator` uses a holistic approach to detect master and intermediate master failures. Based on information gained from the topology itself, it recognizes a variety of failure scenarios.

Configurable, it may choose to perform automated recovery (or allow the user to choose type of manual recovery). Intermediate master recovery achieved internally to `orchestrator`. Master failover supported by pre/post failure hooks.

Recovery process utilizes _orchestrator's_ understanding of the topology and of its ability to perform refactoring. It is based on _state_ as opposed to _configuration_: `orchestrator` picks the best recovery method by investigating/evaluating the topology at the time of
recovery itself.

![Orchestrator screenshot](images/orchestrator-simple.png)

#### Credits, attributions

Authored by [Shlomi Noach](https://github.com/shlomi-noach)

This project was originally initiated at [Outbrain](http://outbrain.com), who were kind enough to release it as open source from its very beginning. We wish to recognize Outbrain for their support of open source software and for their further willingness to collaborate to this particular project's success.

The project was later developed at [Booking.com](http://booking.com) and the company was gracious enough to release further changes into the open source.

At this time the project is being developed at [GitHub](http://github.com).
We will continue to keep it open and supported.

The project accepts pull-requests and encourages them. Thank you for any improvement/assistance!

Additional collaborators & contributors to this Wiki:

- [grierj](https://github.com/grierj)
- Other awesome people
