package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelDescriptionUpdatedEventMessageDetail 
type ChannelDescriptionUpdatedEventMessageDetail struct {
    EventMessageDetail
    // The updated description of the channel.
    channelDescription *string
    // Unique identifier of the channel.
    channelId *string
    // Initiator of the event.
    initiator IdentitySetable
}
// NewChannelDescriptionUpdatedEventMessageDetail instantiates a new ChannelDescriptionUpdatedEventMessageDetail and sets the default values.
func NewChannelDescriptionUpdatedEventMessageDetail()(*ChannelDescriptionUpdatedEventMessageDetail) {
    m := &ChannelDescriptionUpdatedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.channelDescriptionUpdatedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateChannelDescriptionUpdatedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelDescriptionUpdatedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelDescriptionUpdatedEventMessageDetail(), nil
}
// GetChannelDescription gets the channelDescription property value. The updated description of the channel.
func (m *ChannelDescriptionUpdatedEventMessageDetail) GetChannelDescription()(*string) {
    return m.channelDescription
}
// GetChannelId gets the channelId property value. Unique identifier of the channel.
func (m *ChannelDescriptionUpdatedEventMessageDetail) GetChannelId()(*string) {
    return m.channelId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelDescriptionUpdatedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["channelDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelDescription(val)
        }
        return nil
    }
    res["channelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelId(val)
        }
        return nil
    }
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
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *ChannelDescriptionUpdatedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// Serialize serializes information the current object
func (m *ChannelDescriptionUpdatedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("channelDescription", m.GetChannelDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("channelId", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChannelDescription sets the channelDescription property value. The updated description of the channel.
func (m *ChannelDescriptionUpdatedEventMessageDetail) SetChannelDescription(value *string)() {
    m.channelDescription = value
}
// SetChannelId sets the channelId property value. Unique identifier of the channel.
func (m *ChannelDescriptionUpdatedEventMessageDetail) SetChannelId(value *string)() {
    m.channelId = value
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *ChannelDescriptionUpdatedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
