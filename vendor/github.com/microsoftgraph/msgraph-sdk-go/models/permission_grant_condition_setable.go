package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PermissionGrantConditionSetable 
type PermissionGrantConditionSetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientApplicationIds()([]string)
    GetClientApplicationPublisherIds()([]string)
    GetClientApplicationsFromVerifiedPublisherOnly()(*bool)
    GetClientApplicationTenantIds()([]string)
    GetPermissionClassification()(*string)
    GetPermissions()([]string)
    GetPermissionType()(*PermissionType)
    GetResourceApplication()(*string)
    SetClientApplicationIds(value []string)()
    SetClientApplicationPublisherIds(value []string)()
    SetClientApplicationsFromVerifiedPublisherOnly(value *bool)()
    SetClientApplicationTenantIds(value []string)()
    SetPermissionClassification(value *string)()
    SetPermissions(value []string)()
    SetPermissionType(value *PermissionType)()
    SetResourceApplication(value *string)()
}
