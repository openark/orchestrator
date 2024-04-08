package models
import (
    "errors"
)
// Possible values for firewallCertificateRevocationListCheckMethod
type FirewallCertificateRevocationListCheckMethodType int

const (
    // No value configured by Intune, do not override the user-configured device default value
    DEVICEDEFAULT_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE FirewallCertificateRevocationListCheckMethodType = iota
    // Do not check certificate revocation list
    NONE_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE
    // Attempt CRL check and allow a certificate only if the certificate is confirmed by the check
    ATTEMPT_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE
    // Require a successful CRL check before allowing a certificate
    REQUIRE_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE
)

func (i FirewallCertificateRevocationListCheckMethodType) String() string {
    return []string{"deviceDefault", "none", "attempt", "require"}[i]
}
func ParseFirewallCertificateRevocationListCheckMethodType(v string) (any, error) {
    result := DEVICEDEFAULT_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE
        case "none":
            result = NONE_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE
        case "attempt":
            result = ATTEMPT_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE
        case "require":
            result = REQUIRE_FIREWALLCERTIFICATEREVOCATIONLISTCHECKMETHODTYPE
        default:
            return 0, errors.New("Unknown FirewallCertificateRevocationListCheckMethodType value: " + v)
    }
    return &result, nil
}
func SerializeFirewallCertificateRevocationListCheckMethodType(values []FirewallCertificateRevocationListCheckMethodType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
