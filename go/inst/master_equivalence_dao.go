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
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
	"github.com/openark/orchestrator/go/config"
	"github.com/openark/orchestrator/go/db"
)

func WriteMainPositionEquivalence(main1Key *InstanceKey, main1BinlogCoordinates *BinlogCoordinates,
	main2Key *InstanceKey, main2BinlogCoordinates *BinlogCoordinates) error {
	if main1Key.Equals(main2Key) {
		// Not interesting
		return nil
	}
	writeFunc := func() error {
		_, err := db.ExecOrchestrator(`
        	insert into main_position_equivalence (
        			main1_hostname, main1_port, main1_binary_log_file, main1_binary_log_pos,
        			main2_hostname, main2_port, main2_binary_log_file, main2_binary_log_pos,
        			last_suggested)
        		values (?, ?, ?, ?, ?, ?, ?, ?, NOW()) 
        		on duplicate key update last_suggested=values(last_suggested)
				
				`, main1Key.Hostname, main1Key.Port, main1BinlogCoordinates.LogFile, main1BinlogCoordinates.LogPos,
			main2Key.Hostname, main2Key.Port, main2BinlogCoordinates.LogFile, main2BinlogCoordinates.LogPos,
		)
		return log.Errore(err)
	}
	return ExecDBWriteFunc(writeFunc)
}

func GetEquivalentMainCoordinates(instanceCoordinates *InstanceBinlogCoordinates) (result [](*InstanceBinlogCoordinates), err error) {
	query := `
		select 
				main1_hostname as hostname,
				main1_port as port,
				main1_binary_log_file as binlog_file,
				main1_binary_log_pos as binlog_pos
			from 
				main_position_equivalence
			where
				main2_hostname = ?
				and main2_port = ?
				and main2_binary_log_file = ?
				and main2_binary_log_pos = ?
		union
		select 
				main2_hostname as hostname,
				main2_port as port,
				main2_binary_log_file as binlog_file,
				main2_binary_log_pos as binlog_pos
			from 
				main_position_equivalence
			where
				main1_hostname = ?
				and main1_port = ?
				and main1_binary_log_file = ?
				and main1_binary_log_pos = ?
		`
	args := sqlutils.Args(
		instanceCoordinates.Key.Hostname,
		instanceCoordinates.Key.Port,
		instanceCoordinates.Coordinates.LogFile,
		instanceCoordinates.Coordinates.LogPos,
		instanceCoordinates.Key.Hostname,
		instanceCoordinates.Key.Port,
		instanceCoordinates.Coordinates.LogFile,
		instanceCoordinates.Coordinates.LogPos,
	)

	err = db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		equivalentCoordinates := InstanceBinlogCoordinates{}
		equivalentCoordinates.Key.Hostname = m.GetString("hostname")
		equivalentCoordinates.Key.Port = m.GetInt("port")
		equivalentCoordinates.Coordinates.LogFile = m.GetString("binlog_file")
		equivalentCoordinates.Coordinates.LogPos = m.GetInt64("binlog_pos")

		result = append(result, &equivalentCoordinates)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetEquivalentBinlogCoordinatesFor(instanceCoordinates *InstanceBinlogCoordinates, belowKey *InstanceKey) (*BinlogCoordinates, error) {
	possibleCoordinates, err := GetEquivalentMainCoordinates(instanceCoordinates)
	if err != nil {
		return nil, err
	}
	for _, instanceCoordinates := range possibleCoordinates {
		if instanceCoordinates.Key.Equals(belowKey) {
			return &instanceCoordinates.Coordinates, nil
		}
	}
	return nil, nil
}

// ExpireMainPositionEquivalence expires old main_position_equivalence
func ExpireMainPositionEquivalence() error {
	writeFunc := func() error {
		_, err := db.ExecOrchestrator(`
        	delete from main_position_equivalence 
				where last_suggested < NOW() - INTERVAL ? HOUR
				`, config.Config.UnseenInstanceForgetHours,
		)
		return log.Errore(err)
	}
	return ExecDBWriteFunc(writeFunc)
}
