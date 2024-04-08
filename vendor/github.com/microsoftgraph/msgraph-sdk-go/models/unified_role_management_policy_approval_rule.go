package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyApprovalRule 
type UnifiedRoleManagementPolicyApprovalRule struct {
    UnifiedRoleManagementPolicyRule
    // The settings for approval of the role assignment.
    setting ApprovalSettingsable
}
// NewUnifiedRoleManagementPolicyApprovalRule instantiates a new UnifiedRoleManagementPolicyApprovalRule and sets the default values.
func NewUnifiedRoleManagementPolicyApprovalRule()(*UnifiedRoleManagementPolicyApprovalRule) {
    m := &UnifiedRoleManagementPolicyApprovalRule{
        UnifiedRoleManagementPolicyRule: *NewUnifiedRoleManagementPolicyRule(),
    }
    odataTypeValue := "#microsoft.graph.unifiedRoleManagementPolicyApprovalRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUnifiedRoleManagementPolicyApprovalRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyApprovalRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyApprovalRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyApprovalRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementPolicyRule.GetFieldDeserializers()
    res["setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApprovalSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val.(ApprovalSettingsable))
        }
        return nil
    }
    return res
}
// GetSetting gets the setting property value. The settings for approval of the role assignment.
func (m *UnifiedRoleManagementPolicyApprovalRule) GetSetting()(ApprovalSettingsable) {
    return m.setting
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyApprovalRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementPolicyRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSetting sets the setting property value. The settings for approval of the role assignment.
func (m *UnifiedRoleManagementPolicyApprovalRule) SetSetting(value ApprovalSettingsable)() {
    m.setting = value
}
