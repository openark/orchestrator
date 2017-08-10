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

package inst

import (
	"fmt"
)

// CandidatePromotionRule describe the promotion preference/rule for an instance.
// It maps to promotion_rule column in candidate_database_instance
type CandidatePromotionRule string

const (
	MustPromoteRule      CandidatePromotionRule = "must"
	PreferPromoteRule                           = "prefer"
	NeutralPromoteRule                          = "neutral"
	PreferNotPromoteRule                        = "prefer_not"
	MustNotPromoteRule                          = "must_not"
)

// ParseCandidatePromotionRule returns a CandidatePromotionRule by name.
// It returns an error if there is no known rule by the given name.
func ParseCandidatePromotionRule(ruleName string) (CandidatePromotionRule, error) {
	switch ruleName {
	case "prefer", "neutral", "prefer_not", "must_not":
		return CandidatePromotionRule(ruleName), nil
	case "must":
		return CandidatePromotionRule(""), fmt.Errorf("CandidatePromotionRule: %v not supported yet", ruleName)
	default:
		return CandidatePromotionRule(""), fmt.Errorf("Invalid CandidatePromotionRule: %v", ruleName)
	}
}
