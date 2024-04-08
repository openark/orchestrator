package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentWindowsHelloForBusinessConfigurationable 
type DeviceEnrollmentWindowsHelloForBusinessConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnhancedBiometricsState()(*Enablement)
    GetPinExpirationInDays()(*int32)
    GetPinLowercaseCharactersUsage()(*WindowsHelloForBusinessPinUsage)
    GetPinMaximumLength()(*int32)
    GetPinMinimumLength()(*int32)
    GetPinPreviousBlockCount()(*int32)
    GetPinSpecialCharactersUsage()(*WindowsHelloForBusinessPinUsage)
    GetPinUppercaseCharactersUsage()(*WindowsHelloForBusinessPinUsage)
    GetRemotePassportEnabled()(*bool)
    GetSecurityDeviceRequired()(*bool)
    GetState()(*Enablement)
    GetUnlockWithBiometricsEnabled()(*bool)
    SetEnhancedBiometricsState(value *Enablement)()
    SetPinExpirationInDays(value *int32)()
    SetPinLowercaseCharactersUsage(value *WindowsHelloForBusinessPinUsage)()
    SetPinMaximumLength(value *int32)()
    SetPinMinimumLength(value *int32)()
    SetPinPreviousBlockCount(value *int32)()
    SetPinSpecialCharactersUsage(value *WindowsHelloForBusinessPinUsage)()
    SetPinUppercaseCharactersUsage(value *WindowsHelloForBusinessPinUsage)()
    SetRemotePassportEnabled(value *bool)()
    SetSecurityDeviceRequired(value *bool)()
    SetState(value *Enablement)()
    SetUnlockWithBiometricsEnabled(value *bool)()
}
