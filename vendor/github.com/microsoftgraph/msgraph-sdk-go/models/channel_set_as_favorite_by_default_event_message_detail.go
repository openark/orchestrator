package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelSetAsFavoriteByDefaultEventMessageDetail 
type ChannelSetAsFavoriteByDefaultEventMessageDetail struct {
    EventMessageDetail
    // Unique identifier of the channel.
    channelId *string
    // Initiator of the event.
    initiator IdentitySetable
}
// NewChannelSetAsFavoriteByDefaultEventMessageDetail instantiates a new ChannelSetAsFavoriteByDefaultEventMessageDetail and sets the default values.
func NewChannelSetAsFavoriteByDefaultEventMessageDetail()(*ChannelSetAsFavoriteByDefaultEventMessageDetail) {
    m := &ChannelSetAsFavoriteByDefaultEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.channelSetAsFavoriteByDefaultEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateChannelSetAsFavoriteByDefaultEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelSetAsFavoriteByDefaultEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelSetAsFavoriteByDefaultEventMessageDetail(), nil
}
// GetChannelId gets the channelId property value. Unique identifier of the channel.
func (m *ChannelSetAsFavoriteByDefaultEventMessageDetail) GetChannelId()(*string) {
    return m.channelId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelSetAsFavoriteByDefaultEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
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
func (m *ChannelSetAsFavoriteByDefaultEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// Serialize serializes information the current object
func (m *ChannelSetAsFavoriteByDefaultEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
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
// SetChannelId sets the channelId property value. Unique identifier of the channel.
func (m *ChannelSetAsFavoriteByDefaultEventMessageDetail) SetChannelId(value *string)() {
    m.channelId = value
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *ChannelSetAsFavoriteByDefaultEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
