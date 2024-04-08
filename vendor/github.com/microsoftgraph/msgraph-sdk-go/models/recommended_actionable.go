package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecommendedActionable 
type RecommendedActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionWebUrl()(*string)
    GetOdataType()(*string)
    GetPotentialScoreImpact()(*float64)
    GetTitle()(*string)
    SetActionWebUrl(value *string)()
    SetOdataType(value *string)()
    SetPotentialScoreImpact(value *float64)()
    SetTitle(value *string)()
}
