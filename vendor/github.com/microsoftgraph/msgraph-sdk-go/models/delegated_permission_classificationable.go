package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedPermissionClassificationable 
type DelegatedPermissionClassificationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassification()(*PermissionClassificationType)
    GetPermissionId()(*string)
    GetPermissionName()(*string)
    SetClassification(value *PermissionClassificationType)()
    SetPermissionId(value *string)()
    SetPermissionName(value *string)()
}
