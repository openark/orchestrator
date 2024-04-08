package externalconnectors
import (
    "errors"
)
// 
type IdentityType int

const (
    USER_IDENTITYTYPE IdentityType = iota
    GROUP_IDENTITYTYPE
    EXTERNALGROUP_IDENTITYTYPE
    UNKNOWNFUTUREVALUE_IDENTITYTYPE
)

func (i IdentityType) String() string {
    return []string{"user", "group", "externalGroup", "unknownFutureValue"}[i]
}
func ParseIdentityType(v string) (any, error) {
    result := USER_IDENTITYTYPE
    switch v {
        case "user":
            result = USER_IDENTITYTYPE
        case "group":
            result = GROUP_IDENTITYTYPE
        case "externalGroup":
            result = EXTERNALGROUP_IDENTITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_IDENTITYTYPE
        default:
            return 0, errors.New("Unknown IdentityType value: " + v)
    }
    return &result, nil
}
func SerializeIdentityType(values []IdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
