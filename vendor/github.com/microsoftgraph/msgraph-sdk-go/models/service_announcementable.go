package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncementable 
type ServiceAnnouncementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHealthOverviews()([]ServiceHealthable)
    GetIssues()([]ServiceHealthIssueable)
    GetMessages()([]ServiceUpdateMessageable)
    SetHealthOverviews(value []ServiceHealthable)()
    SetIssues(value []ServiceHealthIssueable)()
    SetMessages(value []ServiceUpdateMessageable)()
}
