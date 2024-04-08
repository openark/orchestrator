package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10CompliancePolicyable 
type Windows10CompliancePolicyable interface {
    DeviceCompliancePolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBitLockerEnabled()(*bool)
    GetCodeIntegrityEnabled()(*bool)
    GetEarlyLaunchAntiMalwareDriverEnabled()(*bool)
    GetMobileOsMaximumVersion()(*string)
    GetMobileOsMinimumVersion()(*string)
    GetOsMaximumVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPasswordBlockSimple()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumCharacterSetCount()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeLock()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredToUnlockFromIdle()(*bool)
    GetPasswordRequiredType()(*RequiredPasswordType)
    GetRequireHealthyDeviceReport()(*bool)
    GetSecureBootEnabled()(*bool)
    GetStorageRequireEncryption()(*bool)
    SetBitLockerEnabled(value *bool)()
    SetCodeIntegrityEnabled(value *bool)()
    SetEarlyLaunchAntiMalwareDriverEnabled(value *bool)()
    SetMobileOsMaximumVersion(value *string)()
    SetMobileOsMinimumVersion(value *string)()
    SetOsMaximumVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPasswordBlockSimple(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumCharacterSetCount(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeLock(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredToUnlockFromIdle(value *bool)()
    SetPasswordRequiredType(value *RequiredPasswordType)()
    SetRequireHealthyDeviceReport(value *bool)()
    SetSecureBootEnabled(value *bool)()
    SetStorageRequireEncryption(value *bool)()
}
