package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultManagedAppProtectionable 
type DefaultManagedAppProtectionable interface {
    ManagedAppProtectionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDataEncryptionType()(*ManagedAppDataEncryptionType)
    GetApps()([]ManagedMobileAppable)
    GetCustomSettings()([]KeyValuePairable)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool)
    GetEncryptAppData()(*bool)
    GetFaceIdBlocked()(*bool)
    GetMinimumRequiredPatchVersion()(*string)
    GetMinimumRequiredSdkVersion()(*string)
    GetMinimumWarningPatchVersion()(*string)
    GetScreenCaptureBlocked()(*bool)
    SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)()
    SetApps(value []ManagedMobileAppable)()
    SetCustomSettings(value []KeyValuePairable)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)()
    SetEncryptAppData(value *bool)()
    SetFaceIdBlocked(value *bool)()
    SetMinimumRequiredPatchVersion(value *string)()
    SetMinimumRequiredSdkVersion(value *string)()
    SetMinimumWarningPatchVersion(value *string)()
    SetScreenCaptureBlocked(value *bool)()
}
