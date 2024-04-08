package models
import (
    "errors"
)
// 
type AppliedConditionalAccessPolicyResult int

const (
    SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT AppliedConditionalAccessPolicyResult = iota
    FAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT
    NOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT
    NOTENABLED_APPLIEDCONDITIONALACCESSPOLICYRESULT
    UNKNOWN_APPLIEDCONDITIONALACCESSPOLICYRESULT
    UNKNOWNFUTUREVALUE_APPLIEDCONDITIONALACCESSPOLICYRESULT
)

func (i AppliedConditionalAccessPolicyResult) String() string {
    return []string{"success", "failure", "notApplied", "notEnabled", "unknown", "unknownFutureValue"}[i]
}
func ParseAppliedConditionalAccessPolicyResult(v string) (any, error) {
    result := SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT
    switch v {
        case "success":
            result = SUCCESS_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "failure":
            result = FAILURE_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "notApplied":
            result = NOTAPPLIED_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "notEnabled":
            result = NOTENABLED_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "unknown":
            result = UNKNOWN_APPLIEDCONDITIONALACCESSPOLICYRESULT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPLIEDCONDITIONALACCESSPOLICYRESULT
        default:
            return 0, errors.New("Unknown AppliedConditionalAccessPolicyResult value: " + v)
    }
    return &result, nil
}
func SerializeAppliedConditionalAccessPolicyResult(values []AppliedConditionalAccessPolicyResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
