package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookApplication 
type WorkbookApplication struct {
    Entity
    // Returns the calculation mode used in the workbook. Possible values are: Automatic, AutomaticExceptTables, Manual.
    calculationMode *string
}
// NewWorkbookApplication instantiates a new workbookApplication and sets the default values.
func NewWorkbookApplication()(*WorkbookApplication) {
    m := &WorkbookApplication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookApplication(), nil
}
// GetCalculationMode gets the calculationMode property value. Returns the calculation mode used in the workbook. Possible values are: Automatic, AutomaticExceptTables, Manual.
func (m *WorkbookApplication) GetCalculationMode()(*string) {
    return m.calculationMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calculationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalculationMode(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WorkbookApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("calculationMode", m.GetCalculationMode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCalculationMode sets the calculationMode property value. Returns the calculation mode used in the workbook. Possible values are: Automatic, AutomaticExceptTables, Manual.
func (m *WorkbookApplication) SetCalculationMode(value *string)() {
    m.calculationMode = value
}
