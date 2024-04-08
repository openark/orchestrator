package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The denominator property
    denominator iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The numerator property
    numerator iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDenominator gets the denominator property value. The denominator property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) GetDenominator()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.denominator
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["denominator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDenominator(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["numerator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumerator(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetNumerator gets the numerator property value. The numerator property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) GetNumerator()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.numerator
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("denominator", m.GetDenominator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numerator", m.GetNumerator())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDenominator sets the denominator property value. The denominator property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) SetDenominator(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.denominator = value
}
// SetNumerator sets the numerator property value. The numerator property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphQuotientQuotientPostRequestBody) SetNumerator(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.numerator = value
}
