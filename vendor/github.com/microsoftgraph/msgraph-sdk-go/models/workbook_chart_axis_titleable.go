package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAxisTitleable 
type WorkbookChartAxisTitleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(WorkbookChartAxisTitleFormatable)
    GetText()(*string)
    GetVisible()(*bool)
    SetFormat(value WorkbookChartAxisTitleFormatable)()
    SetText(value *string)()
    SetVisible(value *bool)()
}
