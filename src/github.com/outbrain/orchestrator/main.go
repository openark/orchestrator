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
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/app"
	"github.com/outbrain/orchestrator/config"
)

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

	switch {
	case len(flag.Args()) == 0 || flag.Arg(0) == "cli":
		app.Cli(*command, *strict, *instance, *sibling, *owner, *reason, *pattern)
	case flag.Arg(0) == "http":
		app.Http(*discovery)
	default:
		log.Error("Usage: orchestrator --options... [cli|http]")
	}
}
