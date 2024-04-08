package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatRenamedEventMessageDetail 
type ChatRenamedEventMessageDetail struct {
    EventMessageDetail
    // The updated name of the chat.
    chatDisplayName *string
    // Unique identifier of the chat.
    chatId *string
    // Initiator of the event.
    initiator IdentitySetable
}
// NewChatRenamedEventMessageDetail instantiates a new ChatRenamedEventMessageDetail and sets the default values.
func NewChatRenamedEventMessageDetail()(*ChatRenamedEventMessageDetail) {
    m := &ChatRenamedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.chatRenamedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateChatRenamedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatRenamedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatRenamedEventMessageDetail(), nil
}
// GetChatDisplayName gets the chatDisplayName property value. The updated name of the chat.
func (m *ChatRenamedEventMessageDetail) GetChatDisplayName()(*string) {
    return m.chatDisplayName
}
// GetChatId gets the chatId property value. Unique identifier of the chat.
func (m *ChatRenamedEventMessageDetail) GetChatId()(*string) {
    return m.chatId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatRenamedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["chatDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatDisplayName(val)
        }
        return nil
    }
    res["chatId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatId(val)
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
func (m *ChatRenamedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// Serialize serializes information the current object
func (m *ChatRenamedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("chatDisplayName", m.GetChatDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("chatId", m.GetChatId())
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
// SetChatDisplayName sets the chatDisplayName property value. The updated name of the chat.
func (m *ChatRenamedEventMessageDetail) SetChatDisplayName(value *string)() {
    m.chatDisplayName = value
}
// SetChatId sets the chatId property value. Unique identifier of the chat.
func (m *ChatRenamedEventMessageDetail) SetChatId(value *string)() {
    m.chatId = value
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *ChatRenamedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
