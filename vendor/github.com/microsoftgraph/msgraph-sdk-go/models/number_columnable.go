package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NumberColumnable 
type NumberColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDecimalPlaces()(*string)
    GetDisplayAs()(*string)
    GetMaximum()(*float64)
    GetMinimum()(*float64)
    GetOdataType()(*string)
    SetDecimalPlaces(value *string)()
    SetDisplayAs(value *string)()
    SetMaximum(value *float64)()
    SetMinimum(value *float64)()
    SetOdataType(value *string)()
}
