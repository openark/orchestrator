Orchestrator Vagrant Instructions
=================================

Orchestrator's Vagrant defaults to installing five (5) CentOS 6.6 boxes in the following replication topology, with a separate "admin" node:

```
db1<->db2
 |     |
 v     v
db3   db4 
```

It is possible to override what gets installed by the use of environmental variables:

```
VAGRANT_SERVER_URL defaults to 'https://atlas.hashicorp.com'
VAGRANT_DEFAULT_PROVIDER defaults to 'virtualbox'
VAGRANT_BOX defaults to 'nrel/CentOS-6.6-x86_64
```

The MySQL configuration is such that it is the minimum required to set up replication. For testing such features as delayed replication, RBR/SBR, GTID, it is simply a matter of editing the `vagrant/dbX-my.cnf` before running the `vagrant up` command.

FAQ
===

Q: By default, there are still a lot of steps that I have to do within each virtual machine to get going

A: That is by design.  Vagrant will execute `db-post-install.sh`, `dbX-post-install.sh`, and `admin-post-install.sh` in the `vagrant/` directory (they are `.gitignore`'d) for any custom work that you want to have done (i.e. build Orchestrator, etc etc)

Q: I run some other distribution of Linux.  Why don't you support that?

A: Pull Requests are welcome!  If you update any of the `vagrant/*.sh` scripts, they must work with at least CentOS 6 and Ubuntu 12

Tips & Tricks
=============

Specify GTID Usage
------------------

If you want to use GTID Replication, you must update all of the `vagrant/dbX-my.cnf` files with the following options in `[mysqld]`:

```
enforce-gtid-consistency
gtid-mode=ON
```

Specify RBR/SBR
---------------

`vagrant/dbX-my.cnf` files are copied directly to the virtual machines.  If you'd like to specify SBR/RBR/MIXED, add one of the following lines to the `[mysqld]` section of the `my.cnf` template:

```
binlog_format=MIXED
binlog_format=STATEMENT
binlog_format=ROW
```

This is not global because we want to be able to test out non-standard replication configurations.

Use VMWare vs. VirtualBox
-------------------------

```
%> export VAGRANT_DEFAULT_PROVIDER='vmware_fusion'
%> vagrant up
```

CentOS
------

This is the default.  Nothing special is required:

```
%> vagrant up
```

Ubuntu
------

```
%> export VAGRANT_SERVER_URL="https://atlas.hashicorp.com"
%> export VAGRANT_BOX='chef/ubuntu-12.04'
%> vagrant up
```

TO DO
=====

- Support other MySQL's (5.5, 5.7, MariaDB)
- Support customizable replication configurations
- Better `my.cnf` templates

