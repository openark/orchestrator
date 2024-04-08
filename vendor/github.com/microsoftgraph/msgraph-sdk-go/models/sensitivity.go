package models
import (
    "errors"
)
// 
type Sensitivity int

const (
    NORMAL_SENSITIVITY Sensitivity = iota
    PERSONAL_SENSITIVITY
    PRIVATE_SENSITIVITY
    CONFIDENTIAL_SENSITIVITY
)

func (i Sensitivity) String() string {
    return []string{"normal", "personal", "private", "confidential"}[i]
}
func ParseSensitivity(v string) (any, error) {
    result := NORMAL_SENSITIVITY
    switch v {
        case "normal":
            result = NORMAL_SENSITIVITY
        case "personal":
            result = PERSONAL_SENSITIVITY
        case "private":
            result = PRIVATE_SENSITIVITY
        case "confidential":
            result = CONFIDENTIAL_SENSITIVITY
        default:
            return 0, errors.New("Unknown Sensitivity value: " + v)
    }
    return &result, nil
}
func SerializeSensitivity(values []Sensitivity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
