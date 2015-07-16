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

package app

import (
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/util"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
	"github.com/outbrain/orchestrator/go/logic"
	"net"
	"os"
	"os/user"
	"strings"
)

var thisInstanceKey *inst.InstanceKey
var knownCommands []string

func cliCommand(command string) string {
	knownCommands = append(knownCommands, command)
	return command
}

func getInstanceKey(instanceKey *inst.InstanceKey) *inst.InstanceKey {
	if instanceKey == nil {
		instanceKey = thisInstanceKey
	}
	if instanceKey == nil {
		log.Fatal("Cannot deduce instance key")
	}
	return instanceKey
}

func getClusterName(clusterAlias string, instanceKey *inst.InstanceKey) (clusterName string) {
	var err error
	if clusterAlias != "" {
		clusterName, err = inst.ReadClusterByAlias(clusterAlias)
		if err != nil {
			log.Fatale(err)
		}
	} else {
		// deduce cluster by instance
		if instanceKey == nil {
			instanceKey = thisInstanceKey
		}
		if instanceKey == nil {
			log.Fatalf("Unable to get cluster instances: unresolved instance")
		}
		instance, _, err := inst.ReadInstance(instanceKey)
		if err != nil {
			log.Fatale(err)
		}
		if instance == nil {
			log.Fatalf("Instance not found: %+v", *instanceKey)
		}
		clusterName = instance.ClusterName
	}
	if clusterName == "" {
		log.Fatalf("Unable to determine cluster name")
	}
	return clusterName
}

