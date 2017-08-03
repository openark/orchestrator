package orcraft

import (
	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"
	test "github.com/openark/golib/tests"
	"testing"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.MarkConfigurationLoaded()
	log.SetLevel(log.ERROR)
}

func TestPeerAPI(t *testing.T) {
	{
		config.Config.UseSSL = false
		config.Config.ListenAddress = ":2020"
		api := peerAPI("10.0.0.10")
		test.S(t).ExpectEquals(api, "http://10.0.0.10:2020/api")
	}
	{
		config.Config.UseSSL = true
		config.Config.ListenAddress = ":2020"
		api := peerAPI("10.0.0.10")
		test.S(t).ExpectEquals(api, "https://10.0.0.10:2020/api")
	}
}
