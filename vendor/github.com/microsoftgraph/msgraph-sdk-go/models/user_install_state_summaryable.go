package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserInstallStateSummaryable 
type UserInstallStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceStates()([]DeviceInstallStateable)
    GetFailedDeviceCount()(*int32)
    GetInstalledDeviceCount()(*int32)
    GetNotInstalledDeviceCount()(*int32)
    GetUserName()(*string)
    SetDeviceStates(value []DeviceInstallStateable)()
    SetFailedDeviceCount(value *int32)()
    SetInstalledDeviceCount(value *int32)()
    SetNotInstalledDeviceCount(value *int32)()
    SetUserName(value *string)()
}
