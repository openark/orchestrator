package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarSharingMessageable 
type CalendarSharingMessageable interface {
    Messageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanAccept()(*bool)
    GetSharingMessageAction()(CalendarSharingMessageActionable)
    GetSharingMessageActions()([]CalendarSharingMessageActionable)
    GetSuggestedCalendarName()(*string)
    SetCanAccept(value *bool)()
    SetSharingMessageAction(value CalendarSharingMessageActionable)()
    SetSharingMessageActions(value []CalendarSharingMessageActionable)()
    SetSuggestedCalendarName(value *string)()
}
