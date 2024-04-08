package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateAuthenticationModeConfigurationable 
type X509CertificateAuthenticationModeConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetRules()([]X509CertificateRuleable)
    GetX509CertificateAuthenticationDefaultMode()(*X509CertificateAuthenticationMode)
    SetOdataType(value *string)()
    SetRules(value []X509CertificateRuleable)()
    SetX509CertificateAuthenticationDefaultMode(value *X509CertificateAuthenticationMode)()
}
