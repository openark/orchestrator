package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedDriveItemable 
type SharedDriveItemable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDriveItem()(DriveItemable)
    GetItems()([]DriveItemable)
    GetList()(Listable)
    GetListItem()(ListItemable)
    GetOwner()(IdentitySetable)
    GetPermission()(Permissionable)
    GetRoot()(DriveItemable)
    GetSite()(Siteable)
    SetDriveItem(value DriveItemable)()
    SetItems(value []DriveItemable)()
    SetList(value Listable)()
    SetListItem(value ListItemable)()
    SetOwner(value IdentitySetable)()
    SetPermission(value Permissionable)()
    SetRoot(value DriveItemable)()
    SetSite(value Siteable)()
}
