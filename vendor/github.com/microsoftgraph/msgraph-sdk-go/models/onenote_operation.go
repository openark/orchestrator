package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteOperation 
type OnenoteOperation struct {
    Operation
    // The error returned by the operation.
    error OnenoteOperationErrorable
    // The operation percent complete if the operation is still in running status.
    percentComplete *string
    // The resource id.
    resourceId *string
    // The resource URI for the object. For example, the resource URI for a copied page or section.
    resourceLocation *string
}
// NewOnenoteOperation instantiates a new OnenoteOperation and sets the default values.
func NewOnenoteOperation()(*OnenoteOperation) {
    m := &OnenoteOperation{
        Operation: *NewOperation(),
    }
    return m
}
// CreateOnenoteOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenoteOperation(), nil
}
// GetError gets the error property value. The error returned by the operation.
func (m *OnenoteOperation) GetError()(OnenoteOperationErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Operation.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnenoteOperationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(OnenoteOperationErrorable))
        }
        return nil
    }
    res["percentComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentComplete(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
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
    return res
}
// GetPercentComplete gets the percentComplete property value. The operation percent complete if the operation is still in running status.
func (m *OnenoteOperation) GetPercentComplete()(*string) {
    return m.percentComplete
}
// GetResourceId gets the resourceId property value. The resource id.
func (m *OnenoteOperation) GetResourceId()(*string) {
    return m.resourceId
}
// GetResourceLocation gets the resourceLocation property value. The resource URI for the object. For example, the resource URI for a copied page or section.
func (m *OnenoteOperation) GetResourceLocation()(*string) {
    return m.resourceLocation
}
// Serialize serializes information the current object
func (m *OnenoteOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Operation.Serialize(writer)
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
        err = writer.WriteStringValue("percentComplete", m.GetPercentComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
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
    return nil
}
// SetError sets the error property value. The error returned by the operation.
func (m *OnenoteOperation) SetError(value OnenoteOperationErrorable)() {
    m.error = value
}
// SetPercentComplete sets the percentComplete property value. The operation percent complete if the operation is still in running status.
func (m *OnenoteOperation) SetPercentComplete(value *string)() {
    m.percentComplete = value
}
// SetResourceId sets the resourceId property value. The resource id.
func (m *OnenoteOperation) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetResourceLocation sets the resourceLocation property value. The resource URI for the object. For example, the resource URI for a copied page or section.
func (m *OnenoteOperation) SetResourceLocation(value *string)() {
    m.resourceLocation = value
}
