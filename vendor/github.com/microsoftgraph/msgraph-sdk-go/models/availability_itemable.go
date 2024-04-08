package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AvailabilityItemable 
type AvailabilityItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDateTime()(DateTimeTimeZoneable)
    GetOdataType()(*string)
    GetServiceId()(*string)
    GetStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*BookingsAvailabilityStatus)
    SetEndDateTime(value DateTimeTimeZoneable)()
    SetOdataType(value *string)()
    SetServiceId(value *string)()
    SetStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *BookingsAvailabilityStatus)()
}
