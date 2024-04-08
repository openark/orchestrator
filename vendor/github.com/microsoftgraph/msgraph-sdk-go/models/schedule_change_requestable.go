package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScheduleChangeRequestable 
type ScheduleChangeRequestable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTo()(*ScheduleChangeRequestActor)
    GetManagerActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagerActionMessage()(*string)
    GetManagerUserId()(*string)
    GetSenderDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSenderMessage()(*string)
    GetSenderUserId()(*string)
    GetState()(*ScheduleChangeState)
    SetAssignedTo(value *ScheduleChangeRequestActor)()
    SetManagerActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagerActionMessage(value *string)()
    SetManagerUserId(value *string)()
    SetSenderDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSenderMessage(value *string)()
    SetSenderUserId(value *string)()
    SetState(value *ScheduleChangeState)()
}
