package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttendanceRecordable 
type AttendanceRecordable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendanceIntervals()([]AttendanceIntervalable)
    GetEmailAddress()(*string)
    GetIdentity()(Identityable)
    GetRole()(*string)
    GetTotalAttendanceInSeconds()(*int32)
    SetAttendanceIntervals(value []AttendanceIntervalable)()
    SetEmailAddress(value *string)()
    SetIdentity(value Identityable)()
    SetRole(value *string)()
    SetTotalAttendanceInSeconds(value *int32)()
}
