package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageAttachment 
type ChatMessageAttachment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The content of the attachment. If the attachment is a rich card, set the property to the rich card object. This property and contentUrl are mutually exclusive.
    content *string
    // The media type of the content attachment. It can have the following values: reference: Attachment is a link to another file. Populate the contentURL with the link to the object.Any contentTypes supported by the Bot Framework's Attachment objectapplication/vnd.microsoft.card.codesnippet: A code snippet. application/vnd.microsoft.card.announcement: An announcement header.
    contentType *string
    // URL for the content of the attachment. Supported protocols: http, https, file and data.
    contentUrl *string
    // Read-only. Unique id of the attachment.
    id *string
    // Name of the attachment.
    name *string
    // The OdataType property
    odataType *string
    // The ID of the Teams app that is associated with the attachment. The property is specifically used to attribute a Teams message card to the specified app.
    teamsAppId *string
    // URL to a thumbnail image that the channel can use if it supports using an alternative, smaller form of content or contentUrl. For example, if you set contentType to application/word and set contentUrl to the location of the Word document, you might include a thumbnail image that represents the document. The channel could display the thumbnail image instead of the document. When the user clicks the image, the channel would open the document.
    thumbnailUrl *string
}
// NewChatMessageAttachment instantiates a new chatMessageAttachment and sets the default values.
func NewChatMessageAttachment()(*ChatMessageAttachment) {
    m := &ChatMessageAttachment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChatMessageAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageAttachment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessageAttachment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. The content of the attachment. If the attachment is a rich card, set the property to the rich card object. This property and contentUrl are mutually exclusive.
func (m *ChatMessageAttachment) GetContent()(*string) {
    return m.content
}
// GetContentType gets the contentType property value. The media type of the content attachment. It can have the following values: reference: Attachment is a link to another file. Populate the contentURL with the link to the object.Any contentTypes supported by the Bot Framework's Attachment objectapplication/vnd.microsoft.card.codesnippet: A code snippet. application/vnd.microsoft.card.announcement: An announcement header.
func (m *ChatMessageAttachment) GetContentType()(*string) {
    return m.contentType
}
// GetContentUrl gets the contentUrl property value. URL for the content of the attachment. Supported protocols: http, https, file and data.
func (m *ChatMessageAttachment) GetContentUrl()(*string) {
    return m.contentUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
    res["contentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["teamsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Read-only. Unique id of the attachment.
func (m *ChatMessageAttachment) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. Name of the attachment.
func (m *ChatMessageAttachment) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ChatMessageAttachment) GetOdataType()(*string) {
    return m.odataType
}
// GetTeamsAppId gets the teamsAppId property value. The ID of the Teams app that is associated with the attachment. The property is specifically used to attribute a Teams message card to the specified app.
func (m *ChatMessageAttachment) GetTeamsAppId()(*string) {
    return m.teamsAppId
}
// GetThumbnailUrl gets the thumbnailUrl property value. URL to a thumbnail image that the channel can use if it supports using an alternative, smaller form of content or contentUrl. For example, if you set contentType to application/word and set contentUrl to the location of the Word document, you might include a thumbnail image that represents the document. The channel could display the thumbnail image instead of the document. When the user clicks the image, the channel would open the document.
func (m *ChatMessageAttachment) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// Serialize serializes information the current object
func (m *ChatMessageAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
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
        err := writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
        err := writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
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
func (m *ChatMessageAttachment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. The content of the attachment. If the attachment is a rich card, set the property to the rich card object. This property and contentUrl are mutually exclusive.
func (m *ChatMessageAttachment) SetContent(value *string)() {
    m.content = value
}
// SetContentType sets the contentType property value. The media type of the content attachment. It can have the following values: reference: Attachment is a link to another file. Populate the contentURL with the link to the object.Any contentTypes supported by the Bot Framework's Attachment objectapplication/vnd.microsoft.card.codesnippet: A code snippet. application/vnd.microsoft.card.announcement: An announcement header.
func (m *ChatMessageAttachment) SetContentType(value *string)() {
    m.contentType = value
}
// SetContentUrl sets the contentUrl property value. URL for the content of the attachment. Supported protocols: http, https, file and data.
func (m *ChatMessageAttachment) SetContentUrl(value *string)() {
    m.contentUrl = value
}
// SetId sets the id property value. Read-only. Unique id of the attachment.
func (m *ChatMessageAttachment) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. Name of the attachment.
func (m *ChatMessageAttachment) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChatMessageAttachment) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTeamsAppId sets the teamsAppId property value. The ID of the Teams app that is associated with the attachment. The property is specifically used to attribute a Teams message card to the specified app.
func (m *ChatMessageAttachment) SetTeamsAppId(value *string)() {
    m.teamsAppId = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. URL to a thumbnail image that the channel can use if it supports using an alternative, smaller form of content or contentUrl. For example, if you set contentType to application/word and set contentUrl to the location of the Word document, you might include a thumbnail image that represents the document. The channel could display the thumbnail image instead of the document. When the user clicks the image, the channel would open the document.
func (m *ChatMessageAttachment) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
