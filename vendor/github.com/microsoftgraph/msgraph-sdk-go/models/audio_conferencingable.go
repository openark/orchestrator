package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AudioConferencingable 
type AudioConferencingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConferenceId()(*string)
    GetDialinUrl()(*string)
    GetOdataType()(*string)
    GetTollFreeNumber()(*string)
    GetTollFreeNumbers()([]string)
    GetTollNumber()(*string)
    GetTollNumbers()([]string)
    SetConferenceId(value *string)()
    SetDialinUrl(value *string)()
    SetOdataType(value *string)()
    SetTollFreeNumber(value *string)()
    SetTollFreeNumbers(value []string)()
    SetTollNumber(value *string)()
    SetTollNumbers(value []string)()
}
