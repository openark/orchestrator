package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleDefinitionable 
type RoleDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsBuiltIn()(*bool)
    GetRoleAssignments()([]RoleAssignmentable)
    GetRolePermissions()([]RolePermissionable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsBuiltIn(value *bool)()
    SetRoleAssignments(value []RoleAssignmentable)()
    SetRolePermissions(value []RolePermissionable)()
}
