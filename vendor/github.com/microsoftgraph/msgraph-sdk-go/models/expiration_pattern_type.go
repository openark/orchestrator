package models
import (
    "errors"
)
// 
type ExpirationPatternType int

const (
    NOTSPECIFIED_EXPIRATIONPATTERNTYPE ExpirationPatternType = iota
    NOEXPIRATION_EXPIRATIONPATTERNTYPE
    AFTERDATETIME_EXPIRATIONPATTERNTYPE
    AFTERDURATION_EXPIRATIONPATTERNTYPE
)

func (i ExpirationPatternType) String() string {
    return []string{"notSpecified", "noExpiration", "afterDateTime", "afterDuration"}[i]
}
func ParseExpirationPatternType(v string) (any, error) {
    result := NOTSPECIFIED_EXPIRATIONPATTERNTYPE
    switch v {
        case "notSpecified":
            result = NOTSPECIFIED_EXPIRATIONPATTERNTYPE
        case "noExpiration":
            result = NOEXPIRATION_EXPIRATIONPATTERNTYPE
        case "afterDateTime":
            result = AFTERDATETIME_EXPIRATIONPATTERNTYPE
        case "afterDuration":
            result = AFTERDURATION_EXPIRATIONPATTERNTYPE
        default:
            return 0, errors.New("Unknown ExpirationPatternType value: " + v)
    }
    return &result, nil
}
func SerializeExpirationPatternType(values []ExpirationPatternType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
