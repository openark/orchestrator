package app 

import (
	"fmt"
	"os/user"
	"github.com/outbrain/orchestrator/inst"	
	"github.com/outbrain/orchestrator/logic"
	"github.com/outbrain/log"
)


// Cli initiates a command line interface, executing requested command.
func Cli(command string, instance string, sibling string, owner string, reason string) {
	
	instanceKey, err := inst.ParseInstanceKey(instance)
	// There always has to be an instance; all commands require one.
	if err != nil {log.Fatal("Cannot deduce instance:", instance)}
	siblingKey, err := inst.ParseInstanceKey(sibling)
	if err != nil {siblingKey = nil}

	if len(owner) == 0 {
		// get os username as owner
		usr, err := user.Current()
    	if err != nil {
        	log.Fatale(err)
    	}
    	owner = usr.Username
	}
		
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
		case "begin-maintenance": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", instance)}
			if owner == "" {log.Fatal("--owner option required")}
			if reason == "" {log.Fatal("--reason option required")}
			maintenanceKey, err := inst.BeginMaintenance(instanceKey, owner, reason)
			if err == nil {log.Infof("Maintenance key: %+v", maintenanceKey)}
			if err != nil {log.Errore(err)}
		}
		case "end-maintenance": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", instance)}
			err := inst.EndMaintenanceByInstanceKey(instanceKey)
			if err != nil {log.Errore(err)}
		}
		case "topology": {
			if instanceKey == nil {log.Fatal("Cannot deduce instance:", instance)}
			output, err := inst.AsciiTopology(instance)
			if err != nil {
				log.Errore(err)
			} else {
				fmt.Println(output)
			}
		}
		case "continuous": {
			orchestrator.ContinuousDiscovery()
		}
		default: log.Fatal("Unknown command:", command) 
	}
}
