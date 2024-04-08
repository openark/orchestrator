package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosMinimumOperatingSystemable 
type IosMinimumOperatingSystemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetV100()(*bool)
    GetV110()(*bool)
    GetV120()(*bool)
    GetV130()(*bool)
    GetV140()(*bool)
    GetV150()(*bool)
    GetV80()(*bool)
    GetV90()(*bool)
    SetOdataType(value *string)()
    SetV100(value *bool)()
    SetV110(value *bool)()
    SetV120(value *bool)()
    SetV130(value *bool)()
    SetV140(value *bool)()
    SetV150(value *bool)()
    SetV80(value *bool)()
    SetV90(value *bool)()
}
