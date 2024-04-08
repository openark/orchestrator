package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonOrGroupColumnable 
type PersonOrGroupColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowMultipleSelection()(*bool)
    GetChooseFromType()(*string)
    GetDisplayAs()(*string)
    GetOdataType()(*string)
    SetAllowMultipleSelection(value *bool)()
    SetChooseFromType(value *string)()
    SetDisplayAs(value *string)()
    SetOdataType(value *string)()
}
