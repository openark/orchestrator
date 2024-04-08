package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PhysicalOfficeAddressable 
type PhysicalOfficeAddressable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountryOrRegion()(*string)
    GetOdataType()(*string)
    GetOfficeLocation()(*string)
    GetPostalCode()(*string)
    GetState()(*string)
    GetStreet()(*string)
    SetCity(value *string)()
    SetCountryOrRegion(value *string)()
    SetOdataType(value *string)()
    SetOfficeLocation(value *string)()
    SetPostalCode(value *string)()
    SetState(value *string)()
    SetStreet(value *string)()
}
