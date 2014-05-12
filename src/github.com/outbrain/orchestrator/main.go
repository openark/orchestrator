package main 

import (
	"flag"
	"github.com/outbrain/orchestrator/app"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/log"
)

func main() {
	configFile := flag.String("config", "", "config file name")
	command := flag.String("c", "", "command (move-up|make-child-of)")
	instance := flag.String("i", "", "instance, host:port")
	sibling := flag.String("s", "", "sibling instance, host:port")
	owner := flag.String("owner", "", "operation owner")
	reason := flag.String("reason", "", "operation reason")
	discovery := flag.Bool("discovery", false, "auto discovery mode")
	verbose := flag.Bool("verbose", false, "verbose")
	debug := flag.Bool("debug", false, "debug mode (very verbose)")
	flag.Parse();
	
	log.SetLevel(log.ERROR)
	if *verbose {
		log.SetLevel(log.INFO)
	}
	if *debug {
		log.SetLevel(log.DEBUG)
	}

	log.Info("starting")

	if len(*configFile) > 0 {
		config.ForceRead(*configFile)
	} else {
		config.Read("/etc/orchestrator.conf.json", "conf/orchestrator.conf.json", "orchestrator.conf.json")
	}

	switch {
		case len(flag.Args()) == 0 || flag.Arg(0) == "cli": 
			app.Cli(*command, *instance, *sibling, *owner, *reason)
		case flag.Arg(0) == "http": 
			app.Http(*discovery)
		default:
			log.Error("Usage: orchestrator --options... [cli|http]")
	}
}
