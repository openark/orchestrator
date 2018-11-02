# Configuration: backend

Let orchestrator know where to find backend database. In this setup `orchestrator` will serve HTTP on port `:3000`.

```json
{
  "Debug": false,
  "ListenAddress": ":3000",
}
```

You may choose either a `MySQL` backend or a `SQLite` backend. See [High availability](high-availability.md) page for scenarios, possibilities and reasons to using either.

## MySQL backend

You will need to set up schema & credentials:

```json
{
  "MySQLOrchestratorHost": "orchestrator.backend.master.com",
  "MySQLOrchestratorPort": 3306,
  "MySQLOrchestratorDatabase": "orchestrator",
  "MySQLOrchestratorCredentialsConfigFile": "/etc/mysql/orchestrator-backend.cnf",
}
```

Notice `MySQLOrchestratorCredentialsConfigFile`. It will be of the form:
```
[client]
user=orchestrator_srv
password=${ORCHESTRATOR_PASSWORD}
```

where either `user` or `password` can be in plaintext or get their value from the environment.


Alternatively, you may choose to use plaintext credentials in the config file:

```json
{
  "MySQLOrchestratorUser": "orchestrator_srv",
  "MySQLOrchestratorPassword": "orc_server_password",
}
```

#### MySQL backend DB setup

For a MySQL backend DB, you will need to grant the necessary privileges:

```
CREATE USER 'orchestrator_srv'@'orc_host' IDENTIFIED BY 'orc_server_password';
GRANT ALL ON orchestrator.* TO 'orchestrator_srv'@'orc_host';
```

## SQLite backend

Default backend is `MySQL`. To setup `SQLite`, use:

```json
{
  "BackendDB": "sqlite",
  "SQLite3DataFile": "/var/lib/orchestrator/orchestrator.db",  
}
```

`SQLite` is embedded within `orchestrator`.

If the file indicated by `SQLite3DataFile` does not exist, `orchestrator` will create it. It will need write permissions on given path/file.
