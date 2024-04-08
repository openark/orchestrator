package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookGeoCoordinatesable 
type OutlookGeoCoordinatesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccuracy()(*float64)
    GetAltitude()(*float64)
    GetAltitudeAccuracy()(*float64)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    GetOdataType()(*string)
    SetAccuracy(value *float64)()
    SetAltitude(value *float64)()
    SetAltitudeAccuracy(value *float64)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
    SetOdataType(value *string)()
}
