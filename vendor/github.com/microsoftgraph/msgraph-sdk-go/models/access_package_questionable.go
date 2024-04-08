package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageQuestionable 
type AccessPackageQuestionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsAnswerEditable()(*bool)
    GetIsRequired()(*bool)
    GetLocalizations()([]AccessPackageLocalizedTextable)
    GetSequence()(*int32)
    GetText()(*string)
    SetIsAnswerEditable(value *bool)()
    SetIsRequired(value *bool)()
    SetLocalizations(value []AccessPackageLocalizedTextable)()
    SetSequence(value *int32)()
    SetText(value *string)()
}
