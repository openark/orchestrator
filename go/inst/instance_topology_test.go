package inst

import (
	"github.com/outbrain/golib/log"
	test "github.com/outbrain/golib/tests"
	"github.com/outbrain/orchestrator/go/config"
	"testing"
)

var (
	i710Key = InstanceKey{Hostname: "i710", Port: 3306}
	i720Key = InstanceKey{Hostname: "i720", Port: 3306}
	i730Key = InstanceKey{Hostname: "i730", Port: 3306}
	i810Key = InstanceKey{Hostname: "i810", Port: 3306}
	i820Key = InstanceKey{Hostname: "i820", Port: 3306}
	i830Key = InstanceKey{Hostname: "i830", Port: 3306}
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	log.SetLevel(log.ERROR)
}

func generateTestInstances() (instances [](*Instance), instancesMap map[string](*Instance)) {
	i710 := Instance{Key: i710Key, ServerID: 710, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000007", LogPos: 10}}
	i720 := Instance{Key: i720Key, ServerID: 720, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000007", LogPos: 20}}
	i730 := Instance{Key: i730Key, ServerID: 730, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000007", LogPos: 30}}
	i810 := Instance{Key: i810Key, ServerID: 810, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000008", LogPos: 10}}
	i820 := Instance{Key: i820Key, ServerID: 820, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000008", LogPos: 20}}
	i830 := Instance{Key: i830Key, ServerID: 830, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000008", LogPos: 30}}
	instances = [](*Instance){&i710, &i720, &i730, &i810, &i820, &i830}
	for _, instance := range instances {
		instance.Version = "5.6.7"
		instance.Binlog_format = "STATEMENT"
	}
	instancesMap = make(map[string](*Instance))
	for _, instance := range instances {
		instancesMap[instance.Key.StringCode()] = instance
	}
	return instances, instancesMap
}

func applyGeneralGoodToGoReplicationParams(instances [](*Instance)) {
	for _, instance := range instances {
		instance.IsLastCheckValid = true
		instance.LogBinEnabled = true
		instance.LogSlaveUpdatesEnabled = true
	}
}

func TestInitial(t *testing.T) {
	test.S(t).ExpectTrue(true)
}

func TestSortInstances(t *testing.T) {
	instances, _ := generateTestInstances()
	sortInstances(instances)
	test.S(t).ExpectEquals(instances[0].Key, i830Key)
	test.S(t).ExpectEquals(instances[1].Key, i820Key)
	test.S(t).ExpectEquals(instances[2].Key, i810Key)
	test.S(t).ExpectEquals(instances[3].Key, i730Key)
	test.S(t).ExpectEquals(instances[4].Key, i720Key)
	test.S(t).ExpectEquals(instances[5].Key, i710Key)
}

func TestSortInstancesSameCoordinatesDifferingBinlogFormats(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	for _, instance := range instances {
		instance.ExecBinlogCoordinates = instances[0].ExecBinlogCoordinates
		instance.Binlog_format = "MIXED"
	}
	instancesMap[i810Key.StringCode()].Binlog_format = "STATEMENT"
	instancesMap[i720Key.StringCode()].Binlog_format = "ROW"
	sortInstances(instances)
	test.S(t).ExpectEquals(instances[0].Key, i810Key)
	test.S(t).ExpectEquals(instances[5].Key, i720Key)
}

func TestSortInstancesSameCoordinatesDifferingVersions(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	for _, instance := range instances {
		instance.ExecBinlogCoordinates = instances[0].ExecBinlogCoordinates
	}
	instancesMap[i810Key.StringCode()].Version = "5.5.1"
	instancesMap[i720Key.StringCode()].Version = "5.7.8"
	sortInstances(instances)
	test.S(t).ExpectEquals(instances[0].Key, i810Key)
	test.S(t).ExpectEquals(instances[5].Key, i720Key)
}

func TestGetPriorityMajorVersionForCandidate(t *testing.T) {
	instances, instancesMap := generateTestInstances()

	priorityMajorVersion, err := getPriorityMajorVersionForCandidate(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(priorityMajorVersion, "5.6")

	instancesMap[i810Key.StringCode()].Version = "5.5.1"
	instancesMap[i720Key.StringCode()].Version = "5.7.8"
	priorityMajorVersion, err = getPriorityMajorVersionForCandidate(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(priorityMajorVersion, "5.6")

	instancesMap[i710Key.StringCode()].Version = "5.7.8"
	instancesMap[i720Key.StringCode()].Version = "5.7.8"
	instancesMap[i730Key.StringCode()].Version = "5.7.8"
	instancesMap[i830Key.StringCode()].Version = "5.7.8"
	priorityMajorVersion, err = getPriorityMajorVersionForCandidate(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(priorityMajorVersion, "5.7")
}

func TestGetPriorityBinlogFormatForCandidate(t *testing.T) {
	instances, instancesMap := generateTestInstances()

	priorityBinlogFormat, err := getPriorityBinlogFormatForCandidate(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(priorityBinlogFormat, "STATEMENT")

	instancesMap[i810Key.StringCode()].Binlog_format = "MIXED"
	instancesMap[i720Key.StringCode()].Binlog_format = "ROW"
	priorityBinlogFormat, err = getPriorityBinlogFormatForCandidate(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(priorityBinlogFormat, "STATEMENT")

	instancesMap[i710Key.StringCode()].Binlog_format = "ROW"
	instancesMap[i720Key.StringCode()].Binlog_format = "ROW"
	instancesMap[i730Key.StringCode()].Binlog_format = "ROW"
	instancesMap[i830Key.StringCode()].Binlog_format = "ROW"
	priorityBinlogFormat, err = getPriorityBinlogFormatForCandidate(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(priorityBinlogFormat, "ROW")
}

func TestIsGenerallyValidAsBinlogSource(t *testing.T) {
	instances, _ := generateTestInstances()
	for _, instance := range instances {
		test.S(t).ExpectFalse(isGenerallyValidAsBinlogSource(instance))
	}
	applyGeneralGoodToGoReplicationParams(instances)
	for _, instance := range instances {
		test.S(t).ExpectTrue(isGenerallyValidAsBinlogSource(instance))
	}
}

func TestIsGenerallyValidAsCandidateSlave(t *testing.T) {
	instances, _ := generateTestInstances()
	for _, instance := range instances {
		test.S(t).ExpectFalse(isGenerallyValidAsCandidateSlave(instance))
	}
	for _, instance := range instances {
		instance.IsLastCheckValid = true
		instance.LogBinEnabled = true
		instance.LogSlaveUpdatesEnabled = false
	}
	for _, instance := range instances {
		test.S(t).ExpectFalse(isGenerallyValidAsCandidateSlave(instance))
	}
	applyGeneralGoodToGoReplicationParams(instances)
	for _, instance := range instances {
		test.S(t).ExpectTrue(isGenerallyValidAsCandidateSlave(instance))
	}
}

func TestIsBannedFromBeingCandidateSlave(t *testing.T) {
	instances, _ := generateTestInstances()
	for _, instance := range instances {
		test.S(t).ExpectFalse(isBannedFromBeingCandidateSlave(instance))
	}
}

func TestChooseCandidateSlaveNoCandidateSlave(t *testing.T) {
	instances, _ := generateTestInstances()
	for _, instance := range instances {
		instance.IsLastCheckValid = true
		instance.LogBinEnabled = true
		instance.LogSlaveUpdatesEnabled = false
	}
	_, _, _, _, _, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNotNil(err)
}

func TestChooseCandidateSlave(t *testing.T) {
	instances, _ := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i830Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 0)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 5)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlave2(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instancesMap[i830Key.StringCode()].LogSlaveUpdatesEnabled = false
	instancesMap[i820Key.StringCode()].LogBinEnabled = false
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i810Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 2)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 3)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlaveSameCoordinatesDifferentVersions(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	for _, instance := range instances {
		instance.ExecBinlogCoordinates = instances[0].ExecBinlogCoordinates
	}
	instancesMap[i810Key.StringCode()].Version = "5.5.1"
	instancesMap[i720Key.StringCode()].Version = "5.7.8"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i810Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 0)
	test.S(t).ExpectEquals(len(equalSlaves), 5)
	test.S(t).ExpectEquals(len(laterSlaves), 0)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlavePriorityVersionNoLoss(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instancesMap[i830Key.StringCode()].Version = "5.5.1"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i830Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 0)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 5)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlavePriorityVersionLosesOne(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instancesMap[i830Key.StringCode()].Version = "5.7.8"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i820Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 1)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 4)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlavePriorityVersionLosesTwo(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instancesMap[i830Key.StringCode()].Version = "5.7.8"
	instancesMap[i820Key.StringCode()].Version = "5.7.18"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i810Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 2)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 3)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlavePriorityVersionHigherVersionOverrides(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instancesMap[i830Key.StringCode()].Version = "5.7.8"
	instancesMap[i820Key.StringCode()].Version = "5.7.18"
	instancesMap[i810Key.StringCode()].Version = "5.7.5"
	instancesMap[i730Key.StringCode()].Version = "5.7.30"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i830Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 0)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 3)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 2)
}

func TestChooseCandidateSlaveLosesOneDueToBinlogFormat(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	for _, instance := range instances {
		instance.Binlog_format = "ROW"
	}
	instancesMap[i730Key.StringCode()].Binlog_format = "STATEMENT"

	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i830Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 0)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 4)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 1)
}

func TestChooseCandidateSlavePriorityBinlogFormatNoLoss(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	for _, instance := range instances {
		instance.Binlog_format = "MIXED"
	}
	instancesMap[i830Key.StringCode()].Binlog_format = "STATEMENT"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i830Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 0)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 5)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlavePriorityBinlogFormatLosesOne(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instancesMap[i830Key.StringCode()].Binlog_format = "ROW"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i820Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 1)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 4)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlavePriorityBinlogFormatLosesTwo(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instancesMap[i830Key.StringCode()].Binlog_format = "ROW"
	instancesMap[i820Key.StringCode()].Binlog_format = "ROW"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i810Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 2)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 3)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 0)
}

func TestChooseCandidateSlavePriorityBinlogFormatRowOverrides(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	applyGeneralGoodToGoReplicationParams(instances)
	instancesMap[i830Key.StringCode()].Binlog_format = "ROW"
	instancesMap[i820Key.StringCode()].Binlog_format = "ROW"
	instancesMap[i810Key.StringCode()].Binlog_format = "ROW"
	instancesMap[i730Key.StringCode()].Binlog_format = "ROW"
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, cannotReplicateSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i830Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 0)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 3)
	test.S(t).ExpectEquals(len(cannotReplicateSlaves), 2)
}
