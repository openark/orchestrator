package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppliedConditionalAccessPolicyable 
type AppliedConditionalAccessPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetEnforcedGrantControls()([]string)
    GetEnforcedSessionControls()([]string)
    GetId()(*string)
    GetOdataType()(*string)
    GetResult()(*AppliedConditionalAccessPolicyResult)
    SetDisplayName(value *string)()
    SetEnforcedGrantControls(value []string)()
    SetEnforcedSessionControls(value []string)()
    SetId(value *string)()
    SetOdataType(value *string)()
    SetResult(value *AppliedConditionalAccessPolicyResult)()
}
