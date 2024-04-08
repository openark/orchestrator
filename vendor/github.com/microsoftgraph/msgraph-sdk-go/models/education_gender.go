package models
import (
    "errors"
)
// 
type EducationGender int

const (
    FEMALE_EDUCATIONGENDER EducationGender = iota
    MALE_EDUCATIONGENDER
    OTHER_EDUCATIONGENDER
    UNKNOWNFUTUREVALUE_EDUCATIONGENDER
)

func (i EducationGender) String() string {
    return []string{"female", "male", "other", "unknownFutureValue"}[i]
}
func ParseEducationGender(v string) (any, error) {
    result := FEMALE_EDUCATIONGENDER
    switch v {
        case "female":
            result = FEMALE_EDUCATIONGENDER
        case "male":
            result = MALE_EDUCATIONGENDER
        case "other":
            result = OTHER_EDUCATIONGENDER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EDUCATIONGENDER
        default:
            return 0, errors.New("Unknown EducationGender value: " + v)
    }
    return &result, nil
}
func SerializeEducationGender(values []EducationGender) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
