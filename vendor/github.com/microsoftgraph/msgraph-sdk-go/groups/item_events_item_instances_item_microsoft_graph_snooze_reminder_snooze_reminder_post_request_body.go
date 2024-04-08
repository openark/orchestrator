package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody 
type ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The NewReminderTime property
    newReminderTime iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable
}
// NewItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody instantiates a new ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody and sets the default values.
func NewItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody()(*ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) {
    m := &ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) GetNewReminderTime()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable) {
    return m.newReminderTime
}
// Serialize serializes information the current object
func (m *ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewReminderTime sets the newReminderTime property value. The NewReminderTime property
func (m *ItemEventsItemInstancesItemMicrosoftGraphSnoozeReminderSnoozeReminderPostRequestBody) SetNewReminderTime(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable)() {
    m.newReminderTime = value
}
