package main 

import (
	"flag"
	"log"
	"github.com/outbrain/config"
	"github.com/outbrain/inst"
	"github.com/outbrain/orch"
)

func main() {
	
	configFile := flag.String("config", "", "config file name")
	command := flag.String("c", "", "command (move-up|make-child-of)")
	instance := flag.String("i", "", "instance, host:port")
	sibling := flag.String("s", "", "sibling instance, host:port")
	flag.Parse();
	if len(flag.Args()) > 0 {
		log.Fatal("Unexpected arguments:", flag.Args()) 
	}

	if len(*configFile) > 0 {
		config.ForceRead(*configFile)
	} else {
		config.Read("/etc/orchestrator.conf.json", "conf/orchestrator.conf.json", "orchestrator.conf.json")
	}

	instanceKey := inst.ParseInstanceKey(*instance)
	siblingKey := inst.ParseInstanceKey(*sibling)
	
	if len(*command) == 0 {
		log.Fatal("expected command (-c)")
	}
	switch *command {
		case "move-up": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", *instance)}
			_, err := inst.MoveUp(instanceKey)
			if err != nil {log.Println(err)}
		}
		case "move-below": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", *instance)}
			if siblingKey == nil {log.Fatal("Cannot deduce sibling:", *sibling)}
			_, err := inst.MoveBelow(instanceKey, siblingKey)
			if err != nil {log.Println(err)}
		}
		case "discover": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", *instance)}
			orchestrator.StartDiscovery(*instanceKey)
		}
		default: log.Fatal("Unknown command:", *command) 
	}
}
