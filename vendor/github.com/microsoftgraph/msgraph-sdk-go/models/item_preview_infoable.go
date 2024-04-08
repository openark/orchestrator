package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPreviewInfoable 
type ItemPreviewInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGetUrl()(*string)
    GetOdataType()(*string)
    GetPostParameters()(*string)
    GetPostUrl()(*string)
    SetGetUrl(value *string)()
    SetOdataType(value *string)()
    SetPostParameters(value *string)()
    SetPostUrl(value *string)()
}
