/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/math"
	"github.com/outbrain/orchestrator/go/app"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
)

var AppVersion string

const prompt string = `
orchestrator [-c command] [-i instance] [-d destination] [--verbose|--debug] [... cli ] | http

Cheatsheet:
    Run orchestrator in HTTP mode:

        orchestrator --debug http

    See all possible commands:

        orchestrator -c help

    Usage for most commands:
        orchestrator -c <command> [-i <instance.fqdn>] [-d <destination.fqdn>] [--verbose|--debug]

    -i (instance):
        instance on which to operate, in "hostname" or "hostname:port" format.
        Default port is 3306 (or DefaultInstancePort in config)
        For some commands this argument can be ommitted altogether, and the
        value is implicitly the local hostname.
    -d (Destination)
        destination instance (used when moving replicas around or when failing over)
    -s (Sibling/Subinstance/deStination) - synonym to "-d"

    -c (command):
        Listed below are all available commands; all of which apply for CLI execution (ignored by HTTP mode).
        Different flags are required for different commands; see specific documentation per commmand.

    Topology refactoring, generic aka "smart" commands
        These operations let orchestrator pick the best course of action for relocating slaves. It may choose to use
        standard binlog file:pos math, GTID, Pseudo-GTID, or take advantage of binlog servers, or combine two or more
        methods in a multi-step operation.
        In case a of a multi-step operation, failure may result in slaves only moving halfway to destination point. Nonetheless
        they will be in a valid position.

        relocate
            Relocate a slave beneath another (destination) instance. The choice of destination is almost arbitrary;
            it must not be a child/descendant of the instance, but otherwise it can be anywhere, and can be a normal slave
            or a binlog server. Orchestrator will choose the best course of action to relocate the slave.
            No action taken when destination instance cannot act as master (e.g. has no binary logs, is of incompatible version, incompatible binlog format etc.)
            Examples:

            orchestrator -c relocate -i slave.to.relocate.com -d instance.that.becomes.its.master

            orchestrator -c relocate -d destination.instance.that.becomes.its.master
                -i not given, implicitly assumed local hostname

            (this command was previously named "relocate-below")

        relocate-slaves
            Relocates all or part of the slaves of a given instance under another (destination) instance. This is
            typically much faster than relocating slaves one by one.
            Orchestrator chooses the best course of action to relocation the slaves. It may choose a multi-step operations.
            Some slaves may succeed and some may fail the operation.
            The instance (slaves' master) itself may be crashed or inaccessible. It is not contacted throughout the operation.
            Examples:

            orchestrator -c relocate-slaves -i instance.whose.slaves.will.relocate -d instance.that.becomes.their.master

            orchestrator -c relocate-slaves -i instance.whose.slaves.will.relocate -d instance.that.becomes.their.master --pattern=regexp.filter
                only apply to those instances that match given regex

    Topology refactoring using classic MySQL replication commands
        (ie STOP SLAVE; START SLAVE UNTIL; CHANGE MASTER TO; ...)
        These commands require connected topology: slaves that are up and running; a lagging, stopped or
        failed slave will disable use of most these commands. At least one, and typically two or more slaves
        will be stopped for a short time during these operations.

        move-up
            Move a slave one level up the topology; makes it replicate from its grandparent and become sibling of
            its parent. It is OK if the instance's master is not replicating. Examples:

            orchestrator -c move-up -i slave.to.move.up.com:3306

            orchestrator -c move-up
                -i not given, implicitly assumed local hostname

        move-up-slaves
            Moves slaves of the given instance one level up the topology, making them siblings of given instance.
            This is a (faster) shortcut to executing move-up on all slaves of given instance.
            Examples:

            orchestrator -c move-up-slaves -i slave.whose.subslaves.will.move.up.com[:3306]

            orchestrator -c move-up-slaves -i slave.whose.subslaves.will.move.up.com[:3306] --pattern=regexp.filter
                only apply to those instances that match given regex

        move-below
            Moves a slave beneath its sibling. Both slaves must be actively replicating from same master.
            The sibling will become instance's master. No action taken when sibling cannot act as master
            (e.g. has no binary logs, is of incompatible version, incompatible binlog format etc.)
            Example:

            orchestrator -c move-below -i slave.to.move.com -d sibling.slave.under.which.to.move.com

            orchestrator -c move-below -d sibling.slave.under.which.to.move.com
                -i not given, implicitly assumed local hostname

        move-equivalent
            Moves a slave beneath another server, based on previously recorded "equivalence coordinates". Such coordinates
            are obtained whenever orchestrator issues a CHANGE MASTER TO. The "before" and "after" masters coordinates are
            persisted. In such cases where the newly relocated slave is unable to replicate (e.g. firewall issues) it is then
            easy to revert the relocation via "move-equivalent".
            The command works if and only if orchestrator has an exact mapping between the slave's current replication coordinates
            and some other coordinates.
            Example:

            orchestrator -c move-equivalent -i slave.to.revert.master.position.com -d master.to.move.to.com

        enslave-siblings
            Turn all siblings of a slave into its sub-slaves. No action taken for siblings that cannot become
            slaves of given instance (e.g. incompatible versions, binlog format etc.). This is a (faster) shortcut
            to executing move-below for all siblings of the given instance. Example:

            orchestrator -c enslave-siblings -i slave.whose.siblings.will.move.below.com

        enslave-master
            Turn an instance into a master of its own master; essentially switch the two. Slaves of each of the two
            involved instances are unaffected, and continue to replicate as they were.
            The instance's master must itself be a slave. It does not necessarily have to be actively replicating.

            orchestrator -c enslave-master -i slave.that.will.switch.places.with.its.master.com

        repoint
            Make the given instance replicate from another instance without changing the binglog coordinates. There
            are little sanity checks to this and this is a risky operation. Use cases are: a rename of the master's
            host, a corruption in relay-logs, move from beneath MaxScale & Binlog-server. Examples:

            orchestrator -c repoint -i slave.to.operate.on.com -d new.master.com

            orchestrator -c repoint -i slave.to.operate.on.com
                The above will repoint the slave back to its existing master without change

            orchestrator -c repoint
                -i not given, implicitly assumed local hostname

        repoint-slaves
            Repoint all slaves of given instance to replicate back from the instance. This is a convenience method
            which implies a one-by-one "repoint" command on each slave.

            orchestrator -c repoint-slaves -i instance.whose.slaves.will.be.repointed.com

            orchestrator -c repoint-slaves
                -i not given, implicitly assumed local hostname

        make-co-master
            Create a master-master replication. Given instance is a slave which replicates directly from a master.
            The master is then turned to be a slave of the instance. The master is expected to not be a slave.
            The read_only property of the slve is unaffected by this operation. Examples:

            orchestrator -c make-co-master -i slave.to.turn.into.co.master.com

            orchestrator -c make-co-master
                -i not given, implicitly assumed local hostname

        get-candidate-slave
            Information command suggesting the most up-to-date slave of a given instance, which can be promoted
            as local master to its siblings. If replication is up and running, this command merely gives an
            estimate, since slaves advance and progress continuously in different pace. If all slaves of given
            instance have broken replication (e.g. because given instance is dead), then this command provides
            with a definitve candidate, which could act as a replace master. See also regroup-slaves. Example:

            orchestrator -c get-candidate-slave -i instance.with.slaves.one.of.which.may.be.candidate.com

        regroup-slaves-bls
            Given an instance that has Binlog Servers for slaves, promote one such Binlog Server over its other
            Binlog Server siblings.

            Example:

            orchestrator -c regroup-slaves-bls -i instance.with.binlog.server.slaves.com

            --debug is your friend.


    Topology refactoring using GTID
        These operations only work if GTID (either Oracle or MariaDB variants) is enabled on your servers.

        move-gtid
            Move a slave beneath another (destination) instance. Orchestrator will reject the operation if GTID is
            not enabled on the slave, or is not supported by the would-be master.
            You may try and move the slave under any other instance; there are no constraints on the family ties the
            two may have, though you should be careful as not to try and replicate from a descendant (making an
            impossible loop).
            Examples:

            orchestrator -c move-gtid -i slave.to.move.com -d instance.that.becomes.its.master

            orchestrator -c match -d destination.instance.that.becomes.its.master
                -i not given, implicitly assumed local hostname

        move-slaves-gtid
            Moves all slaves of a given instance under another (destination) instance using GTID. This is a (faster)
            shortcut to moving each slave via "move-gtid".
            Orchestrator will only move those slaves configured with GTID (either Oracle or MariaDB variants) and under the
            condition the would-be master supports GTID.
            Examples:

            orchestrator -c move-slaves-gtid -i instance.whose.slaves.will.relocate -d instance.that.becomes.their.master

            orchestrator -c move-slaves-gtid -i instance.whose.slaves.will.relocate -d instance.that.becomes.their.master --pattern=regexp.filter
                only apply to those instances that match given regex

        regroup-slaves-gtid
            Given an instance (possibly a crashed one; it is never being accessed), pick one of its slave and make it
            local master of its siblings, using GTID. The rules are similar to those in the "regroup-slaves" command.
            Example:

            orchestrator -c regroup-slaves-gtid -i instance.with.gtid.and.slaves.one.of.which.will.turn.local.master.if.possible

            --debug is your friend.


    Topology refactoring using Pseudo-GTID
        These operations require that the topology's master is periodically injected with pseudo-GTID,
        and that the PseudoGTIDPattern configuration is setup accordingly. Also consider setting
        DetectPseudoGTIDQuery.
        Operations via Pseudo-GTID are typically slower, since they involve scanning of binary/relay logs.
        They impose less constraints on topology locations and affect less servers. Only servers that
        are being relocateed have their replication stopped. Their masters or destinations are unaffected.

        match
            Matches a slave beneath another (destination) instance. The choice of destination is almost arbitrary;
            it must not be a child/descendant of the instance. But otherwise they don't have to be direct siblings,
            and in fact (if you know what you're doing), they don't actually have to belong to the same topology.
            The operation expects the relocated instance to be "behind" the destination instance. It only finds out
            whether this is the case by the end; the operation is cancelled in the event this is not the case.
            No action taken when destination instance cannot act as master (e.g. has no binary logs, is of incompatible version, incompatible binlog format etc.)
            Examples:

            orchestrator -c match -i slave.to.relocate.com -d instance.that.becomes.its.master

            orchestrator -c match -d destination.instance.that.becomes.its.master
                -i not given, implicitly assumed local hostname

            (this command was previously named "match-below")

        match-slaves
            Matches all slaves of a given instance under another (destination) instance. This is a (faster) shortcut
            to matching said slaves one by one under the destination instance. In fact, this bulk operation is highly
            optimized and can execute in orders of magnitue faster, depeding on the nu,ber of slaves involved and their
            respective position behind the instance (the more slaves, the more savings).
            The instance itself may be crashed or inaccessible. It is not contacted throughout the operation. Examples:

            orchestrator -c match-slaves -i instance.whose.slaves.will.relocate -d instance.that.becomes.their.master

            orchestrator -c match-slaves -i instance.whose.slaves.will.relocate -d instance.that.becomes.their.master --pattern=regexp.filter
                only apply to those instances that match given regex

            (this command was previously named "multi-match-slaves")

        match-up
            Transport the slave one level up the hierarchy, making it child of its grandparent. This is
            similar in essence to move-up, only based on Pseudo-GTID. The master of the given instance
            does not need to be alive or connected (and could in fact be crashed). It is never contacted.
            Grandparent instance must be alive and accessible.
            Examples:

            orchestrator -c match-up -i slave.to.match.up.com:3306

            orchestrator -c match-up
                -i not given, implicitly assumed local hostname

        match-up-slaves
            Matches slaves of the given instance one level up the topology, making them siblings of given instance.
            This is a (faster) shortcut to executing match-up on all slaves of given instance. The instance need
            not be alive / accessib;e / functional. It can be crashed.
            Example:

            orchestrator -c match-up-slaves -i slave.whose.subslaves.will.match.up.com

            orchestrator -c match-up-slaves -i slave.whose.subslaves.will.match.up.com[:3306] --pattern=regexp.filter
                only apply to those instances that match given regex

        rematch
            Reconnect a slave onto its master, via PSeudo-GTID. The use case for this operation is a non-crash-safe
            replication configuration (e.g. MySQL 5.5) with sync_binlog=1 and log_slave_updates. This operation
            implies crash-safe-replication and makes it possible for the slave to reconnect. Example:

            orchestrator -c rematch -i slave.to.rematch.under.its.master

        regroup-slaves
            Given an instance (possibly a crashed one; it is never being accessed), pick one of its slave and make it
            local master of its siblings, using Pseudo-GTID. It is uncertain that there *is* a slave that will be able to
            become master to all its siblings. But if there is one, orchestrator will pick such one. There are many
            constraints, most notably the replication positions of all slaves, whether they use log_slave_updates, and
            otherwise version compatabilities etc.
            As many slaves that can be regrouped under promoted slves are operated on. The rest are untouched.
            This command is useful in the event of a crash. For example, in the event that a master dies, this operation
            can promote a candidate replacement and set up the remaining topology to correctly replicate from that
            replacement slave. Example:

            orchestrator -c regroup-slaves -i instance.with.slaves.one.of.which.will.turn.local.master.if.possible

            --debug is your friend.

    General replication commands
        These commands issue various statements that relate to replication.

        enable-gtid
            If possible, enable GTID replication. This works on Oracle (>= 5.6, gtid-mode=1) and MariaDB (>= 10.0).
            Replication is stopped for a short duration so as to reconfigure as GTID. In case of error replication remains
            stopped. Example:

            orchestrator -c enable-gtid -i slave.compatible.with.gtid.com

        disable-gtid
            Assuming slave replicates via GTID, disable GTID replication and resume standard file:pos replication. Example:

            orchestrator -c disable-gtid -i slave.replicating.via.gtid.com

        reset-master-gtid-remove-own-uuid
            Assuming GTID is enabled, Reset master on instance, remove GTID entries generated by the instance.
            This operation is only allowed on Oracle-GTID enabled servers that have no slaves.
            Is is used for cleaning up the GTID mess incurred by mistakenly issuing queries on the slave (even such
            queries as "FLUSH ENGINE LOGS" that happen to write to binary logs). Example:

            orchestrator -c reset-master-gtid-remove-own-uuid -i slave.running.with.gtid.com

        stop-slave
            Issues a STOP SLAVE; command. Example:

            orchestrator -c stop-slave -i slave.to.be.stopped.com

        start-slave
            Issues a START SLAVE; command. Example:

            orchestrator -c start-slave -i slave.to.be.started.com

        restart-slave
            Issues STOP SLAVE + START SLAVE; Example:

            orchestrator -c restart-slave -i slave.to.be.started.com

        skip-query
            On a failed replicating slave, skips a single query and attempts to resume replication.
            Only applies when the replication seems to be broken on SQL thread (e.g. on duplicate
            key error). Also works in GTID mode. Example:

            orchestrator -c skip-query -i slave.with.broken.sql.thread.com

        reset-slave
            Issues a RESET SLAVE command. Destructive to replication. Example:

            orchestrator -c reset-slave -i slave.to.reset.com

        detach-slave
            Stops replication and modifies binlog position into an impossible, yet reversible, value.
            This effectively means the replication becomes broken. See reattach-slave. Example:

            orchestrator -c detach-slave -i slave.whose.replication.will.break.com

            Issuing this on an already detached slave will do nothing.

        reattach-slave
            Undo a detach-slave operation. Reverses the binlog change into the original values, and
            resumes replication. Example:

            orchestrator -c reattach-slave -i detahced.slave.whose.replication.will.amend.com

            Issuing this on an attached (i.e. normal) slave will do nothing.

        detach-slave-master-host
            Stops replication and modifies Master_Host into an impossible, yet reversible, value.
            This effectively means the replication becomes broken. See reattach-slave-master-host. Example:

            orchestrator -c detach-slave-master-host -i slave.whose.replication.will.break.com

            Issuing this on an already detached slave will do nothing.

        reattach-slave-master-host
            Undo a detach-slave-master-host operation. Reverses the hostname change into the original value, and
            resumes replication. Example:

            orchestrator -c reattach-slave-master-host -i detahced.slave.whose.replication.will.amend.com

            Issuing this on an attached (i.e. normal) slave will do nothing.

		restart-slave-statements
			Prints a list of statements to execute to stop then restore slave to same execution state.
			Provide --statement for injected statement.
			This is useful for issuing a command that can only be executed whiel slave is stopped. Such
			commands are any of CHANGE MASTER TO.
			Orchestrator will not execute given commands, only print them as courtesy. It may not have
			the privileges to execute them in the first place. Example:

			orchestrator -c restart-slave-statements -i some.slave.com -statement="change master to master_heartbeat_period=5"

    General instance commands
        Applying general instance configuration and state

        set-read-only
            Turn an instance read-only, via SET GLOBAL read_only := 1. Examples:

            orchestrator -c set-read-only -i instance.to.turn.read.only.com

            orchestrator -c set-read-only
                -i not given, implicitly assumed local hostname

        set-writeable
            Turn an instance writeable, via SET GLOBAL read_only := 0. Example:

            orchestrator -c set-writeable -i instance.to.turn.writeable.com

            orchestrator -c set-writeable
                -i not given, implicitly assumed local hostname

    Binlog commands
        Commands that investigate/work on binary logs

        flush-binary-logs
            Flush binary logs on an instance. Examples:

            orchestrator -c flush-binary-logs -i instance.with.binary.logs.com

            orchestrator -c flush-binary-logs -i instance.with.binary.logs.com --binlog=mysql-bin.002048
                Flushes binary logs until reaching given number. Fails when current number is larger than input

        purge-binary-logs
            Purge binary logs on an instance. Examples:

            orchestrator -c purge-binary-logs -i instance.with.binary.logs.com --binlog mysql-bin.002048

                Purges binary logs until given log

        last-pseudo-gtid
            Information command; an authoritative way of detecting whether a Pseudo-GTID event exist for an instance,
            and if so, output the last Pseudo-GTID entry and its location. Example:

            orchestrator -c last-pseudo-gtid -i instance.with.possible.pseudo-gtid.injection

        find-binlog-entry
            Get binlog file:pos of entry given by --pattern (exact full match, not a regular expression) in a given instance.
            This will search the instance's binary logs starting with most recent, and terminate as soon as an exact match is found.
            The given input is not a regular expression. It must fully match the entry (not a substring).
            This is most useful when looking for uniquely identifyable values, such as Pseudo-GTID. Example:

            orchestrator -c find-binlog-entry -i instance.to.search.on.com --pattern "insert into my_data (my_column) values ('distinct_value_01234_56789')"

                Prints out the binlog file:pos where the entry is found, or errors if unfound.

        correlate-binlog-pos
            Given an instance (-i) and binlog coordinates (--binlog=file:pos), find the correlated coordinates in another instance (-d).
            "Correlated coordinates" are those that present the same point-in-time of sequence of binary log events, untangling
            the mess of different binlog file:pos coordinates on different servers.
            This operation relies on Pseudo-GTID: your servers must have been pre-injected with PSeudo-GTID entries as these are
            being used as binlog markers in the correlation process.
            You must provide a valid file:pos in the binlogs of the source instance (-i), and in response get the correlated
            coordinates in the binlogs of the destination instance (-d). This operation does not work on relay logs.
            Example:

            orchestrator -c correlate-binlog-pos  -i instance.with.binary.log.com --binlog=mysql-bin.002366:14127 -d other.instance.with.binary.logs.com

                Prints out correlated coordinates, e.g.: "mysql-bin.002302:14220", or errors out.

    Pool commands
        Orchestrator provides with getter/setter commands for handling pools. It does not on its own investigate pools,
        but merely accepts and provides association of an instance (host:port) and a pool (any_name).

        submit-pool-instances
            Submit a pool name with a list of instances in that pool. This removes any previous instances associated with
            that pool. Expecting comma delimited list of instances

            orchestrator -c submit-pool-instances --pool name_of_pool -i pooled.instance1.com,pooled.instance2.com:3306,pooled.instance3.com

        cluster-pool-instances
            List all pools and their associated instances. Output is in tab delimited format, and lists:
            cluster_name, cluster_alias, pool_name, pooled instance
            Example:

            orchestrator -c cluster-pool-instances

				which-heuristic-cluster-pool-instances
						List instances belonging to a cluster, which are also in some pool or in a specific given pool.
						Not all instances are listed: unreachable, downtimed instances ar left out. Only those that should be
						responsive and healthy are listed. This serves applications in getting information about instances
						that could be queried (this complements a proxy behavior in providing the *list* of instances).
						Examples:

						orchestrator -c which-heuristic-cluster-pool-instances --alias mycluster
								Get the instances of a specific cluster, no specific pool

						orchestrator -c which-heuristic-cluster-pool-instances --alias mycluster --pool some_pool
								Get the instances of a specific cluster and which belong to a given pool

						orchestrator -c which-heuristic-cluster-pool-instances -i instance.belonging.to.a.cluster
								Cluster inferred by given instance

						orchestrator -c which-heuristic-cluster-pool-instances
								Cluster inferred by local hostname

    Information commands
        These commands provide/store information about topologies, replication connections, or otherwise orchstrator's
        "inventory".

        find
            Find instances whose hostname matches given regex pattern. Example:

            orchestrator -c find -pattern "backup.*us-east"

        clusters
            List all clusters known to orchestrator. A cluster (aka topology, aka chain) is identified by its
            master (or one of its master if more than one exists). Example:

            orchesrtator -c clusters
                -i not given, implicitly assumed local hostname

        all-clusters-masters
            List of writeable masters, one per cluster.
			For most single-master topologies, this is trivially the master.
			For active-active master-master topologies, this ensures only one of
			the masters is returned.

			Example:

            orchestrator -c all-clusters-masters

        topology
            Show an ascii-graph of a replication topology, given a member of that topology. Example:

            orchestrator -c topology -i instance.belonging.to.a.topology.com

            orchestrator -c topology
                -i not given, implicitly assumed local hostname

            Instance must be already known to orchestrator. Topology is generated by orchestrator's mapping
            and not from synchronuous investigation of the instances. The generated topology may include
            instances that are dead, or whose replication is broken.

        all-instances
            List the complete known set of instances. Similar to '-c find -pattern "."'
			Example:

            orchestrator -c all-instances

        which-instance
            Output the fully-qualified hostname:port representation of the given instance, or error if unknown
            to orchestrator. Examples:

            orchestrator -c which-instance -i instance.to.check.com

            orchestrator -c which-instance
                -i not given, implicitly assumed local hostname

        which-cluster
            Output the name of the cluster an instance belongs to, or error if unknown to orchestrator. Examples:

            orchestrator -c which-cluster -i instance.to.check.com

            orchestrator -c which-cluster
                -i not given, implicitly assumed local hostname

        which-cluster-instances
            Output the list of instances participating in same cluster as given instance; output is one line
            per instance, in hostname:port format. Examples:

            orchestrator -c which-cluster-instances -i instance.to.check.com

            orchestrator -c which-cluster-instances
                -i not given, implicitly assumed local hostname

            orchestrator -c which-cluster-instances -alias some_alias
                assuming some_alias is a known cluster alias (see ClusterNameToAlias or DetectClusterAliasQuery configuration)

        which-cluster-domain
            Output the domain name of given cluster, indicated by instance or alias. This depends on
						the DetectClusterDomainQuery configuration. Example:

            orchestrator -c which-cluster-domain -i instance.to.check.com

            orchestrator -c which-cluster-domain
                -i not given, implicitly assumed local hostname

            orchestrator -c which-cluster-domain -alias some_alias
                assuming some_alias is a known cluster alias (see ClusterNameToAlias or DetectClusterAliasQuery configuration)

				which-heuristic-domain-instance
						Returns the instance associated as the cluster's writer with a cluster's domain name.
						Given a cluster, orchestrator looks for the domain name indicated by this cluster, and proceeds to search for
						a stord key-value attribute for that domain name. This would be the writer host for the given domain.
						See also set-heuristic-domain-instance, this is meant to be a temporary service mimicking in micro-scale a
						service discovery functionality.
						Example:

						orchestrator -c which-heuristic-domain-instance -alias some_alias
							Detects the domain name for given cluster, reads from key-value store the writer host associated with the domain name.

						orchestrator -c which-heuristic-domain-instance -i instance.of.some.cluster
							Cluster is inferred by a member instance (the instance is not necessarily the master)

				which-cluster-master
						Output the name of the active master in a given cluster, indicated by instance or alias.
						An "active" master is one that is writable and is not marked as downtimed due to a topology recovery.
						Examples:

            orchestrator -c which-cluster-master -i instance.to.check.com

            orchestrator -c which-cluster-master
                -i not given, implicitly assumed local hostname

            orchestrator -c which-cluster-master -alias some_alias
                assuming some_alias is a known cluster alias (see ClusterNameToAlias or DetectClusterAliasQuery configuration)

        which-cluster-osc-slaves
            Output a list of slaves in same cluster as given instance, that would server as good candidates as control slaves
            for a pt-online-schema-change operation.
            Those slaves would be used for replication delay so as to throtthe osc operation. Selected slaves will include,
            where possible: intermediate masters, their slaves, 3rd level slaves, direct non-intermediate-master slaves.

            orchestrator -c which-cluster-osc-slaves -i instance.to.check.com

            orchestrator -c which-cluster-osc-slaves
                -i not given, implicitly assumed local hostname

            orchestrator -c which-cluster-osc-slaves -alias some_alias
                assuming some_alias is a known cluster alias (see ClusterNameToAlias or DetectClusterAliasQuery configuration)

				which-lost-in-recovery
						List instances marked as downtimed for being lost in a recovery process. This depends on the configuration
						of MasterFailoverLostInstancesDowntimeMinutes. The output of this command lists heuristically recent
						"lost" instances that probabaly should be recycled. Note that when the 'downtime' flag expires (or
						is reset by '-c end-downtime') an instance no longer appears on this list.
						The topology recovery process injects a magic hint when downtiming lost instances, that is picked up
						by this command. Examples:

						orchestrator -c which-lost-in-recovery
								Lists all heuristically-recent known lost instances

        which-master
            Output the fully-qualified hostname:port representation of a given instance's master. Examples:

            orchestrator -c which-master -i a.known.slave.com

            orchestrator -c which-master
                -i not given, implicitly assumed local hostname

        which-slaves
            Output the fully-qualified hostname:port list of slaves (one per line) of a given instance (or empty
            list if instance is not a master to anyone). Examples:

            orchestrator -c which-slaves -i a.known.instance.com

            orchestrator -c which-slaves
                -i not given, implicitly assumed local hostname

        get-cluster-heuristic-lag
            For a given cluster (indicated by an instance or alias), output a heuristic "representative" lag of that cluster.
            The output is obtained by examining the slaves that are member of "which-cluster-osc-slaves"-command, and
            getting the maximum slave lag of those slaves. Recall that those slaves are a subset of the entire cluster,
            and that they are ebing polled periodically. Hence the output of this command is not necessarily up-to-date
            and does not represent all slaves in cluster. Examples:

            orchestrator -c get-cluster-heuristic-lag -i instance.that.is.part.of.cluster.com

            orchestrator -c get-cluster-heuristic-lag
                -i not given, implicitly assumed local host, cluster implied

            orchestrator -c get-cluster-heuristic-lag -alias some_alias
                assuming some_alias is a known cluster alias (see ClusterNameToAlias or DetectClusterAliasQuery configuration)

        instance-status
            Output short status on a given instance (name, replication status, noteable configuration). Example2:

            orchestrator -c instance-status -i instance.to.investigate.com

            orchestrator -c instance-status
                -i not given, implicitly assumed local hostname

        snapshot-topologies
            Take a snapshot of existing topologies. This will record minimal replication topology data: the identity
            of an instance, its master and its cluster.
            Taking a snapshot later allows for reviewing changes in topologies. One might wish to invoke this command
            on a daily basis, and later be able to solve questions like 'where was this instacne replicating from before
            we moved it?', 'which instances were replication from this instance a week ago?' etc. Example:

            orchestrator -c snapshot-topologies

    Orchestrator instance management
        These command dig into the way orchestrator manages instances and operations on instances

        discover
            Request that orchestrator cotacts given instance, reads its status, and upsert it into
            orchestrator's respository. Examples:

            orchestrator -c discover -i instance.to.discover.com:3306

            orchestrator -c discover -i cname.of.instance

            orchestrator -c discover
                -i not given, implicitly assumed local hostname

            Orchestrator will resolve CNAMEs and VIPs.

        forget
            Request that orchestrator removed given instance from its repository. If the instance is alive
            and connected through replication to otherwise known and live instances, orchestrator will
            re-discover it by nature of its discovery process. Instances are auto-removed via config's
            UnseenAgentForgetHours. If you happen to know a machine is decommisioned, for example, it
            can be nice to remove it from the repository before it auto-expires. Example:

            orchestrator -c forget -i instance.to.forget.com

            Orchestrator will *not* resolve CNAMEs and VIPs for given instance.

        begin-maintenance
            Request a maintenance lock on an instance. Topology changes require placing locks on the minimal set of
            affected instances, so as to avoid an incident of two uncoordinated operations on a smae instance (leading
            to possible chaos). Locks are placed in the backend database, and so multiple orchestrator instances are safe.
            Operations automatically acquire locks and release them. This command manually acquires a lock, and will
            block other operations on the instance until lock is released.
            Note that orchestrator automatically assumes locks to be expired after MaintenanceExpireMinutes (in config).
            Examples:

            orchestrator -c begin-maintenance -i instance.to.lock.com --duration=3h --reason="load testing; do not disturb"
                accepted duration format: 10s, 30m, 24h, 3d, 4w

            orchestrator -c begin-maintenance -i instance.to.lock.com --reason="load testing; do not disturb"
                --duration not given; default to config's MaintenanceExpireMinutes

        end-maintenance
            Remove maintenance lock; such lock may have been gained by an explicit begin-maintenance command implicitly
            by a topology change. You should generally only remove locks you have placed manually; orchestrator will
            automatically expire locks after MaintenanceExpireMinutes (in config).
            Example:

            orchestrator -c end-maintenance -i locked.instance.com

        begin-downtime
            Mark an instance as downtimed. A downtimed instance is assumed to be taken care of, and recovery-analysis does
            not apply for such an instance. As result, no recommendation for recovery, and no automated-recovery are issued
            on a downtimed instance.
            Downtime is different than maintanence in that it places no lock (mainenance uses an exclusive lock on the instance).
            It is OK to downtime an instance that is already downtimed -- the new begin-downtime command will override whatever
            previous downtime attributes there were on downtimes instance.
            Note that orchestrator automatically assumes downtime to be expired after MaintenanceExpireMinutes (in config).
            Examples:

            orchestrator -c begin-downtime -i instance.to.downtime.com --duration=3h --reason="dba handling; do not do recovery"
                accepted duration format: 10s, 30m, 24h, 3d, 4w

            orchestrator -c begin-downtime -i instance.to.lock.com --reason="dba handling; do not do recovery"
                --duration not given; default to config's MaintenanceExpireMinutes

        end-downtime
            Indicate an instance is no longer downtimed. Typically you should not need to use this since
            a downtime is always bounded by a duration and auto-expires. But you may use this to forcibly
            indicate the active downtime should be expired now.
            Example:

            orchestrator -c end-downtime -i downtimed.instance.com

    Crash recovery commands

        recover
            Do auto-recovery given a dead instance. Orchestrator chooses the best course of action.
            The given instance must be acknowledged as dead and have slaves, or else there's nothing to do.
            See "replication-analysis" command.
            Orchestrator executes external processes as configured by *Processes variables.
            --debug is your friend. Example:

            orchestrator -c recover -i dead.instance.com --debug

        recover-lite
            Do auto-recovery given a dead instance. Orchestrator chooses the best course of action, exactly
            as in "-c recover". Orchestratir will *not* execute external processes.

            orchestrator -c recover-lite -i dead.instance.com --debug

				force-master-takeover
						Forcibly discard master and promote another (direct child) instance instead, even if everything is running well.
						This allows for planned switchover.
						NOTE:
						- You must specify the instance to promote via "-d"
						- Promoted instance must be a direct child of the existing master
						- This will not work in a master-master configuration
						- Orchestrator just treats this command as a DeadMaster failover scenario
						- It is STRONGLY suggested that you first relocate everything below your chosen instance-to-promote.
						  It *is* a planned failover thing.
						- Otherwise orchestrator will do its thing in moving instances around, hopefully promoting your requested
						  server on top.
						- Orchestrator will issue all relevant pre-failover and post-failover external processes.
						- In this command orchestrator will not issue 'SET GLOBAL read_only=1' on the existing master, nor will
						  it issue a 'FLUSH TABLES WITH READ LOCK'. Please see the 'graceful-master-takeover' command.
						Examples:

						orchestrator -c force-master-takeover -alias mycluster -d immediate.child.of.master.com
								Indicate cluster by alias. Orchestrator automatically figures out the master

						orchestrator -c force-master-takeover -i instance.in.relevant.cluster.com -d immediate.child.of.master.com
								Indicate cluster by an instance. You don't structly need to specify the master, orchestrator
								will infer the master's identify.

				graceful-master-takeover
						Gracefully discard master and promote another (direct child) instance instead, even if everything is running well.
						This allows for planned switchover.
						NOTE:
						- Promoted instance must be a direct child of the existing master
						- Promoted instance must be the *only* direct child of the existing master. It *is* a planned failover thing.
						- Orchestrator will first issue a "set global read_only=1" on existing master
						- It will promote candidate master to the binlog positions of the existing master after issuing the above
						- There _could_ still be statements issued and executed on the existing master by SUPER users, but those are ignored.
						- Orchestrator then proceeds to handle a DeadMaster failover scenario
						- Orchestrator will issue all relevant pre-failover and post-failover external processes.
						Examples:

						orchestrator -c graceful-master-takeover -alias mycluster
								Indicate cluster by alias. Orchestrator automatically figures out the master and verifies it has a single direct replica

						orchestrator -c force-master-takeover -i instance.in.relevant.cluster.com
								Indicate cluster by an instance. You don't structly need to specify the master, orchestrator
								will infer the master's identify.

        replication-analysis
            Request an analysis of potential crash incidents in all known topologies.
            Output format is not yet stabilized and may change in the future. Do not trust the output
            for automated parsing. Use web API instead, at this time. Example:

            orchestrator -c replication-analysis

        ack-cluster-recoveries
            Acknowledge recoveries for a given cluster; this unblocks pending future recoveries.
            Acknowledging a recovery requires a comment (supply via --reason). Acknowledgement clears the in-active-period
            flag for affected recoveries, which in turn affects any blocking recoveries.
            Multiple recoveries may be affected. Only unacknowledged recoveries will be affected.
            Examples:

            orchestrator -c ack-cluster-recoveries -i instance.in.a.cluster.com --reason="dba has taken taken necessary steps"
                 Cluster is indicated by any of its members. The recovery need not necessarily be on/to given instance.

            orchestrator -c ack-cluster-recoveries -alias some_alias --reason="dba has taken taken necessary steps"
                 Cluster indicated by alias

        ack-instance-recoveries
            Acknowledge recoveries for a given instance; this unblocks pending future recoveries.
            Acknowledging a recovery requires a comment (supply via --reason). Acknowledgement clears the in-active-period
            flag for affected recoveries, which in turn affects any blocking recoveries.
            Multiple recoveries may be affected. Only unacknowledged recoveries will be affected.
            Example:

            orchestrator -c ack-cluster-recoveries -i instance.that.failed.com --reason="dba has taken taken necessary steps"

    Instance meta commands

        register-candidate
            Indicate that a specific instance is a preferred candidate for master promotion. Upon a dead master
            recovery, orchestrator will do its best to promote instances that are marked as candidates. However
            orchestrator cannot guarantee this will always work. Issues like version compatabilities, binlog format
            etc. are limiting factors.
            You will want to mark an instance as a candidate when: it is replicating directly from the master, has
            binary logs and log_slave_updates is enabled, uses same binlog_format as its siblings, compatible version
            as its siblings. If you're using DataCenterPattern & PhysicalEnvironmentPattern (see configuration),
            you would further wish to make sure you have a candidate in each data center.
            Orchestrator first promotes the best-possible slave, and only then replaces it with your candidate,
            and only if both in same datcenter and physical enviroment.
            An instance needs to continuously be marked as candidate, so as to make sure orchestrator is not wasting
            time with stale instances. Orchestrator periodically clears candidate-registration for instances that have
            not been registeres for over CandidateInstanceExpireMinutes (see config).
            Example:

            orchestrator -c register-candidate -i candidate.instance.com

            orchestrator -c register-candidate
                -i not given, implicitly assumed local hostname

        register-hostname-unresolve
            Assigns the given instance a virtual (aka "unresolved") name. When moving slaves under an instance with assigned
            "unresolve" name, orchestrator issues a CHANGE MASTER TO MASTER_HOST='<the unresovled name instead of the fqdn>' ...
            This is useful in cases where your master is behind virtual IP (e.g. active/passive masters with shared storage or DRBD,
            e.g. binlog servers sharing common VIP).
            A "repoint" command is useful after "register-hostname-unresolve": you can repoint slaves of the instance to their exact
            same location, and orchestrator will swap the fqdn of their master with the unresolved name.
            Such registration must be periodic. Orchestrator automatically expires such registration after ExpiryHostnameResolvesMinutes.
            Example:

            orchestrator -c register-hostname-unresolve -i instance.fqdn.com --hostname=virtual.name.com

        deregister-hostname-unresolve
            Explicitly deregister/dosassociate a hostname with an "unresolved" name. Orchestrator merely remvoes the association, but does
            not touch any slave at this point. A "repoint" command can be useful right after calling this command to change slave's master host
            name (assumed to be an "unresolved" name, such as a VIP) with the real fqdn of the master host.
            Example:

            orchestrator -c deregister-hostname-unresolve -i instance.fqdn.com

				set-heuristic-domain-instance
						This is a temporary (sync your watches, watch for next ice age) command which registers the cluster domain name of a given cluster
						with the master/writer host for that cluster. It is a one-time-master-discovery operation.
						At this time orchestrator may also act as a small & simple key-value store (recall the "temporary" indication).
						Master failover operations will overwrite the domain instance identity. Orchestrator so turns into a mini master-discovery
						service (I said "TEMPORARY"). Really there are other tools for the job. See also: which-heuristic-domain-instance
						Example:

						orchestrator -c set-heuristic-domain-instance --alias some_alias
								Detects the domain name for given cluster, identifies the writer master of the cluster, associates the two in key-value store

						orchestrator -c set-heuristic-domain-instance -i instance.of.some.cluster
								Cluster is inferred by a member instance (the instance is not necessarily the master)

    Misc commands

        continuous
            Enter continuous mode, and actively poll for instances, diagnose problems, do maintenance etc.
            This type of work is typically done in HTTP mode. However nothing prevents orchestrator from
            doing it in command line. Invoking with "continuous" will run indefinitely. Example:

            orchestrator -c continuous

				active-nodes
						List orchestrator nodes or processes that are actively running or have most recently
						executed. Output is in hostname:token format, where "token" is an internal unique identifier
						of an orchestrator process. Example:

						orchestrator -c active-nodes

				access-token
						When running HTTP with "AuthenticationMethod" : "token", receive a new access token.
						This token must be utilized within "AccessTokenUseExpirySeconds" and can then be used
						until "AccessTokenExpiryMinutes" have passed.
						In "token" authentication method a user is read-only unless able to provide with a fresh token.
						A token may only be used once (two users must get two distinct tokens).
						Submitting a token is done via "/web/access-token?publicToken=<received_token>". The token is then stored
						in HTTP cookie.

						orchestrator -c access-token

        reset-hostname-resolve-cache
            Clear the hostname resolve cache; it will be refilled by following host discoveries

            orchestrator -c reset-hostname-resolve-cache

        resolve
            Utility command to resolve a CNAME and return resolved hostname name. Example:

            orchestrator -c resolve -i cname.to.resolve

        redeploy-internal-db
						Force internal schema migration to current backend structure. Orchestrator keeps track of the deployed
						versions and will not reissue a migration for a version already deployed. Normally you should not use
						this command, and it is provided mostly for building and testing purposes. Nonetheless it is safe to
						use and at most it wastes some cycles.
    `

