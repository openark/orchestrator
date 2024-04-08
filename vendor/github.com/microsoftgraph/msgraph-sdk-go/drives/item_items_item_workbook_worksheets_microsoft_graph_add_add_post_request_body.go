package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody 
type ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
}
// NewItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody instantiates a new ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody()(*ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody) {
    m := &ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetName gets the name property value. The name property
func (m *ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *ItemItemsItemWorkbookWorksheetsMicrosoftGraphAddAddPostRequestBody) SetName(value *string)() {
    m.name = value
}
