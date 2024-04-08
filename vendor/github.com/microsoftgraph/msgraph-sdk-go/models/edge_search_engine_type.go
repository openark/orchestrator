package models
import (
    "errors"
)
// Allows IT admind to set a predefined default search engine for MDM-Controlled devices
type EdgeSearchEngineType int

const (
    // Uses factory settings of Edge to assign the default search engine as per the user market
    DEFAULTESCAPED_EDGESEARCHENGINETYPE EdgeSearchEngineType = iota
    // Sets Bing as the default search engine
    BING_EDGESEARCHENGINETYPE
)

func (i EdgeSearchEngineType) String() string {
    return []string{"default", "bing"}[i]
}
func ParseEdgeSearchEngineType(v string) (any, error) {
    result := DEFAULTESCAPED_EDGESEARCHENGINETYPE
    switch v {
        case "default":
            result = DEFAULTESCAPED_EDGESEARCHENGINETYPE
        case "bing":
            result = BING_EDGESEARCHENGINETYPE
        default:
            return 0, errors.New("Unknown EdgeSearchEngineType value: " + v)
    }
    return &result, nil
}
func SerializeEdgeSearchEngineType(values []EdgeSearchEngineType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
