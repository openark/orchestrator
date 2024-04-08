package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterLocationable 
type PrinterLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAltitudeInMeters()(*int32)
    GetBuilding()(*string)
    GetCity()(*string)
    GetCountryOrRegion()(*string)
    GetFloor()(*string)
    GetFloorDescription()(*string)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    GetOdataType()(*string)
    GetOrganization()([]string)
    GetPostalCode()(*string)
    GetRoomDescription()(*string)
    GetRoomName()(*string)
    GetSite()(*string)
    GetStateOrProvince()(*string)
    GetStreetAddress()(*string)
    GetSubdivision()([]string)
    GetSubunit()([]string)
    SetAltitudeInMeters(value *int32)()
    SetBuilding(value *string)()
    SetCity(value *string)()
    SetCountryOrRegion(value *string)()
    SetFloor(value *string)()
    SetFloorDescription(value *string)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
    SetOdataType(value *string)()
    SetOrganization(value []string)()
    SetPostalCode(value *string)()
    SetRoomDescription(value *string)()
    SetRoomName(value *string)()
    SetSite(value *string)()
    SetStateOrProvince(value *string)()
    SetStreetAddress(value *string)()
    SetSubdivision(value []string)()
    SetSubunit(value []string)()
}
