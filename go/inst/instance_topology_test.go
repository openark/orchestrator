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
	i710 := Instance{Key: i710Key, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000007", LogPos: 10}}
	i720 := Instance{Key: i720Key, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000007", LogPos: 20}}
	i730 := Instance{Key: i730Key, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000007", LogPos: 30}}
	i810 := Instance{Key: i810Key, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000008", LogPos: 10}}
	i820 := Instance{Key: i820Key, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000008", LogPos: 20}}
	i830 := Instance{Key: i830Key, ExecBinlogCoordinates: BinlogCoordinates{LogFile: "mysql.000008", LogPos: 30}}
	instances = [](*Instance){&i710, &i720, &i730, &i810, &i820, &i830}
	instancesMap = make(map[string](*Instance))
	for _, instance := range instances {
		instancesMap[instance.Key.StringCode()] = instance
	}
	return instances, instancesMap
}

func TestInitial(t *testing.T) {
	test.S(t).ExpectTrue(true)
}

func TestSortInstances(t *testing.T) {
	instances, _ := generateTestInstances()
	for _, instance := range instances {
		instance.Version = "5.6.7"
	}
	sortInstances(instances)
	test.S(t).ExpectTrue(instances[0].Key.Equals(&i830Key))
	test.S(t).ExpectTrue(instances[1].Key.Equals(&i820Key))
	test.S(t).ExpectTrue(instances[2].Key.Equals(&i810Key))
	test.S(t).ExpectTrue(instances[3].Key.Equals(&i730Key))
	test.S(t).ExpectTrue(instances[4].Key.Equals(&i720Key))
	test.S(t).ExpectTrue(instances[5].Key.Equals(&i710Key))
}

func TestIsGenerallyValidAsBinlogSource(t *testing.T) {
	instances, _ := generateTestInstances()
	for _, instance := range instances {
		test.S(t).ExpectFalse(isGenerallyValidAsBinlogSource(instance))
	}
	for _, instance := range instances {
		instance.IsLastCheckValid = true
		instance.LogBinEnabled = true
		instance.LogSlaveUpdatesEnabled = true
	}
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
	for _, instance := range instances {
		instance.IsLastCheckValid = true
		instance.LogBinEnabled = true
		instance.LogSlaveUpdatesEnabled = true
	}
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
	_, _, _, _, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNotNil(err)
}

func TestChooseCandidateSlave(t *testing.T) {
	instances, _ := generateTestInstances()
	for _, instance := range instances {
		instance.IsLastCheckValid = true
		instance.LogBinEnabled = true
		instance.LogSlaveUpdatesEnabled = true
	}
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i830Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 0)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 5)
}

func TestChooseCandidateSlave2(t *testing.T) {
	instances, instancesMap := generateTestInstances()
	for _, instance := range instances {
		instance.IsLastCheckValid = true
		instance.LogBinEnabled = true
		instance.LogSlaveUpdatesEnabled = true
	}
	instancesMap[i830Key.StringCode()].LogSlaveUpdatesEnabled = false
	instancesMap[i820Key.StringCode()].LogBinEnabled = false
	instances = sortedSlaves(instances, false)
	candidate, aheadSlaves, equalSlaves, laterSlaves, err := chooseCandidateSlave(instances)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(candidate.Key, i810Key)
	test.S(t).ExpectEquals(len(aheadSlaves), 2)
	test.S(t).ExpectEquals(len(equalSlaves), 0)
	test.S(t).ExpectEquals(len(laterSlaves), 3)
}