// main is the application's entry point. It will either spawn a CLI or HTTP itnerfaces.
func main() {
	configFile := flag.String("config", "", "config file name")
	command := flag.String("c", "", "command, required. See full list of commands via 'orchestrator -c help'")
	strict := flag.Bool("strict", false, "strict mode (more checks, slower)")
	instance := flag.String("i", "", "instance, host_fqdn[:port] (e.g. db.company.com:3306, db.company.com)")
	sibling := flag.String("s", "", "sibling instance, host_fqdn[:port]")
	destination := flag.String("d", "", "destination instance, host_fqdn[:port] (synonym to -s)")
	owner := flag.String("owner", "", "operation owner")
	reason := flag.String("reason", "", "operation reason")
	duration := flag.String("duration", "", "maintenance duration (format: 59s, 59m, 23h, 6d, 4w)")
	pattern := flag.String("pattern", "", "regular expression pattern")
	clusterAlias := flag.String("alias", "", "cluster alias")
	pool := flag.String("pool", "", "Pool logical name (applies for pool-related commands)")
	hostnameFlag := flag.String("hostname", "", "Hostname/fqdn/CNAME/VIP (applies for hostname/resolve related commands)")
	discovery := flag.Bool("discovery", true, "auto discovery mode")
	quiet := flag.Bool("quiet", false, "quiet")
	verbose := flag.Bool("verbose", false, "verbose")
	debug := flag.Bool("debug", false, "debug mode (very verbose)")
	stack := flag.Bool("stack", false, "add stack trace upon error")
	config.RuntimeCLIFlags.Databaseless = flag.Bool("databaseless", false, "EXPERIMENTAL! Work without backend database")
	config.RuntimeCLIFlags.SkipUnresolve = flag.Bool("skip-unresolve", false, "Do not unresolve a host name")
	config.RuntimeCLIFlags.SkipUnresolveCheck = flag.Bool("skip-unresolve-check", false, "Skip/ignore checking an unresolve mapping (via hostname_unresolve table) resolves back to same hostname")
	config.RuntimeCLIFlags.Noop = flag.Bool("noop", false, "Dry run; do not perform destructing operations")
	config.RuntimeCLIFlags.BinlogFile = flag.String("binlog", "", "Binary log file name")
	config.RuntimeCLIFlags.Statement = flag.String("statement", "", "Statement/hint")
	config.RuntimeCLIFlags.GrabElection = flag.Bool("grab-election", false, "Grab leadership (only applies to continuous mode)")
	config.RuntimeCLIFlags.PromotionRule = flag.String("promotion-rule", "prefer", "Promotion rule for register-andidate (prefer|neutral|must_not)")
	config.RuntimeCLIFlags.Version = flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *destination != "" && *sibling != "" {
		log.Fatalf("-s and -d are synonyms, yet both were specified. You're probably doing the wrong thing.")
	}
	switch *config.RuntimeCLIFlags.PromotionRule {
	case "prefer", "neutral", "must_not":
		{
			// OK
		}
	default:
		{
			log.Fatalf("-promotion-rule only supports prefer|neutral|must_not")
		}
	}
	if *destination == "" {
		*destination = *sibling
	}

	log.SetLevel(log.ERROR)
	if *verbose {
		log.SetLevel(log.INFO)
	}
	if *debug {
		log.SetLevel(log.DEBUG)
	}
	if *stack {
		log.SetPrintStackTrace(*stack)
	}
	log.Info("starting orchestrator") // FIXME and add the version which is currently in build.sh

	if *config.RuntimeCLIFlags.Version {
		fmt.Println(AppVersion)
		return
	}

	runtime.GOMAXPROCS(math.MinInt(4, runtime.NumCPU()))

	if len(*configFile) > 0 {
		config.ForceRead(*configFile)
	} else {
		config.Read("/etc/orchestrator.conf.json", "conf/orchestrator.conf.json", "orchestrator.conf.json")
	}
	if *config.RuntimeCLIFlags.Databaseless {
		config.Config.DatabaselessMode__experimental = true
	}
	if config.Config.Debug {
		log.SetLevel(log.DEBUG)
	}
	if *quiet {
		// Override!!
		log.SetLevel(log.ERROR)
	}
	if config.Config.EnableSyslog {
		log.EnableSyslogWriter("orchestrator")
		log.SetSyslogLevel(log.INFO)
	}
	if config.Config.AuditToSyslog {
		inst.EnableAuditSyslog()
	}
	config.RuntimeCLIFlags.ConfiguredVersion = AppVersion

	inst.InitializeInstanceDao()

	if len(flag.Args()) == 0 && *command == "" {
		// No command, no argument: just prompt
		fmt.Println(prompt)
		return
	}
	switch {
	case len(flag.Args()) == 0 || flag.Arg(0) == "cli":
		app.CliWrapper(*command, *strict, *instance, *destination, *owner, *reason, *duration, *pattern, *clusterAlias, *pool, *hostnameFlag)
	case flag.Arg(0) == "http":
		app.Http(*discovery)
	default:
		fmt.Fprintln(os.Stderr, `Usage:
  orchestrator --options... [cli|http]
See complete list of commands:
  orchestrator -c help
Full blown documentation:
  orchestrator
`)
		os.Exit(1)
	}
}
