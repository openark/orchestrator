package models
import (
    "errors"
)
// 
type PlannerContainerType int

const (
    GROUP_PLANNERCONTAINERTYPE PlannerContainerType = iota
    UNKNOWNFUTUREVALUE_PLANNERCONTAINERTYPE
    ROSTER_PLANNERCONTAINERTYPE
)

func (i PlannerContainerType) String() string {
    return []string{"group", "unknownFutureValue", "roster"}[i]
}
func ParsePlannerContainerType(v string) (any, error) {
    result := GROUP_PLANNERCONTAINERTYPE
    switch v {
        case "group":
            result = GROUP_PLANNERCONTAINERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERCONTAINERTYPE
        case "roster":
            result = ROSTER_PLANNERCONTAINERTYPE
        default:
            return 0, errors.New("Unknown PlannerContainerType value: " + v)
    }
    return &result, nil
}
func SerializePlannerContainerType(values []PlannerContainerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
