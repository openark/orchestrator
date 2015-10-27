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
	"github.com/outbrain/golib/math"
	"github.com/outbrain/orchestrator/go/agent"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
	ometrics "github.com/outbrain/orchestrator/go/metrics"
	"github.com/outbrain/orchestrator/go/process"
	"github.com/pmylund/go-cache"
	"github.com/rcrowley/go-metrics"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxConcurrency = 5
)

// discoveryInstanceKeys is a channel of instanceKey-s that were requested for discovery.
// It can be continuously updated as discovery process progresses.
var discoveryInstanceKeys chan inst.InstanceKey = make(chan inst.InstanceKey, maxConcurrency)

var discoveriesCounter = metrics.NewCounter()
var failedDiscoveriesCounter = metrics.NewCounter()
var discoveryQueueLengthGauge = metrics.NewGauge()
var discoveryRecentCountGauge = metrics.NewGauge()
var isElectedGauge = metrics.NewGauge()

var isElectedNode = false

var recentDiscoveryOperationKeys = cache.New(time.Duration(config.Config.DiscoveryPollSeconds/2)*time.Second, time.Second)

func init() {
	isElectedNode = false

	metrics.Register("discoveries.attempt", discoveriesCounter)
	metrics.Register("discoveries.fail", failedDiscoveriesCounter)
	metrics.Register("discoveries.queue_length", discoveryQueueLengthGauge)
	metrics.Register("discoveries.recent_count", discoveryRecentCountGauge)
	metrics.Register("elect.is_elected", isElectedGauge)

	ometrics.OnGraphiteTick(func() { discoveryQueueLengthGauge.Update(int64(len(discoveryInstanceKeys))) })
	ometrics.OnGraphiteTick(func() { discoveryRecentCountGauge.Update(int64(recentDiscoveryOperationKeys.ItemCount())) })
	ometrics.OnGraphiteTick(func() { isElectedGauge.Update(int64(math.TernaryInt(isElectedNode, 1, 0))) })
}

// acceptSignals registers for OS signals
func acceptSignals() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGHUP:
				log.Debugf("Received SIGHUP. Reloading configuration")
				config.Reload()
			}
		}
	}()
}

// handleDiscoveryRequests iterates the discoveryInstanceKeys channel and calls upon
// instance discovery per entry.
func handleDiscoveryRequests(pendingTokens chan bool, completedTokens chan bool) {
	for instanceKey := range discoveryInstanceKeys {
		accountedDiscoverInstance(instanceKey, pendingTokens, completedTokens)
	}
}

// accountedDiscoverInstance will call upon DiscoverInstance and will keep track of
// discovery tokens such that management of multiple discoveries can figure out
// whether all instances in a topology are accounted for.
func accountedDiscoverInstance(instanceKey inst.InstanceKey, pendingTokens chan bool, completedTokens chan bool) {
	if pendingTokens != nil {
		pendingTokens <- true
	}
	go func() {
		discoverInstance(instanceKey)
		if completedTokens != nil {
			completedTokens <- true
		}
	}()
}

// discoverInstance will attempt discovering an instance (unless it is already up to date) and will
// list down its master and slaves (if any) for further discovery.
func discoverInstance(instanceKey inst.InstanceKey) {
	instanceKey.Formalize()
	if !instanceKey.IsValid() {
		return
	}

	if existsInCacheError := recentDiscoveryOperationKeys.Add(instanceKey.DisplayString(), true, cache.DefaultExpiration); existsInCacheError != nil {
		// Just recently attempted
		return
	}

	instance, found, err := inst.ReadInstance(&instanceKey)

	if found && instance.IsUpToDate && instance.IsLastCheckValid {
		// we've already discovered this one. Skip!
		return
	}
	discoveriesCounter.Inc(1)
	// First we've ever heard of this instance. Continue investigation:
	instance, err = inst.ReadTopologyInstance(&instanceKey)
	// panic can occur (IO stuff). Therefore it may happen
	// that instance is nil. Check it.
	if instance == nil {
		failedDiscoveriesCounter.Inc(1)
		log.Warningf("instance is nil in discoverInstance. key=%+v, error=%+v", instanceKey, err)
		return
	}

	log.Debugf("Discovered host: %+v, master: %+v", instance.Key, instance.MasterKey)

	if !isElectedNode {
		// Maybe this node was elected before, but isn't elected anymore.
		// If not elected, stop drilling down to further investigate slaves.
		return
	}

	// Investigate slaves:
	for _, slaveKey := range instance.SlaveHosts.GetInstanceKeys() {
		discoveryInstanceKeys <- slaveKey
	}
	// Investigate master:
	discoveryInstanceKeys <- instance.MasterKey
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

	accountedDiscoverInstance(instanceKey, pendingTokens, completedTokens)
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
	if config.Config.DatabaselessMode__experimental {
		log.Fatal("Cannot execute continuous mode in databaseless mode")
	}

	log.Infof("Starting continuous discovery")
	inst.LoadHostnameResolveCacheFromDatabase()
	go handleDiscoveryRequests(nil, nil)
	tick := time.Tick(time.Duration(config.Config.DiscoveryPollSeconds) * time.Second)
	forgetUnseenTick := time.Tick(time.Minute)
	recoverTick := time.Tick(time.Duration(config.Config.RecoveryPollSeconds) * time.Second)

	var snapshotTopologiesTick <-chan time.Time
	if config.Config.SnapshotTopologiesIntervalHours > 0 {
		snapshotTopologiesTick = time.Tick(time.Duration(config.Config.SnapshotTopologiesIntervalHours) * time.Hour)
	}

	go ometrics.InitGraphiteMetrics()
	go acceptSignals()

	for {
		select {
		case <-tick:
			go func() {
				if isElectedNode, _ = process.AttemptElection(); isElectedNode {
					instanceKeys, _ := inst.ReadOutdatedInstanceKeys()
					log.Debugf("outdated keys: %+v", instanceKeys)
					for _, instanceKey := range instanceKeys {
						discoveryInstanceKeys <- instanceKey
					}
				} else {
					log.Debugf("Not elected as active node; polling")
				}
			}()
		case <-forgetUnseenTick:
			// See if we should also forget objects (lower frequency)
			go func() {
				if isElectedNode {
					go inst.ForgetLongUnseenInstances()
					go inst.ForgetUnseenInstancesDifferentlyResolved()
					go inst.ForgetExpiredHostnameResolves()
					go inst.DeleteInvalidHostnameResolves()
					go inst.ReviewUnseenInstances()
					go inst.InjectUnseenMasters()
					go inst.ResolveUnknownMasterHostnameResolves()
					go inst.UpdateClusterAliases()
					go inst.ExpireMaintenance()
					go inst.ExpireDowntime()
					go inst.ExpireCandidateInstances()
					go inst.ExpireHostnameUnresolve()
					go inst.ExpireClusterDomainName()
					go inst.ExpireAudit()
					go inst.ExpireMasterPositionEquivalence()
					go inst.FlushNontrivialResolveCacheToDatabase()
				}
				if !isElectedNode {
					// Take this opportunity to refresh yourself
					inst.LoadHostnameResolveCacheFromDatabase()
				}
				go inst.ReadClusterAliases()
				go process.HealthTest()
			}()
		case <-recoverTick:
			go func() {
				if isElectedNode {
					go ClearActiveFailureDetections()
					go ClearActiveRecoveries()
					go ExpireBlockedRecoveries()
					go CheckAndRecover(nil, nil, false)
				}
			}()
		case <-snapshotTopologiesTick:
			go func() {
				go inst.SnapshotTopologies()
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
	for range tick {
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
