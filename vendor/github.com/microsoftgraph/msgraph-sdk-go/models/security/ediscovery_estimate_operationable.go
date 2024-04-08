package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryEstimateOperationable 
type EdiscoveryEstimateOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIndexedItemCount()(*int64)
    GetIndexedItemsSize()(*int64)
    GetMailboxCount()(*int32)
    GetSearch()(EdiscoverySearchable)
    GetSiteCount()(*int32)
    GetUnindexedItemCount()(*int64)
    GetUnindexedItemsSize()(*int64)
    SetIndexedItemCount(value *int64)()
    SetIndexedItemsSize(value *int64)()
    SetMailboxCount(value *int32)()
    SetSearch(value EdiscoverySearchable)()
    SetSiteCount(value *int32)()
    SetUnindexedItemCount(value *int64)()
    SetUnindexedItemsSize(value *int64)()
}
