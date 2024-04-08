package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Driveable 
type Driveable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBundles()([]DriveItemable)
    GetDriveType()(*string)
    GetFollowing()([]DriveItemable)
    GetItems()([]DriveItemable)
    GetList()(Listable)
    GetOwner()(IdentitySetable)
    GetQuota()(Quotaable)
    GetRoot()(DriveItemable)
    GetSharePointIds()(SharepointIdsable)
    GetSpecial()([]DriveItemable)
    GetSystem()(SystemFacetable)
    SetBundles(value []DriveItemable)()
    SetDriveType(value *string)()
    SetFollowing(value []DriveItemable)()
    SetItems(value []DriveItemable)()
    SetList(value Listable)()
    SetOwner(value IdentitySetable)()
    SetQuota(value Quotaable)()
    SetRoot(value DriveItemable)()
    SetSharePointIds(value SharepointIdsable)()
    SetSpecial(value []DriveItemable)()
    SetSystem(value SystemFacetable)()
}
