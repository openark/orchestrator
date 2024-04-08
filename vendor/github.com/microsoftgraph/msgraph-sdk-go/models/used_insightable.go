package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsedInsightable 
type UsedInsightable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastUsed()(UsageDetailsable)
    GetResource()(Entityable)
    GetResourceReference()(ResourceReferenceable)
    GetResourceVisualization()(ResourceVisualizationable)
    SetLastUsed(value UsageDetailsable)()
    SetResource(value Entityable)()
    SetResourceReference(value ResourceReferenceable)()
    SetResourceVisualization(value ResourceVisualizationable)()
}
