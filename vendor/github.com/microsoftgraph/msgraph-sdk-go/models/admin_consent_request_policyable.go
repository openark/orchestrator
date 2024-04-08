package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminConsentRequestPolicyable 
type AdminConsentRequestPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsEnabled()(*bool)
    GetNotifyReviewers()(*bool)
    GetRemindersEnabled()(*bool)
    GetRequestDurationInDays()(*int32)
    GetReviewers()([]AccessReviewReviewerScopeable)
    GetVersion()(*int32)
    SetIsEnabled(value *bool)()
    SetNotifyReviewers(value *bool)()
    SetRemindersEnabled(value *bool)()
    SetRequestDurationInDays(value *int32)()
    SetReviewers(value []AccessReviewReviewerScopeable)()
    SetVersion(value *int32)()
}
