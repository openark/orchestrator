package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceUpdateMessage 
type ServiceUpdateMessage struct {
    ServiceAnnouncementBase
    // The expected deadline of the action for the message.
    actionRequiredByDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of serviceAnnouncementAttachments.
    attachments []ServiceAnnouncementAttachmentable
    // The zip file that contains all attachments for a message.
    attachmentsArchive []byte
    // The body property
    body ItemBodyable
    // The category property
    category *ServiceUpdateCategory
    // Indicates whether the message has any attachment.
    hasAttachments *bool
    // Indicates whether the message describes a major update for the service.
    isMajorChange *bool
    // The affected services by the service message.
    services []string
    // The severity property
    severity *ServiceUpdateSeverity
    // A collection of tags for the service message. Tags are provided by the service team/support team who post the message to tell whether this message contains privacy data, or whether this message is for a service new feature update, and so on.
    tags []string
    // Represents user viewpoints data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions.
    viewPoint ServiceUpdateMessageViewpointable
}
// NewServiceUpdateMessage instantiates a new ServiceUpdateMessage and sets the default values.
func NewServiceUpdateMessage()(*ServiceUpdateMessage) {
    m := &ServiceUpdateMessage{
        ServiceAnnouncementBase: *NewServiceAnnouncementBase(),
    }
    odataTypeValue := "#microsoft.graph.serviceUpdateMessage"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServiceUpdateMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceUpdateMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceUpdateMessage(), nil
}
// GetActionRequiredByDateTime gets the actionRequiredByDateTime property value. The expected deadline of the action for the message.
func (m *ServiceUpdateMessage) GetActionRequiredByDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.actionRequiredByDateTime
}
// GetAttachments gets the attachments property value. A collection of serviceAnnouncementAttachments.
func (m *ServiceUpdateMessage) GetAttachments()([]ServiceAnnouncementAttachmentable) {
    return m.attachments
}
// GetAttachmentsArchive gets the attachmentsArchive property value. The zip file that contains all attachments for a message.
func (m *ServiceUpdateMessage) GetAttachmentsArchive()([]byte) {
    return m.attachmentsArchive
}
// GetBody gets the body property value. The body property
func (m *ServiceUpdateMessage) GetBody()(ItemBodyable) {
    return m.body
}
// GetCategory gets the category property value. The category property
func (m *ServiceUpdateMessage) GetCategory()(*ServiceUpdateCategory) {
    return m.category
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceUpdateMessage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ServiceAnnouncementBase.GetFieldDeserializers()
    res["actionRequiredByDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionRequiredByDateTime(val)
        }
        return nil
    }
    res["attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceAnnouncementAttachmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceAnnouncementAttachmentable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceAnnouncementAttachmentable)
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["attachmentsArchive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentsArchive(val)
        }
        return nil
    }
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
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ServiceUpdateCategory))
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
    res["isMajorChange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMajorChange(val)
        }
        return nil
    }
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetServices(res)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceUpdateSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*ServiceUpdateSeverity))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["viewPoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceUpdateMessageViewpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewPoint(val.(ServiceUpdateMessageViewpointable))
        }
        return nil
    }
    return res
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether the message has any attachment.
func (m *ServiceUpdateMessage) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetIsMajorChange gets the isMajorChange property value. Indicates whether the message describes a major update for the service.
func (m *ServiceUpdateMessage) GetIsMajorChange()(*bool) {
    return m.isMajorChange
}
// GetServices gets the services property value. The affected services by the service message.
func (m *ServiceUpdateMessage) GetServices()([]string) {
    return m.services
}
// GetSeverity gets the severity property value. The severity property
func (m *ServiceUpdateMessage) GetSeverity()(*ServiceUpdateSeverity) {
    return m.severity
}
// GetTags gets the tags property value. A collection of tags for the service message. Tags are provided by the service team/support team who post the message to tell whether this message contains privacy data, or whether this message is for a service new feature update, and so on.
func (m *ServiceUpdateMessage) GetTags()([]string) {
    return m.tags
}
// GetViewPoint gets the viewPoint property value. Represents user viewpoints data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions.
func (m *ServiceUpdateMessage) GetViewPoint()(ServiceUpdateMessageViewpointable) {
    return m.viewPoint
}
// Serialize serializes information the current object
func (m *ServiceUpdateMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ServiceAnnouncementBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("actionRequiredByDateTime", m.GetActionRequiredByDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetAttachments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("attachmentsArchive", m.GetAttachmentsArchive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
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
        err = writer.WriteBoolValue("isMajorChange", m.GetIsMajorChange())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        err = writer.WriteCollectionOfStringValues("services", m.GetServices())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("viewPoint", m.GetViewPoint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionRequiredByDateTime sets the actionRequiredByDateTime property value. The expected deadline of the action for the message.
func (m *ServiceUpdateMessage) SetActionRequiredByDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.actionRequiredByDateTime = value
}
// SetAttachments sets the attachments property value. A collection of serviceAnnouncementAttachments.
func (m *ServiceUpdateMessage) SetAttachments(value []ServiceAnnouncementAttachmentable)() {
    m.attachments = value
}
// SetAttachmentsArchive sets the attachmentsArchive property value. The zip file that contains all attachments for a message.
func (m *ServiceUpdateMessage) SetAttachmentsArchive(value []byte)() {
    m.attachmentsArchive = value
}
// SetBody sets the body property value. The body property
func (m *ServiceUpdateMessage) SetBody(value ItemBodyable)() {
    m.body = value
}
// SetCategory sets the category property value. The category property
func (m *ServiceUpdateMessage) SetCategory(value *ServiceUpdateCategory)() {
    m.category = value
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether the message has any attachment.
func (m *ServiceUpdateMessage) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetIsMajorChange sets the isMajorChange property value. Indicates whether the message describes a major update for the service.
func (m *ServiceUpdateMessage) SetIsMajorChange(value *bool)() {
    m.isMajorChange = value
}
// SetServices sets the services property value. The affected services by the service message.
func (m *ServiceUpdateMessage) SetServices(value []string)() {
    m.services = value
}
// SetSeverity sets the severity property value. The severity property
func (m *ServiceUpdateMessage) SetSeverity(value *ServiceUpdateSeverity)() {
    m.severity = value
}
// SetTags sets the tags property value. A collection of tags for the service message. Tags are provided by the service team/support team who post the message to tell whether this message contains privacy data, or whether this message is for a service new feature update, and so on.
func (m *ServiceUpdateMessage) SetTags(value []string)() {
    m.tags = value
}
// SetViewPoint sets the viewPoint property value. Represents user viewpoints data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions.
func (m *ServiceUpdateMessage) SetViewPoint(value ServiceUpdateMessageViewpointable)() {
    m.viewPoint = value
}
