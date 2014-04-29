package orchestrator

import (
	"fmt"
	"time"
	"github.com/outbrain/inst"
	"github.com/outbrain/config"
	"github.com/outbrain/log"
)

const (
	maxConcurrency = 5
)
var discoveryInstanceKeys chan inst.InstanceKey = make(chan inst.InstanceKey, maxConcurrency)


func handleDiscoveryRequests(pendingTokens chan bool, completedTokens chan bool) {
    for instanceKey := range discoveryInstanceKeys {
        AccountedDiscoverInstance(instanceKey, pendingTokens, completedTokens)
    }
}


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

func DiscoverInstance(instanceKey inst.InstanceKey) {
	instanceKey.Formalize()
	if !instanceKey.IsValid() {
		return
	}
	
	instance, found, err := inst.ReadInstance(&instanceKey)
	
	if found && instance.IsUpToDate && instance.IsLastSeenValid {
		// we've already discovered this one. Skip!
		goto Cleanup
	}
	// First we've ever heard of this instance. Continue investigation:
	instance, err = inst.ReadTopologyInstance(&instanceKey)
	if	err	!=	nil	{goto Cleanup}

	fmt.Printf("host: %+v, master: %+v\n", instance.Key, *instance.GetMasterInstanceKey())

	// Investigate slaves:
	for _, slaveKey := range instance.GetSlaveInstanceKeys() {
		discoveryInstanceKeys <- slaveKey
	}
	// Investigate master:
	discoveryInstanceKeys <- *instance.GetMasterInstanceKey()
	
	
	Cleanup:
}


func StartDiscovery(instanceKey inst.InstanceKey) {
	log.Infof("Starting discovery at %+v", instanceKey)
	pendingTokens := make(chan bool, 5)
	completedTokens := make(chan bool, 5)

	AccountedDiscoverInstance(instanceKey, pendingTokens, completedTokens) 
	go handleDiscoveryRequests(pendingTokens, completedTokens)
	
	// Block until all are complete
	for {
		select {
			case <- pendingTokens:
				<- completedTokens
			default:
				return
		}
	}
}


func ContinuousDiscovery() {
	log.Infof("Starting continuous discovery")
	go handleDiscoveryRequests(nil, nil)
    tick := time.Tick(time.Duration(config.Config.DiscoveryPollSeconds) * time.Second)
    for _ = range tick {
		instanceKeys, _ := inst.ReadOutdatedInstanceKeys()
		log.Debugf("outdated keys: %+v", instanceKeys)
		for _, instanceKey := range instanceKeys {
			discoveryInstanceKeys <- instanceKey
		}
	}
}
