package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceVisualizationable 
type ResourceVisualizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainerDisplayName()(*string)
    GetContainerType()(*string)
    GetContainerWebUrl()(*string)
    GetMediaType()(*string)
    GetOdataType()(*string)
    GetPreviewImageUrl()(*string)
    GetPreviewText()(*string)
    GetTitle()(*string)
    GetType()(*string)
    SetContainerDisplayName(value *string)()
    SetContainerType(value *string)()
    SetContainerWebUrl(value *string)()
    SetMediaType(value *string)()
    SetOdataType(value *string)()
    SetPreviewImageUrl(value *string)()
    SetPreviewText(value *string)()
    SetTitle(value *string)()
    SetType(value *string)()
}
