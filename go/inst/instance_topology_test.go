package inst

import (
	"github.com/outbrain/golib/log"
	test "github.com/outbrain/golib/tests"
	"github.com/outbrain/orchestrator/go/config"
	"testing"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	log.SetLevel(log.ERROR)
}

func TestInitial(t *testing.T) {
	test.S(t).ExpectTrue(true)
}
