package http

import (
	"fmt"
	"testing"

	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"
	test "github.com/openark/golib/tests"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.MarkConfigurationLoaded()
	log.SetLevel(log.ERROR)
}

func TestKnownWebPaths(t *testing.T) {
	config.Config.WebInterfaceReadOnly = true
	test.S(t).ExpectTrue(isForcedReadOnly(fmt.Sprintf("%s/web/clusters/", config.Config.URLPrefix)))

	config.Config.WebInterfaceReadOnly = false
	test.S(t).ExpectEquals(isForcedReadOnly(fmt.Sprintf("%s/web/clusters/", config.Config.URLPrefix)), false)
}
