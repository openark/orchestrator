package orchestrator

import (
	"fmt"
	"time"
	"github.com/outbrain/orchestrator/inst"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/log"
)

const (
	maxConcurrency = 5
)

// discoveryInstanceKeys is a channel of instanceKey-s that were requested for discovery.
// It can be continuously updated as discovery process progresses.
var discoveryInstanceKeys chan inst.InstanceKey = make(chan inst.InstanceKey, maxConcurrency)


// handleDiscoveryRequests iterates the discoveryInstanceKeys channel and calls upon
// instance discovery per entry.
func handleDiscoveryRequests(pendingTokens chan bool, completedTokens chan bool) {
    for instanceKey := range discoveryInstanceKeys {
        AccountedDiscoverInstance(instanceKey, pendingTokens, completedTokens)
    }
}


// AccountedDiscoverInstance will call upon DiscoverInstance and will keep track of 
// discovery tokens such that management of multiple discoveries can figure out
// whether all instances in a topology are accounted for.
func AccountedDiscoverInstance(instanceKey inst.InstanceKey, pendingTokens chan bool, completedTokens chan bool) {
	if pendingTokens != nil {
		pendingTokens <- true
	}
	go func () {
		DiscoverInstance(instanceKey)
		if completedTokens != nil {
			completedTokens <- true
		}
	}()
}


// DiscoverInstance will attempt discovering an instance (unless it is already up to date) and will
// list down its master and slaves (if any) for further discovery.
func DiscoverInstance(instanceKey inst.InstanceKey) {
	instanceKey.Formalize()
	if !instanceKey.IsValid() {
		return
	}
	
	instance, found, err := inst.ReadInstance(&instanceKey)
	
	if found && instance.IsUpToDate && instance.IsLastCheckValid {
		// we've already discovered this one. Skip!
		goto Cleanup
	}
	// First we've ever heard of this instance. Continue investigation:
	instance, err = inst.ReadTopologyInstance(&instanceKey)
	// panic can occur (IO stuff). Therefore it may happen
	// that instance is nil. Check it.
	if	err	!= nil || instance == nil {goto Cleanup}

	fmt.Printf("host: %+v, master: %+v\n", instance.Key, instance.MasterKey)

	// Investigate slaves:
	for _, slaveKey := range instance.SlaveHosts.GetInstanceKeys() {
		discoveryInstanceKeys <- slaveKey
	}
	// Investigate master:
	discoveryInstanceKeys <- instance.MasterKey
	
	
	Cleanup:
}


// Start discovery begins a one time asynchronuous discovery process for the given
// instance and all of its topology connected instances.
// That is, the instance will be investigated for master and slaves, and the routines will follow on
// each and every such found master/slave.
// In essense, assuming all slaves in a replication topology are running, and given a single instance
// in such topology, this function will detect the entire topology.
func StartDiscovery(instanceKey inst.InstanceKey) {
	log.Infof("Starting discovery at %+v", instanceKey)
	pendingTokens := make(chan bool, maxConcurrency)
	completedTokens := make(chan bool, maxConcurrency)

	AccountedDiscoverInstance(instanceKey, pendingTokens, completedTokens) 
	go handleDiscoveryRequests(pendingTokens, completedTokens)
	
	// Block until all are complete
	for {
		select {
			case <- pendingTokens:
				<- completedTokens
			default:
				inst.AuditOperation("start-discovery", &instanceKey, "")
				return
		}
	}
}

// ContinuousDiscovery starts an asynchronuous infinite discovery process where instances are
// periodically investigated and their status captured, and long since unseen instances are
// purged and forgotten.
func ContinuousDiscovery() {
	log.Infof("Starting continuous discovery")
	go handleDiscoveryRequests(nil, nil)
    tick := time.Tick(time.Duration(config.Config.DiscoveryPollSeconds) * time.Second)
    forgetUnseenTick := time.Tick(time.Hour)
    for _ = range tick {
		instanceKeys, _ := inst.ReadOutdatedInstanceKeys()
		log.Debugf("outdated keys: %+v", instanceKeys)
		for _, instanceKey := range instanceKeys {
			discoveryInstanceKeys <- instanceKey
		}
    	// See if we should also forget instances (lower frequency)
		select {
			case <- forgetUnseenTick:
		    	inst.ForgetLongUnseenInstances()
			default:
		}
	}
}
