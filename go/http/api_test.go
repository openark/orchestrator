package http

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-martini/martini"

	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"
	test "github.com/openark/golib/tests"
)

var pathEvaluate = "relocate-slaves"

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.MarkConfigurationLoaded()
	log.SetLevel(log.ERROR)
}

func TestGetSynonymPath(t *testing.T) {
	api := HttpAPI{}

	{
		path := pathEvaluate
		synonym := api.getSynonymPath(path)
		test.S(t).ExpectEquals(synonym, "relocate-replicas")
	}
	{
		path := pathEvaluate + "/:host/:port"
		synonym := api.getSynonymPath(path)
		test.S(t).ExpectEquals(synonym, "relocate-replicas/:host/:port")
	}
}

func TestKnownPaths(t *testing.T) {
	m := martini.Classic()
	api := HttpAPI{}

	api.RegisterRequests(m)

	pathsMap := make(map[string]bool)
	for _, path := range registeredPaths {
		pathBase := strings.Split(path, "/")[0]
		pathsMap[pathBase] = true
	}
	test.S(t).ExpectTrue(pathsMap["health"])
	test.S(t).ExpectTrue(pathsMap["lb-check"])
	test.S(t).ExpectTrue(pathsMap["relocate"])
	test.S(t).ExpectTrue(pathsMap["relocate-slaves"])

	for path, synonym := range apiSynonyms {
		test.S(t).ExpectTrue(pathsMap[path])
		test.S(t).ExpectTrue(pathsMap[synonym])
	}

	config.Config.WebInterfaceReadOnly = true
	test.S(t).ExpectEquals(isForcedReadOnly(fmt.Sprintf("%s/api/%s", config.Config.URLPrefix, api.getSynonymPath(pathEvaluate))), false)
}
