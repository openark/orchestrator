/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

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

	"github.com/github/orchestrator/go/db"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

func SetInstanceTag(tag *DatabaseInstanceTag) (err error) {
	_, err = db.ExecOrchestrator(`
			insert
				into database_instance_tags (
					hostname, port, tag_name, tag_value, last_updated,
				) VALUES (
					?, ?, ?, ?, NOW()
				)
				on duplicate key update
					tag_value=values(tag_value),
					last_updated=values(last_updated)
			`,
		tag.Key.Hostname,
		tag.Key.Port,
		tag.TagName,
		tag.TagValue,
	)
	return err
}

// EndDowntime will remove downtime flag from an instance
func DeleteInstanceTag(tag *DatabaseInstanceTag) (tagExisted bool, err error) {
	res, err := db.ExecOrchestrator(`
			delete from
				database_instance_tags
			where
				hostname = ?
				and port = ?
				and tag_name = ?
			`,
		tag.Key.Hostname,
		tag.Key.Port,
		tag.TagName,
	)
	if err != nil {
		return tagExisted, log.Errore(err)
	}

	if affected, _ := res.RowsAffected(); affected > 0 {
		tagExisted = true
		AuditOperation("delete-instance-tag", tag.Key, "")
	}
	return tagExisted, err
}

func InstanceTagExists(tag *DatabaseInstanceTag) (tagExists bool, err error) {
	query := `
		select
			count(*) > 0 as tag_exists
		from
			database_instance_tags
		where
			hostname = ?
			and port = ?
			and tag_name = ?
			`
	args := sqlutils.Args(tag.Key.Hostname, tag.Key.Port, tag.TagName)
	err = db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		tagExists = m.GetBool("tagExists")
		return nil
	})

	return tagExists, log.Errore(err)
}

func ReadInstanceTag(tag *DatabaseInstanceTag) (tagExists bool, err error) {
	query := `
		select
			tag_value
		from
			database_instance_tags
		where
			hostname = ?
			and port = ?
			and tag_name = ?
			`
	args := sqlutils.Args(tag.Key.Hostname, tag.Key.Port, tag.TagName)
	err = db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		tag.TagValue = m.GetString("tag_value")
		tagExists = true
		return nil
	})

	return tagExists, log.Errore(err)
}

func GetInstanceKeysByTagName(tagName string) (tagged *InstanceKeyMap, err error) {
	tagged = NewInstanceKeyMap()
	query := `
		select
			hostname,
			port
		from
			database_instance_tags
		where
			tag_name = ?
		`
	args := sqlutils.Args(tagName)
	err = db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		key, _ := NewResolveInstanceKey(m.GetString("hostname"), m.GetInt("port"))
		tagged.AddKey(*key)
		return nil
	})
	return tagged, log.Errore(err)
}

func GetInstanceKeysByTagNameAndValue(tagName, tagValue string) (tagged *InstanceKeyMap, err error) {
	tagged = NewInstanceKeyMap()
	query := `
		select
			hostname,
			port
		from
			database_instance_tags
		where
			tag_name = ?
			and tag_value = ?
		`
	args := sqlutils.Args(tagName, tagValue)
	err = db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		key, _ := NewResolveInstanceKey(m.GetString("hostname"), m.GetInt("port"))
		tagged.AddKey(*key)
		return nil
	})
	return tagged, log.Errore(err)
}

func GetInstanceKeysByTag(tag *DatabaseInstanceTag) (tagged *InstanceKeyMap, err error) {
	args := sqlutils.Args(tag.TagName)
	tagValueClause := ``
	if tag.HasValue {
		tagValueClause = `and tag_value = ?`
		args = append(args, tag.TagValue)
	}
	tagged = NewInstanceKeyMap()
	query := fmt.Sprintf(`
		select
			hostname,
			port
		from
			database_instance_tags
		where
			tag_name = ?
			%s
		`, tagValueClause)
	err = db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		key, _ := NewResolveInstanceKey(m.GetString("hostname"), m.GetInt("port"))
		tagged.AddKey(*key)
		return nil
	})
	return tagged, log.Errore(err)
}
