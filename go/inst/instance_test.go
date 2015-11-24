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
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
	. "gopkg.in/check.v1"
	"testing"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	log.SetLevel(log.ERROR)
}

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

var key1 inst.InstanceKey = inst.InstanceKey{Hostname: "host1", Port: 3306}
var key2 inst.InstanceKey = inst.InstanceKey{Hostname: "host2", Port: 3306}

var instance1 inst.Instance = inst.Instance{Key: key1}
var instance2 inst.Instance = inst.Instance{Key: key2}

func (s *TestSuite) TestInstanceKeyEquals(c *C) {
	i1 := inst.Instance{
		Key: inst.InstanceKey{
			Hostname: "sql00.db",
			Port:     3306,
		},
		Version: "5.6",
	}
	i2 := inst.Instance{
		Key: inst.InstanceKey{
			Hostname: "sql00.db",
			Port:     3306,
		},
		Version: "5.5",
	}

	c.Assert(i1.Key, Equals, i2.Key)

	i2.Key.Port = 3307
	c.Assert(i1.Key, Not(Equals), i2.Key)
}

func (s *TestSuite) TestIsSmallerMajorVersion(c *C) {
	i55 := inst.Instance{Version: "5.5"}
	i5517 := inst.Instance{Version: "5.5.17"}
	i56 := inst.Instance{Version: "5.6"}

	c.Assert(i55.IsSmallerMajorVersion(&i5517), Not(Equals), true)
	c.Assert(i56.IsSmallerMajorVersion(&i5517), Not(Equals), true)
	c.Assert(i55.IsSmallerMajorVersion(&i56), Equals, true)
}

func (s *TestSuite) TestIsVersion(c *C) {
	i51 := inst.Instance{Version: "5.1.19"}
	i55 := inst.Instance{Version: "5.5.17-debug"}
	i56 := inst.Instance{Version: "5.6.20"}
	i57 := inst.Instance{Version: "5.7.8-log"}

	c.Assert(i51.IsMySQL51(), Equals, true)
	c.Assert(i55.IsMySQL55(), Equals, true)
	c.Assert(i56.IsMySQL56(), Equals, true)
	c.Assert(i55.IsMySQL56(), Equals, false)
	c.Assert(i57.IsMySQL57(), Equals, true)
	c.Assert(i56.IsMySQL57(), Equals, false)
}

func (s *TestSuite) TestBinlogCoordinates(c *C) {
	c1 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 104}
	c2 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 104}
	c3 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 5000}
	c4 := inst.BinlogCoordinates{LogFile: "mysql-bin.00112", LogPos: 104}

	c.Assert(c1.Equals(&c2), Equals, true)
	c.Assert(c1.Equals(&c3), Equals, false)
	c.Assert(c1.Equals(&c4), Equals, false)
	c.Assert(c1.SmallerThan(&c2), Equals, false)
	c.Assert(c1.SmallerThan(&c3), Equals, true)
	c.Assert(c1.SmallerThan(&c4), Equals, true)
	c.Assert(c3.SmallerThan(&c4), Equals, true)
	c.Assert(c3.SmallerThan(&c2), Equals, false)
	c.Assert(c4.SmallerThan(&c2), Equals, false)
	c.Assert(c4.SmallerThan(&c3), Equals, false)

	c.Assert(c1.SmallerThanOrEquals(&c2), Equals, true)
	c.Assert(c1.SmallerThanOrEquals(&c3), Equals, true)
}

func (s *TestSuite) TestBinlogPrevious(c *C) {
	c1 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 104}
	cres, err := c1.PreviousFileCoordinates()

	c.Assert(err, IsNil)
	c.Assert(c1.Type, Equals, cres.Type)
	c.Assert(cres.LogFile, Equals, "mysql-bin.00016")

	c2 := inst.BinlogCoordinates{LogFile: "mysql-bin.00100", LogPos: 104}
	cres, err = c2.PreviousFileCoordinates()

	c.Assert(err, IsNil)
	c.Assert(c1.Type, Equals, cres.Type)
	c.Assert(cres.LogFile, Equals, "mysql-bin.00099")

	c3 := inst.BinlogCoordinates{LogFile: "mysql.00.prod.com.00100", LogPos: 104}
	cres, err = c3.PreviousFileCoordinates()

	c.Assert(err, IsNil)
	c.Assert(c1.Type, Equals, cres.Type)
	c.Assert(cres.LogFile, Equals, "mysql.00.prod.com.00099")

	c4 := inst.BinlogCoordinates{LogFile: "mysql.00.prod.com.00000", LogPos: 104}
	_, err = c4.PreviousFileCoordinates()

	c.Assert(err, Not(IsNil))
}

