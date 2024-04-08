package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyConfigurationDefault 
type CrossTenantAccessPolicyConfigurationDefault struct {
    Entity
    // Defines your default configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
    b2bCollaborationInbound CrossTenantAccessPolicyB2BSettingable
    // Defines your default configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
    b2bCollaborationOutbound CrossTenantAccessPolicyB2BSettingable
    // Defines your default configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
    b2bDirectConnectInbound CrossTenantAccessPolicyB2BSettingable
    // Defines your default configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
    b2bDirectConnectOutbound CrossTenantAccessPolicyB2BSettingable
    // Determines the default configuration for trusting other Conditional Access claims from external Azure AD organizations.
    inboundTrust CrossTenantAccessPolicyInboundTrustable
    // If true, the default configuration is set to the system default configuration. If false, the default settings have been customized.
    isServiceDefault *bool
}
// NewCrossTenantAccessPolicyConfigurationDefault instantiates a new crossTenantAccessPolicyConfigurationDefault and sets the default values.
func NewCrossTenantAccessPolicyConfigurationDefault()(*CrossTenantAccessPolicyConfigurationDefault) {
    m := &CrossTenantAccessPolicyConfigurationDefault{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCrossTenantAccessPolicyConfigurationDefaultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyConfigurationDefaultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyConfigurationDefault(), nil
}
// GetB2bCollaborationInbound gets the b2bCollaborationInbound property value. Defines your default configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationDefault) GetB2bCollaborationInbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bCollaborationInbound
}
// GetB2bCollaborationOutbound gets the b2bCollaborationOutbound property value. Defines your default configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationDefault) GetB2bCollaborationOutbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bCollaborationOutbound
}
// GetB2bDirectConnectInbound gets the b2bDirectConnectInbound property value. Defines your default configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationDefault) GetB2bDirectConnectInbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bDirectConnectInbound
}
// GetB2bDirectConnectOutbound gets the b2bDirectConnectOutbound property value. Defines your default configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationDefault) GetB2bDirectConnectOutbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bDirectConnectOutbound
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyConfigurationDefault) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["b2bCollaborationInbound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bCollaborationInbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bCollaborationOutbound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bCollaborationOutbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bDirectConnectInbound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bDirectConnectInbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["b2bDirectConnectOutbound"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyB2BSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB2bDirectConnectOutbound(val.(CrossTenantAccessPolicyB2BSettingable))
        }
        return nil
    }
    res["inboundTrust"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCrossTenantAccessPolicyInboundTrustFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundTrust(val.(CrossTenantAccessPolicyInboundTrustable))
        }
        return nil
    }
    res["isServiceDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsServiceDefault(val)
        }
        return nil
    }
    return res
}
// GetInboundTrust gets the inboundTrust property value. Determines the default configuration for trusting other Conditional Access claims from external Azure AD organizations.
func (m *CrossTenantAccessPolicyConfigurationDefault) GetInboundTrust()(CrossTenantAccessPolicyInboundTrustable) {
    return m.inboundTrust
}
// GetIsServiceDefault gets the isServiceDefault property value. If true, the default configuration is set to the system default configuration. If false, the default settings have been customized.
func (m *CrossTenantAccessPolicyConfigurationDefault) GetIsServiceDefault()(*bool) {
    return m.isServiceDefault
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyConfigurationDefault) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("b2bCollaborationInbound", m.GetB2bCollaborationInbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bCollaborationOutbound", m.GetB2bCollaborationOutbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bDirectConnectInbound", m.GetB2bDirectConnectInbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("b2bDirectConnectOutbound", m.GetB2bDirectConnectOutbound())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inboundTrust", m.GetInboundTrust())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isServiceDefault", m.GetIsServiceDefault())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetB2bCollaborationInbound sets the b2bCollaborationInbound property value. Defines your default configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationDefault) SetB2bCollaborationInbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bCollaborationInbound = value
}
// SetB2bCollaborationOutbound sets the b2bCollaborationOutbound property value. Defines your default configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationDefault) SetB2bCollaborationOutbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bCollaborationOutbound = value
}
// SetB2bDirectConnectInbound sets the b2bDirectConnectInbound property value. Defines your default configuration for users from other organizations accessing your resources via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationDefault) SetB2bDirectConnectInbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bDirectConnectInbound = value
}
// SetB2bDirectConnectOutbound sets the b2bDirectConnectOutbound property value. Defines your default configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationDefault) SetB2bDirectConnectOutbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bDirectConnectOutbound = value
}
// SetInboundTrust sets the inboundTrust property value. Determines the default configuration for trusting other Conditional Access claims from external Azure AD organizations.
func (m *CrossTenantAccessPolicyConfigurationDefault) SetInboundTrust(value CrossTenantAccessPolicyInboundTrustable)() {
    m.inboundTrust = value
}
// SetIsServiceDefault sets the isServiceDefault property value. If true, the default configuration is set to the system default configuration. If false, the default settings have been customized.
func (m *CrossTenantAccessPolicyConfigurationDefault) SetIsServiceDefault(value *bool)() {
    m.isServiceDefault = value
}
