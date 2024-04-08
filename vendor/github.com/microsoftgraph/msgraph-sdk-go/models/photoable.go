package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Photoable 
type Photoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCameraMake()(*string)
    GetCameraModel()(*string)
    GetExposureDenominator()(*float64)
    GetExposureNumerator()(*float64)
    GetFNumber()(*float64)
    GetFocalLength()(*float64)
    GetIso()(*int32)
    GetOdataType()(*string)
    GetOrientation()(*int32)
    GetTakenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCameraMake(value *string)()
    SetCameraModel(value *string)()
    SetExposureDenominator(value *float64)()
    SetExposureNumerator(value *float64)()
    SetFNumber(value *float64)()
    SetFocalLength(value *float64)()
    SetIso(value *int32)()
    SetOdataType(value *string)()
    SetOrientation(value *int32)()
    SetTakenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
