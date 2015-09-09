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

package inst

import (
	"fmt"
	"github.com/outbrain/orchestrator/go/config"
	"strconv"
	"strings"
)

// InstanceKey is an instance indicator, identifued by hostname and port
type InstanceKey struct {
	Hostname string
	Port     int
}

// ParseInstanceKey will parse an InstanceKey from a string representation such as 127.0.0.1:3306
func NewRawInstanceKey(hostPort string) (*InstanceKey, error) {
	tokens := strings.SplitN(hostPort, ":", 2)
	if len(tokens) != 2 {
		return nil, fmt.Errorf("Cannot parse InstanceKey from %s. Expected format is host:port", hostPort)
	}
	instanceKey := &InstanceKey{Hostname: tokens[0]}
	var err error
	if instanceKey.Port, err = strconv.Atoi(tokens[1]); err != nil {
		return instanceKey, fmt.Errorf("Invalid port: %s", tokens[1])
	}

	return instanceKey, nil
}

// NewInstanceKeyFromStrings creates a new InstanceKey by resolving hostname and port.
// hostname is normalized via ResolveHostname. port is tested to be valid integer.
func NewInstanceKeyFromStrings(hostname string, port string) (*InstanceKey, error) {
	instanceKey := &InstanceKey{}
	var err error
	if instanceKey.Port, err = strconv.Atoi(port); err != nil {
		return instanceKey, fmt.Errorf("Invalid port: %s", port)
	}

	if instanceKey.Hostname, err = ResolveHostname(hostname); err != nil {
		return instanceKey, err
	}

	return instanceKey, nil
}

// ParseInstanceKey will parse an InstanceKey from a string representation such as 127.0.0.1:3306
func ParseInstanceKey(hostPort string) (*InstanceKey, error) {
	tokens := strings.SplitN(hostPort, ":", 2)
	if len(tokens) != 2 {
		return nil, fmt.Errorf("Cannot parse InstanceKey from %s. Expected format is host:port", hostPort)
	}
	return NewInstanceKeyFromStrings(tokens[0], tokens[1])
}

// ParseInstanceKeyLoose will parse an InstanceKey from a string representation such as 127.0.0.1:3306.
// The port part is optional
func ParseInstanceKeyLoose(hostPort string) (*InstanceKey, error) {
	if !strings.Contains(hostPort, ":") {
		return &InstanceKey{Hostname: hostPort, Port: config.Config.DefaultInstancePort}, nil
	}
	return ParseInstanceKey(hostPort)
}

// Formalize this key by getting CNAME for hostname
func (this *InstanceKey) Formalize() *InstanceKey {
	this.Hostname, _ = ResolveHostname(this.Hostname)
	return this
}

// Equals tests equality between this key and another key
func (this *InstanceKey) Equals(other *InstanceKey) bool {
	return this.Hostname == other.Hostname && this.Port == other.Port
}

// IsValid uses simple heuristics to see whether this key represents an actual instance
func (this *InstanceKey) IsValid() bool {
	return len(this.Hostname) > 0 && this.Port > 0
}

// DisplayString returns a user-friendly string representation of this key
func (this *InstanceKey) DisplayString() string {
	return fmt.Sprintf("%s:%d", this.Hostname, this.Port)
}
