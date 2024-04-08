package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskyUserable 
type RiskyUserable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHistory()([]RiskyUserHistoryItemable)
    GetIsDeleted()(*bool)
    GetIsProcessing()(*bool)
    GetRiskDetail()(*RiskDetail)
    GetRiskLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRiskLevel()(*RiskLevel)
    GetRiskState()(*RiskState)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetHistory(value []RiskyUserHistoryItemable)()
    SetIsDeleted(value *bool)()
    SetIsProcessing(value *bool)()
    SetRiskDetail(value *RiskDetail)()
    SetRiskLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRiskLevel(value *RiskLevel)()
    SetRiskState(value *RiskState)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
