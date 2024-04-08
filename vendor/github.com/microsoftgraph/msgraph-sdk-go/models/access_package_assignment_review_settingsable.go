package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentReviewSettingsable 
type AccessPackageAssignmentReviewSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationBehavior()(*AccessReviewExpirationBehavior)
    GetFallbackReviewers()([]SubjectSetable)
    GetIsEnabled()(*bool)
    GetIsRecommendationEnabled()(*bool)
    GetIsReviewerJustificationRequired()(*bool)
    GetIsSelfReview()(*bool)
    GetOdataType()(*string)
    GetPrimaryReviewers()([]SubjectSetable)
    GetSchedule()(EntitlementManagementScheduleable)
    SetExpirationBehavior(value *AccessReviewExpirationBehavior)()
    SetFallbackReviewers(value []SubjectSetable)()
    SetIsEnabled(value *bool)()
    SetIsRecommendationEnabled(value *bool)()
    SetIsReviewerJustificationRequired(value *bool)()
    SetIsSelfReview(value *bool)()
    SetOdataType(value *string)()
    SetPrimaryReviewers(value []SubjectSetable)()
    SetSchedule(value EntitlementManagementScheduleable)()
}
