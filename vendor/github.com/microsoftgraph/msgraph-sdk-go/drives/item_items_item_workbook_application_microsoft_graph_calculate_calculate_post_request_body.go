package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody 
type ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The calculationType property
    calculationType *string
}
// NewItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody instantiates a new ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody and sets the default values.
func NewItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody()(*ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody) {
    m := &ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCalculationType gets the calculationType property value. The calculationType property
func (m *ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody) GetCalculationType()(*string) {
    return m.calculationType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calculationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalculationType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("calculationType", m.GetCalculationType())
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
func (m *ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCalculationType sets the calculationType property value. The calculationType property
func (m *ItemItemsItemWorkbookApplicationMicrosoftGraphCalculateCalculatePostRequestBody) SetCalculationType(value *string)() {
    m.calculationType = value
}
