package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// B2xIdentityUserFlowable 
type B2xIdentityUserFlowable interface {
    IdentityUserFlowable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiConnectorConfiguration()(UserFlowApiConnectorConfigurationable)
    GetIdentityProviders()([]IdentityProviderable)
    GetLanguages()([]UserFlowLanguageConfigurationable)
    GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignmentable)
    GetUserFlowIdentityProviders()([]IdentityProviderBaseable)
    SetApiConnectorConfiguration(value UserFlowApiConnectorConfigurationable)()
    SetIdentityProviders(value []IdentityProviderable)()
    SetLanguages(value []UserFlowLanguageConfigurationable)()
    SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignmentable)()
    SetUserFlowIdentityProviders(value []IdentityProviderBaseable)()
}
