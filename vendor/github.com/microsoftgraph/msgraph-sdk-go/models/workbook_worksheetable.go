package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookWorksheetable 
type WorkbookWorksheetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCharts()([]WorkbookChartable)
    GetName()(*string)
    GetNames()([]WorkbookNamedItemable)
    GetPivotTables()([]WorkbookPivotTableable)
    GetPosition()(*int32)
    GetProtection()(WorkbookWorksheetProtectionable)
    GetTables()([]WorkbookTableable)
    GetVisibility()(*string)
    SetCharts(value []WorkbookChartable)()
    SetName(value *string)()
    SetNames(value []WorkbookNamedItemable)()
    SetPivotTables(value []WorkbookPivotTableable)()
    SetPosition(value *int32)()
    SetProtection(value WorkbookWorksheetProtectionable)()
    SetTables(value []WorkbookTableable)()
    SetVisibility(value *string)()
}
