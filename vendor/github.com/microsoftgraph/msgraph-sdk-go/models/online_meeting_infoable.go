package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeetingInfoable 
type OnlineMeetingInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConferenceId()(*string)
    GetJoinUrl()(*string)
    GetOdataType()(*string)
    GetPhones()([]Phoneable)
    GetQuickDial()(*string)
    GetTollFreeNumbers()([]string)
    GetTollNumber()(*string)
    SetConferenceId(value *string)()
    SetJoinUrl(value *string)()
    SetOdataType(value *string)()
    SetPhones(value []Phoneable)()
    SetQuickDial(value *string)()
    SetTollFreeNumbers(value []string)()
    SetTollNumber(value *string)()
}
