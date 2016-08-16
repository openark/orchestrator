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

// Size of discovery queue. Should be greater than the total number
// of machines being discovered.
const queueCapacity = 100000

type Queue struct {
	sync.Mutex
	queue     chan inst.InstanceKey
	knownKeys map[inst.InstanceKey]time.Time
}

func NewQueue() *Queue {
	q := new(Queue)
	q.knownKeys = make(map[inst.InstanceKey]time.Time)
	q.queue = make(chan inst.InstanceKey, queueCapacity)
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
	q.queue <- key
}

// pop the entry and remove it from known keys;
// blocks if queue is empty.
func (q *Queue) Pop() inst.InstanceKey {
	q.Lock()
	queue := q.queue
	q.Unlock()

	key := <-queue

	q.Lock()
	defer q.Unlock()

	// alarm if have been waiting for too long
	timeOnQueue := time.Since(q.knownKeys[key])
	if timeOnQueue > time.Duration(config.Config.InstancePollSeconds)*time.Second {
		log.Warningf("key %v spent %.4fs waiting on a discoveryQueue", key, timeOnQueue.Seconds())
	}

	delete(q.knownKeys, key)
	return key
}
