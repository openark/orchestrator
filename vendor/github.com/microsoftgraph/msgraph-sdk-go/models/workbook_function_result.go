package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookFunctionResult 
type WorkbookFunctionResult struct {
    Entity
    // The error property
    error *string
    // The value property
    value Jsonable
}
// NewWorkbookFunctionResult instantiates a new WorkbookFunctionResult and sets the default values.
func NewWorkbookFunctionResult()(*WorkbookFunctionResult) {
    m := &WorkbookFunctionResult{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookFunctionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFunctionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookFunctionResult(), nil
}
// GetError gets the error property value. The error property
func (m *WorkbookFunctionResult) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFunctionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *WorkbookFunctionResult) GetValue()(Jsonable) {
    return m.value
}
// Serialize serializes information the current object
func (m *WorkbookFunctionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. The error property
func (m *WorkbookFunctionResult) SetError(value *string)() {
    m.error = value
}
// SetValue sets the value property value. The value property
func (m *WorkbookFunctionResult) SetValue(value Jsonable)() {
    m.value = value
}
