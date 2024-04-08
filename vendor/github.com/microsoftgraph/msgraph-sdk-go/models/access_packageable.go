package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageable 
type AccessPackageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackagesIncompatibleWith()([]AccessPackageable)
    GetAssignmentPolicies()([]AccessPackageAssignmentPolicyable)
    GetCatalog()(AccessPackageCatalogable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIncompatibleAccessPackages()([]AccessPackageable)
    GetIncompatibleGroups()([]Groupable)
    GetIsHidden()(*bool)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessPackagesIncompatibleWith(value []AccessPackageable)()
    SetAssignmentPolicies(value []AccessPackageAssignmentPolicyable)()
    SetCatalog(value AccessPackageCatalogable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIncompatibleAccessPackages(value []AccessPackageable)()
    SetIncompatibleGroups(value []Groupable)()
    SetIsHidden(value *bool)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
