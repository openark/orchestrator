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
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
	. "gopkg.in/check.v1"
	"testing"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
}

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

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

func (s *TestSuite) TestCanReplicateFrom(c *C) {
	i55 := inst.Instance{Version: "5.5"}
	i56 := inst.Instance{Version: "5.6"}

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

	canReplicate, _ = i56.CanReplicateFrom(&i55)
	c.Assert(canReplicate, Equals, true)
	canReplicate, _ = i55.CanReplicateFrom(&i56)
	c.Assert(canReplicate, Equals, false)

	iStatement := inst.Instance{Binlog_format: "STATEMENT", ServerID: 1, Version: "5.5", LogBinEnabled: true, LogSlaveUpdatesEnabled: true}
	iRow := inst.Instance{Binlog_format: "ROW", ServerID: 2, Version: "5.5", LogBinEnabled: true, LogSlaveUpdatesEnabled: true}
	canReplicate, _ = iRow.CanReplicateFrom(&iStatement)
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
