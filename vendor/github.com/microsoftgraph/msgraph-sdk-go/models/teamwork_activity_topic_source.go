package models
import (
    "errors"
)
// 
type TeamworkActivityTopicSource int

const (
    ENTITYURL_TEAMWORKACTIVITYTOPICSOURCE TeamworkActivityTopicSource = iota
    TEXT_TEAMWORKACTIVITYTOPICSOURCE
)

func (i TeamworkActivityTopicSource) String() string {
    return []string{"entityUrl", "text"}[i]
}
func ParseTeamworkActivityTopicSource(v string) (any, error) {
    result := ENTITYURL_TEAMWORKACTIVITYTOPICSOURCE
    switch v {
        case "entityUrl":
            result = ENTITYURL_TEAMWORKACTIVITYTOPICSOURCE
        case "text":
            result = TEXT_TEAMWORKACTIVITYTOPICSOURCE
        default:
            return 0, errors.New("Unknown TeamworkActivityTopicSource value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkActivityTopicSource(values []TeamworkActivityTopicSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
