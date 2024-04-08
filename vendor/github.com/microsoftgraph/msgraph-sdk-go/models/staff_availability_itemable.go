package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StaffAvailabilityItemable 
type StaffAvailabilityItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailabilityItems()([]AvailabilityItemable)
    GetOdataType()(*string)
    GetStaffId()(*string)
    SetAvailabilityItems(value []AvailabilityItemable)()
    SetOdataType(value *string)()
    SetStaffId(value *string)()
}
