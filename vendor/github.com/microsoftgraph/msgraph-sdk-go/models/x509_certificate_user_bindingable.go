package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateUserBindingable 
type X509CertificateUserBindingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPriority()(*int32)
    GetUserProperty()(*string)
    GetX509CertificateField()(*string)
    SetOdataType(value *string)()
    SetPriority(value *int32)()
    SetUserProperty(value *string)()
    SetX509CertificateField(value *string)()
}
