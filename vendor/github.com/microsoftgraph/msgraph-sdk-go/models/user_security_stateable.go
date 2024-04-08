package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSecurityStateable 
type UserSecurityStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAadUserId()(*string)
    GetAccountName()(*string)
    GetDomainName()(*string)
    GetEmailRole()(*EmailRole)
    GetIsVpn()(*bool)
    GetLogonDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogonId()(*string)
    GetLogonIp()(*string)
    GetLogonLocation()(*string)
    GetLogonType()(*LogonType)
    GetOdataType()(*string)
    GetOnPremisesSecurityIdentifier()(*string)
    GetRiskScore()(*string)
    GetUserAccountType()(*UserAccountSecurityType)
    GetUserPrincipalName()(*string)
    SetAadUserId(value *string)()
    SetAccountName(value *string)()
    SetDomainName(value *string)()
    SetEmailRole(value *EmailRole)()
    SetIsVpn(value *bool)()
    SetLogonDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogonId(value *string)()
    SetLogonIp(value *string)()
    SetLogonLocation(value *string)()
    SetLogonType(value *LogonType)()
    SetOdataType(value *string)()
    SetOnPremisesSecurityIdentifier(value *string)()
    SetRiskScore(value *string)()
    SetUserAccountType(value *UserAccountSecurityType)()
    SetUserPrincipalName(value *string)()
}
