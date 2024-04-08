package models
import (
    "errors"
)
// 
type ServiceHealthClassificationType int

const (
    ADVISORY_SERVICEHEALTHCLASSIFICATIONTYPE ServiceHealthClassificationType = iota
    INCIDENT_SERVICEHEALTHCLASSIFICATIONTYPE
    UNKNOWNFUTUREVALUE_SERVICEHEALTHCLASSIFICATIONTYPE
)

func (i ServiceHealthClassificationType) String() string {
    return []string{"advisory", "incident", "unknownFutureValue"}[i]
}
func ParseServiceHealthClassificationType(v string) (any, error) {
    result := ADVISORY_SERVICEHEALTHCLASSIFICATIONTYPE
    switch v {
        case "advisory":
            result = ADVISORY_SERVICEHEALTHCLASSIFICATIONTYPE
        case "incident":
            result = INCIDENT_SERVICEHEALTHCLASSIFICATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SERVICEHEALTHCLASSIFICATIONTYPE
        default:
            return 0, errors.New("Unknown ServiceHealthClassificationType value: " + v)
    }
    return &result, nil
}
func SerializeServiceHealthClassificationType(values []ServiceHealthClassificationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
