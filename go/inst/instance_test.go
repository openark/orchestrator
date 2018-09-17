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

var key1 = InstanceKey{Hostname: "host1", Port: 3306}
var key2 = InstanceKey{Hostname: "host2", Port: 3306}
var key3 = InstanceKey{Hostname: "host3", Port: 3306}

var instance1 = Instance{Key: key1}
var instance2 = Instance{Key: key2}
var instance3 = Instance{Key: key3}

func TestInstanceKeyEquals(t *testing.T) {
	i1 := Instance{
		Key: InstanceKey{
			Hostname: "sql00.db",
			Port:     3306,
		},
		Version: "5.6",
	}
	i2 := Instance{
		Key: InstanceKey{
			Hostname: "sql00.db",
			Port:     3306,
		},
		Version: "5.5",
	}

	test.S(t).ExpectEquals(i1.Key, i2.Key)

	i2.Key.Port = 3307
	test.S(t).ExpectNotEquals(i1.Key, i2.Key)
}

func TestIsSmallerMajorVersion(t *testing.T) {
	i55 := Instance{Version: "5.5"}
	i5517 := Instance{Version: "5.5.17"}
	i56 := Instance{Version: "5.6"}

	test.S(t).ExpectFalse(i55.IsSmallerMajorVersion(&i5517))
	test.S(t).ExpectFalse(i56.IsSmallerMajorVersion(&i5517))
	test.S(t).ExpectTrue(i55.IsSmallerMajorVersion(&i56))
}

func TestIsVersion(t *testing.T) {
	i51 := Instance{Version: "5.1.19"}
	i55 := Instance{Version: "5.5.17-debug"}
	i56 := Instance{Version: "5.6.20"}
	i57 := Instance{Version: "5.7.8-log"}

	test.S(t).ExpectTrue(i51.IsMySQL51())
	test.S(t).ExpectTrue(i55.IsMySQL55())
	test.S(t).ExpectTrue(i56.IsMySQL56())
	test.S(t).ExpectFalse(i55.IsMySQL56())
	test.S(t).ExpectTrue(i57.IsMySQL57())
	test.S(t).ExpectFalse(i56.IsMySQL57())
}

func TestIsSmallerBinlogFormat(t *testing.T) {
	iStatement := &Instance{Key: key1, Binlog_format: "STATEMENT"}
	iRow := &Instance{Key: key2, Binlog_format: "ROW"}
	iMixed := &Instance{Key: key3, Binlog_format: "MIXED"}
	test.S(t).ExpectTrue(iStatement.IsSmallerBinlogFormat(iRow))
	test.S(t).ExpectFalse(iStatement.IsSmallerBinlogFormat(iStatement))
	test.S(t).ExpectFalse(iRow.IsSmallerBinlogFormat(iStatement))

	test.S(t).ExpectTrue(iStatement.IsSmallerBinlogFormat(iMixed))
	test.S(t).ExpectTrue(iMixed.IsSmallerBinlogFormat(iRow))
	test.S(t).ExpectFalse(iMixed.IsSmallerBinlogFormat(iStatement))
	test.S(t).ExpectFalse(iRow.IsSmallerBinlogFormat(iMixed))
}

