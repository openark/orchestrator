package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceHealthIssueable 
type ServiceHealthIssueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceAnnouncementBaseable
    GetClassification()(*ServiceHealthClassificationType)
    GetFeature()(*string)
    GetFeatureGroup()(*string)
    GetImpactDescription()(*string)
    GetIsResolved()(*bool)
    GetOrigin()(*ServiceHealthOrigin)
    GetPosts()([]ServiceHealthIssuePostable)
    GetService()(*string)
    GetStatus()(*ServiceHealthStatus)
    SetClassification(value *ServiceHealthClassificationType)()
    SetFeature(value *string)()
    SetFeatureGroup(value *string)()
    SetImpactDescription(value *string)()
    SetIsResolved(value *bool)()
    SetOrigin(value *ServiceHealthOrigin)()
    SetPosts(value []ServiceHealthIssuePostable)()
    SetService(value *string)()
    SetStatus(value *ServiceHealthStatus)()
}
