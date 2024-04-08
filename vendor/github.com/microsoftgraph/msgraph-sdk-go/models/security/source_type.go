package security
import (
    "errors"
)
// 
type SourceType int

const (
    MAILBOX_SOURCETYPE SourceType = iota
    SITE_SOURCETYPE
    UNKNOWNFUTUREVALUE_SOURCETYPE
)

func (i SourceType) String() string {
    return []string{"mailbox", "site", "unknownFutureValue"}[i]
}
func ParseSourceType(v string) (any, error) {
    result := MAILBOX_SOURCETYPE
    switch v {
        case "mailbox":
            result = MAILBOX_SOURCETYPE
        case "site":
            result = SITE_SOURCETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SOURCETYPE
        default:
            return 0, errors.New("Unknown SourceType value: " + v)
    }
    return &result, nil
}
func SerializeSourceType(values []SourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
