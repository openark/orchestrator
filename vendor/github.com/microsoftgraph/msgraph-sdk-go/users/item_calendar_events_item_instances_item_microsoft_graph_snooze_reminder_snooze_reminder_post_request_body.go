package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody 
type ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The NewReminderTime property
    newReminderTime iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable
}
// NewItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody instantiates a new ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody and sets the default values.
func NewItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody()(*ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) {
    m := &ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newReminderTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewReminderTime(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetNewReminderTime gets the newReminderTime property value. The NewReminderTime property
func (m *ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetNewReminderTime()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable) {
    return m.newReminderTime
}
// Serialize serializes information the current object
func (m *ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newReminderTime", m.GetNewReminderTime())
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
func (m *ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewReminderTime sets the newReminderTime property value. The NewReminderTime property
func (m *ItemCalendarEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) SetNewReminderTime(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable)() {
    m.newReminderTime = value
}
