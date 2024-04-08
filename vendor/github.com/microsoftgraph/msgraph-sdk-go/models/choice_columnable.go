package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChoiceColumnable 
type ChoiceColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowTextEntry()(*bool)
    GetChoices()([]string)
    GetDisplayAs()(*string)
    GetOdataType()(*string)
    SetAllowTextEntry(value *bool)()
    SetChoices(value []string)()
    SetDisplayAs(value *string)()
    SetOdataType(value *string)()
}
