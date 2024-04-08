package models
import (
    "errors"
)
// 
type LocationUniqueIdType int

const (
    UNKNOWN_LOCATIONUNIQUEIDTYPE LocationUniqueIdType = iota
    LOCATIONSTORE_LOCATIONUNIQUEIDTYPE
    DIRECTORY_LOCATIONUNIQUEIDTYPE
    PRIVATE_LOCATIONUNIQUEIDTYPE
    BING_LOCATIONUNIQUEIDTYPE
)

func (i LocationUniqueIdType) String() string {
    return []string{"unknown", "locationStore", "directory", "private", "bing"}[i]
}
func ParseLocationUniqueIdType(v string) (any, error) {
    result := UNKNOWN_LOCATIONUNIQUEIDTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_LOCATIONUNIQUEIDTYPE
        case "locationStore":
            result = LOCATIONSTORE_LOCATIONUNIQUEIDTYPE
        case "directory":
            result = DIRECTORY_LOCATIONUNIQUEIDTYPE
        case "private":
            result = PRIVATE_LOCATIONUNIQUEIDTYPE
        case "bing":
            result = BING_LOCATIONUNIQUEIDTYPE
        default:
            return 0, errors.New("Unknown LocationUniqueIdType value: " + v)
    }
    return &result, nil
}
func SerializeLocationUniqueIdType(values []LocationUniqueIdType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
