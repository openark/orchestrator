package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AggregationOptionable 
type AggregationOptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBucketDefinition()(BucketAggregationDefinitionable)
    GetField()(*string)
    GetOdataType()(*string)
    GetSize()(*int32)
    SetBucketDefinition(value BucketAggregationDefinitionable)()
    SetField(value *string)()
    SetOdataType(value *string)()
    SetSize(value *int32)()
}
