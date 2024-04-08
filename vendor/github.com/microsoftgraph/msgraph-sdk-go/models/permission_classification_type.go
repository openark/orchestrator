package models
import (
    "errors"
)
// 
type PermissionClassificationType int

const (
    LOW_PERMISSIONCLASSIFICATIONTYPE PermissionClassificationType = iota
    MEDIUM_PERMISSIONCLASSIFICATIONTYPE
    HIGH_PERMISSIONCLASSIFICATIONTYPE
    UNKNOWNFUTUREVALUE_PERMISSIONCLASSIFICATIONTYPE
)

func (i PermissionClassificationType) String() string {
    return []string{"low", "medium", "high", "unknownFutureValue"}[i]
}
func ParsePermissionClassificationType(v string) (any, error) {
    result := LOW_PERMISSIONCLASSIFICATIONTYPE
    switch v {
        case "low":
            result = LOW_PERMISSIONCLASSIFICATIONTYPE
        case "medium":
            result = MEDIUM_PERMISSIONCLASSIFICATIONTYPE
        case "high":
            result = HIGH_PERMISSIONCLASSIFICATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PERMISSIONCLASSIFICATIONTYPE
        default:
            return 0, errors.New("Unknown PermissionClassificationType value: " + v)
    }
    return &result, nil
}
func SerializePermissionClassificationType(values []PermissionClassificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
