package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody 
type ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The Message property
    message iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable
}
// NewItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody instantiates a new ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody and sets the default values.
func NewItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody()(*ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) {
    m := &ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The Message property
func (m *ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) GetMessage()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable) {
    return m.message
}
// Serialize serializes information the current object
func (m *ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
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
func (m *ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetMessage sets the message property value. The Message property
func (m *ItemMailFoldersItemMessagesItemMicrosoftGraphReplyReplyPostRequestBody) SetMessage(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable)() {
    m.message = value
}
