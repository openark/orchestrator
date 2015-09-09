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
	"encoding/json"
	"strings"
)

// InstanceKeyMap is a convenience struct for listing InstanceKey-s
type InstanceKeyMap map[InstanceKey]bool

func NewInstanceKeyMap() *InstanceKeyMap {
	return &InstanceKeyMap{}
}

// GetInstanceKeys returns keys in this map in the form of an array
func (this *InstanceKeyMap) AddKey(key InstanceKey) {
	(*this)[key] = true
}

// GetInstanceKeys returns keys in this map in the form of an array
func (this *InstanceKeyMap) AddKeys(keys []InstanceKey) {
	for _, key := range keys {
		this.AddKey(key)
	}
}

// GetInstanceKeys returns keys in this map in the form of an array
func (this *InstanceKeyMap) GetInstanceKeys() []InstanceKey {
	res := []InstanceKey{}
	for key := range *this {
		res = append(res, key)
	}
	return res
}

// MarshalJSON will marshal this map as JSON
func (this *InstanceKeyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.GetInstanceKeys())
}

// MarshalJSON will marshal this map as JSON
func (this *InstanceKeyMap) ToJSON() (string, error) {
	bytes, err := this.MarshalJSON()
	return string(bytes), err
}

// MarshalJSON will marshal this map as JSON
func (this *InstanceKeyMap) ToJSONString() string {
	s, _ := this.ToJSON()
	return s
}

// ReadJson unmarshalls a json into this map
func (this *InstanceKeyMap) ReadJson(jsonString string) error {
	var keys []InstanceKey
	err := json.Unmarshal([]byte(jsonString), &keys)
	if err != nil {
		return err
	}
	this.AddKeys(keys)
	return err
}

// ReadJson unmarshalls a json into this map
func (this *InstanceKeyMap) ReadCommaDelimitedList(list string) error {
	tokens := strings.Split(list, ",")
	for _, token := range tokens {
		key, err := ParseInstanceKey(token)
		if err != nil {
			return err
		}
		this.AddKey(*key)
	}
	return nil
}
