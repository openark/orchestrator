package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementSettingsable 
type EntitlementManagementSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDurationUntilExternalUserDeletedAfterBlocked()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetExternalUserLifecycleAction()(*AccessPackageExternalUserLifecycleAction)
    SetDurationUntilExternalUserDeletedAfterBlocked(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetExternalUserLifecycleAction(value *AccessPackageExternalUserLifecycleAction)()
}
