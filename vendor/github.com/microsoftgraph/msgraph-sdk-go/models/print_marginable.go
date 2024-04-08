package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintMarginable 
type PrintMarginable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBottom()(*int32)
    GetLeft()(*int32)
    GetOdataType()(*string)
    GetRight()(*int32)
    GetTop()(*int32)
    SetBottom(value *int32)()
    SetLeft(value *int32)()
    SetOdataType(value *string)()
    SetRight(value *int32)()
    SetTop(value *int32)()
}
