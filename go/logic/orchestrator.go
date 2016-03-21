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
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/agent"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
	ometrics "github.com/outbrain/orchestrator/go/metrics"
	"github.com/outbrain/orchestrator/go/process"
	"github.com/pmylund/go-cache"
	"github.com/rcrowley/go-metrics"
)

const (
	maxConcurrency = 5
)

// discoveryInstanceKeys is a channel of instanceKey-s that were requested for discovery.
// It can be continuously updated as discovery process progresses.
var discoveryInstanceKeys = make(chan inst.InstanceKey, maxConcurrency)

var discoveriesCounter = metrics.NewCounter()
var failedDiscoveriesCounter = metrics.NewCounter()
var discoveryQueueLengthGauge = metrics.NewGauge()
var discoveryRecentCountGauge = metrics.NewGauge()
var isElectedGauge = metrics.NewGauge()

var isElectedNode int64 = 0

var recentDiscoveryOperationKeys *cache.Cache

func init() {
	metrics.Register("discoveries.attempt", discoveriesCounter)
	metrics.Register("discoveries.fail", failedDiscoveriesCounter)
	metrics.Register("discoveries.queue_length", discoveryQueueLengthGauge)
	metrics.Register("discoveries.recent_count", discoveryRecentCountGauge)
	metrics.Register("elect.is_elected", isElectedGauge)

	ometrics.OnGraphiteTick(func() { discoveryQueueLengthGauge.Update(int64(len(discoveryInstanceKeys))) })
	ometrics.OnGraphiteTick(func() {
		if recentDiscoveryOperationKeys == nil {
			return
		}
		discoveryRecentCountGauge.Update(int64(recentDiscoveryOperationKeys.ItemCount()))
	})
	ometrics.OnGraphiteTick(func() { isElectedGauge.Update(int64(atomic.LoadInt64(&isElectedNode))) })
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
				inst.AuditOperation("reload-configuration", nil, "Triggered via SIGHUP")
			}
		}
	}()
}

// handleDiscoveryRequests iterates the discoveryInstanceKeys channel and calls upon
// instance discovery per entry.
func handleDiscoveryRequests() {
	for instanceKey := range discoveryInstanceKeys {
		// Possibly this used to be the elected node, but has been demoted, while still
		// the queue is full.
		// Just don't process the queue when not elected.
		if atomic.LoadInt64(&isElectedNode) == 1 {
			go discoverInstance(instanceKey)
		} else {
			log.Debugf("Node apparently demoted. Skipping discovery of %+v. Remaining queue size: %+v", instanceKey, len(discoveryInstanceKeys))
		}
	}
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

	log.Debugf("Discovered host: %+v, master: %+v, version: %+v", instance.Key, instance.MasterKey, instance.Version)

	if atomic.LoadInt64(&isElectedNode) == 0 {
		// Maybe this node was elected before, but isn't elected anymore.
		// If not elected, stop drilling up/down the topology
		return
	}

	// Investigate slaves:
	for _, slaveKey := range instance.SlaveHosts.GetInstanceKeys() {
		slaveKey := slaveKey
		if slaveKey.IsValid() {
			discoveryInstanceKeys <- slaveKey
		}
	}
	// Investigate master:
	if instance.MasterKey.IsValid() {
		discoveryInstanceKeys <- instance.MasterKey
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
	recentDiscoveryOperationKeys = cache.New(time.Duration(config.Config.InstancePollSeconds)*time.Second, time.Second)

	inst.LoadHostnameResolveCache()
	go handleDiscoveryRequests()

	discoveryTick := time.Tick(time.Duration(config.Config.GetDiscoveryPollSeconds()) * time.Second)
	instancePollTick := time.Tick(time.Duration(config.Config.InstancePollSeconds) * time.Second)
	caretakingTick := time.Tick(time.Minute)
	recoveryTick := time.Tick(time.Duration(config.Config.RecoveryPollSeconds) * time.Second)
	var snapshotTopologiesTick <-chan time.Time
	if config.Config.SnapshotTopologiesIntervalHours > 0 {
		snapshotTopologiesTick = time.Tick(time.Duration(config.Config.SnapshotTopologiesIntervalHours) * time.Hour)
	}

	go ometrics.InitGraphiteMetrics()
	go acceptSignals()

	if *config.RuntimeCLIFlags.GrabElection {
		process.GrabElection()
	}
	for {
		select {
		case <-discoveryTick:
			go func() {
				wasAlreadyElected := atomic.LoadInt64(&isElectedNode)
				myIsElectedNode, err := process.AttemptElection()
				if err != nil {
					log.Errore(err)
				}
				if myIsElectedNode {
					atomic.StoreInt64(&isElectedNode, 1)
				} else {
					atomic.StoreInt64(&isElectedNode, 0)
				}

				if myIsElectedNode {
					instanceKeys, err := inst.ReadOutdatedInstanceKeys()
					if err != nil {
						log.Errore(err)
					}

					log.Debugf("outdated keys: %+v", instanceKeys)
					for _, instanceKey := range instanceKeys {
						instanceKey := instanceKey

						go func() {
							if instanceKey.IsValid() {
								discoveryInstanceKeys <- instanceKey
							}
						}()
					}
					if wasAlreadyElected == 0 {
						// Just turned to be leader!
						go process.RegisterNode("", "", false)
					}
				} else {
					log.Debugf("Not elected as active node; polling")
				}
			}()
		case <-instancePollTick:
			go func() {
				// This tick does NOT do instance poll (these are handled by the oversmapling discoveryTick)
				// But rather should invoke such routinely operations that need to be as (or roughly as) frequent
				// as instance poll
				if atomic.LoadInt64(&isElectedNode) == 1 {
					go inst.UpdateInstanceRecentRelaylogHistory()
					go inst.RecordInstanceCoordinatesHistory()
				}
			}()
		case <-caretakingTick:
			// Various periodic internal maintenance tasks
			go func() {
				if atomic.LoadInt64(&isElectedNode) == 1 {
					go inst.RecordInstanceBinlogFileHistory()
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
					go process.ExpireNodesHistory()
					go process.ExpireAccessTokens()
				} else {
					// Take this opportunity to refresh yourself
					go inst.LoadHostnameResolveCache()
				}
			}()
		case <-recoveryTick:
			go func() {
				if atomic.LoadInt64(&isElectedNode) == 1 {
					go ClearActiveFailureDetections()
					go ClearActiveRecoveries()
					go ExpireBlockedRecoveries()
					go AcknowledgeCrashedRecoveries()
					go inst.ExpireInstanceAnalysisChangelog()
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

	tick := time.Tick(time.Duration(config.Config.GetDiscoveryPollSeconds()) * time.Second)
	caretakingTick := time.Tick(time.Hour)
	for range tick {
		agentsHosts, _ := agent.ReadOutdatedAgentsHosts()
		log.Debugf("outdated agents hosts: %+v", agentsHosts)
		for _, hostname := range agentsHosts {
			go pollAgent(hostname)
		}
		// See if we should also forget agents (lower frequency)
		select {
		case <-caretakingTick:
			agent.ForgetLongUnseenAgents()
			agent.FailStaleSeeds()
		default:
		}
	}
}

func discoverSeededAgents() {
	for seededAgent := range agent.SeededAgents {
		instanceKey := &inst.InstanceKey{Hostname: seededAgent.Hostname, Port: int(seededAgent.MySQLPort)}
		go inst.ReadTopologyInstance(instanceKey)
	}
}
