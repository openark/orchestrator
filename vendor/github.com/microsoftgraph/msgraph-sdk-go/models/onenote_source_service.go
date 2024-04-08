package models
import (
    "errors"
)
// 
type OnenoteSourceService int

const (
    UNKNOWN_ONENOTESOURCESERVICE OnenoteSourceService = iota
    ONEDRIVE_ONENOTESOURCESERVICE
    ONEDRIVEFORBUSINESS_ONENOTESOURCESERVICE
    ONPREMONEDRIVEFORBUSINESS_ONENOTESOURCESERVICE
)

func (i OnenoteSourceService) String() string {
    return []string{"Unknown", "OneDrive", "OneDriveForBusiness", "OnPremOneDriveForBusiness"}[i]
}
func ParseOnenoteSourceService(v string) (any, error) {
    result := UNKNOWN_ONENOTESOURCESERVICE
    switch v {
        case "Unknown":
            result = UNKNOWN_ONENOTESOURCESERVICE
        case "OneDrive":
            result = ONEDRIVE_ONENOTESOURCESERVICE
        case "OneDriveForBusiness":
            result = ONEDRIVEFORBUSINESS_ONENOTESOURCESERVICE
        case "OnPremOneDriveForBusiness":
            result = ONPREMONEDRIVEFORBUSINESS_ONENOTESOURCESERVICE
        default:
            return 0, errors.New("Unknown OnenoteSourceService value: " + v)
    }
    return &result, nil
}
func SerializeOnenoteSourceService(values []OnenoteSourceService) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
