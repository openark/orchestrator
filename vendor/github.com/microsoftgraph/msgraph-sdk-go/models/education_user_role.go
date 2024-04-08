package models
import (
    "errors"
)
// 
type EducationUserRole int

const (
    STUDENT_EDUCATIONUSERROLE EducationUserRole = iota
    TEACHER_EDUCATIONUSERROLE
    NONE_EDUCATIONUSERROLE
    UNKNOWNFUTUREVALUE_EDUCATIONUSERROLE
)

func (i EducationUserRole) String() string {
    return []string{"student", "teacher", "none", "unknownFutureValue"}[i]
}
func ParseEducationUserRole(v string) (any, error) {
    result := STUDENT_EDUCATIONUSERROLE
    switch v {
        case "student":
            result = STUDENT_EDUCATIONUSERROLE
        case "teacher":
            result = TEACHER_EDUCATIONUSERROLE
        case "none":
            result = NONE_EDUCATIONUSERROLE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EDUCATIONUSERROLE
        default:
            return 0, errors.New("Unknown EducationUserRole value: " + v)
    }
    return &result, nil
}
func SerializeEducationUserRole(values []EducationUserRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
