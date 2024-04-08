package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody 
type ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The days property
    days iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The holidays property
    holidays iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The startDate property
    startDate iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
    // The weekend property
    weekend iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
}
// NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody()(*ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) {
    m := &ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDays gets the days property value. The days property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetDays()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.days
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["days"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDays(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
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
    res["weekend"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeekend(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
        }
        return nil
    }
    return res
}
// GetHolidays gets the holidays property value. The holidays property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetHolidays()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.holidays
}
// GetStartDate gets the startDate property value. The startDate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetStartDate()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.startDate
}
// GetWeekend gets the weekend property value. The weekend property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) GetWeekend()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
    return m.weekend
}
// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("days", m.GetDays())
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
        err := writer.WriteObjectValue("weekend", m.GetWeekend())
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
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDays sets the days property value. The days property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetDays(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.days = value
}
// SetHolidays sets the holidays property value. The holidays property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetHolidays(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.holidays = value
}
// SetStartDate sets the startDate property value. The startDate property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetStartDate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.startDate = value
}
// SetWeekend sets the weekend property value. The weekend property
func (m *ItemItemsItemWorkbookFunctionsMicrosoftGraphWorkDay_IntlWorkDay_IntlPostRequestBody) SetWeekend(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)() {
    m.weekend = value
}