func TestCanReplicateFrom(t *testing.T) {
	i55 := Instance{Key: key1, Version: "5.5"}
	i56 := Instance{Key: key2, Version: "5.6"}

	var canReplicate bool
	canReplicate, _ = i56.CanReplicateFrom(&i55)
	test.S(t).ExpectEquals(canReplicate, false) //binlog not yet enabled

	i55.LogBinEnabled = true
	i55.LogSlaveUpdatesEnabled = true
	i56.LogBinEnabled = true
	i56.LogSlaveUpdatesEnabled = true

	canReplicate, _ = i56.CanReplicateFrom(&i55)
	test.S(t).ExpectEquals(canReplicate, false) //serverid not set
	i55.ServerID = 55
	i56.ServerID = 56

	canReplicate, err := i56.CanReplicateFrom(&i55)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectTrue(canReplicate)
	canReplicate, _ = i55.CanReplicateFrom(&i56)
	test.S(t).ExpectFalse(canReplicate)

	iStatement := Instance{Key: key1, Binlog_format: "STATEMENT", ServerID: 1, Version: "5.5", LogBinEnabled: true, LogSlaveUpdatesEnabled: true}
	iRow := Instance{Key: key2, Binlog_format: "ROW", ServerID: 2, Version: "5.5", LogBinEnabled: true, LogSlaveUpdatesEnabled: true}
	canReplicate, err = iRow.CanReplicateFrom(&iStatement)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectTrue(canReplicate)
	canReplicate, _ = iStatement.CanReplicateFrom(&iRow)
	test.S(t).ExpectFalse(canReplicate)
}

func TestNewInstanceKeyFromStrings(t *testing.T) {
	i, err := NewInstanceKeyFromStrings("127.0.0.1", "3306")
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(i.Hostname, "127.0.0.1")
	test.S(t).ExpectEquals(i.Port, 3306)
}

func TestNewInstanceKeyFromStringsFail(t *testing.T) {
	_, err := NewInstanceKeyFromStrings("127.0.0.1", "3306x")
	test.S(t).ExpectNotNil(err)
}

func TestParseInstanceKey(t *testing.T) {
	i, err := ParseInstanceKey("127.0.0.1:3306")
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(i.Hostname, "127.0.0.1")
	test.S(t).ExpectEquals(i.Port, 3306)
}

func TestInstanceKeyValid(t *testing.T) {
	test.S(t).ExpectTrue(key1.IsValid())
	i, err := ParseInstanceKey("_:3306")
	test.S(t).ExpectNil(err)
	test.S(t).ExpectFalse(i.IsValid())
	i, err = ParseInstanceKey("//myhost:3306")
	test.S(t).ExpectNil(err)
	test.S(t).ExpectFalse(i.IsValid())
}

func TestInstanceKeyDetach(t *testing.T) {
	test.S(t).ExpectFalse(key1.IsDetached())
	detached1 := key1.DetachedKey()
	test.S(t).ExpectTrue(detached1.IsDetached())
	detached2 := key1.DetachedKey()
	test.S(t).ExpectTrue(detached2.IsDetached())
	test.S(t).ExpectTrue(detached1.Equals(detached2))

	reattached1 := detached1.ReattachedKey()
	test.S(t).ExpectFalse(reattached1.IsDetached())
	test.S(t).ExpectTrue(reattached1.Equals(&key1))
	reattached2 := reattached1.ReattachedKey()
	test.S(t).ExpectFalse(reattached2.IsDetached())
	test.S(t).ExpectTrue(reattached1.Equals(reattached2))
}

func TestInstanceKeyMapToJSON(t *testing.T) {
	m := *NewInstanceKeyMap()
	m.AddKey(key1)
	m.AddKey(key2)
	json, err := m.ToJSON()
	test.S(t).ExpectNil(err)
	ok := (json == `[{"Hostname":"host1","Port":3306},{"Hostname":"host2","Port":3306}]`) || (json == `[{"Hostname":"host2","Port":3306},{"Hostname":"host1","Port":3306}]`)
	test.S(t).ExpectTrue(ok)
}

func TestInstanceKeyMapReadJSON(t *testing.T) {
	json := `[{"Hostname":"host1","Port":3306},{"Hostname":"host2","Port":3306}]`
	m := *NewInstanceKeyMap()
	m.ReadJson(json)
	test.S(t).ExpectEquals(len(m), 2)
	test.S(t).ExpectTrue(m[key1])
	test.S(t).ExpectTrue(m[key2])
}

