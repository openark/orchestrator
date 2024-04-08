package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceGeoLocationable 
type DeviceGeoLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAltitude()(*float64)
    GetHeading()(*float64)
    GetHorizontalAccuracy()(*float64)
    GetLastCollectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    GetOdataType()(*string)
    GetSpeed()(*float64)
    GetVerticalAccuracy()(*float64)
    SetAltitude(value *float64)()
    SetHeading(value *float64)()
    SetHorizontalAccuracy(value *float64)()
    SetLastCollectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
    SetOdataType(value *string)()
    SetSpeed(value *float64)()
    SetVerticalAccuracy(value *float64)()
}
