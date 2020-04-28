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

package inst

import (
	"github.com/openark/orchestrator/go/config"
	"github.com/openark/orchestrator/go/db"
	. "gopkg.in/check.v1"
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

// This test suite assumes one master and three direct replicas, as follows;
// This was setup with mysqlsandbox (using MySQL 5.5.32, not that it matters) via:
// $ make_replication_sandbox --how_many_nodes=3 --replication_directory=55orchestrator /path/to/sandboxes/5.5.32
// modify below to fit your own environment
var masterKey = InstanceKey{
	Hostname: "127.0.0.1",
	Port:     22987,
}
var slave1Key = InstanceKey{
	Hostname: "127.0.0.1",
	Port:     22988,
}
var slave2Key = InstanceKey{
	Hostname: "127.0.0.1",
	Port:     22989,
}
var slave3Key = InstanceKey{
	Hostname: "127.0.0.1",
	Port:     22990,
}

func clearTestMaintenance() {
	_, _ = db.ExecOrchestrator("update database_instance_maintenance set maintenance_active=null, end_timestamp=NOW() where owner = ?", "unittest")
}

// The test also assumes one backend MySQL server.
func (s *TestSuite) SetUpSuite(c *C) {
	config.Config.MySQLTopologyUser = "msandbox"
	config.Config.MySQLTopologyPassword = "msandbox"
	config.Config.MySQLOrchestratorHost = "127.0.0.1"
	config.Config.MySQLOrchestratorPort = 5622
	config.Config.MySQLOrchestratorDatabase = "orchestrator"
	config.Config.MySQLOrchestratorUser = "msandbox"
	config.Config.MySQLOrchestratorPassword = "msandbox"
	config.Config.DiscoverByShowSlaveHosts = true

	_, _ = db.ExecOrchestrator("delete from database_instance where hostname = ? and port = ?", masterKey.Hostname, masterKey.Port)
	_, _ = db.ExecOrchestrator("delete from database_instance where hostname = ? and port = ?", slave1Key.Hostname, slave1Key.Port)
	_, _ = db.ExecOrchestrator("delete from database_instance where hostname = ? and port = ?", slave2Key.Hostname, slave2Key.Port)
	_, _ = db.ExecOrchestrator("delete from database_instance where hostname = ? and port = ?", slave3Key.Hostname, slave3Key.Port)

	ExecInstance(&masterKey, "drop database if exists orchestrator_test")
	ExecInstance(&masterKey, "create database orchestrator_test")
	ExecInstance(&masterKey, `create table orchestrator_test.test_table(
			name    varchar(128) charset ascii not null primary key,
			value   varchar(128) charset ascii not null
		)`)
	rand.Seed(time.Now().UTC().UnixNano())
}

func (s *TestSuite) TestReadTopologyMaster(c *C) {
	key := masterKey
	i, _ := ReadTopologyInstance(&key)

	c.Assert(i.Key.Hostname, Equals, key.Hostname)
	c.Assert(i.IsReplica(), Equals, false)
	c.Assert(len(i.SlaveHosts), Equals, 3)
	c.Assert(len(i.SlaveHosts.GetInstanceKeys()), Equals, len(i.SlaveHosts))
}

func (s *TestSuite) TestReadTopologySlave(c *C) {
	key := slave3Key
	i, _ := ReadTopologyInstance(&key)
	c.Assert(i.Key.Hostname, Equals, key.Hostname)
	c.Assert(i.IsReplica(), Equals, true)
	c.Assert(len(i.SlaveHosts), Equals, 0)
}

func (s *TestSuite) TestReadTopologyAndInstanceMaster(c *C) {
	i, _ := ReadTopologyInstance(&masterKey)
	iRead, found, _ := ReadInstance(&masterKey)
	c.Assert(found, Equals, true)
	c.Assert(iRead.Key.Hostname, Equals, i.Key.Hostname)
	c.Assert(iRead.Version, Equals, i.Version)
	c.Assert(len(iRead.SlaveHosts), Equals, len(i.SlaveHosts))
}

func (s *TestSuite) TestReadTopologyAndInstanceSlave(c *C) {
	i, _ := ReadTopologyInstance(&slave1Key)
	iRead, found, _ := ReadInstance(&slave1Key)
	c.Assert(found, Equals, true)
	c.Assert(iRead.Key.Hostname, Equals, i.Key.Hostname)
	c.Assert(iRead.Version, Equals, i.Version)
}

func (s *TestSuite) TestGetMasterOfASlave(c *C) {
	i, err := ReadTopologyInstance(&slave1Key)
	c.Assert(err, IsNil)
	master, err := GetInstanceMaster(i)
	c.Assert(err, IsNil)
	c.Assert(master.IsReplica(), Equals, false)
	c.Assert(master.Key.Port, Equals, 22987)
}

func (s *TestSuite) TestSlavesAreSiblings(c *C) {
	i0, _ := ReadTopologyInstance(&slave1Key)
	i1, _ := ReadTopologyInstance(&slave2Key)
	c.Assert(InstancesAreSiblings(i0, i1), Equals, true)
}

func (s *TestSuite) TestNonSiblings(c *C) {
	i0, _ := ReadTopologyInstance(&masterKey)
	i1, _ := ReadTopologyInstance(&slave1Key)
	c.Assert(InstancesAreSiblings(i0, i1), Not(Equals), true)
}

func (s *TestSuite) TestInstanceIsMasterOf(c *C) {
	i0, _ := ReadTopologyInstance(&masterKey)
	i1, _ := ReadTopologyInstance(&slave1Key)
	c.Assert(InstanceIsMasterOf(i0, i1), Equals, true)
}

func (s *TestSuite) TestStopStartSlave(c *C) {

	i, _ := ReadTopologyInstance(&slave1Key)
	c.Assert(i.ReplicaRunning(), Equals, true)
	i, _ = StopSlaveNicely(&i.Key, 0)

	c.Assert(i.ReplicaRunning(), Equals, false)
	c.Assert(i.SQLThreadUpToDate(), Equals, true)

	i, _ = StartSlave(&i.Key)
	c.Assert(i.ReplicaRunning(), Equals, true)
}

func (s *TestSuite) TestReadTopologyUnexisting(c *C) {
	key := InstanceKey{
		Hostname: "127.0.0.1",
		Port:     22999,
	}
	_, err := ReadTopologyInstance(&key)

	c.Assert(err, Not(IsNil))
}

func (s *TestSuite) TestMoveBelowAndBack(c *C) {
	clearTestMaintenance()
	// become child
	slave1, err := MoveBelow(&slave1Key, &slave2Key)
	c.Assert(err, IsNil)

	c.Assert(slave1.MasterKey.Equals(&slave2Key), Equals, true)
	c.Assert(slave1.ReplicaRunning(), Equals, true)

	// And back; keep topology intact
	slave1, _ = MoveUp(&slave1Key)
	slave2, _ := ReadTopologyInstance(&slave2Key)

	c.Assert(InstancesAreSiblings(slave1, slave2), Equals, true)
	c.Assert(slave1.ReplicaRunning(), Equals, true)

}

func (s *TestSuite) TestMoveBelowAndBackComplex(c *C) {
	clearTestMaintenance()

	// become child
	slave1, _ := MoveBelow(&slave1Key, &slave2Key)

	c.Assert(slave1.MasterKey.Equals(&slave2Key), Equals, true)
	c.Assert(slave1.ReplicaRunning(), Equals, true)

	// Now let's have fun. Stop replica2 (which is now parent of replica1), execute queries on master,
	// move s1 back under master, start all, verify queries.

	_, err := StopSlave(&slave2Key)
	c.Assert(err, IsNil)

	randValue := rand.Int()
	_, err = ExecInstance(&masterKey, `replace into orchestrator_test.test_table (name, value) values ('TestMoveBelowAndBackComplex', ?)`, randValue)
	c.Assert(err, IsNil)
	master, err := ReadTopologyInstance(&masterKey)
	c.Assert(err, IsNil)

	// And back; keep topology intact
	slave1, err = MoveUp(&slave1Key)
	c.Assert(err, IsNil)
	_, err = MasterPosWait(&slave1Key, &master.SelfBinlogCoordinates)
	c.Assert(err, IsNil)
	slave2, err := ReadTopologyInstance(&slave2Key)
	c.Assert(err, IsNil)
	_, err = MasterPosWait(&slave2Key, &master.SelfBinlogCoordinates)
	c.Assert(err, IsNil)
	// Now check for value!
	var value1, value2 int
	ScanInstanceRow(&slave1Key, `select value from orchestrator_test.test_table where name='TestMoveBelowAndBackComplex'`, &value1)
	ScanInstanceRow(&slave2Key, `select value from orchestrator_test.test_table where name='TestMoveBelowAndBackComplex'`, &value2)

	c.Assert(InstancesAreSiblings(slave1, slave2), Equals, true)
	c.Assert(value1, Equals, randValue)
	c.Assert(value2, Equals, randValue)
}

func (s *TestSuite) TestFailMoveBelow(c *C) {
	clearTestMaintenance()
	_, _ = ExecInstance(&slave2Key, `set global binlog_format:='ROW'`)
	_, err := MoveBelow(&slave1Key, &slave2Key)
	_, _ = ExecInstance(&slave2Key, `set global binlog_format:='STATEMENT'`)
	c.Assert(err, Not(IsNil))
}

func (s *TestSuite) TestMakeCoMasterAndBack(c *C) {
	clearTestMaintenance()

	slave1, err := MakeCoMaster(&slave1Key)
	c.Assert(err, IsNil)

	// Now master & replica1 expected to be co-masters. Check!
	master, _ := ReadTopologyInstance(&masterKey)
	c.Assert(master.IsReplicaOf(slave1), Equals, true)
	c.Assert(slave1.IsReplicaOf(master), Equals, true)

	// reset - restore to original state
	master, err = ResetSlaveOperation(&masterKey)
	slave1, _ = ReadTopologyInstance(&slave1Key)
	c.Assert(err, IsNil)
	c.Assert(master.MasterKey.Hostname, Equals, "_")
}

func (s *TestSuite) TestFailMakeCoMaster(c *C) {
	clearTestMaintenance()
	_, err := MakeCoMaster(&masterKey)
	c.Assert(err, Not(IsNil))
}

func (s *TestSuite) TestMakeCoMasterAndBackAndFailOthersToBecomeCoMasters(c *C) {
	clearTestMaintenance()

	slave1, err := MakeCoMaster(&slave1Key)
	c.Assert(err, IsNil)

	// Now master & replica1 expected to be co-masters. Check!
	master, _, _ := ReadInstance(&masterKey)
	c.Assert(master.IsReplicaOf(slave1), Equals, true)
	c.Assert(slave1.IsReplicaOf(master), Equals, true)

	// Verify can't have additional co-masters
	_, err = MakeCoMaster(&masterKey)
	c.Assert(err, Not(IsNil))
	_, err = MakeCoMaster(&slave1Key)
	c.Assert(err, Not(IsNil))
	_, err = MakeCoMaster(&slave2Key)
	c.Assert(err, Not(IsNil))

	// reset replica - restore to original state
	master, err = ResetSlaveOperation(&masterKey)
	c.Assert(err, IsNil)
	c.Assert(master.MasterKey.Hostname, Equals, "_")
}

func (s *TestSuite) TestDiscover(c *C) {
	var err error
	_, err = db.ExecOrchestrator("delete from database_instance where hostname = ? and port = ?", masterKey.Hostname, masterKey.Port)
	_, err = db.ExecOrchestrator("delete from database_instance where hostname = ? and port = ?", slave1Key.Hostname, slave1Key.Port)
	_, err = db.ExecOrchestrator("delete from database_instance where hostname = ? and port = ?", slave2Key.Hostname, slave2Key.Port)
	_, err = db.ExecOrchestrator("delete from database_instance where hostname = ? and port = ?", slave3Key.Hostname, slave3Key.Port)
	_, found, _ := ReadInstance(&masterKey)
	c.Assert(found, Equals, false)
	_, _ = ReadTopologyInstance(&slave1Key)
	_, found, err = ReadInstance(&slave1Key)
	c.Assert(found, Equals, true)
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestForgetMaster(c *C) {
	_, _ = ReadTopologyInstance(&masterKey)
	_, found, _ := ReadInstance(&masterKey)
	c.Assert(found, Equals, true)
	ForgetInstance(&masterKey)
	_, found, _ = ReadInstance(&masterKey)
	c.Assert(found, Equals, false)
}

func (s *TestSuite) TestBeginMaintenance(c *C) {
	clearTestMaintenance()
	_, _ = ReadTopologyInstance(&masterKey)
	_, err := BeginMaintenance(&masterKey, "unittest", "TestBeginMaintenance")

	c.Assert(err, IsNil)
}

func (s *TestSuite) TestBeginEndMaintenance(c *C) {
	clearTestMaintenance()
	_, _ = ReadTopologyInstance(&masterKey)
	k, err := BeginMaintenance(&masterKey, "unittest", "TestBeginEndMaintenance")
	c.Assert(err, IsNil)
	err = EndMaintenance(k)
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestFailBeginMaintenanceTwice(c *C) {
	clearTestMaintenance()
	_, _ = ReadTopologyInstance(&masterKey)
	_, err := BeginMaintenance(&masterKey, "unittest", "TestFailBeginMaintenanceTwice")
	c.Assert(err, IsNil)
	_, err = BeginMaintenance(&masterKey, "unittest", "TestFailBeginMaintenanceTwice")
	c.Assert(err, Not(IsNil))
}

func (s *TestSuite) TestFailEndMaintenanceTwice(c *C) {
	clearTestMaintenance()
	_, _ = ReadTopologyInstance(&masterKey)
	k, err := BeginMaintenance(&masterKey, "unittest", "TestFailEndMaintenanceTwice")
	c.Assert(err, IsNil)
	err = EndMaintenance(k)
	c.Assert(err, IsNil)
	err = EndMaintenance(k)
	c.Assert(err, Not(IsNil))
}

func (s *TestSuite) TestFailMoveBelowUponMaintenance(c *C) {
	clearTestMaintenance()
	_, _ = ReadTopologyInstance(&slave1Key)
	k, err := BeginMaintenance(&slave1Key, "unittest", "TestBeginEndMaintenance")
	c.Assert(err, IsNil)

	_, err = MoveBelow(&slave1Key, &slave2Key)
	c.Assert(err, Not(IsNil))

	err = EndMaintenance(k)
	c.Assert(err, IsNil)
}

func (s *TestSuite) TestFailMoveBelowUponSlaveStopped(c *C) {
	clearTestMaintenance()

	slave1, _ := ReadTopologyInstance(&slave1Key)
	c.Assert(slave1.ReplicaRunning(), Equals, true)
	slave1, _ = StopSlaveNicely(&slave1.Key, 0)
	c.Assert(slave1.ReplicaRunning(), Equals, false)

	_, err := MoveBelow(&slave1Key, &slave2Key)
	c.Assert(err, Not(IsNil))

	_, _ = StartSlave(&slave1.Key)
}

func (s *TestSuite) TestFailMoveBelowUponOtherSlaveStopped(c *C) {
	clearTestMaintenance()

	slave1, _ := ReadTopologyInstance(&slave1Key)
	c.Assert(slave1.ReplicaRunning(), Equals, true)
	slave1, _ = StopSlaveNicely(&slave1.Key, 0)
	c.Assert(slave1.ReplicaRunning(), Equals, false)

	_, err := MoveBelow(&slave2Key, &slave1Key)
	c.Assert(err, Not(IsNil))

	_, _ = StartSlave(&slave1.Key)
}
