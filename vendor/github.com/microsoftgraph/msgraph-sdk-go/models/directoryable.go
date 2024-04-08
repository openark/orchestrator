package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Directoryable 
type Directoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnits()([]AdministrativeUnitable)
    GetDeletedItems()([]DirectoryObjectable)
    GetFederationConfigurations()([]IdentityProviderBaseable)
    SetAdministrativeUnits(value []AdministrativeUnitable)()
    SetDeletedItems(value []DirectoryObjectable)()
    SetFederationConfigurations(value []IdentityProviderBaseable)()
}
