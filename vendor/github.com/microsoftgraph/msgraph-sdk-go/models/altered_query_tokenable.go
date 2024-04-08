package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlteredQueryTokenable 
type AlteredQueryTokenable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLength()(*int32)
    GetOdataType()(*string)
    GetOffset()(*int32)
    GetSuggestion()(*string)
    SetLength(value *int32)()
    SetOdataType(value *string)()
    SetOffset(value *int32)()
    SetSuggestion(value *string)()
}
