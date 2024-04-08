package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodFeatureConfigurationable 
type AuthenticationMethodFeatureConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeTarget()(FeatureTargetable)
    GetIncludeTarget()(FeatureTargetable)
    GetOdataType()(*string)
    GetState()(*AdvancedConfigState)
    SetExcludeTarget(value FeatureTargetable)()
    SetIncludeTarget(value FeatureTargetable)()
    SetOdataType(value *string)()
    SetState(value *AdvancedConfigState)()
}
