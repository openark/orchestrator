package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventMessageable 
type EventMessageable interface {
    Messageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDateTime()(DateTimeTimeZoneable)
    GetEvent()(Eventable)
    GetIsAllDay()(*bool)
    GetIsDelegated()(*bool)
    GetIsOutOfDate()(*bool)
    GetLocation()(Locationable)
    GetMeetingMessageType()(*MeetingMessageType)
    GetRecurrence()(PatternedRecurrenceable)
    GetStartDateTime()(DateTimeTimeZoneable)
    GetType()(*EventType)
    SetEndDateTime(value DateTimeTimeZoneable)()
    SetEvent(value Eventable)()
    SetIsAllDay(value *bool)()
    SetIsDelegated(value *bool)()
    SetIsOutOfDate(value *bool)()
    SetLocation(value Locationable)()
    SetMeetingMessageType(value *MeetingMessageType)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
    SetType(value *EventType)()
}
