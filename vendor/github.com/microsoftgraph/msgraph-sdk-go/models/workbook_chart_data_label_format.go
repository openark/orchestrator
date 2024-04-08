package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartDataLabelFormat 
type WorkbookChartDataLabelFormat struct {
    Entity
    // Represents the fill format of the current chart data label. Read-only.
    fill WorkbookChartFillable
    // Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
    font WorkbookChartFontable
}
// NewWorkbookChartDataLabelFormat instantiates a new workbookChartDataLabelFormat and sets the default values.
func NewWorkbookChartDataLabelFormat()(*WorkbookChartDataLabelFormat) {
    m := &WorkbookChartDataLabelFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartDataLabelFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartDataLabelFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartDataLabelFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartDataLabelFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFillFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(WorkbookChartFillable))
        }
        return nil
    }
    res["font"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFontFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFont(val.(WorkbookChartFontable))
        }
        return nil
    }
    return res
}
// GetFill gets the fill property value. Represents the fill format of the current chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) GetFill()(WorkbookChartFillable) {
    return m.fill
}
// GetFont gets the font property value. Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) GetFont()(WorkbookChartFontable) {
    return m.font
}
// Serialize serializes information the current object
func (m *WorkbookChartDataLabelFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fill", m.GetFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFill sets the fill property value. Represents the fill format of the current chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) SetFill(value WorkbookChartFillable)() {
    m.fill = value
}
// SetFont sets the font property value. Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) SetFont(value WorkbookChartFontable)() {
    m.font = value
}
