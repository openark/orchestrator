package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConversationThread 
type ConversationThread struct {
    Entity
    // The Cc: recipients for the thread. Returned only on $select.
    ccRecipients []Recipientable
    // Indicates whether any of the posts within this thread has at least one attachment. Returned by default.
    hasAttachments *bool
    // Indicates if the thread is locked. Returned by default.
    isLocked *bool
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default.
    lastDeliveredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The posts property
    posts []Postable
    // A short summary from the body of the latest post in this conversation. Returned by default.
    preview *string
    // The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. Returned by default.
    topic *string
    // The To: recipients for the thread. Returned only on $select.
    toRecipients []Recipientable
    // All the users that sent a message to this thread. Returned by default.
    uniqueSenders []string
}
// NewConversationThread instantiates a new conversationThread and sets the default values.
func NewConversationThread()(*ConversationThread) {
    m := &ConversationThread{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConversationThreadFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConversationThreadFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConversationThread(), nil
}
// GetCcRecipients gets the ccRecipients property value. The Cc: recipients for the thread. Returned only on $select.
func (m *ConversationThread) GetCcRecipients()([]Recipientable) {
    return m.ccRecipients
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConversationThread) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ccRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetCcRecipients(res)
        }
        return nil
    }
    res["hasAttachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAttachments(val)
        }
        return nil
    }
    res["isLocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocked(val)
        }
        return nil
    }
    res["lastDeliveredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDeliveredDateTime(val)
        }
        return nil
    }
    res["posts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Postable, len(val))
            for i, v := range val {
                res[i] = v.(Postable)
            }
            m.SetPosts(res)
        }
        return nil
    }
    res["preview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreview(val)
        }
        return nil
    }
    res["topic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val)
        }
        return nil
    }
    res["toRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetToRecipients(res)
        }
        return nil
    }
    res["uniqueSenders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUniqueSenders(res)
        }
        return nil
    }
    return res
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether any of the posts within this thread has at least one attachment. Returned by default.
func (m *ConversationThread) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetIsLocked gets the isLocked property value. Indicates if the thread is locked. Returned by default.
func (m *ConversationThread) GetIsLocked()(*bool) {
    return m.isLocked
}
// GetLastDeliveredDateTime gets the lastDeliveredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default.
func (m *ConversationThread) GetLastDeliveredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastDeliveredDateTime
}
// GetPosts gets the posts property value. The posts property
func (m *ConversationThread) GetPosts()([]Postable) {
    return m.posts
}
// GetPreview gets the preview property value. A short summary from the body of the latest post in this conversation. Returned by default.
func (m *ConversationThread) GetPreview()(*string) {
    return m.preview
}
// GetTopic gets the topic property value. The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. Returned by default.
func (m *ConversationThread) GetTopic()(*string) {
    return m.topic
}
// GetToRecipients gets the toRecipients property value. The To: recipients for the thread. Returned only on $select.
func (m *ConversationThread) GetToRecipients()([]Recipientable) {
    return m.toRecipients
}
// GetUniqueSenders gets the uniqueSenders property value. All the users that sent a message to this thread. Returned by default.
func (m *ConversationThread) GetUniqueSenders()([]string) {
    return m.uniqueSenders
}
// Serialize serializes information the current object
func (m *ConversationThread) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCcRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCcRecipients()))
        for i, v := range m.GetCcRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("ccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocked", m.GetIsLocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastDeliveredDateTime", m.GetLastDeliveredDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPosts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPosts()))
        for i, v := range m.GetPosts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("posts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preview", m.GetPreview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("topic", m.GetTopic())
        if err != nil {
            return err
        }
    }
    if m.GetToRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetToRecipients()))
        for i, v := range m.GetToRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("toRecipients", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUniqueSenders() != nil {
        err = writer.WriteCollectionOfStringValues("uniqueSenders", m.GetUniqueSenders())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCcRecipients sets the ccRecipients property value. The Cc: recipients for the thread. Returned only on $select.
func (m *ConversationThread) SetCcRecipients(value []Recipientable)() {
    m.ccRecipients = value
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether any of the posts within this thread has at least one attachment. Returned by default.
func (m *ConversationThread) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetIsLocked sets the isLocked property value. Indicates if the thread is locked. Returned by default.
func (m *ConversationThread) SetIsLocked(value *bool)() {
    m.isLocked = value
}
// SetLastDeliveredDateTime sets the lastDeliveredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default.
func (m *ConversationThread) SetLastDeliveredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDeliveredDateTime = value
}
// SetPosts sets the posts property value. The posts property
func (m *ConversationThread) SetPosts(value []Postable)() {
    m.posts = value
}
// SetPreview sets the preview property value. A short summary from the body of the latest post in this conversation. Returned by default.
func (m *ConversationThread) SetPreview(value *string)() {
    m.preview = value
}
// SetTopic sets the topic property value. The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. Returned by default.
func (m *ConversationThread) SetTopic(value *string)() {
    m.topic = value
}
// SetToRecipients sets the toRecipients property value. The To: recipients for the thread. Returned only on $select.
func (m *ConversationThread) SetToRecipients(value []Recipientable)() {
    m.toRecipients = value
}
// SetUniqueSenders sets the uniqueSenders property value. All the users that sent a message to this thread. Returned by default.
func (m *ConversationThread) SetUniqueSenders(value []string)() {
    m.uniqueSenders = value
}
