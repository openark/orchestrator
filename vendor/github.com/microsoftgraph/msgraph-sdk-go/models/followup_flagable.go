package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FollowupFlagable 
type FollowupFlagable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDateTime()(DateTimeTimeZoneable)
    GetDueDateTime()(DateTimeTimeZoneable)
    GetFlagStatus()(*FollowupFlagStatus)
    GetOdataType()(*string)
    GetStartDateTime()(DateTimeTimeZoneable)
    SetCompletedDateTime(value DateTimeTimeZoneable)()
    SetDueDateTime(value DateTimeTimeZoneable)()
    SetFlagStatus(value *FollowupFlagStatus)()
    SetOdataType(value *string)()
    SetStartDateTime(value DateTimeTimeZoneable)()
}
