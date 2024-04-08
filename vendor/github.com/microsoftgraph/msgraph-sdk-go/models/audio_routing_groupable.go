package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AudioRoutingGroupable 
type AudioRoutingGroupable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReceivers()([]string)
    GetRoutingMode()(*RoutingMode)
    GetSources()([]string)
    SetReceivers(value []string)()
    SetRoutingMode(value *RoutingMode)()
    SetSources(value []string)()
}
