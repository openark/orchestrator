package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventMessageRequest 
type EventMessageRequest struct {
    EventMessage
    // True if the meeting organizer allows invitees to propose a new time when responding, false otherwise. Optional. Default is true.
    allowNewTimeProposals *bool
    // The meetingRequestType property
    meetingRequestType *MeetingRequestType
    // If the meeting update changes the meeting end time, this property specifies the previous meeting end time.
    previousEndDateTime DateTimeTimeZoneable
    // If the meeting update changes the meeting location, this property specifies the previous meeting location.
    previousLocation Locationable
    // If the meeting update changes the meeting start time, this property specifies the previous meeting start time.
    previousStartDateTime DateTimeTimeZoneable
    // Set to true if the sender would like the invitee to send a response to the requested meeting.
    responseRequested *bool
}
// NewEventMessageRequest instantiates a new EventMessageRequest and sets the default values.
func NewEventMessageRequest()(*EventMessageRequest) {
    m := &EventMessageRequest{
        EventMessage: *NewEventMessage(),
    }
    odataTypeValue := "#microsoft.graph.eventMessageRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEventMessageRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEventMessageRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventMessageRequest(), nil
}
// GetAllowNewTimeProposals gets the allowNewTimeProposals property value. True if the meeting organizer allows invitees to propose a new time when responding, false otherwise. Optional. Default is true.
func (m *EventMessageRequest) GetAllowNewTimeProposals()(*bool) {
    return m.allowNewTimeProposals
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EventMessageRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessage.GetFieldDeserializers()
    res["allowNewTimeProposals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowNewTimeProposals(val)
        }
        return nil
    }
    res["meetingRequestType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingRequestType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingRequestType(val.(*MeetingRequestType))
        }
        return nil
    }
    res["previousEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["previousLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousLocation(val.(Locationable))
        }
        return nil
    }
    res["previousStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["responseRequested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseRequested(val)
        }
        return nil
    }
    return res
}
// GetMeetingRequestType gets the meetingRequestType property value. The meetingRequestType property
func (m *EventMessageRequest) GetMeetingRequestType()(*MeetingRequestType) {
    return m.meetingRequestType
}
// GetPreviousEndDateTime gets the previousEndDateTime property value. If the meeting update changes the meeting end time, this property specifies the previous meeting end time.
func (m *EventMessageRequest) GetPreviousEndDateTime()(DateTimeTimeZoneable) {
    return m.previousEndDateTime
}
// GetPreviousLocation gets the previousLocation property value. If the meeting update changes the meeting location, this property specifies the previous meeting location.
func (m *EventMessageRequest) GetPreviousLocation()(Locationable) {
    return m.previousLocation
}
// GetPreviousStartDateTime gets the previousStartDateTime property value. If the meeting update changes the meeting start time, this property specifies the previous meeting start time.
func (m *EventMessageRequest) GetPreviousStartDateTime()(DateTimeTimeZoneable) {
    return m.previousStartDateTime
}
// GetResponseRequested gets the responseRequested property value. Set to true if the sender would like the invitee to send a response to the requested meeting.
func (m *EventMessageRequest) GetResponseRequested()(*bool) {
    return m.responseRequested
}
// Serialize serializes information the current object
func (m *EventMessageRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowNewTimeProposals", m.GetAllowNewTimeProposals())
        if err != nil {
            return err
        }
    }
    if m.GetMeetingRequestType() != nil {
        cast := (*m.GetMeetingRequestType()).String()
        err = writer.WriteStringValue("meetingRequestType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("previousEndDateTime", m.GetPreviousEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("previousLocation", m.GetPreviousLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("previousStartDateTime", m.GetPreviousStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("responseRequested", m.GetResponseRequested())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowNewTimeProposals sets the allowNewTimeProposals property value. True if the meeting organizer allows invitees to propose a new time when responding, false otherwise. Optional. Default is true.
func (m *EventMessageRequest) SetAllowNewTimeProposals(value *bool)() {
    m.allowNewTimeProposals = value
}
// SetMeetingRequestType sets the meetingRequestType property value. The meetingRequestType property
func (m *EventMessageRequest) SetMeetingRequestType(value *MeetingRequestType)() {
    m.meetingRequestType = value
}
// SetPreviousEndDateTime sets the previousEndDateTime property value. If the meeting update changes the meeting end time, this property specifies the previous meeting end time.
func (m *EventMessageRequest) SetPreviousEndDateTime(value DateTimeTimeZoneable)() {
    m.previousEndDateTime = value
}
// SetPreviousLocation sets the previousLocation property value. If the meeting update changes the meeting location, this property specifies the previous meeting location.
func (m *EventMessageRequest) SetPreviousLocation(value Locationable)() {
    m.previousLocation = value
}
// SetPreviousStartDateTime sets the previousStartDateTime property value. If the meeting update changes the meeting start time, this property specifies the previous meeting start time.
func (m *EventMessageRequest) SetPreviousStartDateTime(value DateTimeTimeZoneable)() {
    m.previousStartDateTime = value
}
// SetResponseRequested sets the responseRequested property value. Set to true if the sender would like the invitee to send a response to the requested meeting.
func (m *EventMessageRequest) SetResponseRequested(value *bool)() {
    m.responseRequested = value
}
