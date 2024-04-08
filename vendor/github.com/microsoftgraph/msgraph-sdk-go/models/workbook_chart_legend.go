package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartLegend 
type WorkbookChartLegend struct {
    Entity
    // Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
    format WorkbookChartLegendFormatable
    // Boolean value for whether the chart legend should overlap with the main body of the chart.
    overlay *bool
    // Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
    position *string
    // A boolean value the represents the visibility of a ChartLegend object.
    visible *bool
}
// NewWorkbookChartLegend instantiates a new workbookChartLegend and sets the default values.
func NewWorkbookChartLegend()(*WorkbookChartLegend) {
    m := &WorkbookChartLegend{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartLegendFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartLegendFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartLegend(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartLegend) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartLegendFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartLegendFormatable))
        }
        return nil
    }
    res["overlay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverlay(val)
        }
        return nil
    }
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    res["visible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisible(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
func (m *WorkbookChartLegend) GetFormat()(WorkbookChartLegendFormatable) {
    return m.format
}
// GetOverlay gets the overlay property value. Boolean value for whether the chart legend should overlap with the main body of the chart.
func (m *WorkbookChartLegend) GetOverlay()(*bool) {
    return m.overlay
}
// GetPosition gets the position property value. Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
func (m *WorkbookChartLegend) GetPosition()(*string) {
    return m.position
}
// GetVisible gets the visible property value. A boolean value the represents the visibility of a ChartLegend object.
func (m *WorkbookChartLegend) GetVisible()(*bool) {
    return m.visible
}
// Serialize serializes information the current object
func (m *WorkbookChartLegend) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("overlay", m.GetOverlay())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("visible", m.GetVisible())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Represents the formatting of a chart legend, which includes fill and font formatting. Read-only.
func (m *WorkbookChartLegend) SetFormat(value WorkbookChartLegendFormatable)() {
    m.format = value
}
// SetOverlay sets the overlay property value. Boolean value for whether the chart legend should overlap with the main body of the chart.
func (m *WorkbookChartLegend) SetOverlay(value *bool)() {
    m.overlay = value
}
// SetPosition sets the position property value. Represents the position of the legend on the chart. The possible values are: Top, Bottom, Left, Right, Corner, Custom.
func (m *WorkbookChartLegend) SetPosition(value *string)() {
    m.position = value
}
// SetVisible sets the visible property value. A boolean value the represents the visibility of a ChartLegend object.
func (m *WorkbookChartLegend) SetVisible(value *bool)() {
    m.visible = value
}
