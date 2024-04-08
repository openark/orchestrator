package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestable 
type SubjectRightsRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTo()(Identityable)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDataSubject()(DataSubjectable)
    GetDataSubjectType()(*DataSubjectType)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHistory()([]SubjectRightsRequestHistoryable)
    GetInsight()(SubjectRightsRequestDetailable)
    GetInternalDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotes()([]AuthoredNoteable)
    GetRegulations()([]string)
    GetStages()([]SubjectRightsRequestStageDetailable)
    GetStatus()(*SubjectRightsRequestStatus)
    GetTeam()(Teamable)
    GetType()(*SubjectRightsRequestType)
    SetAssignedTo(value Identityable)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDataSubject(value DataSubjectable)()
    SetDataSubjectType(value *DataSubjectType)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHistory(value []SubjectRightsRequestHistoryable)()
    SetInsight(value SubjectRightsRequestDetailable)()
    SetInternalDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotes(value []AuthoredNoteable)()
    SetRegulations(value []string)()
    SetStages(value []SubjectRightsRequestStageDetailable)()
    SetStatus(value *SubjectRightsRequestStatus)()
    SetTeam(value Teamable)()
    SetType(value *SubjectRightsRequestType)()
}
