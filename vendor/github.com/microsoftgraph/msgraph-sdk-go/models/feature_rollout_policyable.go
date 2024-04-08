package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureRolloutPolicyable 
type FeatureRolloutPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppliesTo()([]DirectoryObjectable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFeature()(*StagedFeatureName)
    GetIsAppliedToOrganization()(*bool)
    GetIsEnabled()(*bool)
    SetAppliesTo(value []DirectoryObjectable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFeature(value *StagedFeatureName)()
    SetIsAppliedToOrganization(value *bool)()
    SetIsEnabled(value *bool)()
}
