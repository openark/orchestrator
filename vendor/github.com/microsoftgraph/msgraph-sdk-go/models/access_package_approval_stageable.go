package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageApprovalStageable 
type AccessPackageApprovalStageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDurationBeforeAutomaticDenial()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDurationBeforeEscalation()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetEscalationApprovers()([]SubjectSetable)
    GetFallbackEscalationApprovers()([]SubjectSetable)
    GetFallbackPrimaryApprovers()([]SubjectSetable)
    GetIsApproverJustificationRequired()(*bool)
    GetIsEscalationEnabled()(*bool)
    GetOdataType()(*string)
    GetPrimaryApprovers()([]SubjectSetable)
    SetDurationBeforeAutomaticDenial(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDurationBeforeEscalation(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetEscalationApprovers(value []SubjectSetable)()
    SetFallbackEscalationApprovers(value []SubjectSetable)()
    SetFallbackPrimaryApprovers(value []SubjectSetable)()
    SetIsApproverJustificationRequired(value *bool)()
    SetIsEscalationEnabled(value *bool)()
    SetOdataType(value *string)()
    SetPrimaryApprovers(value []SubjectSetable)()
}
