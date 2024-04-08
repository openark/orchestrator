package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody 
type ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color property
    color *string
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody()(*ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody) GetColor()(*string) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
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
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color property
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemLegendFormatFillMicrosoftGraphSetSolidColorSetSolidColorPostRequestBody) SetColor(value *string)() {
    m.color = value
}
