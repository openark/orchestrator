package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookWorksheetProtectionOptionsable 
type WorkbookWorksheetProtectionOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAutoFilter()(*bool)
    GetAllowDeleteColumns()(*bool)
    GetAllowDeleteRows()(*bool)
    GetAllowFormatCells()(*bool)
    GetAllowFormatColumns()(*bool)
    GetAllowFormatRows()(*bool)
    GetAllowInsertColumns()(*bool)
    GetAllowInsertHyperlinks()(*bool)
    GetAllowInsertRows()(*bool)
    GetAllowPivotTables()(*bool)
    GetAllowSort()(*bool)
    GetOdataType()(*string)
    SetAllowAutoFilter(value *bool)()
    SetAllowDeleteColumns(value *bool)()
    SetAllowDeleteRows(value *bool)()
    SetAllowFormatCells(value *bool)()
    SetAllowFormatColumns(value *bool)()
    SetAllowFormatRows(value *bool)()
    SetAllowInsertColumns(value *bool)()
    SetAllowInsertHyperlinks(value *bool)()
    SetAllowInsertRows(value *bool)()
    SetAllowPivotTables(value *bool)()
    SetAllowSort(value *bool)()
    SetOdataType(value *string)()
}
