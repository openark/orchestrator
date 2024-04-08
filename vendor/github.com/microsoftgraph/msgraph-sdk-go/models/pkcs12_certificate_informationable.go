package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Pkcs12CertificateInformationable 
type Pkcs12CertificateInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsActive()(*bool)
    GetNotAfter()(*int64)
    GetNotBefore()(*int64)
    GetOdataType()(*string)
    GetThumbprint()(*string)
    SetIsActive(value *bool)()
    SetNotAfter(value *int64)()
    SetNotBefore(value *int64)()
    SetOdataType(value *string)()
    SetThumbprint(value *string)()
}
