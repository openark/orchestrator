package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedAppProtectionable 
type AndroidManagedAppProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TargetedManagedAppProtectionable
    GetApps()([]ManagedMobileAppable)
    GetCustomBrowserDisplayName()(*string)
    GetCustomBrowserPackageId()(*string)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool)
    GetEncryptAppData()(*bool)
    GetMinimumRequiredPatchVersion()(*string)
    GetMinimumWarningPatchVersion()(*string)
    GetScreenCaptureBlocked()(*bool)
    SetApps(value []ManagedMobileAppable)()
    SetCustomBrowserDisplayName(value *string)()
    SetCustomBrowserPackageId(value *string)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)()
    SetEncryptAppData(value *bool)()
    SetMinimumRequiredPatchVersion(value *string)()
    SetMinimumWarningPatchVersion(value *string)()
    SetScreenCaptureBlocked(value *bool)()
}
