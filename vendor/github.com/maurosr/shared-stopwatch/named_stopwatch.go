/*
Copyright (c) 2017, Mauro Schilman, courtesy Booking.com
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Package stopwatch implements a shared stopwatch
package stopwatch

import (
	"fmt"
	"sync"
	"time"
)

type NamedStopwatch struct {
	sync.RWMutex
	stopwatches map[string]*Stopwatch
}

func NewNamedStopwatch() *NamedStopwatch {
	ns := new(NamedStopwatch)
	ns.stopwatches = make(map[string]*Stopwatch)

	return ns
}

func (ns *NamedStopwatch) Add(names ...string) {
	ns.Lock()
	defer ns.Unlock()

	for _, name := range names {
		if _, ok := ns.stopwatches[name]; !ok {
			ns.stopwatches[name] = NewStopwatch()
		}
	}
}

func (ns *NamedStopwatch) Delete(name string) {
	ns.Lock()
	defer ns.Unlock()

	delete(ns.stopwatches, name)
}

func (ns *NamedStopwatch) Exists(name string) bool {
	ns.RLock()
	defer ns.RUnlock()

	_, found := ns.stopwatches[name]

	return found
}

func (ns *NamedStopwatch) IsRunning(name string) bool {
	ns.RLock()
	defer ns.RUnlock()

	s, ok := ns.stopwatches[name]

	return ok && s.IsRunning()
}

func (ns *NamedStopwatch) Keys() []string {
	ns.RLock()
	defer ns.RUnlock()

	keys := []string{}
	for k := range ns.stopwatches {
		keys = append(keys, k)
	}
	return keys
}

func (ns *NamedStopwatch) Start(name string) func() error {
	ns.Lock()
	defer ns.Unlock()

	if s, ok := ns.stopwatches[name]; ok {
		return s.Start()
	} else {
		return nil
	}
}

func (ns *NamedStopwatch) Elapsed(name string) time.Duration {
	ns.RLock()
	defer ns.RUnlock()

	if s, ok := ns.stopwatches[name]; ok {
		return s.Elapsed()
	} else {
		return time.Duration(0)
	}
}

func (ns *NamedStopwatch) Reset(name string) error {
	ns.Lock()
	defer ns.Unlock()

	if s, ok := ns.stopwatches[name]; ok {
		return s.Reset()
	} else {
		return fmt.Errorf("No stopwatch with name %v", name)
	}
}
