package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppable 
type TeamsAppable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDefinitions()([]TeamsAppDefinitionable)
    GetDisplayName()(*string)
    GetDistributionMethod()(*TeamsAppDistributionMethod)
    GetExternalId()(*string)
    SetAppDefinitions(value []TeamsAppDefinitionable)()
    SetDisplayName(value *string)()
    SetDistributionMethod(value *TeamsAppDistributionMethod)()
    SetExternalId(value *string)()
}
