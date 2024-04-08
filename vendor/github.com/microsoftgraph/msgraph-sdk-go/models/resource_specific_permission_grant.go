package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceSpecificPermissionGrant 
type ResourceSpecificPermissionGrant struct {
    DirectoryObject
    // ID of the service principal of the Azure AD app that has been granted access. Read-only.
    clientAppId *string
    // ID of the Azure AD app that has been granted access. Read-only.
    clientId *string
    // The name of the resource-specific permission. Read-only.
    permission *string
    // The type of permission. Possible values are: Application, Delegated. Read-only.
    permissionType *string
    // ID of the Azure AD app that is hosting the resource. Read-only.
    resourceAppId *string
}
// NewResourceSpecificPermissionGrant instantiates a new resourceSpecificPermissionGrant and sets the default values.
func NewResourceSpecificPermissionGrant()(*ResourceSpecificPermissionGrant) {
    m := &ResourceSpecificPermissionGrant{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.resourceSpecificPermissionGrant"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateResourceSpecificPermissionGrantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceSpecificPermissionGrantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceSpecificPermissionGrant(), nil
}
// GetClientAppId gets the clientAppId property value. ID of the service principal of the Azure AD app that has been granted access. Read-only.
func (m *ResourceSpecificPermissionGrant) GetClientAppId()(*string) {
    return m.clientAppId
}
// GetClientId gets the clientId property value. ID of the Azure AD app that has been granted access. Read-only.
func (m *ResourceSpecificPermissionGrant) GetClientId()(*string) {
    return m.clientId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceSpecificPermissionGrant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["clientAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAppId(val)
        }
        return nil
    }
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val)
        }
        return nil
    }
    res["permissionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionType(val)
        }
        return nil
    }
    res["resourceAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAppId(val)
        }
        return nil
    }
    return res
}
// GetPermission gets the permission property value. The name of the resource-specific permission. Read-only.
func (m *ResourceSpecificPermissionGrant) GetPermission()(*string) {
    return m.permission
}
// GetPermissionType gets the permissionType property value. The type of permission. Possible values are: Application, Delegated. Read-only.
func (m *ResourceSpecificPermissionGrant) GetPermissionType()(*string) {
    return m.permissionType
}
// GetResourceAppId gets the resourceAppId property value. ID of the Azure AD app that is hosting the resource. Read-only.
func (m *ResourceSpecificPermissionGrant) GetResourceAppId()(*string) {
    return m.resourceAppId
}
// Serialize serializes information the current object
func (m *ResourceSpecificPermissionGrant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientAppId", m.GetClientAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permission", m.GetPermission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionType", m.GetPermissionType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceAppId", m.GetResourceAppId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientAppId sets the clientAppId property value. ID of the service principal of the Azure AD app that has been granted access. Read-only.
func (m *ResourceSpecificPermissionGrant) SetClientAppId(value *string)() {
    m.clientAppId = value
}
// SetClientId sets the clientId property value. ID of the Azure AD app that has been granted access. Read-only.
func (m *ResourceSpecificPermissionGrant) SetClientId(value *string)() {
    m.clientId = value
}
// SetPermission sets the permission property value. The name of the resource-specific permission. Read-only.
func (m *ResourceSpecificPermissionGrant) SetPermission(value *string)() {
    m.permission = value
}
// SetPermissionType sets the permissionType property value. The type of permission. Possible values are: Application, Delegated. Read-only.
func (m *ResourceSpecificPermissionGrant) SetPermissionType(value *string)() {
    m.permissionType = value
}
// SetResourceAppId sets the resourceAppId property value. ID of the Azure AD app that is hosting the resource. Read-only.
func (m *ResourceSpecificPermissionGrant) SetResourceAppId(value *string)() {
    m.resourceAppId = value
}
