package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncementAttachment 
type ServiceAnnouncementAttachment struct {
    Entity
    // The attachment content.
    content []byte
    // The contentType property
    contentType *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name property
    name *string
    // The size property
    size *int32
}
// NewServiceAnnouncementAttachment instantiates a new serviceAnnouncementAttachment and sets the default values.
func NewServiceAnnouncementAttachment()(*ServiceAnnouncementAttachment) {
    m := &ServiceAnnouncementAttachment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceAnnouncementAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncementAttachment(), nil
}
// GetContent gets the content property value. The attachment content.
func (m *ServiceAnnouncementAttachment) GetContent()([]byte) {
    return m.content
}
// GetContentType gets the contentType property value. The contentType property
func (m *ServiceAnnouncementAttachment) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAnnouncementAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ServiceAnnouncementAttachment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetName gets the name property value. The name property
func (m *ServiceAnnouncementAttachment) GetName()(*string) {
    return m.name
}
// GetSize gets the size property value. The size property
func (m *ServiceAnnouncementAttachment) GetSize()(*int32) {
    return m.size
}
// Serialize serializes information the current object
func (m *ServiceAnnouncementAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The attachment content.
func (m *ServiceAnnouncementAttachment) SetContent(value []byte)() {
    m.content = value
}
// SetContentType sets the contentType property value. The contentType property
func (m *ServiceAnnouncementAttachment) SetContentType(value *string)() {
    m.contentType = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ServiceAnnouncementAttachment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetName sets the name property value. The name property
func (m *ServiceAnnouncementAttachment) SetName(value *string)() {
    m.name = value
}
// SetSize sets the size property value. The size property
func (m *ServiceAnnouncementAttachment) SetSize(value *int32)() {
    m.size = value
}
