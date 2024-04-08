package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody 
type ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The contentType property
    contentType *string
}
// NewItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody instantiates a new ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody and sets the default values.
func NewItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody()(*ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody) {
    m := &ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentType gets the contentType property value. The contentType property
func (m *ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// Serialize serializes information the current object
func (m *ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
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
func (m *ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentType sets the contentType property value. The contentType property
func (m *ItemContentTypesMicrosoftGraphAddCopyAddCopyPostRequestBody) SetContentType(value *string)() {
    m.contentType = value
}
