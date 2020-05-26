# Docker

Multiple Dockerfiles are available, to:

- Build and test `orchestrator`
- Create distribution files
- Run a minimal `orchestrator` daemon
- Run a full blown CI environment

[`script/dock`](https://github.com/openark/orchestrator/blob/master/script/dock) is a convenience script to build/spawn each of these docker images.

## Build and test

First, it should be noted that you can let GitHub Actions do all the work for you: it will build, put through testing, and generate an artifact, an `orchestrator` binary for Linux `amd64` for you, all from GitHub's platform. You do not strictly need Docker nor a development environment on your computer. See [CI](ci.md).

If you wish to build and test on your host, but do not want to set up a development environment, use:
```shell
$ script/dock test
```

This will use [`docker/Dockerfile.test`](https://github.com/openark/orchestrator/blob/master/docker/Dockerfile.test) to build, unit test, integration test, run doc validation on your behalf.

## Build and run

Run this command:
```shell
$ script/dock alpine
```
which uses [`docker/Dockerfile`](https://github.com/openark/orchestrator/blob/master/docker/Dockerfile) to build `orchestrator` on an Alpine Linux, and run the service. Docker will map port `:3000` onto your machine, you may browse onto `http://127.0.0.1:3000` to access the orchestrator web interface.

The following environment variables are available and take effect if no config
file is bind mounted into container at `/etc/orchestrator.conf.json`

* `ORC_TOPOLOGY_USER`: defaults to `orchestrator`
* `ORC_TOPOLOGY_PASSWORD`: defaults to `orchestrator`
* `ORC_DB_HOST`: defaults to `db`
* `ORC_DB_PORT`: defaults to `3306`
* `ORC_DB_NAME`: defaults to `orchestrator`
* `ORC_USER`: defaulsts to `orc_server_user`
* `ORC_PASSWORD`: defaults to `orc_server_password`

To set these variables you could add these to an environment file where you add them like `key=value` (one pair per line). You can then pass this enviroment file to the docker command adding `--env-file=path/to/env-file` to the `docker run` command.

## Create package files

Run this command:
```shell
$ script/dock pkg
```
To create (via [`fpm`](https://fpm.readthedocs.io/en/latest/)) release packages:
- `.deb`
- `.rpm`
- `.tgz`

for Linux `amd64`, with `Systemd` or `SysVinit`, all binaries or just client scripts. It uses the same methods as used for [official releases](https://github.com/openark/orchestrator/releases).

Uses [`Dockerfile.packaging`](https://github.com/openark/orchestrator/blob/master/docker/Dockerfile.packaging)

## Run full CI environment

Execute:
```
$ script/dock system
```

to run a full blown environment (see [ci-env.md](ci-env.md)), consisting of:
- MySQL replication topology (via [`dbdeployer`](https://www.dbdeployer.com/)) with heartbeat injection
- `orchestrator` as a service
- `HAProxy`
- `Consul`
- `consul-template`

All wired to work together. It's a good playground for testing `orchestrator`'s functionality.

Tips:

- port `13306` routes to current topology master
- MySQL topology available on ports `10111, 10112, 10113, 10114`
- Connect to MySQL with user: `ci`, password: `ci`. e.g.:
  `mysqladmin -uci -pci -h 127.0.0.1 --port 13306 processlist`
- Use `redeploy-ci-env` to re-create the MySQL topology, and recreate and restart the heartbeat, consul, consul-template and haproxy services. This resets the services to their original state.

Uses [`Dockerfile.system`](https://github.com/openark/orchestrator/blob/master/docker/Dockerfile.system)
