package models
import (
    "errors"
)
// 
type RecurrenceRangeType int

const (
    ENDDATE_RECURRENCERANGETYPE RecurrenceRangeType = iota
    NOEND_RECURRENCERANGETYPE
    NUMBERED_RECURRENCERANGETYPE
)

func (i RecurrenceRangeType) String() string {
    return []string{"endDate", "noEnd", "numbered"}[i]
}
func ParseRecurrenceRangeType(v string) (any, error) {
    result := ENDDATE_RECURRENCERANGETYPE
    switch v {
        case "endDate":
            result = ENDDATE_RECURRENCERANGETYPE
        case "noEnd":
            result = NOEND_RECURRENCERANGETYPE
        case "numbered":
            result = NUMBERED_RECURRENCERANGETYPE
        default:
            return 0, errors.New("Unknown RecurrenceRangeType value: " + v)
    }
    return &result, nil
}
func SerializeRecurrenceRangeType(values []RecurrenceRangeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
