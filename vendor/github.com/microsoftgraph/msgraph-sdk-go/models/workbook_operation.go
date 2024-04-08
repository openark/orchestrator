package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookOperation 
type WorkbookOperation struct {
    Entity
    // The error returned by the operation.
    error WorkbookOperationErrorable
    // The resource URI for the result.
    resourceLocation *string
    // The status property
    status *WorkbookOperationStatus
}
// NewWorkbookOperation instantiates a new workbookOperation and sets the default values.
func NewWorkbookOperation()(*WorkbookOperation) {
    m := &WorkbookOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookOperation(), nil
}
// GetError gets the error property value. The error returned by the operation.
func (m *WorkbookOperation) GetError()(WorkbookOperationErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookOperationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(WorkbookOperationErrorable))
        }
        return nil
    }
    res["resourceLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceLocation(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkbookOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*WorkbookOperationStatus))
        }
        return nil
    }
    return res
}
// GetResourceLocation gets the resourceLocation property value. The resource URI for the result.
func (m *WorkbookOperation) GetResourceLocation()(*string) {
    return m.resourceLocation
}
// GetStatus gets the status property value. The status property
func (m *WorkbookOperation) GetStatus()(*WorkbookOperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *WorkbookOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceLocation", m.GetResourceLocation())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. The error returned by the operation.
func (m *WorkbookOperation) SetError(value WorkbookOperationErrorable)() {
    m.error = value
}
// SetResourceLocation sets the resourceLocation property value. The resource URI for the result.
func (m *WorkbookOperation) SetResourceLocation(value *string)() {
    m.resourceLocation = value
}
// SetStatus sets the status property value. The status property
func (m *WorkbookOperation) SetStatus(value *WorkbookOperationStatus)() {
    m.status = value
}
