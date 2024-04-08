package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequestable 
type AccessPackageAssignmentRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAnswers()([]AccessPackageAnswerable)
    GetAssignment()(AccessPackageAssignmentable)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestor()(AccessPackageSubjectable)
    GetRequestType()(*AccessPackageRequestType)
    GetSchedule()(EntitlementManagementScheduleable)
    GetState()(*AccessPackageRequestState)
    GetStatus()(*string)
    SetAccessPackage(value AccessPackageable)()
    SetAnswers(value []AccessPackageAnswerable)()
    SetAssignment(value AccessPackageAssignmentable)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestor(value AccessPackageSubjectable)()
    SetRequestType(value *AccessPackageRequestType)()
    SetSchedule(value EntitlementManagementScheduleable)()
    SetState(value *AccessPackageRequestState)()
    SetStatus(value *string)()
}
