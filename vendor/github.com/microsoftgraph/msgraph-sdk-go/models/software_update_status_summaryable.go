package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SoftwareUpdateStatusSummaryable 
type SoftwareUpdateStatusSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantDeviceCount()(*int32)
    GetCompliantUserCount()(*int32)
    GetConflictDeviceCount()(*int32)
    GetConflictUserCount()(*int32)
    GetDisplayName()(*string)
    GetErrorDeviceCount()(*int32)
    GetErrorUserCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetNonCompliantUserCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetNotApplicableUserCount()(*int32)
    GetRemediatedDeviceCount()(*int32)
    GetRemediatedUserCount()(*int32)
    GetUnknownDeviceCount()(*int32)
    GetUnknownUserCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetCompliantUserCount(value *int32)()
    SetConflictDeviceCount(value *int32)()
    SetConflictUserCount(value *int32)()
    SetDisplayName(value *string)()
    SetErrorDeviceCount(value *int32)()
    SetErrorUserCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetNonCompliantUserCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetNotApplicableUserCount(value *int32)()
    SetRemediatedDeviceCount(value *int32)()
    SetRemediatedUserCount(value *int32)()
    SetUnknownDeviceCount(value *int32)()
    SetUnknownUserCount(value *int32)()
}
