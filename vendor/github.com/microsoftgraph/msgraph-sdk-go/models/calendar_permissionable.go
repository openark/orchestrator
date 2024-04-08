package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarPermissionable 
type CalendarPermissionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedRoles()([]CalendarRoleType)
    GetEmailAddress()(EmailAddressable)
    GetIsInsideOrganization()(*bool)
    GetIsRemovable()(*bool)
    GetRole()(*CalendarRoleType)
    SetAllowedRoles(value []CalendarRoleType)()
    SetEmailAddress(value EmailAddressable)()
    SetIsInsideOrganization(value *bool)()
    SetIsRemovable(value *bool)()
    SetRole(value *CalendarRoleType)()
}
