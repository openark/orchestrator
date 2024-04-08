package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedAppable 
type DetectedAppable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceCount()(*int32)
    GetDisplayName()(*string)
    GetManagedDevices()([]ManagedDeviceable)
    GetPlatform()(*DetectedAppPlatformType)
    GetPublisher()(*string)
    GetSizeInByte()(*int64)
    GetVersion()(*string)
    SetDeviceCount(value *int32)()
    SetDisplayName(value *string)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetPlatform(value *DetectedAppPlatformType)()
    SetPublisher(value *string)()
    SetSizeInByte(value *int64)()
    SetVersion(value *string)()
}
