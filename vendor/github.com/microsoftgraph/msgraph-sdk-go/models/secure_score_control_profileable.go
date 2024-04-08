package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecureScoreControlProfileable 
type SecureScoreControlProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionType()(*string)
    GetActionUrl()(*string)
    GetAzureTenantId()(*string)
    GetComplianceInformation()([]ComplianceInformationable)
    GetControlCategory()(*string)
    GetControlStateUpdates()([]SecureScoreControlStateUpdateable)
    GetDeprecated()(*bool)
    GetImplementationCost()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaxScore()(*float64)
    GetRank()(*int32)
    GetRemediation()(*string)
    GetRemediationImpact()(*string)
    GetService()(*string)
    GetThreats()([]string)
    GetTier()(*string)
    GetTitle()(*string)
    GetUserImpact()(*string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetActionType(value *string)()
    SetActionUrl(value *string)()
    SetAzureTenantId(value *string)()
    SetComplianceInformation(value []ComplianceInformationable)()
    SetControlCategory(value *string)()
    SetControlStateUpdates(value []SecureScoreControlStateUpdateable)()
    SetDeprecated(value *bool)()
    SetImplementationCost(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaxScore(value *float64)()
    SetRank(value *int32)()
    SetRemediation(value *string)()
    SetRemediationImpact(value *string)()
    SetService(value *string)()
    SetThreats(value []string)()
    SetTier(value *string)()
    SetTitle(value *string)()
    SetUserImpact(value *string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
