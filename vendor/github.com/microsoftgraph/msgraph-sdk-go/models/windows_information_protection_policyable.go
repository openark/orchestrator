package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionPolicyable 
type WindowsInformationProtectionPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsInformationProtectionable
    GetDaysWithoutContactBeforeUnenroll()(*int32)
    GetMdmEnrollmentUrl()(*string)
    GetMinutesOfInactivityBeforeDeviceLock()(*int32)
    GetNumberOfPastPinsRemembered()(*int32)
    GetPasswordMaximumAttemptCount()(*int32)
    GetPinExpirationDays()(*int32)
    GetPinLowercaseLetters()(*WindowsInformationProtectionPinCharacterRequirements)
    GetPinMinimumLength()(*int32)
    GetPinSpecialCharacters()(*WindowsInformationProtectionPinCharacterRequirements)
    GetPinUppercaseLetters()(*WindowsInformationProtectionPinCharacterRequirements)
    GetRevokeOnMdmHandoffDisabled()(*bool)
    GetWindowsHelloForBusinessBlocked()(*bool)
    SetDaysWithoutContactBeforeUnenroll(value *int32)()
    SetMdmEnrollmentUrl(value *string)()
    SetMinutesOfInactivityBeforeDeviceLock(value *int32)()
    SetNumberOfPastPinsRemembered(value *int32)()
    SetPasswordMaximumAttemptCount(value *int32)()
    SetPinExpirationDays(value *int32)()
    SetPinLowercaseLetters(value *WindowsInformationProtectionPinCharacterRequirements)()
    SetPinMinimumLength(value *int32)()
    SetPinSpecialCharacters(value *WindowsInformationProtectionPinCharacterRequirements)()
    SetPinUppercaseLetters(value *WindowsInformationProtectionPinCharacterRequirements)()
    SetRevokeOnMdmHandoffDisabled(value *bool)()
    SetWindowsHelloForBusinessBlocked(value *bool)()
}
