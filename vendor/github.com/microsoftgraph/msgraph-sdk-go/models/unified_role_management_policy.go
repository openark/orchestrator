package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicy 
type UnifiedRoleManagementPolicy struct {
    Entity
    // Description for the policy.
    description *string
    // Display name for the policy.
    displayName *string
    // The list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
    effectiveRules []UnifiedRoleManagementPolicyRuleable
    // This can only be set to true for a single tenant-wide policy which will apply to all scopes and roles. Set the scopeId to / and scopeType to Directory. Supports $filter (eq, ne).
    isOrganizationDefault *bool
    // The identity who last modified the role setting.
    lastModifiedBy Identityable
    // The time when the role setting was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of rules like approval rules and expiration rules. Supports $expand.
    rules []UnifiedRoleManagementPolicyRuleable
    // The identifier of the scope where the policy is created. Can be / for the tenant or a group ID. Required.
    scopeId *string
    // The type of the scope where the policy is created. One of Directory, DirectoryRole. Required.
    scopeType *string
}
// NewUnifiedRoleManagementPolicy instantiates a new unifiedRoleManagementPolicy and sets the default values.
func NewUnifiedRoleManagementPolicy()(*UnifiedRoleManagementPolicy) {
    m := &UnifiedRoleManagementPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicy(), nil
}
// GetDescription gets the description property value. Description for the policy.
func (m *UnifiedRoleManagementPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name for the policy.
func (m *UnifiedRoleManagementPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetEffectiveRules gets the effectiveRules property value. The list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
func (m *UnifiedRoleManagementPolicy) GetEffectiveRules()([]UnifiedRoleManagementPolicyRuleable) {
    return m.effectiveRules
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["effectiveRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyRuleable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementPolicyRuleable)
            }
            m.SetEffectiveRules(res)
        }
        return nil
    }
    res["isOrganizationDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizationDefault(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(Identityable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementPolicyRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementPolicyRuleable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleManagementPolicyRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    res["scopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeId(val)
        }
        return nil
    }
    res["scopeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeType(val)
        }
        return nil
    }
    return res
}
// GetIsOrganizationDefault gets the isOrganizationDefault property value. This can only be set to true for a single tenant-wide policy which will apply to all scopes and roles. Set the scopeId to / and scopeType to Directory. Supports $filter (eq, ne).
func (m *UnifiedRoleManagementPolicy) GetIsOrganizationDefault()(*bool) {
    return m.isOrganizationDefault
}
// GetLastModifiedBy gets the lastModifiedBy property value. The identity who last modified the role setting.
func (m *UnifiedRoleManagementPolicy) GetLastModifiedBy()(Identityable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The time when the role setting was last modified.
func (m *UnifiedRoleManagementPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRules gets the rules property value. The collection of rules like approval rules and expiration rules. Supports $expand.
func (m *UnifiedRoleManagementPolicy) GetRules()([]UnifiedRoleManagementPolicyRuleable) {
    return m.rules
}
// GetScopeId gets the scopeId property value. The identifier of the scope where the policy is created. Can be / for the tenant or a group ID. Required.
func (m *UnifiedRoleManagementPolicy) GetScopeId()(*string) {
    return m.scopeId
}
// GetScopeType gets the scopeType property value. The type of the scope where the policy is created. One of Directory, DirectoryRole. Required.
func (m *UnifiedRoleManagementPolicy) GetScopeType()(*string) {
    return m.scopeType
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetEffectiveRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEffectiveRules()))
        for i, v := range m.GetEffectiveRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("effectiveRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizationDefault", m.GetIsOrganizationDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeId", m.GetScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scopeType", m.GetScopeType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description for the policy.
func (m *UnifiedRoleManagementPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name for the policy.
func (m *UnifiedRoleManagementPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEffectiveRules sets the effectiveRules property value. The list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a rule to disable approval. Supports $expand.
func (m *UnifiedRoleManagementPolicy) SetEffectiveRules(value []UnifiedRoleManagementPolicyRuleable)() {
    m.effectiveRules = value
}
// SetIsOrganizationDefault sets the isOrganizationDefault property value. This can only be set to true for a single tenant-wide policy which will apply to all scopes and roles. Set the scopeId to / and scopeType to Directory. Supports $filter (eq, ne).
func (m *UnifiedRoleManagementPolicy) SetIsOrganizationDefault(value *bool)() {
    m.isOrganizationDefault = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The identity who last modified the role setting.
func (m *UnifiedRoleManagementPolicy) SetLastModifiedBy(value Identityable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The time when the role setting was last modified.
func (m *UnifiedRoleManagementPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRules sets the rules property value. The collection of rules like approval rules and expiration rules. Supports $expand.
func (m *UnifiedRoleManagementPolicy) SetRules(value []UnifiedRoleManagementPolicyRuleable)() {
    m.rules = value
}
// SetScopeId sets the scopeId property value. The identifier of the scope where the policy is created. Can be / for the tenant or a group ID. Required.
func (m *UnifiedRoleManagementPolicy) SetScopeId(value *string)() {
    m.scopeId = value
}
// SetScopeType sets the scopeType property value. The type of the scope where the policy is created. One of Directory, DirectoryRole. Required.
func (m *UnifiedRoleManagementPolicy) SetScopeType(value *string)() {
    m.scopeType = value
}
