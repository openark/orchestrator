package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequestRequirementsable 
type AccessPackageAssignmentRequestRequirementsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowCustomAssignmentSchedule()(*bool)
    GetIsApprovalRequiredForAdd()(*bool)
    GetIsApprovalRequiredForUpdate()(*bool)
    GetOdataType()(*string)
    GetPolicyDescription()(*string)
    GetPolicyDisplayName()(*string)
    GetPolicyId()(*string)
    GetQuestions()([]AccessPackageQuestionable)
    GetSchedule()(EntitlementManagementScheduleable)
    SetAllowCustomAssignmentSchedule(value *bool)()
    SetIsApprovalRequiredForAdd(value *bool)()
    SetIsApprovalRequiredForUpdate(value *bool)()
    SetOdataType(value *string)()
    SetPolicyDescription(value *string)()
    SetPolicyDisplayName(value *string)()
    SetPolicyId(value *string)()
    SetQuestions(value []AccessPackageQuestionable)()
    SetSchedule(value EntitlementManagementScheduleable)()
}
