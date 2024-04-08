package models
import (
    "errors"
)
// 
type PostType int

const (
    REGULAR_POSTTYPE PostType = iota
    QUICK_POSTTYPE
    STRATEGIC_POSTTYPE
    UNKNOWNFUTUREVALUE_POSTTYPE
)

func (i PostType) String() string {
    return []string{"regular", "quick", "strategic", "unknownFutureValue"}[i]
}
func ParsePostType(v string) (any, error) {
    result := REGULAR_POSTTYPE
    switch v {
        case "regular":
            result = REGULAR_POSTTYPE
        case "quick":
            result = QUICK_POSTTYPE
        case "strategic":
            result = STRATEGIC_POSTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_POSTTYPE
        default:
            return 0, errors.New("Unknown PostType value: " + v)
    }
    return &result, nil
}
func SerializePostType(values []PostType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
