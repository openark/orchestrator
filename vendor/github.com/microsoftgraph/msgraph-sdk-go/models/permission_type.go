package models
import (
    "errors"
)
// 
type PermissionType int

const (
    APPLICATION_PERMISSIONTYPE PermissionType = iota
    DELEGATED_PERMISSIONTYPE
    DELEGATEDUSERCONSENTABLE_PERMISSIONTYPE
)

func (i PermissionType) String() string {
    return []string{"application", "delegated", "delegatedUserConsentable"}[i]
}
func ParsePermissionType(v string) (any, error) {
    result := APPLICATION_PERMISSIONTYPE
    switch v {
        case "application":
            result = APPLICATION_PERMISSIONTYPE
        case "delegated":
            result = DELEGATED_PERMISSIONTYPE
        case "delegatedUserConsentable":
            result = DELEGATEDUSERCONSENTABLE_PERMISSIONTYPE
        default:
            return 0, errors.New("Unknown PermissionType value: " + v)
    }
    return &result, nil
}
func SerializePermissionType(values []PermissionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
