package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFirewallNetworkProfile windows Firewall Profile Policies.
type WindowsFirewallNetworkProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Configures the firewall to merge authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
    authorizedApplicationRulesFromGroupPolicyMerged *bool
    // Configures the firewall to merge connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
    connectionSecurityRulesFromGroupPolicyMerged *bool
    // State Management Setting.
    firewallEnabled *StateManagementSetting
    // Configures the firewall to merge global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
    globalPortRulesFromGroupPolicyMerged *bool
    // Configures the firewall to block all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
    inboundConnectionsBlocked *bool
    // Prevents the firewall from displaying notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
    inboundNotificationsBlocked *bool
    // Configures the firewall to block all incoming traffic regardless of other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
    incomingTrafficBlocked *bool
    // The OdataType property
    odataType *string
    // Configures the firewall to block all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
    outboundConnectionsBlocked *bool
    // Configures the firewall to merge Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
    policyRulesFromGroupPolicyMerged *bool
    // Configures the firewall to allow the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
    securedPacketExemptionAllowed *bool
    // Prevent the server from operating in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
    stealthModeBlocked *bool
    // Configures the firewall to block unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
    unicastResponsesToMulticastBroadcastsBlocked *bool
}
// NewWindowsFirewallNetworkProfile instantiates a new windowsFirewallNetworkProfile and sets the default values.
func NewWindowsFirewallNetworkProfile()(*WindowsFirewallNetworkProfile) {
    m := &WindowsFirewallNetworkProfile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsFirewallNetworkProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFirewallNetworkProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsFirewallNetworkProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsFirewallNetworkProfile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthorizedApplicationRulesFromGroupPolicyMerged gets the authorizedApplicationRulesFromGroupPolicyMerged property value. Configures the firewall to merge authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetAuthorizedApplicationRulesFromGroupPolicyMerged()(*bool) {
    return m.authorizedApplicationRulesFromGroupPolicyMerged
}
// GetConnectionSecurityRulesFromGroupPolicyMerged gets the connectionSecurityRulesFromGroupPolicyMerged property value. Configures the firewall to merge connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetConnectionSecurityRulesFromGroupPolicyMerged()(*bool) {
    return m.connectionSecurityRulesFromGroupPolicyMerged
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsFirewallNetworkProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authorizedApplicationRulesFromGroupPolicyMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedApplicationRulesFromGroupPolicyMerged(val)
        }
        return nil
    }
    res["connectionSecurityRulesFromGroupPolicyMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionSecurityRulesFromGroupPolicyMerged(val)
        }
        return nil
    }
    res["firewallEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStateManagementSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallEnabled(val.(*StateManagementSetting))
        }
        return nil
    }
    res["globalPortRulesFromGroupPolicyMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobalPortRulesFromGroupPolicyMerged(val)
        }
        return nil
    }
    res["inboundConnectionsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundConnectionsBlocked(val)
        }
        return nil
    }
    res["inboundNotificationsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundNotificationsBlocked(val)
        }
        return nil
    }
    res["incomingTrafficBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncomingTrafficBlocked(val)
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
    res["outboundConnectionsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutboundConnectionsBlocked(val)
        }
        return nil
    }
    res["policyRulesFromGroupPolicyMerged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyRulesFromGroupPolicyMerged(val)
        }
        return nil
    }
    res["securedPacketExemptionAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecuredPacketExemptionAllowed(val)
        }
        return nil
    }
    res["stealthModeBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStealthModeBlocked(val)
        }
        return nil
    }
    res["unicastResponsesToMulticastBroadcastsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnicastResponsesToMulticastBroadcastsBlocked(val)
        }
        return nil
    }
    return res
}
// GetFirewallEnabled gets the firewallEnabled property value. State Management Setting.
func (m *WindowsFirewallNetworkProfile) GetFirewallEnabled()(*StateManagementSetting) {
    return m.firewallEnabled
}
// GetGlobalPortRulesFromGroupPolicyMerged gets the globalPortRulesFromGroupPolicyMerged property value. Configures the firewall to merge global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetGlobalPortRulesFromGroupPolicyMerged()(*bool) {
    return m.globalPortRulesFromGroupPolicyMerged
}
// GetInboundConnectionsBlocked gets the inboundConnectionsBlocked property value. Configures the firewall to block all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetInboundConnectionsBlocked()(*bool) {
    return m.inboundConnectionsBlocked
}
// GetInboundNotificationsBlocked gets the inboundNotificationsBlocked property value. Prevents the firewall from displaying notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetInboundNotificationsBlocked()(*bool) {
    return m.inboundNotificationsBlocked
}
// GetIncomingTrafficBlocked gets the incomingTrafficBlocked property value. Configures the firewall to block all incoming traffic regardless of other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetIncomingTrafficBlocked()(*bool) {
    return m.incomingTrafficBlocked
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsFirewallNetworkProfile) GetOdataType()(*string) {
    return m.odataType
}
// GetOutboundConnectionsBlocked gets the outboundConnectionsBlocked property value. Configures the firewall to block all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
func (m *WindowsFirewallNetworkProfile) GetOutboundConnectionsBlocked()(*bool) {
    return m.outboundConnectionsBlocked
}
// GetPolicyRulesFromGroupPolicyMerged gets the policyRulesFromGroupPolicyMerged property value. Configures the firewall to merge Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) GetPolicyRulesFromGroupPolicyMerged()(*bool) {
    return m.policyRulesFromGroupPolicyMerged
}
// GetSecuredPacketExemptionAllowed gets the securedPacketExemptionAllowed property value. Configures the firewall to allow the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
func (m *WindowsFirewallNetworkProfile) GetSecuredPacketExemptionAllowed()(*bool) {
    return m.securedPacketExemptionAllowed
}
// GetStealthModeBlocked gets the stealthModeBlocked property value. Prevent the server from operating in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetStealthModeBlocked()(*bool) {
    return m.stealthModeBlocked
}
// GetUnicastResponsesToMulticastBroadcastsBlocked gets the unicastResponsesToMulticastBroadcastsBlocked property value. Configures the firewall to block unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) GetUnicastResponsesToMulticastBroadcastsBlocked()(*bool) {
    return m.unicastResponsesToMulticastBroadcastsBlocked
}
// Serialize serializes information the current object
func (m *WindowsFirewallNetworkProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("authorizedApplicationRulesFromGroupPolicyMerged", m.GetAuthorizedApplicationRulesFromGroupPolicyMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("connectionSecurityRulesFromGroupPolicyMerged", m.GetConnectionSecurityRulesFromGroupPolicyMerged())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallEnabled() != nil {
        cast := (*m.GetFirewallEnabled()).String()
        err := writer.WriteStringValue("firewallEnabled", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("globalPortRulesFromGroupPolicyMerged", m.GetGlobalPortRulesFromGroupPolicyMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inboundConnectionsBlocked", m.GetInboundConnectionsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inboundNotificationsBlocked", m.GetInboundNotificationsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("incomingTrafficBlocked", m.GetIncomingTrafficBlocked())
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
        err := writer.WriteBoolValue("outboundConnectionsBlocked", m.GetOutboundConnectionsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("policyRulesFromGroupPolicyMerged", m.GetPolicyRulesFromGroupPolicyMerged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("securedPacketExemptionAllowed", m.GetSecuredPacketExemptionAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("stealthModeBlocked", m.GetStealthModeBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("unicastResponsesToMulticastBroadcastsBlocked", m.GetUnicastResponsesToMulticastBroadcastsBlocked())
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
func (m *WindowsFirewallNetworkProfile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthorizedApplicationRulesFromGroupPolicyMerged sets the authorizedApplicationRulesFromGroupPolicyMerged property value. Configures the firewall to merge authorized application rules from group policy with those from local store instead of ignoring the local store rules. When AuthorizedApplicationRulesFromGroupPolicyNotMerged and AuthorizedApplicationRulesFromGroupPolicyMerged are both true, AuthorizedApplicationRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetAuthorizedApplicationRulesFromGroupPolicyMerged(value *bool)() {
    m.authorizedApplicationRulesFromGroupPolicyMerged = value
}
// SetConnectionSecurityRulesFromGroupPolicyMerged sets the connectionSecurityRulesFromGroupPolicyMerged property value. Configures the firewall to merge connection security rules from group policy with those from local store instead of ignoring the local store rules. When ConnectionSecurityRulesFromGroupPolicyNotMerged and ConnectionSecurityRulesFromGroupPolicyMerged are both true, ConnectionSecurityRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetConnectionSecurityRulesFromGroupPolicyMerged(value *bool)() {
    m.connectionSecurityRulesFromGroupPolicyMerged = value
}
// SetFirewallEnabled sets the firewallEnabled property value. State Management Setting.
func (m *WindowsFirewallNetworkProfile) SetFirewallEnabled(value *StateManagementSetting)() {
    m.firewallEnabled = value
}
// SetGlobalPortRulesFromGroupPolicyMerged sets the globalPortRulesFromGroupPolicyMerged property value. Configures the firewall to merge global port rules from group policy with those from local store instead of ignoring the local store rules. When GlobalPortRulesFromGroupPolicyNotMerged and GlobalPortRulesFromGroupPolicyMerged are both true, GlobalPortRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetGlobalPortRulesFromGroupPolicyMerged(value *bool)() {
    m.globalPortRulesFromGroupPolicyMerged = value
}
// SetInboundConnectionsBlocked sets the inboundConnectionsBlocked property value. Configures the firewall to block all incoming connections by default. When InboundConnectionsRequired and InboundConnectionsBlocked are both true, InboundConnectionsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetInboundConnectionsBlocked(value *bool)() {
    m.inboundConnectionsBlocked = value
}
// SetInboundNotificationsBlocked sets the inboundNotificationsBlocked property value. Prevents the firewall from displaying notifications when an application is blocked from listening on a port. When InboundNotificationsRequired and InboundNotificationsBlocked are both true, InboundNotificationsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetInboundNotificationsBlocked(value *bool)() {
    m.inboundNotificationsBlocked = value
}
// SetIncomingTrafficBlocked sets the incomingTrafficBlocked property value. Configures the firewall to block all incoming traffic regardless of other policy settings. When IncomingTrafficRequired and IncomingTrafficBlocked are both true, IncomingTrafficBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetIncomingTrafficBlocked(value *bool)() {
    m.incomingTrafficBlocked = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsFirewallNetworkProfile) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOutboundConnectionsBlocked sets the outboundConnectionsBlocked property value. Configures the firewall to block all outgoing connections by default. When OutboundConnectionsRequired and OutboundConnectionsBlocked are both true, OutboundConnectionsBlocked takes priority. This setting will get applied to Windows releases version 1809 and above.
func (m *WindowsFirewallNetworkProfile) SetOutboundConnectionsBlocked(value *bool)() {
    m.outboundConnectionsBlocked = value
}
// SetPolicyRulesFromGroupPolicyMerged sets the policyRulesFromGroupPolicyMerged property value. Configures the firewall to merge Firewall Rule policies from group policy with those from local store instead of ignoring the local store rules. When PolicyRulesFromGroupPolicyNotMerged and PolicyRulesFromGroupPolicyMerged are both true, PolicyRulesFromGroupPolicyMerged takes priority.
func (m *WindowsFirewallNetworkProfile) SetPolicyRulesFromGroupPolicyMerged(value *bool)() {
    m.policyRulesFromGroupPolicyMerged = value
}
// SetSecuredPacketExemptionAllowed sets the securedPacketExemptionAllowed property value. Configures the firewall to allow the host computer to respond to unsolicited network traffic of that traffic is secured by IPSec even when stealthModeBlocked is set to true. When SecuredPacketExemptionBlocked and SecuredPacketExemptionAllowed are both true, SecuredPacketExemptionAllowed takes priority.
func (m *WindowsFirewallNetworkProfile) SetSecuredPacketExemptionAllowed(value *bool)() {
    m.securedPacketExemptionAllowed = value
}
// SetStealthModeBlocked sets the stealthModeBlocked property value. Prevent the server from operating in stealth mode. When StealthModeRequired and StealthModeBlocked are both true, StealthModeBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetStealthModeBlocked(value *bool)() {
    m.stealthModeBlocked = value
}
// SetUnicastResponsesToMulticastBroadcastsBlocked sets the unicastResponsesToMulticastBroadcastsBlocked property value. Configures the firewall to block unicast responses to multicast broadcast traffic. When UnicastResponsesToMulticastBroadcastsRequired and UnicastResponsesToMulticastBroadcastsBlocked are both true, UnicastResponsesToMulticastBroadcastsBlocked takes priority.
func (m *WindowsFirewallNetworkProfile) SetUnicastResponsesToMulticastBroadcastsBlocked(value *bool)() {
    m.unicastResponsesToMulticastBroadcastsBlocked = value
}
