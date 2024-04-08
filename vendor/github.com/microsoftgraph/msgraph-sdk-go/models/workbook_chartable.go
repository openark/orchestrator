package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartable 
type WorkbookChartable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAxes()(WorkbookChartAxesable)
    GetDataLabels()(WorkbookChartDataLabelsable)
    GetFormat()(WorkbookChartAreaFormatable)
    GetHeight()(*float64)
    GetLeft()(*float64)
    GetLegend()(WorkbookChartLegendable)
    GetName()(*string)
    GetSeries()([]WorkbookChartSeriesable)
    GetTitle()(WorkbookChartTitleable)
    GetTop()(*float64)
    GetWidth()(*float64)
    GetWorksheet()(WorkbookWorksheetable)
    SetAxes(value WorkbookChartAxesable)()
    SetDataLabels(value WorkbookChartDataLabelsable)()
    SetFormat(value WorkbookChartAreaFormatable)()
    SetHeight(value *float64)()
    SetLeft(value *float64)()
    SetLegend(value WorkbookChartLegendable)()
    SetName(value *string)()
    SetSeries(value []WorkbookChartSeriesable)()
    SetTitle(value WorkbookChartTitleable)()
    SetTop(value *float64)()
    SetWidth(value *float64)()
    SetWorksheet(value WorkbookWorksheetable)()
}
