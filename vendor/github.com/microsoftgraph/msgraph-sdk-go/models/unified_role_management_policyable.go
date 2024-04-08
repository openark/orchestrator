package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyable 
type UnifiedRoleManagementPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEffectiveRules()([]UnifiedRoleManagementPolicyRuleable)
    GetIsOrganizationDefault()(*bool)
    GetLastModifiedBy()(Identityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRules()([]UnifiedRoleManagementPolicyRuleable)
    GetScopeId()(*string)
    GetScopeType()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEffectiveRules(value []UnifiedRoleManagementPolicyRuleable)()
    SetIsOrganizationDefault(value *bool)()
    SetLastModifiedBy(value Identityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRules(value []UnifiedRoleManagementPolicyRuleable)()
    SetScopeId(value *string)()
    SetScopeType(value *string)()
}
