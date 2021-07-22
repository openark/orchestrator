package inst

import (
	"github.com/openark/golib/log"
	test "github.com/openark/golib/tests"
	"github.com/openark/orchestrator/go/config"
	"testing"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.MarkConfigurationLoaded()
	log.SetLevel(log.ERROR)
}

func newTestReplica(key string, masterKey string, lastCheckValid bool, semiSyncPriority uint, promotionRule CandidatePromotionRule, replicationState ReplicationThreadState) *Instance {
	return &Instance{
		Key:                       InstanceKey{Hostname: key, Port: 3306},
		MasterKey:                 InstanceKey{Hostname: masterKey, Port: 3306},
		ReadBinlogCoordinates:     BinlogCoordinates{LogFile: "mysql.000001", LogPos: 10},
		ReplicationSQLThreadState: replicationState,
		ReplicationIOThreadState:  replicationState,
		IsLastCheckValid:          lastCheckValid,
		SemiSyncPriority:          semiSyncPriority,
		PromotionRule:             promotionRule,
	}
}

func expectInstancesMatch(t *testing.T, actual []*Instance, expected []*Instance) {
	if len(expected) != len(actual) {
		t.Fatalf("Actual instance list %+v does not match expected list %+v", actual, expected)
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Fatalf("Actual instance %+v does not match expected %+v", actual[i], expected[i])
		}
	}
}

func TestClassifyAndPrioritizeReplicas_NoPrioritiesSamePromotionRule_NameTiebreaker(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica3, replica2, replica1} // inverse order!

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica1, replica2, replica3})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}

func TestClassifyAndPrioritizeReplicas_WithPriorities(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, 3, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, 2, NeutralPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica1, replica2, replica3}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica2, replica3, replica1})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}

func TestClassifyAndPrioritizeReplicas_WithPrioritiesAndPromotionRules_PriorityTakesPrecedence(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, 1, PreferPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, 3, MustNotPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, 2, NeutralPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica1, replica2, replica3}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica2, replica3, replica1})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}

func TestClassifyAndPrioritizeReplicas_LastCheckInvalidAndNotReplication(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, 1, MustNotPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, 1, NeutralPromoteRule, ReplicationThreadStateStopped)
	replica3 := newTestReplica("replica3", "master1", false, 1, PreferPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica1, replica2, replica3}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica1})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{replica2, replica3})
}

func TestClassifyAndPrioritizeReplicas_NonReplicatingReplica(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, 1, MustNotPromoteRule, ReplicationThreadStateStopped)
	replicas := []*Instance{replica1, replica2, replica3}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, &replica3.Key) // Non-replicating instance
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica1, replica2, replica3})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}

func TestDetermineSemiSyncReplicaActionsForExactTopology_EnableSomeDisableSome(t *testing.T) {
	master := &Instance{SemiSyncMasterWaitForReplicaCount: 2}
	replica1 := &Instance{SemiSyncReplicaEnabled: true}
	replica2 := &Instance{SemiSyncReplicaEnabled: false}
	replica3 := &Instance{SemiSyncReplicaEnabled: true}
	replica4 := &Instance{SemiSyncReplicaEnabled: false}
	replica5 := &Instance{SemiSyncReplicaEnabled: true}
	possibleSemiSyncReplicas := []*Instance{replica1, replica2, replica3, replica4}
	asyncReplicas := []*Instance{replica5}
	actions := determineSemiSyncReplicaActionsForExactTopology(master, possibleSemiSyncReplicas, asyncReplicas)
	test.S(t).ExpectTrue(len(actions) == 3)
	test.S(t).ExpectTrue(actions[replica2])
	test.S(t).ExpectFalse(actions[replica3])
	test.S(t).ExpectFalse(actions[replica5])
}

func TestDetermineSemiSyncReplicaActionsForExactTopology_NoActions(t *testing.T) {
	master := &Instance{SemiSyncMasterWaitForReplicaCount: 1, SemiSyncMasterClients: 1}
	replica1 := &Instance{SemiSyncReplicaEnabled: true}
	replica2 := &Instance{SemiSyncReplicaEnabled: false}
	replica3 := &Instance{SemiSyncReplicaEnabled: false}
	possibleSemiSyncReplicas := []*Instance{replica1, replica2, replica3}
	asyncReplicas := []*Instance{}
	actions := determineSemiSyncReplicaActionsForExactTopology(master, possibleSemiSyncReplicas, asyncReplicas)
	test.S(t).ExpectTrue(len(actions) == 0)
}

func TestDetermineSemiSyncReplicaActionsForEnoughTopology_MoreThanWaitCountNoActions(t *testing.T) {
	master := &Instance{SemiSyncMasterWaitForReplicaCount: 1, SemiSyncMasterClients: 3}
	replica1 := &Instance{SemiSyncReplicaEnabled: true}
	replica2 := &Instance{SemiSyncReplicaEnabled: true}
	replica3 := &Instance{SemiSyncReplicaEnabled: true}
	possibleSemiSyncReplicas := []*Instance{replica1, replica2, replica3}
	actions := determineSemiSyncReplicaActionsForEnoughTopology(master, possibleSemiSyncReplicas)
	test.S(t).ExpectTrue(len(actions) == 0)
}

func TestDetermineSemiSyncReplicaActionsForEnoughTopology_LessThanWaitCountEnableOne(t *testing.T) {
	master := &Instance{SemiSyncMasterWaitForReplicaCount: 2, SemiSyncMasterClients: 1}
	replica1 := &Instance{SemiSyncReplicaEnabled: false}
	replica2 := &Instance{SemiSyncReplicaEnabled: true}
	replica3 := &Instance{SemiSyncReplicaEnabled: false}
	possibleSemiSyncReplicas := []*Instance{replica1, replica2, replica3}
	actions := determineSemiSyncReplicaActionsForEnoughTopology(master, possibleSemiSyncReplicas)
	test.S(t).ExpectTrue(len(actions) == 1)
	test.S(t).ExpectTrue(actions[replica1])
}
