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
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/inst"
	"github.com/outbrain/orchestrator/logic"
	"net"
	"os"
	"os/user"
	"strings"
)

// Cli initiates a command line interface, executing requested command.
func Cli(command string, strict bool, instance string, sibling string, owner string, reason string, pattern string) {

	if instance != "" && !strings.Contains(instance, ":") {
		instance = fmt.Sprintf("%s:%d", instance, config.Config.DefaultInstancePort)
	}
	instanceKey, err := inst.ParseInstanceKey(instance)
	if err != nil {
		instanceKey = nil
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
		log.Fatal("expected command (-c) (discover|forget|continuous|move-up|move-below|make-co-master|match-below|reset-slave|set-read-only|set-writeable|begin-maintenance|end-maintenance|clusters|topology|resolve)")
	}
	switch command {
	case "move-up":
		{
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.MoveUp(instanceKey)
			if err != nil {
				log.Fatale(err)
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
	case "enslave-sublings-simple":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, _, err := inst.EnslaveSiblingsSimple(instanceKey)
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
			fmt.Println(instance.Key.DisplayString())
		}
	case "get-candidate-slave":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}

			instance, _, _, _, err := inst.GetCandidateSlave(instanceKey, strict, true)
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

			matchedSlaves, _, err := inst.MultiMatchSlaves(instanceKey, siblingKey)
			if err != nil {
				log.Fatale(err)
			} else {
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

			matchedSlaves, _, err := inst.MatchUpSlaves(instanceKey)
			if err != nil {
				log.Fatale(err)
			} else {
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

			lostSlaves, equalSlaves, aheadSlaves, promotedSlave, err := inst.RegroupSlaves(instanceKey)
			if err != nil {
				log.Fatale(err)
			} else {
				fmt.Println(fmt.Sprintf("promoted slave: %s, lost: %d, trivial: %d, pseudo-gtid: %d",
					promotedSlave.Key.DisplayString(), len(lostSlaves), len(equalSlaves), len(aheadSlaves)))
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
			if instanceKey == nil {
				instanceKey = thisInstanceKey
			}
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			err := inst.ForgetInstance(instanceKey)
			if err != nil {
				log.Fatale(err)
			}
			fmt.Println(instanceKey.DisplayString())
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
			maintenanceKey, err := inst.BeginMaintenance(instanceKey, inst.GetMaintenanceOwner(), reason)
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
