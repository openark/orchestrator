package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifiedDomainable 
type VerifiedDomainable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapabilities()(*string)
    GetIsDefault()(*bool)
    GetIsInitial()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetType()(*string)
    SetCapabilities(value *string)()
    SetIsDefault(value *bool)()
    SetIsInitial(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetType(value *string)()
}
