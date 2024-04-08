package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The lookupValue property
    lookupValue iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The rangeLookup property
    rangeLookup iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The rowIndexNum property
    rowIndexNum iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The tableArray property
    tableArray iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["rowIndexNum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowIndexNum(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) GetLookupValue()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.lookupValue
}
// GetRangeLookup gets the rangeLookup property value. The rangeLookup property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) GetRangeLookup()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.rangeLookup
}
// GetRowIndexNum gets the rowIndexNum property value. The rowIndexNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) GetRowIndexNum()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.rowIndexNum
}
// GetTableArray gets the tableArray property value. The tableArray property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) GetTableArray()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.tableArray
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("rowIndexNum", m.GetRowIndexNum())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLookupValue sets the lookupValue property value. The lookupValue property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) SetLookupValue(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.lookupValue = value
}
// SetRangeLookup sets the rangeLookup property value. The rangeLookup property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) SetRangeLookup(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.rangeLookup = value
}
// SetRowIndexNum sets the rowIndexNum property value. The rowIndexNum property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) SetRowIndexNum(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.rowIndexNum = value
}
// SetTableArray sets the tableArray property value. The tableArray property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphHlookupHlookupPostRequestBody) SetTableArray(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.tableArray = value
}
