package security
import (
    "errors"
)
// 
type AdditionalDataOptions int

const (
    ALLVERSIONS_ADDITIONALDATAOPTIONS AdditionalDataOptions = iota
    LINKEDFILES_ADDITIONALDATAOPTIONS
    UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS
)

func (i AdditionalDataOptions) String() string {
    return []string{"allVersions", "linkedFiles", "unknownFutureValue"}[i]
}
func ParseAdditionalDataOptions(v string) (any, error) {
    result := ALLVERSIONS_ADDITIONALDATAOPTIONS
    switch v {
        case "allVersions":
            result = ALLVERSIONS_ADDITIONALDATAOPTIONS
        case "linkedFiles":
            result = LINKEDFILES_ADDITIONALDATAOPTIONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS
        default:
            return 0, errors.New("Unknown AdditionalDataOptions value: " + v)
    }
    return &result, nil
}
func SerializeAdditionalDataOptions(values []AdditionalDataOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
