package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookTableSortable 
type WorkbookTableSortable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFields()([]WorkbookSortFieldable)
    GetMatchCase()(*bool)
    GetMethod()(*string)
    SetFields(value []WorkbookSortFieldable)()
    SetMatchCase(value *bool)()
    SetMethod(value *string)()
}
