package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookWorksheet 
type WorkbookWorksheet struct {
    Entity
    // Returns collection of charts that are part of the worksheet. Read-only.
    charts []WorkbookChartable
    // The display name of the worksheet.
    name *string
    // Returns collection of names that are associated with the worksheet. Read-only.
    names []WorkbookNamedItemable
    // Collection of PivotTables that are part of the worksheet.
    pivotTables []WorkbookPivotTableable
    // The zero-based position of the worksheet within the workbook.
    position *int32
    // Returns sheet protection object for a worksheet. Read-only.
    protection WorkbookWorksheetProtectionable
    // Collection of tables that are part of the worksheet. Read-only.
    tables []WorkbookTableable
    // The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
    visibility *string
}
// NewWorkbookWorksheet instantiates a new workbookWorksheet and sets the default values.
func NewWorkbookWorksheet()(*WorkbookWorksheet) {
    m := &WorkbookWorksheet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookWorksheetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookWorksheetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookWorksheet(), nil
}
// GetCharts gets the charts property value. Returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) GetCharts()([]WorkbookChartable) {
    return m.charts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookWorksheet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["charts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookChartable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookChartable)
            }
            m.SetCharts(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookNamedItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookNamedItemable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookNamedItemable)
            }
            m.SetNames(res)
        }
        return nil
    }
    res["pivotTables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookPivotTableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookPivotTableable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookPivotTableable)
            }
            m.SetPivotTables(res)
        }
        return nil
    }
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    res["protection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookWorksheetProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtection(val.(WorkbookWorksheetProtectionable))
        }
        return nil
    }
    res["tables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookTableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookTableable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookTableable)
            }
            m.SetTables(res)
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The display name of the worksheet.
func (m *WorkbookWorksheet) GetName()(*string) {
    return m.name
}
// GetNames gets the names property value. Returns collection of names that are associated with the worksheet. Read-only.
func (m *WorkbookWorksheet) GetNames()([]WorkbookNamedItemable) {
    return m.names
}
// GetPivotTables gets the pivotTables property value. Collection of PivotTables that are part of the worksheet.
func (m *WorkbookWorksheet) GetPivotTables()([]WorkbookPivotTableable) {
    return m.pivotTables
}
// GetPosition gets the position property value. The zero-based position of the worksheet within the workbook.
func (m *WorkbookWorksheet) GetPosition()(*int32) {
    return m.position
}
// GetProtection gets the protection property value. Returns sheet protection object for a worksheet. Read-only.
func (m *WorkbookWorksheet) GetProtection()(WorkbookWorksheetProtectionable) {
    return m.protection
}
// GetTables gets the tables property value. Collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) GetTables()([]WorkbookTableable) {
    return m.tables
}
// GetVisibility gets the visibility property value. The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
func (m *WorkbookWorksheet) GetVisibility()(*string) {
    return m.visibility
}
// Serialize serializes information the current object
func (m *WorkbookWorksheet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCharts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCharts()))
        for i, v := range m.GetCharts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("charts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetNames() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNames()))
        for i, v := range m.GetNames() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("names", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPivotTables() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPivotTables()))
        for i, v := range m.GetPivotTables() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("pivotTables", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("protection", m.GetProtection())
        if err != nil {
            return err
        }
    }
    if m.GetTables() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTables()))
        for i, v := range m.GetTables() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tables", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCharts sets the charts property value. Returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) SetCharts(value []WorkbookChartable)() {
    m.charts = value
}
// SetName sets the name property value. The display name of the worksheet.
func (m *WorkbookWorksheet) SetName(value *string)() {
    m.name = value
}
// SetNames sets the names property value. Returns collection of names that are associated with the worksheet. Read-only.
func (m *WorkbookWorksheet) SetNames(value []WorkbookNamedItemable)() {
    m.names = value
}
// SetPivotTables sets the pivotTables property value. Collection of PivotTables that are part of the worksheet.
func (m *WorkbookWorksheet) SetPivotTables(value []WorkbookPivotTableable)() {
    m.pivotTables = value
}
// SetPosition sets the position property value. The zero-based position of the worksheet within the workbook.
func (m *WorkbookWorksheet) SetPosition(value *int32)() {
    m.position = value
}
// SetProtection sets the protection property value. Returns sheet protection object for a worksheet. Read-only.
func (m *WorkbookWorksheet) SetProtection(value WorkbookWorksheetProtectionable)() {
    m.protection = value
}
// SetTables sets the tables property value. Collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) SetTables(value []WorkbookTableable)() {
    m.tables = value
}
// SetVisibility sets the visibility property value. The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
func (m *WorkbookWorksheet) SetVisibility(value *string)() {
    m.visibility = value
}
