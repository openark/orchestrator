package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodsPolicyable 
type AuthenticationMethodsPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethodConfigurations()([]AuthenticationMethodConfigurationable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPolicyVersion()(*string)
    GetReconfirmationInDays()(*int32)
    GetRegistrationEnforcement()(RegistrationEnforcementable)
    SetAuthenticationMethodConfigurations(value []AuthenticationMethodConfigurationable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPolicyVersion(value *string)()
    SetReconfirmationInDays(value *int32)()
    SetRegistrationEnforcement(value RegistrationEnforcementable)()
}
