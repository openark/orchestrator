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
)

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
	return BinlogCoordinates{LogFile: this.Coordinates.LogFile, LogPos: this.NextEventPos}
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

//
func (this *BinlogEventCursor) NextEvent() (*BinlogEvent, error) {
	if len(this.cachedEvents) == 0 {
		return nil, nil
	}
	if this.currentEventIndex+1 < len(this.cachedEvents) {
		this.currentEventIndex++
		event := &this.cachedEvents[this.currentEventIndex]
		this.nextCoordinates = event.NextBinlogCoordinates()
		return event, nil
	} else {
		var err error
		this.cachedEvents, err = this.fetchNextEvents(this.cachedEvents[len(this.cachedEvents)-1].NextBinlogCoordinates())
		if err != nil {
			return nil, err
		}
		this.currentEventIndex = -1
		return this.NextEvent()
	}
}

//
func (this *BinlogEventCursor) NextRealEvent() (*BinlogEvent, error) {
	event, err := this.NextEvent()
	if err != nil {
		return event, err
	}
	if event == nil {
		return event, err
	}
	if _, found := skippedEventTypes[event.EventType]; found {
		return this.NextRealEvent()
	}
	return event, err
}

//
func (this *BinlogEventCursor) NextCoordinates() (BinlogCoordinates, error) {
	if this.nextCoordinates.LogPos == 0 {
		return this.nextCoordinates, errors.New("Next coordinates unfound")
	}
	return this.nextCoordinates, nil
}
