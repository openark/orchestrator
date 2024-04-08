package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingStaffMemberable 
type BookingStaffMemberable interface {
    BookingStaffMemberBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityIsAffectedByPersonalCalendar()(*bool)
    GetDisplayName()(*string)
    GetEmailAddress()(*string)
    GetIsEmailNotificationEnabled()(*bool)
    GetRole()(*BookingStaffRole)
    GetTimeZone()(*string)
    GetUseBusinessHours()(*bool)
    GetWorkingHours()([]BookingWorkHoursable)
    SetAvailabilityIsAffectedByPersonalCalendar(value *bool)()
    SetDisplayName(value *string)()
    SetEmailAddress(value *string)()
    SetIsEmailNotificationEnabled(value *bool)()
    SetRole(value *BookingStaffRole)()
    SetTimeZone(value *string)()
    SetUseBusinessHours(value *bool)()
    SetWorkingHours(value []BookingWorkHoursable)()
}
