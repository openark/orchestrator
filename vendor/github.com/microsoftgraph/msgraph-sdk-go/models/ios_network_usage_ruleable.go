package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosNetworkUsageRuleable 
type IosNetworkUsageRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCellularDataBlocked()(*bool)
    GetCellularDataBlockWhenRoaming()(*bool)
    GetManagedApps()([]AppListItemable)
    GetOdataType()(*string)
    SetCellularDataBlocked(value *bool)()
    SetCellularDataBlockWhenRoaming(value *bool)()
    SetManagedApps(value []AppListItemable)()
    SetOdataType(value *string)()
}
