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
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/math"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
	"github.com/pmylund/go-cache"
	"regexp"
	"strings"
	"time"
)

const maxEmptyBinlogFiles int = 10
const maxEventInfoDisplayLength int = 200

var instancePseudoGTIDEntryCache = cache.New(time.Duration(10)*time.Minute, time.Minute)

func getInstancePseudoGTIDKey(instance *Instance, entry string) string {
	return fmt.Sprintf("%s;%s", instance.Key.DisplayString(), entry)
}

// Try and find the last position of a pseudo GTID query entry in the given binary log.
// Also return the full text of that entry.
// maxCoordinates is the position beyond which we should not read. This is relevant when reading relay logs; in particular,
// the last relay log. We must be careful not to scan for Pseudo-GTID entries past the position executed by the SQL thread.
// maxCoordinates == nil means no limit.
func getLastPseudoGTIDEntryInBinlog(instanceKey *InstanceKey, binlog string, binlogType BinlogType, maxCoordinates *BinlogCoordinates) (*BinlogCoordinates, string, error) {
	if binlog == "" {
		return nil, "", log.Errorf("getLastPseudoGTIDEntryInBinlog: empty binlog file name for %+v. maxCoordinates = %+v", *instanceKey, maxCoordinates)
	}
	binlogCoordinates := BinlogCoordinates{LogFile: binlog, LogPos: 0, Type: binlogType}
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return nil, "", err
	}

	moreRowsExpected := true
	var nextPos int64 = 0
	step := 0

	entryText := ""
	for moreRowsExpected {
		query := ""
		if binlogCoordinates.Type == BinaryLog {
			query = fmt.Sprintf("show binlog events in '%s' FROM %d LIMIT %d", binlog, nextPos, config.Config.BinlogEventsChunkSize)
		} else {
			query = fmt.Sprintf("show relaylog events in '%s' LIMIT %d,%d", binlog, (step * config.Config.BinlogEventsChunkSize), config.Config.BinlogEventsChunkSize)
		}

		moreRowsExpected = false
		queryRowsFunc := sqlutils.QueryRowsMap
		if config.Config.BufferBinlogEvents {
			queryRowsFunc = sqlutils.QueryRowsMapBuffered
		}
		err = queryRowsFunc(db, query, func(m sqlutils.RowMap) error {
			moreRowsExpected = true
			nextPos = m.GetInt64("End_log_pos")
			binlogEntryInfo := m.GetString("Info")
			if matched, _ := regexp.MatchString(config.Config.PseudoGTIDPattern, binlogEntryInfo); matched {
				if maxCoordinates != nil && maxCoordinates.SmallerThan(&BinlogCoordinates{LogFile: binlog, LogPos: m.GetInt64("Pos")}) {
					// past the limitation
					moreRowsExpected = false
					return nil
				}
				binlogCoordinates.LogPos = m.GetInt64("Pos")
				entryText = binlogEntryInfo
				// Found a match. But we keep searching: we're interested in the LAST entry, and, alas,
				// we can only search in ASCENDING order...
			}
			return nil
		})
		if err != nil {
			return nil, "", err
		}
		step++
	}

	// Not found? return nil. an error is reserved to SQL problems.
	if binlogCoordinates.LogPos == 0 {
		return nil, "", nil
	}
	return &binlogCoordinates, entryText, err
}

func getLastPseudoGTIDEntryInInstance(instance *Instance, exhaustiveSearch bool) (*BinlogCoordinates, string, error) {
	// Look for last GTID in instance:
	currentBinlog := instance.SelfBinlogCoordinates

	var err error = nil
	for err == nil {
		log.Debugf("Searching for latest pseudo gtid entry in binlog %+v of %+v", currentBinlog.LogFile, instance.Key)
		resultCoordinates, entryInfo, err := getLastPseudoGTIDEntryInBinlog(&instance.Key, currentBinlog.LogFile, BinaryLog, nil)
		if err != nil {
			return nil, "", err
		}
		if resultCoordinates != nil {
			log.Debugf("Found pseudo gtid entry in %+v: %+v", instance.Key, resultCoordinates)
			return resultCoordinates, entryInfo, err
		}
		if !exhaustiveSearch {
			break
		}
		currentBinlog, err = currentBinlog.PreviousFileCoordinates()
	}
	return nil, "", log.Errorf("Cannot find pseudo GTID entry in binlogs of %+v", instance.Key)
}

