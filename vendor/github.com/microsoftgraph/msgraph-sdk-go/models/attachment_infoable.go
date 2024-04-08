package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttachmentInfoable 
type AttachmentInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachmentType()(*AttachmentType)
    GetContentType()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetSize()(*int64)
    SetAttachmentType(value *AttachmentType)()
    SetContentType(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetSize(value *int64)()
}
