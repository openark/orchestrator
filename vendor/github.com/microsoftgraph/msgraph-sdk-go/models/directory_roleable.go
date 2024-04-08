package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryRoleable 
type DirectoryRoleable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetMembers()([]DirectoryObjectable)
    GetRoleTemplateId()(*string)
    GetScopedMembers()([]ScopedRoleMembershipable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetMembers(value []DirectoryObjectable)()
    SetRoleTemplateId(value *string)()
    SetScopedMembers(value []ScopedRoleMembershipable)()
}
