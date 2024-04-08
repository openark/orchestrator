package models
import (
    "errors"
)
// 
type TeamsAsyncOperationStatus int

const (
    INVALID_TEAMSASYNCOPERATIONSTATUS TeamsAsyncOperationStatus = iota
    NOTSTARTED_TEAMSASYNCOPERATIONSTATUS
    INPROGRESS_TEAMSASYNCOPERATIONSTATUS
    SUCCEEDED_TEAMSASYNCOPERATIONSTATUS
    FAILED_TEAMSASYNCOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONSTATUS
)

func (i TeamsAsyncOperationStatus) String() string {
    return []string{"invalid", "notStarted", "inProgress", "succeeded", "failed", "unknownFutureValue"}[i]
}
func ParseTeamsAsyncOperationStatus(v string) (any, error) {
    result := INVALID_TEAMSASYNCOPERATIONSTATUS
    switch v {
        case "invalid":
            result = INVALID_TEAMSASYNCOPERATIONSTATUS
        case "notStarted":
            result = NOTSTARTED_TEAMSASYNCOPERATIONSTATUS
        case "inProgress":
            result = INPROGRESS_TEAMSASYNCOPERATIONSTATUS
        case "succeeded":
            result = SUCCEEDED_TEAMSASYNCOPERATIONSTATUS
        case "failed":
            result = FAILED_TEAMSASYNCOPERATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown TeamsAsyncOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAsyncOperationStatus(values []TeamsAsyncOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
