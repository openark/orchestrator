package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StsPolicyable 
type StsPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyBaseable
    GetAppliesTo()([]DirectoryObjectable)
    GetDefinition()([]string)
    GetIsOrganizationDefault()(*bool)
    SetAppliesTo(value []DirectoryObjectable)()
    SetDefinition(value []string)()
    SetIsOrganizationDefault(value *bool)()
}
