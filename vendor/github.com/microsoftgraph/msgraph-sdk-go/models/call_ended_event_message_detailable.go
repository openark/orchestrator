package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallEndedEventMessageDetailable 
type CallEndedEventMessageDetailable interface {
    EventMessageDetailable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetCallEventType()(*TeamworkCallEventType)
    GetCallId()(*string)
    GetCallParticipants()([]CallParticipantInfoable)
    GetInitiator()(IdentitySetable)
    SetCallDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetCallEventType(value *TeamworkCallEventType)()
    SetCallId(value *string)()
    SetCallParticipants(value []CallParticipantInfoable)()
    SetInitiator(value IdentitySetable)()
}
