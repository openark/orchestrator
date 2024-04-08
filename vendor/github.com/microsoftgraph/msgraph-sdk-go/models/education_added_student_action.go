package models
import (
    "errors"
)
// 
type EducationAddedStudentAction int

const (
    NONE_EDUCATIONADDEDSTUDENTACTION EducationAddedStudentAction = iota
    ASSIGNIFOPEN_EDUCATIONADDEDSTUDENTACTION
    UNKNOWNFUTUREVALUE_EDUCATIONADDEDSTUDENTACTION
)

func (i EducationAddedStudentAction) String() string {
    return []string{"none", "assignIfOpen", "unknownFutureValue"}[i]
}
func ParseEducationAddedStudentAction(v string) (any, error) {
    result := NONE_EDUCATIONADDEDSTUDENTACTION
    switch v {
        case "none":
            result = NONE_EDUCATIONADDEDSTUDENTACTION
        case "assignIfOpen":
            result = ASSIGNIFOPEN_EDUCATIONADDEDSTUDENTACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EDUCATIONADDEDSTUDENTACTION
        default:
            return 0, errors.New("Unknown EducationAddedStudentAction value: " + v)
    }
    return &result, nil
}
func SerializeEducationAddedStudentAction(values []EducationAddedStudentAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
