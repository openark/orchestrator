package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAccountUserConversationMember 
type MicrosoftAccountUserConversationMember struct {
    ConversationMember
    // Microsoft Account ID of the user.
    userId *string
}
// NewMicrosoftAccountUserConversationMember instantiates a new MicrosoftAccountUserConversationMember and sets the default values.
func NewMicrosoftAccountUserConversationMember()(*MicrosoftAccountUserConversationMember) {
    m := &MicrosoftAccountUserConversationMember{
        ConversationMember: *NewConversationMember(),
    }
    odataTypeValue := "#microsoft.graph.microsoftAccountUserConversationMember"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftAccountUserConversationMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftAccountUserConversationMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAccountUserConversationMember(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftAccountUserConversationMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConversationMember.GetFieldDeserializers()
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetUserId gets the userId property value. Microsoft Account ID of the user.
func (m *MicrosoftAccountUserConversationMember) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *MicrosoftAccountUserConversationMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConversationMember.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserId sets the userId property value. Microsoft Account ID of the user.
func (m *MicrosoftAccountUserConversationMember) SetUserId(value *string)() {
    m.userId = value
}
