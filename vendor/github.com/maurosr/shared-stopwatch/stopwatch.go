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

type Stopwatch struct {
	sync.RWMutex

	elapsedTime time.Duration
	startTime   time.Time

	running  map[uint]interface{}
	runCount uint
}

// NewStopwatch creates and returns a new shared stopwatch.
// This stopwatch can be safely shared between different go routines.
// It can be started multiple times, and each time returns a stop callback. It will only stop when all the stop
// callbacks are executed. It measures the elapsed run time, i.e. any time where it had been
// started but the corresponding stop callback hadn't yet been executed.
func NewStopwatch() *Stopwatch {
	s := new(Stopwatch)
	s.running = make(map[uint]interface{})
	s.runCount = 0

	return s
}

// Starts the stopwatch if it was not running. In any case, adds a 'run' reference count.
// Returns a stop callback function that must be executed right after finishing the timed operation. This function
// returns an error if is tried to execute more than once, but won't panic.
func (s *Stopwatch) Start() func() error {
	s.Lock()
	defer s.Unlock()

	wasRunning := s.isRunning()

	s.running[s.runCount] = struct{}{}
	s.runCount++

	if !wasRunning {
		s.startTime = time.Now()
	}

	return stopCallback(s, s.runCount-1)
}

func stopCallback(s *Stopwatch, idx uint) func() error {
	return func() error {
		var elapsedTime time.Duration = time.Since(s.startTime)
		var err error = nil
		s.Lock()
		defer s.Unlock()

		if _, ok := s.running[idx]; ok {
			delete(s.running, idx)

			if !s.isRunning() {
				s.elapsedTime += elapsedTime
				s.startTime = time.Time{}
			}
		} else {
			err = fmt.Errorf("Stop button already pressed!")
		}

		return err
	}
}

func (s *Stopwatch) isRunning() bool {
	return len(s.running) != 0
}

// Checks whether this stopwatch is running, i.e. it was started but not all stop callbacks have been executed.
func (s *Stopwatch) IsRunning() bool {
	s.RLock()
	defer s.RUnlock()

	return s.isRunning()
}

// Returns the elapsed run time. Since it is a shared stopwatch, elapsed run time means all time in which any routine
// that started the stopwatch was running.
func (s *Stopwatch) Elapsed() time.Duration {
	s.RLock()
	defer s.RUnlock()

	elapsed := s.elapsedTime

	if s.isRunning() {
		elapsed += time.Since(s.startTime)
	}

	return elapsed
}

// Resets the stopwatch if it was not running, returns an error otherwise.
func (s *Stopwatch) Reset() error {
	var err error = nil

	s.Lock()
	defer s.Unlock()

	if !s.isRunning() {
		s.elapsedTime = 0
		s.startTime = time.Time{}
		s.runCount = 0
	} else {
		err = fmt.Errorf("Can't reset a stopwatch if it is running")
	}

	return err
}
