package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScheduleItemable 
type ScheduleItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnd()(DateTimeTimeZoneable)
    GetIsPrivate()(*bool)
    GetLocation()(*string)
    GetOdataType()(*string)
    GetStart()(DateTimeTimeZoneable)
    GetStatus()(*FreeBusyStatus)
    GetSubject()(*string)
    SetEnd(value DateTimeTimeZoneable)()
    SetIsPrivate(value *bool)()
    SetLocation(value *string)()
    SetOdataType(value *string)()
    SetStart(value DateTimeTimeZoneable)()
    SetStatus(value *FreeBusyStatus)()
    SetSubject(value *string)()
}
