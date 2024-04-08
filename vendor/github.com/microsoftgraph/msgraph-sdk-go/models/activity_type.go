package models
import (
    "errors"
)
// 
type ActivityType int

const (
    SIGNIN_ACTIVITYTYPE ActivityType = iota
    USER_ACTIVITYTYPE
    UNKNOWNFUTUREVALUE_ACTIVITYTYPE
    SERVICEPRINCIPAL_ACTIVITYTYPE
)

func (i ActivityType) String() string {
    return []string{"signin", "user", "unknownFutureValue", "servicePrincipal"}[i]
}
func ParseActivityType(v string) (any, error) {
    result := SIGNIN_ACTIVITYTYPE
    switch v {
        case "signin":
            result = SIGNIN_ACTIVITYTYPE
        case "user":
            result = USER_ACTIVITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACTIVITYTYPE
        case "servicePrincipal":
            result = SERVICEPRINCIPAL_ACTIVITYTYPE
        default:
            return 0, errors.New("Unknown ActivityType value: " + v)
    }
    return &result, nil
}
func SerializeActivityType(values []ActivityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
