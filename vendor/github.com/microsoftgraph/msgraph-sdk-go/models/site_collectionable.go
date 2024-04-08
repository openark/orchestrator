package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SiteCollectionable 
type SiteCollectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataLocationCode()(*string)
    GetHostname()(*string)
    GetOdataType()(*string)
    GetRoot()(Rootable)
    SetDataLocationCode(value *string)()
    SetHostname(value *string)()
    SetOdataType(value *string)()
    SetRoot(value Rootable)()
}
