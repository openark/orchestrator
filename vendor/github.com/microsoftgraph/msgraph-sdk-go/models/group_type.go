package models
import (
    "errors"
)
// 
type GroupType int

const (
    UNIFIEDGROUPS_GROUPTYPE GroupType = iota
    AZUREAD_GROUPTYPE
    UNKNOWNFUTUREVALUE_GROUPTYPE
)

func (i GroupType) String() string {
    return []string{"unifiedGroups", "azureAD", "unknownFutureValue"}[i]
}
func ParseGroupType(v string) (any, error) {
    result := UNIFIEDGROUPS_GROUPTYPE
    switch v {
        case "unifiedGroups":
            result = UNIFIEDGROUPS_GROUPTYPE
        case "azureAD":
            result = AZUREAD_GROUPTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_GROUPTYPE
        default:
            return 0, errors.New("Unknown GroupType value: " + v)
    }
    return &result, nil
}
func SerializeGroupType(values []GroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
