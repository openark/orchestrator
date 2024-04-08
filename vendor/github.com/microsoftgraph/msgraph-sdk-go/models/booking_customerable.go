package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomerable 
type BookingCustomerable interface {
    BookingCustomerBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddresses()([]PhysicalAddressable)
    GetDisplayName()(*string)
    GetEmailAddress()(*string)
    GetPhones()([]Phoneable)
    SetAddresses(value []PhysicalAddressable)()
    SetDisplayName(value *string)()
    SetEmailAddress(value *string)()
    SetPhones(value []Phoneable)()
}
