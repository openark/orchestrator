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

package agent

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"errors"
	"encoding/json"
			
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/orchestrator/db"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/log"
)



func readResponse(res *http.Response, err error) ([]byte, error) {
	if err != nil { return nil, err }
	
	defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
	if err != nil { return nil, err }

	if res.Status == "500" { 
		return body, errors.New("Response Status 500")
	}
	
	return body, nil
}

// SubmitAgent submits a new agent for listing
func SubmitAgent(hostname string, port int, token string) (string, error) {
	db,	err	:=	db.OpenOrchestrator()
	if err != nil {return "", log.Errore(err)}
	
	_, err = sqlutils.Exec(db, `
			replace 
				into host_agent (
					hostname, port, token, last_submitted
				) VALUES (
					?, ?, ?, NOW()
				)
			`,
			hostname,
			port,
		 	token,
		 )
	if err != nil {return "", log.Errore(err)}
	
	return hostname, err
}



// ForgetLongUnseenInstances will remove entries of all instacnes that have long since been last seen.
func ForgetLongUnseenAgents() error {
	db,	err	:=	db.OpenOrchestrator()
	if err != nil {return log.Errore(err)}
	
	_, err = sqlutils.Exec(db, `
			delete 
				from host_agent 
			where 
				last_submitted < NOW() - interval ? hour`,
			config.Config.UnseenAgentForgetHours,
		 )
	return err		 
}



// ReadAgents returns a list of all known agents
func ReadAgents() ([]Agent, error) {
	res := []Agent{}
	query := `
		select 
			hostname,
			port,
			token,
			last_submitted
		from 
			host_agent
		order by
			hostname
		`;
	db,	err	:=	db.OpenOrchestrator()
    if err != nil {goto Cleanup}
    
    err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
    	agent := Agent{}
    	agent.Hostname = m.GetString("hostname")
    	agent.Port = m.GetInt("port")
    	agent.Token = "" 
    	agent.LastSubmitted = m.GetString("last_submitted") 

    	res = append(res, agent)
    	return nil       	
   	})
	Cleanup:

	if err != nil	{
		log.Errore(err)
	}
	return res, err

}


// ReadAgents returns a list of all known agents
func GetAgent(hostname string) (Agent, error) {
	agent := Agent{}
	token := ""
	query := fmt.Sprintf(`
		select 
			hostname,
			port,
			token,
			last_submitted
		from 
			host_agent
		where
			hostname = '%s'
		`, hostname);
	db,	err	:=	db.OpenOrchestrator()
    if err != nil {goto Cleanup}
    
    err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
    	agent.Hostname = m.GetString("hostname")
    	agent.Port = m.GetInt("port")
    	agent.LastSubmitted = m.GetString("last_submitted")
    	token = m.GetString("token")
    	
    	return nil
   	})

	if token == "" {
		return agent, log.Errorf("Cannot get agent/token: %s", hostname) 
	}

	// All seems to be in order. Now make some inquiries from orchestrator-agent service:
	{
		uri := fmt.Sprintf("http://%s:%d/api", agent.Hostname, agent.Port)
		log.Debugf("orchestrator-agent uri: %s", uri)

		{		
			availableLocalSnapshotsUri := fmt.Sprintf("%s/available-snapshots-local?token=%s", uri, token)
			body, err := readResponse(http.Get(availableLocalSnapshotsUri))
			if err != nil {return agent, log.Errore(err)}
	
			err = json.Unmarshal(body, &agent.AvailableLocalSnapshots)
			if err != nil {return agent, log.Errore(err)}
		}
		{		
			availableSnapshotsUri := fmt.Sprintf("%s/available-snapshots?token=%s", uri, token)
			body, err := readResponse(http.Get(availableSnapshotsUri))
			if err != nil {return agent, log.Errore(err)}
	
			err = json.Unmarshal(body, &agent.AvailableSnapshots)
			if err != nil {return agent, log.Errore(err)}
		}
		{
			lvSnapshotsUri := fmt.Sprintf("%s/lvs/%s?token=%s", uri, config.Config.SnapshotVolumesFilter, token)		
			body, err := readResponse(http.Get(lvSnapshotsUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.LogicalVolumes)
			}
			//if err != nil {return agent, log.Errore(err)}
		}
		{
			mountUri := fmt.Sprintf("%s/mount?token=%s", uri, token)		
			body, err := readResponse(http.Get(mountUri))
			if err == nil {
				err = json.Unmarshal(body, &agent.MountPoint)
			}
			//if err != nil {return agent, log.Errore(err)}
		}
	}
	Cleanup:
	if err != nil{
		return agent, log.Errore(err)
	}
	return agent, err

}
