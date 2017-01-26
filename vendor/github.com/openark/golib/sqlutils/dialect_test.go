/*
   Copyright 2017 GitHub Inc.

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

package sqlutils

import (
	"testing"

	test "github.com/openark/golib/tests"
)

func init() {
}

func TestIsCreateTable(t *testing.T) {
	test.S(t).ExpectTrue(isCreateTable("create table t(id int)"))
	test.S(t).ExpectTrue(isCreateTable(" create table t(id int)"))
	test.S(t).ExpectTrue(isCreateTable("CREATE  TABLE t(id int)"))
	test.S(t).ExpectTrue(isCreateTable(`
		create table t(id int)
		`))
	test.S(t).ExpectFalse(isCreateTable("where create table t(id int)"))
	test.S(t).ExpectFalse(isCreateTable("insert"))
}

func TestToSqlite3CreateTable(t *testing.T) {
	{
		statement := "create table t(id int)"
		result, err := ToSqlite3CreateTable(statement)
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(result, statement)
	}
	{
		statement := "create table t(id int, v varchar(123) CHARACTER SET ascii NOT NULL default '')"
		result, err := ToSqlite3CreateTable(statement)
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(result, "create table t(id int, v varchar(123) NOT NULL default '')")
	}
	{
		statement := "create table t(id int, v varchar ( 123 ) CHARACTER SET ascii NOT NULL default '')"
		result, err := ToSqlite3CreateTable(statement)
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(result, "create table t(id int, v varchar(123) NOT NULL default '')")
	}
}
