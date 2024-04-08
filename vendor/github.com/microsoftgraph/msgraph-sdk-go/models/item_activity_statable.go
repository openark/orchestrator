package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActivityStatable 
type ItemActivityStatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(ItemActionStatable)
    GetActivities()([]ItemActivityable)
    GetCreate()(ItemActionStatable)
    GetDelete()(ItemActionStatable)
    GetEdit()(ItemActionStatable)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIncompleteData()(IncompleteDataable)
    GetIsTrending()(*bool)
    GetMove()(ItemActionStatable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccess(value ItemActionStatable)()
    SetActivities(value []ItemActivityable)()
    SetCreate(value ItemActionStatable)()
    SetDelete(value ItemActionStatable)()
    SetEdit(value ItemActionStatable)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIncompleteData(value IncompleteDataable)()
    SetIsTrending(value *bool)()
    SetMove(value ItemActionStatable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
