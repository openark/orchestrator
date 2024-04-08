package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageAttachmentable 
type ChatMessageAttachmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(*string)
    GetContentType()(*string)
    GetContentUrl()(*string)
    GetId()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetTeamsAppId()(*string)
    GetThumbnailUrl()(*string)
    SetContent(value *string)()
    SetContentType(value *string)()
    SetContentUrl(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetTeamsAppId(value *string)()
    SetThumbnailUrl(value *string)()
}
