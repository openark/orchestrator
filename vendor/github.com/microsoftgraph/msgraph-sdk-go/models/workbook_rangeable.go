package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRangeable 
type WorkbookRangeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetAddressLocal()(*string)
    GetCellCount()(*int32)
    GetColumnCount()(*int32)
    GetColumnHidden()(*bool)
    GetColumnIndex()(*int32)
    GetFormat()(WorkbookRangeFormatable)
    GetFormulas()(Jsonable)
    GetFormulasLocal()(Jsonable)
    GetFormulasR1C1()(Jsonable)
    GetHidden()(*bool)
    GetNumberFormat()(Jsonable)
    GetRowCount()(*int32)
    GetRowHidden()(*bool)
    GetRowIndex()(*int32)
    GetSort()(WorkbookRangeSortable)
    GetText()(Jsonable)
    GetValues()(Jsonable)
    GetValueTypes()(Jsonable)
    GetWorksheet()(WorkbookWorksheetable)
    SetAddress(value *string)()
    SetAddressLocal(value *string)()
    SetCellCount(value *int32)()
    SetColumnCount(value *int32)()
    SetColumnHidden(value *bool)()
    SetColumnIndex(value *int32)()
    SetFormat(value WorkbookRangeFormatable)()
    SetFormulas(value Jsonable)()
    SetFormulasLocal(value Jsonable)()
    SetFormulasR1C1(value Jsonable)()
    SetHidden(value *bool)()
    SetNumberFormat(value Jsonable)()
    SetRowCount(value *int32)()
    SetRowHidden(value *bool)()
    SetRowIndex(value *int32)()
    SetSort(value WorkbookRangeSortable)()
    SetText(value Jsonable)()
    SetValues(value Jsonable)()
    SetValueTypes(value Jsonable)()
    SetWorksheet(value WorkbookWorksheetable)()
}
