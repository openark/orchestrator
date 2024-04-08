package security
import (
    "errors"
)
// 
type DefenderAvStatus int

const (
    NOTREPORTING_DEFENDERAVSTATUS DefenderAvStatus = iota
    DISABLED_DEFENDERAVSTATUS
    NOTUPDATED_DEFENDERAVSTATUS
    UPDATED_DEFENDERAVSTATUS
    UNKNOWN_DEFENDERAVSTATUS
    NOTSUPPORTED_DEFENDERAVSTATUS
    UNKNOWNFUTUREVALUE_DEFENDERAVSTATUS
)

func (i DefenderAvStatus) String() string {
    return []string{"notReporting", "disabled", "notUpdated", "updated", "unknown", "notSupported", "unknownFutureValue"}[i]
}
func ParseDefenderAvStatus(v string) (any, error) {
    result := NOTREPORTING_DEFENDERAVSTATUS
    switch v {
        case "notReporting":
            result = NOTREPORTING_DEFENDERAVSTATUS
        case "disabled":
            result = DISABLED_DEFENDERAVSTATUS
        case "notUpdated":
            result = NOTUPDATED_DEFENDERAVSTATUS
        case "updated":
            result = UPDATED_DEFENDERAVSTATUS
        case "unknown":
            result = UNKNOWN_DEFENDERAVSTATUS
        case "notSupported":
            result = NOTSUPPORTED_DEFENDERAVSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEFENDERAVSTATUS
        default:
            return 0, errors.New("Unknown DefenderAvStatus value: " + v)
    }
    return &result, nil
}
func SerializeDefenderAvStatus(values []DefenderAvStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
