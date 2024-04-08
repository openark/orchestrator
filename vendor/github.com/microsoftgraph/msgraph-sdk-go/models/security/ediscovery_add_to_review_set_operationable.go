package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryAddToReviewSetOperationable 
type EdiscoveryAddToReviewSetOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewSet()(EdiscoveryReviewSetable)
    GetSearch()(EdiscoverySearchable)
    SetReviewSet(value EdiscoveryReviewSetable)()
    SetSearch(value EdiscoverySearchable)()
}
