package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ThumbnailSetable 
type ThumbnailSetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLarge()(Thumbnailable)
    GetMedium()(Thumbnailable)
    GetSmall()(Thumbnailable)
    GetSource()(Thumbnailable)
    SetLarge(value Thumbnailable)()
    SetMedium(value Thumbnailable)()
    SetSmall(value Thumbnailable)()
    SetSource(value Thumbnailable)()
}
