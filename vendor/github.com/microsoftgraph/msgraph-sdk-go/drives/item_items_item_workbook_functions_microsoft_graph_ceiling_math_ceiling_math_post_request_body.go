package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The mode property
    mode iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The number property
    number iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The significance property
    significance iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMode(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["significance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignificance(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetMode gets the mode property value. The mode property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) GetMode()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.mode
}
// GetNumber gets the number property value. The number property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) GetNumber()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.number
}
// GetSignificance gets the significance property value. The significance property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) GetSignificance()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.significance
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("mode", m.GetMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("significance", m.GetSignificance())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMode sets the mode property value. The mode property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) SetMode(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.mode = value
}
// SetNumber sets the number property value. The number property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) SetNumber(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.number = value
}
// SetSignificance sets the significance property value. The significance property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphCeiling_MathCeiling_MathPostRequestBody) SetSignificance(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.significance = value
}
