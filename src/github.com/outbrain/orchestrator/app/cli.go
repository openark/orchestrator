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
	"github.com/outbrain/log"
	"github.com/outbrain/orchestrator/inst"
	"github.com/outbrain/orchestrator/logic"
	"net"
	"os/user"
	"strings"
)

// Cli initiates a command line interface, executing requested command.
func Cli(command string, instance string, sibling string, owner string, reason string) {

	instanceKey, err := inst.ParseInstanceKey(instance)
	if err != nil {
		instanceKey = nil
	}
	siblingKey, err := inst.ParseInstanceKey(sibling)
	if err != nil {
		siblingKey = nil
	}

	if len(owner) == 0 {
		// get os username as owner
		usr, err := user.Current()
		if err != nil {
			log.Fatale(err)
		}
		owner = usr.Username
	}

	if len(command) == 0 {
		log.Fatal("expected command (-c) (discover|forget|continuous|move-up|move-below|make-co-master|reset-slave|set-read-only|set-writeable|begin-maintenance|end-maintenance|clusters|topology|resolve)")
	}
	switch command {
	case "move-up":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.MoveUp(instanceKey)
			if err != nil {
				log.Errore(err)
			}
		}
	case "move-below":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			if siblingKey == nil {
				log.Fatal("Cannot deduce sibling:", sibling)
			}
			_, err := inst.MoveBelow(instanceKey, siblingKey)
			if err != nil {
				log.Errore(err)
			}
		}
	case "make-co-master":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.MakeCoMaster(instanceKey)
			if err != nil {
				log.Errore(err)
			}
		}
	case "reset-slave":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.ResetSlaveOperation(instanceKey)
			if err != nil {
				log.Errore(err)
			}
		}
	case "set-read-only":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.SetReadOnly(instanceKey, true)
			if err != nil {
				log.Errore(err)
			}
		}
	case "set-writeable":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			_, err := inst.SetReadOnly(instanceKey, false)
			if err != nil {
				log.Errore(err)
			}
		}
	case "discover":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			orchestrator.StartDiscovery(*instanceKey)
		}
	case "forget":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			inst.ForgetInstance(instanceKey)
		}
	case "begin-maintenance":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			if owner == "" {
				log.Fatal("--owner option required")
			}
			if reason == "" {
				log.Fatal("--reason option required")
			}
			maintenanceKey, err := inst.BeginMaintenance(instanceKey, owner, reason)
			if err == nil {
				log.Infof("Maintenance key: %+v", maintenanceKey)
			}
			if err != nil {
				log.Errore(err)
			}
		}
	case "end-maintenance":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			err := inst.EndMaintenanceByInstanceKey(instanceKey)
			if err != nil {
				log.Errore(err)
			}
		}
	case "clusters":
		{
			clusters, err := inst.ReadClusters()
			if err != nil {
				log.Errore(err)
			} else {
				fmt.Println(strings.Join(clusters, "\n"))
			}
		}
	case "topology":
		{
			if instanceKey == nil {
				log.Fatal("Cannot deduce instance:", instance)
			}
			output, err := inst.AsciiTopology(instance)
			if err != nil {
				log.Errore(err)
			} else {
				fmt.Println(output)
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
