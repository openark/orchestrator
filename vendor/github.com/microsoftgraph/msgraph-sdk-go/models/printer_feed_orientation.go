package models
import (
    "errors"
)
// 
type PrinterFeedOrientation int

const (
    LONGEDGEFIRST_PRINTERFEEDORIENTATION PrinterFeedOrientation = iota
    SHORTEDGEFIRST_PRINTERFEEDORIENTATION
    UNKNOWNFUTUREVALUE_PRINTERFEEDORIENTATION
)

func (i PrinterFeedOrientation) String() string {
    return []string{"longEdgeFirst", "shortEdgeFirst", "unknownFutureValue"}[i]
}
func ParsePrinterFeedOrientation(v string) (any, error) {
    result := LONGEDGEFIRST_PRINTERFEEDORIENTATION
    switch v {
        case "longEdgeFirst":
            result = LONGEDGEFIRST_PRINTERFEEDORIENTATION
        case "shortEdgeFirst":
            result = SHORTEDGEFIRST_PRINTERFEEDORIENTATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRINTERFEEDORIENTATION
        default:
            return 0, errors.New("Unknown PrinterFeedOrientation value: " + v)
    }
    return &result, nil
}
func SerializePrinterFeedOrientation(values []PrinterFeedOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
