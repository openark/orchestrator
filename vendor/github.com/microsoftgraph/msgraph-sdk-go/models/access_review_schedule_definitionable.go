package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewScheduleDefinitionable 
type AccessReviewScheduleDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalNotificationRecipients()([]AccessReviewNotificationRecipientItemable)
    GetCreatedBy()(UserIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescriptionForAdmins()(*string)
    GetDescriptionForReviewers()(*string)
    GetDisplayName()(*string)
    GetFallbackReviewers()([]AccessReviewReviewerScopeable)
    GetInstanceEnumerationScope()(AccessReviewScopeable)
    GetInstances()([]AccessReviewInstanceable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewers()([]AccessReviewReviewerScopeable)
    GetScope()(AccessReviewScopeable)
    GetSettings()(AccessReviewScheduleSettingsable)
    GetStageSettings()([]AccessReviewStageSettingsable)
    GetStatus()(*string)
    SetAdditionalNotificationRecipients(value []AccessReviewNotificationRecipientItemable)()
    SetCreatedBy(value UserIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescriptionForAdmins(value *string)()
    SetDescriptionForReviewers(value *string)()
    SetDisplayName(value *string)()
    SetFallbackReviewers(value []AccessReviewReviewerScopeable)()
    SetInstanceEnumerationScope(value AccessReviewScopeable)()
    SetInstances(value []AccessReviewInstanceable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewers(value []AccessReviewReviewerScopeable)()
    SetScope(value AccessReviewScopeable)()
    SetSettings(value AccessReviewScheduleSettingsable)()
    SetStageSettings(value []AccessReviewStageSettingsable)()
    SetStatus(value *string)()
}
