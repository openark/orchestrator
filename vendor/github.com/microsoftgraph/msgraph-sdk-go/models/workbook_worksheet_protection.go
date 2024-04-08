package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookWorksheetProtection 
type WorkbookWorksheetProtection struct {
    Entity
    // Sheet protection options. Read-only.
    options WorkbookWorksheetProtectionOptionsable
    // Indicates if the worksheet is protected.  Read-only.
    protected *bool
}
// NewWorkbookWorksheetProtection instantiates a new workbookWorksheetProtection and sets the default values.
func NewWorkbookWorksheetProtection()(*WorkbookWorksheetProtection) {
    m := &WorkbookWorksheetProtection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookWorksheetProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookWorksheetProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookWorksheetProtection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookWorksheetProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookWorksheetProtectionOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptions(val.(WorkbookWorksheetProtectionOptionsable))
        }
        return nil
    }
    res["protected"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtected(val)
        }
        return nil
    }
    return res
}
// GetOptions gets the options property value. Sheet protection options. Read-only.
func (m *WorkbookWorksheetProtection) GetOptions()(WorkbookWorksheetProtectionOptionsable) {
    return m.options
}
// GetProtected gets the protected property value. Indicates if the worksheet is protected.  Read-only.
func (m *WorkbookWorksheetProtection) GetProtected()(*bool) {
    return m.protected
}
// Serialize serializes information the current object
func (m *WorkbookWorksheetProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("options", m.GetOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("protected", m.GetProtected())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOptions sets the options property value. Sheet protection options. Read-only.
func (m *WorkbookWorksheetProtection) SetOptions(value WorkbookWorksheetProtectionOptionsable)() {
    m.options = value
}
// SetProtected sets the protected property value. Indicates if the worksheet is protected.  Read-only.
func (m *WorkbookWorksheetProtection) SetProtected(value *bool)() {
    m.protected = value
}
