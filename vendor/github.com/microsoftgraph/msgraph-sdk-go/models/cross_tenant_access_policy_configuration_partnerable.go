package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyConfigurationPartnerable 
type CrossTenantAccessPolicyConfigurationPartnerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetB2bCollaborationInbound()(CrossTenantAccessPolicyB2BSettingable)
    GetB2bCollaborationOutbound()(CrossTenantAccessPolicyB2BSettingable)
    GetB2bDirectConnectInbound()(CrossTenantAccessPolicyB2BSettingable)
    GetB2bDirectConnectOutbound()(CrossTenantAccessPolicyB2BSettingable)
    GetInboundTrust()(CrossTenantAccessPolicyInboundTrustable)
    GetIsServiceProvider()(*bool)
    GetOdataType()(*string)
    GetTenantId()(*string)
    SetB2bCollaborationInbound(value CrossTenantAccessPolicyB2BSettingable)()
    SetB2bCollaborationOutbound(value CrossTenantAccessPolicyB2BSettingable)()
    SetB2bDirectConnectInbound(value CrossTenantAccessPolicyB2BSettingable)()
    SetB2bDirectConnectOutbound(value CrossTenantAccessPolicyB2BSettingable)()
    SetInboundTrust(value CrossTenantAccessPolicyInboundTrustable)()
    SetIsServiceProvider(value *bool)()
    SetOdataType(value *string)()
    SetTenantId(value *string)()
}
