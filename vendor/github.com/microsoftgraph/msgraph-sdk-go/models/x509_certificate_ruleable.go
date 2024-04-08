package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateRuleable 
type X509CertificateRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentifier()(*string)
    GetOdataType()(*string)
    GetX509CertificateAuthenticationMode()(*X509CertificateAuthenticationMode)
    GetX509CertificateRuleType()(*X509CertificateRuleType)
    SetIdentifier(value *string)()
    SetOdataType(value *string)()
    SetX509CertificateAuthenticationMode(value *X509CertificateAuthenticationMode)()
    SetX509CertificateRuleType(value *X509CertificateRuleType)()
}
