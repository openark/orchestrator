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

func TestNewNamed(t *testing.T) {
	s := NewNamedStopwatch()
	if s == nil {
		t.Errorf("Error: NewStopwatch() returned: %v, expected: non-nil value", s)
	}
}

func TestAddAndDeleteNamed(t *testing.T) {
	s := NewNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	for _, name := range names {
		if !s.Exists(name) {
			t.Errorf("Stopwatch %v was not added", name)
		}
	}

	s.Delete(names[0])
	if s.Exists(names[0]) {
		t.Errorf("Stopwatch %v was not successfully deleted", names[0])
	}
}

func TestStartNamed(t *testing.T) {
	s := NewNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	s.Start(names[0])
	elapsed := s.Elapsed(names[0])
	if elapsed <= time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: > 0", elapsed)
	}
}

func TestElapsedNamed(t *testing.T) {
	s := NewNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	elapsed0 := s.Elapsed(names[0])

	if elapsed0 != time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: 0", elapsed0)
	}

	stopCallback := s.Start(names[0])
	elapsed1 := s.Elapsed(names[0])

	if elapsed1 <= time.Duration(0) {
		t.Errorf("Error: Elapsed time returned: %v, expected: > 0", elapsed0)
	}

	stopCallback()
	elapsed2 := s.Elapsed(names[0])
	elapsed3 := s.Elapsed(names[0])

	if elapsed2 < elapsed1 {
		t.Errorf("Error: Elapsed time is not monotonically incresing: elapsed2(%v) <= elapsed1(%v)", elapsed2, elapsed1)
	}

	if elapsed2 != elapsed3 {
		t.Errorf("Error: Elapsed time continues increasing after stop: elapsed2(%v) != elapsed3(%v)", elapsed2, elapsed3)
	}
}

func TestStopNamed(t *testing.T) {
	s := NewNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	stopCallback := s.Start(names[0])
	err := stopCallback()

	if err != nil {
		t.Error(err)
	}

	elapsed1 := s.Elapsed(names[0])
	elapsed2 := s.Elapsed(names[0])
	if elapsed1 != elapsed2 {
		t.Errorf("Error: Elapsed time continues increasing after stop: elapsed2(%v) != elapsed3(%v)", elapsed1, elapsed2)
	}

	err = stopCallback()

	if err.Error() != "Stop button already pressed!" {
		t.Errorf("Error: wrong error returned when stopping twice with the same callback: '%v', expected: 'Stop button already pressed!'", err)
	}
}

func TestConcurrencyNamed(t *testing.T) {
	s := NewNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	checkpoint := make(chan interface{})
	end := make(chan interface{})

	var elapsed1, elapsed2, elapsed3 time.Duration

	go func(tt *testing.T) {
		stopCallback1 := s.Start(names[0])
		tt.Log("Routine 1 started stopwatch")

		checkpoint <- struct{}{}
		<-checkpoint

		err1 := stopCallback1()
		tt.Log("Routine 1 stopped stopwatch")

		if err1 != nil {
			tt.Errorf("Error: stopping in routine 1: %v", err1)
		}
		elapsed1 = s.Elapsed(names[0])
		tt.Logf("Routine 1 elapsed time: %v", elapsed1)

		checkpoint <- struct{}{}
	}(t)

	go func(tt *testing.T) {
		<-checkpoint

		stopCallback2 := s.Start(names[0])
		tt.Log("Routine 2 started stopwatch")

		checkpoint <- struct{}{}
		<-checkpoint

		err2 := stopCallback2()
		tt.Log("Routine 2 stopped stopwatch")

		if err2 != nil {
			tt.Errorf("Error: stopping in routine 2: %v", err2)
		}
		elapsed2 = s.Elapsed(names[0])
		tt.Logf("Routine 2 elapsed time: %v", elapsed2)

		if elapsed2 <= elapsed1 {
			tt.Errorf("Error: Stopwatch stopped before all routines stop")
		}

		end <- struct{}{}
	}(t)

	<-end

	elapsed3 = s.Elapsed(names[0])
	t.Logf("Final elapsed time: %v", elapsed3)
	if elapsed2 != elapsed3 {
		t.Errorf("Error: Stopwatch did not stop after all routines stop")
	}
}

func TestResetNamed(t *testing.T) {
	s := NewNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	stopC := s.Start(names[0])

	err := s.Reset(names[0])
	if err.Error() != "Can't reset a stopwatch if it is running" {
		t.Errorf("Error: wrong error returned when resetting a running stopwatch: '%v', expected: 'Can't reset a stopwatch if it is running'", err)
	}

	stopC()
	err = s.Reset(names[0])
	if err != nil {
		t.Errorf("Error: reset failed: %v", err)
	}
	elapsed := s.Elapsed(names[0])
	if elapsed != time.Duration(0) {
		t.Errorf("Error: elapsed time after reset %v, expected 0", elapsed)
	}
}

func TestIsRunningNamed(t *testing.T) {
	s := NewNamedStopwatch()
	names := []string{"s1", "s2", "s3"}

	s.Add(names...)

	if s.IsRunning(names[0]) {
		t.Errorf("Error: Stopwatch running before started")
	}
	stopC := s.Start(names[0])
	if !s.IsRunning(names[0]) {
		t.Errorf("Error: Stopwatch not running after started")
	}

	stopC2 := s.Start(names[0])

	stopC()
	if !s.IsRunning(names[0]) {
		t.Errorf("Error: Stopwatch stopped before all stop callbacks executed")
	}
	stopC2()
	if s.IsRunning(names[0]) {
		t.Errorf("Error: Stopwatch did not stop after all stop callbacks executed")
	}
}
