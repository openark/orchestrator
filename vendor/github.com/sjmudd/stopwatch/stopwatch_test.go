/*
Copyright (c) 2016, Simon J Mudd
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

// Package stopwatch implements simple stopwatch functionality
package stopwatch

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New(nil)
	if s == nil {
		t.Error("TestNew(): New() returned %v, expecting non-nil value", s)
	}
}

func TestStart(t *testing.T) {
	s := Start(nil)
	elapsed := s.ElapsedSeconds() // convert to seconds should be ok for now
	if elapsed == 0 {
		t.Error("TestStart(): s.Elapsed() returned %+v, expecting > 0", elapsed)
	}
}

func TestStop(t *testing.T) {
	s := Start(nil)
	s.Stop()
	elapsed1 := s.ElapsedSeconds() // convert to seconds should be ok for now
	elapsed2 := s.ElapsedSeconds() // convert to seconds should be ok for now
	if elapsed1 != elapsed2 {
		t.Error("TestStop(): elapsed1 (%v) != elapsed2 (%v). Expecting them to be the same", elapsed1, elapsed2)
	}
}

func TestReset(t *testing.T) {
	s := New(nil)
	s.Start()
}

func TestIsRunning(t *testing.T) {
	s := New(nil)
	if s.IsRunning() {
		t.Error("TestIsRunning() returns %v after New(), expecting false", s.IsRunning())
	}
	s.Start()
	if !s.IsRunning() {
		t.Error("TestIsRunning() returns %v after Start(), expecting true", s.IsRunning())
	}
	s.Start()
	if !s.IsRunning() {
		t.Error("TestIsRunning() returns %v after 2nd Start(), expecting true", s.IsRunning())
	}
	s.Stop()
	if s.IsRunning() {
		t.Error("TestIsRunning() returns %v after 2nd Stop(), expecting false", s.IsRunning())
	}
}
