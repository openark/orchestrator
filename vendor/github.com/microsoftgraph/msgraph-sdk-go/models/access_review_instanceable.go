package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceable 
type AccessReviewInstanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContactedReviewers()([]AccessReviewReviewerable)
    GetDecisions()([]AccessReviewInstanceDecisionItemable)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFallbackReviewers()([]AccessReviewReviewerScopeable)
    GetReviewers()([]AccessReviewReviewerScopeable)
    GetScope()(AccessReviewScopeable)
    GetStages()([]AccessReviewStageable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*string)
    SetContactedReviewers(value []AccessReviewReviewerable)()
    SetDecisions(value []AccessReviewInstanceDecisionItemable)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFallbackReviewers(value []AccessReviewReviewerScopeable)()
    SetReviewers(value []AccessReviewReviewerScopeable)()
    SetScope(value AccessReviewScopeable)()
    SetStages(value []AccessReviewStageable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *string)()
}
