package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileCompliancePolicyable 
type AndroidWorkProfileCompliancePolicyable interface {
    DeviceCompliancePolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceThreatProtectionEnabled()(*bool)
    GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel)
    GetMinAndroidSecurityPatchLevel()(*string)
    GetOsMaximumVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeLock()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredType()(*AndroidRequiredPasswordType)
    GetSecurityBlockJailbrokenDevices()(*bool)
    GetSecurityDisableUsbDebugging()(*bool)
    GetSecurityPreventInstallAppsFromUnknownSources()(*bool)
    GetSecurityRequireCompanyPortalAppIntegrity()(*bool)
    GetSecurityRequireGooglePlayServices()(*bool)
    GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool)
    GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool)
    GetSecurityRequireUpToDateSecurityProviders()(*bool)
    GetSecurityRequireVerifyApps()(*bool)
    GetStorageRequireEncryption()(*bool)
    SetDeviceThreatProtectionEnabled(value *bool)()
    SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)()
    SetMinAndroidSecurityPatchLevel(value *string)()
    SetOsMaximumVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeLock(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredType(value *AndroidRequiredPasswordType)()
    SetSecurityBlockJailbrokenDevices(value *bool)()
    SetSecurityDisableUsbDebugging(value *bool)()
    SetSecurityPreventInstallAppsFromUnknownSources(value *bool)()
    SetSecurityRequireCompanyPortalAppIntegrity(value *bool)()
    SetSecurityRequireGooglePlayServices(value *bool)()
    SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)()
    SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)()
    SetSecurityRequireUpToDateSecurityProviders(value *bool)()
    SetSecurityRequireVerifyApps(value *bool)()
    SetStorageRequireEncryption(value *bool)()
}
