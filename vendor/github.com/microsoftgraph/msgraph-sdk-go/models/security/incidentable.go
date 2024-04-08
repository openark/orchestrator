package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Incidentable 
type Incidentable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlerts()([]Alertable)
    GetAssignedTo()(*string)
    GetClassification()(*AlertClassification)
    GetComments()([]AlertCommentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomTags()([]string)
    GetDetermination()(*AlertDetermination)
    GetDisplayName()(*string)
    GetIncidentWebUrl()(*string)
    GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRedirectIncidentId()(*string)
    GetSeverity()(*AlertSeverity)
    GetStatus()(*IncidentStatus)
    GetTenantId()(*string)
    SetAlerts(value []Alertable)()
    SetAssignedTo(value *string)()
    SetClassification(value *AlertClassification)()
    SetComments(value []AlertCommentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomTags(value []string)()
    SetDetermination(value *AlertDetermination)()
    SetDisplayName(value *string)()
    SetIncidentWebUrl(value *string)()
    SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRedirectIncidentId(value *string)()
    SetSeverity(value *AlertSeverity)()
    SetStatus(value *IncidentStatus)()
    SetTenantId(value *string)()
}
