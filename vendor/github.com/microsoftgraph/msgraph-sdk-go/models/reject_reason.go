package models
import (
    "errors"
)
// 
type RejectReason int

const (
    NONE_REJECTREASON RejectReason = iota
    BUSY_REJECTREASON
    FORBIDDEN_REJECTREASON
    UNKNOWNFUTUREVALUE_REJECTREASON
)

func (i RejectReason) String() string {
    return []string{"none", "busy", "forbidden", "unknownFutureValue"}[i]
}
func ParseRejectReason(v string) (any, error) {
    result := NONE_REJECTREASON
    switch v {
        case "none":
            result = NONE_REJECTREASON
        case "busy":
            result = BUSY_REJECTREASON
        case "forbidden":
            result = FORBIDDEN_REJECTREASON
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REJECTREASON
        default:
            return 0, errors.New("Unknown RejectReason value: " + v)
    }
    return &result, nil
}
func SerializeRejectReason(values []RejectReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
