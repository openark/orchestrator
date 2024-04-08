package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintUsageable 
type PrintUsageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedBlackAndWhiteJobCount()(*int64)
    GetCompletedColorJobCount()(*int64)
    GetIncompleteJobCount()(*int64)
    GetUsageDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetCompletedBlackAndWhiteJobCount(value *int64)()
    SetCompletedColorJobCount(value *int64)()
    SetIncompleteJobCount(value *int64)()
    SetUsageDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
