package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaStreamable 
type MediaStreamable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDirection()(*MediaDirection)
    GetLabel()(*string)
    GetMediaType()(*Modality)
    GetOdataType()(*string)
    GetServerMuted()(*bool)
    GetSourceId()(*string)
    SetDirection(value *MediaDirection)()
    SetLabel(value *string)()
    SetMediaType(value *Modality)()
    SetOdataType(value *string)()
    SetServerMuted(value *bool)()
    SetSourceId(value *string)()
}
