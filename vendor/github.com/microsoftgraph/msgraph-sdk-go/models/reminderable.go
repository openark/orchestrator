package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Reminderable 
type Reminderable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeKey()(*string)
    GetEventEndTime()(DateTimeTimeZoneable)
    GetEventId()(*string)
    GetEventLocation()(Locationable)
    GetEventStartTime()(DateTimeTimeZoneable)
    GetEventSubject()(*string)
    GetEventWebLink()(*string)
    GetOdataType()(*string)
    GetReminderFireTime()(DateTimeTimeZoneable)
    SetChangeKey(value *string)()
    SetEventEndTime(value DateTimeTimeZoneable)()
    SetEventId(value *string)()
    SetEventLocation(value Locationable)()
    SetEventStartTime(value DateTimeTimeZoneable)()
    SetEventSubject(value *string)()
    SetEventWebLink(value *string)()
    SetOdataType(value *string)()
    SetReminderFireTime(value DateTimeTimeZoneable)()
}
