package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserTrainingEventInfoable 
type UserTrainingEventInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetLatestTrainingStatus()(*TrainingStatus)
    GetOdataType()(*string)
    GetTrainingAssignedProperties()(UserTrainingContentEventInfoable)
    GetTrainingCompletedProperties()(UserTrainingContentEventInfoable)
    GetTrainingUpdatedProperties()(UserTrainingContentEventInfoable)
    SetDisplayName(value *string)()
    SetLatestTrainingStatus(value *TrainingStatus)()
    SetOdataType(value *string)()
    SetTrainingAssignedProperties(value UserTrainingContentEventInfoable)()
    SetTrainingCompletedProperties(value UserTrainingContentEventInfoable)()
    SetTrainingUpdatedProperties(value UserTrainingContentEventInfoable)()
}
