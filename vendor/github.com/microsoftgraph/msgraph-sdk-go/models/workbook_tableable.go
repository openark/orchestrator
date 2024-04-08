package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookTableable 
type WorkbookTableable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColumns()([]WorkbookTableColumnable)
    GetHighlightFirstColumn()(*bool)
    GetHighlightLastColumn()(*bool)
    GetLegacyId()(*string)
    GetName()(*string)
    GetRows()([]WorkbookTableRowable)
    GetShowBandedColumns()(*bool)
    GetShowBandedRows()(*bool)
    GetShowFilterButton()(*bool)
    GetShowHeaders()(*bool)
    GetShowTotals()(*bool)
    GetSort()(WorkbookTableSortable)
    GetStyle()(*string)
    GetWorksheet()(WorkbookWorksheetable)
    SetColumns(value []WorkbookTableColumnable)()
    SetHighlightFirstColumn(value *bool)()
    SetHighlightLastColumn(value *bool)()
    SetLegacyId(value *string)()
    SetName(value *string)()
    SetRows(value []WorkbookTableRowable)()
    SetShowBandedColumns(value *bool)()
    SetShowBandedRows(value *bool)()
    SetShowFilterButton(value *bool)()
    SetShowHeaders(value *bool)()
    SetShowTotals(value *bool)()
    SetSort(value WorkbookTableSortable)()
    SetStyle(value *string)()
    SetWorksheet(value WorkbookWorksheetable)()
}
