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

var (
	createTableCharset = rmap(`(?i)varchar[\s]*[(][\s]*([0-9]+)[\s]*[)] (character set|charset) [\S]+`, `varchar(${1})`)
)

var (
	identifyCreateStatement = regexp.MustCompile(regexpSpaces(`(?i)^[\s]*create table`))
)

func rmap(regexpExpression string, replacement string) regexpMap {
	return regexpMap{
		r:           regexp.MustCompile(regexpSpaces(regexpExpression)),
		replacement: replacement,
	}
}

func regexpSpaces(statement string) string {
	return strings.Replace(statement, " ", `[ ]+`, -1)
}

func isCreateTable(statement string) bool {
	return identifyCreateStatement.MatchString(statement)
}

func ToSqlite3CreateTable(statement string) (string, error) {
	statement = createTableCharset.process(statement)
	return statement, nil
}

func ToSqlite3Dialect(statement string) (translated string, err error) {
	return statement, err
}
