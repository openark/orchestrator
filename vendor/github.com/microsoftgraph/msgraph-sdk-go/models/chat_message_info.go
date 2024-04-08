package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageInfo 
type ChatMessageInfo struct {
    Entity
    // Body of the chatMessage. This will still contain markers for @mentions and attachments even though the object does not return @mentions and attachments.
    body ItemBodyable
    // Date time object representing the time at which message was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
    eventDetail EventMessageDetailable
    // Information about the sender of the message.
    from ChatMessageFromIdentitySetable
    // If set to true, the original message has been deleted.
    isDeleted *bool
    // The messageType property
    messageType *ChatMessageType
}
// NewChatMessageInfo instantiates a new chatMessageInfo and sets the default values.
func NewChatMessageInfo()(*ChatMessageInfo) {
    m := &ChatMessageInfo{
        Entity: *NewEntity(),
    }
    return m
}
// CreateChatMessageInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageInfo(), nil
}
// GetBody gets the body property value. Body of the chatMessage. This will still contain markers for @mentions and attachments even though the object does not return @mentions and attachments.
func (m *ChatMessageInfo) GetBody()(ItemBodyable) {
    return m.body
}
// GetCreatedDateTime gets the createdDateTime property value. Date time object representing the time at which message was created.
func (m *ChatMessageInfo) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetEventDetail gets the eventDetail property value. Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
func (m *ChatMessageInfo) GetEventDetail()(EventMessageDetailable) {
    return m.eventDetail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(ItemBodyable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["eventDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEventMessageDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDetail(val.(EventMessageDetailable))
        }
        return nil
    }
    res["from"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessageFromIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val.(ChatMessageFromIdentitySetable))
        }
        return nil
    }
    res["isDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["messageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageType(val.(*ChatMessageType))
        }
        return nil
    }
    return res
}
// GetFrom gets the from property value. Information about the sender of the message.
func (m *ChatMessageInfo) GetFrom()(ChatMessageFromIdentitySetable) {
    return m.from
}
// GetIsDeleted gets the isDeleted property value. If set to true, the original message has been deleted.
func (m *ChatMessageInfo) GetIsDeleted()(*bool) {
    return m.isDeleted
}
// GetMessageType gets the messageType property value. The messageType property
func (m *ChatMessageInfo) GetMessageType()(*ChatMessageType) {
    return m.messageType
}
// Serialize serializes information the current object
func (m *ChatMessageInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("eventDetail", m.GetEventDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    if m.GetMessageType() != nil {
        cast := (*m.GetMessageType()).String()
        err = writer.WriteStringValue("messageType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBody sets the body property value. Body of the chatMessage. This will still contain markers for @mentions and attachments even though the object does not return @mentions and attachments.
func (m *ChatMessageInfo) SetBody(value ItemBodyable)() {
    m.body = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date time object representing the time at which message was created.
func (m *ChatMessageInfo) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetEventDetail sets the eventDetail property value. Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
func (m *ChatMessageInfo) SetEventDetail(value EventMessageDetailable)() {
    m.eventDetail = value
}
// SetFrom sets the from property value. Information about the sender of the message.
func (m *ChatMessageInfo) SetFrom(value ChatMessageFromIdentitySetable)() {
    m.from = value
}
// SetIsDeleted sets the isDeleted property value. If set to true, the original message has been deleted.
func (m *ChatMessageInfo) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// SetMessageType sets the messageType property value. The messageType property
func (m *ChatMessageInfo) SetMessageType(value *ChatMessageType)() {
    m.messageType = value
}
