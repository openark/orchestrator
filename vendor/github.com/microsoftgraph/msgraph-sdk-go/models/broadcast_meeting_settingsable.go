package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BroadcastMeetingSettingsable 
type BroadcastMeetingSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedAudience()(*BroadcastMeetingAudience)
    GetCaptions()(BroadcastMeetingCaptionSettingsable)
    GetIsAttendeeReportEnabled()(*bool)
    GetIsQuestionAndAnswerEnabled()(*bool)
    GetIsRecordingEnabled()(*bool)
    GetIsVideoOnDemandEnabled()(*bool)
    GetOdataType()(*string)
    SetAllowedAudience(value *BroadcastMeetingAudience)()
    SetCaptions(value BroadcastMeetingCaptionSettingsable)()
    SetIsAttendeeReportEnabled(value *bool)()
    SetIsQuestionAndAnswerEnabled(value *bool)()
    SetIsRecordingEnabled(value *bool)()
    SetIsVideoOnDemandEnabled(value *bool)()
    SetOdataType(value *string)()
}
