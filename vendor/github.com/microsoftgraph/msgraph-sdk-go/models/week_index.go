package models
import (
    "errors"
)
// 
type WeekIndex int

const (
    FIRST_WEEKINDEX WeekIndex = iota
    SECOND_WEEKINDEX
    THIRD_WEEKINDEX
    FOURTH_WEEKINDEX
    LAST_WEEKINDEX
)

func (i WeekIndex) String() string {
    return []string{"first", "second", "third", "fourth", "last"}[i]
}
func ParseWeekIndex(v string) (any, error) {
    result := FIRST_WEEKINDEX
    switch v {
        case "first":
            result = FIRST_WEEKINDEX
        case "second":
            result = SECOND_WEEKINDEX
        case "third":
            result = THIRD_WEEKINDEX
        case "fourth":
            result = FOURTH_WEEKINDEX
        case "last":
            result = LAST_WEEKINDEX
        default:
            return 0, errors.New("Unknown WeekIndex value: " + v)
    }
    return &result, nil
}
func SerializeWeekIndex(values []WeekIndex) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
