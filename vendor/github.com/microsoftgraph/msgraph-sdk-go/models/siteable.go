package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Siteable 
type Siteable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalytics()(ItemAnalyticsable)
    GetColumns()([]ColumnDefinitionable)
    GetContentTypes()([]ContentTypeable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetDrives()([]Driveable)
    GetError()(PublicErrorable)
    GetExternalColumns()([]ColumnDefinitionable)
    GetItems()([]BaseItemable)
    GetLists()([]Listable)
    GetOnenote()(Onenoteable)
    GetOperations()([]RichLongRunningOperationable)
    GetPermissions()([]Permissionable)
    GetRoot()(Rootable)
    GetSharepointIds()(SharepointIdsable)
    GetSiteCollection()(SiteCollectionable)
    GetSites()([]Siteable)
    SetAnalytics(value ItemAnalyticsable)()
    SetColumns(value []ColumnDefinitionable)()
    SetContentTypes(value []ContentTypeable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetDrives(value []Driveable)()
    SetError(value PublicErrorable)()
    SetExternalColumns(value []ColumnDefinitionable)()
    SetItems(value []BaseItemable)()
    SetLists(value []Listable)()
    SetOnenote(value Onenoteable)()
    SetOperations(value []RichLongRunningOperationable)()
    SetPermissions(value []Permissionable)()
    SetRoot(value Rootable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSiteCollection(value SiteCollectionable)()
    SetSites(value []Siteable)()
}
