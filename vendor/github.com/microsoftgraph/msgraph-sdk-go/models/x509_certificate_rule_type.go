package models
import (
    "errors"
)
// 
type X509CertificateRuleType int

const (
    ISSUERSUBJECT_X509CERTIFICATERULETYPE X509CertificateRuleType = iota
    POLICYOID_X509CERTIFICATERULETYPE
    UNKNOWNFUTUREVALUE_X509CERTIFICATERULETYPE
)

func (i X509CertificateRuleType) String() string {
    return []string{"issuerSubject", "policyOID", "unknownFutureValue"}[i]
}
func ParseX509CertificateRuleType(v string) (any, error) {
    result := ISSUERSUBJECT_X509CERTIFICATERULETYPE
    switch v {
        case "issuerSubject":
            result = ISSUERSUBJECT_X509CERTIFICATERULETYPE
        case "policyOID":
            result = POLICYOID_X509CERTIFICATERULETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_X509CERTIFICATERULETYPE
        default:
            return 0, errors.New("Unknown X509CertificateRuleType value: " + v)
    }
    return &result, nil
}
func SerializeX509CertificateRuleType(values []X509CertificateRuleType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
