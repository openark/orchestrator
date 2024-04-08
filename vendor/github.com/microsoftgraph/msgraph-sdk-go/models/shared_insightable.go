package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedInsightable 
type SharedInsightable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastShared()(SharingDetailable)
    GetLastSharedMethod()(Entityable)
    GetResource()(Entityable)
    GetResourceReference()(ResourceReferenceable)
    GetResourceVisualization()(ResourceVisualizationable)
    GetSharingHistory()([]SharingDetailable)
    SetLastShared(value SharingDetailable)()
    SetLastSharedMethod(value Entityable)()
    SetResource(value Entityable)()
    SetResourceReference(value ResourceReferenceable)()
    SetResourceVisualization(value ResourceVisualizationable)()
    SetSharingHistory(value []SharingDetailable)()
}
