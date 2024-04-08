package models
import (
    "errors"
)
// Indicates the type of execution context the app runs in.
type RunAsAccountType int

const (
    // System context
    SYSTEM_RUNASACCOUNTTYPE RunAsAccountType = iota
    // User context
    USER_RUNASACCOUNTTYPE
)

func (i RunAsAccountType) String() string {
    return []string{"system", "user"}[i]
}
func ParseRunAsAccountType(v string) (any, error) {
    result := SYSTEM_RUNASACCOUNTTYPE
    switch v {
        case "system":
            result = SYSTEM_RUNASACCOUNTTYPE
        case "user":
            result = USER_RUNASACCOUNTTYPE
        default:
            return 0, errors.New("Unknown RunAsAccountType value: " + v)
    }
    return &result, nil
}
func SerializeRunAsAccountType(values []RunAsAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
