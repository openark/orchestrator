package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookTable 
type WorkbookTable struct {
    Entity
    // Represents a collection of all the columns in the table. Read-only.
    columns []WorkbookTableColumnable
    // Indicates whether the first column contains special formatting.
    highlightFirstColumn *bool
    // Indicates whether the last column contains special formatting.
    highlightLastColumn *bool
    // Legacy Id used in older Excle clients. The value of the identifier remains the same even when the table is renamed. This property should be interpreted as an opaque string value and should not be parsed to any other type. Read-only.
    legacyId *string
    // Name of the table.
    name *string
    // Represents a collection of all the rows in the table. Read-only.
    rows []WorkbookTableRowable
    // Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones to make reading the table easier.
    showBandedColumns *bool
    // Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to make reading the table easier.
    showBandedRows *bool
    // Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if the table contains a header row.
    showFilterButton *bool
    // Indicates whether the header row is visible or not. This value can be set to show or remove the header row.
    showHeaders *bool
    // Indicates whether the total row is visible or not. This value can be set to show or remove the total row.
    showTotals *bool
    // Represents the sorting for the table. Read-only.
    sort WorkbookTableSortable
    // Constant value that represents the Table style. The possible values are: TableStyleLight1 thru TableStyleLight21, TableStyleMedium1 thru TableStyleMedium28, TableStyleStyleDark1 thru TableStyleStyleDark11. A custom user-defined style present in the workbook can also be specified.
    style *string
    // The worksheet containing the current table. Read-only.
    worksheet WorkbookWorksheetable
}
// NewWorkbookTable instantiates a new workbookTable and sets the default values.
func NewWorkbookTable()(*WorkbookTable) {
    m := &WorkbookTable{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookTableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookTableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookTable(), nil
}
// GetColumns gets the columns property value. Represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTable) GetColumns()([]WorkbookTableColumnable) {
    return m.columns
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookTable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["columns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookTableColumnFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookTableColumnable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookTableColumnable)
            }
            m.SetColumns(res)
        }
        return nil
    }
    res["highlightFirstColumn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighlightFirstColumn(val)
        }
        return nil
    }
    res["highlightLastColumn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighlightLastColumn(val)
        }
        return nil
    }
    res["legacyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegacyId(val)
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
    res["rows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookTableRowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookTableRowable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookTableRowable)
            }
            m.SetRows(res)
        }
        return nil
    }
    res["showBandedColumns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowBandedColumns(val)
        }
        return nil
    }
    res["showBandedRows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowBandedRows(val)
        }
        return nil
    }
    res["showFilterButton"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowFilterButton(val)
        }
        return nil
    }
    res["showHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowHeaders(val)
        }
        return nil
    }
    res["showTotals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowTotals(val)
        }
        return nil
    }
    res["sort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookTableSortFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSort(val.(WorkbookTableSortable))
        }
        return nil
    }
    res["style"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStyle(val)
        }
        return nil
    }
    res["worksheet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookWorksheetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorksheet(val.(WorkbookWorksheetable))
        }
        return nil
    }
    return res
}
// GetHighlightFirstColumn gets the highlightFirstColumn property value. Indicates whether the first column contains special formatting.
func (m *WorkbookTable) GetHighlightFirstColumn()(*bool) {
    return m.highlightFirstColumn
}
// GetHighlightLastColumn gets the highlightLastColumn property value. Indicates whether the last column contains special formatting.
func (m *WorkbookTable) GetHighlightLastColumn()(*bool) {
    return m.highlightLastColumn
}
// GetLegacyId gets the legacyId property value. Legacy Id used in older Excle clients. The value of the identifier remains the same even when the table is renamed. This property should be interpreted as an opaque string value and should not be parsed to any other type. Read-only.
func (m *WorkbookTable) GetLegacyId()(*string) {
    return m.legacyId
}
// GetName gets the name property value. Name of the table.
func (m *WorkbookTable) GetName()(*string) {
    return m.name
}
// GetRows gets the rows property value. Represents a collection of all the rows in the table. Read-only.
func (m *WorkbookTable) GetRows()([]WorkbookTableRowable) {
    return m.rows
}
// GetShowBandedColumns gets the showBandedColumns property value. Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones to make reading the table easier.
func (m *WorkbookTable) GetShowBandedColumns()(*bool) {
    return m.showBandedColumns
}
// GetShowBandedRows gets the showBandedRows property value. Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to make reading the table easier.
func (m *WorkbookTable) GetShowBandedRows()(*bool) {
    return m.showBandedRows
}
// GetShowFilterButton gets the showFilterButton property value. Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if the table contains a header row.
func (m *WorkbookTable) GetShowFilterButton()(*bool) {
    return m.showFilterButton
}
// GetShowHeaders gets the showHeaders property value. Indicates whether the header row is visible or not. This value can be set to show or remove the header row.
func (m *WorkbookTable) GetShowHeaders()(*bool) {
    return m.showHeaders
}
// GetShowTotals gets the showTotals property value. Indicates whether the total row is visible or not. This value can be set to show or remove the total row.
func (m *WorkbookTable) GetShowTotals()(*bool) {
    return m.showTotals
}
// GetSort gets the sort property value. Represents the sorting for the table. Read-only.
func (m *WorkbookTable) GetSort()(WorkbookTableSortable) {
    return m.sort
}
// GetStyle gets the style property value. Constant value that represents the Table style. The possible values are: TableStyleLight1 thru TableStyleLight21, TableStyleMedium1 thru TableStyleMedium28, TableStyleStyleDark1 thru TableStyleStyleDark11. A custom user-defined style present in the workbook can also be specified.
func (m *WorkbookTable) GetStyle()(*string) {
    return m.style
}
// GetWorksheet gets the worksheet property value. The worksheet containing the current table. Read-only.
func (m *WorkbookTable) GetWorksheet()(WorkbookWorksheetable) {
    return m.worksheet
}
// Serialize serializes information the current object
func (m *WorkbookTable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColumns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("highlightFirstColumn", m.GetHighlightFirstColumn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("highlightLastColumn", m.GetHighlightLastColumn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("legacyId", m.GetLegacyId())
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
    if m.GetRows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRows()))
        for i, v := range m.GetRows() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rows", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showBandedColumns", m.GetShowBandedColumns())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showBandedRows", m.GetShowBandedRows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showFilterButton", m.GetShowFilterButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showHeaders", m.GetShowHeaders())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showTotals", m.GetShowTotals())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sort", m.GetSort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("style", m.GetStyle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("worksheet", m.GetWorksheet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColumns sets the columns property value. Represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTable) SetColumns(value []WorkbookTableColumnable)() {
    m.columns = value
}
// SetHighlightFirstColumn sets the highlightFirstColumn property value. Indicates whether the first column contains special formatting.
func (m *WorkbookTable) SetHighlightFirstColumn(value *bool)() {
    m.highlightFirstColumn = value
}
// SetHighlightLastColumn sets the highlightLastColumn property value. Indicates whether the last column contains special formatting.
func (m *WorkbookTable) SetHighlightLastColumn(value *bool)() {
    m.highlightLastColumn = value
}
// SetLegacyId sets the legacyId property value. Legacy Id used in older Excle clients. The value of the identifier remains the same even when the table is renamed. This property should be interpreted as an opaque string value and should not be parsed to any other type. Read-only.
func (m *WorkbookTable) SetLegacyId(value *string)() {
    m.legacyId = value
}
// SetName sets the name property value. Name of the table.
func (m *WorkbookTable) SetName(value *string)() {
    m.name = value
}
// SetRows sets the rows property value. Represents a collection of all the rows in the table. Read-only.
func (m *WorkbookTable) SetRows(value []WorkbookTableRowable)() {
    m.rows = value
}
// SetShowBandedColumns sets the showBandedColumns property value. Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones to make reading the table easier.
func (m *WorkbookTable) SetShowBandedColumns(value *bool)() {
    m.showBandedColumns = value
}
// SetShowBandedRows sets the showBandedRows property value. Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to make reading the table easier.
func (m *WorkbookTable) SetShowBandedRows(value *bool)() {
    m.showBandedRows = value
}
// SetShowFilterButton sets the showFilterButton property value. Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if the table contains a header row.
func (m *WorkbookTable) SetShowFilterButton(value *bool)() {
    m.showFilterButton = value
}
// SetShowHeaders sets the showHeaders property value. Indicates whether the header row is visible or not. This value can be set to show or remove the header row.
func (m *WorkbookTable) SetShowHeaders(value *bool)() {
    m.showHeaders = value
}
// SetShowTotals sets the showTotals property value. Indicates whether the total row is visible or not. This value can be set to show or remove the total row.
func (m *WorkbookTable) SetShowTotals(value *bool)() {
    m.showTotals = value
}
// SetSort sets the sort property value. Represents the sorting for the table. Read-only.
func (m *WorkbookTable) SetSort(value WorkbookTableSortable)() {
    m.sort = value
}
// SetStyle sets the style property value. Constant value that represents the Table style. The possible values are: TableStyleLight1 thru TableStyleLight21, TableStyleMedium1 thru TableStyleMedium28, TableStyleStyleDark1 thru TableStyleStyleDark11. A custom user-defined style present in the workbook can also be specified.
func (m *WorkbookTable) SetStyle(value *string)() {
    m.style = value
}
// SetWorksheet sets the worksheet property value. The worksheet containing the current table. Read-only.
func (m *WorkbookTable) SetWorksheet(value WorkbookWorksheetable)() {
    m.worksheet = value
}
