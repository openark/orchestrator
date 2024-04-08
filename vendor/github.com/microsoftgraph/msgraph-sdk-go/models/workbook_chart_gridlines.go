package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartGridlines 
type WorkbookChartGridlines struct {
    Entity
    // Represents the formatting of chart gridlines. Read-only.
    format WorkbookChartGridlinesFormatable
    // Boolean value representing if the axis gridlines are visible or not.
    visible *bool
}
// NewWorkbookChartGridlines instantiates a new workbookChartGridlines and sets the default values.
func NewWorkbookChartGridlines()(*WorkbookChartGridlines) {
    m := &WorkbookChartGridlines{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartGridlinesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartGridlinesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartGridlines(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartGridlines) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartGridlinesFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartGridlinesFormatable))
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
// GetFormat gets the format property value. Represents the formatting of chart gridlines. Read-only.
func (m *WorkbookChartGridlines) GetFormat()(WorkbookChartGridlinesFormatable) {
    return m.format
}
// GetVisible gets the visible property value. Boolean value representing if the axis gridlines are visible or not.
func (m *WorkbookChartGridlines) GetVisible()(*bool) {
    return m.visible
}
// Serialize serializes information the current object
func (m *WorkbookChartGridlines) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("visible", m.GetVisible())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Represents the formatting of chart gridlines. Read-only.
func (m *WorkbookChartGridlines) SetFormat(value WorkbookChartGridlinesFormatable)() {
    m.format = value
}
// SetVisible sets the visible property value. Boolean value representing if the axis gridlines are visible or not.
func (m *WorkbookChartGridlines) SetVisible(value *bool)() {
    m.visible = value
}
