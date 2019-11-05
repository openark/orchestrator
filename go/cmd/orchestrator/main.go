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

	"github.com/github/orchestrator/go/app"
	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/inst"
	"github.com/openark/golib/log"
)

var AppVersion, GitCommit string

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
	config.RuntimeCLIFlags.SkipBinlogSearch = flag.Bool("skip-binlog-search", false, "when matching via Pseudo-GTID, only use relay logs. This can save the hassle of searching for a non-existend pseudo-GTID entry, for example in servers with replication filters.")
	config.RuntimeCLIFlags.SkipUnresolve = flag.Bool("skip-unresolve", false, "Do not unresolve a host name")
	config.RuntimeCLIFlags.SkipUnresolveCheck = flag.Bool("skip-unresolve-check", false, "Skip/ignore checking an unresolve mapping (via hostname_unresolve table) resolves back to same hostname")
	config.RuntimeCLIFlags.Noop = flag.Bool("noop", false, "Dry run; do not perform destructing operations")
	config.RuntimeCLIFlags.BinlogFile = flag.String("binlog", "", "Binary log file name")
	config.RuntimeCLIFlags.Statement = flag.String("statement", "", "Statement/hint")
	config.RuntimeCLIFlags.GrabElection = flag.Bool("grab-election", false, "Grab leadership (only applies to continuous mode)")
	config.RuntimeCLIFlags.PromotionRule = flag.String("promotion-rule", "prefer", "Promotion rule for register-andidate (prefer|neutral|prefer_not|must_not)")
	config.RuntimeCLIFlags.Version = flag.Bool("version", false, "Print version and exit")
	config.RuntimeCLIFlags.SkipContinuousRegistration = flag.Bool("skip-continuous-registration", false, "Skip cli commands performaing continuous registration (to reduce orchestratrator backend db load")
	config.RuntimeCLIFlags.EnableDatabaseUpdate = flag.Bool("enable-database-update", false, "Enable database update, overrides SkipOrchestratorDatabaseUpdate")
	config.RuntimeCLIFlags.IgnoreRaftSetup = flag.Bool("ignore-raft-setup", false, "Override RaftEnabled for CLI invocation (CLI by default not allowed for raft setups). NOTE: operations by CLI invocation may not reflect in all raft nodes.")
	config.RuntimeCLIFlags.Tag = flag.String("tag", "", "tag to add ('tagname' or 'tagname=tagvalue') or to search ('tagname' or 'tagname=tagvalue' or comma separated 'tag0,tag1=val1,tag2' for intersection of all)")
	config.RuntimeCLIFlags.CommandPayload = flag.String("command-payload", "", "CommandPayload")
	flag.Parse()

	if *destination != "" && *sibling != "" {
		log.Fatalf("-s and -d are synonyms, yet both were specified. You're probably doing the wrong thing.")
	}
	switch *config.RuntimeCLIFlags.PromotionRule {
	case "prefer", "neutral", "prefer_not", "must_not":
		{
			// OK
		}
	default:
		{
			log.Fatalf("-promotion-rule only supports prefer|neutral|prefer_not|must_not")
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
	if *config.RuntimeCLIFlags.Version {
		fmt.Println(AppVersion)
		fmt.Println(GitCommit)
		return
	}

	startText := "starting orchestrator"
	if AppVersion != "" {
		startText += ", version: " + AppVersion
	}
	if GitCommit != "" {
		startText += ", git commit: " + GitCommit
	}
	log.Info(startText)

	if len(*configFile) > 0 {
		config.ForceRead(*configFile)
	} else {
		config.Read("/etc/orchestrator.conf.json", "conf/orchestrator.conf.json", "orchestrator.conf.json")
	}
	if *config.RuntimeCLIFlags.EnableDatabaseUpdate {
		config.Config.SkipOrchestratorDatabaseUpdate = false
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
	config.MarkConfigurationLoaded()

	if len(flag.Args()) == 0 && *command == "" {
		// No command, no argument: just prompt
		fmt.Println(app.AppPrompt)
		return
	}

	helpTopic := ""
	if flag.Arg(0) == "help" {
		if flag.Arg(1) != "" {
			helpTopic = flag.Arg(1)
		}
		if helpTopic == "" {
			helpTopic = *command
		}
		if helpTopic == "" {
			// hacky way to make the CLI kick in as if the user typed `orchestrator -c help cli`
			*command = "help"
			flag.Args()[0] = "cli"
		}
	}

	switch {
	case helpTopic != "":
		app.HelpCommand(helpTopic)
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
  orchestrator`)
		os.Exit(1)
	}
}
