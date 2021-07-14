package inst

import (
	"github.com/openark/golib/log"
	"github.com/openark/orchestrator/go/config"
	"testing"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.MarkConfigurationLoaded()
	log.SetLevel(log.ERROR)
}

func newTestReplica(key string, masterKey string, lastCheckValid bool, downtimed bool, semiSyncPriority uint, promotionRule CandidatePromotionRule, replicationState ReplicationThreadState) *Instance {
	return &Instance{
		Key:                       InstanceKey{Hostname: key, Port: 3306},
		MasterKey:                 InstanceKey{Hostname: masterKey, Port: 3306},
		ReadBinlogCoordinates:     BinlogCoordinates{LogFile: "mysql.000001", LogPos: 10},
		ReplicationSQLThreadState: replicationState,
		ReplicationIOThreadState:  replicationState,
		IsLastCheckValid:          lastCheckValid,
		IsDowntimed:               downtimed,
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
	replica1 := newTestReplica("replica1", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica3, replica2, replica1} // inverse order!

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica1, replica2, replica3})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}

func TestClassifyAndPrioritizeReplicas_WithPriorities(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, false, 3, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, false, 2, NeutralPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica1, replica2, replica3}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica2, replica3, replica1})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}

func TestClassifyAndPrioritizeReplicas_WithPrioritiesAndPromotionRules_PriorityTakesPrecedence(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, false, 1, PreferPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, false, 3, MustNotPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, false, 2, NeutralPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica1, replica2, replica3}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica2, replica3, replica1})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}

func TestClassifyAndPrioritizeReplicas_LastCheckInvalidAndNotReplication(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, false, 1, MustNotPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateStopped)
	replica3 := newTestReplica("replica3", "master1", false, false, 1, PreferPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica1, replica2, replica3}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica1})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{replica2, replica3})
}

func TestClassifyAndPrioritizeReplicas_Downtimed(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, true, 1, MustNotPromoteRule, ReplicationThreadStateRunning)
	replica4 := newTestReplica("replica4", "master1", true, false, 0, MustNotPromoteRule, ReplicationThreadStateRunning)
	replicas := []*Instance{replica1, replica2, replica3, replica4}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, nil)
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica1, replica2})
	expectInstancesMatch(t, asyncReplicas, []*Instance{replica3, replica4})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}

func TestClassifyAndPrioritizeReplicas_NonReplicatingReplica(t *testing.T) {
	replica1 := newTestReplica("replica1", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica2 := newTestReplica("replica2", "master1", true, false, 1, NeutralPromoteRule, ReplicationThreadStateRunning)
	replica3 := newTestReplica("replica3", "master1", true, false, 1, MustNotPromoteRule, ReplicationThreadStateStopped)
	replicas := []*Instance{replica1, replica2, replica3}

	possibleSemiSyncReplicas, asyncReplicas, excludedReplicas := classifyAndPrioritizeReplicas(replicas, &replica3.Key) // Non-replicating instance
	expectInstancesMatch(t, possibleSemiSyncReplicas, []*Instance{replica1, replica2, replica3})
	expectInstancesMatch(t, asyncReplicas, []*Instance{})
	expectInstancesMatch(t, excludedReplicas, []*Instance{})
}
