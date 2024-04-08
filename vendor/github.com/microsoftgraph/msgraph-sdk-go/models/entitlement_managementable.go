package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementable 
type EntitlementManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageAssignmentApprovals()([]Approvalable)
    GetAccessPackages()([]AccessPackageable)
    GetAssignmentPolicies()([]AccessPackageAssignmentPolicyable)
    GetAssignmentRequests()([]AccessPackageAssignmentRequestable)
    GetAssignments()([]AccessPackageAssignmentable)
    GetCatalogs()([]AccessPackageCatalogable)
    GetConnectedOrganizations()([]ConnectedOrganizationable)
    GetSettings()(EntitlementManagementSettingsable)
    SetAccessPackageAssignmentApprovals(value []Approvalable)()
    SetAccessPackages(value []AccessPackageable)()
    SetAssignmentPolicies(value []AccessPackageAssignmentPolicyable)()
    SetAssignmentRequests(value []AccessPackageAssignmentRequestable)()
    SetAssignments(value []AccessPackageAssignmentable)()
    SetCatalogs(value []AccessPackageCatalogable)()
    SetConnectedOrganizations(value []ConnectedOrganizationable)()
    SetSettings(value EntitlementManagementSettingsable)()
}
