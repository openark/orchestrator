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
	"errors"
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/db"
	"regexp"
)

const binlogEventsChunkSize int = 10000

// Try and find the last position of a pseudo GTID query entry in the given binary log.
// Also return the full text of that entry
func GetLastPseudoGTIDEntryInBinlog(instanceKey *InstanceKey, binlog string) (BinlogCoordinates, string, error) {
	binlogCoordinates := BinlogCoordinates{LogFile: binlog, LogPos: 0}
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return binlogCoordinates, "", err
	}

	moreRowsExpected := true
	step := 0

	entryText := ""
	for moreRowsExpected {
		query := fmt.Sprintf("show binlog events in '%s' LIMIT %d,%d", binlog, (step * binlogEventsChunkSize), binlogEventsChunkSize)

		moreRowsExpected = false
		err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
			moreRowsExpected = true
			binlogEntryInfo := m.GetString("Info")
			if matched, _ := regexp.MatchString(config.Config.PseudoGTIDPattern, binlogEntryInfo); matched {
				binlogCoordinates.LogPos = m.GetInt64("Pos")
				entryText = binlogEntryInfo
				// Found a match. But we keep searching: we're interested in the LAST entry, and, alas,
				// we can only search in ASCENDING order...
			}
			return nil
		})
		if err != nil {
			return binlogCoordinates, "", err
		}
		step++
	}

	if binlogCoordinates.LogPos == 0 {
		return binlogCoordinates, "", errors.New(fmt.Sprintf("Cannot find pseudo GTID entry in binlog '%s'", binlog))
	}
	return binlogCoordinates, entryText, err
}

func GetLastPseudoGTIDEntryInInstance(instance *Instance) (*BinlogCoordinates, string, error) {
	// Look for last GTID in instance:
	instanceBinlogs := instance.GetBinaryLogs()

	for i := len(instanceBinlogs) - 1; i >= 0; i-- {
		resultCoordinates, entryInfo, err := GetLastPseudoGTIDEntryInBinlog(&instance.Key, instanceBinlogs[i])
		if resultCoordinates.LogPos != 0 && err == nil {
			log.Debugf("Found pseudo gtid entry in %+v: %+v", instance.Key, resultCoordinates)
			return &resultCoordinates, entryInfo, err
		}
	}
	return nil, "", log.Errorf("Cannot find pseudo GTID entry in binlogs of %+v", instance.Key)
}

// Given a binlog entry text (query), search it in the give nbinary log of a given instance
func SearchPseudoGTIDEntryInBinlog(instanceKey *InstanceKey, binlog string, entryText string) (BinlogCoordinates, error) {
	binlogCoordinates := BinlogCoordinates{LogFile: binlog, LogPos: 0}
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return binlogCoordinates, err
	}

	moreRowsExpected := true
	step := 0

	for moreRowsExpected {
		query := fmt.Sprintf("show binlog events in '%s' LIMIT %d,%d", binlog, (step * binlogEventsChunkSize), binlogEventsChunkSize)
		moreRowsExpected = false
		err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
			if binlogCoordinates.LogPos != 0 {
				return nil
				// moreRowsExpected reamins false, this quits the loop
			}
			moreRowsExpected = true
			if m.GetString("Info") == entryText {
				// found it!
				binlogCoordinates.LogPos = m.GetInt64("Pos")
			}
			return nil
		})
		if err != nil {
			return binlogCoordinates, err
		}
		step++
	}

	if binlogCoordinates.LogPos == 0 {
		return binlogCoordinates, errors.New(fmt.Sprintf("Cannot match pseudo GTID entry in binlog '%s'", binlog))
	}
	return binlogCoordinates, err
}

func SearchPseudoGTIDEntryInInstance(instance *Instance, entryText string) (*BinlogCoordinates, error) {
	// Look for GTID entry in other-instance:
	binlogs := instance.GetBinaryLogs()
	for i := len(binlogs) - 1; i >= 0; i-- {
		resultCoordinates, err := SearchPseudoGTIDEntryInBinlog(&instance.Key, binlogs[i], entryText)
		if resultCoordinates.LogPos != 0 && err == nil {
			log.Debugf("Matched entry in %+v: %+v", instance.Key, resultCoordinates)
			return &resultCoordinates, nil
		}
	}
	return nil, log.Errorf("Cannot match pseudo GTID entry in binlogs of %+v", instance.Key)
}

