# Configuration: backend

Let orchestrator know where to find backend database. In this setup `orchestrator` will server HTTP on port `:3000`.

```json
{
  "Debug": false,
  "ListenAddress": ":3000",

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

### Backend DB

For a MySQL backend DB, you will need to grant the necessary privileges:

```
CREATE USER 'orchestrator_srv'@'orc_host' IDENTIFIED BY 'orc_server_password';
GRANT ALL ON orchestrator.* TO 'orchestrator_srv'@'orc_host';
```
