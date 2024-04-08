package models
import (
    "errors"
)
// 
type SimulationAutomationStatus int

const (
    UNKNOWN_SIMULATIONAUTOMATIONSTATUS SimulationAutomationStatus = iota
    DRAFT_SIMULATIONAUTOMATIONSTATUS
    NOTRUNNING_SIMULATIONAUTOMATIONSTATUS
    RUNNING_SIMULATIONAUTOMATIONSTATUS
    COMPLETED_SIMULATIONAUTOMATIONSTATUS
    UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONSTATUS
)

func (i SimulationAutomationStatus) String() string {
    return []string{"unknown", "draft", "notRunning", "running", "completed", "unknownFutureValue"}[i]
}
func ParseSimulationAutomationStatus(v string) (any, error) {
    result := UNKNOWN_SIMULATIONAUTOMATIONSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_SIMULATIONAUTOMATIONSTATUS
        case "draft":
            result = DRAFT_SIMULATIONAUTOMATIONSTATUS
        case "notRunning":
            result = NOTRUNNING_SIMULATIONAUTOMATIONSTATUS
        case "running":
            result = RUNNING_SIMULATIONAUTOMATIONSTATUS
        case "completed":
            result = COMPLETED_SIMULATIONAUTOMATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SIMULATIONAUTOMATIONSTATUS
        default:
            return 0, errors.New("Unknown SimulationAutomationStatus value: " + v)
    }
    return &result, nil
}
func SerializeSimulationAutomationStatus(values []SimulationAutomationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
