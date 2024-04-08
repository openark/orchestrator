package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Phoneable 
type Phoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLanguage()(*string)
    GetNumber()(*string)
    GetOdataType()(*string)
    GetRegion()(*string)
    GetType()(*PhoneType)
    SetLanguage(value *string)()
    SetNumber(value *string)()
    SetOdataType(value *string)()
    SetRegion(value *string)()
    SetType(value *PhoneType)()
}
