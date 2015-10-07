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

package logic

import (
	"github.com/outbrain/orchestrator/go/inst"
)

// AsyncRequest represents an entry in the async_request table
type AsyncRequest struct {
	Id                  int64
	Story               string
	Command             string
	OperatedInstanceKey *inst.InstanceKey
	DestinationKey      *inst.InstanceKey
	Pattern             string
	GTIDHint            inst.OperationGTIDHint
}

func NewEmptyAsyncRequest() *AsyncRequest {
	asyncRequest := &AsyncRequest{}
	asyncRequest.GTIDHint = inst.GTIDHintNeutral
	return asyncRequest
}

func NewAsyncRequest(story string, command string, instanceKey *inst.InstanceKey, destinationKey *inst.InstanceKey, pattern string, gtidHint inst.OperationGTIDHint) *AsyncRequest {
	asyncRequest := NewEmptyAsyncRequest()
	asyncRequest.Story = story
	asyncRequest.Command = command
	asyncRequest.OperatedInstanceKey = instanceKey
	asyncRequest.DestinationKey = destinationKey
	asyncRequest.Pattern = pattern
	asyncRequest.GTIDHint = gtidHint
	return asyncRequest
}

func NewSimpleAsyncRequest(story string, command string, instanceKey *inst.InstanceKey) *AsyncRequest {
	asyncRequest := NewEmptyAsyncRequest()
	asyncRequest.Story = story
	asyncRequest.Command = command
	asyncRequest.OperatedInstanceKey = instanceKey
	asyncRequest.DestinationKey = nil
	asyncRequest.Pattern = ""
	return asyncRequest
}
