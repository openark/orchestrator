package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningSystem 
type ProvisioningSystem struct {
    Identity
    // Details of the system.
    details DetailsInfoable
}
// NewProvisioningSystem instantiates a new ProvisioningSystem and sets the default values.
func NewProvisioningSystem()(*ProvisioningSystem) {
    m := &ProvisioningSystem{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.provisioningSystem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProvisioningSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningSystem(), nil
}
// GetDetails gets the details property value. Details of the system.
func (m *ProvisioningSystem) GetDetails()(DetailsInfoable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetailsInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(DetailsInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProvisioningSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetails sets the details property value. Details of the system.
func (m *ProvisioningSystem) SetDetails(value DetailsInfoable)() {
    m.details = value
}
