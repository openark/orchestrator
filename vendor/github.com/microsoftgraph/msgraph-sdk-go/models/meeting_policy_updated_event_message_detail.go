package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingPolicyUpdatedEventMessageDetail 
type MeetingPolicyUpdatedEventMessageDetail struct {
    EventMessageDetail
    // Initiator of the event.
    initiator IdentitySetable
    // Represents whether the meeting chat is enabled or not.
    meetingChatEnabled *bool
    // Unique identifier of the meeting chat.
    meetingChatId *string
}
// NewMeetingPolicyUpdatedEventMessageDetail instantiates a new MeetingPolicyUpdatedEventMessageDetail and sets the default values.
func NewMeetingPolicyUpdatedEventMessageDetail()(*MeetingPolicyUpdatedEventMessageDetail) {
    m := &MeetingPolicyUpdatedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.meetingPolicyUpdatedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMeetingPolicyUpdatedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingPolicyUpdatedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingPolicyUpdatedEventMessageDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingPolicyUpdatedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["initiator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
        }
        return nil
    }
    res["meetingChatEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingChatEnabled(val)
        }
        return nil
    }
    res["meetingChatId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingChatId(val)
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *MeetingPolicyUpdatedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// GetMeetingChatEnabled gets the meetingChatEnabled property value. Represents whether the meeting chat is enabled or not.
func (m *MeetingPolicyUpdatedEventMessageDetail) GetMeetingChatEnabled()(*bool) {
    return m.meetingChatEnabled
}
// GetMeetingChatId gets the meetingChatId property value. Unique identifier of the meeting chat.
func (m *MeetingPolicyUpdatedEventMessageDetail) GetMeetingChatId()(*string) {
    return m.meetingChatId
}
// Serialize serializes information the current object
func (m *MeetingPolicyUpdatedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("meetingChatEnabled", m.GetMeetingChatEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("meetingChatId", m.GetMeetingChatId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *MeetingPolicyUpdatedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
// SetMeetingChatEnabled sets the meetingChatEnabled property value. Represents whether the meeting chat is enabled or not.
func (m *MeetingPolicyUpdatedEventMessageDetail) SetMeetingChatEnabled(value *bool)() {
    m.meetingChatEnabled = value
}
// SetMeetingChatId sets the meetingChatId property value. Unique identifier of the meeting chat.
func (m *MeetingPolicyUpdatedEventMessageDetail) SetMeetingChatId(value *string)() {
    m.meetingChatId = value
}
