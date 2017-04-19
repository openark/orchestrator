package config

import (
	"testing"

	"github.com/openark/golib/log"
	test "github.com/openark/golib/tests"
)

func init() {
	Config().HostnameResolveMethod = "none"
	log.SetLevel(log.ERROR)
}

func TestReplicationLagQuery(t *testing.T) {
	{
		c := defaultConfiguration()
		c.SlaveLagQuery = "select 3"
		c.ReplicationLagQuery = "select 4"
		err := c.postReadAdjustments()
		test.S(t).ExpectNotNil(err)
	}
	{
		c := defaultConfiguration()
		c.SlaveLagQuery = "select 3"
		c.ReplicationLagQuery = "select 3"
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
	}
	{
		c := defaultConfiguration()
		c.SlaveLagQuery = "select 3"
		c.ReplicationLagQuery = ""
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(c.ReplicationLagQuery, "select 3")
	}
}

func TestPostponeReplicaRecoveryOnLagMinutes(t *testing.T) {
	{
		c := defaultConfiguration()
		c.PostponeSlaveRecoveryOnLagMinutes = 3
		c.PostponeReplicaRecoveryOnLagMinutes = 5
		err := c.postReadAdjustments()
		test.S(t).ExpectNotNil(err)
	}
	{
		c := defaultConfiguration()
		c.PostponeSlaveRecoveryOnLagMinutes = 3
		c.PostponeReplicaRecoveryOnLagMinutes = 3
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
	}
	{
		c := defaultConfiguration()
		c.PostponeSlaveRecoveryOnLagMinutes = 3
		c.PostponeReplicaRecoveryOnLagMinutes = 0
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(c.PostponeReplicaRecoveryOnLagMinutes, uint(3))
	}
}

func TestMasterFailoverDetachReplicaMasterHost(t *testing.T) {
	{
		c := defaultConfiguration()
		c.MasterFailoverDetachSlaveMasterHost = false
		c.MasterFailoverDetachReplicaMasterHost = false
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectFalse(c.MasterFailoverDetachReplicaMasterHost)
	}
	{
		c := defaultConfiguration()
		c.MasterFailoverDetachSlaveMasterHost = false
		c.MasterFailoverDetachReplicaMasterHost = true
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectTrue(c.MasterFailoverDetachReplicaMasterHost)
	}
	{
		c := defaultConfiguration()
		c.MasterFailoverDetachSlaveMasterHost = true
		c.MasterFailoverDetachReplicaMasterHost = false
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectTrue(c.MasterFailoverDetachReplicaMasterHost)
	}
}

func TestMasterFailoverDetachDetachLostReplicasAfterMasterFailover(t *testing.T) {
	{
		c := defaultConfiguration()
		c.DetachLostSlavesAfterMasterFailover = false
		c.DetachLostReplicasAfterMasterFailover = false
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectFalse(c.DetachLostReplicasAfterMasterFailover)
	}
	{
		c := defaultConfiguration()
		c.DetachLostSlavesAfterMasterFailover = false
		c.DetachLostReplicasAfterMasterFailover = true
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectTrue(c.DetachLostReplicasAfterMasterFailover)
	}
	{
		c := defaultConfiguration()
		c.DetachLostSlavesAfterMasterFailover = true
		c.DetachLostReplicasAfterMasterFailover = false
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectTrue(c.DetachLostReplicasAfterMasterFailover)
	}
}

func TestRecoveryPeriodBlock(t *testing.T) {
	{
		c := defaultConfiguration()
		c.RecoveryPeriodBlockSeconds = 0
		c.RecoveryPeriodBlockMinutes = 0
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(c.RecoveryPeriodBlockSeconds, 0)
	}
	{
		c := defaultConfiguration()
		c.RecoveryPeriodBlockSeconds = 30
		c.RecoveryPeriodBlockMinutes = 1
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(c.RecoveryPeriodBlockSeconds, 30)
	}
	{
		c := defaultConfiguration()
		c.RecoveryPeriodBlockSeconds = 0
		c.RecoveryPeriodBlockMinutes = 2
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(c.RecoveryPeriodBlockSeconds, 120)
	}
	{
		c := defaultConfiguration()
		c.RecoveryPeriodBlockSeconds = 15
		c.RecoveryPeriodBlockMinutes = 0
		err := c.postReadAdjustments()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(c.RecoveryPeriodBlockSeconds, 15)
	}
}
