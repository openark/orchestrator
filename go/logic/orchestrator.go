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

	"github.com/github/orchestrator/go/agent"
	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/discovery"
	"github.com/github/orchestrator/go/inst"
	ometrics "github.com/github/orchestrator/go/metrics"
	"github.com/github/orchestrator/go/process"
	"github.com/outbrain/golib/log"
	"github.com/patrickmn/go-cache"
	"github.com/rcrowley/go-metrics"
	"github.com/sjmudd/stopwatch"
)

// discoveryQueue is a channel of deduplicated instanceKey-s
// that were requested for discovery.  It can be continuously updated
// as discovery process progresses.
var discoveryQueue *discovery.Queue

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

	ometrics.OnGraphiteTick(func() { discoveryQueueLengthGauge.Update(int64(discoveryQueue.Len())) })
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
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGHUP:
				log.Debugf("Received SIGHUP. Reloading configuration")
				config.Reload()
				discovery.MC.Reload()
				inst.AuditOperation("reload-configuration", nil, "Triggered via SIGHUP")
			case syscall.SIGTERM:
				log.Debugf("Received SIGTERM. Shutting down orchestrator")
				discovery.MC.Shutdown()
				// probably should poke other go routines to stop cleanly here ...
				inst.AuditOperation("shutdown", nil, "Triggered via SIGTERM")
				os.Exit(0)
			}
		}
	}()
}

// handleDiscoveryRequests iterates the discoveryQueue channel and calls upon
// instance discovery per entry.
func handleDiscoveryRequests() {
	discoveryQueue = discovery.NewQueue()

	// create a pool of discovery workers
	for i := uint(0); i < config.Config.DiscoveryMaxConcurrency; i++ {
		go func() {
			for {
				instanceKey := discoveryQueue.Consume()
				// Possibly this used to be the elected node, but has
				// been demoted, while still the queue is full.
				if atomic.LoadInt64(&isElectedNode) != 1 {
					log.Debugf("Node apparently demoted. Skipping discovery of %+v. "+
						"Remaining queue size: %+v", instanceKey, discoveryQueue.Len())

					discoveryQueue.Release(instanceKey)

					continue
				}

				discoverInstance(instanceKey)

				discoveryQueue.Release(instanceKey)
			}
		}()
	}
}

// discoverInstance will attempt to discover (poll) an instance (unless
// it is already up to date) and will also ensure that its master and
// replicas (if any) are also checked.
func discoverInstance(instanceKey inst.InstanceKey) {
	start := time.Now()
	defer func() {
		discoveryTime := time.Since(start)
		if discoveryTime > time.Duration(config.Config.InstancePollSeconds)*time.Second {
			log.Warningf("discoverInstance for key %v took %.4fs", instanceKey, discoveryTime.Seconds())
		}
	}()

	latency := stopwatch.NewNamedStopwatch()
	latency.Add("total")
	latency.Add("instance")
	latency.Add("backend")
	latency.Start("total")

	instanceKey.Formalize()
	if !instanceKey.IsValid() {
		return
	}

	if existsInCacheError := recentDiscoveryOperationKeys.Add(instanceKey.DisplayString(), true, cache.DefaultExpiration); existsInCacheError != nil {
		// Just recently attempted
		return
	}

	latency.Start("backend")
	instance, found, err := inst.ReadInstance(&instanceKey)
	latency.Stop("backend")
	if found && instance.IsUpToDate && instance.IsLastCheckValid {
		// we've already discovered this one. Skip!
		return
	}

	discoveriesCounter.Inc(1)

	// First we've ever heard of this instance. Continue investigation:
	instance, err = inst.ReadTopologyInstanceBufferable(&instanceKey, config.Config.BufferInstanceWrites, latency)
	// panic can occur (IO stuff). Therefore it may happen
	// that instance is nil. Check it.
	if instance == nil {
		failedDiscoveriesCounter.Inc(1)
		discovery.MC.Append(&discovery.Metric{
			Timestamp:       time.Now(),
			InstanceKey:     instanceKey,
			TotalLatency:    latency.Elapsed("total"),
			BackendLatency:  latency.Elapsed("backend"),
			InstanceLatency: latency.Elapsed("instance"),
			Err:             err,
		})
		log.Warningf("discoverInstance(%+v) instance is nil in %.3fs (Backend: %.3fs, Instance: %.3fs), error=%+v",
			instanceKey,
			latency.ElapsedSeconds("total"),
			latency.ElapsedSeconds("backend"),
			latency.ElapsedSeconds("instance"),
			err)
		return
	}

	discovery.MC.Append(&discovery.Metric{
		Timestamp:       time.Now(),
		InstanceKey:     instanceKey,
		TotalLatency:    latency.Elapsed("total"),
		BackendLatency:  latency.Elapsed("backend"),
		InstanceLatency: latency.Elapsed("instance"),
		Err:             nil,
	})
	log.Debugf("Discovered host: %+v, master: %+v, version: %+v in %.3fs (Backend: %.3fs, Instance: %.3fs)",
		instance.Key,
		instance.MasterKey,
		instance.Version,
		latency.ElapsedSeconds("total"),
		latency.ElapsedSeconds("backend"),
		latency.ElapsedSeconds("instance"))

	if atomic.LoadInt64(&isElectedNode) == 0 {
		// Maybe this node was elected before, but isn't elected anymore.
		// If not elected, stop drilling up/down the topology
		return
	}

	// Investigate replicas:
	for _, replicaKey := range instance.SlaveHosts.GetInstanceKeys() {
		replicaKey := replicaKey
		if replicaKey.IsValid() {
			discoveryQueue.Push(replicaKey)
		}
	}
	// Investigate master:
	if instance.MasterKey.IsValid() {
		discoveryQueue.Push(instance.MasterKey)
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

	// record discovery Metrics for the time shown, after which they get thrown away.
	// metrics are available via an API call.
	discovery.MC = discovery.NewMetricCollection(time.Duration(config.Config.DiscoveryCollectionRetentionSeconds) * time.Second)

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

					if len(instanceKeys) > config.Config.MaxOutdatedKeysToShow {
						log.Debugf("polling %d outdated keys", len(instanceKeys))
					} else {
						log.Debugf("outdated keys: %+v", instanceKeys)
					}

					for _, instanceKey := range instanceKeys {
						instanceKey := instanceKey

						if instanceKey.IsValid() {
							discoveryQueue.Push(instanceKey)
						}
					}
					if wasAlreadyElected == 0 {
						// Just turned to be leader!
						go process.RegisterNode("", "", false)
					}
				} else {
					hostname, _, _, err := process.ElectedNode()
					if err == nil {
						log.Debugf("Not elected as active node; active node: %v; polling", hostname)
					} else {
						log.Debugf("Not elected as active node; active node: Unable to determine: %v; polling", err)
					}
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
					go inst.ExpirePoolInstances()
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
