package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseUnitsDetailable 
type LicenseUnitsDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*int32)
    GetOdataType()(*string)
    GetSuspended()(*int32)
    GetWarning()(*int32)
    SetEnabled(value *int32)()
    SetOdataType(value *string)()
    SetSuspended(value *int32)()
    SetWarning(value *int32)()
}
