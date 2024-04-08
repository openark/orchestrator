package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlternativeSecurityIdable 
type AlternativeSecurityIdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentityProvider()(*string)
    GetKey()([]byte)
    GetOdataType()(*string)
    GetType()(*int32)
    SetIdentityProvider(value *string)()
    SetKey(value []byte)()
    SetOdataType(value *string)()
    SetType(value *int32)()
}
