package models
import (
    "errors"
)
// 
type TeamsAppDistributionMethod int

const (
    STORE_TEAMSAPPDISTRIBUTIONMETHOD TeamsAppDistributionMethod = iota
    ORGANIZATION_TEAMSAPPDISTRIBUTIONMETHOD
    SIDELOADED_TEAMSAPPDISTRIBUTIONMETHOD
    UNKNOWNFUTUREVALUE_TEAMSAPPDISTRIBUTIONMETHOD
)

func (i TeamsAppDistributionMethod) String() string {
    return []string{"store", "organization", "sideloaded", "unknownFutureValue"}[i]
}
func ParseTeamsAppDistributionMethod(v string) (any, error) {
    result := STORE_TEAMSAPPDISTRIBUTIONMETHOD
    switch v {
        case "store":
            result = STORE_TEAMSAPPDISTRIBUTIONMETHOD
        case "organization":
            result = ORGANIZATION_TEAMSAPPDISTRIBUTIONMETHOD
        case "sideloaded":
            result = SIDELOADED_TEAMSAPPDISTRIBUTIONMETHOD
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSAPPDISTRIBUTIONMETHOD
        default:
            return 0, errors.New("Unknown TeamsAppDistributionMethod value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAppDistributionMethod(values []TeamsAppDistributionMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
