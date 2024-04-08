package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Printable 
type Printable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectors()([]PrintConnectorable)
    GetOdataType()(*string)
    GetOperations()([]PrintOperationable)
    GetPrinters()([]Printerable)
    GetServices()([]PrintServiceable)
    GetSettings()(PrintSettingsable)
    GetShares()([]PrinterShareable)
    GetTaskDefinitions()([]PrintTaskDefinitionable)
    SetConnectors(value []PrintConnectorable)()
    SetOdataType(value *string)()
    SetOperations(value []PrintOperationable)()
    SetPrinters(value []Printerable)()
    SetServices(value []PrintServiceable)()
    SetSettings(value PrintSettingsable)()
    SetShares(value []PrinterShareable)()
    SetTaskDefinitions(value []PrintTaskDefinitionable)()
}
