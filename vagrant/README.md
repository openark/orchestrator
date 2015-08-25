Orchestrator Vagrant Instructions
=================================

Orchestrator's Vagrant defaults to installing five (5) CentOS 6.6 boxes in the following replication topology, with a separate "admin" node:

db1<->db2
 |     |
 v     v
db3   db4 

It is possible to override what gets installed by the use of environmental variables:

VAGRANT_SERVER_URL defaults to 'https://atlas.hashicorp.com'
VAGRANT_DEFAULT_PROVIDER defaults to 'virtualbox'
VAGRANT_BOX defaults to 'nrel/CentOS-6.6-x86_64

The MySQL configuration is such that it is the minimum required to set up replication. For testing such features as delayed replication, RBR/SBR, GTID, it is simply a matter of editing the vagrant/dbX-my.cnf before running the `vagrant up` command.
