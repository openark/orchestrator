package security
import (
    "errors"
)
// 
type DeviceRiskScore int

const (
    NONE_DEVICERISKSCORE DeviceRiskScore = iota
    INFORMATIONAL_DEVICERISKSCORE
    LOW_DEVICERISKSCORE
    MEDIUM_DEVICERISKSCORE
    HIGH_DEVICERISKSCORE
    UNKNOWNFUTUREVALUE_DEVICERISKSCORE
)

func (i DeviceRiskScore) String() string {
    return []string{"none", "informational", "low", "medium", "high", "unknownFutureValue"}[i]
}
func ParseDeviceRiskScore(v string) (any, error) {
    result := NONE_DEVICERISKSCORE
    switch v {
        case "none":
            result = NONE_DEVICERISKSCORE
        case "informational":
            result = INFORMATIONAL_DEVICERISKSCORE
        case "low":
            result = LOW_DEVICERISKSCORE
        case "medium":
            result = MEDIUM_DEVICERISKSCORE
        case "high":
            result = HIGH_DEVICERISKSCORE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICERISKSCORE
        default:
            return 0, errors.New("Unknown DeviceRiskScore value: " + v)
    }
    return &result, nil
}
func SerializeDeviceRiskScore(values []DeviceRiskScore) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
