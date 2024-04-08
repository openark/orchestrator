package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookTableColumnable 
type WorkbookTableColumnable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFilter()(WorkbookFilterable)
    GetIndex()(*int32)
    GetName()(*string)
    GetValues()(Jsonable)
    SetFilter(value WorkbookFilterable)()
    SetIndex(value *int32)()
    SetName(value *string)()
    SetValues(value Jsonable)()
}
