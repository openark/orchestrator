package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntegerRangeable 
type IntegerRangeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnd()(*int64)
    GetOdataType()(*string)
    GetStart()(*int64)
    SetEnd(value *int64)()
    SetOdataType(value *string)()
    SetStart(value *int64)()
}
