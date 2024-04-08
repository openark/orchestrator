package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommsOperationable 
type CommsOperationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientContext()(*string)
    GetResultInfo()(ResultInfoable)
    GetStatus()(*OperationStatus)
    SetClientContext(value *string)()
    SetResultInfo(value ResultInfoable)()
    SetStatus(value *OperationStatus)()
}
