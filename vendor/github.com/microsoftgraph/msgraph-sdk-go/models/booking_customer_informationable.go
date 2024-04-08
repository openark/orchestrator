package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomerInformationable 
type BookingCustomerInformationable interface {
    BookingCustomerInformationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomerId()(*string)
    GetCustomQuestionAnswers()([]BookingQuestionAnswerable)
    GetEmailAddress()(*string)
    GetLocation()(Locationable)
    GetName()(*string)
    GetNotes()(*string)
    GetPhone()(*string)
    GetTimeZone()(*string)
    SetCustomerId(value *string)()
    SetCustomQuestionAnswers(value []BookingQuestionAnswerable)()
    SetEmailAddress(value *string)()
    SetLocation(value Locationable)()
    SetName(value *string)()
    SetNotes(value *string)()
    SetPhone(value *string)()
    SetTimeZone(value *string)()
}
