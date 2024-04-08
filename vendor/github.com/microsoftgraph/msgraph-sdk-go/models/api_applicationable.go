package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiApplicationable 
type ApiApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptMappedClaims()(*bool)
    GetKnownClientApplications()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetOauth2PermissionScopes()([]PermissionScopeable)
    GetOdataType()(*string)
    GetPreAuthorizedApplications()([]PreAuthorizedApplicationable)
    GetRequestedAccessTokenVersion()(*int32)
    SetAcceptMappedClaims(value *bool)()
    SetKnownClientApplications(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetOauth2PermissionScopes(value []PermissionScopeable)()
    SetOdataType(value *string)()
    SetPreAuthorizedApplications(value []PreAuthorizedApplicationable)()
    SetRequestedAccessTokenVersion(value *int32)()
}
