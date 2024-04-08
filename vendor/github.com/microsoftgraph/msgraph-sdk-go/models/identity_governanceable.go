package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernanceable 
type IdentityGovernanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessReviews()(AccessReviewSetable)
    GetAppConsent()(AppConsentApprovalRouteable)
    GetEntitlementManagement()(EntitlementManagementable)
    GetOdataType()(*string)
    GetTermsOfUse()(TermsOfUseContainerable)
    SetAccessReviews(value AccessReviewSetable)()
    SetAppConsent(value AppConsentApprovalRouteable)()
    SetEntitlementManagement(value EntitlementManagementable)()
    SetOdataType(value *string)()
    SetTermsOfUse(value TermsOfUseContainerable)()
}
