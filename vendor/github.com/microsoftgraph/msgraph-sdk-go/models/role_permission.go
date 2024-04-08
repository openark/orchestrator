package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RolePermission contains the set of ResourceActions determining the allowed and not allowed permissions for each role.
type RolePermission struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Resource Actions each containing a set of allowed and not allowed permissions.
    resourceActions []ResourceActionable
}
// NewRolePermission instantiates a new rolePermission and sets the default values.
func NewRolePermission()(*RolePermission) {
    m := &RolePermission{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRolePermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRolePermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRolePermission(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RolePermission) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RolePermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["resourceActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceActionable, len(val))
            for i, v := range val {
                res[i] = v.(ResourceActionable)
            }
            m.SetResourceActions(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RolePermission) GetOdataType()(*string) {
    return m.odataType
}
// GetResourceActions gets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) GetResourceActions()([]ResourceActionable) {
    return m.resourceActions
}
// Serialize serializes information the current object
func (m *RolePermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetResourceActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceActions()))
        for i, v := range m.GetResourceActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("resourceActions", cast)
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
func (m *RolePermission) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RolePermission) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResourceActions sets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) SetResourceActions(value []ResourceActionable)() {
    m.resourceActions = value
}
