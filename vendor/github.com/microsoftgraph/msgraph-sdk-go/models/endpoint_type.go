package models
import (
    "errors"
)
// 
type EndpointType int

const (
    DEFAULTESCAPED_ENDPOINTTYPE EndpointType = iota
    VOICEMAIL_ENDPOINTTYPE
    SKYPEFORBUSINESS_ENDPOINTTYPE
    SKYPEFORBUSINESSVOIPPHONE_ENDPOINTTYPE
    UNKNOWNFUTUREVALUE_ENDPOINTTYPE
)

func (i EndpointType) String() string {
    return []string{"default", "voicemail", "skypeForBusiness", "skypeForBusinessVoipPhone", "unknownFutureValue"}[i]
}
func ParseEndpointType(v string) (any, error) {
    result := DEFAULTESCAPED_ENDPOINTTYPE
    switch v {
        case "default":
            result = DEFAULTESCAPED_ENDPOINTTYPE
        case "voicemail":
            result = VOICEMAIL_ENDPOINTTYPE
        case "skypeForBusiness":
            result = SKYPEFORBUSINESS_ENDPOINTTYPE
        case "skypeForBusinessVoipPhone":
            result = SKYPEFORBUSINESSVOIPPHONE_ENDPOINTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENDPOINTTYPE
        default:
            return 0, errors.New("Unknown EndpointType value: " + v)
    }
    return &result, nil
}
func SerializeEndpointType(values []EndpointType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
