package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodsRegistrationCampaignable 
type AuthenticationMethodsRegistrationCampaignable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeTargets()([]ExcludeTargetable)
    GetIncludeTargets()([]AuthenticationMethodsRegistrationCampaignIncludeTargetable)
    GetOdataType()(*string)
    GetSnoozeDurationInDays()(*int32)
    GetState()(*AdvancedConfigState)
    SetExcludeTargets(value []ExcludeTargetable)()
    SetIncludeTargets(value []AuthenticationMethodsRegistrationCampaignIncludeTargetable)()
    SetOdataType(value *string)()
    SetSnoozeDurationInDays(value *int32)()
    SetState(value *AdvancedConfigState)()
}
