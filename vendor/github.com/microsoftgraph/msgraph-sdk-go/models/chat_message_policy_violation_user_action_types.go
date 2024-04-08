package models
import (
    "errors"
)
// 
type ChatMessagePolicyViolationUserActionTypes int

const (
    NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES ChatMessagePolicyViolationUserActionTypes = iota
    OVERRIDE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
    REPORTFALSEPOSITIVE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
)

func (i ChatMessagePolicyViolationUserActionTypes) String() string {
    return []string{"none", "override", "reportFalsePositive"}[i]
}
func ParseChatMessagePolicyViolationUserActionTypes(v string) (any, error) {
    result := NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
    switch v {
        case "none":
            result = NONE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
        case "override":
            result = OVERRIDE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
        case "reportFalsePositive":
            result = REPORTFALSEPOSITIVE_CHATMESSAGEPOLICYVIOLATIONUSERACTIONTYPES
        default:
            return 0, errors.New("Unknown ChatMessagePolicyViolationUserActionTypes value: " + v)
    }
    return &result, nil
}
func SerializeChatMessagePolicyViolationUserActionTypes(values []ChatMessagePolicyViolationUserActionTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
