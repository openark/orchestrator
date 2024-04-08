package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyAssignmentable 
type UnifiedRoleManagementPolicyAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPolicy()(UnifiedRoleManagementPolicyable)
    GetPolicyId()(*string)
    GetRoleDefinitionId()(*string)
    GetScopeId()(*string)
    GetScopeType()(*string)
    SetPolicy(value UnifiedRoleManagementPolicyable)()
    SetPolicyId(value *string)()
    SetRoleDefinitionId(value *string)()
    SetScopeId(value *string)()
    SetScopeType(value *string)()
}
