package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Shiftable 
type Shiftable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDraftShift()(ShiftItemable)
    GetSchedulingGroupId()(*string)
    GetSharedShift()(ShiftItemable)
    GetUserId()(*string)
    SetDraftShift(value ShiftItemable)()
    SetSchedulingGroupId(value *string)()
    SetSharedShift(value ShiftItemable)()
    SetUserId(value *string)()
}
