package orchestrator 

import (
	"github.com/outbrain/inst"
	"github.com/outbrain/orch"
	
	"github.com/outbrain/log"
)


func Cli(command string, instance string, sibling string) {
	
	instanceKey := inst.ParseInstanceKey(instance)
	siblingKey := inst.ParseInstanceKey(sibling)
	
	if len(command) == 0 {
		log.Fatal("expected command (-c) (discover|forget|continuous|move-up|move-below)")
	}
	switch command {
		case "move-up": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", instance)}
			_, err := inst.MoveUp(instanceKey)
			if err != nil {log.Errore( err)}
		}
		case "move-below": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", instance)}
			if siblingKey == nil {log.Fatal("Cannot deduce sibling:", sibling)}
			_, err := inst.MoveBelow(instanceKey, siblingKey)
			if err != nil {log.Errore(err)}
		}
		case "discover": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", instance)}
			orchestrator.StartDiscovery(*instanceKey)
		}
		case "forget": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", instance)}
			inst.ForgetInstance(instanceKey)
		}
		case "continuous": {
			orchestrator.ContinuousDiscovery()
		}
		default: log.Fatal("Unknown command:", command) 
	}
}
