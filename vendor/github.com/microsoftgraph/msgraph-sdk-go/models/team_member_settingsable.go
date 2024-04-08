package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamMemberSettingsable 
type TeamMemberSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAddRemoveApps()(*bool)
    GetAllowCreatePrivateChannels()(*bool)
    GetAllowCreateUpdateChannels()(*bool)
    GetAllowCreateUpdateRemoveConnectors()(*bool)
    GetAllowCreateUpdateRemoveTabs()(*bool)
    GetAllowDeleteChannels()(*bool)
    GetOdataType()(*string)
    SetAllowAddRemoveApps(value *bool)()
    SetAllowCreatePrivateChannels(value *bool)()
    SetAllowCreateUpdateChannels(value *bool)()
    SetAllowCreateUpdateRemoveConnectors(value *bool)()
    SetAllowCreateUpdateRemoveTabs(value *bool)()
    SetAllowDeleteChannels(value *bool)()
    SetOdataType(value *string)()
}
