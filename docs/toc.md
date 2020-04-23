# Table of Contents

#### Introduction
- [About](about.md)
- [License](license.md)
- [Download](download.md)
- [Requirements](requirements.md)

#### Setup
- [Installation](install.md): installing the service/binary
- [Configuration](configuration.md): breakdown of major configuration variables by topic.

#### Use
- [Execution](execution.md): running the `orchestrator` service.
- [Executing via command line](executing-via-command-line.md)
- [Using the web interface](using-the-web-interface.md)
- [Using the web API](using-the-web-api.md): achieving automation via HTTP GET requests
- [Using orchestrator-client](orchestrator-client.md): a no binary/config needed script that wraps API calls
- [Scripting samples](script-samples.md)

#### Deployment
- [High availability](high-availability.md): making `orchestrator` highly available
- [Deployment](deployment.md) instructions, hints and tips
- [Shared backend DB](deployment-shared-backend.md) deployment
- [orchestrator/raft](deployment-raft.md) deployment

#### Failure detection & recovery
- [Failure detection](failure-detection.md): how `orchestrator` detects failure, types of failures it can handle
- [Topology recovery](topology-recovery.md): recovery process, promotion and hooks.
- [Key-Value stores](kv.md): master discovery for your apps

#### Operation
- [Status Checks](status-checks.md)
- [Tags](tags.md)

#### Development
- [Understanding CI](ci.md)
- [Building and testing](build.md)
- [System test environment](ci-env.md)
- [Docker](docker.md)

#### Various
- [Security](security.md)
- [SSL and TLS](ssl-and-tls.md)
- [Pseudo GTID](pseudo-gtid.md): refactoring and high availability without using GTID.
- [Agents](agents.md)

#### Meta
- [Risks](risks.md)
- [Gotchas](gotchas.md)
- [Supported topologies and versions](supported-topologies-and-versions.md)
- [Bugs](bugs.md)
- [Contributions](contributions.md)
- [Who uses Orchestrator?](users.md)
- [Presentations](presentations.md)

#### Quick guides

- [FAQ](faq.md)
- [First steps](first-steps.md), a quick introduction to `orchestrator`
- Orchestrator for [developers](developers.md). Read this if you wish to develop/contribute to `orchestrator`.
