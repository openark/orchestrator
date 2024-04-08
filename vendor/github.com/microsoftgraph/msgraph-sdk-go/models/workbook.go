package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Workbook 
type Workbook struct {
    Entity
    // The application property
    application WorkbookApplicationable
    // The comments property
    comments []WorkbookCommentable
    // The functions property
    functions WorkbookFunctionsable
    // Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
    names []WorkbookNamedItemable
    // The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
    operations []WorkbookOperationable
    // Represents a collection of tables associated with the workbook. Read-only.
    tables []WorkbookTableable
    // Represents a collection of worksheets associated with the workbook. Read-only.
    worksheets []WorkbookWorksheetable
}
// NewWorkbook instantiates a new workbook and sets the default values.
func NewWorkbook()(*Workbook) {
    m := &Workbook{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbook(), nil
}
// GetApplication gets the application property value. The application property
func (m *Workbook) GetApplication()(WorkbookApplicationable) {
    return m.application
}
// GetComments gets the comments property value. The comments property
func (m *Workbook) GetComments()([]WorkbookCommentable) {
    return m.comments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Workbook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(WorkbookApplicationable))
        }
        return nil
    }
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookCommentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookCommentable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookCommentable)
            }
            m.SetComments(res)
        }
        return nil
    }
    res["functions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookFunctionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunctions(val.(WorkbookFunctionsable))
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
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookOperationable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookOperationable)
            }
            m.SetOperations(res)
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
    res["worksheets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkbookWorksheetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookWorksheetable, len(val))
            for i, v := range val {
                res[i] = v.(WorkbookWorksheetable)
            }
            m.SetWorksheets(res)
        }
        return nil
    }
    return res
}
// GetFunctions gets the functions property value. The functions property
func (m *Workbook) GetFunctions()(WorkbookFunctionsable) {
    return m.functions
}
// GetNames gets the names property value. Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
func (m *Workbook) GetNames()([]WorkbookNamedItemable) {
    return m.names
}
// GetOperations gets the operations property value. The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
func (m *Workbook) GetOperations()([]WorkbookOperationable) {
    return m.operations
}
// GetTables gets the tables property value. Represents a collection of tables associated with the workbook. Read-only.
func (m *Workbook) GetTables()([]WorkbookTableable) {
    return m.tables
}
// GetWorksheets gets the worksheets property value. Represents a collection of worksheets associated with the workbook. Read-only.
func (m *Workbook) GetWorksheets()([]WorkbookWorksheetable) {
    return m.worksheets
}
// Serialize serializes information the current object
func (m *Workbook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    if m.GetComments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComments()))
        for i, v := range m.GetComments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("comments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("functions", m.GetFunctions())
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
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
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
    if m.GetWorksheets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorksheets()))
        for i, v := range m.GetWorksheets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("worksheets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The application property
func (m *Workbook) SetApplication(value WorkbookApplicationable)() {
    m.application = value
}
// SetComments sets the comments property value. The comments property
func (m *Workbook) SetComments(value []WorkbookCommentable)() {
    m.comments = value
}
// SetFunctions sets the functions property value. The functions property
func (m *Workbook) SetFunctions(value WorkbookFunctionsable)() {
    m.functions = value
}
// SetNames sets the names property value. Represents a collection of workbooks scoped named items (named ranges and constants). Read-only.
func (m *Workbook) SetNames(value []WorkbookNamedItemable)() {
    m.names = value
}
// SetOperations sets the operations property value. The status of workbook operations. Getting an operation collection is not supported, but you can get the status of a long-running operation if the Location header is returned in the response. Read-only.
func (m *Workbook) SetOperations(value []WorkbookOperationable)() {
    m.operations = value
}
// SetTables sets the tables property value. Represents a collection of tables associated with the workbook. Read-only.
func (m *Workbook) SetTables(value []WorkbookTableable)() {
    m.tables = value
}
// SetWorksheets sets the worksheets property value. Represents a collection of worksheets associated with the workbook. Read-only.
func (m *Workbook) SetWorksheets(value []WorkbookWorksheetable)() {
    m.worksheets = value
}
