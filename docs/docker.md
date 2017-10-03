# Docker

## Building the docker image
```
docker build -t orchestrator:latest .
```

## Running the docker image
The docker images exposes orcestrator on port 3000.

```
docker run -p 3000:3000 orchestrator
```

The folowing environment variables are available and take effect if no config
file si bind mounted into container at `/etc/orchestrator.conf.json`

* `ORC_TOPOLOGY_USER`: defaults to `orchestrator`
* `ORC_TOPOLOGY_PASSWORD`: defaults to `orchestrator`
* `ORC_DB_HOST`: defaults to `db`
* `ORC_DB_PORT`: defaults to `3306`
* `ORC_DB_NAME`: defaults to `orchestrator`
* `ORC_USER`: defaulsts to `orc_server_user`
* `ORC_PASSWORD`: defaults to `orc_server_password`

To set these variables you could add these to an environment-file where you add them like `key=value` (one pair per line). You can then pass this enviroment-file to the docker command adding `--env-file=path/to/env-file` to the `docker run` command.
