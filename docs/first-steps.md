# First Steps with Orchestrator

You have `Orchestrator` installed, deployed and configured. What can you do with it?

A walk through of common commands, mostly on the CLI side

#### Must

##### Discover

You need to discover your MySQL hosts. Either browse to your `http://orchestrator:3000/web/discover` page and submit an instance for discovery, or:

	$ orchestrator-client -c discover -i some.mysql.instance.com:3306

The `:3306` is not required, since the `DefaultInstancePort` configuration is `3306`. You may also:

	$ orchestrator-client -c discover -i some.mysql.instance.com

This discovers a single instance. But: do you also have an `orchestrator` service running? It will pick up from there and
will interrogate this instance for its master and replicas, recursively moving on until the entire topology is revealed.

#### Information

We now assume you have topologies known to `orchestrator` (you have _discovered_ it). Let's say `some.mysql.instance.com`
belongs to one topology. `a.replica.3.instance.com` belongs to another. You may ask the following questions:

	$ orchestrator-client -c clusters
	topology1.master.instance.com:3306
	topology2.master.instance.com:3306

	$ orchestrator-client -c which-master -i some.mysql.instance.com
	some.master.instance.com:3306

	$ orchestrator-client -c which-replicas -i some.mysql.instance.com
	a.replica.instance.com:3306
	another.replica.instance.com:3306

	$ orchestrator-client -c which-cluster -i a.replica.3.instance.com
	topology2.master.instance.com:3306

	$ orchestrator-client -c which-cluster-instances -i a.replica.3.instance.com
	topology2.master.instance.com:3306
	a.replica.1.instance.com:3306
	a.replica.2.instance.com:3306
	a.replica.3.instance.com:3306
	a.replica.4.instance.com:3306
	a.replica.5.instance.com:3306
	a.replica.6.instance.com:3306
	a.replica.7.instance.com:3306
	a.replica.8.instance.com:3306

	$ orchestrator-client -c topology -i a.replica.3.instance.com
	topology2.master.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.1.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.2.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	  + a.replica.3.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	  + a.replica.4.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	  + a.replica.5.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	+ a.replica.6.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	  + a.replica.7.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.8.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]

#### Move stuff around

You may move servers around using various commands. The generic "figure things out automatically" commands are
`relocate` and `relocate-replicas`:

	# Move a.replica.3.instance.com to replicate from a.replica.4.instance.com

	$ orchestrator-client -c relocate -i a.replica.3.instance.com:3306 -d a.replica.4.instance.com
	a.replica.3.instance.com:3306<a.replica.4.instance.com:3306

	$ orchestrator-client -c topology -i a.replica.3.instance.com
	topology2.master.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.1.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.2.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	  + a.replica.4.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	    + a.replica.3.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	  + a.replica.5.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	+ a.replica.6.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	  + a.replica.7.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.8.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]

	# Move the replicas of a.replica.2.instance.com to replicate from a.replica.6.instance.com

	$ orchestrator-client -c relocate-replicas -i a.replica.2.instance.com:3306 -d a.replica.6.instance.com
	a.replica.4.instance.com:3306
	a.replica.5.instance.com:3306

	$ orchestrator-client -c topology -i a.replica.3.instance.com
	topology2.master.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.1.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.2.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.6.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	  + a.replica.4.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	    + a.replica.3.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	  + a.replica.5.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	  + a.replica.7.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.8.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]

`relocate` and `relocate-replicas` automatically figure out how to repoint a replica. Perhaps via GTID; perhaps normal binlog file:pos.
Or maybe there's Pseudo GTID, or is there a binlog server involved? Other variations also supported.

If you want to have greater control:
 - Normal file:pos operations are done via `move-up`, `move-below`
 - Pseudo-GTID specific replica relocation, use `match`, `match-replicas`, `regroup-replicas`.
 - Binlog server operations are typically done with `repoint`, `repoint-replicas`

#### Replication control

You are easily able to see what the following do:

	$ orchestrator-client -c stop-replica -i a.replica.8.instance.com
	$ orchestrator-client -c start-replica -i a.replica.8.instance.com
	$ orchestrator-client -c restart-replica -i a.replica.8.instance.com
	$ orchestrator-client -c set-read-only -i a.replica.8.instance.com
	$ orchestrator-client -c set-writeable -i a.replica.8.instance.com

Break replication by messing with a replica's master host:

	$ orchestrator-client -c detach-replica -i a.replica.8.instance.com

Don't worry, this is reversible:

	$ orchestrator-client -c reattach-replica -i a.replica.8.instance.com

#### Crash analysis & recovery

Are your clusters healthy?

	$ orchestrator-client -c replication-analysis
	some.master.instance.com:3306 (cluster some.master.instance.com:3306): DeadMaster
	a.replica.6.instance.com:3306 (cluster topology2.master.instance.com:3306): DeadIntermediateMaster

	$ orchestrator-client -c topology -i a.replica.6.instance.com
	topology2.master.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.1.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.2.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.6.instance.com:3306 [last check invalid,5.6.17-log,STATEMENT,>>]
	  + a.replica.4.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	    + a.replica.3.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	  + a.replica.5.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	  + a.replica.7.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.8.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]

Ask `orchestrator` to recover the above dead intermediate master:

	$ orchestrator-client -c recover -i a.replica.6.instance.com:3306
	a.replica.8.instance.com:3306

	$ orchestrator-client -c topology -i a.replica.8.instance.com
	topology2.master.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.1.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.2.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	+ a.replica.6.instance.com:3306 [last check invalid,5.6.17-log,STATEMENT,>>]
	+ a.replica.8.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	  + a.replica.4.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]
	    + a.replica.3.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	  + a.replica.5.instance.com:3306 [OK,5.6.17-log,STATEMENT]
	  + a.replica.7.instance.com:3306 [OK,5.6.17-log,STATEMENT,>>]

#### More

The above should get you up and running. For more please consult the [Manual](toc.md). For CLI commands listing just run:

	orchestrator-client -help
