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
		fmt.Println(app.AppPrompt)
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
