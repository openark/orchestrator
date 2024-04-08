package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10TeamGeneralConfigurationable 
type Windows10TeamGeneralConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureOperationalInsightsBlockTelemetry()(*bool)
    GetAzureOperationalInsightsWorkspaceId()(*string)
    GetAzureOperationalInsightsWorkspaceKey()(*string)
    GetConnectAppBlockAutoLaunch()(*bool)
    GetMaintenanceWindowBlocked()(*bool)
    GetMaintenanceWindowDurationInHours()(*int32)
    GetMaintenanceWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetMiracastBlocked()(*bool)
    GetMiracastChannel()(*MiracastChannel)
    GetMiracastRequirePin()(*bool)
    GetSettingsBlockMyMeetingsAndFiles()(*bool)
    GetSettingsBlockSessionResume()(*bool)
    GetSettingsBlockSigninSuggestions()(*bool)
    GetSettingsDefaultVolume()(*int32)
    GetSettingsScreenTimeoutInMinutes()(*int32)
    GetSettingsSessionTimeoutInMinutes()(*int32)
    GetSettingsSleepTimeoutInMinutes()(*int32)
    GetWelcomeScreenBackgroundImageUrl()(*string)
    GetWelcomeScreenBlockAutomaticWakeUp()(*bool)
    GetWelcomeScreenMeetingInformation()(*WelcomeScreenMeetingInformation)
    SetAzureOperationalInsightsBlockTelemetry(value *bool)()
    SetAzureOperationalInsightsWorkspaceId(value *string)()
    SetAzureOperationalInsightsWorkspaceKey(value *string)()
    SetConnectAppBlockAutoLaunch(value *bool)()
    SetMaintenanceWindowBlocked(value *bool)()
    SetMaintenanceWindowDurationInHours(value *int32)()
    SetMaintenanceWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetMiracastBlocked(value *bool)()
    SetMiracastChannel(value *MiracastChannel)()
    SetMiracastRequirePin(value *bool)()
    SetSettingsBlockMyMeetingsAndFiles(value *bool)()
    SetSettingsBlockSessionResume(value *bool)()
    SetSettingsBlockSigninSuggestions(value *bool)()
    SetSettingsDefaultVolume(value *int32)()
    SetSettingsScreenTimeoutInMinutes(value *int32)()
    SetSettingsSessionTimeoutInMinutes(value *int32)()
    SetSettingsSleepTimeoutInMinutes(value *int32)()
    SetWelcomeScreenBackgroundImageUrl(value *string)()
    SetWelcomeScreenBlockAutomaticWakeUp(value *bool)()
    SetWelcomeScreenMeetingInformation(value *WelcomeScreenMeetingInformation)()
}
