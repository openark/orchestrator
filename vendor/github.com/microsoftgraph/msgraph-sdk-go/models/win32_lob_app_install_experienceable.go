package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppInstallExperienceable 
type Win32LobAppInstallExperienceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceRestartBehavior()(*Win32LobAppRestartBehavior)
    GetOdataType()(*string)
    GetRunAsAccount()(*RunAsAccountType)
    SetDeviceRestartBehavior(value *Win32LobAppRestartBehavior)()
    SetOdataType(value *string)()
    SetRunAsAccount(value *RunAsAccountType)()
}
