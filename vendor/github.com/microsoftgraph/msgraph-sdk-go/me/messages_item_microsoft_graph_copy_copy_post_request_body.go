package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessagesItemMicrosoftGraphCopyCopyPostRequestBody 
type MessagesItemMicrosoftGraphCopyCopyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The DestinationId property
    destinationId *string
}
// NewMessagesItemMicrosoftGraphCopyCopyPostRequestBody instantiates a new MessagesItemMicrosoftGraphCopyCopyPostRequestBody and sets the default values.
func NewMessagesItemMicrosoftGraphCopyCopyPostRequestBody()(*MessagesItemMicrosoftGraphCopyCopyPostRequestBody) {
    m := &MessagesItemMicrosoftGraphCopyCopyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessagesItemMicrosoftGraphCopyCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessagesItemMicrosoftGraphCopyCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessagesItemMicrosoftGraphCopyCopyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessagesItemMicrosoftGraphCopyCopyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestinationId gets the destinationId property value. The DestinationId property
func (m *MessagesItemMicrosoftGraphCopyCopyPostRequestBody) GetDestinationId()(*string) {
    return m.destinationId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessagesItemMicrosoftGraphCopyCopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MessagesItemMicrosoftGraphCopyCopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationId", m.GetDestinationId())
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
func (m *MessagesItemMicrosoftGraphCopyCopyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestinationId sets the destinationId property value. The DestinationId property
func (m *MessagesItemMicrosoftGraphCopyCopyPostRequestBody) SetDestinationId(value *string)() {
    m.destinationId = value
}
