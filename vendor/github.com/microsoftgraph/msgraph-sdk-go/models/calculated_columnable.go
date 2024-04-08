package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalculatedColumnable 
type CalculatedColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(*string)
    GetFormula()(*string)
    GetOdataType()(*string)
    GetOutputType()(*string)
    SetFormat(value *string)()
    SetFormula(value *string)()
    SetOdataType(value *string)()
    SetOutputType(value *string)()
}
