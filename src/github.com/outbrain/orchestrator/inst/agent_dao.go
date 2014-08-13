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
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/orchestrator/db"
	"github.com/outbrain/log"
)


// SubmitAgent submits a new agent for listing
func SubmitAgent(hostname string, token string) (string, error) {
	db,	err	:=	db.OpenOrchestrator()
	if err != nil {return "", log.Errore(err)}
	
	_, err = sqlutils.Exec(db, `
			replace 
				into host_agent (
					hostname, token, last_submitted
				) VALUES (
					?, ?, NOW()
				)
			`,
			hostname,
		 	token,
		 )
	if err != nil {return "", log.Errore(err)}
	
	return hostname, err
}
