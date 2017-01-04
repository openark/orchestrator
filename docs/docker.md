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

* `ORC_ADDRESS`: set the `ListenAddress`, defaults to `:3000`
* `ORC_TOPOLOGY_USER`: defaults to `orcestrator`
* `ORC_TOPOLOGY_PASSOWRD`: defaults to `orchestrator`
* `ORC_DB_HOST`: defaults to `db`
* `ORC_DB_PORT`: defaults to `3306`
* `ORC_DB_NAME`: defaults to `orchestrator`
* `ORC_USER`: defaulsts to `orc_server_user`
* `ORC_PASSWORD`: defaults to `orc_server_password`
