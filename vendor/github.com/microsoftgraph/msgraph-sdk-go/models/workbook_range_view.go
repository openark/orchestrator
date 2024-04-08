package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRangeView 
type WorkbookRangeView struct {
    Entity
    // Represents the cell addresses
    cellAddresses Jsonable
    // Returns the number of visible columns. Read-only.
    columnCount *int32
    // Represents the formula in A1-style notation.
    formulas Jsonable
    // Represents the formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
    formulasLocal Jsonable
    // Represents the formula in R1C1-style notation.
    formulasR1C1 Jsonable
    // Index of the range.
    index *int32
    // Represents Excel's number format code for the given cell. Read-only.
    numberFormat Jsonable
    // Returns the number of visible rows. Read-only.
    rowCount *int32
    // Represents a collection of range views associated with the range. Read-only. Read-only.
    rows []WorkbookRangeViewable
    // Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
    text Jsonable
    // Represents the raw values of the specified range view. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
    values Jsonable
    // Represents the type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
    valueTypes Jsonable
}
// NewWorkbookRangeView instantiates a new WorkbookRangeView and sets the default values.
func NewWorkbookRangeView()(*WorkbookRangeView) {
    m := &WorkbookRangeView{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookRangeViewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookRangeViewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRangeView(), nil
}
// GetCellAddresses gets the cellAddresses property value. Represents the cell addresses
func (m *WorkbookRangeView) GetCellAddresses()(Jsonable) {
    return m.cellAddresses
}
// GetColumnCount gets the columnCount property value. Returns the number of visible columns. Read-only.
func (m *WorkbookRangeView) GetColumnCount()(*int32) {
    return m.columnCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRangeView) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cellAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellAddresses(val.(Jsonable))
        }
        return nil
    }
    res["columnCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnCount(val)
        }
        return nil
    }
    res["formulas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulas(val.(Jsonable))
        }
        return nil
    }
    res["formulasLocal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasLocal(val.(Jsonable))
        }
        return nil
    }
    res["formulasR1C1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasR1C1(val.(Jsonable))
        }
        return nil
    }
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["numberFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberFormat(val.(Jsonable))
        }
        return nil
    }
    res["rowCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowCount(val)
        }
        return nil
    }
    res["rows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookRangeViewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookRangeViewable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookRangeViewable)
            }
            m.SetRows(res)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(Jsonable))
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(Jsonable))
        }
        return nil
    }
    res["valueTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueTypes(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetFormulas gets the formulas property value. Represents the formula in A1-style notation.
func (m *WorkbookRangeView) GetFormulas()(Jsonable) {
    return m.formulas
}
// GetFormulasLocal gets the formulasLocal property value. Represents the formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
func (m *WorkbookRangeView) GetFormulasLocal()(Jsonable) {
    return m.formulasLocal
}
// GetFormulasR1C1 gets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
func (m *WorkbookRangeView) GetFormulasR1C1()(Jsonable) {
    return m.formulasR1C1
}
// GetIndex gets the index property value. Index of the range.
func (m *WorkbookRangeView) GetIndex()(*int32) {
    return m.index
}
// GetNumberFormat gets the numberFormat property value. Represents Excel's number format code for the given cell. Read-only.
func (m *WorkbookRangeView) GetNumberFormat()(Jsonable) {
    return m.numberFormat
}
// GetRowCount gets the rowCount property value. Returns the number of visible rows. Read-only.
func (m *WorkbookRangeView) GetRowCount()(*int32) {
    return m.rowCount
}
// GetRows gets the rows property value. Represents a collection of range views associated with the range. Read-only. Read-only.
func (m *WorkbookRangeView) GetRows()([]WorkbookRangeViewable) {
    return m.rows
}
// GetText gets the text property value. Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
func (m *WorkbookRangeView) GetText()(Jsonable) {
    return m.text
}
// GetValues gets the values property value. Represents the raw values of the specified range view. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookRangeView) GetValues()(Jsonable) {
    return m.values
}
// GetValueTypes gets the valueTypes property value. Represents the type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
func (m *WorkbookRangeView) GetValueTypes()(Jsonable) {
    return m.valueTypes
}
// Serialize serializes information the current object
func (m *WorkbookRangeView) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cellAddresses", m.GetCellAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("columnCount", m.GetColumnCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulas", m.GetFormulas())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulasLocal", m.GetFormulasLocal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulasR1C1", m.GetFormulasR1C1())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("numberFormat", m.GetNumberFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rowCount", m.GetRowCount())
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
        err = writer.WriteObjectValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("valueTypes", m.GetValueTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCellAddresses sets the cellAddresses property value. Represents the cell addresses
func (m *WorkbookRangeView) SetCellAddresses(value Jsonable)() {
    m.cellAddresses = value
}
// SetColumnCount sets the columnCount property value. Returns the number of visible columns. Read-only.
func (m *WorkbookRangeView) SetColumnCount(value *int32)() {
    m.columnCount = value
}
// SetFormulas sets the formulas property value. Represents the formula in A1-style notation.
func (m *WorkbookRangeView) SetFormulas(value Jsonable)() {
    m.formulas = value
}
// SetFormulasLocal sets the formulasLocal property value. Represents the formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
func (m *WorkbookRangeView) SetFormulasLocal(value Jsonable)() {
    m.formulasLocal = value
}
// SetFormulasR1C1 sets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
func (m *WorkbookRangeView) SetFormulasR1C1(value Jsonable)() {
    m.formulasR1C1 = value
}
// SetIndex sets the index property value. Index of the range.
func (m *WorkbookRangeView) SetIndex(value *int32)() {
    m.index = value
}
// SetNumberFormat sets the numberFormat property value. Represents Excel's number format code for the given cell. Read-only.
func (m *WorkbookRangeView) SetNumberFormat(value Jsonable)() {
    m.numberFormat = value
}
// SetRowCount sets the rowCount property value. Returns the number of visible rows. Read-only.
func (m *WorkbookRangeView) SetRowCount(value *int32)() {
    m.rowCount = value
}
// SetRows sets the rows property value. Represents a collection of range views associated with the range. Read-only. Read-only.
func (m *WorkbookRangeView) SetRows(value []WorkbookRangeViewable)() {
    m.rows = value
}
// SetText sets the text property value. Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
func (m *WorkbookRangeView) SetText(value Jsonable)() {
    m.text = value
}
// SetValues sets the values property value. Represents the raw values of the specified range view. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookRangeView) SetValues(value Jsonable)() {
    m.values = value
}
// SetValueTypes sets the valueTypes property value. Represents the type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
func (m *WorkbookRangeView) SetValueTypes(value Jsonable)() {
    m.valueTypes = value
}
