# Orchestrator for Developers

`orchestrator` is open source and accepts pull requests.

If you would like to build `orchestrator` on your own machine, or eventually submit PRs, follow this guide.

### Requirements

- `orchestrator` is built on Linux and OS/X. I have no hint about MS Windows, and the build is incompatible with Windows.

- `Go 1.9` or higher.

### Clone repo

`orchestrator` is built with `Go`. `Go` is picky about where you place your code (tl;dr: in a well structured path under `$GOPATH`). However the good news is we have the scripting to go around that (no pun intended).

### Easy clone + builds

Clone the repo anywhere on your filesystem via

	git clone git@github.com:openark/orchestrator.git
	cd orchestrator

Ensure you have proper build tools installed (CentOS referenced below)

    yum install gcc make perl-Digest-SHA

Then you can build `orchestrator` via

	script/build

You will find the binary as `bin/orchestrator`

This is the same script used by CI to build & test `orchestrator`.

### Not as easy clone + builds

Why would you want this? Because this will empower you with building `.DEB`, `.rpm` packages for both Linux and OS/X.

- Make sure `GOPATH` is set
- Issue:

	  go get github.com/github/orchestrator/...
	  cd $GOPATH:/src/github.com/github/orchestrator

- Compile or run via:

	  go build -i go/cmd/orchestrator/main.go
	  go run go/cmd/orchestrator/main.go

- Create packages via:

	  ./build.sh

	To create packages you will need to have:

	 - [fpm](https://github.com/jordansissel/fpm), which assumes you have `ruby` and `ruby-gems`
	 - `rpmbuild`
	 - `go`, `gofmt` in path
	 - `tar`

### DB setup

As per [configuration: backend](configuration-backend.md), you may use either `SQLite` or `MySQL` as `orchestrator`'s backend database.

#### Sqlite

No special setup is required. Make sure to configure database file path.

#### MySQL

`MySQL` can be installed anywhere and you don't necessarily need it to run on your local box. I usually use [mysqlsandbox](http://mysqlsandbox.net/) for local installations. You may choose to just install mysql-server on your dev machine, or run a docker container, a VM, what have you.

Once your backend MySQL setup is complete, issue:

    CREATE DATABASE IF NOT EXISTS orchestrator;
    CREATE USER 'orc_server_user'@'%' IDENTIFIED BY 'orc_server_password';
    GRANT ALL PRIVILEGES ON `orchestrator`.* TO 'orc_server_user'@'%';

`orchestrator` uses a configuration file whose search path is either `/etc/orchestrator.conf.json`,  `conf/orchestrator.conf.json` or `orchestrator.conf.json`.
The repository includes a file called `conf/orchestrator-sample.conf.json` with some basic settings. Issue:

	cp conf/orchestrator-sample.conf.json conf/orchestrator.conf.json

The `conf/orchestrator.conf.json` file is not part of the repository and there is in fact a `.gitignore` entry for this file.

Verify `orchestrator.conf.json` matches the above as follows:

    ...
    "MySQLOrchestratorHost": "127.0.0.1",
    "MySQLOrchestratorPort": 3306,
    "MySQLOrchestratorDatabase": "orchestrator",
    "MySQLOrchestratorUser": "orc_server_user",
    "MySQLOrchestratorPassword": "orc_server_password",
    ...

Edit the above as as fit for your MySQL backend install.

#### Executing from dev environment

You should now be able to:

	go run go/cmd/orchestrator/main.go http

or, if you used the easy clone + build process:

	bin/orchestrator http

This will also invoke initial setup of your database environment (creating necessary tables in the `orchestrator` schema).

Browse into `http://localhost:3000` or replace `localhoast` with your dev hostname. You should see the orchestrator (empty) clusters dashboard.

Now to make stuff interesting.

### Grant access to orchestrator on all your MySQL servers

For `orchestrator` to detect your replication topologies, it must also have an account on each and every topology. At this stage this has to be the
same account (same user, same password) for all topologies. On each of your masters, issue the following:

    CREATE USER 'orchestrator'@'orch_host' IDENTIFIED BY 'orch_topology_password';
    GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orchestrator'@'orch_host';

> REPLICATION SLAVE is required if you intend to use [Pseudo GTID](#pseudo-gtid)

Replace `orch_host` with hostname or orchestrator machine (or do your wildcards thing). Choose your password wisely. Edit `orchestrator.conf.json` to match:

    "MySQLTopologyUser": "orchestrator",
    "MySQLTopologyPassword": "orch_topology_password",

### Discovering MySQL instances

Go to the `Discovery` page at `http://localhost:3000/web/discover`. Type in a hostname & port for a known MySQL instance, preferably one that is part of a larger topology (again I like using _MySQLSandbox_ for such test environments). Submit it.

Depending on your configuration (`InstancePollSeconds`) this may take a few seconds to a minute for
`orchestrator` to fully scan the replication topology this instance belongs to, and present it under the [clusters dashboard](http://localhost:3000/web/clusters/).

If you've made it this far, you've done 90% of the work. You may consider configuring Pseudo GTID queries, DC awareness etc. See
"want to have" sub-sections under [configuration](Orchestrator-Manual#configuration).

### Customizations

There are some hooks in the Orchestrator web frontend which can be used to add customizations via CSS and JavaScript.

The corresponding files to edit are `resources/public/css/custom.css` and `resources/public/js/custom.js`.

You can find available hooks via `grep -r 'orchestrator:' resources/public/js`.

Please note that all APIs and structures are bound to change and any customizations are unsupported. Please file issues against uncustomized versions.

### Forking and Pull-Requesting

If you want to submit [pull-requests](https://help.github.com/articles/using-pull-requests/) you should first fork `http://github.com/github/orchestrator`.

Setting up the environment is basically the same, except you don't want to

	go get github.com/github/orchestrator/...

But instead clone your own repository.

Assume you fork onto `github.com/you-are-awesome/orchestrator`. _Golang_ has tight coupling between source code import paths and actual URIs. This leads to much confusion. Please consult [Forking Golang repositories on GitHub and managing the import path](http://code.openark.org/blog/development/forking-golang-repositories-on-github-and-managing-the-import-path) as for ways to solve
that coupling.

Very briefly, you will either want to:

	go get github.com/github/orchestrator/...
	git remote add awesome-fork https://github.com/you-are-awesome/orchestrator.git

Or you will workaround as follows:

	cd $GOPATH
	mkdir -p {src,bin,pkg}
	mkdir -p src/github.com/github/
	cd src/github.com/github/
	git clone git@github.com:you-are-awesome/orchestrator.git # OR: git clone https://github.com/you-are-awesome/orchestrator.git
	cd orchestrator/


You will have a fork of `orchestrator` to which you can push your changes and from which you can send pull requests.
It is best that you first consult (use the [project issues](https://github.com/github/orchestrator/issues)) whether some kind of development would indeed be merged.

You will need to license your code in [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0) or compatible.

Thank you for considering contributions to `orchestrator`!
