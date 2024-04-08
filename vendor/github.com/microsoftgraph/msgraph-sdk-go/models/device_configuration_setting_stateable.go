package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationSettingStateable 
type DeviceConfigurationSettingStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentValue()(*string)
    GetErrorCode()(*int64)
    GetErrorDescription()(*string)
    GetInstanceDisplayName()(*string)
    GetOdataType()(*string)
    GetSetting()(*string)
    GetSettingName()(*string)
    GetSources()([]SettingSourceable)
    GetState()(*ComplianceStatus)
    GetUserEmail()(*string)
    GetUserId()(*string)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetCurrentValue(value *string)()
    SetErrorCode(value *int64)()
    SetErrorDescription(value *string)()
    SetInstanceDisplayName(value *string)()
    SetOdataType(value *string)()
    SetSetting(value *string)()
    SetSettingName(value *string)()
    SetSources(value []SettingSourceable)()
    SetState(value *ComplianceStatus)()
    SetUserEmail(value *string)()
    SetUserId(value *string)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}
