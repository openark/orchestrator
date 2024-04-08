package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskDetectionable 
type RiskDetectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(*ActivityType)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAdditionalInfo()(*string)
    GetCorrelationId()(*string)
    GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDetectionTimingType()(*RiskDetectionTimingType)
    GetIpAddress()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocation()(SignInLocationable)
    GetRequestId()(*string)
    GetRiskDetail()(*RiskDetail)
    GetRiskEventType()(*string)
    GetRiskLevel()(*RiskLevel)
    GetRiskState()(*RiskState)
    GetSource()(*string)
    GetTokenIssuerType()(*TokenIssuerType)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetActivity(value *ActivityType)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAdditionalInfo(value *string)()
    SetCorrelationId(value *string)()
    SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDetectionTimingType(value *RiskDetectionTimingType)()
    SetIpAddress(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocation(value SignInLocationable)()
    SetRequestId(value *string)()
    SetRiskDetail(value *RiskDetail)()
    SetRiskEventType(value *string)()
    SetRiskLevel(value *RiskLevel)()
    SetRiskState(value *RiskState)()
    SetSource(value *string)()
    SetTokenIssuerType(value *TokenIssuerType)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
