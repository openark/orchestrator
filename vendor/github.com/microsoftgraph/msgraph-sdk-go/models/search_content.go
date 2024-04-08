package models
import (
    "errors"
)
// 
type SearchContent int

const (
    SHAREDCONTENT_SEARCHCONTENT SearchContent = iota
    PRIVATECONTENT_SEARCHCONTENT
    UNKNOWNFUTUREVALUE_SEARCHCONTENT
)

func (i SearchContent) String() string {
    return []string{"sharedContent", "privateContent", "unknownFutureValue"}[i]
}
func ParseSearchContent(v string) (any, error) {
    result := SHAREDCONTENT_SEARCHCONTENT
    switch v {
        case "sharedContent":
            result = SHAREDCONTENT_SEARCHCONTENT
        case "privateContent":
            result = PRIVATECONTENT_SEARCHCONTENT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SEARCHCONTENT
        default:
            return 0, errors.New("Unknown SearchContent value: " + v)
    }
    return &result, nil
}
func SerializeSearchContent(values []SearchContent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
