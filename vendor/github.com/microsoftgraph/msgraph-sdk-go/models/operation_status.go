package models
import (
    "errors"
)
// 
type OperationStatus int

const (
    NOTSTARTED_OPERATIONSTATUS OperationStatus = iota
    RUNNING_OPERATIONSTATUS
    COMPLETED_OPERATIONSTATUS
    FAILED_OPERATIONSTATUS
)

func (i OperationStatus) String() string {
    return []string{"NotStarted", "Running", "Completed", "Failed"}[i]
}
func ParseOperationStatus(v string) (any, error) {
    result := NOTSTARTED_OPERATIONSTATUS
    switch v {
        case "NotStarted":
            result = NOTSTARTED_OPERATIONSTATUS
        case "Running":
            result = RUNNING_OPERATIONSTATUS
        case "Completed":
            result = COMPLETED_OPERATIONSTATUS
        case "Failed":
            result = FAILED_OPERATIONSTATUS
        default:
            return 0, errors.New("Unknown OperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeOperationStatus(values []OperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
