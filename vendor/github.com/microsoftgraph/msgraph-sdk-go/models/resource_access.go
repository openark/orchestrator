package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceAccess 
type ResourceAccess struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of an app role or delegated permission exposed by the resource application. For delegated permissions, this should match the id property of one of the delegated permissions in the oauth2PermissionScopes collection of the resource application's service principal. For app roles (application permissions), this should match the id property of an app role in the appRoles collection of the resource application's service principal.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The OdataType property
    odataType *string
    // Specifies whether the id property references a delegated permission or an app role (application permission). The possible values are: Scope (for delegated permissions) or Role (for app roles).
    typeEscaped *string
}
// NewResourceAccess instantiates a new resourceAccess and sets the default values.
func NewResourceAccess()(*ResourceAccess) {
    m := &ResourceAccess{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResourceAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceAccess(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceAccess) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier of an app role or delegated permission exposed by the resource application. For delegated permissions, this should match the id property of one of the delegated permissions in the oauth2PermissionScopes collection of the resource application's service principal. For app roles (application permissions), this should match the id property of an app role in the appRoles collection of the resource application's service principal.
func (m *ResourceAccess) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ResourceAccess) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. Specifies whether the id property references a delegated permission or an app role (application permission). The possible values are: Scope (for delegated permissions) or Role (for app roles).
func (m *ResourceAccess) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ResourceAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceAccess) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The unique identifier of an app role or delegated permission exposed by the resource application. For delegated permissions, this should match the id property of one of the delegated permissions in the oauth2PermissionScopes collection of the resource application's service principal. For app roles (application permissions), this should match the id property of an app role in the appRoles collection of the resource application's service principal.
func (m *ResourceAccess) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ResourceAccess) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. Specifies whether the id property references a delegated permission or an app role (application permission). The possible values are: Scope (for delegated permissions) or Role (for app roles).
func (m *ResourceAccess) SetType(value *string)() {
    m.typeEscaped = value
}
