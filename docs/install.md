# Installation

For production deployments, see [Orchestrator deployment](deployment.md). The following text walks you through the manual way of installation and the necessary configuration to make it work.

The following assumes you will be using the same machine for both the `orchestrator` binary and the MySQL backend.
If not, replace `127.0.0.1` with appropriate host name. Replace `orch_backend_password` with your own super secret password.

#### Extract orchestrator binary and files

- Extract from tarball

  Extract the archive you've downloaded from https://github.com/openark/orchestrator/releases
  For example, let's assume you wish to install `orchestrator` under `/usr/local/orchestrator`:

      sudo mkdir -p /usr/local
      sudo cd /usr/local
      sudo tar xzfv orchestrator-1.0.tar.gz

- Install from `RPM`

  Installs onto `/usr/local/orchestrator`. Execute:

      sudo rpm -i orchestrator-1.0-1.x86_64.rpm


- Install from `DEB`

  Installs onto `/usr/local/orchestrator`. Execute:

      sudo dpkg -i orchestrator_1.0_amd64.deb

- Install from repository

  `orchestrator` packages can be found in https://packagecloud.io/github/orchestrator


#### Setup backend MySQL server

Setup a MySQL server for backend, and invoke the following:

    CREATE DATABASE IF NOT EXISTS orchestrator;
    CREATE USER 'orchestrator'@'127.0.0.1' IDENTIFIED BY 'orch_backend_password';
    GRANT ALL PRIVILEGES ON `orchestrator`.* TO 'orchestrator'@'127.0.0.1';

`Orchestrator` uses a configuration file, located in either `/etc/orchestrator.conf.json` or relative path to binary `conf/orchestrator.conf.json` or
`orchestrator.conf.json`.

Tip: the installed package includes a file called `orchestrator.conf.json.sample` with some basic settings which you can use as baseline for `orchestrator.conf.json`. It is found in `/usr/local/orchestrator/orchestrator-sample.conf.json` and you may also find `/usr/local/orchestrator/orchestrator-sample-sqlite.conf.json` which has a SQLite-oriented configuration. Those sample files are also available [on the `orchestrator` repository](https://github.com/openark/orchestrator/tree/master/conf).

Edit `orchestrator.conf.json` to match the above as follows:

    ...
    "MySQLOrchestratorHost": "127.0.0.1",
    "MySQLOrchestratorPort": 3306,
    "MySQLOrchestratorDatabase": "orchestrator",
    "MySQLOrchestratorUser": "orchestrator",
    "MySQLOrchestratorPassword": "orch_backend_password",
    ...

#### Grant access to orchestrator on all your MySQL servers
For `orchestrator` to detect your replication topologies, it must also have an account on each and every topology. At this stage this has to be the
same account (same user, same password) for all topologies. On each of your masters, issue the following:

    CREATE USER 'orchestrator'@'orch_host' IDENTIFIED BY 'orch_topology_password';
    GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orchestrator'@'orch_host';
    GRANT SELECT ON mysql.slave_master_info TO 'orchestrator'@'orch_host';
    GRANT SELECT ON ndbinfo.processes TO 'orchestrator'@'orch_host'; -- Only for NDB Cluster

> `REPLICATION SLAVE` is required for `SHOW SLAVE HOSTS`, and for scanning binary logs in favor of [Pseudo GTID](#pseudo-gtid)
> `RELOAD` required for `RESET SLAVE` operation
> `PROCESS` required to see replica processes in `SHOW PROCESSLIST`
> On MySQL 5.6 and above, and if using `master_info_repository = 'TABLE'`, let orchestrator have access
> to the `mysql.slave_master_info` table. This will allow orchestrator to grab replication credentials if need be.

Replace `orch_host` with hostname or orchestrator machine (or do your wildcards thing). Choose your password wisely. Edit `orchestrator.conf.json` to match:

    "MySQLTopologyUser": "orchestrator",
    "MySQLTopologyPassword": "orch_topology_password",


Consider moving `conf/orchestrator.conf.json` to `/etc/orchestrator.conf.json` (both locations are valid)

To execute `orchestrator` in command line mode or in HTTP API only, all you need is the `orchestrator` binary.
To enjoy the rich web interface, including topology visualizations and drag-and-drop topology changes, you will need
the `resources` directory and all that is underneath it. If you're unsure, don't touch; things are already in place.
