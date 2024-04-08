package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttachmentItem 
type AttachmentItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type of attachment. Possible values are: file, item, reference. Required.
    attachmentType *AttachmentType
    // The CID or Content-Id of the attachment for referencing in case of in-line attachments using <img src='cid:contentId'> tag in HTML messages. Optional.
    contentId *string
    // The nature of the data in the attachment. Optional.
    contentType *string
    // true if the attachment is an inline attachment; otherwise, false. Optional.
    isInline *bool
    // The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
    name *string
    // The OdataType property
    odataType *string
    // The length of the attachment in bytes. Required.
    size *int64
}
// NewAttachmentItem instantiates a new attachmentItem and sets the default values.
func NewAttachmentItem()(*AttachmentItem) {
    m := &AttachmentItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAttachmentItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttachmentItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttachmentItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttachmentItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachmentType gets the attachmentType property value. The type of attachment. Possible values are: file, item, reference. Required.
func (m *AttachmentItem) GetAttachmentType()(*AttachmentType) {
    return m.attachmentType
}
// GetContentId gets the contentId property value. The CID or Content-Id of the attachment for referencing in case of in-line attachments using <img src='cid:contentId'> tag in HTML messages. Optional.
func (m *AttachmentItem) GetContentId()(*string) {
    return m.contentId
}
// GetContentType gets the contentType property value. The nature of the data in the attachment. Optional.
func (m *AttachmentItem) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttachmentItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttachmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentType(val.(*AttachmentType))
        }
        return nil
    }
    res["contentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentId(val)
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
    res["isInline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInline(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
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
// GetIsInline gets the isInline property value. true if the attachment is an inline attachment; otherwise, false. Optional.
func (m *AttachmentItem) GetIsInline()(*bool) {
    return m.isInline
}
// GetName gets the name property value. The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
func (m *AttachmentItem) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttachmentItem) GetOdataType()(*string) {
    return m.odataType
}
// GetSize gets the size property value. The length of the attachment in bytes. Required.
func (m *AttachmentItem) GetSize()(*int64) {
    return m.size
}
// Serialize serializes information the current object
func (m *AttachmentItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttachmentType() != nil {
        cast := (*m.GetAttachmentType()).String()
        err := writer.WriteStringValue("attachmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentId", m.GetContentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isInline", m.GetIsInline())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttachmentItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachmentType sets the attachmentType property value. The type of attachment. Possible values are: file, item, reference. Required.
func (m *AttachmentItem) SetAttachmentType(value *AttachmentType)() {
    m.attachmentType = value
}
// SetContentId sets the contentId property value. The CID or Content-Id of the attachment for referencing in case of in-line attachments using <img src='cid:contentId'> tag in HTML messages. Optional.
func (m *AttachmentItem) SetContentId(value *string)() {
    m.contentId = value
}
// SetContentType sets the contentType property value. The nature of the data in the attachment. Optional.
func (m *AttachmentItem) SetContentType(value *string)() {
    m.contentType = value
}
// SetIsInline sets the isInline property value. true if the attachment is an inline attachment; otherwise, false. Optional.
func (m *AttachmentItem) SetIsInline(value *bool)() {
    m.isInline = value
}
// SetName sets the name property value. The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
func (m *AttachmentItem) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttachmentItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSize sets the size property value. The length of the attachment in bytes. Required.
func (m *AttachmentItem) SetSize(value *int64)() {
    m.size = value
}
