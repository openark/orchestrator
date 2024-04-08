package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The decimalSeparator property
    decimalSeparator iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The groupSeparator property
    groupSeparator iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The text property
    text iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDecimalSeparator gets the decimalSeparator property value. The decimalSeparator property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) GetDecimalSeparator()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.decimalSeparator
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decimalSeparator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecimalSeparator(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["groupSeparator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupSeparator(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetGroupSeparator gets the groupSeparator property value. The groupSeparator property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) GetGroupSeparator()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.groupSeparator
}
// GetText gets the text property value. The text property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) GetText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.text
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("decimalSeparator", m.GetDecimalSeparator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("groupSeparator", m.GetGroupSeparator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("text", m.GetText())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDecimalSeparator sets the decimalSeparator property value. The decimalSeparator property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) SetDecimalSeparator(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.decimalSeparator = value
}
// SetGroupSeparator sets the groupSeparator property value. The groupSeparator property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) SetGroupSeparator(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.groupSeparator = value
}
// SetText sets the text property value. The text property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNumberValueNumberValuePostRequestBody) SetText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.text = value
}
