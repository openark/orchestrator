package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Participant 
type Participant struct {
    Entity
    // The info property
    info ParticipantInfoable
    // true if the participant is in lobby.
    isInLobby *bool
    // true if the participant is muted (client or server muted).
    isMuted *bool
    // The list of media streams.
    mediaStreams []MediaStreamable
    // A blob of data provided by the participant in the roster.
    metadata *string
    // Information about whether the participant has recording capability.
    recordingInfo RecordingInfoable
}
// NewParticipant instantiates a new participant and sets the default values.
func NewParticipant()(*Participant) {
    m := &Participant{
        Entity: *NewEntity(),
    }
    return m
}
// CreateParticipantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParticipantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipant(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Participant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["info"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfo(val.(ParticipantInfoable))
        }
        return nil
    }
    res["isInLobby"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInLobby(val)
        }
        return nil
    }
    res["isMuted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMuted(val)
        }
        return nil
    }
    res["mediaStreams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMediaStreamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MediaStreamable, len(val))
            for i, v := range val {
                res[i] = v.(MediaStreamable)
            }
            m.SetMediaStreams(res)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val)
        }
        return nil
    }
    res["recordingInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecordingInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordingInfo(val.(RecordingInfoable))
        }
        return nil
    }
    return res
}
// GetInfo gets the info property value. The info property
func (m *Participant) GetInfo()(ParticipantInfoable) {
    return m.info
}
// GetIsInLobby gets the isInLobby property value. true if the participant is in lobby.
func (m *Participant) GetIsInLobby()(*bool) {
    return m.isInLobby
}
// GetIsMuted gets the isMuted property value. true if the participant is muted (client or server muted).
func (m *Participant) GetIsMuted()(*bool) {
    return m.isMuted
}
// GetMediaStreams gets the mediaStreams property value. The list of media streams.
func (m *Participant) GetMediaStreams()([]MediaStreamable) {
    return m.mediaStreams
}
// GetMetadata gets the metadata property value. A blob of data provided by the participant in the roster.
func (m *Participant) GetMetadata()(*string) {
    return m.metadata
}
// GetRecordingInfo gets the recordingInfo property value. Information about whether the participant has recording capability.
func (m *Participant) GetRecordingInfo()(RecordingInfoable) {
    return m.recordingInfo
}
// Serialize serializes information the current object
func (m *Participant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInLobby", m.GetIsInLobby())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMuted", m.GetIsMuted())
        if err != nil {
            return err
        }
    }
    if m.GetMediaStreams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMediaStreams()))
        for i, v := range m.GetMediaStreams() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mediaStreams", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recordingInfo", m.GetRecordingInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInfo sets the info property value. The info property
func (m *Participant) SetInfo(value ParticipantInfoable)() {
    m.info = value
}
// SetIsInLobby sets the isInLobby property value. true if the participant is in lobby.
func (m *Participant) SetIsInLobby(value *bool)() {
    m.isInLobby = value
}
// SetIsMuted sets the isMuted property value. true if the participant is muted (client or server muted).
func (m *Participant) SetIsMuted(value *bool)() {
    m.isMuted = value
}
// SetMediaStreams sets the mediaStreams property value. The list of media streams.
func (m *Participant) SetMediaStreams(value []MediaStreamable)() {
    m.mediaStreams = value
}
// SetMetadata sets the metadata property value. A blob of data provided by the participant in the roster.
func (m *Participant) SetMetadata(value *string)() {
    m.metadata = value
}
// SetRecordingInfo sets the recordingInfo property value. Information about whether the participant has recording capability.
func (m *Participant) SetRecordingInfo(value RecordingInfoable)() {
    m.recordingInfo = value
}
