package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRangeViewable 
type WorkbookRangeViewable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCellAddresses()(Jsonable)
    GetColumnCount()(*int32)
    GetFormulas()(Jsonable)
    GetFormulasLocal()(Jsonable)
    GetFormulasR1C1()(Jsonable)
    GetIndex()(*int32)
    GetNumberFormat()(Jsonable)
    GetRowCount()(*int32)
    GetRows()([]WorkbookRangeViewable)
    GetText()(Jsonable)
    GetValues()(Jsonable)
    GetValueTypes()(Jsonable)
    SetCellAddresses(value Jsonable)()
    SetColumnCount(value *int32)()
    SetFormulas(value Jsonable)()
    SetFormulasLocal(value Jsonable)()
    SetFormulasR1C1(value Jsonable)()
    SetIndex(value *int32)()
    SetNumberFormat(value Jsonable)()
    SetRowCount(value *int32)()
    SetRows(value []WorkbookRangeViewable)()
    SetText(value Jsonable)()
    SetValues(value Jsonable)()
    SetValueTypes(value Jsonable)()
}
