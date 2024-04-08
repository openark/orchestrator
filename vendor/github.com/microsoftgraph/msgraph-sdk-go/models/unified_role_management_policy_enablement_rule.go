package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyEnablementRule 
type UnifiedRoleManagementPolicyEnablementRule struct {
    UnifiedRoleManagementPolicyRule
    // The collection of rules that are enabled for this policy rule. For example, MultiFactorAuthentication, Ticketing, and Justification.
    enabledRules []string
}
// NewUnifiedRoleManagementPolicyEnablementRule instantiates a new UnifiedRoleManagementPolicyEnablementRule and sets the default values.
func NewUnifiedRoleManagementPolicyEnablementRule()(*UnifiedRoleManagementPolicyEnablementRule) {
    m := &UnifiedRoleManagementPolicyEnablementRule{
        UnifiedRoleManagementPolicyRule: *NewUnifiedRoleManagementPolicyRule(),
    }
    odataTypeValue := "#microsoft.graph.unifiedRoleManagementPolicyEnablementRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUnifiedRoleManagementPolicyEnablementRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyEnablementRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyEnablementRule(), nil
}
// GetEnabledRules gets the enabledRules property value. The collection of rules that are enabled for this policy rule. For example, MultiFactorAuthentication, Ticketing, and Justification.
func (m *UnifiedRoleManagementPolicyEnablementRule) GetEnabledRules()([]string) {
    return m.enabledRules
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyEnablementRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementPolicyRule.GetFieldDeserializers()
    res["enabledRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnabledRules(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyEnablementRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementPolicyRule.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEnabledRules() != nil {
        err = writer.WriteCollectionOfStringValues("enabledRules", m.GetEnabledRules())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnabledRules sets the enabledRules property value. The collection of rules that are enabled for this policy rule. For example, MultiFactorAuthentication, Ticketing, and Justification.
func (m *UnifiedRoleManagementPolicyEnablementRule) SetEnabledRules(value []string)() {
    m.enabledRules = value
}
