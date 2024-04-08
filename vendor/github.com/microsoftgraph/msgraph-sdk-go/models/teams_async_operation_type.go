package models
import (
    "errors"
)
// 
type TeamsAsyncOperationType int

const (
    INVALID_TEAMSASYNCOPERATIONTYPE TeamsAsyncOperationType = iota
    CLONETEAM_TEAMSASYNCOPERATIONTYPE
    ARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
    UNARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
    CREATETEAM_TEAMSASYNCOPERATIONTYPE
    UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONTYPE
    TEAMIFYGROUP_TEAMSASYNCOPERATIONTYPE
    CREATECHANNEL_TEAMSASYNCOPERATIONTYPE
)

func (i TeamsAsyncOperationType) String() string {
    return []string{"invalid", "cloneTeam", "archiveTeam", "unarchiveTeam", "createTeam", "unknownFutureValue", "teamifyGroup", "createChannel"}[i]
}
func ParseTeamsAsyncOperationType(v string) (any, error) {
    result := INVALID_TEAMSASYNCOPERATIONTYPE
    switch v {
        case "invalid":
            result = INVALID_TEAMSASYNCOPERATIONTYPE
        case "cloneTeam":
            result = CLONETEAM_TEAMSASYNCOPERATIONTYPE
        case "archiveTeam":
            result = ARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
        case "unarchiveTeam":
            result = UNARCHIVETEAM_TEAMSASYNCOPERATIONTYPE
        case "createTeam":
            result = CREATETEAM_TEAMSASYNCOPERATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSASYNCOPERATIONTYPE
        case "teamifyGroup":
            result = TEAMIFYGROUP_TEAMSASYNCOPERATIONTYPE
        case "createChannel":
            result = CREATECHANNEL_TEAMSASYNCOPERATIONTYPE
        default:
            return 0, errors.New("Unknown TeamsAsyncOperationType value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAsyncOperationType(values []TeamsAsyncOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