// Cli initiates a command line interface, executing requested command.
func Cli(command string, strict bool, instance string, sibling string, owner string, reason string, duration string, pattern string, clusterAlias string, pool string, hostnameFlag string) {

	if instance != "" && !strings.Contains(instance, ":") {
		instance = fmt.Sprintf("%s:%d", instance, config.Config.DefaultInstancePort)
	}
	instanceKey, err := inst.ParseInstanceKey(instance)
	if err != nil {
		instanceKey = nil
	}
	rawInstanceKey, err := inst.NewRawInstanceKey(instance)
	if err != nil {
		rawInstanceKey = nil
	}

	if sibling != "" && !strings.Contains(sibling, ":") {
		sibling = fmt.Sprintf("%s:%d", sibling, config.Config.DefaultInstancePort)
	}
	siblingKey, err := inst.ParseInstanceKey(sibling)
	if err != nil {
		siblingKey = nil
	}

	if hostname, err := os.Hostname(); err == nil {
		thisInstanceKey = &inst.InstanceKey{Hostname: hostname, Port: int(config.Config.DefaultInstancePort)}
	}

	if len(owner) == 0 {
		// get os username as owner
		usr, err := user.Current()
		if err != nil {
			log.Fatale(err)
		}
		owner = usr.Username
	}
	inst.SetMaintenanceOwner(owner)

	switch command {
	case cliCommand("skip-query"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.SkipQuery(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("move-up"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			instance, err := inst.MoveUp(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(fmt.Sprintf("%s<%s", instanceKey.DisplayString(), instance.MasterKey.DisplayString()))
		}
	case cliCommand("move-up-slaves"):
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}

			movedSlaves, _, err, errs := inst.MoveUpSlaves(instanceKey, pattern)
			if err != nil {
				log.Fatale(err)
			} else {
				for _, e := range errs {
					log.Errore(e)
				}
				for _, slave := range movedSlaves {
					fmt.Println(slave.Key.DisplayString())
				}
			}
		}
	case cliCommand("move-below"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			if siblingKey == nil {
				log.Fatal("Cannot deduce sibling:", sibling)
			}
			_, err := inst.MoveBelow(instanceKey, siblingKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(fmt.Sprintf("%s<%s", instanceKey.DisplayString(), siblingKey.DisplayString()))
		}
	case cliCommand("repoint"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			// siblingKey can be null, in which case the instance repoints to its existing master
			instance, err := inst.Repoint(instanceKey, siblingKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(fmt.Sprintf("%s<%s", instanceKey.DisplayString(), instance.MasterKey.DisplayString()))
		}
	case cliCommand("repoint-slaves"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			repointedSlaves, err, errs := inst.RepointSlaves(instanceKey, pattern)
			if err != nil {
				log.Fatale(err)
			} else {
				for _, e := range errs {
					log.Errore(e)
				}
				for _, slave := range repointedSlaves {
					fmt.Println(fmt.Sprintf("%s<%s", slave.Key.DisplayString(), instanceKey.DisplayString()))
				}
			}
		}
	case cliCommand("enslave-siblings"):
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, _, err := inst.EnslaveSiblings(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("enslave-master"):
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.EnslaveMaster(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("make-co-master"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.MakeCoMaster(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("match-below"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			if siblingKey == nil {
				log.Fatal("Cannot deduce sibling:", sibling)
			}
			_, _, err := inst.MatchBelow(instanceKey, siblingKey, true, true)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(fmt.Sprintf("%s<%s", instanceKey.DisplayString(), siblingKey.DisplayString()))
		}
	case cliCommand("match-up"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			instance, _, err := inst.MatchUp(instanceKey, true, true)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(fmt.Sprintf("%s<%s", instanceKey.DisplayString(), instance.MasterKey.DisplayString()))
		}
	case cliCommand("rematch"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			instance, _, err := inst.RematchSlave(instanceKey, true, true)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(fmt.Sprintf("%s<%s", instanceKey.DisplayString(), instance.MasterKey.DisplayString()))
		}
	case cliCommand("get-candidate-slave"):
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}

			instance, _, _, _, err := inst.GetCandidateSlave(instanceKey, false)
			if err != nil {
				log.Fatale(err)
			} else {
				fmt.Println(instance.Key.DisplayString())
			}
		}
	case cliCommand("multi-match-slaves"):
		{
			// Move all slaves of "instance" beneath "sibling"
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			if siblingKey == nil {
				log.Fatal("Cannot deduce sibling:", sibling)
			}

			matchedSlaves, _, err, errs := inst.MultiMatchSlaves(instanceKey, siblingKey, pattern)
			if err != nil {
				log.Fatale(err)
			} else {
				for _, e := range errs {
					log.Errore(e)
				}
				for _, slave := range matchedSlaves {
					fmt.Println(slave.Key.DisplayString())
				}
			}
		}
	case cliCommand("match-up-slaves"):
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}

			matchedSlaves, _, err, errs := inst.MatchUpSlaves(instanceKey, pattern)
			if err != nil {
				log.Fatale(err)
			} else {
				for _, e := range errs {
					log.Errore(e)
				}
				for _, slave := range matchedSlaves {
					fmt.Println(slave.Key.DisplayString())
				}
			}
		}
	case cliCommand("regroup-slaves"):
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}

			lostSlaves, equalSlaves, aheadSlaves, promotedSlave, err := inst.RegroupSlaves(instanceKey, func(candidateSlave *inst.Instance) { fmt.Println(candidateSlave.Key.DisplayString()) })
			fmt.Println(fmt.Sprintf("%s lost: %d, trivial: %d, pseudo-gtid: %d",
				promotedSlave.Key.DisplayString(), len(lostSlaves), len(equalSlaves), len(aheadSlaves)))
			if err != nil {
				log.Fatale(err)
			}
		}
	case cliCommand("recover"):
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}

			actionTaken, promotedInstance, err := logic.CheckAndRecover(instanceKey, siblingKey, true)
			if err != nil {
				log.Fatale(err)
			}
			if actionTaken {
				fmt.Println(promotedInstance.Key.DisplayString())
			}
		}
	case cliCommand("last-pseudo-gtid"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatalf("Unresolved instance")
			}
			instance, err := inst.ReadTopologyInstance(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			if instance == nil {
				log.Fatalf("Instance not found: %+v", *instanceKey)
			}
			coordinates, text, err := inst.FindLastPseudoGTIDEntry(instance, instance.RelaylogCoordinates, strict, nil)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(fmt.Sprintf("%+v:%s", *coordinates, text))
		}

	case cliCommand("stop-slave"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.StopSlave(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("start-slave"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.StartSlave(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("reset-slave"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.ResetSlaveOperation(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("detach-slave"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.DetachSlaveOperation(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("reattach-slave"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.ReattachSlaveOperation(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("set-read-only"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.SetReadOnly(instanceKey, true)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("set-writeable"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.SetReadOnly(instanceKey, false)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("discover"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			instance, err := inst.ReadTopologyInstance(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instance.Key.DisplayString())
		}
	case cliCommand("forget"):
		{
			if rawInstanceKey == nil {
				rawInstanceKey = thisInstanceKey
			}
			if rawInstanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			err := inst.ForgetInstance(rawInstanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(rawInstanceKey.DisplayString())
		}
	case cliCommand("begin-maintenance"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			if reason == "" {
				log.Fatal("--reason option required")
			}
			var durationSeconds int = 0
			if duration != "" {
				durationSeconds, err = util.SimpleTimeToSeconds(duration)
				if err != nil {
					log.Fatale(err)
				}
				if durationSeconds < 0 {
					log.Fatalf("Duration value must be non-negative. Given value: %d", durationSeconds)
				}
			}
			maintenanceKey, err := inst.BeginBoundedMaintenance(instanceKey, inst.GetMaintenanceOwner(), reason, uint(durationSeconds))
			if err == nil {
				log.Infof("Maintenance key: %+v", maintenanceKey)
				log.Infof("Maintenance duration: %d seconds", durationSeconds)
			}
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("end-maintenance"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			err := inst.EndMaintenanceByInstanceKey(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("begin-downtime"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			if reason == "" {
				log.Fatal("--reason option required")
			}
			var durationSeconds int = 0
			if duration != "" {
				durationSeconds, err = util.SimpleTimeToSeconds(duration)
				if err != nil {
					log.Fatale(err)
				}
				if durationSeconds < 0 {
					log.Fatalf("Duration value must be non-negative. Given value: %d", durationSeconds)
				}
			}
			err := inst.BeginDowntime(instanceKey, inst.GetMaintenanceOwner(), reason, uint(durationSeconds))
			if err == nil {
				log.Infof("Downtime duration: %d seconds", durationSeconds)
			} else {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("end-downtime"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			err := inst.EndDowntime(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("register-candidate"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			err := inst.RegisterCandidateInstance(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("register-hostname-unresolve"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			err := inst.RegisterHostnameUnresolve(instanceKey, hostnameFlag)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("deregister-hostname-unresolve"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			err := inst.DeregisterHostnameUnresolve(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case cliCommand("submit-pool-instances"):
		{
			if pool == "" {
				log.Fatal("Please submit --pool")
			}
			err := inst.ApplyPoolInstances(pool, instance)
			if err != nil {
				log.Fatale(err)
			}
		}
	case cliCommand("cluster-pool-instances"):
		{
			clusterPoolInstances, err := inst.ReadAllClusterPoolInstances()
			if err != nil {
				log.Fatale(err)
			}
			for _, clusterPoolInstance := range clusterPoolInstances {
				fmt.Println(fmt.Sprintf("%s\t%s\t%s\t%s:%d", clusterPoolInstance.ClusterName, clusterPoolInstance.ClusterAlias, clusterPoolInstance.Pool, clusterPoolInstance.Hostname, clusterPoolInstance.Port))
			}
		}
	case cliCommand("clusters"):
		{
			clusters, err := inst.ReadClusters()
			if err != nil {
				log.Fatale(err)
			} else {
				fmt.Println(strings.Join(clusters, "\n"))
			}
		}
	case cliCommand("find"):
		{
			if pattern == "" {
				log.Fatal("No pattern given")
			}
			instances, err := inst.FindInstances(pattern)
			if err != nil {
				log.Fatale(err)
			} else {
				for _, instance := range instances {
					fmt.Println(instance.Key.DisplayString())
				}
			}
		}
	case cliCommand("topology"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			output, err := inst.ASCIITopology(instanceKey, pattern)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(output)
		}
	case cliCommand("which-instance"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatalf("Unable to get master: unresolved instance")
			}
			instance, _, err := inst.ReadInstance(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			if instance == nil {
				log.Fatalf("Instance not found: %+v", *instanceKey)
			}
			fmt.Println(instance.Key.DisplayString())
		}
	case cliCommand("which-master"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatalf("Unable to get master: unresolved instance")
			}
			instance, _, err := inst.ReadInstance(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			if instance == nil {
				log.Fatalf("Instance not found: %+v", *instanceKey)
			}
			fmt.Println(instance.MasterKey.DisplayString())
		}
	case cliCommand("which-cluster"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatalf("Unable to get cluster: unresolved instance")
			}
			instance, _, err := inst.ReadInstance(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			if instance == nil {
				log.Fatalf("Instance not found: %+v", *instanceKey)
			}
			fmt.Println(instance.ClusterName)
		}
	case cliCommand("which-cluster-instances"):
		{
			clusterName := getClusterName(clusterAlias, instanceKey)
			instances, err := inst.ReadClusterInstances(clusterName)
			if err != nil {
				log.Fatale(err)
			}
			for _, clusterInstance := range instances {
				fmt.Println(clusterInstance.Key.DisplayString())
			}
		}
	case cliCommand("which-cluster-osc-slaves"):
		{
			clusterName := getClusterName(clusterAlias, instanceKey)
			instances, err := inst.GetClusterOSCSlaves(clusterName)
			if err != nil {
				log.Fatale(err)
			}
			for _, clusterInstance := range instances {
				fmt.Println(clusterInstance.Key.DisplayString())
			}
		}
	case cliCommand("which-slaves"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatalf("Unable to get slaves: unresolved instance")
			}
			slaves, err := inst.ReadSlaveInstances(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			for _, slave := range slaves {
				fmt.Println(slave.Key.DisplayString())
			}
		}
	case cliCommand("snapshot-topologies"):
		{
			err := inst.SnapshotTopologies()
			if err != nil {
				log.Fatale(err)
			}
		}
	case cliCommand("instance-status"):
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatalf("Unable to get status: unresolved instance")
			}
			instance, _, err := inst.ReadInstance(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			if instance == nil {
				log.Fatalf("Instance not found: %+v", *instanceKey)
			}
			fmt.Println(instance.HumanReadableDescription())
		}
	case cliCommand("get-cluster-heuristic-lag"):
		{
			clusterName := getClusterName(clusterAlias, instanceKey)
			lag, err := inst.GetClusterHeuristicLag(clusterName)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(lag)
		}
	case cliCommand("replication-analysis"):
		{
			analysis, err := inst.GetReplicationAnalysis(false)
			if err != nil {
				log.Fatale(err)
			}
			for _, entry := range analysis {
				fmt.Println(fmt.Sprintf("%s (cluster %s): %s", entry.AnalyzedInstanceKey.DisplayString(), entry.ClusterDetails.ClusterName, entry.Analysis))
			}
		}
	case cliCommand("continuous"):
		{
			logic.ContinuousDiscovery()
		}
	case cliCommand("reset-hostname-resolve-cache"):
		{
			err := inst.ResetHostnameResolveCache()
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println("hostname resolve cache cleared")
		}
	case cliCommand("resolve"):
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			if conn, err := net.Dial("tcp", instanceKey.DisplayString()); err == nil {
				conn.Close()
			} else {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	default:
		log.Fatalf("Unknown command: \"%s\". Available commands (-c):\n\t%v", command, strings.Join(knownCommands, "\n\t"))
	}
}
