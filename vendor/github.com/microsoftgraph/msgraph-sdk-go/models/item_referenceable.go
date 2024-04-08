package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemReferenceable 
type ItemReferenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDriveId()(*string)
    GetDriveType()(*string)
    GetId()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetPath()(*string)
    GetShareId()(*string)
    GetSharepointIds()(SharepointIdsable)
    GetSiteId()(*string)
    SetDriveId(value *string)()
    SetDriveType(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPath(value *string)()
    SetShareId(value *string)()
    SetSharepointIds(value SharepointIdsable)()
    SetSiteId(value *string)()
}
