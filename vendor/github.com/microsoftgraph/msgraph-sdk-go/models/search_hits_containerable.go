package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchHitsContainerable 
type SearchHitsContainerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAggregations()([]SearchAggregationable)
    GetHits()([]SearchHitable)
    GetMoreResultsAvailable()(*bool)
    GetOdataType()(*string)
    GetTotal()(*int32)
    SetAggregations(value []SearchAggregationable)()
    SetHits(value []SearchHitable)()
    SetMoreResultsAvailable(value *bool)()
    SetOdataType(value *string)()
    SetTotal(value *int32)()
}
