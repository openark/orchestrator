package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImageInfoable 
type ImageInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddImageQuery()(*bool)
    GetAlternateText()(*string)
    GetAlternativeText()(*string)
    GetIconUrl()(*string)
    GetOdataType()(*string)
    SetAddImageQuery(value *bool)()
    SetAlternateText(value *string)()
    SetAlternativeText(value *string)()
    SetIconUrl(value *string)()
    SetOdataType(value *string)()
}
