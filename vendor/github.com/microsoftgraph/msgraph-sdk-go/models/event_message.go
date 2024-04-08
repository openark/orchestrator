package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventMessage 
type EventMessage struct {
    Message
    // The endDateTime property
    endDateTime DateTimeTimeZoneable
    // The event associated with the event message. The assumption for attendees or room resources is that the Calendar Attendant is set to automatically update the calendar with an event when meeting request event messages arrive. Navigation property.  Read-only.
    event Eventable
    // The isAllDay property
    isAllDay *bool
    // The isDelegated property
    isDelegated *bool
    // The isOutOfDate property
    isOutOfDate *bool
    // The location property
    location Locationable
    // The meetingMessageType property
    meetingMessageType *MeetingMessageType
    // The recurrence property
    recurrence PatternedRecurrenceable
    // The startDateTime property
    startDateTime DateTimeTimeZoneable
    // The type property
    typeEscaped *EventType
}
// NewEventMessage instantiates a new EventMessage and sets the default values.
func NewEventMessage()(*EventMessage) {
    m := &EventMessage{
        Message: *NewMessage(),
    }
    odataTypeValue := "#microsoft.graph.eventMessage"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEventMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEventMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.eventMessageRequest":
                        return NewEventMessageRequest(), nil
                    case "#microsoft.graph.eventMessageResponse":
                        return NewEventMessageResponse(), nil
                }
            }
        }
    }
    return NewEventMessage(), nil
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *EventMessage) GetEndDateTime()(DateTimeTimeZoneable) {
    return m.endDateTime
}
// GetEvent gets the event property value. The event associated with the event message. The assumption for attendees or room resources is that the Calendar Attendant is set to automatically update the calendar with an event when meeting request event messages arrive. Navigation property.  Read-only.
func (m *EventMessage) GetEvent()(Eventable) {
    return m.event
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EventMessage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Message.GetFieldDeserializers()
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["event"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvent(val.(Eventable))
        }
        return nil
    }
    res["isAllDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAllDay(val)
        }
        return nil
    }
    res["isDelegated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDelegated(val)
        }
        return nil
    }
    res["isOutOfDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOutOfDate(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(Locationable))
        }
        return nil
    }
    res["meetingMessageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingMessageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingMessageType(val.(*MeetingMessageType))
        }
        return nil
    }
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*EventType))
        }
        return nil
    }
    return res
}
// GetIsAllDay gets the isAllDay property value. The isAllDay property
func (m *EventMessage) GetIsAllDay()(*bool) {
    return m.isAllDay
}
// GetIsDelegated gets the isDelegated property value. The isDelegated property
func (m *EventMessage) GetIsDelegated()(*bool) {
    return m.isDelegated
}
// GetIsOutOfDate gets the isOutOfDate property value. The isOutOfDate property
func (m *EventMessage) GetIsOutOfDate()(*bool) {
    return m.isOutOfDate
}
// GetLocation gets the location property value. The location property
func (m *EventMessage) GetLocation()(Locationable) {
    return m.location
}
// GetMeetingMessageType gets the meetingMessageType property value. The meetingMessageType property
func (m *EventMessage) GetMeetingMessageType()(*MeetingMessageType) {
    return m.meetingMessageType
}
// GetRecurrence gets the recurrence property value. The recurrence property
func (m *EventMessage) GetRecurrence()(PatternedRecurrenceable) {
    return m.recurrence
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *EventMessage) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// GetType gets the type property value. The type property
func (m *EventMessage) GetType()(*EventType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *EventMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Message.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("event", m.GetEvent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAllDay", m.GetIsAllDay())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDelegated", m.GetIsDelegated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOutOfDate", m.GetIsOutOfDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    if m.GetMeetingMessageType() != nil {
        cast := (*m.GetMeetingMessageType()).String()
        err = writer.WriteStringValue("meetingMessageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *EventMessage) SetEndDateTime(value DateTimeTimeZoneable)() {
    m.endDateTime = value
}
// SetEvent sets the event property value. The event associated with the event message. The assumption for attendees or room resources is that the Calendar Attendant is set to automatically update the calendar with an event when meeting request event messages arrive. Navigation property.  Read-only.
func (m *EventMessage) SetEvent(value Eventable)() {
    m.event = value
}
// SetIsAllDay sets the isAllDay property value. The isAllDay property
func (m *EventMessage) SetIsAllDay(value *bool)() {
    m.isAllDay = value
}
// SetIsDelegated sets the isDelegated property value. The isDelegated property
func (m *EventMessage) SetIsDelegated(value *bool)() {
    m.isDelegated = value
}
// SetIsOutOfDate sets the isOutOfDate property value. The isOutOfDate property
func (m *EventMessage) SetIsOutOfDate(value *bool)() {
    m.isOutOfDate = value
}
// SetLocation sets the location property value. The location property
func (m *EventMessage) SetLocation(value Locationable)() {
    m.location = value
}
// SetMeetingMessageType sets the meetingMessageType property value. The meetingMessageType property
func (m *EventMessage) SetMeetingMessageType(value *MeetingMessageType)() {
    m.meetingMessageType = value
}
// SetRecurrence sets the recurrence property value. The recurrence property
func (m *EventMessage) SetRecurrence(value PatternedRecurrenceable)() {
    m.recurrence = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *EventMessage) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
// SetType sets the type property value. The type property
func (m *EventMessage) SetType(value *EventType)() {
    m.typeEscaped = value
}
