package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody 
type ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address property
    address *string
    // The hasHeaders property
    hasHeaders *bool
}
// NewItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody()(*ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The address property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) GetAddress()(*string) {
    return m.address
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["hasHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasHeaders(val)
        }
        return nil
    }
    return res
}
// GetHasHeaders gets the hasHeaders property value. The hasHeaders property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) GetHasHeaders()(*bool) {
    return m.hasHeaders
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasHeaders", m.GetHasHeaders())
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
func (m *ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The address property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) SetAddress(value *string)() {
    m.address = value
}
// SetHasHeaders sets the hasHeaders property value. The hasHeaders property
func (m *ItemItemsItemWorkbookWorksheetsItemTablesMicrosoftGraphAddAddPostRequestBody) SetHasHeaders(value *bool)() {
    m.hasHeaders = value
}
