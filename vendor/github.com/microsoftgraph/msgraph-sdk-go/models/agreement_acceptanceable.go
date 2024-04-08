package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgreementAcceptanceable 
type AgreementAcceptanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgreementFileId()(*string)
    GetAgreementId()(*string)
    GetDeviceDisplayName()(*string)
    GetDeviceId()(*string)
    GetDeviceOSType()(*string)
    GetDeviceOSVersion()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetState()(*AgreementAcceptanceState)
    GetUserDisplayName()(*string)
    GetUserEmail()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetAgreementFileId(value *string)()
    SetAgreementId(value *string)()
    SetDeviceDisplayName(value *string)()
    SetDeviceId(value *string)()
    SetDeviceOSType(value *string)()
    SetDeviceOSVersion(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetState(value *AgreementAcceptanceState)()
    SetUserDisplayName(value *string)()
    SetUserEmail(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