func TestEmptyInstanceKeyMapToCommaDelimitedList(t *testing.T) {
	m := *NewInstanceKeyMap()
	res := m.ToCommaDelimitedList()

	test.S(t).ExpectEquals(res, "")
}

func TestInstanceKeyMapToCommaDelimitedList(t *testing.T) {
	m := *NewInstanceKeyMap()
	m.AddKey(key1)
	m.AddKey(key2)
	res := m.ToCommaDelimitedList()

	ok := (res == `host1:3306,host2:3306`) || (res == `host2:3306,host1:3306`)
	test.S(t).ExpectTrue(ok)
}

func TestNextGTID(t *testing.T) {
	{
		i := Instance{ExecutedGtidSet: "4f6d62ed-df65-11e3-b395-60672090eb04:1,b9b4712a-df64-11e3-b391-60672090eb04:1-6"}
		nextGTID, err := i.NextGTID()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(nextGTID, "b9b4712a-df64-11e3-b391-60672090eb04:7")
	}
	{
		i := Instance{ExecutedGtidSet: "b9b4712a-df64-11e3-b391-60672090eb04:1-6"}
		nextGTID, err := i.NextGTID()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(nextGTID, "b9b4712a-df64-11e3-b391-60672090eb04:7")
	}
	{
		i := Instance{ExecutedGtidSet: "b9b4712a-df64-11e3-b391-60672090eb04:6"}
		nextGTID, err := i.NextGTID()
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(nextGTID, "b9b4712a-df64-11e3-b391-60672090eb04:7")
	}
}

func TestRemoveInstance(t *testing.T) {
	{
		instances := [](*Instance){&instance1, &instance2}
		test.S(t).ExpectEquals(len(instances), 2)
		instances = RemoveNilInstances(instances)
		test.S(t).ExpectEquals(len(instances), 2)
	}
	{
		instances := [](*Instance){&instance1, nil, &instance2}
		test.S(t).ExpectEquals(len(instances), 3)
		instances = RemoveNilInstances(instances)
		test.S(t).ExpectEquals(len(instances), 2)
	}
	{
		instances := [](*Instance){&instance1, &instance2}
		test.S(t).ExpectEquals(len(instances), 2)
		instances = RemoveInstance(instances, &key1)
		test.S(t).ExpectEquals(len(instances), 1)
		instances = RemoveInstance(instances, &key1)
		test.S(t).ExpectEquals(len(instances), 1)
		instances = RemoveInstance(instances, &key2)
		test.S(t).ExpectEquals(len(instances), 0)
		instances = RemoveInstance(instances, &key2)
		test.S(t).ExpectEquals(len(instances), 0)
	}
}

func TestHumanReadableDescription(t *testing.T) {
	i57 := Instance{Version: "5.7.8-log"}
	{
		desc := i57.HumanReadableDescription()
		test.S(t).ExpectEquals(desc, "[unknown,invalid,5.7.8-log,rw,nobinlog]")
	}
	{
		i57.UsingPseudoGTID = true
		i57.LogBinEnabled = true
		i57.Binlog_format = "ROW"
		i57.LogSlaveUpdatesEnabled = true
		desc := i57.HumanReadableDescription()
		test.S(t).ExpectEquals(desc, "[unknown,invalid,5.7.8-log,rw,ROW,>>,P-GTID]")
	}
}

func TestTabulatedDescription(t *testing.T) {
	i57 := Instance{Version: "5.7.8-log"}
	{
		desc := i57.TabulatedDescription("|")
		test.S(t).ExpectEquals(desc, "unknown|invalid|5.7.8-log|rw|nobinlog|")
	}
	{
		i57.UsingPseudoGTID = true
		i57.LogBinEnabled = true
		i57.Binlog_format = "ROW"
		i57.LogSlaveUpdatesEnabled = true
		desc := i57.TabulatedDescription("|")
		test.S(t).ExpectEquals(desc, "unknown|invalid|5.7.8-log|rw|ROW|>>,P-GTID")
	}
}
