package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalIdentity 
type ServicePrincipalIdentity struct {
    Identity
    // The application identifier of the service principal.
    appId *string
}
// NewServicePrincipalIdentity instantiates a new ServicePrincipalIdentity and sets the default values.
func NewServicePrincipalIdentity()(*ServicePrincipalIdentity) {
    m := &ServicePrincipalIdentity{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.servicePrincipalIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServicePrincipalIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalIdentity(), nil
}
// GetAppId gets the appId property value. The application identifier of the service principal.
func (m *ServicePrincipalIdentity) GetAppId()(*string) {
    return m.appId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServicePrincipalIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The application identifier of the service principal.
func (m *ServicePrincipalIdentity) SetAppId(value *string)() {
    m.appId = value
}
