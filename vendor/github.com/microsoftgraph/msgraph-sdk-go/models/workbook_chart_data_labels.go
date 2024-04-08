package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartDataLabels 
type WorkbookChartDataLabels struct {
    Entity
    // Represents the format of chart data labels, which includes fill and font formatting. Read-only.
    format WorkbookChartDataLabelFormatable
    // DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
    position *string
    // String representing the separator used for the data labels on a chart.
    separator *string
    // Boolean value representing if the data label bubble size is visible or not.
    showBubbleSize *bool
    // Boolean value representing if the data label category name is visible or not.
    showCategoryName *bool
    // Boolean value representing if the data label legend key is visible or not.
    showLegendKey *bool
    // Boolean value representing if the data label percentage is visible or not.
    showPercentage *bool
    // Boolean value representing if the data label series name is visible or not.
    showSeriesName *bool
    // Boolean value representing if the data label value is visible or not.
    showValue *bool
}
// NewWorkbookChartDataLabels instantiates a new workbookChartDataLabels and sets the default values.
func NewWorkbookChartDataLabels()(*WorkbookChartDataLabels) {
    m := &WorkbookChartDataLabels{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartDataLabelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartDataLabelsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartDataLabels(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartDataLabels) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartDataLabelFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartDataLabelFormatable))
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
    res["separator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeparator(val)
        }
        return nil
    }
    res["showBubbleSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowBubbleSize(val)
        }
        return nil
    }
    res["showCategoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowCategoryName(val)
        }
        return nil
    }
    res["showLegendKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowLegendKey(val)
        }
        return nil
    }
    res["showPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowPercentage(val)
        }
        return nil
    }
    res["showSeriesName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowSeriesName(val)
        }
        return nil
    }
    res["showValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowValue(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Represents the format of chart data labels, which includes fill and font formatting. Read-only.
func (m *WorkbookChartDataLabels) GetFormat()(WorkbookChartDataLabelFormatable) {
    return m.format
}
// GetPosition gets the position property value. DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
func (m *WorkbookChartDataLabels) GetPosition()(*string) {
    return m.position
}
// GetSeparator gets the separator property value. String representing the separator used for the data labels on a chart.
func (m *WorkbookChartDataLabels) GetSeparator()(*string) {
    return m.separator
}
// GetShowBubbleSize gets the showBubbleSize property value. Boolean value representing if the data label bubble size is visible or not.
func (m *WorkbookChartDataLabels) GetShowBubbleSize()(*bool) {
    return m.showBubbleSize
}
// GetShowCategoryName gets the showCategoryName property value. Boolean value representing if the data label category name is visible or not.
func (m *WorkbookChartDataLabels) GetShowCategoryName()(*bool) {
    return m.showCategoryName
}
// GetShowLegendKey gets the showLegendKey property value. Boolean value representing if the data label legend key is visible or not.
func (m *WorkbookChartDataLabels) GetShowLegendKey()(*bool) {
    return m.showLegendKey
}
// GetShowPercentage gets the showPercentage property value. Boolean value representing if the data label percentage is visible or not.
func (m *WorkbookChartDataLabels) GetShowPercentage()(*bool) {
    return m.showPercentage
}
// GetShowSeriesName gets the showSeriesName property value. Boolean value representing if the data label series name is visible or not.
func (m *WorkbookChartDataLabels) GetShowSeriesName()(*bool) {
    return m.showSeriesName
}
// GetShowValue gets the showValue property value. Boolean value representing if the data label value is visible or not.
func (m *WorkbookChartDataLabels) GetShowValue()(*bool) {
    return m.showValue
}
// Serialize serializes information the current object
func (m *WorkbookChartDataLabels) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("separator", m.GetSeparator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showBubbleSize", m.GetShowBubbleSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showCategoryName", m.GetShowCategoryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showLegendKey", m.GetShowLegendKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showPercentage", m.GetShowPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showSeriesName", m.GetShowSeriesName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showValue", m.GetShowValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Represents the format of chart data labels, which includes fill and font formatting. Read-only.
func (m *WorkbookChartDataLabels) SetFormat(value WorkbookChartDataLabelFormatable)() {
    m.format = value
}
// SetPosition sets the position property value. DataLabelPosition value that represents the position of the data label. The possible values are: None, Center, InsideEnd, InsideBase, OutsideEnd, Left, Right, Top, Bottom, BestFit, Callout.
func (m *WorkbookChartDataLabels) SetPosition(value *string)() {
    m.position = value
}
// SetSeparator sets the separator property value. String representing the separator used for the data labels on a chart.
func (m *WorkbookChartDataLabels) SetSeparator(value *string)() {
    m.separator = value
}
// SetShowBubbleSize sets the showBubbleSize property value. Boolean value representing if the data label bubble size is visible or not.
func (m *WorkbookChartDataLabels) SetShowBubbleSize(value *bool)() {
    m.showBubbleSize = value
}
// SetShowCategoryName sets the showCategoryName property value. Boolean value representing if the data label category name is visible or not.
func (m *WorkbookChartDataLabels) SetShowCategoryName(value *bool)() {
    m.showCategoryName = value
}
// SetShowLegendKey sets the showLegendKey property value. Boolean value representing if the data label legend key is visible or not.
func (m *WorkbookChartDataLabels) SetShowLegendKey(value *bool)() {
    m.showLegendKey = value
}
// SetShowPercentage sets the showPercentage property value. Boolean value representing if the data label percentage is visible or not.
func (m *WorkbookChartDataLabels) SetShowPercentage(value *bool)() {
    m.showPercentage = value
}
// SetShowSeriesName sets the showSeriesName property value. Boolean value representing if the data label series name is visible or not.
func (m *WorkbookChartDataLabels) SetShowSeriesName(value *bool)() {
    m.showSeriesName = value
}
// SetShowValue sets the showValue property value. Boolean value representing if the data label value is visible or not.
func (m *WorkbookChartDataLabels) SetShowValue(value *bool)() {
    m.showValue = value
}
