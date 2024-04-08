package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Authenticationable 
type Authenticationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailMethods()([]EmailAuthenticationMethodable)
    GetFido2Methods()([]Fido2AuthenticationMethodable)
    GetMethods()([]AuthenticationMethodable)
    GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethodable)
    GetOperations()([]LongRunningOperationable)
    GetPasswordMethods()([]PasswordAuthenticationMethodable)
    GetPhoneMethods()([]PhoneAuthenticationMethodable)
    GetSoftwareOathMethods()([]SoftwareOathAuthenticationMethodable)
    GetTemporaryAccessPassMethods()([]TemporaryAccessPassAuthenticationMethodable)
    GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethodable)
    SetEmailMethods(value []EmailAuthenticationMethodable)()
    SetFido2Methods(value []Fido2AuthenticationMethodable)()
    SetMethods(value []AuthenticationMethodable)()
    SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethodable)()
    SetOperations(value []LongRunningOperationable)()
    SetPasswordMethods(value []PasswordAuthenticationMethodable)()
    SetPhoneMethods(value []PhoneAuthenticationMethodable)()
    SetSoftwareOathMethods(value []SoftwareOathAuthenticationMethodable)()
    SetTemporaryAccessPassMethods(value []TemporaryAccessPassAuthenticationMethodable)()
    SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethodable)()
}