func getLastPseudoGTIDEntryInRelayLogs(instance *Instance, recordedInstanceRelayLogCoordinates BinlogCoordinates, exhaustiveSearch bool) (*BinlogCoordinates, string, error) {
	// Look for last GTID in relay logs:
	// Since MySQL does not provide with a SHOW RELAY LOGS command, we heuristically srtart from current
	// relay log (indiciated by Relay_log_file) and walk backwards.
	// Eventually we will hit a relay log name which does not exist.
	currentRelayLog := recordedInstanceRelayLogCoordinates
	var err error = nil
	for err == nil {
		log.Debugf("Searching for latest pseudo gtid entry in relaylog %+v of %+v, up to pos %+v", currentRelayLog.LogFile, instance.Key, recordedInstanceRelayLogCoordinates)
		if resultCoordinates, entryInfo, err := getLastPseudoGTIDEntryInBinlog(&instance.Key, currentRelayLog.LogFile, RelayLog, &recordedInstanceRelayLogCoordinates); err != nil {
			return nil, "", err
		} else if resultCoordinates != nil {
			log.Debugf("Found pseudo gtid entry in %+v: %+v", instance.Key, resultCoordinates)
			return resultCoordinates, entryInfo, err
		}
		if !exhaustiveSearch {
			break
		}
		currentRelayLog, err = currentRelayLog.PreviousFileCoordinates()
	}
	return nil, "", log.Errorf("Cannot find pseudo GTID entry in relay logs of %+v", instance.Key)
}

// Given a binlog entry text (query), search it in the given binary log of a given instance
func SearchPseudoGTIDEntryInBinlog(instanceKey *InstanceKey, binlog string, entryText string, entriesMonotonic bool) (BinlogCoordinates, bool, error) {
	binlogCoordinates := BinlogCoordinates{LogFile: binlog, LogPos: 0, Type: BinaryLog}
	if binlog == "" {
		return binlogCoordinates, false, log.Errorf("SearchPseudoGTIDEntryInBinlog: empty binlog file name for %+v", *instanceKey)
	}

	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return binlogCoordinates, false, err
	}

	moreRowsExpected := true
	skipRestOfBinlog := false
	var nextPos int64 = 0

	//	commandToken := math.TernaryString(binlogCoordinates.Type == BinaryLog, "binlog", "relaylog")
	for moreRowsExpected {
		query := fmt.Sprintf("show binlog events in '%s' FROM %d LIMIT %d", binlog, nextPos, config.Config.BinlogEventsChunkSize)

		// We don't know in advance when we will hit the end of the binlog. We will implicitly understand it when our
		// `show binlog events` query does not return any row.
		moreRowsExpected = false
		queryRowsFunc := sqlutils.QueryRowsMap
		if config.Config.BufferBinlogEvents {
			queryRowsFunc = sqlutils.QueryRowsMapBuffered
		}
		err = queryRowsFunc(db, query, func(m sqlutils.RowMap) error {
			if binlogCoordinates.LogPos != 0 {
				// Entry found!
				skipRestOfBinlog = true
				return nil
			}
			if skipRestOfBinlog {
				return nil
			}
			moreRowsExpected = true
			nextPos = m.GetInt64("End_log_pos")
			binlogEntryInfo := m.GetString("Info")
			//
			if binlogEntryInfo == entryText {
				// found it!
				binlogCoordinates.LogPos = m.GetInt64("Pos")
			} else if entriesMonotonic {
				// More heavyweight computation here. Need to verify whether the binlog entry we have is a pseudo-gtid entry
				if matched, _ := regexp.MatchString(config.Config.PseudoGTIDPattern, binlogEntryInfo); matched {
					if binlogEntryInfo > entryText {
						// Entries ascending, and current entry is larger than the one we are searching for.
						// There is no need to scan further on. We can skip the entire binlog
						log.Debugf(`Pseudo GTID entries are monotonic and we hit "%+v" > "%+v"; skipping binlog %+v`, m.GetString("Info"), entryText, binlogCoordinates.LogFile)
						skipRestOfBinlog = true
						return nil
					}
				}
			}
			return nil
		})
		if err != nil {
			return binlogCoordinates, (binlogCoordinates.LogPos != 0), err
		}
		if skipRestOfBinlog {
			return binlogCoordinates, (binlogCoordinates.LogPos != 0), err
		}
	}

	return binlogCoordinates, (binlogCoordinates.LogPos != 0), err
}

