package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallRecordingEventMessageDetailable 
type CallRecordingEventMessageDetailable interface {
    EventMessageDetailable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallId()(*string)
    GetCallRecordingDisplayName()(*string)
    GetCallRecordingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetCallRecordingStatus()(*CallRecordingStatus)
    GetCallRecordingUrl()(*string)
    GetInitiator()(IdentitySetable)
    GetMeetingOrganizer()(IdentitySetable)
    SetCallId(value *string)()
    SetCallRecordingDisplayName(value *string)()
    SetCallRecordingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetCallRecordingStatus(value *CallRecordingStatus)()
    SetCallRecordingUrl(value *string)()
    SetInitiator(value IdentitySetable)()
    SetMeetingOrganizer(value IdentitySetable)()
}
