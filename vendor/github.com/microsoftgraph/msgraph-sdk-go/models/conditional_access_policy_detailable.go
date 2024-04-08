package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessPolicyDetailable 
type ConditionalAccessPolicyDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConditions()(ConditionalAccessConditionSetable)
    GetGrantControls()(ConditionalAccessGrantControlsable)
    GetOdataType()(*string)
    GetSessionControls()(ConditionalAccessSessionControlsable)
    SetConditions(value ConditionalAccessConditionSetable)()
    SetGrantControls(value ConditionalAccessGrantControlsable)()
    SetOdataType(value *string)()
    SetSessionControls(value ConditionalAccessSessionControlsable)()
}
