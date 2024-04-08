package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserTrainingContentEventInfoable 
type UserTrainingContentEventInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrowser()(*string)
    GetContentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIpAddress()(*string)
    GetOdataType()(*string)
    GetOsPlatformDeviceDetails()(*string)
    GetPotentialScoreImpact()(*float64)
    SetBrowser(value *string)()
    SetContentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIpAddress(value *string)()
    SetOdataType(value *string)()
    SetOsPlatformDeviceDetails(value *string)()
    SetPotentialScoreImpact(value *float64)()
}
