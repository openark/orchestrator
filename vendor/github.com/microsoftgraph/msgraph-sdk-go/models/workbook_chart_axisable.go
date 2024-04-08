package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAxisable 
type WorkbookChartAxisable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(WorkbookChartAxisFormatable)
    GetMajorGridlines()(WorkbookChartGridlinesable)
    GetMajorUnit()(Jsonable)
    GetMaximum()(Jsonable)
    GetMinimum()(Jsonable)
    GetMinorGridlines()(WorkbookChartGridlinesable)
    GetMinorUnit()(Jsonable)
    GetTitle()(WorkbookChartAxisTitleable)
    SetFormat(value WorkbookChartAxisFormatable)()
    SetMajorGridlines(value WorkbookChartGridlinesable)()
    SetMajorUnit(value Jsonable)()
    SetMaximum(value Jsonable)()
    SetMinimum(value Jsonable)()
    SetMinorGridlines(value WorkbookChartGridlinesable)()
    SetMinorUnit(value Jsonable)()
    SetTitle(value WorkbookChartAxisTitleable)()
}
