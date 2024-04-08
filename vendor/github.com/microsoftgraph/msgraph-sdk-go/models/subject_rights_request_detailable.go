package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestDetailable 
type SubjectRightsRequestDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludedItemCount()(*int64)
    GetInsightCounts()([]KeyValuePairable)
    GetItemCount()(*int64)
    GetItemNeedReview()(*int64)
    GetOdataType()(*string)
    GetProductItemCounts()([]KeyValuePairable)
    GetSignedOffItemCount()(*int64)
    GetTotalItemSize()(*int64)
    SetExcludedItemCount(value *int64)()
    SetInsightCounts(value []KeyValuePairable)()
    SetItemCount(value *int64)()
    SetItemNeedReview(value *int64)()
    SetOdataType(value *string)()
    SetProductItemCounts(value []KeyValuePairable)()
    SetSignedOffItemCount(value *int64)()
    SetTotalItemSize(value *int64)()
}
