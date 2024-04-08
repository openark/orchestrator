package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cost property
    cost iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The endPeriod property
    endPeriod iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The factor property
    factor iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The life property
    life iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The noSwitch property
    noSwitch iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The salvage property
    salvage iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The startPeriod property
    startPeriod iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCost gets the cost property value. The cost property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetCost()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.cost
}
// GetEndPeriod gets the endPeriod property value. The endPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetEndPeriod()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.endPeriod
}
// GetFactor gets the factor property value. The factor property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetFactor()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.factor
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCost(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["endPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndPeriod(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["factor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactor(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["life"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLife(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["noSwitch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoSwitch(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["salvage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSalvage(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["startPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartPeriod(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetLife gets the life property value. The life property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetLife()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.life
}
// GetNoSwitch gets the noSwitch property value. The noSwitch property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetNoSwitch()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.noSwitch
}
// GetSalvage gets the salvage property value. The salvage property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetSalvage()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.salvage
}
// GetStartPeriod gets the startPeriod property value. The startPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) GetStartPeriod()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.startPeriod
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cost", m.GetCost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("endPeriod", m.GetEndPeriod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("factor", m.GetFactor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("life", m.GetLife())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("noSwitch", m.GetNoSwitch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("salvage", m.GetSalvage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startPeriod", m.GetStartPeriod())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCost sets the cost property value. The cost property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) SetCost(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.cost = value
}
// SetEndPeriod sets the endPeriod property value. The endPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) SetEndPeriod(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.endPeriod = value
}
// SetFactor sets the factor property value. The factor property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) SetFactor(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.factor = value
}
// SetLife sets the life property value. The life property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) SetLife(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.life = value
}
// SetNoSwitch sets the noSwitch property value. The noSwitch property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) SetNoSwitch(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.noSwitch = value
}
// SetSalvage sets the salvage property value. The salvage property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) SetSalvage(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.salvage = value
}
// SetStartPeriod sets the startPeriod property value. The startPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphVdbVdbPostRequestBody) SetStartPeriod(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.startPeriod = value
}
