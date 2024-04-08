package models
import (
    "errors"
)
// 
type RecipientScopeType int

const (
    NONE_RECIPIENTSCOPETYPE RecipientScopeType = iota
    INTERNAL_RECIPIENTSCOPETYPE
    EXTERNAL_RECIPIENTSCOPETYPE
    EXTERNALPARTNER_RECIPIENTSCOPETYPE
    EXTERNALNONPARTNER_RECIPIENTSCOPETYPE
)

func (i RecipientScopeType) String() string {
    return []string{"none", "internal", "external", "externalPartner", "externalNonPartner"}[i]
}
func ParseRecipientScopeType(v string) (any, error) {
    result := NONE_RECIPIENTSCOPETYPE
    switch v {
        case "none":
            result = NONE_RECIPIENTSCOPETYPE
        case "internal":
            result = INTERNAL_RECIPIENTSCOPETYPE
        case "external":
            result = EXTERNAL_RECIPIENTSCOPETYPE
        case "externalPartner":
            result = EXTERNALPARTNER_RECIPIENTSCOPETYPE
        case "externalNonPartner":
            result = EXTERNALNONPARTNER_RECIPIENTSCOPETYPE
        default:
            return 0, errors.New("Unknown RecipientScopeType value: " + v)
    }
    return &result, nil
}
func SerializeRecipientScopeType(values []RecipientScopeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
