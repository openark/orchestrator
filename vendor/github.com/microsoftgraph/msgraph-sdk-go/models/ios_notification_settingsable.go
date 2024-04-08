package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosNotificationSettingsable 
type IosNotificationSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertType()(*IosNotificationAlertType)
    GetAppName()(*string)
    GetBadgesEnabled()(*bool)
    GetBundleID()(*string)
    GetEnabled()(*bool)
    GetOdataType()(*string)
    GetPublisher()(*string)
    GetShowInNotificationCenter()(*bool)
    GetShowOnLockScreen()(*bool)
    GetSoundsEnabled()(*bool)
    SetAlertType(value *IosNotificationAlertType)()
    SetAppName(value *string)()
    SetBadgesEnabled(value *bool)()
    SetBundleID(value *string)()
    SetEnabled(value *bool)()
    SetOdataType(value *string)()
    SetPublisher(value *string)()
    SetShowInNotificationCenter(value *bool)()
    SetShowOnLockScreen(value *bool)()
    SetSoundsEnabled(value *bool)()
}
