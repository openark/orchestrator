package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosStoreAppable 
type IosStoreAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableDeviceType()(IosDeviceTypeable)
    GetAppStoreUrl()(*string)
    GetBundleId()(*string)
    GetMinimumSupportedOperatingSystem()(IosMinimumOperatingSystemable)
    SetApplicableDeviceType(value IosDeviceTypeable)()
    SetAppStoreUrl(value *string)()
    SetBundleId(value *string)()
    SetMinimumSupportedOperatingSystem(value IosMinimumOperatingSystemable)()
}
