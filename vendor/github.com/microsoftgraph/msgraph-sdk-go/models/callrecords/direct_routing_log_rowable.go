package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectRoutingLogRowable 
type DirectRoutingLogRowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalleeNumber()(*string)
    GetCallEndSubReason()(*int32)
    GetCallerNumber()(*string)
    GetCallType()(*string)
    GetCorrelationId()(*string)
    GetDuration()(*int32)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailureDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFinalSipCode()(*int32)
    GetFinalSipCodePhrase()(*string)
    GetId()(*string)
    GetInviteDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMediaBypassEnabled()(*bool)
    GetMediaPathLocation()(*string)
    GetOdataType()(*string)
    GetSignalingLocation()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSuccessfulCall()(*bool)
    GetTrunkFullyQualifiedDomainName()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetCalleeNumber(value *string)()
    SetCallEndSubReason(value *int32)()
    SetCallerNumber(value *string)()
    SetCallType(value *string)()
    SetCorrelationId(value *string)()
    SetDuration(value *int32)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailureDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFinalSipCode(value *int32)()
    SetFinalSipCodePhrase(value *string)()
    SetId(value *string)()
    SetInviteDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMediaBypassEnabled(value *bool)()
    SetMediaPathLocation(value *string)()
    SetOdataType(value *string)()
    SetSignalingLocation(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSuccessfulCall(value *bool)()
    SetTrunkFullyQualifiedDomainName(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
