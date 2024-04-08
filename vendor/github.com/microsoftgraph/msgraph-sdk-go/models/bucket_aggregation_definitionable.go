package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BucketAggregationDefinitionable 
type BucketAggregationDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsDescending()(*bool)
    GetMinimumCount()(*int32)
    GetOdataType()(*string)
    GetPrefixFilter()(*string)
    GetRanges()([]BucketAggregationRangeable)
    GetSortBy()(*BucketAggregationSortProperty)
    SetIsDescending(value *bool)()
    SetMinimumCount(value *int32)()
    SetOdataType(value *string)()
    SetPrefixFilter(value *string)()
    SetRanges(value []BucketAggregationRangeable)()
    SetSortBy(value *BucketAggregationSortProperty)()
}
