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

package logic

import (
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/agent"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
	"time"
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
	go func() {
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
	if err != nil || instance == nil {
		log.Warningf("instance is nil in DiscoverInstance. key=%+v, error=%+v", instanceKey, err)
		goto Cleanup
	}

	log.Debugf("Discovered host: %+v, master: %+v", instance.Key, instance.MasterKey)

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
		case <-pendingTokens:
			<-completedTokens
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
	inst.LoadHostnameResolveCacheFromDatabase()
	go handleDiscoveryRequests(nil, nil)
	tick := time.Tick(time.Duration(config.Config.DiscoveryPollSeconds) * time.Second)
	forgetUnseenTick := time.Tick(time.Minute)
	recoverTick := time.Tick(10 * time.Second)

	var snapshotTopologiesTick <-chan time.Time
	if config.Config.SnapshotTopologiesIntervalHours > 0 {
		snapshotTopologiesTick = time.Tick(time.Duration(config.Config.SnapshotTopologiesIntervalHours) * time.Hour)
	}

	elected := false
	_ = CreateElectionAnchor(false)
	for {
		select {
		case <-tick:
			if elected, _ = AttemptElection(); elected {
				instanceKeys, _ := inst.ReadOutdatedInstanceKeys()
				log.Debugf("outdated keys: %+v", instanceKeys)
				for _, instanceKey := range instanceKeys {
					discoveryInstanceKeys <- instanceKey
				}
			} else {
				log.Debugf("Not elected as active node; polling")
			}
		case <-forgetUnseenTick:
			// See if we should also forget objects (lower frequency)
			go func() {
				if elected {
					inst.ForgetLongUnseenInstances()
					inst.ForgetUnseenInstancesDifferentlyResolved()
					inst.ForgetExpiredHostnameResolves()
					inst.DeleteInvalidHostnameResolves()
					inst.ReviewUnseenInstances()
					inst.InjectUnseenMasters()
					inst.ResolveUnknownMasterHostnameResolves()
					inst.ExpireMaintenance()
					inst.ExpireDowntime()
					inst.ExpireCandidateInstances()
					inst.ExpireHostnameUnresolve()
				}
				if !elected {
					// Take this opportunity to refresh yourself
					inst.LoadHostnameResolveCacheFromDatabase()
				}
				inst.ReadClusterAliases()
				HealthTest()
			}()
		case <-recoverTick:
			go func() {
				if elected {
					ClearActiveFailureDetections()
					ClearActiveRecoveries()
					CheckAndRecover(nil, nil, false)
				}
			}()
		case <-snapshotTopologiesTick:
			go func() {
				inst.SnapshotTopologies()
			}()
		}
	}
}

func pollAgent(hostname string) error {
	polledAgent, err := agent.GetAgent(hostname)
	agent.UpdateAgentLastChecked(hostname)

	if err != nil {
		return log.Errore(err)
	}

	err = agent.UpdateAgentInfo(hostname, polledAgent)
	if err != nil {
		return log.Errore(err)
	}

	return nil
}

// ContinuousAgentsPoll starts an asynchronuous infinite process where agents are
// periodically investigated and their status captured, and long since unseen agents are
// purged and forgotten.
func ContinuousAgentsPoll() {
	log.Infof("Starting continuous agents poll")

	go discoverSeededAgents()

	tick := time.Tick(time.Duration(config.Config.DiscoveryPollSeconds) * time.Second)
	forgetUnseenTick := time.Tick(time.Hour)
	for _ = range tick {
		agentsHosts, _ := agent.ReadOutdatedAgentsHosts()
		log.Debugf("outdated agents hosts: %+v", agentsHosts)
		for _, hostname := range agentsHosts {
			go pollAgent(hostname)
		}
		// See if we should also forget agents (lower frequency)
		select {
		case <-forgetUnseenTick:
			agent.ForgetLongUnseenAgents()
			agent.FailStaleSeeds()
		default:
		}
	}
}

func discoverSeededAgents() {
	for seededAgent := range agent.SeededAgents {
		instanceKey := inst.InstanceKey{Hostname: seededAgent.Hostname, Port: int(seededAgent.MySQLPort)}
		go StartDiscovery(instanceKey)
	}
}
