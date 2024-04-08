package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The instanceNum property
    instanceNum iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The newText property
    newText iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The oldText property
    oldText iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The text property
    text iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["instanceNum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceNum(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["newText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewText(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["oldText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldText(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
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
// GetInstanceNum gets the instanceNum property value. The instanceNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) GetInstanceNum()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.instanceNum
}
// GetNewText gets the newText property value. The newText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) GetNewText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.newText
}
// GetOldText gets the oldText property value. The oldText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) GetOldText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.oldText
}
// GetText gets the text property value. The text property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) GetText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.text
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("instanceNum", m.GetInstanceNum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("newText", m.GetNewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("oldText", m.GetOldText())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInstanceNum sets the instanceNum property value. The instanceNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) SetInstanceNum(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.instanceNum = value
}
// SetNewText sets the newText property value. The newText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) SetNewText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.newText = value
}
// SetOldText sets the oldText property value. The oldText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) SetOldText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.oldText = value
}
// SetText sets the text property value. The text property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphSubstituteSubstitutePostRequestBody) SetText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.text = value
}
