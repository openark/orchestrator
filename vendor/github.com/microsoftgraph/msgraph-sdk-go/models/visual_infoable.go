package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VisualInfoable 
type VisualInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttribution()(ImageInfoable)
    GetBackgroundColor()(*string)
    GetContent()(Jsonable)
    GetDescription()(*string)
    GetDisplayText()(*string)
    GetOdataType()(*string)
    SetAttribution(value ImageInfoable)()
    SetBackgroundColor(value *string)()
    SetContent(value Jsonable)()
    SetDescription(value *string)()
    SetDisplayText(value *string)()
    SetOdataType(value *string)()
}
