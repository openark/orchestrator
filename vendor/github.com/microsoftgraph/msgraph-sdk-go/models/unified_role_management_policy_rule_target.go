package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyRuleTarget 
type UnifiedRoleManagementPolicyRuleTarget struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type of caller that's the target of the policy rule. Allowed values are: None, Admin, EndUser.
    caller *string
    // The list of role settings that are enforced and cannot be overridden by child scopes. Use All for all settings.
    enforcedSettings []string
    // The list of role settings that can be inherited by child scopes. Use All for all settings.
    inheritableSettings []string
    // The role assignment type that's the target of policy rule. Allowed values are: Eligibility, Assignment.
    level *string
    // The OdataType property
    odataType *string
    // The role management operations that are the target of the policy rule. Allowed values are: All, Activate, Deactivate, Assign, Update, Remove, Extend, Renew.
    operations []UnifiedRoleManagementPolicyRuleTargetOperations
    // The targetObjects property
    targetObjects []DirectoryObjectable
}
// NewUnifiedRoleManagementPolicyRuleTarget instantiates a new unifiedRoleManagementPolicyRuleTarget and sets the default values.
func NewUnifiedRoleManagementPolicyRuleTarget()(*UnifiedRoleManagementPolicyRuleTarget) {
    m := &UnifiedRoleManagementPolicyRuleTarget{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUnifiedRoleManagementPolicyRuleTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyRuleTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyRuleTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCaller gets the caller property value. The type of caller that's the target of the policy rule. Allowed values are: None, Admin, EndUser.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetCaller()(*string) {
    return m.caller
}
// GetEnforcedSettings gets the enforcedSettings property value. The list of role settings that are enforced and cannot be overridden by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetEnforcedSettings()([]string) {
    return m.enforcedSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyRuleTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["caller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaller(val)
        }
        return nil
    }
    res["enforcedSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnforcedSettings(res)
        }
        return nil
    }
    res["inheritableSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInheritableSettings(res)
        }
        return nil
    }
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
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
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseUnifiedRoleManagementPolicyRuleTargetOperations)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyRuleTargetOperations, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnifiedRoleManagementPolicyRuleTargetOperations))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["targetObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetTargetObjects(res)
        }
        return nil
    }
    return res
}
// GetInheritableSettings gets the inheritableSettings property value. The list of role settings that can be inherited by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetInheritableSettings()([]string) {
    return m.inheritableSettings
}
// GetLevel gets the level property value. The role assignment type that's the target of policy rule. Allowed values are: Eligibility, Assignment.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetLevel()(*string) {
    return m.level
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UnifiedRoleManagementPolicyRuleTarget) GetOdataType()(*string) {
    return m.odataType
}
// GetOperations gets the operations property value. The role management operations that are the target of the policy rule. Allowed values are: All, Activate, Deactivate, Assign, Update, Remove, Extend, Renew.
func (m *UnifiedRoleManagementPolicyRuleTarget) GetOperations()([]UnifiedRoleManagementPolicyRuleTargetOperations) {
    return m.operations
}
// GetTargetObjects gets the targetObjects property value. The targetObjects property
func (m *UnifiedRoleManagementPolicyRuleTarget) GetTargetObjects()([]DirectoryObjectable) {
    return m.targetObjects
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyRuleTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("caller", m.GetCaller())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcedSettings() != nil {
        err := writer.WriteCollectionOfStringValues("enforcedSettings", m.GetEnforcedSettings())
        if err != nil {
            return err
        }
    }
    if m.GetInheritableSettings() != nil {
        err := writer.WriteCollectionOfStringValues("inheritableSettings", m.GetInheritableSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("level", m.GetLevel())
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
    if m.GetOperations() != nil {
        err := writer.WriteCollectionOfStringValues("operations", SerializeUnifiedRoleManagementPolicyRuleTargetOperations(m.GetOperations()))
        if err != nil {
            return err
        }
    }
    if m.GetTargetObjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetObjects()))
        for i, v := range m.GetTargetObjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("targetObjects", cast)
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
func (m *UnifiedRoleManagementPolicyRuleTarget) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCaller sets the caller property value. The type of caller that's the target of the policy rule. Allowed values are: None, Admin, EndUser.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetCaller(value *string)() {
    m.caller = value
}
// SetEnforcedSettings sets the enforcedSettings property value. The list of role settings that are enforced and cannot be overridden by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetEnforcedSettings(value []string)() {
    m.enforcedSettings = value
}
// SetInheritableSettings sets the inheritableSettings property value. The list of role settings that can be inherited by child scopes. Use All for all settings.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetInheritableSettings(value []string)() {
    m.inheritableSettings = value
}
// SetLevel sets the level property value. The role assignment type that's the target of policy rule. Allowed values are: Eligibility, Assignment.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetLevel(value *string)() {
    m.level = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UnifiedRoleManagementPolicyRuleTarget) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperations sets the operations property value. The role management operations that are the target of the policy rule. Allowed values are: All, Activate, Deactivate, Assign, Update, Remove, Extend, Renew.
func (m *UnifiedRoleManagementPolicyRuleTarget) SetOperations(value []UnifiedRoleManagementPolicyRuleTargetOperations)() {
    m.operations = value
}
// SetTargetObjects sets the targetObjects property value. The targetObjects property
func (m *UnifiedRoleManagementPolicyRuleTarget) SetTargetObjects(value []DirectoryObjectable)() {
    m.targetObjects = value
}
