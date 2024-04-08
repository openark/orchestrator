package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultUserRolePermissionsable 
type DefaultUserRolePermissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedToCreateApps()(*bool)
    GetAllowedToCreateSecurityGroups()(*bool)
    GetAllowedToReadOtherUsers()(*bool)
    GetOdataType()(*string)
    GetPermissionGrantPoliciesAssigned()([]string)
    SetAllowedToCreateApps(value *bool)()
    SetAllowedToCreateSecurityGroups(value *bool)()
    SetAllowedToReadOtherUsers(value *bool)()
    SetOdataType(value *string)()
    SetPermissionGrantPoliciesAssigned(value []string)()
}
