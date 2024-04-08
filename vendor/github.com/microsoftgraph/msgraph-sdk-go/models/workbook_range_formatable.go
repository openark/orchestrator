package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRangeFormatable 
type WorkbookRangeFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBorders()([]WorkbookRangeBorderable)
    GetColumnWidth()(*float64)
    GetFill()(WorkbookRangeFillable)
    GetFont()(WorkbookRangeFontable)
    GetHorizontalAlignment()(*string)
    GetProtection()(WorkbookFormatProtectionable)
    GetRowHeight()(*float64)
    GetVerticalAlignment()(*string)
    GetWrapText()(*bool)
    SetBorders(value []WorkbookRangeBorderable)()
    SetColumnWidth(value *float64)()
    SetFill(value WorkbookRangeFillable)()
    SetFont(value WorkbookRangeFontable)()
    SetHorizontalAlignment(value *string)()
    SetProtection(value WorkbookFormatProtectionable)()
    SetRowHeight(value *float64)()
    SetVerticalAlignment(value *string)()
    SetWrapText(value *bool)()
}