func (s *TestSuite) TestBinlogCoordinatesAsKey(c *C) {
	m := make(map[inst.BinlogCoordinates]bool)

	c1 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 104}
	c2 := inst.BinlogCoordinates{LogFile: "mysql-bin.00022", LogPos: 104}
	c3 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 104}
	c4 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 222}

	m[c1] = true
	m[c2] = true
	m[c3] = true
	m[c4] = true

	c.Assert(len(m), Equals, 3)
}

func (s *TestSuite) TestBinlogFileNumber(c *C) {
	c1 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 104}
	c2 := inst.BinlogCoordinates{LogFile: "mysql-bin.00022", LogPos: 104}

	c.Assert(c1.FileNumberDistance(&c1), Equals, 0)
	c.Assert(c1.FileNumberDistance(&c2), Equals, 5)
	c.Assert(c2.FileNumberDistance(&c1), Equals, -5)
}

func (s *TestSuite) TestBinlogFileNumberDistance(c *C) {
	c1 := inst.BinlogCoordinates{LogFile: "mysql-bin.00017", LogPos: 104}
	fileNum, numLen := c1.FileNumber()

	c.Assert(fileNum, Equals, 17)
	c.Assert(numLen, Equals, 5)
}

func (s *TestSuite) TestCanReplicateFrom(c *C) {
	i55 := inst.Instance{Key: key1, Version: "5.5"}
	i56 := inst.Instance{Key: key2, Version: "5.6"}

	var canReplicate bool
	canReplicate, _ = i56.CanReplicateFrom(&i55)
	c.Assert(canReplicate, Equals, false) //binlog not yet enabled

	i55.LogBinEnabled = true
	i55.LogSlaveUpdatesEnabled = true
	i56.LogBinEnabled = true
	i56.LogSlaveUpdatesEnabled = true

	canReplicate, _ = i56.CanReplicateFrom(&i55)
	c.Assert(canReplicate, Equals, false) //serverid not set
	i55.ServerID = 55
	i56.ServerID = 56

	canReplicate, err := i56.CanReplicateFrom(&i55)
	c.Assert(err, IsNil)
	c.Assert(canReplicate, Equals, true)
	canReplicate, _ = i55.CanReplicateFrom(&i56)
	c.Assert(canReplicate, Equals, false)

	iStatement := inst.Instance{Key: key1, Binlog_format: "STATEMENT", ServerID: 1, Version: "5.5", LogBinEnabled: true, LogSlaveUpdatesEnabled: true}
	iRow := inst.Instance{Key: key2, Binlog_format: "ROW", ServerID: 2, Version: "5.5", LogBinEnabled: true, LogSlaveUpdatesEnabled: true}
	canReplicate, err = iRow.CanReplicateFrom(&iStatement)
	c.Assert(err, IsNil)
	c.Assert(canReplicate, Equals, true)
	canReplicate, _ = iStatement.CanReplicateFrom(&iRow)
	c.Assert(canReplicate, Equals, false)
}

func (s *TestSuite) TestNewInstanceKeyFromStrings(c *C) {
	i, err := inst.NewInstanceKeyFromStrings("127.0.0.1", "3306")
	c.Assert(err, IsNil)
	c.Assert(i.Hostname, Equals, "127.0.0.1")
	c.Assert(i.Port, Equals, 3306)
}

func (s *TestSuite) TestNewInstanceKeyFromStringsFail(c *C) {
	_, err := inst.NewInstanceKeyFromStrings("127.0.0.1", "3306x")
	c.Assert(err, Not(IsNil))
}

func (s *TestSuite) TestParseInstanceKey(c *C) {
	i, err := inst.ParseInstanceKey("127.0.0.1:3306")
	c.Assert(err, IsNil)
	c.Assert(i.Hostname, Equals, "127.0.0.1")
	c.Assert(i.Port, Equals, 3306)
}

