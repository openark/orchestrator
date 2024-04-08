package models
import (
    "errors"
)
// 
type MicrosoftAuthenticatorAuthenticationMode int

const (
    DEVICEBASEDPUSH_MICROSOFTAUTHENTICATORAUTHENTICATIONMODE MicrosoftAuthenticatorAuthenticationMode = iota
    PUSH_MICROSOFTAUTHENTICATORAUTHENTICATIONMODE
    ANY_MICROSOFTAUTHENTICATORAUTHENTICATIONMODE
)

func (i MicrosoftAuthenticatorAuthenticationMode) String() string {
    return []string{"deviceBasedPush", "push", "any"}[i]
}
func ParseMicrosoftAuthenticatorAuthenticationMode(v string) (any, error) {
    result := DEVICEBASEDPUSH_MICROSOFTAUTHENTICATORAUTHENTICATIONMODE
    switch v {
        case "deviceBasedPush":
            result = DEVICEBASEDPUSH_MICROSOFTAUTHENTICATORAUTHENTICATIONMODE
        case "push":
            result = PUSH_MICROSOFTAUTHENTICATORAUTHENTICATIONMODE
        case "any":
            result = ANY_MICROSOFTAUTHENTICATORAUTHENTICATIONMODE
        default:
            return 0, errors.New("Unknown MicrosoftAuthenticatorAuthenticationMode value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftAuthenticatorAuthenticationMode(values []MicrosoftAuthenticatorAuthenticationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
