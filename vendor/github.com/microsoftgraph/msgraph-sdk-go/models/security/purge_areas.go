package security
import (
    "errors"
)
// 
type PurgeAreas int

const (
    MAILBOXES_PURGEAREAS PurgeAreas = iota
    TEAMSMESSAGES_PURGEAREAS
    UNKNOWNFUTUREVALUE_PURGEAREAS
)

func (i PurgeAreas) String() string {
    return []string{"mailboxes", "teamsMessages", "unknownFutureValue"}[i]
}
func ParsePurgeAreas(v string) (any, error) {
    result := MAILBOXES_PURGEAREAS
    switch v {
        case "mailboxes":
            result = MAILBOXES_PURGEAREAS
        case "teamsMessages":
            result = TEAMSMESSAGES_PURGEAREAS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PURGEAREAS
        default:
            return 0, errors.New("Unknown PurgeAreas value: " + v)
    }
    return &result, nil
}
func SerializePurgeAreas(values []PurgeAreas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