// SearchPseudoGTIDEntryInInstance will search for a specific text entry within the binary logs of a given instance.
func SearchPseudoGTIDEntryInInstance(instance *Instance, entryText string, entriesMonotonic bool) (*BinlogCoordinates, error) {
	cacheKey := getInstancePseudoGTIDKey(instance, entryText)
	coords, found := instancePseudoGTIDEntryCache.Get(cacheKey)
	if found {
		// This is wonderful. We can skip the tedious GTID search in the binary log
		log.Debugf("Found instance Pseudo GTID entry coordinates in cache: %+v, %+v, %+v", instance.Key, entryText, coords)
		return coords.(*BinlogCoordinates), nil
	}

	// Look for GTID entry in given instance:
	log.Debugf("Searching for given pseudo gtid entry in %+v. entriesMonotonic=%+v", instance.Key, entriesMonotonic)
	currentBinlog := instance.SelfBinlogCoordinates
	var err error = nil
	for {
		log.Debugf("Searching for given pseudo gtid entry in binlog %+v of %+v", currentBinlog.LogFile, instance.Key)
		// loop iteration per binary log. This might turn to be a heavyweight operation. We wish to throttle the operation such that
		// the instance does not suffer. If it is a slave, we will only act as long as it's not lagging too much.
		if instance.SlaveRunning() {
			for {
				log.Debugf("%+v is a replicating slave. Verifying lag", instance.Key)
				instance, err = ReadTopologyInstance(&instance.Key)
				if err != nil {
					break
				}
				if instance.HasReasonableMaintenanceReplicationLag() {
					// is good to go!
					break
				}
				log.Debugf("lag is too high on %+v. Throttling the search for pseudo gtid entry", instance.Key)
				time.Sleep(time.Duration(config.Config.ReasonableMaintenanceReplicationLagSeconds) * time.Second)
			}
		}
		var resultCoordinates BinlogCoordinates
		var found bool = false
		resultCoordinates, found, err = SearchPseudoGTIDEntryInBinlog(&instance.Key, currentBinlog.LogFile, entryText, entriesMonotonic)
		if err != nil {
			break
		}
		if found {
			log.Debugf("Matched entry in %+v: %+v", instance.Key, resultCoordinates)
			instancePseudoGTIDEntryCache.Set(cacheKey, &resultCoordinates, 0)
			return &resultCoordinates, nil
		}
		// Got here? Unfound. Keep looking
		currentBinlog, err = currentBinlog.PreviousFileCoordinates()
		if err != nil {
			break
		}
		log.Debugf("- Will move next to binlog %+v", currentBinlog.LogFile)
	}

	return nil, log.Errorf("Cannot match pseudo GTID entry in binlogs of %+v; err: %+v", instance.Key, err)
}

// Read (as much as possible of) a chunk of binary log events starting the given startingCoordinates
func readBinlogEventsChunk(instanceKey *InstanceKey, startingCoordinates BinlogCoordinates) ([]BinlogEvent, error) {
	events := []BinlogEvent{}
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return events, err
	}
	commandToken := math.TernaryString(startingCoordinates.Type == BinaryLog, "binlog", "relaylog")
	if startingCoordinates.LogFile == "" {
		return events, log.Errorf("readBinlogEventsChunk: empty binlog file name for %+v.", *instanceKey)
	}
	query := fmt.Sprintf("show %s events in '%s' FROM %d LIMIT %d", commandToken, startingCoordinates.LogFile, startingCoordinates.LogPos, config.Config.BinlogEventsChunkSize)
	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		binlogEvent := BinlogEvent{}
		binlogEvent.Coordinates.LogFile = m.GetString("Log_name")
		binlogEvent.Coordinates.LogPos = m.GetInt64("Pos")
		binlogEvent.Coordinates.Type = startingCoordinates.Type
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
func getNextBinlogEventsChunk(instance *Instance, startingCoordinates BinlogCoordinates, numEmptyBinlogs int) ([]BinlogEvent, error) {
	if numEmptyBinlogs > maxEmptyBinlogFiles {
		log.Debugf("Reached maxEmptyBinlogFiles (%d) at %+v", maxEmptyBinlogFiles, startingCoordinates)
		// Give up and return empty results
		return []BinlogEvent{}, nil
	}
	coordinatesExceededCurrent := false
	switch startingCoordinates.Type {
	case BinaryLog:
		coordinatesExceededCurrent = instance.SelfBinlogCoordinates.FileSmallerThan(&startingCoordinates)
	case RelayLog:
		coordinatesExceededCurrent = instance.RelaylogCoordinates.FileSmallerThan(&startingCoordinates)
	}
	if coordinatesExceededCurrent {
		// We're past the last file. This is a non-error: there are no more events.
		log.Debugf("Coordinates overflow: %+v; terminating search", startingCoordinates)
		return []BinlogEvent{}, nil
	}
	events, err := readBinlogEventsChunk(&instance.Key, startingCoordinates)
	if err != nil {
		return events, err
	}
	if len(events) > 0 {
		log.Debugf("Returning %d events at %+v", len(events), startingCoordinates)
		return events, nil
	}

	// events are empty
	if nextCoordinates, err := instance.GetNextBinaryLog(startingCoordinates); err == nil {
		log.Debugf("Recursing into %+v", nextCoordinates)
		return getNextBinlogEventsChunk(instance, nextCoordinates, numEmptyBinlogs+1)
	}
	// on error
	return events, err
}

