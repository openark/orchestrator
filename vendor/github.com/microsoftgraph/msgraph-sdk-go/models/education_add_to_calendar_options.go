package models
import (
    "errors"
)
// 
type EducationAddToCalendarOptions int

const (
    NONE_EDUCATIONADDTOCALENDAROPTIONS EducationAddToCalendarOptions = iota
    STUDENTSANDPUBLISHER_EDUCATIONADDTOCALENDAROPTIONS
    STUDENTSANDTEAMOWNERS_EDUCATIONADDTOCALENDAROPTIONS
    UNKNOWNFUTUREVALUE_EDUCATIONADDTOCALENDAROPTIONS
    STUDENTSONLY_EDUCATIONADDTOCALENDAROPTIONS
)

func (i EducationAddToCalendarOptions) String() string {
    return []string{"none", "studentsAndPublisher", "studentsAndTeamOwners", "unknownFutureValue", "studentsOnly"}[i]
}
func ParseEducationAddToCalendarOptions(v string) (any, error) {
    result := NONE_EDUCATIONADDTOCALENDAROPTIONS
    switch v {
        case "none":
            result = NONE_EDUCATIONADDTOCALENDAROPTIONS
        case "studentsAndPublisher":
            result = STUDENTSANDPUBLISHER_EDUCATIONADDTOCALENDAROPTIONS
        case "studentsAndTeamOwners":
            result = STUDENTSANDTEAMOWNERS_EDUCATIONADDTOCALENDAROPTIONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EDUCATIONADDTOCALENDAROPTIONS
        case "studentsOnly":
            result = STUDENTSONLY_EDUCATIONADDTOCALENDAROPTIONS
        default:
            return 0, errors.New("Unknown EducationAddToCalendarOptions value: " + v)
    }
    return &result, nil
}
func SerializeEducationAddToCalendarOptions(values []EducationAddToCalendarOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
