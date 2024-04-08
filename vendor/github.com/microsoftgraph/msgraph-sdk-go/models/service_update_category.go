package models
import (
    "errors"
)
// 
type ServiceUpdateCategory int

const (
    PREVENTORFIXISSUE_SERVICEUPDATECATEGORY ServiceUpdateCategory = iota
    PLANFORCHANGE_SERVICEUPDATECATEGORY
    STAYINFORMED_SERVICEUPDATECATEGORY
    UNKNOWNFUTUREVALUE_SERVICEUPDATECATEGORY
)

func (i ServiceUpdateCategory) String() string {
    return []string{"preventOrFixIssue", "planForChange", "stayInformed", "unknownFutureValue"}[i]
}
func ParseServiceUpdateCategory(v string) (any, error) {
    result := PREVENTORFIXISSUE_SERVICEUPDATECATEGORY
    switch v {
        case "preventOrFixIssue":
            result = PREVENTORFIXISSUE_SERVICEUPDATECATEGORY
        case "planForChange":
            result = PLANFORCHANGE_SERVICEUPDATECATEGORY
        case "stayInformed":
            result = STAYINFORMED_SERVICEUPDATECATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SERVICEUPDATECATEGORY
        default:
            return 0, errors.New("Unknown ServiceUpdateCategory value: " + v)
    }
    return &result, nil
}
func SerializeServiceUpdateCategory(values []ServiceUpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
