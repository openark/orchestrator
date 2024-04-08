package models
import (
    "errors"
)
// 
type ProvisioningStatusErrorCategory int

const (
    FAILURE_PROVISIONINGSTATUSERRORCATEGORY ProvisioningStatusErrorCategory = iota
    NONSERVICEFAILURE_PROVISIONINGSTATUSERRORCATEGORY
    SUCCESS_PROVISIONINGSTATUSERRORCATEGORY
    UNKNOWNFUTUREVALUE_PROVISIONINGSTATUSERRORCATEGORY
)

func (i ProvisioningStatusErrorCategory) String() string {
    return []string{"failure", "nonServiceFailure", "success", "unknownFutureValue"}[i]
}
func ParseProvisioningStatusErrorCategory(v string) (any, error) {
    result := FAILURE_PROVISIONINGSTATUSERRORCATEGORY
    switch v {
        case "failure":
            result = FAILURE_PROVISIONINGSTATUSERRORCATEGORY
        case "nonServiceFailure":
            result = NONSERVICEFAILURE_PROVISIONINGSTATUSERRORCATEGORY
        case "success":
            result = SUCCESS_PROVISIONINGSTATUSERRORCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PROVISIONINGSTATUSERRORCATEGORY
        default:
            return 0, errors.New("Unknown ProvisioningStatusErrorCategory value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningStatusErrorCategory(values []ProvisioningStatusErrorCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
