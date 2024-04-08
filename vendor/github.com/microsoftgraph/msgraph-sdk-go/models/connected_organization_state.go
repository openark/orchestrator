package models
import (
    "errors"
)
// 
type ConnectedOrganizationState int

const (
    CONFIGURED_CONNECTEDORGANIZATIONSTATE ConnectedOrganizationState = iota
    PROPOSED_CONNECTEDORGANIZATIONSTATE
    UNKNOWNFUTUREVALUE_CONNECTEDORGANIZATIONSTATE
)

func (i ConnectedOrganizationState) String() string {
    return []string{"configured", "proposed", "unknownFutureValue"}[i]
}
func ParseConnectedOrganizationState(v string) (any, error) {
    result := CONFIGURED_CONNECTEDORGANIZATIONSTATE
    switch v {
        case "configured":
            result = CONFIGURED_CONNECTEDORGANIZATIONSTATE
        case "proposed":
            result = PROPOSED_CONNECTEDORGANIZATIONSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONNECTEDORGANIZATIONSTATE
        default:
            return 0, errors.New("Unknown ConnectedOrganizationState value: " + v)
    }
    return &result, nil
}
func SerializeConnectedOrganizationState(values []ConnectedOrganizationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
