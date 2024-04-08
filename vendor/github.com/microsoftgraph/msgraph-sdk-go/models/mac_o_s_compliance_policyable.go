package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSCompliancePolicyable 
type MacOSCompliancePolicyable interface {
    DeviceCompliancePolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceThreatProtectionEnabled()(*bool)
    GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel)
    GetFirewallBlockAllIncoming()(*bool)
    GetFirewallEnabled()(*bool)
    GetFirewallEnableStealthMode()(*bool)
    GetOsMaximumVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPasswordBlockSimple()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumCharacterSetCount()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeLock()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredType()(*RequiredPasswordType)
    GetStorageRequireEncryption()(*bool)
    GetSystemIntegrityProtectionEnabled()(*bool)
    SetDeviceThreatProtectionEnabled(value *bool)()
    SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)()
    SetFirewallBlockAllIncoming(value *bool)()
    SetFirewallEnabled(value *bool)()
    SetFirewallEnableStealthMode(value *bool)()
    SetOsMaximumVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPasswordBlockSimple(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumCharacterSetCount(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeLock(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredType(value *RequiredPasswordType)()
    SetStorageRequireEncryption(value *bool)()
    SetSystemIntegrityProtectionEnabled(value *bool)()
}
