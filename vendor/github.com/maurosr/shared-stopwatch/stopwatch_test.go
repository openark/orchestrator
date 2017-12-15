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
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	s := NewStopwatch()
	if s == nil {
		t.Errorf("Error: NewStopwatch() returned: %v, expected: non-nil value", s)
	}
}

func TestStart(t *testing.T) {
	s := NewStopwatch()
	s.Start()
	elapsed := s.Elapsed()
	if elapsed <= time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: > 0", elapsed)
	}
}

func TestElapsed(t *testing.T) {
	s := NewStopwatch()
	elapsed0 := s.Elapsed()

	if elapsed0 != time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: 0", elapsed0)
	}

	stopCallback := s.Start()
	elapsed1 := s.Elapsed()

	if elapsed1 <= time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: > 0", elapsed0)
	}

	stopCallback()
	elapsed2 := s.Elapsed()
	elapsed3 := s.Elapsed()

	if elapsed2 < elapsed1 {
		t.Errorf("Error: Elapsed time is not monotonically incresing: elapsed2(%v) <= elapsed1(%v)", elapsed2, elapsed1)
	}

	if elapsed2 != elapsed3 {
		t.Errorf("Error: Elapsed time continues increasing after stop: elapsed2(%v) != elapsed3(%v)", elapsed2, elapsed3)
	}
}

func TestStop(t *testing.T) {
	s := NewStopwatch()
	stopCallback := s.Start()
	err := stopCallback()

	if err != nil {
		t.Error(err)
	}

	elapsed1 := s.Elapsed()
	elapsed2 := s.Elapsed()
	if elapsed1 != elapsed2 {
		t.Errorf("Error: Elapsed time continues increasing after stop: elapsed2(%v) != elapsed3(%v)", elapsed1, elapsed2)
	}

	err = stopCallback()

	if err.Error() != "Stop button already pressed!" {
		t.Errorf("Error: wrong error returned when stopping twice with the same callback: '%v', expected: 'Stop button already pressed!'", err)
	}
}

func TestConcurrency(t *testing.T) {
	s := NewStopwatch()
	checkpoint := make(chan interface{})
	end := make(chan interface{})

	var elapsed1, elapsed2, elapsed3 time.Duration

	go func(tt *testing.T) {
		stopCallback1 := s.Start()
		tt.Log("Routine 1 started stopwatch")

		checkpoint <- struct{}{}
		<-checkpoint

		err1 := stopCallback1()
		tt.Log("Routine 1 stopped stopwatch")

		if err1 != nil {
			tt.Errorf("Error: stopping in routine 1: %v", err1)
		}
		elapsed1 = s.Elapsed()
		tt.Logf("Routine 1 elapsed time: %v", elapsed1)

		checkpoint <- struct{}{}
	}(t)

	go func(tt *testing.T) {
		<-checkpoint

		stopCallback2 := s.Start()
		tt.Log("Routine 2 started stopwatch")

		checkpoint <- struct{}{}
		<-checkpoint

		err2 := stopCallback2()
		tt.Log("Routine 2 stopped stopwatch")

		if err2 != nil {
			tt.Errorf("Error: stopping in routine 2: %v", err2)
		}
		elapsed2 = s.Elapsed()
		tt.Logf("Routine 2 elapsed time: %v", elapsed2)

		if elapsed2 <= elapsed1 {
			tt.Errorf("Error: Stopwatch stopped before all routines stop")
		}

		end <- struct{}{}
	}(t)

	<-end

	elapsed3 = s.Elapsed()
	t.Logf("Final elapsed time: %v", elapsed3)
	if elapsed2 != elapsed3 {
		t.Errorf("Error: Stopwatch did not stop after all routines stop")
	}
}

func TestReset(t *testing.T) {
	s := NewStopwatch()
	stopC := s.Start()

	err := s.Reset()
	if err.Error() != "Can't reset a stopwatch if it is running" {
		t.Errorf("Error: wrong error returned when resetting a running stopwatch: '%v', expected: 'Can't reset a stopwatch if it is running'", err)
	}

	stopC()
	err = s.Reset()
	if err != nil {
		t.Errorf("Error: reset failed: %v", err)
	}
	elapsed := s.Elapsed()
	if elapsed != time.Duration(0) {
		t.Errorf("Error: elapsed time after reset %v, expected 0", elapsed)
	}
}

func TestIsRunning(t *testing.T) {
	s := NewStopwatch()
	if s.IsRunning() {
		t.Errorf("Error: Stopwatch running before started")
	}
	stopC := s.Start()
	if !s.IsRunning() {
		t.Errorf("Error: Stopwatch not running after started")
	}

	stopC2 := s.Start()

	stopC()
	if !s.IsRunning() {
		t.Errorf("Error: Stopwatch stopped before all stop callbacks executed")
	}
	stopC2()
	if s.IsRunning() {
		t.Errorf("Error: Stopwatch did not stop after all stop callbacks executed")
	}
}
