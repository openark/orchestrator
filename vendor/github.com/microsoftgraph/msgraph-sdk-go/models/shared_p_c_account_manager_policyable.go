package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedPCAccountManagerPolicyable 
type SharedPCAccountManagerPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountDeletionPolicy()(*SharedPCAccountDeletionPolicyType)
    GetCacheAccountsAboveDiskFreePercentage()(*int32)
    GetInactiveThresholdDays()(*int32)
    GetOdataType()(*string)
    GetRemoveAccountsBelowDiskFreePercentage()(*int32)
    SetAccountDeletionPolicy(value *SharedPCAccountDeletionPolicyType)()
    SetCacheAccountsAboveDiskFreePercentage(value *int32)()
    SetInactiveThresholdDays(value *int32)()
    SetOdataType(value *string)()
    SetRemoveAccountsBelowDiskFreePercentage(value *int32)()
}
