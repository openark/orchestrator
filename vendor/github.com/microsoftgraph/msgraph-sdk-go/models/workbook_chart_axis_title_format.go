package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAxisTitleFormat 
type WorkbookChartAxisTitleFormat struct {
    Entity
    // Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
    font WorkbookChartFontable
}
// NewWorkbookChartAxisTitleFormat instantiates a new workbookChartAxisTitleFormat and sets the default values.
func NewWorkbookChartAxisTitleFormat()(*WorkbookChartAxisTitleFormat) {
    m := &WorkbookChartAxisTitleFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxisTitleFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxisTitleFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartAxisTitleFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxisTitleFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
// GetFont gets the font property value. Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
func (m *WorkbookChartAxisTitleFormat) GetFont()(WorkbookChartFontable) {
    return m.font
}
// Serialize serializes information the current object
func (m *WorkbookChartAxisTitleFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFont sets the font property value. Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
func (m *WorkbookChartAxisTitleFormat) SetFont(value WorkbookChartFontable)() {
    m.font = value
}
