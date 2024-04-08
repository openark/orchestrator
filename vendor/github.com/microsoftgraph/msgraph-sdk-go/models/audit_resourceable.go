package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditResourceable 
type AuditResourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuditResourceType()(*string)
    GetDisplayName()(*string)
    GetModifiedProperties()([]AuditPropertyable)
    GetOdataType()(*string)
    GetResourceId()(*string)
    SetAuditResourceType(value *string)()
    SetDisplayName(value *string)()
    SetModifiedProperties(value []AuditPropertyable)()
    SetOdataType(value *string)()
    SetResourceId(value *string)()
}
