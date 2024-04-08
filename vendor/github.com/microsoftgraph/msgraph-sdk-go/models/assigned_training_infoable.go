package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignedTrainingInfoable 
type AssignedTrainingInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedUserCount()(*int32)
    GetCompletedUserCount()(*int32)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    SetAssignedUserCount(value *int32)()
    SetCompletedUserCount(value *int32)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
}
