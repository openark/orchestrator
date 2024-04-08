package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionAppable 
type WindowsInformationProtectionAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDenied()(*bool)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetProductName()(*string)
    GetPublisherName()(*string)
    SetDenied(value *bool)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetProductName(value *string)()
    SetPublisherName(value *string)()
}
