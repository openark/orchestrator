package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserFeedbackable 
type UserFeedbackable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetRating()(*UserFeedbackRating)
    GetText()(*string)
    GetTokens()(FeedbackTokenSetable)
    SetOdataType(value *string)()
    SetRating(value *UserFeedbackRating)()
    SetText(value *string)()
    SetTokens(value FeedbackTokenSetable)()
}
