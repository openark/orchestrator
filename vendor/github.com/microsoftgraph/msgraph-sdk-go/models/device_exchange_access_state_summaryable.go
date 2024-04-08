package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceExchangeAccessStateSummaryable 
type DeviceExchangeAccessStateSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedDeviceCount()(*int32)
    GetBlockedDeviceCount()(*int32)
    GetOdataType()(*string)
    GetQuarantinedDeviceCount()(*int32)
    GetUnavailableDeviceCount()(*int32)
    GetUnknownDeviceCount()(*int32)
    SetAllowedDeviceCount(value *int32)()
    SetBlockedDeviceCount(value *int32)()
    SetOdataType(value *string)()
    SetQuarantinedDeviceCount(value *int32)()
    SetUnavailableDeviceCount(value *int32)()
    SetUnknownDeviceCount(value *int32)()
}
