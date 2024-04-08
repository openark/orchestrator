package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartGridlinesFormat 
type WorkbookChartGridlinesFormat struct {
    Entity
    // Represents chart line formatting. Read-only.
    line WorkbookChartLineFormatable
}
// NewWorkbookChartGridlinesFormat instantiates a new workbookChartGridlinesFormat and sets the default values.
func NewWorkbookChartGridlinesFormat()(*WorkbookChartGridlinesFormat) {
    m := &WorkbookChartGridlinesFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartGridlinesFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartGridlinesFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartGridlinesFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartGridlinesFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartLineFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val.(WorkbookChartLineFormatable))
        }
        return nil
    }
    return res
}
// GetLine gets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartGridlinesFormat) GetLine()(WorkbookChartLineFormatable) {
    return m.line
}
// Serialize serializes information the current object
func (m *WorkbookChartGridlinesFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("line", m.GetLine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLine sets the line property value. Represents chart line formatting. Read-only.
func (m *WorkbookChartGridlinesFormat) SetLine(value WorkbookChartLineFormatable)() {
    m.line = value
}
