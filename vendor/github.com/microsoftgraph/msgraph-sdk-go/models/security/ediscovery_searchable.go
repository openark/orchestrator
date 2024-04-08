package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoverySearchable 
type EdiscoverySearchable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Searchable
    GetAdditionalSources()([]DataSourceable)
    GetAddToReviewSetOperation()(EdiscoveryAddToReviewSetOperationable)
    GetCustodianSources()([]DataSourceable)
    GetDataSourceScopes()(*DataSourceScopes)
    GetLastEstimateStatisticsOperation()(EdiscoveryEstimateOperationable)
    GetNoncustodialSources()([]EdiscoveryNoncustodialDataSourceable)
    SetAdditionalSources(value []DataSourceable)()
    SetAddToReviewSetOperation(value EdiscoveryAddToReviewSetOperationable)()
    SetCustodianSources(value []DataSourceable)()
    SetDataSourceScopes(value *DataSourceScopes)()
    SetLastEstimateStatisticsOperation(value EdiscoveryEstimateOperationable)()
    SetNoncustodialSources(value []EdiscoveryNoncustodialDataSourceable)()
}
