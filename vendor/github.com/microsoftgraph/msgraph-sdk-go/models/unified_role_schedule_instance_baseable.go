package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleScheduleInstanceBaseable 
type UnifiedRoleScheduleInstanceBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppScope()(AppScopeable)
    GetAppScopeId()(*string)
    GetDirectoryScope()(DirectoryObjectable)
    GetDirectoryScopeId()(*string)
    GetPrincipal()(DirectoryObjectable)
    GetPrincipalId()(*string)
    GetRoleDefinition()(UnifiedRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    SetAppScope(value AppScopeable)()
    SetAppScopeId(value *string)()
    SetDirectoryScope(value DirectoryObjectable)()
    SetDirectoryScopeId(value *string)()
    SetPrincipal(value DirectoryObjectable)()
    SetPrincipalId(value *string)()
    SetRoleDefinition(value UnifiedRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
}
