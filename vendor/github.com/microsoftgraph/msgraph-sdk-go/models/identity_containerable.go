package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityContainerable 
type IdentityContainerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiConnectors()([]IdentityApiConnectorable)
    GetB2xUserFlows()([]B2xIdentityUserFlowable)
    GetConditionalAccess()(ConditionalAccessRootable)
    GetIdentityProviders()([]IdentityProviderBaseable)
    GetUserFlowAttributes()([]IdentityUserFlowAttributeable)
    SetApiConnectors(value []IdentityApiConnectorable)()
    SetB2xUserFlows(value []B2xIdentityUserFlowable)()
    SetConditionalAccess(value ConditionalAccessRootable)()
    SetIdentityProviders(value []IdentityProviderBaseable)()
    SetUserFlowAttributes(value []IdentityUserFlowAttributeable)()
}
