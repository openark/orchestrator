package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateWindowsDeviceAccountActionParameterable 
type UpdateWindowsDeviceAccountActionParameterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalendarSyncEnabled()(*bool)
    GetDeviceAccount()(WindowsDeviceAccountable)
    GetDeviceAccountEmail()(*string)
    GetExchangeServer()(*string)
    GetOdataType()(*string)
    GetPasswordRotationEnabled()(*bool)
    GetSessionInitiationProtocalAddress()(*string)
    SetCalendarSyncEnabled(value *bool)()
    SetDeviceAccount(value WindowsDeviceAccountable)()
    SetDeviceAccountEmail(value *string)()
    SetExchangeServer(value *string)()
    SetOdataType(value *string)()
    SetPasswordRotationEnabled(value *bool)()
    SetSessionInitiationProtocalAddress(value *string)()
}
