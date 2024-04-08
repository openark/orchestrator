package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CalendarViewItemMicrosoftGraphDeclineDeclinePostRequestBodyable 
type CalendarViewItemMicrosoftGraphDeclineDeclinePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetProposedNewTime()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeSlotable)
    GetSendResponse()(*bool)
    SetComment(value *string)()
    SetProposedNewTime(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TimeSlotable)()
    SetSendResponse(value *bool)()
}
