package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppPolicyDeploymentSummaryPerAppable 
type ManagedAppPolicyDeploymentSummaryPerAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationAppliedUserCount()(*int32)
    GetMobileAppIdentifier()(MobileAppIdentifierable)
    GetOdataType()(*string)
    SetConfigurationAppliedUserCount(value *int32)()
    SetMobileAppIdentifier(value MobileAppIdentifierable)()
    SetOdataType(value *string)()
}
