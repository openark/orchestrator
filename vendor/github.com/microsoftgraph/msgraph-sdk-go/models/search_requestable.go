package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchRequestable 
type SearchRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAggregationFilters()([]string)
    GetAggregations()([]AggregationOptionable)
    GetContentSources()([]string)
    GetEnableTopResults()(*bool)
    GetEntityTypes()([]EntityType)
    GetFields()([]string)
    GetFrom()(*int32)
    GetOdataType()(*string)
    GetQuery()(SearchQueryable)
    GetQueryAlterationOptions()(SearchAlterationOptionsable)
    GetRegion()(*string)
    GetResultTemplateOptions()(ResultTemplateOptionable)
    GetSharePointOneDriveOptions()(SharePointOneDriveOptionsable)
    GetSize()(*int32)
    GetSortProperties()([]SortPropertyable)
    SetAggregationFilters(value []string)()
    SetAggregations(value []AggregationOptionable)()
    SetContentSources(value []string)()
    SetEnableTopResults(value *bool)()
    SetEntityTypes(value []EntityType)()
    SetFields(value []string)()
    SetFrom(value *int32)()
    SetOdataType(value *string)()
    SetQuery(value SearchQueryable)()
    SetQueryAlterationOptions(value SearchAlterationOptionsable)()
    SetRegion(value *string)()
    SetResultTemplateOptions(value ResultTemplateOptionable)()
    SetSharePointOneDriveOptions(value SharePointOneDriveOptionsable)()
    SetSize(value *int32)()
    SetSortProperties(value []SortPropertyable)()
}
