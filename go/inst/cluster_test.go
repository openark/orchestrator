/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package inst

import (
	"fmt"

	"github.com/openark/golib/log"
	test "github.com/openark/golib/tests"
	"github.com/openark/orchestrator/go/config"
	"testing"
)

var mainKey = InstanceKey{Hostname: "host1", Port: 3306}

func init() {
	config.Config.HostnameResolveMethod = "none"
	config.Config.KVClusterMainPrefix = "test/main/"
	config.MarkConfigurationLoaded()
	log.SetLevel(log.ERROR)
}

func TestGetClusterMainKVKey(t *testing.T) {
	kvKey := GetClusterMainKVKey("foo")
	test.S(t).ExpectEquals(kvKey, "test/main/foo")
}

func TestGetClusterMainKVPair(t *testing.T) {
	{
		kvPair := getClusterMainKVPair("myalias", &mainKey)
		test.S(t).ExpectNotNil(kvPair)
		test.S(t).ExpectEquals(kvPair.Key, "test/main/myalias")
		test.S(t).ExpectEquals(kvPair.Value, mainKey.StringCode())
	}
	{
		kvPair := getClusterMainKVPair("", &mainKey)
		test.S(t).ExpectTrue(kvPair == nil)
	}
	{
		kvPair := getClusterMainKVPair("myalias", nil)
		test.S(t).ExpectTrue(kvPair == nil)
	}
}

func TestGetClusterMainKVPairs(t *testing.T) {
	kvPairs := GetClusterMainKVPairs("myalias", &mainKey)
	test.S(t).ExpectTrue(len(kvPairs) >= 2)

	{
		kvPair := kvPairs[0]
		test.S(t).ExpectEquals(kvPair.Key, "test/main/myalias")
		test.S(t).ExpectEquals(kvPair.Value, mainKey.StringCode())
	}
	{
		kvPair := kvPairs[1]
		test.S(t).ExpectEquals(kvPair.Key, "test/main/myalias/hostname")
		test.S(t).ExpectEquals(kvPair.Value, mainKey.Hostname)
	}
	{
		kvPair := kvPairs[2]
		test.S(t).ExpectEquals(kvPair.Key, "test/main/myalias/port")
		test.S(t).ExpectEquals(kvPair.Value, fmt.Sprintf("%d", mainKey.Port))
	}
}

func TestGetClusterMainKVPairs2(t *testing.T) {
	kvPairs := GetClusterMainKVPairs("", &mainKey)
	test.S(t).ExpectEquals(len(kvPairs), 0)
}
