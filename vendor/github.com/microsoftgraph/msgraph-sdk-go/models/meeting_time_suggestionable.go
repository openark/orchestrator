package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingTimeSuggestionable 
type MeetingTimeSuggestionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendeeAvailability()([]AttendeeAvailabilityable)
    GetConfidence()(*float64)
    GetLocations()([]Locationable)
    GetMeetingTimeSlot()(TimeSlotable)
    GetOdataType()(*string)
    GetOrder()(*int32)
    GetOrganizerAvailability()(*FreeBusyStatus)
    GetSuggestionReason()(*string)
    SetAttendeeAvailability(value []AttendeeAvailabilityable)()
    SetConfidence(value *float64)()
    SetLocations(value []Locationable)()
    SetMeetingTimeSlot(value TimeSlotable)()
    SetOdataType(value *string)()
    SetOrder(value *int32)()
    SetOrganizerAvailability(value *FreeBusyStatus)()
    SetSuggestionReason(value *string)()
}