func (s *TestSuite) TestInstanceKeyValid(c *C) {
	c.Assert(key1.IsValid(), Equals, true)
	i, err := inst.ParseInstanceKey("_:3306")
	c.Assert(err, IsNil)
	c.Assert(i.IsValid(), Equals, false)
	i, err = inst.ParseInstanceKey("//myhost:3306")
	c.Assert(err, Not(IsNil))
	c.Assert(i.IsValid(), Equals, false)
}

func (s *TestSuite) TestInstanceKeyDetach(c *C) {
	c.Assert(key1.IsDetached(), Equals, false)
	detached1 := key1.DetachedKey()
	c.Assert(detached1.IsDetached(), Equals, true)
	detached2 := key1.DetachedKey()
	c.Assert(detached2.IsDetached(), Equals, true)
	c.Assert(detached1.Equals(detached2), Equals, true)

	reattached1 := detached1.ReattachedKey()
	c.Assert(reattached1.IsDetached(), Equals, false)
	c.Assert(reattached1.Equals(&key1), Equals, true)
	reattached2 := reattached1.ReattachedKey()
	c.Assert(reattached2.IsDetached(), Equals, false)
	c.Assert(reattached1.Equals(reattached2), Equals, true)
}

func (s *TestSuite) TestInstanceKeyMapToJSON(c *C) {
	m := *inst.NewInstanceKeyMap()
	m.AddKey(key1)
	m.AddKey(key2)
	json, err := m.ToJSON()
	c.Assert(err, IsNil)
	ok := (json == `[{"Hostname":"host1","Port":3306},{"Hostname":"host2","Port":3306}]`) || (json == `[{"Hostname":"host2","Port":3306},{"Hostname":"host1","Port":3306}]`)
	c.Assert(ok, Equals, true)
}

func (s *TestSuite) TestInstanceKeyMapReadJSON(c *C) {
	json := `[{"Hostname":"host1","Port":3306},{"Hostname":"host2","Port":3306}]`
	m := *inst.NewInstanceKeyMap()
	m.ReadJson(json)
	c.Assert(len(m), Equals, 2)
	c.Assert(m[key1], Equals, true)
	c.Assert(m[key2], Equals, true)
}

func (s *TestSuite) TestEmptyInstanceKeyMapToCommaDelimitedList(c *C) {
	m := *inst.NewInstanceKeyMap()
	res := m.ToCommaDelimitedList()

	c.Assert(res, Equals, "")
}

func (s *TestSuite) TestInstanceKeyMapToCommaDelimitedList(c *C) {
	m := *inst.NewInstanceKeyMap()
	m.AddKey(key1)
	m.AddKey(key2)
	res := m.ToCommaDelimitedList()

	ok := (res == `host1:3306,host2:3306`) || (res == `host2:3306,host1:3306`)
	c.Assert(ok, Equals, true)
}

func (s *TestSuite) TestNextGTID(c *C) {
	{
		i := inst.Instance{ExecutedGtidSet: "4f6d62ed-df65-11e3-b395-60672090eb04:1,b9b4712a-df64-11e3-b391-60672090eb04:1-6"}
		nextGTID, err := i.NextGTID()
		c.Assert(err, IsNil)
		c.Assert(nextGTID, Equals, "b9b4712a-df64-11e3-b391-60672090eb04:7")
	}
	{
		i := inst.Instance{ExecutedGtidSet: "b9b4712a-df64-11e3-b391-60672090eb04:1-6"}
		nextGTID, err := i.NextGTID()
		c.Assert(err, IsNil)
		c.Assert(nextGTID, Equals, "b9b4712a-df64-11e3-b391-60672090eb04:7")
	}
	{
		i := inst.Instance{ExecutedGtidSet: "b9b4712a-df64-11e3-b391-60672090eb04:6"}
		nextGTID, err := i.NextGTID()
		c.Assert(err, IsNil)
		c.Assert(nextGTID, Equals, "b9b4712a-df64-11e3-b391-60672090eb04:7")
	}
}

