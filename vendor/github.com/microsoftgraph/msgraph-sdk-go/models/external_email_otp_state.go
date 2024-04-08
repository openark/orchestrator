package models
import (
    "errors"
)
// 
type ExternalEmailOtpState int

const (
    DEFAULTESCAPED_EXTERNALEMAILOTPSTATE ExternalEmailOtpState = iota
    ENABLED_EXTERNALEMAILOTPSTATE
    DISABLED_EXTERNALEMAILOTPSTATE
    UNKNOWNFUTUREVALUE_EXTERNALEMAILOTPSTATE
)

func (i ExternalEmailOtpState) String() string {
    return []string{"default", "enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseExternalEmailOtpState(v string) (any, error) {
    result := DEFAULTESCAPED_EXTERNALEMAILOTPSTATE
    switch v {
        case "default":
            result = DEFAULTESCAPED_EXTERNALEMAILOTPSTATE
        case "enabled":
            result = ENABLED_EXTERNALEMAILOTPSTATE
        case "disabled":
            result = DISABLED_EXTERNALEMAILOTPSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EXTERNALEMAILOTPSTATE
        default:
            return 0, errors.New("Unknown ExternalEmailOtpState value: " + v)
    }
    return &result, nil
}
func SerializeExternalEmailOtpState(values []ExternalEmailOtpState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
