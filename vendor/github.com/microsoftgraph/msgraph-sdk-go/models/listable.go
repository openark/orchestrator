package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Listable 
type Listable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColumns()([]ColumnDefinitionable)
    GetContentTypes()([]ContentTypeable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetItems()([]ListItemable)
    GetList()(ListInfoable)
    GetOperations()([]RichLongRunningOperationable)
    GetSharepointIds()(SharepointIdsable)
    GetSubscriptions()([]Subscriptionable)
    GetSystem()(SystemFacetable)
    SetColumns(value []ColumnDefinitionable)()
    SetContentTypes(value []ContentTypeable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetItems(value []ListItemable)()
    SetList(value ListInfoable)()
    SetOperations(value []RichLongRunningOperationable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSubscriptions(value []Subscriptionable)()
    SetSystem(value SystemFacetable)()
}
