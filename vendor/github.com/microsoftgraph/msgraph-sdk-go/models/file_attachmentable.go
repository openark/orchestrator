package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileAttachmentable 
type FileAttachmentable interface {
    Attachmentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentBytes()([]byte)
    GetContentId()(*string)
    GetContentLocation()(*string)
    SetContentBytes(value []byte)()
    SetContentId(value *string)()
    SetContentLocation(value *string)()
}
