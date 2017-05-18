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

// NewNamedStopwatch creates an empty Stopwatch list
func TestNewNamedStopwatch(t *testing.T) {
	n := NewNamedStopwatch()
	if n == nil {
		t.Error("TestNewNamedStopwatch() returns nil")
	}
	len := len(n.Keys())
	if len != 0 {
		t.Error("TestNewNamedStopwatch() len(n) = %d, expecting %d", len, 1)
	}
}

var testNames []string

func init() {
	testNames = []string{"S1", "S2", "S3", "S4"}
}

// uses Add and Exists
func TestAddNamedStopwatch(t *testing.T) {
	n := NewNamedStopwatch()

	for i, v := range testNames {
		n.Add(v)
		len := len(n.Keys())
		if len != (i + 1) {
			t.Errorf("TestAddNamedStopwatch(): len(n.Keys()) = %d after adding %q, expecting %d", len, v, 1+i)
		}
		if !n.Exists(v) {
			t.Errorf("TestAddNamedStopwatch(): %q not exists after adding it.", v)
		}
	}
}

// checkAddMany behaves the same as Add
func TestAddManyNamedStopwatch(t *testing.T) {
	n1 := NewNamedStopwatch()
	n2 := NewNamedStopwatch()

	names := []string{}
	for i := range testNames {
		names = append(names, testNames[i])
	}
	n1.AddMany(names)
	size := len(n1.Keys())
	if size != len(testNames) {
		t.Errorf("TestAddManyNamedStopwatch(): len(n1.Keys()) = %d, expecting %d", size, 1+i)
	}

	for i, v := range testNames {
		if !n2.Exists(v) {
			t.Errorf("TestAddManyNamedStopwatch(): name %q does not exist.", v)
		}
	}
}
