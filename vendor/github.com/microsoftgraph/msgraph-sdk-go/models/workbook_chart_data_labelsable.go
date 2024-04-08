package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartDataLabelsable 
type WorkbookChartDataLabelsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(WorkbookChartDataLabelFormatable)
    GetPosition()(*string)
    GetSeparator()(*string)
    GetShowBubbleSize()(*bool)
    GetShowCategoryName()(*bool)
    GetShowLegendKey()(*bool)
    GetShowPercentage()(*bool)
    GetShowSeriesName()(*bool)
    GetShowValue()(*bool)
    SetFormat(value WorkbookChartDataLabelFormatable)()
    SetPosition(value *string)()
    SetSeparator(value *string)()
    SetShowBubbleSize(value *bool)()
    SetShowCategoryName(value *bool)()
    SetShowLegendKey(value *bool)()
    SetShowPercentage(value *bool)()
    SetShowSeriesName(value *bool)()
    SetShowValue(value *bool)()
}
