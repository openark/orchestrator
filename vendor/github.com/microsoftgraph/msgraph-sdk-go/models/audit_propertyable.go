package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditPropertyable 
type AuditPropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetNewValue()(*string)
    GetOdataType()(*string)
    GetOldValue()(*string)
    SetDisplayName(value *string)()
    SetNewValue(value *string)()
    SetOdataType(value *string)()
    SetOldValue(value *string)()
}
