package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartFontable 
type WorkbookChartFontable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBold()(*bool)
    GetColor()(*string)
    GetItalic()(*bool)
    GetName()(*string)
    GetSize()(*float64)
    GetUnderline()(*string)
    SetBold(value *bool)()
    SetColor(value *string)()
    SetItalic(value *bool)()
    SetName(value *string)()
    SetSize(value *float64)()
    SetUnderline(value *string)()
}
