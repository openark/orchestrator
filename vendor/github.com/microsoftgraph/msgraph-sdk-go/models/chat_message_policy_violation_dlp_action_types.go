package models
import (
    "errors"
)
// 
type ChatMessagePolicyViolationDlpActionTypes int

const (
    NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES ChatMessagePolicyViolationDlpActionTypes = iota
    NOTIFYSENDER_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
    BLOCKACCESS_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
    BLOCKACCESSEXTERNAL_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
)

func (i ChatMessagePolicyViolationDlpActionTypes) String() string {
    return []string{"none", "notifySender", "blockAccess", "blockAccessExternal"}[i]
}
func ParseChatMessagePolicyViolationDlpActionTypes(v string) (any, error) {
    result := NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
    switch v {
        case "none":
            result = NONE_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
        case "notifySender":
            result = NOTIFYSENDER_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
        case "blockAccess":
            result = BLOCKACCESS_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
        case "blockAccessExternal":
            result = BLOCKACCESSEXTERNAL_CHATMESSAGEPOLICYVIOLATIONDLPACTIONTYPES
        default:
            return 0, errors.New("Unknown ChatMessagePolicyViolationDlpActionTypes value: " + v)
    }
    return &result, nil
}
func SerializeChatMessagePolicyViolationDlpActionTypes(values []ChatMessagePolicyViolationDlpActionTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
