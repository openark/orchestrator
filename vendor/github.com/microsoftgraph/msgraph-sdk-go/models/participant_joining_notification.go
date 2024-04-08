package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParticipantJoiningNotification 
type ParticipantJoiningNotification struct {
    Entity
    // The call property
    call Callable
}
// NewParticipantJoiningNotification instantiates a new ParticipantJoiningNotification and sets the default values.
func NewParticipantJoiningNotification()(*ParticipantJoiningNotification) {
    m := &ParticipantJoiningNotification{
        Entity: *NewEntity(),
    }
    return m
}
// CreateParticipantJoiningNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParticipantJoiningNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipantJoiningNotification(), nil
}
// GetCall gets the call property value. The call property
func (m *ParticipantJoiningNotification) GetCall()(Callable) {
    return m.call
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParticipantJoiningNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["call"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCallFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCall(val.(Callable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ParticipantJoiningNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("call", m.GetCall())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCall sets the call property value. The call property
func (m *ParticipantJoiningNotification) SetCall(value Callable)() {
    m.call = value
}
