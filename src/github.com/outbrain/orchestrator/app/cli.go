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
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/inst"
	"github.com/outbrain/orchestrator/logic"
	"net"
	"os"
	"os/user"
	"strings"
)

// Cli initiates a command line interface, executing requested command.
func Cli(command string, strict bool, instance string, sibling string, owner string, reason string, duration string, pattern string, clusterAlias string) {

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
	var thisInstanceKey *inst.InstanceKey = nil
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

	if len(command) == 0 {
		log.Fatal("expected command (-c) (skip-query|move-up|move-up-slaves|move-below|repoint|repoint-slaves|enslave-siblings|make-co-master|match-below|match-up|rematch|get-candidate-slave|multi-match-slaves|match-up-slaves|regroup-slaves|recover|last-pseudo-gtid|stop-slave|start-slave|reset-slave|detach-slave|reattach-slave|set-read-only|set-writeable|discover|forget|begin-maintenance|end-maintenance|clusters|find|topology|which-instance|which-master|which-cluster|which-cluster-instances|which-slaves|instance-status|replication-analysis|continuous|resolve)")
	}
	switch command {
	case "skip-query":
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
	case "move-up":
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
	case "move-up-slaves":
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
	case "move-below":
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
	case "repoint":
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
	case "repoint-slaves":
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
	case "enslave-siblings":
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
	case "make-co-master":
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
	case "match-below":
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
	case "match-up":
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
	case "rematch":
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
	case "get-candidate-slave":
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
	case "multi-match-slaves":
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
	case "match-up-slaves":
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
	case "regroup-slaves":
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
	case "recover":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}

			actionTaken, err := orchestrator.CheckAndRecover(instanceKey, true)
			if err != nil {
				log.Fatale(err)
			}
			if actionTaken {
				fmt.Println("true")
			}
		}
	case "last-pseudo-gtid":
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
			coordinates, text, err := inst.FindLastPseudoGTIDEntry(instance, instance.RelaylogCoordinates, strict)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(fmt.Sprintf("%+v:%s", *coordinates, text))
		}

	case "stop-slave":
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
	case "start-slave":
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
	case "reset-slave":
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
	case "detach-slave":
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
	case "reattach-slave":
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
	case "set-read-only":
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
	case "set-writeable":
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
	case "discover":
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			orchestrator.StartDiscovery(*instanceKey)
			fmt.Println(instanceKey.DisplayString())
		}
	case "forget":
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
	case "begin-maintenance":
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
			}
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
		}
	case "end-maintenance":
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
	case "clusters":
		{
			clusters, err := inst.ReadClusters()
			if err != nil {
				log.Fatale(err)
			} else {
				fmt.Println(strings.Join(clusters, "\n"))
			}
		}
	case "find":
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
	case "topology":
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			output, err := inst.AsciiTopology(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(output)
		}
	case "which-instance":
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
	case "which-master":
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
	case "which-cluster":
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
	case "which-cluster-instances":
		{
			clusterName := ""
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

			instances, err := inst.ReadClusterInstances(clusterName)
			if err != nil {
				log.Fatale(err)
			}
			for _, clusterInstance := range instances {
				fmt.Println(clusterInstance.Key.DisplayString())
			}
		}
	case "which-slaves":
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
	case "instance-status":
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
	case "replication-analysis":
		{
			analysis, err := inst.GetReplicationAnalysis()
			if err != nil {
				log.Fatale(err)
			}
			for _, entry := range analysis {
				fmt.Println(fmt.Sprintf("%s (cluster %s): %s", entry.AnalyzedInstanceKey.DisplayString(), entry.ClusterName, entry.Analysis))
			}
		}
	case "continuous":
		{
			orchestrator.ContinuousDiscovery()
		}
	case "resolve":
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
		log.Fatal("Unknown command:", command)
	}
}
