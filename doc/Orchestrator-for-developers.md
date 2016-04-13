_Orchestrator_ is open source and accepts pull requests.

If you would like to build _orchestrator_ on your own machine, or eventually submit PRs, follow this guide.

#### Requirements

_Orchestrator_ is built on Linux. OS/X should generally follow same guidelines. I have no hint about MS Windows, and the
build is incompatible with Windows.

#### Go setup

You will need to have a *Go 1.5* environment. *1.5* is required as of Dec 2015 due to [vendor directories](https://golang.org/cmd/go/#hdr-Vendor_Directories) - Go's package dependencies solution.

You will need to

	export GO15VENDOREXPERIMENT=1

This guide assumes you have set your Go environment, along with `GOPATH`.

To get started, issue

	go get github.com/outbrain/orchestrator

Change directory into `$GOPATH:/src/github.com/outbrain/orchestrator`

Issue the following to resolve all dependencies:

	go get ./...

Test that your code builds via

	go run go/cmd/orchestrator/main.go

#### DB setup

_Orchestrator_ requires a MySQL backend to run. This could be installed anywhere. I usually use [mysqlsandbox](http://mysqlsandbox.net/) for local installations. You may choose to just install mysql-server on your dev machine.

Once your backend MySQL setup is complete, issue:

    CREATE DATABASE IF NOT EXISTS orchestrator;
    GRANT ALL PRIVILEGES ON `orchestrator`.* TO 'orchestrator'@'127.0.0.1'
    IDENTIFIED BY 'orch_backend_password';

_Orchestrator_ uses a configuration file whose search path is either `/etc/orchestrator.conf.json`,  `conf/orchestrator.conf.json` or `orchestrator.conf.json`.
The repository includes a file called `conf/orchestrator.conf.json.sample` with some basic settings. Issue:

	cp conf/orchestrator.conf.json.sample conf/orchestrator.conf.json

The `conf/orchestrator.conf.json` file is not part of the repository and there is in fact a `.gitignore` entry for this file.

Edit `orchestrator.conf.json` to match the above as follows:

    ...
    "MySQLOrchestratorHost": "127.0.0.1",
    "MySQLOrchestratorPort": 3306,
    "MySQLOrchestratorDatabase": "orchestrator",
    "MySQLOrchestratorUser": "orchestrator",
    "MySQLOrchestratorPassword": "orch_backend_password",
    ...

Edit the above as as fit for your MySQL backend install.

#### Executing from dev environment

You should now be able to

	go run go/cmd/orchestrator/main.go http

This will also invoke initial setup of your database environment (creating necessary tables in the `orchestrator` schema).

Browse into `http://localhost:3000` or replace `localhoast` with your dev hostname. You should see the orchestrator (empty) clusters dashboard.

Now to make stuff interesting.

#### Grant access to orchestrator on all your MySQL servers
For _orchestrator_ to detect your replication topologies, it must also have an account on each and every topology. At this stage this has to be the
same account (same user, same password) for all topologies. On each of your masters, issue the following:

    GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.*
    TO 'orchestrator'@'orch_host' IDENTIFIED BY 'orch_topology_password';

> REPLICATION SLAVE is required if you intend to use [Pseudo GTID](#pseudo-gtid)

Replace `orch_host` with hostname or orchestrator machine (or do your wildcards thing). Choose your password wisely. Edit `orchestrator.conf.json` to match:

    "MySQLTopologyUser": "orchestrator",
    "MySQLTopologyPassword": "orch_topology_password",

#### Discovering MySQL instances

Go to the `Discovery` page at `http://localhost:3000/web/discover`. Type in a hostname & port for a known MySQL instance, preferably one that is part of a larger topology (again I like using _MySQLSandbox_ for such test environments). Submit it.

Depending on your configuration (`DiscoveryPollSeconds`, `InstancePollSeconds`) this may take a few seconds to a minute for
_orchestrator_ to fully scan the replication topology this instance belongs to, and present it under the [clusters dashboard](http://localhost:3000/web/clusters/).

If you've made it this far, you've done 90% of the work. You may consider configuring Pseudo GTID queries, DC awareness etc. See
"want to have" sub-sections under [configuration](Orchestrator-Manual#configuration).


#### Building

To build an _Orchestrator_ package, use the `build.sh` script:

	bash build.sh

You will need:

 - [fpm](https://github.com/jordansissel/fpm), which assumes you have `ruby` and `ruby-gems`
 - `rpmbuild`
 - `go`, `gofmt` in path
 - `tar`

 Current `build.sh` usage is:

 ```
 usage() {
  echo
  echo "Usage: $0 [-t target ] [-a arch ] [ -p prefix ] [-h] [-d]"
  echo "Options:"
  echo "-h Show this screen"
  echo "-t (linux|darwin) Target OS Default:(linux)"
  echo "-a (amd64|386) Arch Default:(amd64)"
  echo "-d debug output"
  echo "-p build prefix Default:(/usr/local)"
  echo
}
```

### Forking and Pull-Requesting

If you want to submit [pull-requests](https://help.github.com/articles/using-pull-requests/) you should first fork `http://github.com/outbrain/orchestrator`.

Setting up the environment is basically the same, except you don't want to

	go get github.com/outbrain/orchestrator

But instead clone your own repository.

Assume you fork onto `github.com/you-are-awesome/orchestrator`. _Golang_ has tight coupling between source code import paths and actual URIs. This leads to much confusion. Please consult [Forking Golang repositories on GitHub and managing the import path](http://code.openark.org/blog/development/forking-golang-repositories-on-github-and-managing-the-import-path) as for ways to solve
that coupling.

Very briefly, you will either want to:

	go get github.com/outbrain/orchestrator
	git remote add awesome-fork https://github.com/you-are-awesome/orchestrator.git

Or you will workaround as follows:

	cd $GOPATH
	mkdir -p {src,bin,pkg}
	mkdir -p src/github.com/outbrain/
	cd src/github.com/outbrain/
	git clone git@github.com:you-are-awesome/orchestrator.git # OR: git clone https://github.com/you-are-awesome/orchestrator.git
	cd orchestrator/
	go get ./...


You will have a fork of _orchestrator_ to which you can push your changes and from which you can send pull requests.
It is best that you first consult (use the [project issues](https://github.com/outbrain/orchestrator/issues)) whether some kind of development would indeed be merged.

You will need to license your code in [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0) or compatible.

Thank you for considering contributions to _orchestrator_!
