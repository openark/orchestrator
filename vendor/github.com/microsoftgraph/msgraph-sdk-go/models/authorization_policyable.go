package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthorizationPolicyable 
type AuthorizationPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyBaseable
    GetAllowedToSignUpEmailBasedSubscriptions()(*bool)
    GetAllowedToUseSSPR()(*bool)
    GetAllowEmailVerifiedUsersToJoinOrganization()(*bool)
    GetAllowInvitesFrom()(*AllowInvitesFrom)
    GetBlockMsolPowerShell()(*bool)
    GetDefaultUserRolePermissions()(DefaultUserRolePermissionsable)
    GetGuestUserRoleId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetAllowedToSignUpEmailBasedSubscriptions(value *bool)()
    SetAllowedToUseSSPR(value *bool)()
    SetAllowEmailVerifiedUsersToJoinOrganization(value *bool)()
    SetAllowInvitesFrom(value *AllowInvitesFrom)()
    SetBlockMsolPowerShell(value *bool)()
    SetDefaultUserRolePermissions(value DefaultUserRolePermissionsable)()
    SetGuestUserRoleId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
