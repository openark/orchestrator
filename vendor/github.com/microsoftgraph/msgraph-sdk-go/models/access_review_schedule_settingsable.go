package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewScheduleSettingsable 
type AccessReviewScheduleSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplyActions()([]AccessReviewApplyActionable)
    GetAutoApplyDecisionsEnabled()(*bool)
    GetDecisionHistoriesForReviewersEnabled()(*bool)
    GetDefaultDecision()(*string)
    GetDefaultDecisionEnabled()(*bool)
    GetInstanceDurationInDays()(*int32)
    GetJustificationRequiredOnApproval()(*bool)
    GetMailNotificationsEnabled()(*bool)
    GetOdataType()(*string)
    GetRecommendationsEnabled()(*bool)
    GetRecurrence()(PatternedRecurrenceable)
    GetReminderNotificationsEnabled()(*bool)
    SetApplyActions(value []AccessReviewApplyActionable)()
    SetAutoApplyDecisionsEnabled(value *bool)()
    SetDecisionHistoriesForReviewersEnabled(value *bool)()
    SetDefaultDecision(value *string)()
    SetDefaultDecisionEnabled(value *bool)()
    SetInstanceDurationInDays(value *int32)()
    SetJustificationRequiredOnApproval(value *bool)()
    SetMailNotificationsEnabled(value *bool)()
    SetOdataType(value *string)()
    SetRecommendationsEnabled(value *bool)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetReminderNotificationsEnabled(value *bool)()
}
