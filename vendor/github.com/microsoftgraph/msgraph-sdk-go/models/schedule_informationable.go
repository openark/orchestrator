package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScheduleInformationable 
type ScheduleInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityView()(*string)
    GetError()(FreeBusyErrorable)
    GetOdataType()(*string)
    GetScheduleId()(*string)
    GetScheduleItems()([]ScheduleItemable)
    GetWorkingHours()(WorkingHoursable)
    SetAvailabilityView(value *string)()
    SetError(value FreeBusyErrorable)()
    SetOdataType(value *string)()
    SetScheduleId(value *string)()
    SetScheduleItems(value []ScheduleItemable)()
    SetWorkingHours(value WorkingHoursable)()
}
