package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentSetContentable 
type DocumentSetContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentType()(ContentTypeInfoable)
    GetFileName()(*string)
    GetFolderName()(*string)
    GetOdataType()(*string)
    SetContentType(value ContentTypeInfoable)()
    SetFileName(value *string)()
    SetFolderName(value *string)()
    SetOdataType(value *string)()
}
