package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleDefinitionable 
type UnifiedRoleDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInheritsPermissionsFrom()([]UnifiedRoleDefinitionable)
    GetIsBuiltIn()(*bool)
    GetIsEnabled()(*bool)
    GetResourceScopes()([]string)
    GetRolePermissions()([]UnifiedRolePermissionable)
    GetTemplateId()(*string)
    GetVersion()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInheritsPermissionsFrom(value []UnifiedRoleDefinitionable)()
    SetIsBuiltIn(value *bool)()
    SetIsEnabled(value *bool)()
    SetResourceScopes(value []string)()
    SetRolePermissions(value []UnifiedRolePermissionable)()
    SetTemplateId(value *string)()
    SetVersion(value *string)()
}
