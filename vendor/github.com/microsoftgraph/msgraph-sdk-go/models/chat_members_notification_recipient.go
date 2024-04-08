package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMembersNotificationRecipient 
type ChatMembersNotificationRecipient struct {
    TeamworkNotificationRecipient
    // The unique identifier for the chat whose members should receive the notifications.
    chatId *string
}
// NewChatMembersNotificationRecipient instantiates a new ChatMembersNotificationRecipient and sets the default values.
func NewChatMembersNotificationRecipient()(*ChatMembersNotificationRecipient) {
    m := &ChatMembersNotificationRecipient{
        TeamworkNotificationRecipient: *NewTeamworkNotificationRecipient(),
    }
    odataTypeValue := "#microsoft.graph.chatMembersNotificationRecipient"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateChatMembersNotificationRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMembersNotificationRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMembersNotificationRecipient(), nil
}
// GetChatId gets the chatId property value. The unique identifier for the chat whose members should receive the notifications.
func (m *ChatMembersNotificationRecipient) GetChatId()(*string) {
    return m.chatId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMembersNotificationRecipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamworkNotificationRecipient.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *ChatMembersNotificationRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamworkNotificationRecipient.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("chatId", m.GetChatId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChatId sets the chatId property value. The unique identifier for the chat whose members should receive the notifications.
func (m *ChatMembersNotificationRecipient) SetChatId(value *string)() {
    m.chatId = value
}
