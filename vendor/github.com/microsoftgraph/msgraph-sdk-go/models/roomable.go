package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Roomable 
type Roomable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Placeable
    GetAudioDeviceName()(*string)
    GetBookingType()(*BookingType)
    GetBuilding()(*string)
    GetCapacity()(*int32)
    GetDisplayDeviceName()(*string)
    GetEmailAddress()(*string)
    GetFloorLabel()(*string)
    GetFloorNumber()(*int32)
    GetIsWheelChairAccessible()(*bool)
    GetLabel()(*string)
    GetNickname()(*string)
    GetTags()([]string)
    GetVideoDeviceName()(*string)
    SetAudioDeviceName(value *string)()
    SetBookingType(value *BookingType)()
    SetBuilding(value *string)()
    SetCapacity(value *int32)()
    SetDisplayDeviceName(value *string)()
    SetEmailAddress(value *string)()
    SetFloorLabel(value *string)()
    SetFloorNumber(value *int32)()
    SetIsWheelChairAccessible(value *bool)()
    SetLabel(value *string)()
    SetNickname(value *string)()
    SetTags(value []string)()
    SetVideoDeviceName(value *string)()
}
