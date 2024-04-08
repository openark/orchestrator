package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PhoneAuthenticationMethodable 
type PhoneAuthenticationMethodable interface {
    AuthenticationMethodable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPhoneNumber()(*string)
    GetPhoneType()(*AuthenticationPhoneType)
    GetSmsSignInState()(*AuthenticationMethodSignInState)
    SetPhoneNumber(value *string)()
    SetPhoneType(value *AuthenticationPhoneType)()
    SetSmsSignInState(value *AuthenticationMethodSignInState)()
}
