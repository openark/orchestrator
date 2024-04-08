package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAxisTitle 
type WorkbookChartAxisTitle struct {
    Entity
    // Represents the formatting of chart axis title. Read-only.
    format WorkbookChartAxisTitleFormatable
    // Represents the axis title.
    text *string
    // A boolean that specifies the visibility of an axis title.
    visible *bool
}
// NewWorkbookChartAxisTitle instantiates a new workbookChartAxisTitle and sets the default values.
func NewWorkbookChartAxisTitle()(*WorkbookChartAxisTitle) {
    m := &WorkbookChartAxisTitle{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxisTitleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxisTitleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartAxisTitle(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxisTitle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisTitleFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartAxisTitleFormatable))
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
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
// GetFormat gets the format property value. Represents the formatting of chart axis title. Read-only.
func (m *WorkbookChartAxisTitle) GetFormat()(WorkbookChartAxisTitleFormatable) {
    return m.format
}
// GetText gets the text property value. Represents the axis title.
func (m *WorkbookChartAxisTitle) GetText()(*string) {
    return m.text
}
// GetVisible gets the visible property value. A boolean that specifies the visibility of an axis title.
func (m *WorkbookChartAxisTitle) GetVisible()(*bool) {
    return m.visible
}
// Serialize serializes information the current object
func (m *WorkbookChartAxisTitle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("text", m.GetText())
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
// SetFormat sets the format property value. Represents the formatting of chart axis title. Read-only.
func (m *WorkbookChartAxisTitle) SetFormat(value WorkbookChartAxisTitleFormatable)() {
    m.format = value
}
// SetText sets the text property value. Represents the axis title.
func (m *WorkbookChartAxisTitle) SetText(value *string)() {
    m.text = value
}
// SetVisible sets the visible property value. A boolean that specifies the visibility of an axis title.
func (m *WorkbookChartAxisTitle) SetVisible(value *bool)() {
    m.visible = value
}