// Read (as much as possible of) a chink of binary log events starting the given startingCoordinates
func readBinlogEventsChunk(instanceKey *InstanceKey, startingCoordinates BinlogCoordinates) ([]BinlogEvent, error) {
	events := []BinlogEvent{}
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return events, err
	}
	query := fmt.Sprintf("show binlog events in '%s' FROM %d LIMIT %d", startingCoordinates.LogFile, startingCoordinates.LogPos, binlogEventsChunkSize)
	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		binlogEvent := BinlogEvent{}
		binlogEvent.Coordinates.LogFile = m.GetString("Log_name")
		binlogEvent.Coordinates.LogPos = m.GetInt64("Pos")
		binlogEvent.NextEventPos = m.GetInt64("End_log_pos")
		binlogEvent.EventType = m.GetString("Event_type")
		binlogEvent.Info = m.GetString("Info")

		events = append(events, binlogEvent)
		return nil
	})
	return events, err
}

// Return the next chunk of binlog events; skip to next binary log file if need be; return empty result only
// if reached end of binary logs
func getNextBinlogEventsChunk(instance *Instance, startingCoordinates BinlogCoordinates) ([]BinlogEvent, error) {
	events, err := readBinlogEventsChunk(&instance.Key, startingCoordinates)
	if err != nil {
		return events, err
	}
	if len(events) > 0 {
		return events, nil
	}
	// events are empty
	if nextBinlogFile, err := instance.GetNextBinaryLog(startingCoordinates.LogFile); err == nil {
		nextCoordinates := BinlogCoordinates{LogFile: nextBinlogFile, LogPos: 0}
		return getNextBinlogEventsChunk(instance, nextCoordinates)
	}
	// No more log file. We return the empty array: but no error, since there is no error; we've just reached the end.
	// This behaviour is strictly expected by BinlogEventCursor
	return events, nil
}

// GetNextBinlogCoordinatesToMatch is given a twin-coordinates couple for a would-be slave (instanceKey) and another
// instance (otherKey).
// This is part of the match-below process, and is the heart of the operation: matching the binlog events starting
// the twin-coordinates (where both share the same Pseudo-GTID) until "instance" runs out of entries, hopefully
// before "other" runs out.
// If "other" runs out that means "instance" is more advanced in replication than "other", in which case we can't
// turn it into a slave of "other".
// Otherwise "instance" will point to the *next* binlog entry in "other"
func GetNextBinlogCoordinatesToMatch(instance *Instance, instanceCoordinates BinlogCoordinates,
	other *Instance, otherCoordinates BinlogCoordinates) (*BinlogCoordinates, error) {

	fetchNextEvents := func(binlogCoordinates BinlogCoordinates) ([]BinlogEvent, error) {
		return getNextBinlogEventsChunk(instance, binlogCoordinates)
	}
	instanceCursor := NewBinlogEventCursor(instanceCoordinates, fetchNextEvents)

	fetchOtherNextEvents := func(binlogCoordinates BinlogCoordinates) ([]BinlogEvent, error) {
		return getNextBinlogEventsChunk(other, binlogCoordinates)
	}
	otherCursor := NewBinlogEventCursor(otherCoordinates, fetchOtherNextEvents)

	for {
		// Exhaust binlogs on instance. While iterating them, also iterate the otherInstance binlogs.
		// We expect entries on both to match, sequentially, until instance's binlogs are exhausted.
		var instanceEventInfo string
		var otherEventInfo string
		{
			event, err := instanceCursor.NextRealEvent()
			if err != nil {
				return nil, log.Errore(err)
			}
			if event == nil {
				// end of binary logs for instance:
				nextCoordinates, err := otherCursor.NextCoordinates()
				if err != nil {
					return nil, log.Errore(err)
				}
				log.Debugf("Reached end of binary logs for instance. Other coordinates: %+v", nextCoordinates)
				return &nextCoordinates, nil
			}
			instanceEventInfo = event.Info
			log.Debugf("%+v %+v; %+v", event.Coordinates, event.EventType, event.Info)
		}
		{
			event, err := otherCursor.NextRealEvent()
			if err != nil {
				return nil, log.Errore(err)
			}
			if event == nil {
				// end of binary logs for otherInstance: this is unexpected and means instance is more advanced
				// than otherInstance
				return nil, log.Error("Unexpected end of binary logs for assumed master. This means the instance which attempted to be a slave was more advanced. Try the other way round")
			}
			otherEventInfo = event.Info
			log.Debugf("%+v %+v; %+v", event.Coordinates, event.EventType, event.Info)
		}
		// Verify things are sane:
		if instanceEventInfo != otherEventInfo {
			return nil, log.Errorf("Mismatching entries, aborting: %+v <-> %+v", instanceEventInfo, otherEventInfo)
		}
	}

	return nil, log.Error("GetNextBinlogCoordinatesToMatch: unexpected termination")
}
