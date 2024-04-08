package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchAggregationable 
type SearchAggregationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuckets()([]SearchBucketable)
    GetField()(*string)
    GetOdataType()(*string)
    SetBuckets(value []SearchBucketable)()
    SetField(value *string)()
    SetOdataType(value *string)()
}
