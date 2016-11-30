/*
   Copyright 2016 Simon J Mudd

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

	"github.com/outbrain/golib/sqlutils"

	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
)

// CandidateDatabaseInstance contains information about explicit promotion rules for an instance
type CandidateDatabaseInstance struct {
	Hostname      string
	Port          int
	PromotionRule CandidatePromotionRule
}

// String returns a string representation of the CandidateDatabaseInstance struct
func (cdi CandidateDatabaseInstance) String() string {
	return fmt.Sprintf("%s:%d %s", cdi.Hostname, cdi.Port, cdi.PromotionRule)
}

// BulkReadCandidateDatabaseInstance returns a slice of
// CandidateDatabaseInstance converted to JSON.
/*
root@myorchestrator [orchestrator]> select * from candidate_database_instance;
+-------------------+------+---------------------+----------+----------------+
| hostname          | port | last_suggested      | priority | promotion_rule |
+-------------------+------+---------------------+----------+----------------+
| host1.example.com | 3306 | 2016-11-22 17:41:06 |        1 | prefer         |
| host2.example.com | 3306 | 2016-11-22 17:40:24 |        1 | prefer         |
+-------------------+------+---------------------+----------+----------------+
2 rows in set (0.00 sec)
*/
func BulkReadCandidateDatabaseInstance() ([]CandidateDatabaseInstance, error) {
	if config.Config.DatabaselessMode__experimental {
		return nil, nil // no data to return if not using a database
	}

	var candidateDatabaseInstances []CandidateDatabaseInstance

	// Read all promotion rules from the table
	query := `
SELECT	hostname,
	port,
	promotion_rule -- no munging done here yet
FROM	candidate_database_instance
`
	err := db.QueryOrchestrator(query, nil, func(m sqlutils.RowMap) error {
		cdi := CandidateDatabaseInstance{
			Hostname:      m.GetString("hostname"),
			Port:          m.GetInt("port"),
			PromotionRule: CandidatePromotionRule(m.GetString("promotion_rule")),
		}
		// add to end of candidateDatabaseInstances
		candidateDatabaseInstances = append(candidateDatabaseInstances, cdi)

		return nil
	})
	return candidateDatabaseInstances, err
}
