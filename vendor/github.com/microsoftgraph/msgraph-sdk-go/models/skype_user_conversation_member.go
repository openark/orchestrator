package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SkypeUserConversationMember 
type SkypeUserConversationMember struct {
    ConversationMember
    // Skype ID of the user.
    skypeId *string
}
// NewSkypeUserConversationMember instantiates a new SkypeUserConversationMember and sets the default values.
func NewSkypeUserConversationMember()(*SkypeUserConversationMember) {
    m := &SkypeUserConversationMember{
        ConversationMember: *NewConversationMember(),
    }
    odataTypeValue := "#microsoft.graph.skypeUserConversationMember"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSkypeUserConversationMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSkypeUserConversationMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSkypeUserConversationMember(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SkypeUserConversationMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConversationMember.GetFieldDeserializers()
    res["skypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeId(val)
        }
        return nil
    }
    return res
}
// GetSkypeId gets the skypeId property value. Skype ID of the user.
func (m *SkypeUserConversationMember) GetSkypeId()(*string) {
    return m.skypeId
}
// Serialize serializes information the current object
func (m *SkypeUserConversationMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConversationMember.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("skypeId", m.GetSkypeId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSkypeId sets the skypeId property value. Skype ID of the user.
func (m *SkypeUserConversationMember) SetSkypeId(value *string)() {
    m.skypeId = value
}
