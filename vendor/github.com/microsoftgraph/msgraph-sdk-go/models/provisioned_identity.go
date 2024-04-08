package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisionedIdentity 
type ProvisionedIdentity struct {
    Identity
    // Details of the identity.
    details DetailsInfoable
    // Type of identity that has been provisioned, such as 'user' or 'group'.
    identityType *string
}
// NewProvisionedIdentity instantiates a new ProvisionedIdentity and sets the default values.
func NewProvisionedIdentity()(*ProvisionedIdentity) {
    m := &ProvisionedIdentity{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.provisionedIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProvisionedIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisionedIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisionedIdentity(), nil
}
// GetDetails gets the details property value. Details of the identity.
func (m *ProvisionedIdentity) GetDetails()(DetailsInfoable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisionedIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["identityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityType(val)
        }
        return nil
    }
    return res
}
// GetIdentityType gets the identityType property value. Type of identity that has been provisioned, such as 'user' or 'group'.
func (m *ProvisionedIdentity) GetIdentityType()(*string) {
    return m.identityType
}
// Serialize serializes information the current object
func (m *ProvisionedIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("identityType", m.GetIdentityType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetails sets the details property value. Details of the identity.
func (m *ProvisionedIdentity) SetDetails(value DetailsInfoable)() {
    m.details = value
}
// SetIdentityType sets the identityType property value. Type of identity that has been provisioned, such as 'user' or 'group'.
func (m *ProvisionedIdentity) SetIdentityType(value *string)() {
    m.identityType = value
}
