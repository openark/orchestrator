# CI tests

`orchestrator` uses [GitHub Actions](https://help.github.com/en/actions) to run the following CI tests:

- Build (main): build, unit tests, integration tests, docs tests
- Upgrade: test a successful upgrade path
- System: run system tests backed by actual MySQL topology

## Build

Running on pull requests, the [main CI](https://github.com/openark/orchestrator/blob/master/.github/workflows/main.yml) job validates the following:

- Validate source code is formatted using `gofmt`
- Build passes
- Unit tests pass
- Integration tests pass
  - Using `SQLite` backend
  - Using `MySQL` backend
- Documentation tests pass: ensuring pages and links are not orphaned.

The [action](https://github.com/openark/orchestrator/actions?query=workflow%3ACI) completes by producing an artifact: an `orchestrator` binary for Linux `amd64`. The artifact is kept for a couple months per GitHub Actions policy.

## Upgrade

[Upgrade](https://github.com/openark/orchestrator/blob/master/.github/workflows/upgrade.yml) runs on pull requests, tests a successful upgrade path from previous version (ie `master`) to PR's branch. This mainly tests internal database structure changes. The test:

- Checks out `master` and run `orchestrator`, once using `SQLite`, once using `MySQL`
- Checks out `HEAD` (PR's branch) and run `orchestrator` using pre-existing `SQLite` and `MySQL` backends. Expect no error.

## System

[System tests](https://github.com/openark/orchestrator/blob/master/.github/workflows/system.yml) run as a scheduled job. A system test:

- Sets up a [CI environment](https://github.com/openark/orchestrator-ci-env) which includes:
  - A replication topology via [DBDeployer](https://www.dbdeployer.com/), with heartbeat injection
  - [HAProxy](http://www.haproxy.org/)
  - [Consul](https://www.consul.io/)
  - [consul-template](https://github.com/hashicorp/consul-template)
- Deploys `orchestrator` as a service, `orchestrator-client`
- Runs a series of tests where `orchestrator` operates on the topology, e.g. refactors or fails over.
