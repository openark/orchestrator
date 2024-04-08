package models
import (
    "errors"
)
// 
type UserAccountSecurityType int

const (
    UNKNOWN_USERACCOUNTSECURITYTYPE UserAccountSecurityType = iota
    STANDARD_USERACCOUNTSECURITYTYPE
    POWER_USERACCOUNTSECURITYTYPE
    ADMINISTRATOR_USERACCOUNTSECURITYTYPE
    UNKNOWNFUTUREVALUE_USERACCOUNTSECURITYTYPE
)

func (i UserAccountSecurityType) String() string {
    return []string{"unknown", "standard", "power", "administrator", "unknownFutureValue"}[i]
}
func ParseUserAccountSecurityType(v string) (any, error) {
    result := UNKNOWN_USERACCOUNTSECURITYTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_USERACCOUNTSECURITYTYPE
        case "standard":
            result = STANDARD_USERACCOUNTSECURITYTYPE
        case "power":
            result = POWER_USERACCOUNTSECURITYTYPE
        case "administrator":
            result = ADMINISTRATOR_USERACCOUNTSECURITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USERACCOUNTSECURITYTYPE
        default:
            return 0, errors.New("Unknown UserAccountSecurityType value: " + v)
    }
    return &result, nil
}
func SerializeUserAccountSecurityType(values []UserAccountSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
