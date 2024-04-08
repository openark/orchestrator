package models
import (
    "errors"
)
// 
type ClonableTeamParts int

const (
    APPS_CLONABLETEAMPARTS ClonableTeamParts = iota
    TABS_CLONABLETEAMPARTS
    SETTINGS_CLONABLETEAMPARTS
    CHANNELS_CLONABLETEAMPARTS
    MEMBERS_CLONABLETEAMPARTS
)

func (i ClonableTeamParts) String() string {
    return []string{"apps", "tabs", "settings", "channels", "members"}[i]
}
func ParseClonableTeamParts(v string) (any, error) {
    result := APPS_CLONABLETEAMPARTS
    switch v {
        case "apps":
            result = APPS_CLONABLETEAMPARTS
        case "tabs":
            result = TABS_CLONABLETEAMPARTS
        case "settings":
            result = SETTINGS_CLONABLETEAMPARTS
        case "channels":
            result = CHANNELS_CLONABLETEAMPARTS
        case "members":
            result = MEMBERS_CLONABLETEAMPARTS
        default:
            return 0, errors.New("Unknown ClonableTeamParts value: " + v)
    }
    return &result, nil
}
func SerializeClonableTeamParts(values []ClonableTeamParts) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
