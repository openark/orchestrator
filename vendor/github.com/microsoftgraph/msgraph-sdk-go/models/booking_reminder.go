package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingReminder this type represents when and to whom to send an e-mail reminder.
type BookingReminder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The message in the reminder.
    message *string
    // The OdataType property
    odataType *string
    // The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
    offset *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The recipients property
    recipients *BookingReminderRecipients
}
// NewBookingReminder instantiates a new bookingReminder and sets the default values.
func NewBookingReminder()(*BookingReminder) {
    m := &BookingReminder{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBookingReminderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingReminderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingReminder(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingReminder) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingReminder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["offset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffset(val)
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingReminderRecipients)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipients(val.(*BookingReminderRecipients))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message in the reminder.
func (m *BookingReminder) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BookingReminder) GetOdataType()(*string) {
    return m.odataType
}
// GetOffset gets the offset property value. The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
func (m *BookingReminder) GetOffset()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.offset
}
// GetRecipients gets the recipients property value. The recipients property
func (m *BookingReminder) GetRecipients()(*BookingReminderRecipients) {
    return m.recipients
}
// Serialize serializes information the current object
func (m *BookingReminder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("offset", m.GetOffset())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := (*m.GetRecipients()).String()
        err := writer.WriteStringValue("recipients", &cast)
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
func (m *BookingReminder) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. The message in the reminder.
func (m *BookingReminder) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BookingReminder) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOffset sets the offset property value. The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
func (m *BookingReminder) SetOffset(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.offset = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *BookingReminder) SetRecipients(value *BookingReminderRecipients)() {
    m.recipients = value
}
