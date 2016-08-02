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

/*

package discovery manages a queue of discovery requests: an ordered
queue with no duplicates.

push() operation never blocks while pop() blocks on an empty queue.

*/

package discovery

import (
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
	"sync"
	"time"
)

const initialCapacity = 64

type Queue struct {
	sync.Mutex
	queue     chan inst.InstanceKey
	knownKeys map[inst.InstanceKey]time.Time
}

func NewQueue() *Queue {
	q := new(Queue)
	q.knownKeys = make(map[inst.InstanceKey]time.Time)
	q.queue = make(chan inst.InstanceKey, initialCapacity)
	return q
}

func (q *Queue) Len() int {
	q.Lock()
	defer q.Unlock()
	return len(q.queue)
}

// push the key to the slice if it does not exist in known keys;
// silently return otherwise.
func (q *Queue) Push(key inst.InstanceKey) {
	q.Lock()
	defer q.Unlock()

	// is it enqueued already?
	if _, found := q.knownKeys[key]; found {
		return
	}

	q.knownKeys[key] = time.Now()

	// Instead of blocking on a full queue, allocate a larger
	// queue if needed. That's not a very ideomatic but
	// 1) we know the queue is at maximum as large as the number
	// of hosts 2) we really don't want a discovery planner to
	// block. Notice: there's a side effect of closing a chan:
	// zero value will be returned to all waiting goroutines.
	// We handle this in Pop(). Good enough for now.
	if len(q.queue) == cap(q.queue) {
		var newQueue = make(chan inst.InstanceKey, cap(q.queue)*2)
		close(q.queue)
		for key := range q.queue {
			newQueue <- key
		}
		q.queue = newQueue
	}

	q.queue <- key
}

var zeroInstanceKey = inst.InstanceKey{}

// pop the entry and remove it from known keys;
// blocks if queue is empty.
func (q *Queue) Pop() inst.InstanceKey {
	var key inst.InstanceKey
	for {
		q.Lock()
		queue := q.queue
		q.Unlock()
		key = <-queue
		// a zero value key may be returned when a chan is closed (see above)
		if key != zeroInstanceKey {
			break
		}
	}

	q.Lock()
	defer q.Unlock()

	// alarm if have been waiting for too long
	timeOnQueue := time.Since(q.knownKeys[key])
	if timeOnQueue > time.Duration(config.Config.InstancePollSeconds)*time.Second {
		log.Warningf("key %v spent %v waiting on a discoveryQueue", key, time.Since(q.knownKeys[key]))
	}

	delete(q.knownKeys, key)
	return key
}
