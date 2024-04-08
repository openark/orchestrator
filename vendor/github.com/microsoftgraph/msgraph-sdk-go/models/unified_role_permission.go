package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRolePermission 
type UnifiedRolePermission struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Set of tasks that can be performed on a resource. Required.
    allowedResourceActions []string
    // Optional constraints that must be met for the permission to be effective. Not supported for custom roles.
    condition *string
    // Set of tasks that may not be performed on a resource. Not yet supported.
    excludedResourceActions []string
    // The OdataType property
    odataType *string
}
// NewUnifiedRolePermission instantiates a new unifiedRolePermission and sets the default values.
func NewUnifiedRolePermission()(*UnifiedRolePermission) {
    m := &UnifiedRolePermission{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUnifiedRolePermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRolePermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRolePermission(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRolePermission) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowedResourceActions gets the allowedResourceActions property value. Set of tasks that can be performed on a resource. Required.
func (m *UnifiedRolePermission) GetAllowedResourceActions()([]string) {
    return m.allowedResourceActions
}
// GetCondition gets the condition property value. Optional constraints that must be met for the permission to be effective. Not supported for custom roles.
func (m *UnifiedRolePermission) GetCondition()(*string) {
    return m.condition
}
// GetExcludedResourceActions gets the excludedResourceActions property value. Set of tasks that may not be performed on a resource. Not yet supported.
func (m *UnifiedRolePermission) GetExcludedResourceActions()([]string) {
    return m.excludedResourceActions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRolePermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedResourceActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedResourceActions(res)
        }
        return nil
    }
    res["condition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val)
        }
        return nil
    }
    res["excludedResourceActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludedResourceActions(res)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UnifiedRolePermission) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *UnifiedRolePermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedResourceActions() != nil {
        err := writer.WriteCollectionOfStringValues("allowedResourceActions", m.GetAllowedResourceActions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("condition", m.GetCondition())
        if err != nil {
            return err
        }
    }
    if m.GetExcludedResourceActions() != nil {
        err := writer.WriteCollectionOfStringValues("excludedResourceActions", m.GetExcludedResourceActions())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRolePermission) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowedResourceActions sets the allowedResourceActions property value. Set of tasks that can be performed on a resource. Required.
func (m *UnifiedRolePermission) SetAllowedResourceActions(value []string)() {
    m.allowedResourceActions = value
}
// SetCondition sets the condition property value. Optional constraints that must be met for the permission to be effective. Not supported for custom roles.
func (m *UnifiedRolePermission) SetCondition(value *string)() {
    m.condition = value
}
// SetExcludedResourceActions sets the excludedResourceActions property value. Set of tasks that may not be performed on a resource. Not yet supported.
func (m *UnifiedRolePermission) SetExcludedResourceActions(value []string)() {
    m.excludedResourceActions = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UnifiedRolePermission) SetOdataType(value *string)() {
    m.odataType = value
}
