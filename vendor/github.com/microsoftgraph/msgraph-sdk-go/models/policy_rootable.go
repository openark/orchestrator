package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PolicyRootable 
type PolicyRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityBasedTimeoutPolicies()([]ActivityBasedTimeoutPolicyable)
    GetAdminConsentRequestPolicy()(AdminConsentRequestPolicyable)
    GetAuthenticationFlowsPolicy()(AuthenticationFlowsPolicyable)
    GetAuthenticationMethodsPolicy()(AuthenticationMethodsPolicyable)
    GetAuthorizationPolicy()(AuthorizationPolicyable)
    GetClaimsMappingPolicies()([]ClaimsMappingPolicyable)
    GetConditionalAccessPolicies()([]ConditionalAccessPolicyable)
    GetCrossTenantAccessPolicy()(CrossTenantAccessPolicyable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable)
    GetIdentitySecurityDefaultsEnforcementPolicy()(IdentitySecurityDefaultsEnforcementPolicyable)
    GetPermissionGrantPolicies()([]PermissionGrantPolicyable)
    GetRoleManagementPolicies()([]UnifiedRoleManagementPolicyable)
    GetRoleManagementPolicyAssignments()([]UnifiedRoleManagementPolicyAssignmentable)
    GetTokenIssuancePolicies()([]TokenIssuancePolicyable)
    GetTokenLifetimePolicies()([]TokenLifetimePolicyable)
    SetActivityBasedTimeoutPolicies(value []ActivityBasedTimeoutPolicyable)()
    SetAdminConsentRequestPolicy(value AdminConsentRequestPolicyable)()
    SetAuthenticationFlowsPolicy(value AuthenticationFlowsPolicyable)()
    SetAuthenticationMethodsPolicy(value AuthenticationMethodsPolicyable)()
    SetAuthorizationPolicy(value AuthorizationPolicyable)()
    SetClaimsMappingPolicies(value []ClaimsMappingPolicyable)()
    SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)()
    SetCrossTenantAccessPolicy(value CrossTenantAccessPolicyable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)()
    SetIdentitySecurityDefaultsEnforcementPolicy(value IdentitySecurityDefaultsEnforcementPolicyable)()
    SetPermissionGrantPolicies(value []PermissionGrantPolicyable)()
    SetRoleManagementPolicies(value []UnifiedRoleManagementPolicyable)()
    SetRoleManagementPolicyAssignments(value []UnifiedRoleManagementPolicyAssignmentable)()
    SetTokenIssuancePolicies(value []TokenIssuancePolicyable)()
    SetTokenLifetimePolicies(value []TokenLifetimePolicyable)()
}
