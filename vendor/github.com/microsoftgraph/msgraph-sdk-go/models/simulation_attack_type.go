package models
import (
    "errors"
)
// 
type SimulationAttackType int

const (
    UNKNOWN_SIMULATIONATTACKTYPE SimulationAttackType = iota
    SOCIAL_SIMULATIONATTACKTYPE
    CLOUD_SIMULATIONATTACKTYPE
    ENDPOINT_SIMULATIONATTACKTYPE
    UNKNOWNFUTUREVALUE_SIMULATIONATTACKTYPE
)

func (i SimulationAttackType) String() string {
    return []string{"unknown", "social", "cloud", "endpoint", "unknownFutureValue"}[i]
}
func ParseSimulationAttackType(v string) (any, error) {
    result := UNKNOWN_SIMULATIONATTACKTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_SIMULATIONATTACKTYPE
        case "social":
            result = SOCIAL_SIMULATIONATTACKTYPE
        case "cloud":
            result = CLOUD_SIMULATIONATTACKTYPE
        case "endpoint":
            result = ENDPOINT_SIMULATIONATTACKTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SIMULATIONATTACKTYPE
        default:
            return 0, errors.New("Unknown SimulationAttackType value: " + v)
    }
    return &result, nil
}
func SerializeSimulationAttackType(values []SimulationAttackType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
