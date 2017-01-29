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
	"regexp"
	"strings"
)

type regexpMap struct {
	r           *regexp.Regexp
	replacement string
}

func (this *regexpMap) process(text string) (result string) {
	return this.r.ReplaceAllString(text, this.replacement)
}

var createTableConversions = []regexpMap{
	rmap(`(?i) (character set|charset) [\S]+`, ``),
	rmap(`(?i)int unsigned`, `int`),
	rmap(`(?i)int[\s]*[(][\s]*([0-9]+)[\s]*[)] unsigned`, `int`),
	rmap(`(?i)engine[\s]*=[\s]*(innodb|myisam|ndb|memory|tokudb)`, ``),
	rmap(`(?i)DEFAULT CHARSET[\s]*=[\s]*[\S]+`, ``),
	rmap(`(?i)int( not null|) auto_increment`, `integer`),
	rmap(`(?i)comment '[^']*'`, ``),
	rmap(`(?i)after [\S]+`, ``),
	rmap(`(?i)alter table ([\S]+) add (index|key) ([\S]+) (.+)`, `create index $3_$1 on $1 $4`),
	rmap(`(?i)alter table ([\S]+) add unique (index|key) ([\S]+) (.+)`, `create unique index $3_$1 on $1 $4`),
	rmap(`(?i)([\S]+) enum[\s]*([(].*?[)])`, `$1 text check($1 in $2)`),
	rmap(`(?i)([\s\S]+[/][*] sqlite3-skip [*][/][\s\S]+)`, ``),

	/* sqlite3-skip */
}

var insertConversions = []regexpMap{
	rmap(`(?i)insert ignore`, `insert or ignore`),
	rmap(`(?i)now[(][)]`, `datetime('now')`),
}

var generalConversions = []regexpMap{
	rmap(`(?i)now[(][)][\s]*-[\s]*interval [?] ([\S]+)`, `datetime('now', printf('-%d $1', ?))`),
	rmap(`(?i)now[(][)][\s]*[+][\s]*interval [?] ([\S]+)`, `datetime('now', printf('+%d $1', ?))`),
	rmap(`(?i)now[(][)]`, `datetime('now')`),
}

var (
	identifyCreateTableStatement = regexp.MustCompile(regexpSpaces(`(?i)^[\s]*create table`))
	identifyCreateIndexStatement = regexp.MustCompile(regexpSpaces(`(?i)^[\s]*create( unique|) index`))
	identifyAlterTableStatement  = regexp.MustCompile(regexpSpaces(`(?i)^[\s]*alter table`))
	identifyInsertStatement      = regexp.MustCompile(regexpSpaces(`(?i)^[\s]*(insert|replace)`))
)

func rmap(regexpExpression string, replacement string) regexpMap {
	return regexpMap{
		r:           regexp.MustCompile(regexpSpaces(regexpExpression)),
		replacement: replacement,
	}
}

func regexpSpaces(statement string) string {
	return strings.Replace(statement, " ", `[\s]+`, -1)
}

func IsInsert(statement string) bool {
	return identifyInsertStatement.MatchString(statement)
}

func IsCreateTable(statement string) bool {
	return identifyCreateTableStatement.MatchString(statement)
}

func IsCreateIndex(statement string) bool {
	return identifyCreateIndexStatement.MatchString(statement)
}

func IsAlterTable(statement string) bool {
	return identifyAlterTableStatement.MatchString(statement)
}

func applyConversions(statement string, conversions []regexpMap) string {
	for _, rmap := range conversions {
		statement = rmap.process(statement)
	}
	return statement
}

func ToSqlite3CreateTable(statement string) string {
	return applyConversions(statement, createTableConversions)
}

func ToSqlite3Insert(statement string) string {
	return applyConversions(statement, insertConversions)
}

func ToSqlite3Dialect(statement string) (translated string) {
	if IsCreateTable(statement) {
		return ToSqlite3CreateTable(statement)
	}
	if IsAlterTable(statement) {
		return ToSqlite3CreateTable(statement)
	}
	statement = applyConversions(statement, generalConversions)
	if IsInsert(statement) {
		return ToSqlite3Insert(statement)
	}
	return statement
}
