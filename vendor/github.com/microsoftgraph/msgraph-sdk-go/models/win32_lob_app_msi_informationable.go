package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppMsiInformationable 
type Win32LobAppMsiInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPackageType()(*Win32LobAppMsiPackageType)
    GetProductCode()(*string)
    GetProductName()(*string)
    GetProductVersion()(*string)
    GetPublisher()(*string)
    GetRequiresReboot()(*bool)
    GetUpgradeCode()(*string)
    SetOdataType(value *string)()
    SetPackageType(value *Win32LobAppMsiPackageType)()
    SetProductCode(value *string)()
    SetProductName(value *string)()
    SetProductVersion(value *string)()
    SetPublisher(value *string)()
    SetRequiresReboot(value *bool)()
    SetUpgradeCode(value *string)()
}
