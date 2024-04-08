package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityUserFlowAttributeAssignmentable 
type IdentityUserFlowAttributeAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetIsOptional()(*bool)
    GetRequiresVerification()(*bool)
    GetUserAttribute()(IdentityUserFlowAttributeable)
    GetUserAttributeValues()([]UserAttributeValuesItemable)
    GetUserInputType()(*IdentityUserFlowAttributeInputType)
    SetDisplayName(value *string)()
    SetIsOptional(value *bool)()
    SetRequiresVerification(value *bool)()
    SetUserAttribute(value IdentityUserFlowAttributeable)()
    SetUserAttributeValues(value []UserAttributeValuesItemable)()
    SetUserInputType(value *IdentityUserFlowAttributeInputType)()
}
