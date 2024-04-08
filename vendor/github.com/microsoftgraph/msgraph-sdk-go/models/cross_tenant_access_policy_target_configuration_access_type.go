package models
import (
    "errors"
)
// 
type CrossTenantAccessPolicyTargetConfigurationAccessType int

const (
    ALLOWED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE CrossTenantAccessPolicyTargetConfigurationAccessType = iota
    BLOCKED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
    UNKNOWNFUTUREVALUE_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
)

func (i CrossTenantAccessPolicyTargetConfigurationAccessType) String() string {
    return []string{"allowed", "blocked", "unknownFutureValue"}[i]
}
func ParseCrossTenantAccessPolicyTargetConfigurationAccessType(v string) (any, error) {
    result := ALLOWED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
    switch v {
        case "allowed":
            result = ALLOWED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
        case "blocked":
            result = BLOCKED_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CROSSTENANTACCESSPOLICYTARGETCONFIGURATIONACCESSTYPE
        default:
            return 0, errors.New("Unknown CrossTenantAccessPolicyTargetConfigurationAccessType value: " + v)
    }
    return &result, nil
}
func SerializeCrossTenantAccessPolicyTargetConfigurationAccessType(values []CrossTenantAccessPolicyTargetConfigurationAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