func (s *TestSuite) TestOracleGTIDSet(c *C) {
	{
		gtidSetString := `230ea8ea-81e3-11e4-972a-e25ec4bd140a:1-10539,
316d193c-70e5-11e5-adb2-ecf4bb2262ff:1-8935:8984-6124596,
321f5c0d-70e5-11e5-adb2-ecf4bb2262ff:1-56`
		gtidSet, err := inst.ParseGtidSet(gtidSetString)
		c.Assert(err, IsNil)
		c.Assert(len(gtidSet.GtidEntries), Equals, 3)
		c.Assert(gtidSet.String(), Equals, `230ea8ea-81e3-11e4-972a-e25ec4bd140a:1-10539,
316d193c-70e5-11e5-adb2-ecf4bb2262ff:1-8935:8984-6124596,
321f5c0d-70e5-11e5-adb2-ecf4bb2262ff:1-56`)

		c.Assert(gtidSet.GtidEntries[0].UUID, Equals, `230ea8ea-81e3-11e4-972a-e25ec4bd140a`)
		c.Assert(gtidSet.GtidEntries[1].UUID, Equals, `316d193c-70e5-11e5-adb2-ecf4bb2262ff`)
		c.Assert(gtidSet.GtidEntries[2].UUID, Equals, `321f5c0d-70e5-11e5-adb2-ecf4bb2262ff`)
		c.Assert(gtidSet.GtidEntries[1].Ranges, Equals, `1-8935:8984-6124596`)

		removed := gtidSet.RemoveUUID(`ffffffff-70e5-11e5-adb2-ecf4bb2262ff`)
		c.Assert(removed, Equals, false)
		c.Assert(len(gtidSet.GtidEntries), Equals, 3)

		removed = gtidSet.RemoveUUID(`316d193c-70e5-11e5-adb2-ecf4bb2262ff`)
		c.Assert(removed, Equals, true)
		c.Assert(len(gtidSet.GtidEntries), Equals, 2)

		c.Assert(gtidSet.String(), Equals, `230ea8ea-81e3-11e4-972a-e25ec4bd140a:1-10539,
321f5c0d-70e5-11e5-adb2-ecf4bb2262ff:1-56`)

		removed = gtidSet.RemoveUUID(`316d193c-70e5-11e5-adb2-ecf4bb2262ff`)
		c.Assert(removed, Equals, false)
		c.Assert(len(gtidSet.GtidEntries), Equals, 2)

		removed = gtidSet.RemoveUUID(`321f5c0d-70e5-11e5-adb2-ecf4bb2262ff`)
		c.Assert(removed, Equals, true)
		c.Assert(len(gtidSet.GtidEntries), Equals, 1)
		c.Assert(gtidSet.String(), Equals, `230ea8ea-81e3-11e4-972a-e25ec4bd140a:1-10539`)

		removed = gtidSet.RemoveUUID(`230ea8ea-81e3-11e4-972a-e25ec4bd140a`)
		c.Assert(removed, Equals, true)
		c.Assert(len(gtidSet.GtidEntries), Equals, 0)

		removed = gtidSet.RemoveUUID(`230ea8ea-81e3-11e4-972a-e25ec4bd140a`)
		c.Assert(removed, Equals, false)
		c.Assert(len(gtidSet.GtidEntries), Equals, 0)
		c.Assert(gtidSet.String(), Equals, ``)
	}
	{
		gtidSetString := `230ea8ea-81e3-11e4-972a-e25ec4bd140a:1-10539,
:1-8935:8984-6124596,
321f5c0d-70e5-11e5-adb2-ecf4bb2262ff:1-56`
		_, err := inst.ParseGtidSet(gtidSetString)
		c.Assert(err, Not(IsNil))
	}
}

func (s *TestSuite) TestRemoveInstance(c *C) {
	{
		instances := [](*inst.Instance){&instance1, &instance2}
		c.Assert(len(instances), Equals, 2)
		instances = inst.RemoveNilInstances(instances)
		c.Assert(len(instances), Equals, 2)
	}
	{
		instances := [](*inst.Instance){&instance1, nil, &instance2}
		c.Assert(len(instances), Equals, 3)
		instances = inst.RemoveNilInstances(instances)
		c.Assert(len(instances), Equals, 2)
	}
	{
		instances := [](*inst.Instance){&instance1, &instance2}
		c.Assert(len(instances), Equals, 2)
		instances = inst.RemoveInstance(instances, &key1)
		c.Assert(len(instances), Equals, 1)
		instances = inst.RemoveInstance(instances, &key1)
		c.Assert(len(instances), Equals, 1)
		instances = inst.RemoveInstance(instances, &key2)
		c.Assert(len(instances), Equals, 0)
		instances = inst.RemoveInstance(instances, &key2)
		c.Assert(len(instances), Equals, 0)
	}
}
