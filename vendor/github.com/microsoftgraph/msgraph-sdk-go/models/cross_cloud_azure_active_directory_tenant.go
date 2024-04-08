package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossCloudAzureActiveDirectoryTenant 
type CrossCloudAzureActiveDirectoryTenant struct {
    IdentitySource
    // The ID of the cloud where the tenant is located, one of microsoftonline.com, microsoftonline.us or partner.microsoftonline.cn. Read only.
    cloudInstance *string
    // The name of the Azure Active Directory tenant. Read only.
    displayName *string
    // The ID of the Azure Active Directory tenant. Read only.
    tenantId *string
}
// NewCrossCloudAzureActiveDirectoryTenant instantiates a new CrossCloudAzureActiveDirectoryTenant and sets the default values.
func NewCrossCloudAzureActiveDirectoryTenant()(*CrossCloudAzureActiveDirectoryTenant) {
    m := &CrossCloudAzureActiveDirectoryTenant{
        IdentitySource: *NewIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.crossCloudAzureActiveDirectoryTenant"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCrossCloudAzureActiveDirectoryTenantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossCloudAzureActiveDirectoryTenantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossCloudAzureActiveDirectoryTenant(), nil
}
// GetCloudInstance gets the cloudInstance property value. The ID of the cloud where the tenant is located, one of microsoftonline.com, microsoftonline.us or partner.microsoftonline.cn. Read only.
func (m *CrossCloudAzureActiveDirectoryTenant) GetCloudInstance()(*string) {
    return m.cloudInstance
}
// GetDisplayName gets the displayName property value. The name of the Azure Active Directory tenant. Read only.
func (m *CrossCloudAzureActiveDirectoryTenant) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossCloudAzureActiveDirectoryTenant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySource.GetFieldDeserializers()
    res["cloudInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudInstance(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. The ID of the Azure Active Directory tenant. Read only.
func (m *CrossCloudAzureActiveDirectoryTenant) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *CrossCloudAzureActiveDirectoryTenant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("cloudInstance", m.GetCloudInstance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudInstance sets the cloudInstance property value. The ID of the cloud where the tenant is located, one of microsoftonline.com, microsoftonline.us or partner.microsoftonline.cn. Read only.
func (m *CrossCloudAzureActiveDirectoryTenant) SetCloudInstance(value *string)() {
    m.cloudInstance = value
}
// SetDisplayName sets the displayName property value. The name of the Azure Active Directory tenant. Read only.
func (m *CrossCloudAzureActiveDirectoryTenant) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTenantId sets the tenantId property value. The ID of the Azure Active Directory tenant. Read only.
func (m *CrossCloudAzureActiveDirectoryTenant) SetTenantId(value *string)() {
    m.tenantId = value
}
