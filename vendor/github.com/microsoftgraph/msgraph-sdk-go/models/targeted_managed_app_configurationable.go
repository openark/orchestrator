package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetedManagedAppConfigurationable 
type TargetedManagedAppConfigurationable interface {
    ManagedAppConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApps()([]ManagedMobileAppable)
    GetAssignments()([]TargetedManagedAppPolicyAssignmentable)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetIsAssigned()(*bool)
    SetApps(value []ManagedMobileAppable)()
    SetAssignments(value []TargetedManagedAppPolicyAssignmentable)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetIsAssigned(value *bool)()
}
