package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyConfigurationPartner 
type CrossTenantAccessPolicyConfigurationPartner struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Defines your partner-specific configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
    b2bCollaborationInbound CrossTenantAccessPolicyB2BSettingable
    // Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
    b2bCollaborationOutbound CrossTenantAccessPolicyB2BSettingable
    // Defines your partner-specific configuration for users from other organizations accessing your resources via Azure B2B direct connect.
    b2bDirectConnectInbound CrossTenantAccessPolicyB2BSettingable
    // Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
    b2bDirectConnectOutbound CrossTenantAccessPolicyB2BSettingable
    // Determines the partner-specific configuration for trusting other Conditional Access claims from external Azure AD organizations.
    inboundTrust CrossTenantAccessPolicyInboundTrustable
    // Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
    isServiceProvider *bool
    // The OdataType property
    odataType *string
    // The tenant identifier for the partner Azure AD organization. Read-only. Key.
    tenantId *string
}
// NewCrossTenantAccessPolicyConfigurationPartner instantiates a new crossTenantAccessPolicyConfigurationPartner and sets the default values.
func NewCrossTenantAccessPolicyConfigurationPartner()(*CrossTenantAccessPolicyConfigurationPartner) {
    m := &CrossTenantAccessPolicyConfigurationPartner{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyConfigurationPartnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyConfigurationPartner(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetB2bCollaborationInbound gets the b2bCollaborationInbound property value. Defines your partner-specific configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetB2bCollaborationInbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bCollaborationInbound
}
// GetB2bCollaborationOutbound gets the b2bCollaborationOutbound property value. Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetB2bCollaborationOutbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bCollaborationOutbound
}
// GetB2bDirectConnectInbound gets the b2bDirectConnectInbound property value. Defines your partner-specific configuration for users from other organizations accessing your resources via Azure B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetB2bDirectConnectInbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bDirectConnectInbound
}
// GetB2bDirectConnectOutbound gets the b2bDirectConnectOutbound property value. Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetB2bDirectConnectOutbound()(CrossTenantAccessPolicyB2BSettingable) {
    return m.b2bDirectConnectOutbound
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyConfigurationPartner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["isServiceProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsServiceProvider(val)
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
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetInboundTrust gets the inboundTrust property value. Determines the partner-specific configuration for trusting other Conditional Access claims from external Azure AD organizations.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetInboundTrust()(CrossTenantAccessPolicyInboundTrustable) {
    return m.inboundTrust
}
// GetIsServiceProvider gets the isServiceProvider property value. Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetIsServiceProvider()(*bool) {
    return m.isServiceProvider
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CrossTenantAccessPolicyConfigurationPartner) GetOdataType()(*string) {
    return m.odataType
}
// GetTenantId gets the tenantId property value. The tenant identifier for the partner Azure AD organization. Read-only. Key.
func (m *CrossTenantAccessPolicyConfigurationPartner) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyConfigurationPartner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("b2bCollaborationInbound", m.GetB2bCollaborationInbound())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("b2bCollaborationOutbound", m.GetB2bCollaborationOutbound())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("b2bDirectConnectInbound", m.GetB2bDirectConnectInbound())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("b2bDirectConnectOutbound", m.GetB2bDirectConnectOutbound())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("inboundTrust", m.GetInboundTrust())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isServiceProvider", m.GetIsServiceProvider())
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
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *CrossTenantAccessPolicyConfigurationPartner) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetB2bCollaborationInbound sets the b2bCollaborationInbound property value. Defines your partner-specific configuration for users from other organizations accessing your resources via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetB2bCollaborationInbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bCollaborationInbound = value
}
// SetB2bCollaborationOutbound sets the b2bCollaborationOutbound property value. Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B collaboration.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetB2bCollaborationOutbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bCollaborationOutbound = value
}
// SetB2bDirectConnectInbound sets the b2bDirectConnectInbound property value. Defines your partner-specific configuration for users from other organizations accessing your resources via Azure B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetB2bDirectConnectInbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bDirectConnectInbound = value
}
// SetB2bDirectConnectOutbound sets the b2bDirectConnectOutbound property value. Defines your partner-specific configuration for users in your organization going outbound to access resources in another organization via Azure AD B2B direct connect.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetB2bDirectConnectOutbound(value CrossTenantAccessPolicyB2BSettingable)() {
    m.b2bDirectConnectOutbound = value
}
// SetInboundTrust sets the inboundTrust property value. Determines the partner-specific configuration for trusting other Conditional Access claims from external Azure AD organizations.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetInboundTrust(value CrossTenantAccessPolicyInboundTrustable)() {
    m.inboundTrust = value
}
// SetIsServiceProvider sets the isServiceProvider property value. Identifies whether the partner-specific configuration is a Cloud Service Provider for your organization.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetIsServiceProvider(value *bool)() {
    m.isServiceProvider = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CrossTenantAccessPolicyConfigurationPartner) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenantId sets the tenantId property value. The tenant identifier for the partner Azure AD organization. Read-only. Key.
func (m *CrossTenantAccessPolicyConfigurationPartner) SetTenantId(value *string)() {
    m.tenantId = value
}
