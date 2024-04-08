package models
import (
    "errors"
)
// 
type ChatMessagePolicyViolationVerdictDetailsTypes int

const (
    NONE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES ChatMessagePolicyViolationVerdictDetailsTypes = iota
    ALLOWFALSEPOSITIVEOVERRIDE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
    ALLOWOVERRIDEWITHOUTJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
    ALLOWOVERRIDEWITHJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
)

func (i ChatMessagePolicyViolationVerdictDetailsTypes) String() string {
    return []string{"none", "allowFalsePositiveOverride", "allowOverrideWithoutJustification", "allowOverrideWithJustification"}[i]
}
func ParseChatMessagePolicyViolationVerdictDetailsTypes(v string) (any, error) {
    result := NONE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
    switch v {
        case "none":
            result = NONE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
        case "allowFalsePositiveOverride":
            result = ALLOWFALSEPOSITIVEOVERRIDE_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
        case "allowOverrideWithoutJustification":
            result = ALLOWOVERRIDEWITHOUTJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
        case "allowOverrideWithJustification":
            result = ALLOWOVERRIDEWITHJUSTIFICATION_CHATMESSAGEPOLICYVIOLATIONVERDICTDETAILSTYPES
        default:
            return 0, errors.New("Unknown ChatMessagePolicyViolationVerdictDetailsTypes value: " + v)
    }
    return &result, nil
}
func SerializeChatMessagePolicyViolationVerdictDetailsTypes(values []ChatMessagePolicyViolationVerdictDetailsTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
