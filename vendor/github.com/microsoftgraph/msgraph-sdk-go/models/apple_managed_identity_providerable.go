package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleManagedIdentityProviderable 
type AppleManagedIdentityProviderable interface {
    IdentityProviderBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateData()(*string)
    GetDeveloperId()(*string)
    GetKeyId()(*string)
    GetServiceId()(*string)
    SetCertificateData(value *string)()
    SetDeveloperId(value *string)()
    SetKeyId(value *string)()
    SetServiceId(value *string)()
}
