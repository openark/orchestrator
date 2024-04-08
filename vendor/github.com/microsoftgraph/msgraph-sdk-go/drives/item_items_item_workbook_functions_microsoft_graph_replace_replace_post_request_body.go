package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The newText property
    newText iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The numChars property
    numChars iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The oldText property
    oldText iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The startNum property
    startNum iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["numChars"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumChars(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
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
    res["startNum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartNum(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetNewText gets the newText property value. The newText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) GetNewText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.newText
}
// GetNumChars gets the numChars property value. The numChars property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) GetNumChars()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.numChars
}
// GetOldText gets the oldText property value. The oldText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) GetOldText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.oldText
}
// GetStartNum gets the startNum property value. The startNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) GetStartNum()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.startNum
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newText", m.GetNewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numChars", m.GetNumChars())
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
        err := writer.WriteObjectValue("startNum", m.GetStartNum())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewText sets the newText property value. The newText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) SetNewText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.newText = value
}
// SetNumChars sets the numChars property value. The numChars property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) SetNumChars(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.numChars = value
}
// SetOldText sets the oldText property value. The oldText property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) SetOldText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.oldText = value
}
// SetStartNum sets the startNum property value. The startNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphReplaceReplacePostRequestBody) SetStartNum(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.startNum = value
}
