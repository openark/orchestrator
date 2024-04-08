package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityApiConnectorable 
type IdentityApiConnectorable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationConfiguration()(ApiAuthenticationConfigurationBaseable)
    GetDisplayName()(*string)
    GetTargetUrl()(*string)
    SetAuthenticationConfiguration(value ApiAuthenticationConfigurationBaseable)()
    SetDisplayName(value *string)()
    SetTargetUrl(value *string)()
}
