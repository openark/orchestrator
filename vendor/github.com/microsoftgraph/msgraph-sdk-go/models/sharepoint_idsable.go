package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharepointIdsable 
type SharepointIdsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetListId()(*string)
    GetListItemId()(*string)
    GetListItemUniqueId()(*string)
    GetOdataType()(*string)
    GetSiteId()(*string)
    GetSiteUrl()(*string)
    GetTenantId()(*string)
    GetWebId()(*string)
    SetListId(value *string)()
    SetListItemId(value *string)()
    SetListItemUniqueId(value *string)()
    SetOdataType(value *string)()
    SetSiteId(value *string)()
    SetSiteUrl(value *string)()
    SetTenantId(value *string)()
    SetWebId(value *string)()
}
