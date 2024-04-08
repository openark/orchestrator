package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody 
type ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comment property
    comment *string
    // The formula property
    formula *string
    // The name property
    name *string
}
// NewItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody instantiates a new ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody()(*ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) {
    m := &ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The comment property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["formula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormula(val)
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
    return res
}
// GetFormula gets the formula property value. The formula property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) GetFormula()(*string) {
    return m.formula
}
// GetName gets the name property value. The name property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The comment property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetFormula sets the formula property value. The formula property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) SetFormula(value *string)() {
    m.formula = value
}
// SetName sets the name property value. The name property
func (m *ItemItemsItemWorkbookNamesMicrosoftGraphAddFormulaLocalAddFormulaLocalPostRequestBody) SetName(value *string)() {
    m.name = value
}
