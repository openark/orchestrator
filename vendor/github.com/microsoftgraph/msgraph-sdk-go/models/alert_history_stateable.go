package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertHistoryStateable 
type AlertHistoryStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetAssignedTo()(*string)
    GetComments()([]string)
    GetFeedback()(*AlertFeedback)
    GetOdataType()(*string)
    GetStatus()(*AlertStatus)
    GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUser()(*string)
    SetAppId(value *string)()
    SetAssignedTo(value *string)()
    SetComments(value []string)()
    SetFeedback(value *AlertFeedback)()
    SetOdataType(value *string)()
    SetStatus(value *AlertStatus)()
    SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUser(value *string)()
}
