package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SignInable 
type SignInable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicyable)
    GetClientAppUsed()(*string)
    GetConditionalAccessStatus()(*ConditionalAccessStatus)
    GetCorrelationId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceDetail()(DeviceDetailable)
    GetIpAddress()(*string)
    GetIsInteractive()(*bool)
    GetLocation()(SignInLocationable)
    GetResourceDisplayName()(*string)
    GetResourceId()(*string)
    GetRiskDetail()(*RiskDetail)
    GetRiskEventTypes()([]RiskEventType)
    GetRiskEventTypesV2()([]string)
    GetRiskLevelAggregated()(*RiskLevel)
    GetRiskLevelDuringSignIn()(*RiskLevel)
    GetRiskState()(*RiskState)
    GetStatus()(SignInStatusable)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicyable)()
    SetClientAppUsed(value *string)()
    SetConditionalAccessStatus(value *ConditionalAccessStatus)()
    SetCorrelationId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceDetail(value DeviceDetailable)()
    SetIpAddress(value *string)()
    SetIsInteractive(value *bool)()
    SetLocation(value SignInLocationable)()
    SetResourceDisplayName(value *string)()
    SetResourceId(value *string)()
    SetRiskDetail(value *RiskDetail)()
    SetRiskEventTypes(value []RiskEventType)()
    SetRiskEventTypesV2(value []string)()
    SetRiskLevelAggregated(value *RiskLevel)()
    SetRiskLevelDuringSignIn(value *RiskLevel)()
    SetRiskState(value *RiskState)()
    SetStatus(value SignInStatusable)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
