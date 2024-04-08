package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFirewallNetworkProfileable 
type WindowsFirewallNetworkProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizedApplicationRulesFromGroupPolicyMerged()(*bool)
    GetConnectionSecurityRulesFromGroupPolicyMerged()(*bool)
    GetFirewallEnabled()(*StateManagementSetting)
    GetGlobalPortRulesFromGroupPolicyMerged()(*bool)
    GetInboundConnectionsBlocked()(*bool)
    GetInboundNotificationsBlocked()(*bool)
    GetIncomingTrafficBlocked()(*bool)
    GetOdataType()(*string)
    GetOutboundConnectionsBlocked()(*bool)
    GetPolicyRulesFromGroupPolicyMerged()(*bool)
    GetSecuredPacketExemptionAllowed()(*bool)
    GetStealthModeBlocked()(*bool)
    GetUnicastResponsesToMulticastBroadcastsBlocked()(*bool)
    SetAuthorizedApplicationRulesFromGroupPolicyMerged(value *bool)()
    SetConnectionSecurityRulesFromGroupPolicyMerged(value *bool)()
    SetFirewallEnabled(value *StateManagementSetting)()
    SetGlobalPortRulesFromGroupPolicyMerged(value *bool)()
    SetInboundConnectionsBlocked(value *bool)()
    SetInboundNotificationsBlocked(value *bool)()
    SetIncomingTrafficBlocked(value *bool)()
    SetOdataType(value *string)()
    SetOutboundConnectionsBlocked(value *bool)()
    SetPolicyRulesFromGroupPolicyMerged(value *bool)()
    SetSecuredPacketExemptionAllowed(value *bool)()
    SetStealthModeBlocked(value *bool)()
    SetUnicastResponsesToMulticastBroadcastsBlocked(value *bool)()
}
