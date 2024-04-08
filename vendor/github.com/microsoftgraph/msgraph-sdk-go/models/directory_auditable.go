package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryAuditable 
type DirectoryAuditable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityDisplayName()(*string)
    GetAdditionalDetails()([]KeyValueable)
    GetCategory()(*string)
    GetCorrelationId()(*string)
    GetInitiatedBy()(AuditActivityInitiatorable)
    GetLoggedByService()(*string)
    GetOperationType()(*string)
    GetResult()(*OperationResult)
    GetResultReason()(*string)
    GetTargetResources()([]TargetResourceable)
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityDisplayName(value *string)()
    SetAdditionalDetails(value []KeyValueable)()
    SetCategory(value *string)()
    SetCorrelationId(value *string)()
    SetInitiatedBy(value AuditActivityInitiatorable)()
    SetLoggedByService(value *string)()
    SetOperationType(value *string)()
    SetResult(value *OperationResult)()
    SetResultReason(value *string)()
    SetTargetResources(value []TargetResourceable)()
}
