package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentPolicyable 
type AccessPackageAssignmentPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAllowedTargetScope()(*AllowedTargetScope)
    GetAutomaticRequestSettings()(AccessPackageAutomaticRequestSettingsable)
    GetCatalog()(AccessPackageCatalogable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExpiration()(ExpirationPatternable)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetQuestions()([]AccessPackageQuestionable)
    GetRequestApprovalSettings()(AccessPackageAssignmentApprovalSettingsable)
    GetRequestorSettings()(AccessPackageAssignmentRequestorSettingsable)
    GetReviewSettings()(AccessPackageAssignmentReviewSettingsable)
    GetSpecificAllowedTargets()([]SubjectSetable)
    SetAccessPackage(value AccessPackageable)()
    SetAllowedTargetScope(value *AllowedTargetScope)()
    SetAutomaticRequestSettings(value AccessPackageAutomaticRequestSettingsable)()
    SetCatalog(value AccessPackageCatalogable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExpiration(value ExpirationPatternable)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetQuestions(value []AccessPackageQuestionable)()
    SetRequestApprovalSettings(value AccessPackageAssignmentApprovalSettingsable)()
    SetRequestorSettings(value AccessPackageAssignmentRequestorSettingsable)()
    SetReviewSettings(value AccessPackageAssignmentReviewSettingsable)()
    SetSpecificAllowedTargets(value []SubjectSetable)()
}