// GetNextBinlogCoordinatesToMatch is given a twin-coordinates couple for a would-be slave (instanceKey) and another
// instance (otherKey).
// This is part of the match-below process, and is the heart of the operation: matching the binlog events starting
// the twin-coordinates (where both share the same Pseudo-GTID) until "instance" runs out of entries, hopefully
// before "other" runs out.
// If "other" runs out that means "instance" is more advanced in replication than "other", in which case we can't
// turn it into a slave of "other".
// Otherwise "instance" will point to the *next* binlog entry in "other"
func GetNextBinlogCoordinatesToMatch(instance *Instance, instanceCoordinates BinlogCoordinates, recordedInstanceRelayLogCoordinates BinlogCoordinates,
	other *Instance, otherCoordinates BinlogCoordinates) (*BinlogCoordinates, int, error) {

	fetchNextEvents := func(binlogCoordinates BinlogCoordinates) ([]BinlogEvent, error) {
		return getNextBinlogEventsChunk(instance, binlogCoordinates, 0)
	}
	instanceCursor := NewBinlogEventCursor(instanceCoordinates, fetchNextEvents)

	fetchOtherNextEvents := func(binlogCoordinates BinlogCoordinates) ([]BinlogEvent, error) {
		return getNextBinlogEventsChunk(other, binlogCoordinates, 0)
	}
	otherCursor := NewBinlogEventCursor(otherCoordinates, fetchOtherNextEvents)

	var beautifyCoordinatesLength int = 0
	rpad := func(s string, length int) string {
		if len(s) >= length {
			return s
		}
		return fmt.Sprintf("%s%s", s, strings.Repeat(" ", length-len(s)))
	}

	var lastConsumedEventCoordinates BinlogCoordinates
	var countMatchedEvents int = 0
	for {
		// Exhaust binlogs/relaylogs on instance. While iterating them, also iterate the otherInstance binlogs.
		// We expect entries on both to match, sequentially, until instance's binlogs/relaylogs are exhausted.
		var instanceEventInfo string
		var otherEventInfo string
		{
			// Extract next binlog/relaylog entry from instance:
			event, err := instanceCursor.nextRealEvent()
			if err != nil {
				return nil, 0, log.Errore(err)
			}
			if event != nil {
				lastConsumedEventCoordinates = event.Coordinates
			}

			switch instanceCoordinates.Type {
			case BinaryLog:
				if event == nil {
					// end of binary logs for instance:
					targetMatchCoordinates, err := otherCursor.getNextCoordinates()
					if err != nil {
						return nil, 0, log.Errore(err)
					}
					nextCoordinates, _ := instanceCursor.getNextCoordinates()
					if nextCoordinates.SmallerThan(&instance.SelfBinlogCoordinates) {
						return nil, 0, log.Errorf("Unexpected problem: instance binlog iteration ended before self coordinates. Ended with: %+v, self coordinates: %+v", nextCoordinates, instance.SelfBinlogCoordinates)
					}
					log.Debugf("Reached end of binary logs for instance, at %+v. Other coordinates: %+v", nextCoordinates, targetMatchCoordinates)
					return &targetMatchCoordinates, countMatchedEvents, nil
				}
			case RelayLog:
				// Argghhhh! SHOW RELAY LOG EVENTS IN '...' statement returns CRAPPY values for End_log_pos:
				// instead of returning the end log pos of the current statement in the *relay log*, it shows
				// the end log pos of the matching statement in the *master's binary log*!
				// Yes, there's logic to this. But this means the next-ccordinates are meaningless.
				// As result, in the case where we exhaust (following) the relay log, we cannot do our last
				// nice sanity test that we've indeed reached the Relay_log_pos coordinate; we are only at the
				// last statement, which is SMALLER than Relay_log_pos; and there isn't a "Rotate" entry to make
				// a place holder or anything. The log just ends and we can't be absolutely certain that the next
				// statement is indeed (futuristically) as End_log_pos.
				endOfScan := false
				if event == nil {
					// End of relay log...
					endOfScan = true
					log.Debugf("Reached end of relay log at %+v", recordedInstanceRelayLogCoordinates)
				} else if recordedInstanceRelayLogCoordinates.Equals(&event.Coordinates) {
					// We've passed the maxScanInstanceCoordinates (applies for relay logs)
					endOfScan = true
					log.Debugf("Reached slave relay log coordinates at %+v", recordedInstanceRelayLogCoordinates)
				} else if recordedInstanceRelayLogCoordinates.SmallerThan(&event.Coordinates) {
					return nil, 0, log.Errorf("Unexpected problem: relay log scan passed relay log position without hitting it. Ended with: %+v, relay log position: %+v", event.Coordinates, recordedInstanceRelayLogCoordinates)
				}
				if endOfScan {
					// end of binary logs for instance:
					targetMatchCoordinates, err := otherCursor.getNextCoordinates()
					if err != nil {
						log.Debugf("Cannot otherCursor.getNextCoordinates(). otherCoordinates=%+v, cached events in cursor: %d; index=%d", otherCoordinates, len(otherCursor.cachedEvents), otherCursor.currentEventIndex)
						return nil, 0, log.Errore(err)
					}
					// No further sanity checks (read the above lengthy explanation)
					log.Debugf("Reached limit of relay logs for instance, just after %+v. Other coordinates: %+v", lastConsumedEventCoordinates, targetMatchCoordinates)
					return &targetMatchCoordinates, countMatchedEvents, nil
				}
			}

			instanceEventInfo = event.Info
			coordinatesStr := fmt.Sprintf("%+v", event.Coordinates)
			if len(coordinatesStr) > beautifyCoordinatesLength {
				beautifyCoordinatesLength = len(coordinatesStr)
			}
			log.Debugf("> %+v %+v; %+v", rpad(coordinatesStr, beautifyCoordinatesLength), event.EventType, strings.Split(strings.TrimSpace(instanceEventInfo), "\n")[0])
		}
		{
			// Extract next binlog/relaylog entry from otherInstance (intended master):
			event, err := otherCursor.nextRealEvent()
			if err != nil {
				return nil, 0, log.Errore(err)
			}
			if event == nil {
				// end of binary logs for otherInstance: this is unexpected and means instance is more advanced
				// than otherInstance
				return nil, 0, log.Errorf("Unexpected end of binary logs for assumed master (%+v). This means the instance which attempted to be a slave (%+v) was more advanced. Try the other way round", other.Key, instance.Key)
			}
			otherEventInfo = event.Info
			coordinatesStr := fmt.Sprintf("%+v", event.Coordinates)
			if len(coordinatesStr) > beautifyCoordinatesLength {
				beautifyCoordinatesLength = len(coordinatesStr)
			}
			log.Debugf("< %+v %+v; %+v", rpad(coordinatesStr, beautifyCoordinatesLength), event.EventType, strings.Split(strings.TrimSpace(otherEventInfo), "\n")[0])
		}
		// Verify things are sane (the two extracted entries are identical):
		// (not strictly required by the algorithm but adds such a lovely self-sanity-testing essence)
		if instanceEventInfo != otherEventInfo {
			return nil, 0, log.Errorf("Mismatching entries, aborting: %+v <-> %+v", instanceEventInfo, otherEventInfo)
		}
		countMatchedEvents++
	}
	// Won't get here
}
