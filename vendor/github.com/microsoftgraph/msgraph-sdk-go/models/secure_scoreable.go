package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecureScoreable 
type SecureScoreable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveUserCount()(*int32)
    GetAverageComparativeScores()([]AverageComparativeScoreable)
    GetAzureTenantId()(*string)
    GetControlScores()([]ControlScoreable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentScore()(*float64)
    GetEnabledServices()([]string)
    GetLicensedUserCount()(*int32)
    GetMaxScore()(*float64)
    GetVendorInformation()(SecurityVendorInformationable)
    SetActiveUserCount(value *int32)()
    SetAverageComparativeScores(value []AverageComparativeScoreable)()
    SetAzureTenantId(value *string)()
    SetControlScores(value []ControlScoreable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentScore(value *float64)()
    SetEnabledServices(value []string)()
    SetLicensedUserCount(value *int32)()
    SetMaxScore(value *float64)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
