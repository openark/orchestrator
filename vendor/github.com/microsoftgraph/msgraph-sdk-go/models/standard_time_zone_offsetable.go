package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StandardTimeZoneOffsetable 
type StandardTimeZoneOffsetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDayOccurrence()(*int32)
    GetDayOfWeek()(*DayOfWeek)
    GetMonth()(*int32)
    GetOdataType()(*string)
    GetTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetYear()(*int32)
    SetDayOccurrence(value *int32)()
    SetDayOfWeek(value *DayOfWeek)()
    SetMonth(value *int32)()
    SetOdataType(value *string)()
    SetTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetYear(value *int32)()
}
