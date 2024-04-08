package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationDeviceStateSummaryable 
type DeviceConfigurationDeviceStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantDeviceCount()(*int32)
    GetConflictDeviceCount()(*int32)
    GetErrorDeviceCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetRemediatedDeviceCount()(*int32)
    GetUnknownDeviceCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetConflictDeviceCount(value *int32)()
    SetErrorDeviceCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetRemediatedDeviceCount(value *int32)()
    SetUnknownDeviceCount(value *int32)()
}
