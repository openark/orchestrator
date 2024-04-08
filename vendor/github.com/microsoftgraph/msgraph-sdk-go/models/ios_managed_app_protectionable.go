package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosManagedAppProtectionable 
type IosManagedAppProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TargetedManagedAppProtectionable
    GetAppDataEncryptionType()(*ManagedAppDataEncryptionType)
    GetApps()([]ManagedMobileAppable)
    GetCustomBrowserProtocol()(*string)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetFaceIdBlocked()(*bool)
    GetMinimumRequiredSdkVersion()(*string)
    SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)()
    SetApps(value []ManagedMobileAppable)()
    SetCustomBrowserProtocol(value *string)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetFaceIdBlocked(value *bool)()
    SetMinimumRequiredSdkVersion(value *string)()
}
