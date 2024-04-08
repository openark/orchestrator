package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScoredEmailAddressable 
type ScoredEmailAddressable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetItemId()(*string)
    GetOdataType()(*string)
    GetRelevanceScore()(*float64)
    GetSelectionLikelihood()(*SelectionLikelihoodInfo)
    SetAddress(value *string)()
    SetItemId(value *string)()
    SetOdataType(value *string)()
    SetRelevanceScore(value *float64)()
    SetSelectionLikelihood(value *SelectionLikelihoodInfo)()
}
