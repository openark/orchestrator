package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyExpirationRule 
type UnifiedRoleManagementPolicyExpirationRule struct {
    UnifiedRoleManagementPolicyRule
    // Indicates whether expiration is required or if it's a permanently active assignment or eligibility.
    isExpirationRequired *bool
    // The maximum duration allowed for eligibility or assignment which is not permanent. Required when isExpirationRequired is true.
    maximumDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewUnifiedRoleManagementPolicyExpirationRule instantiates a new UnifiedRoleManagementPolicyExpirationRule and sets the default values.
func NewUnifiedRoleManagementPolicyExpirationRule()(*UnifiedRoleManagementPolicyExpirationRule) {
    m := &UnifiedRoleManagementPolicyExpirationRule{
        UnifiedRoleManagementPolicyRule: *NewUnifiedRoleManagementPolicyRule(),
    }
    odataTypeValue := "#microsoft.graph.unifiedRoleManagementPolicyExpirationRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUnifiedRoleManagementPolicyExpirationRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyExpirationRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyExpirationRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyExpirationRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementPolicyRule.GetFieldDeserializers()
    res["isExpirationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExpirationRequired(val)
        }
        return nil
    }
    res["maximumDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumDuration(val)
        }
        return nil
    }
    return res
}
// GetIsExpirationRequired gets the isExpirationRequired property value. Indicates whether expiration is required or if it's a permanently active assignment or eligibility.
func (m *UnifiedRoleManagementPolicyExpirationRule) GetIsExpirationRequired()(*bool) {
    return m.isExpirationRequired
}
// GetMaximumDuration gets the maximumDuration property value. The maximum duration allowed for eligibility or assignment which is not permanent. Required when isExpirationRequired is true.
func (m *UnifiedRoleManagementPolicyExpirationRule) GetMaximumDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.maximumDuration
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyExpirationRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementPolicyRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isExpirationRequired", m.GetIsExpirationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("maximumDuration", m.GetMaximumDuration())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsExpirationRequired sets the isExpirationRequired property value. Indicates whether expiration is required or if it's a permanently active assignment or eligibility.
func (m *UnifiedRoleManagementPolicyExpirationRule) SetIsExpirationRequired(value *bool)() {
    m.isExpirationRequired = value
}
// SetMaximumDuration sets the maximumDuration property value. The maximum duration allowed for eligibility or assignment which is not permanent. Required when isExpirationRequired is true.
func (m *UnifiedRoleManagementPolicyExpirationRule) SetMaximumDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.maximumDuration = value
}
