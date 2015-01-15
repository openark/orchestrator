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
	"regexp"
)

// Event entries may contains table IDs (can be different for same tables on different servers)
// and also COMMIT transaction IDs (different values on different servers).
// So these need to be removed from the event entry if we're to compare and validate matching
// entries.
var eventInfoTransformations map[*regexp.Regexp]string = map[*regexp.Regexp]string{
	regexp.MustCompile(`(.*) [/][*].*?[*][/](.*$)`): "$1 $2",
	regexp.MustCompile(`(COMMIT) .*$`):              "$1",
	regexp.MustCompile(`(table_id:) [0-9]+ (.*$)`):  "$1 ### $2",
}

var skippedEventTypes map[string]bool = map[string]bool{
	"Format_desc": true,
	"Stop":        true,
	"Rotate":      true,
}

type BinlogEvent struct {
	Coordinates  BinlogCoordinates
	NextEventPos int64
	EventType    string
	Info         string
}

//
func (this *BinlogEvent) NextBinlogCoordinates() BinlogCoordinates {
	return BinlogCoordinates{LogFile: this.Coordinates.LogFile, LogPos: this.NextEventPos, Type: this.Coordinates.Type}
}

//
func (this *BinlogEvent) NormalizeInfo() {
	for reg, replace := range eventInfoTransformations {
		this.Info = reg.ReplaceAllString(this.Info, replace)
	}
}

//
type BinlogEventCursor struct {
	cachedEvents      []BinlogEvent
	currentEventIndex int
	fetchNextEvents   func(BinlogCoordinates) ([]BinlogEvent, error)
	nextCoordinates   BinlogCoordinates
}

// fetchNextEventsFunc expected to return events starting at a given position, and automatically fetch those from next
// binary log when no more rows are found in current log.
// It is expected to return empty array with no error upon end of binlogs
// It is expected to return error upon error...
func NewBinlogEventCursor(startCoordinates BinlogCoordinates, fetchNextEventsFunc func(BinlogCoordinates) ([]BinlogEvent, error)) BinlogEventCursor {
	events, _ := fetchNextEventsFunc(startCoordinates)
	return BinlogEventCursor{
		cachedEvents:      events,
		currentEventIndex: -1,
		fetchNextEvents:   fetchNextEventsFunc,
	}
}

// NextEvent will return the next event entry from binary logs; it will automatically skip to next
// binary log if need be.
// Internally, it uses the cachedEvents array, so that it does not go to the MySQL server upon each call.
// Returns nil upon reaching end of binary logs.
func (this *BinlogEventCursor) NextEvent() (*BinlogEvent, error) {
	if len(this.cachedEvents) == 0 {
		// End of logs
		return nil, nil
	}
	if this.currentEventIndex+1 < len(this.cachedEvents) {
		// We have enough cache to go by
		this.currentEventIndex++
		event := &this.cachedEvents[this.currentEventIndex]
		this.nextCoordinates = event.NextBinlogCoordinates()
		return event, nil
	} else {
		// Cache exhausted; get next bulk of entries and return the next entry
		var err error
		this.cachedEvents, err = this.fetchNextEvents(this.cachedEvents[len(this.cachedEvents)-1].NextBinlogCoordinates())
		if err != nil {
			return nil, err
		}
		this.currentEventIndex = -1
		// While this seems recursive do note that recursion level is at most 1, since we either have
		// entires in the next binlog (no further recursion) or we don't (immediate termination)
		return this.NextEvent()
	}
}

// NextRealEvent returns the next event from binlog that is not meta/control event (these are start-of-binary-log,
// rotate-binary-log etc.)
func (this *BinlogEventCursor) NextRealEvent() (*BinlogEvent, error) {
	event, err := this.NextEvent()
	if err != nil {
		return event, err
	}
	if event == nil {
		return event, err
	}
	if _, found := skippedEventTypes[event.EventType]; found {
		// Recursion will not be deep here. A few entries (end-of-binlog followed by start-of-bin-log) are possible,
		// but we really don't expect a huge sequence of those.
		return this.NextRealEvent()
	}
	event.NormalizeInfo()
	return event, err
}

// NextCoordinates return the binlog coordinates of the next entry as yet unprocessed by the cursor.
// Moreover, when the cursor terminates (consumes last entry), these coordinates indicate what will be the futuristic
// coordinates of the next binlog entry.
// The value of this function is used by match-below to move a slave behind another, after exhausting the shared binlog
// entries of both.
func (this *BinlogEventCursor) NextCoordinates() (BinlogCoordinates, error) {
	if this.nextCoordinates.LogPos == 0 {
		return this.nextCoordinates, errors.New("Next coordinates unfound")
	}
	return this.nextCoordinates, nil
}
