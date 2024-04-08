package models
import (
    "errors"
)
// 
type X509CertificateAuthenticationMode int

const (
    X509CERTIFICATESINGLEFACTOR_X509CERTIFICATEAUTHENTICATIONMODE X509CertificateAuthenticationMode = iota
    X509CERTIFICATEMULTIFACTOR_X509CERTIFICATEAUTHENTICATIONMODE
    UNKNOWNFUTUREVALUE_X509CERTIFICATEAUTHENTICATIONMODE
)

func (i X509CertificateAuthenticationMode) String() string {
    return []string{"x509CertificateSingleFactor", "x509CertificateMultiFactor", "unknownFutureValue"}[i]
}
func ParseX509CertificateAuthenticationMode(v string) (any, error) {
    result := X509CERTIFICATESINGLEFACTOR_X509CERTIFICATEAUTHENTICATIONMODE
    switch v {
        case "x509CertificateSingleFactor":
            result = X509CERTIFICATESINGLEFACTOR_X509CERTIFICATEAUTHENTICATIONMODE
        case "x509CertificateMultiFactor":
            result = X509CERTIFICATEMULTIFACTOR_X509CERTIFICATEAUTHENTICATIONMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_X509CERTIFICATEAUTHENTICATIONMODE
        default:
            return 0, errors.New("Unknown X509CertificateAuthenticationMode value: " + v)
    }
    return &result, nil
}
func SerializeX509CertificateAuthenticationMode(values []X509CertificateAuthenticationMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
