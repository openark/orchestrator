package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// MicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBodyable 
type MicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendees()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable)
    GetIsOrganizerOptional()(*bool)
    GetLocationConstraint()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LocationConstraintable)
    GetMaxCandidates()(*int32)
    GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMinimumAttendeePercentage()(*float64)
    GetReturnSuggestionReasons()(*bool)
    GetTimeConstraint()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeConstraintable)
    SetAttendees(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttendeeBaseable)()
    SetIsOrganizerOptional(value *bool)()
    SetLocationConstraint(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LocationConstraintable)()
    SetMaxCandidates(value *int32)()
    SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMinimumAttendeePercentage(value *float64)()
    SetReturnSuggestionReasons(value *bool)()
    SetTimeConstraint(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeConstraintable)()
}
