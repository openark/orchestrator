package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosDeviceFeaturesConfigurationable 
type IosDeviceFeaturesConfigurationable interface {
    AppleDeviceFeaturesConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssetTagTemplate()(*string)
    GetHomeScreenDockIcons()([]IosHomeScreenItemable)
    GetHomeScreenPages()([]IosHomeScreenPageable)
    GetLockScreenFootnote()(*string)
    GetNotificationSettings()([]IosNotificationSettingsable)
    SetAssetTagTemplate(value *string)()
    SetHomeScreenDockIcons(value []IosHomeScreenItemable)()
    SetHomeScreenPages(value []IosHomeScreenPageable)()
    SetLockScreenFootnote(value *string)()
    SetNotificationSettings(value []IosNotificationSettingsable)()
}
