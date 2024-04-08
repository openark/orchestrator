package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppAssignmentSettingsable 
type Win32LobAppAssignmentSettingsable interface {
    MobileAppAssignmentSettingsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeliveryOptimizationPriority()(*Win32LobAppDeliveryOptimizationPriority)
    GetInstallTimeSettings()(MobileAppInstallTimeSettingsable)
    GetNotifications()(*Win32LobAppNotification)
    GetRestartSettings()(Win32LobAppRestartSettingsable)
    SetDeliveryOptimizationPriority(value *Win32LobAppDeliveryOptimizationPriority)()
    SetInstallTimeSettings(value MobileAppInstallTimeSettingsable)()
    SetNotifications(value *Win32LobAppNotification)()
    SetRestartSettings(value Win32LobAppRestartSettingsable)()
}
