package inst

import (
	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"
	test "github.com/openark/golib/tests"
	"testing"
)

var testCoordinates = BinlogCoordinates{LogFile: "mysql-bin.000010", LogPos: 108}

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.Config.KVClusterMasterPrefix = "test/master/"
	config.MarkConfigurationLoaded()
	log.SetLevel(log.ERROR)
}

func TestDetach(t *testing.T) {
	detachedCoordinates := testCoordinates.Detach()
	test.S(t).ExpectEquals(detachedCoordinates.LogFile, "//mysql-bin.000010:108")
	test.S(t).ExpectEquals(detachedCoordinates.LogPos, testCoordinates.LogPos)
}

func TestDetachedCoordinates(t *testing.T) {
	isDetached, detachedCoordinates := testCoordinates.ExtractDetachedCoordinates()
	test.S(t).ExpectFalse(isDetached)
	test.S(t).ExpectEquals(detachedCoordinates.LogFile, testCoordinates.LogFile)
	test.S(t).ExpectEquals(detachedCoordinates.LogPos, testCoordinates.LogPos)
}

func TestDetachedCoordinates2(t *testing.T) {
	detached := testCoordinates.Detach()
	isDetached, coordinates := detached.ExtractDetachedCoordinates()

	test.S(t).ExpectTrue(isDetached)
	test.S(t).ExpectEquals(coordinates.LogFile, testCoordinates.LogFile)
	test.S(t).ExpectEquals(coordinates.LogPos, testCoordinates.LogPos)
}
