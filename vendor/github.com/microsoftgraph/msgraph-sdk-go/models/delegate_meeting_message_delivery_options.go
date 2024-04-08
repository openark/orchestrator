package models
import (
    "errors"
)
// 
type DelegateMeetingMessageDeliveryOptions int

const (
    SENDTODELEGATEANDINFORMATIONTOPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS DelegateMeetingMessageDeliveryOptions = iota
    SENDTODELEGATEANDPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
    SENDTODELEGATEONLY_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
)

func (i DelegateMeetingMessageDeliveryOptions) String() string {
    return []string{"sendToDelegateAndInformationToPrincipal", "sendToDelegateAndPrincipal", "sendToDelegateOnly"}[i]
}
func ParseDelegateMeetingMessageDeliveryOptions(v string) (any, error) {
    result := SENDTODELEGATEANDINFORMATIONTOPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
    switch v {
        case "sendToDelegateAndInformationToPrincipal":
            result = SENDTODELEGATEANDINFORMATIONTOPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
        case "sendToDelegateAndPrincipal":
            result = SENDTODELEGATEANDPRINCIPAL_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
        case "sendToDelegateOnly":
            result = SENDTODELEGATEONLY_DELEGATEMEETINGMESSAGEDELIVERYOPTIONS
        default:
            return 0, errors.New("Unknown DelegateMeetingMessageDeliveryOptions value: " + v)
    }
    return &result, nil
}
func SerializeDelegateMeetingMessageDeliveryOptions(values []DelegateMeetingMessageDeliveryOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
