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
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/app"
	"github.com/outbrain/orchestrator/config"
)

const prompt string = `
orchestrator [-c command] [-i instance] [--verbose|--debug] [... cli ] | http

-i (instance): 
	instance on which to operate, in "hostname" or "hostname:port" format.
	Default port is 3306 (or DefaultInstancePort in config)
	For some commands this argument can be ommitted altogether, and the
	value is implicitly the local hostname.
-s (Sibling/Subinstance/deStination)
	associated instance. Meaning depends on specific command.
	
-c (command):			

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
			Example:
			
			orchestrator -c move-up-slaves -i slave.whose.subslaves.will.move.up.com[:3306]

		move-below
			Moves a slave beneath its sibling. Both slaves must be actively replicating from same master.
			The sibling will become instance's master. No action taken when sibling cannot act as master 
			(e.g. has no binary logs, is of incompatible version, incompatible binlog format etc.)
			Example:
			
			orchestrator -c move-below -i slave.to.move.com -s sibling.slave.under.which.to.move.com

			orchestrator -c move-below -s sibling.slave.under.which.to.move.com
				-i not given, implicitly assumed local hostname
			
		enslave-siblings
			Turn all siblings of a slave into its sub-slaves. No action taken for siblings that cannot become
			slaves of given instance (e.g. incompatible versions, binlog format etc.). This is a (faster) shortcut
			to executing move-below for all siblings of the given instance. Example:
			
			orchestrator -c enslave-siblings -i slave.whose.siblings.will.move.below.com
			
		repoint
			Make the given instance replicate from another instance without changing the binglog coordinates. There
			are little sanity checks to this and this is a risky operation. Use cases are: a rename of the master's 
			host, a corruption in relay-logs, move from beneath MaxScale & Binlog-server. Examples:
			
			orchestrator -c repoint -i slave.to.operate.on.com -s new.master.com
			
			orchestrator -c repoint -i slave.to.operate.on.com
				The above will repoint the slave back to its existing master without change 
			
			orchestrator -c repoint
				-i not given, implicitly assumed local hostname
			
		make-co-master
			Create a master-master replication. Given instance is a slave which replicates directly from a master.
			The master is then turned to be a slave of the instance. The master is expected to not be a slave.
			The read_only property of the slve is unaffected by this operation. Examples:
	
			orchestrator -c make-co-master -i slave.to.turn.into.co.master.com
			
			orchestrator -c make-co-master
				-i not given, implicitly assumed local hostname
			

	Topology refactoring using Pseudo-GTID
		These operations require that the topology's master is periodically injected with pseudo-GTID,
		and that the PseudoGTIDPattern configuration is setup accordingly. Also consider setting 
		DetectPseudoGTIDQuery.
		

		match-below
		match-up
		rematch
		multi-match-slaves
		match-up-slaves
		regroup-slaves

	get-candidate-slave
	last-pseudo-gtid

	General replication commands
		These commands issue various statements that relate to replication.
		stop-slave
			Issues a STOP SLAVE; command. Example:

			orchestrator -c stop-slave -i slave.to.be.stopped.com
			
		start-slave
			Issues a START SLAVE; command. Example:

			orchestrator -c start-slave -i slave.to.be.started.com
			
		skip-query
			On a failed replicating slave, skips a single query and attempts to resume replication.
			Only applies when the replication seems to be broken on SQL thread (e.g. on duplicate
			key error). Example:

			orchestrator -c skip-query -i slave.with.broken.sql.thread.com
			
		reset-slave
			Issues a RESET SLAVE command. Destructive to replication. Example:

			orchestrator -c reset-slave -i slave.to.reset.com
			
		detach-slave
			Stops replication and modified binlog position into an impossible, yet reversible, value.
			This effectively means the replication becomes broken. See reattach-slave. Example:
			
			orchestrator -c detach-slave -i slave.whose.replication.will.break.com
			
			Issuing this on an already detached slave will do nothing.
			
		reattach-slave
			Undo a detahc-slave operation. Reverses the binlog change into the original values, and 
			resumes replication. Example:
			
			orchestrator -c reattach-slave -i detahced.slave.whose.replication.will.amend.com

			Issuing this on an attached (i.e. normal) slave will do nothing.
	
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
			
	
	Informational commands
		These commands provide information about topologies, replication connections, or otherwise orchstrator's
		"inventory".
		
		find
			Find instances whose hostname matches given regex pattern. Example:
			
			orchestrator -c find -pattern "backup.*us-east"
			
		clusters
			List all clusters known to orchestrator. A cluster (aka topology, aka chain) is identified by its
			master (or one of its master if more than one exists). Example:
			
			orchesrtator -c clusters
				-i not given, implicitly assumed local hostname
			
		topology
			Show an ascii-graph of a replication topology, given a member of that topology. Example:
			
			orchestrator -c topology -i instance.belonging.to.a.topology
			
			orchestrator -c topology
			
			Instance must be already known to orchestrator. Topology is generated by orchestrator's mapping
			and not from synchronuous investigation of the instances. The generated topology may include
			instances that are dead, or whose replication is broken.
			
		which-instance
	which-cluster
	which-cluster-instances
	which-master
	which-slaves
	
	discover
	forget
	
	begin-maintenance
	end-maintenance
	instance-status
	
	replication-analysis
	recover
	
	continuous
	resolve
	`

// main is the application's entry point. It will either spawn a CLI or HTTP itnerfaces.
func main() {
	configFile := flag.String("config", "", "config file name")
	command := flag.String("c", "", "command (discover|forget|continuous|move-up|move-below|begin-maintenance|end-maintenance|clusters|topology)")
	strict := flag.Bool("strict", false, "strict mode (more checks, slower)")
	instance := flag.String("i", "", "instance, host:port")
	sibling := flag.String("s", "", "sibling instance, host:port")
	owner := flag.String("owner", "", "operation owner")
	reason := flag.String("reason", "", "operation reason")
	pattern := flag.String("pattern", "", "regular expression pattern")
	discovery := flag.Bool("discovery", true, "auto discovery mode")
	verbose := flag.Bool("verbose", false, "verbose")
	debug := flag.Bool("debug", false, "debug mode (very verbose)")
	stack := flag.Bool("stack", false, "add stack trace upon error")
	flag.Parse()

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

	log.Info("starting")

	if len(*configFile) > 0 {
		config.ForceRead(*configFile)
	} else {
		config.Read("/etc/orchestrator.conf.json", "conf/orchestrator.conf.json", "orchestrator.conf.json")
	}
	if config.Config.Debug {
		log.SetLevel(log.DEBUG)
	}

	if len(flag.Args()) == 0 && *command == "" {
		// No command, no argument: just prompt
		fmt.Println(prompt)
		return
	}

	switch {
	case len(flag.Args()) == 0 || flag.Arg(0) == "cli":
		app.Cli(*command, *strict, *instance, *sibling, *owner, *reason, *pattern)
	case flag.Arg(0) == "http":
		app.Http(*discovery)
	default:
		log.Error("Usage: orchestrator --options... [cli|http]")
	}
}
