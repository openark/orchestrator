package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharingLinkable 
type SharingLinkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(Identityable)
    GetOdataType()(*string)
    GetPreventsDownload()(*bool)
    GetScope()(*string)
    GetType()(*string)
    GetWebHtml()(*string)
    GetWebUrl()(*string)
    SetApplication(value Identityable)()
    SetOdataType(value *string)()
    SetPreventsDownload(value *bool)()
    SetScope(value *string)()
    SetType(value *string)()
    SetWebHtml(value *string)()
    SetWebUrl(value *string)()
}
