package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentSetVersionItemable 
type DocumentSetVersionItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemId()(*string)
    GetOdataType()(*string)
    GetTitle()(*string)
    GetVersionId()(*string)
    SetItemId(value *string)()
    SetOdataType(value *string)()
    SetTitle(value *string)()
    SetVersionId(value *string)()
}
