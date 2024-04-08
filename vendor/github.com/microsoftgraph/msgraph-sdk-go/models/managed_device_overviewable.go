package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceOverviewable 
type ManagedDeviceOverviewable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceExchangeAccessStateSummary()(DeviceExchangeAccessStateSummaryable)
    GetDeviceOperatingSystemSummary()(DeviceOperatingSystemSummaryable)
    GetDualEnrolledDeviceCount()(*int32)
    GetEnrolledDeviceCount()(*int32)
    GetMdmEnrolledCount()(*int32)
    SetDeviceExchangeAccessStateSummary(value DeviceExchangeAccessStateSummaryable)()
    SetDeviceOperatingSystemSummary(value DeviceOperatingSystemSummaryable)()
    SetDualEnrolledDeviceCount(value *int32)()
    SetEnrolledDeviceCount(value *int32)()
    SetMdmEnrolledCount(value *int32)()
}
