package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteOperationable 
type OnenoteOperationable interface {
    Operationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(OnenoteOperationErrorable)
    GetPercentComplete()(*string)
    GetResourceId()(*string)
    GetResourceLocation()(*string)
    SetError(value OnenoteOperationErrorable)()
    SetPercentComplete(value *string)()
    SetResourceId(value *string)()
    SetResourceLocation(value *string)()
}
