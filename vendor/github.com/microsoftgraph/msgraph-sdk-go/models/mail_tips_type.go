package models
import (
    "errors"
)
// 
type MailTipsType int

const (
    AUTOMATICREPLIES_MAILTIPSTYPE MailTipsType = iota
    MAILBOXFULLSTATUS_MAILTIPSTYPE
    CUSTOMMAILTIP_MAILTIPSTYPE
    EXTERNALMEMBERCOUNT_MAILTIPSTYPE
    TOTALMEMBERCOUNT_MAILTIPSTYPE
    MAXMESSAGESIZE_MAILTIPSTYPE
    DELIVERYRESTRICTION_MAILTIPSTYPE
    MODERATIONSTATUS_MAILTIPSTYPE
    RECIPIENTSCOPE_MAILTIPSTYPE
    RECIPIENTSUGGESTIONS_MAILTIPSTYPE
)

func (i MailTipsType) String() string {
    return []string{"automaticReplies", "mailboxFullStatus", "customMailTip", "externalMemberCount", "totalMemberCount", "maxMessageSize", "deliveryRestriction", "moderationStatus", "recipientScope", "recipientSuggestions"}[i]
}
func ParseMailTipsType(v string) (any, error) {
    result := AUTOMATICREPLIES_MAILTIPSTYPE
    switch v {
        case "automaticReplies":
            result = AUTOMATICREPLIES_MAILTIPSTYPE
        case "mailboxFullStatus":
            result = MAILBOXFULLSTATUS_MAILTIPSTYPE
        case "customMailTip":
            result = CUSTOMMAILTIP_MAILTIPSTYPE
        case "externalMemberCount":
            result = EXTERNALMEMBERCOUNT_MAILTIPSTYPE
        case "totalMemberCount":
            result = TOTALMEMBERCOUNT_MAILTIPSTYPE
        case "maxMessageSize":
            result = MAXMESSAGESIZE_MAILTIPSTYPE
        case "deliveryRestriction":
            result = DELIVERYRESTRICTION_MAILTIPSTYPE
        case "moderationStatus":
            result = MODERATIONSTATUS_MAILTIPSTYPE
        case "recipientScope":
            result = RECIPIENTSCOPE_MAILTIPSTYPE
        case "recipientSuggestions":
            result = RECIPIENTSUGGESTIONS_MAILTIPSTYPE
        default:
            return 0, errors.New("Unknown MailTipsType value: " + v)
    }
    return &result, nil
}
func SerializeMailTipsType(values []MailTipsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
