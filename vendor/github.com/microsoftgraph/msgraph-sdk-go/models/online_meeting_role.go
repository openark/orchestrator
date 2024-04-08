package models
import (
    "errors"
)
// 
type OnlineMeetingRole int

const (
    ATTENDEE_ONLINEMEETINGROLE OnlineMeetingRole = iota
    PRESENTER_ONLINEMEETINGROLE
    UNKNOWNFUTUREVALUE_ONLINEMEETINGROLE
    PRODUCER_ONLINEMEETINGROLE
    COORGANIZER_ONLINEMEETINGROLE
)

func (i OnlineMeetingRole) String() string {
    return []string{"attendee", "presenter", "unknownFutureValue", "producer", "coorganizer"}[i]
}
func ParseOnlineMeetingRole(v string) (any, error) {
    result := ATTENDEE_ONLINEMEETINGROLE
    switch v {
        case "attendee":
            result = ATTENDEE_ONLINEMEETINGROLE
        case "presenter":
            result = PRESENTER_ONLINEMEETINGROLE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ONLINEMEETINGROLE
        case "producer":
            result = PRODUCER_ONLINEMEETINGROLE
        case "coorganizer":
            result = COORGANIZER_ONLINEMEETINGROLE
        default:
            return 0, errors.New("Unknown OnlineMeetingRole value: " + v)
    }
    return &result, nil
}
func SerializeOnlineMeetingRole(values []OnlineMeetingRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
