package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody 
type ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The Message property
    message iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable
}
// NewItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody instantiates a new ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody and sets the default values.
func NewItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody()(*ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) {
    m := &ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) GetMessage()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable) {
    return m.message
}
// Serialize serializes information the current object
func (m *ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetMessage sets the message property value. The Message property
func (m *ItemMessagesItemMicrosoftGraphCreateReplyAllCreateReplyAllPostRequestBody) SetMessage(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable)() {
    m.message = value
}
