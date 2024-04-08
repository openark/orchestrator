package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The endDate property
    endDate iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The holidays property
    holidays iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The startDate property
    startDate iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEndDate gets the endDate property value. The endDate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) GetEndDate()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["holidays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHolidays(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetHolidays gets the holidays property value. The holidays property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) GetHolidays()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.holidays
}
// GetStartDate gets the startDate property value. The startDate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) GetStartDate()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.startDate
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("holidays", m.GetHolidays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startDate", m.GetStartDate())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEndDate sets the endDate property value. The endDate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) SetEndDate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.endDate = value
}
// SetHolidays sets the holidays property value. The holidays property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) SetHolidays(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.holidays = value
}
// SetStartDate sets the startDate property value. The startDate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphNetworkDaysNetworkDaysPostRequestBody) SetStartDate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.startDate = value
}
