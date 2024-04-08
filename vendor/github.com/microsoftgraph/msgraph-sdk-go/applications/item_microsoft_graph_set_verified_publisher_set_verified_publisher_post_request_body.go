package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody 
type ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The verifiedPublisherId property
    verifiedPublisherId *string
}
// NewItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody instantiates a new ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody and sets the default values.
func NewItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody()(*ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) {
    m := &ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["verifiedPublisherId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisherId(val)
        }
        return nil
    }
    return res
}
// GetVerifiedPublisherId gets the verifiedPublisherId property value. The verifiedPublisherId property
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) GetVerifiedPublisherId()(*string) {
    return m.verifiedPublisherId
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("verifiedPublisherId", m.GetVerifiedPublisherId())
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
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetVerifiedPublisherId sets the verifiedPublisherId property value. The verifiedPublisherId property
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) SetVerifiedPublisherId(value *string)() {
    m.verifiedPublisherId = value
}
