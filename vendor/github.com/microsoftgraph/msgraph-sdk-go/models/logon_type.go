package models
import (
    "errors"
)
// 
type LogonType int

const (
    UNKNOWN_LOGONTYPE LogonType = iota
    INTERACTIVE_LOGONTYPE
    REMOTEINTERACTIVE_LOGONTYPE
    NETWORK_LOGONTYPE
    BATCH_LOGONTYPE
    SERVICE_LOGONTYPE
    UNKNOWNFUTUREVALUE_LOGONTYPE
)

func (i LogonType) String() string {
    return []string{"unknown", "interactive", "remoteInteractive", "network", "batch", "service", "unknownFutureValue"}[i]
}
func ParseLogonType(v string) (any, error) {
    result := UNKNOWN_LOGONTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_LOGONTYPE
        case "interactive":
            result = INTERACTIVE_LOGONTYPE
        case "remoteInteractive":
            result = REMOTEINTERACTIVE_LOGONTYPE
        case "network":
            result = NETWORK_LOGONTYPE
        case "batch":
            result = BATCH_LOGONTYPE
        case "service":
            result = SERVICE_LOGONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LOGONTYPE
        default:
            return 0, errors.New("Unknown LogonType value: " + v)
    }
    return &result, nil
}
func SerializeLogonType(values []LogonType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
