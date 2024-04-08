package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSimulationDetailsable 
type UserSimulationDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTrainingsCount()(*int32)
    GetCompletedTrainingsCount()(*int32)
    GetCompromisedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInProgressTrainingsCount()(*int32)
    GetIsCompromised()(*bool)
    GetOdataType()(*string)
    GetReportedPhishDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSimulationEvents()([]UserSimulationEventInfoable)
    GetSimulationUser()(AttackSimulationUserable)
    GetTrainingEvents()([]UserTrainingEventInfoable)
    SetAssignedTrainingsCount(value *int32)()
    SetCompletedTrainingsCount(value *int32)()
    SetCompromisedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInProgressTrainingsCount(value *int32)()
    SetIsCompromised(value *bool)()
    SetOdataType(value *string)()
    SetReportedPhishDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSimulationEvents(value []UserSimulationEventInfoable)()
    SetSimulationUser(value AttackSimulationUserable)()
    SetTrainingEvents(value []UserTrainingEventInfoable)()
}
