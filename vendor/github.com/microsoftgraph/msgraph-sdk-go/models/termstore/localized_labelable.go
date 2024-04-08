package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocalizedLabelable 
type LocalizedLabelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsDefault()(*bool)
    GetLanguageTag()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    SetIsDefault(value *bool)()
    SetLanguageTag(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
}
