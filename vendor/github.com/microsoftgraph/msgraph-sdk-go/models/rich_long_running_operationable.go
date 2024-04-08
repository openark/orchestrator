package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RichLongRunningOperationable 
type RichLongRunningOperationable interface {
    LongRunningOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(PublicErrorable)
    GetPercentageComplete()(*int32)
    GetResourceId()(*string)
    GetType()(*string)
    SetError(value PublicErrorable)()
    SetPercentageComplete(value *int32)()
    SetResourceId(value *string)()
    SetType(value *string)()
}
