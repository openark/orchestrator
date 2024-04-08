package models
import (
    "errors"
)
// 
type DelegatedAdminRelationshipRequestAction int

const (
    LOCKFORAPPROVAL_DELEGATEDADMINRELATIONSHIPREQUESTACTION DelegatedAdminRelationshipRequestAction = iota
    APPROVE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
    TERMINATE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
    UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
)

func (i DelegatedAdminRelationshipRequestAction) String() string {
    return []string{"lockForApproval", "approve", "terminate", "unknownFutureValue"}[i]
}
func ParseDelegatedAdminRelationshipRequestAction(v string) (any, error) {
    result := LOCKFORAPPROVAL_DELEGATEDADMINRELATIONSHIPREQUESTACTION
    switch v {
        case "lockForApproval":
            result = LOCKFORAPPROVAL_DELEGATEDADMINRELATIONSHIPREQUESTACTION
        case "approve":
            result = APPROVE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
        case "terminate":
            result = TERMINATE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DELEGATEDADMINRELATIONSHIPREQUESTACTION
        default:
            return 0, errors.New("Unknown DelegatedAdminRelationshipRequestAction value: " + v)
    }
    return &result, nil
}
func SerializeDelegatedAdminRelationshipRequestAction(values []DelegatedAdminRelationshipRequestAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
