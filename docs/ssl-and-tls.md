# SSL and TLS

Orchestrator supports SSL/TLS for the web interface as HTTPS.  This can be standard server side certificates
or you can configure Orchestrator to validate and filter client provided certificates with Mutual TLS.

Orchestrator also allows for the use of certificates to authenticate with MySQL.

If MySQL is using SSL encryption for replication, Orchestrator will attempt to configure replication with SSL during recovery.

#### HTTPS for the Web/API interface
You can set up SSL/TLS protection like so:

```json
{
    "UseSSL": true,
    "SSLPrivateKeyFile": "PATH_TO_CERT/orchestrator.key",
    "SSLCertFile": "PATH_TO_CERT/orchestrator.crt",
    "SSLCAFile": "PATH_TO_CERT/ca.pem",
}
```

The SSLCAFile is optional if you don't need to specify your certificate authority.  This will enable SSL via
the web interface (and API) so that communications are encrypted, like a normal HTTPS web page.

You can, similarly, set this up for the Agent API if you're using the `Orchestrator Agent` with:

```json
{
    "AgentsUseSSL": true,
    "AgentSSLPrivateKeyFile": "PATH_TO_CERT/orchestrator.key",
    "AgentSSLCertFile": "PATH_TO_CERT/orchestrator.crt",
    "AgentSSLCAFile": "PATH_TO_CERT/ca.pem",
}
```

This can be the same SSL certificate, but it doesn't have to be.

#### Mutual TLS
It also supports the concept of Mutual TLS.  That is, certificates that must be presented and valid for the
client as well as the server.  This is frequently used to protect service to service communication in an
internal network.  The certificates are commonly signed from an internal root certificate.

In this case the certificates must 1) be valid and 2) be for the correct service.  The correct service is dictated by
filtering on the Organizational Unit (OU) of the client certificate.

*Setting up a private root CA is not a trivial task.  It is beyond the scope of these documents to
instruct how to successfully accomplish it*

With that in mind, you can set up Mutual TLS by setting up SSL as above, but also add the following directives:

```json
{
    "UseMutualTLS": true,
    "SSLValidOUs": [ "service1", "service2" ],
}
```

This will turn on client certificate verification and start filtering clients based on their OU.  OU filtering is
mandatory as it's pointless to use Mutual TLS without it.  In this case, `service1` and `service2` would be able
to connect to Orchestrator assuming their certificate was valid and they had an OU with that exact service name.

#### MySQL Authentication
You can also use client certificates to authenticate, or just encrypt, you mysql connection.  You can encrypt the
connection to the MySQL server `Orchestrator` uses with:

```json
{
    "MySQLOrchestratorUseMutualTLS": true,
    "MySQLOrchestratorSSLSkipVerify": true,
    "MySQLOrchestratorSSLPrivateKeyFile": "PATH_TO_CERT/orchestrator-database.key",
    "MySQLOrchestratorSSLCertFile": "PATH_TO_CERT/orchestrator-database.crt",
    "MySQLOrchestratorSSLCAFile": "PATH_TO_CERT/ca.pem",
}
```

Similarly the connections to the topology databases can be encrypted with:

```json
{
    "MySQLTopologyUseMutualTLS": true,
    "MySQLTopologySSLSkipVerify": true,
    "MySQLTopologySSLPrivateKeyFile": "PATH_TO_CERT/orchestrator-database.key",
    "MySQLTopologySSLCertFile": "PATH_TO_CERT/orchestrator-database.crt",
    "MySQLTopologySSLCAFile": "PATH_TO_CERT/ca.pem",
}
```

In this case all of your topology servers must respond to the certificates provided.  There's no current
method to have TLS enabled only for some servers.

#### MySQL SSL Replication
If Orchestrator is able to configure the failed Source to replicate to the newly promoted Source during recovery, it will attempt to configure `Master_SSL=1` if the newly promoted Source was configured that way.

Orchestrator currently does not handle configuring Source SSL certificates for replication during recovery. 
