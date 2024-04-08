package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryReviewTagable 
type EdiscoveryReviewTagable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Tagable
    GetChildSelectability()(*ChildSelectability)
    GetChildTags()([]EdiscoveryReviewTagable)
    GetParent()(EdiscoveryReviewTagable)
    SetChildSelectability(value *ChildSelectability)()
    SetChildTags(value []EdiscoveryReviewTagable)()
    SetParent(value EdiscoveryReviewTagable)()
}
