package inst

import (
	"strconv"
	"testing"

	test "github.com/openark/golib/tests"
)

func mkTestInstance(id uint, version string) Instance {
	hostname := strconv.FormatUint(uint64(id), 10)
	key := InstanceKey{Hostname: hostname, Port: 3306}
	instance := Instance{Key: key, ServerID: id, LogBinEnabled: true, Version: version, VersionComment: "Unknown"}
	return instance
}

func TestSpecial57FilteringWith55MasterAnd57ReplicaWithBinLog(t *testing.T) {
	master55 := mkTestInstance(1, "5.5.23")
	replica57 := mkTestInstance(2, "5.7.23")

	filterInstance, filterOther, err := special57filterProcessing(&replica57, &master55)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectTrue(filterInstance)
	test.S(t).ExpectFalse(filterOther)
}

func TestSpecial57FilteringWith56MasterAnd57ReplicaWithBinLog(t *testing.T) {
	master56 := mkTestInstance(1, "5.6.23")
	replica57 := mkTestInstance(2, "5.7.23")

	filterInstance, filterOther, err := special57filterProcessing(&replica57, &master56)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectTrue(filterInstance)
	test.S(t).ExpectFalse(filterOther)
}

func TestSpecial57FilteringWithSameVersionMasterReplicaWithBinLog(t *testing.T) {
	versions := []string{"5.5.23", "5.6.48", "5.7.32"}

	for _, version := range versions {
		master := mkTestInstance(1, version)
		replica := mkTestInstance(2, version)

		filterInstance, filterOther, err := special57filterProcessing(&replica, &master)
		test.S(t).ExpectNil(err)
		test.S(t).ExpectFalse(filterInstance)
		test.S(t).ExpectFalse(filterOther)
	}
}
