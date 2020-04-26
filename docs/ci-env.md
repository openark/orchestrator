# CI environment

An ancillary project, [orchestrator-ci-env](https://github.com/openark/orchestrator-ci-env), provides a MySQL replication environment with which one may evaluate/test `orchestrator`. Use cases:

- You want to check `orchestrator`'s behavior on a testing environment.
- You want to test failover and master discovery.
- You want to develop changes to `orchestrator` and require a reproducible environment.

You may do all of the above if you already have some staging environment with a MySQL replication topology, or a [dbdeployer](https://www.dbdeployer.com/) setup, but `orchestrator-ci-env` offers a Docker-based environment, reproducible and dependency-free.

`orchestrator-ci-env` is the same environment used in [system CI](ci.md#system) and in [Dockerfile.system](docker.md#run-full-ci-environment).


# Setup

Clone `orchestrator-ci-env` via SSH or HTTPS:
```shell
$ git clone git@github.com:openark/orchestrator-ci-env.git
```
or
```shell
$ git clone https://github.com/openark/orchestrator-ci-env.git
```

# Run environment

Requirement: Docker.

```shell
$ cd orchestrator-ci-env
$ script/dock
```

This will build and run an environment which conists of:

- A replication topology via [DBDeployer](https://www.dbdeployer.com/), with heartbeat injection
- [HAProxy](http://www.haproxy.org/)
- [Consul](https://www.consul.io/)
- [consul-template](https://github.com/hashicorp/consul-template)

Docker will expose these ports:

- `10111`, `10112`, `10113`, `10114`: MySQL hosts in a replication topology. Initially, `10111` is the master.
- `13306`: exposed by HAProxy and routed to current MySQL topology master.

# Run orchestrator with environment

Assuming `orchestrator` is built into `bin/orchestrator` (`./script/build` if not):
```shell
$ bin/orchestrator --config=conf/orchestrator-ci-env.conf.json --debug http
```

[`conf/orchestrator-ci-env.conf.json`](https://github.com/openark/orchestrator/blob/master/conf/orchestrator-ci-env.conf.json) is designed to work with `orchestrator-ci-env`.

You may choose to change the value of `SQLite3DataFile`, which is by default on `/tmp`.

# Running system tests with environment

While `orchestrator` is running (see above), open another terminal in `orchestrator`'s repo path.

Run:
```shell
$ ./tests/system/test.sh
```
for all tests, or
```shell
$ ./tests/system/test.sh <name-or-regex>
```
for a specific test, e.g. `./tests/system/test.sh relocate-single`
