package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Reminder 
type Reminder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
    changeKey *string
    // The date, time and time zone that the event ends.
    eventEndTime DateTimeTimeZoneable
    // The unique ID of the event. Read only.
    eventId *string
    // The location of the event.
    eventLocation Locationable
    // The date, time, and time zone that the event starts.
    eventStartTime DateTimeTimeZoneable
    // The text of the event's subject line.
    eventSubject *string
    // The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
    eventWebLink *string
    // The OdataType property
    odataType *string
    // The date, time, and time zone that the reminder is set to occur.
    reminderFireTime DateTimeTimeZoneable
}
// NewReminder instantiates a new reminder and sets the default values.
func NewReminder()(*Reminder) {
    m := &Reminder{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReminderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReminderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReminder(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Reminder) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChangeKey gets the changeKey property value. Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
func (m *Reminder) GetChangeKey()(*string) {
    return m.changeKey
}
// GetEventEndTime gets the eventEndTime property value. The date, time and time zone that the event ends.
func (m *Reminder) GetEventEndTime()(DateTimeTimeZoneable) {
    return m.eventEndTime
}
// GetEventId gets the eventId property value. The unique ID of the event. Read only.
func (m *Reminder) GetEventId()(*string) {
    return m.eventId
}
// GetEventLocation gets the eventLocation property value. The location of the event.
func (m *Reminder) GetEventLocation()(Locationable) {
    return m.eventLocation
}
// GetEventStartTime gets the eventStartTime property value. The date, time, and time zone that the event starts.
func (m *Reminder) GetEventStartTime()(DateTimeTimeZoneable) {
    return m.eventStartTime
}
// GetEventSubject gets the eventSubject property value. The text of the event's subject line.
func (m *Reminder) GetEventSubject()(*string) {
    return m.eventSubject
}
// GetEventWebLink gets the eventWebLink property value. The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
func (m *Reminder) GetEventWebLink()(*string) {
    return m.eventWebLink
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Reminder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["changeKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["eventEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventEndTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["eventId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventId(val)
        }
        return nil
    }
    res["eventLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventLocation(val.(Locationable))
        }
        return nil
    }
    res["eventStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventStartTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["eventSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventSubject(val)
        }
        return nil
    }
    res["eventWebLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventWebLink(val)
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
    res["reminderFireTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderFireTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Reminder) GetOdataType()(*string) {
    return m.odataType
}
// GetReminderFireTime gets the reminderFireTime property value. The date, time, and time zone that the reminder is set to occur.
func (m *Reminder) GetReminderFireTime()(DateTimeTimeZoneable) {
    return m.reminderFireTime
}
// Serialize serializes information the current object
func (m *Reminder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventEndTime", m.GetEventEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventId", m.GetEventId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventLocation", m.GetEventLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventStartTime", m.GetEventStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventSubject", m.GetEventSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventWebLink", m.GetEventWebLink())
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
        err := writer.WriteObjectValue("reminderFireTime", m.GetReminderFireTime())
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
func (m *Reminder) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChangeKey sets the changeKey property value. Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
func (m *Reminder) SetChangeKey(value *string)() {
    m.changeKey = value
}
// SetEventEndTime sets the eventEndTime property value. The date, time and time zone that the event ends.
func (m *Reminder) SetEventEndTime(value DateTimeTimeZoneable)() {
    m.eventEndTime = value
}
// SetEventId sets the eventId property value. The unique ID of the event. Read only.
func (m *Reminder) SetEventId(value *string)() {
    m.eventId = value
}
// SetEventLocation sets the eventLocation property value. The location of the event.
func (m *Reminder) SetEventLocation(value Locationable)() {
    m.eventLocation = value
}
// SetEventStartTime sets the eventStartTime property value. The date, time, and time zone that the event starts.
func (m *Reminder) SetEventStartTime(value DateTimeTimeZoneable)() {
    m.eventStartTime = value
}
// SetEventSubject sets the eventSubject property value. The text of the event's subject line.
func (m *Reminder) SetEventSubject(value *string)() {
    m.eventSubject = value
}
// SetEventWebLink sets the eventWebLink property value. The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
func (m *Reminder) SetEventWebLink(value *string)() {
    m.eventWebLink = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Reminder) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReminderFireTime sets the reminderFireTime property value. The date, time, and time zone that the reminder is set to occur.
func (m *Reminder) SetReminderFireTime(value DateTimeTimeZoneable)() {
    m.reminderFireTime = value
}
