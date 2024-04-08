package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The colIndexNum property
    colIndexNum iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The lookupValue property
    lookupValue iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The rangeLookup property
    rangeLookup iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The tableArray property
    tableArray iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColIndexNum gets the colIndexNum property value. The colIndexNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) GetColIndexNum()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.colIndexNum
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["colIndexNum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColIndexNum(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["lookupValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLookupValue(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["rangeLookup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRangeLookup(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["tableArray"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTableArray(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetLookupValue gets the lookupValue property value. The lookupValue property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) GetLookupValue()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.lookupValue
}
// GetRangeLookup gets the rangeLookup property value. The rangeLookup property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) GetRangeLookup()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.rangeLookup
}
// GetTableArray gets the tableArray property value. The tableArray property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) GetTableArray()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.tableArray
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("colIndexNum", m.GetColIndexNum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lookupValue", m.GetLookupValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rangeLookup", m.GetRangeLookup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tableArray", m.GetTableArray())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColIndexNum sets the colIndexNum property value. The colIndexNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) SetColIndexNum(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.colIndexNum = value
}
// SetLookupValue sets the lookupValue property value. The lookupValue property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) SetLookupValue(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.lookupValue = value
}
// SetRangeLookup sets the rangeLookup property value. The rangeLookup property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) SetRangeLookup(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.rangeLookup = value
}
// SetTableArray sets the tableArray property value. The tableArray property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVlookupVlookupPostRequestBody) SetTableArray(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.tableArray = value
}
