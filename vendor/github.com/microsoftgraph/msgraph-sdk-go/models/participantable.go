package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Participantable 
type Participantable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInfo()(ParticipantInfoable)
    GetIsInLobby()(*bool)
    GetIsMuted()(*bool)
    GetMediaStreams()([]MediaStreamable)
    GetMetadata()(*string)
    GetRecordingInfo()(RecordingInfoable)
    SetInfo(value ParticipantInfoable)()
    SetIsInLobby(value *bool)()
    SetIsMuted(value *bool)()
    SetMediaStreams(value []MediaStreamable)()
    SetMetadata(value *string)()
    SetRecordingInfo(value RecordingInfoable)()
}
