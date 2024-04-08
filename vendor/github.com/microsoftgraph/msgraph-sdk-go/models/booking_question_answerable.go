package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingQuestionAnswerable 
type BookingQuestionAnswerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnswer()(*string)
    GetAnswerInputType()(*AnswerInputType)
    GetAnswerOptions()([]string)
    GetIsRequired()(*bool)
    GetOdataType()(*string)
    GetQuestion()(*string)
    GetQuestionId()(*string)
    GetSelectedOptions()([]string)
    SetAnswer(value *string)()
    SetAnswerInputType(value *AnswerInputType)()
    SetAnswerOptions(value []string)()
    SetIsRequired(value *bool)()
    SetOdataType(value *string)()
    SetQuestion(value *string)()
    SetQuestionId(value *string)()
    SetSelectedOptions(value []string)()
}
