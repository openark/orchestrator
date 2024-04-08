package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateScheduledInstallable 
type WindowsUpdateScheduledInstallable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsUpdateInstallScheduleTypeable
    GetScheduledInstallDay()(*WeeklySchedule)
    GetScheduledInstallTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    SetScheduledInstallDay(value *WeeklySchedule)()
    SetScheduledInstallTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
}
