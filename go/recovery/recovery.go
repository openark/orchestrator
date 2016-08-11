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

package recovery

// This file holds wrappers around routines to check if global
// recovery is disabled or not.
//
// This is determined by looking in the table
// orchestrator.global_recovery_disable for a value 1.  Note: for
// recoveries to actually happen this must be configured explicitly
// in orchestrator.conf.json. This setting is an emergency brake
// to quickly be able to prevent recoveries happening in some large
// outage type situation.  Maybe this value should be cached etc
// but we won't be doing that many recoveries at once so the load
// on this table is expected to be very low. It should be fine to
// go to the database each time.

import (
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/db"
)

// IsGloballyDisabled returns true if Recoveries are disabled globally
func IsGloballyDisabled() (bool, error) {
	var (
		disabled bool // default is false!
		err      error
	)
	query := `
		SELECT
			COUNT(*) as mycount
		FROM
			global_recovery_disable
		WHERE
			disable_recovery=?
		`
	err = db.QueryOrchestrator(query, sqlutils.Args(1), func(m sqlutils.RowMap) error {
		mycount := m.GetInt("mycount")
		disabled = (mycount > 0)
		return nil
	})
	if err != nil {
		err = log.Errorf("recovery.IsGloballyDisabled(): %v", err)
	}
	return disabled, err
}

// DisableGlobally ensures recoveries are disabled globally
func DisableGlobally() error {
	_, err := db.ExecOrchestrator(`
		INSERT IGNORE INTO global_recovery_disable
			(disable_recovery)
		VALUES  (1)
	`,
	)
	return err
}

// EnableGlobally ensures recoveries are enabled globally
func EnableGlobally() error {
	_, err := db.ExecOrchestrator(`
		DELETE FROM global_recovery_disable -- deliberately no WHERE clause
	`,
	)
	return err
}
