package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The basis property
    basis iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The cost property
    cost iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The datePurchased property
    datePurchased iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The firstPeriod property
    firstPeriod iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The period property
    period iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The rate property
    rate iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The salvage property
    salvage iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBasis gets the basis property value. The basis property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetBasis()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.basis
}
// GetCost gets the cost property value. The cost property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetCost()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.cost
}
// GetDatePurchased gets the datePurchased property value. The datePurchased property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetDatePurchased()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.datePurchased
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["basis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasis(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
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
    res["datePurchased"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatePurchased(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["firstPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstPeriod(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["period"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["rate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRate(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
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
    return res
}
// GetFirstPeriod gets the firstPeriod property value. The firstPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetFirstPeriod()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.firstPeriod
}
// GetPeriod gets the period property value. The period property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetPeriod()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.period
}
// GetRate gets the rate property value. The rate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetRate()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.rate
}
// GetSalvage gets the salvage property value. The salvage property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) GetSalvage()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.salvage
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("basis", m.GetBasis())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cost", m.GetCost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("datePurchased", m.GetDatePurchased())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("firstPeriod", m.GetFirstPeriod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("period", m.GetPeriod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rate", m.GetRate())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBasis sets the basis property value. The basis property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) SetBasis(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.basis = value
}
// SetCost sets the cost property value. The cost property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) SetCost(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.cost = value
}
// SetDatePurchased sets the datePurchased property value. The datePurchased property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) SetDatePurchased(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.datePurchased = value
}
// SetFirstPeriod sets the firstPeriod property value. The firstPeriod property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) SetFirstPeriod(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.firstPeriod = value
}
// SetPeriod sets the period property value. The period property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) SetPeriod(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.period = value
}
// SetRate sets the rate property value. The rate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) SetRate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.rate = value
}
// SetSalvage sets the salvage property value. The salvage property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphAmorDegrcAmorDegrcPostRequestBody) SetSalvage(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.salvage = value
}
