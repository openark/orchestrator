package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ThreatAssessmentRequestable 
type ThreatAssessmentRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*ThreatCategory)
    GetContentType()(*ThreatAssessmentContentType)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpectedAssessment()(*ThreatExpectedAssessment)
    GetRequestSource()(*ThreatAssessmentRequestSource)
    GetResults()([]ThreatAssessmentResultable)
    GetStatus()(*ThreatAssessmentStatus)
    SetCategory(value *ThreatCategory)()
    SetContentType(value *ThreatAssessmentContentType)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpectedAssessment(value *ThreatExpectedAssessment)()
    SetRequestSource(value *ThreatAssessmentRequestSource)()
    SetResults(value []ThreatAssessmentResultable)()
    SetStatus(value *ThreatAssessmentStatus)()
}
