package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedPCConfigurationable 
type SharedPCConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountManagerPolicy()(SharedPCAccountManagerPolicyable)
    GetAllowedAccounts()(*SharedPCAllowedAccountType)
    GetAllowLocalStorage()(*bool)
    GetDisableAccountManager()(*bool)
    GetDisableEduPolicies()(*bool)
    GetDisablePowerPolicies()(*bool)
    GetDisableSignInOnResume()(*bool)
    GetEnabled()(*bool)
    GetIdleTimeBeforeSleepInSeconds()(*int32)
    GetKioskAppDisplayName()(*string)
    GetKioskAppUserModelId()(*string)
    GetMaintenanceStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    SetAccountManagerPolicy(value SharedPCAccountManagerPolicyable)()
    SetAllowedAccounts(value *SharedPCAllowedAccountType)()
    SetAllowLocalStorage(value *bool)()
    SetDisableAccountManager(value *bool)()
    SetDisableEduPolicies(value *bool)()
    SetDisablePowerPolicies(value *bool)()
    SetDisableSignInOnResume(value *bool)()
    SetEnabled(value *bool)()
    SetIdleTimeBeforeSleepInSeconds(value *int32)()
    SetKioskAppDisplayName(value *string)()
    SetKioskAppUserModelId(value *string)()
    SetMaintenanceStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
}
