package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageRuleable 
type MessageRuleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()(MessageRuleActionsable)
    GetConditions()(MessageRulePredicatesable)
    GetDisplayName()(*string)
    GetExceptions()(MessageRulePredicatesable)
    GetHasError()(*bool)
    GetIsEnabled()(*bool)
    GetIsReadOnly()(*bool)
    GetSequence()(*int32)
    SetActions(value MessageRuleActionsable)()
    SetConditions(value MessageRulePredicatesable)()
    SetDisplayName(value *string)()
    SetExceptions(value MessageRulePredicatesable)()
    SetHasError(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsReadOnly(value *bool)()
    SetSequence(value *int32)()
}
