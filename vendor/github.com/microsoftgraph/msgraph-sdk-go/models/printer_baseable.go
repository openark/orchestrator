package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterBaseable 
type PrinterBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapabilities()(PrinterCapabilitiesable)
    GetDefaults()(PrinterDefaultsable)
    GetDisplayName()(*string)
    GetIsAcceptingJobs()(*bool)
    GetJobs()([]PrintJobable)
    GetLocation()(PrinterLocationable)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetStatus()(PrinterStatusable)
    SetCapabilities(value PrinterCapabilitiesable)()
    SetDefaults(value PrinterDefaultsable)()
    SetDisplayName(value *string)()
    SetIsAcceptingJobs(value *bool)()
    SetJobs(value []PrintJobable)()
    SetLocation(value PrinterLocationable)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetStatus(value PrinterStatusable)()
}
