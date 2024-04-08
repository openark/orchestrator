package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingQuestionAnswer 
type BookingQuestionAnswer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The answer given by the user in case the answerInputType is text.
    answer *string
    // The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
    answerInputType *AnswerInputType
    // In case the answerInputType is radioButton, this will consists of a list of possible answer values.
    answerOptions []string
    // Indicates whether it is mandatory to answer the custom question.
    isRequired *bool
    // The OdataType property
    odataType *string
    // The question.
    question *string
    // The ID of the custom question.
    questionId *string
    // The answers selected by the user.
    selectedOptions []string
}
// NewBookingQuestionAnswer instantiates a new bookingQuestionAnswer and sets the default values.
func NewBookingQuestionAnswer()(*BookingQuestionAnswer) {
    m := &BookingQuestionAnswer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBookingQuestionAnswerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingQuestionAnswerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingQuestionAnswer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingQuestionAnswer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnswer gets the answer property value. The answer given by the user in case the answerInputType is text.
func (m *BookingQuestionAnswer) GetAnswer()(*string) {
    return m.answer
}
// GetAnswerInputType gets the answerInputType property value. The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
func (m *BookingQuestionAnswer) GetAnswerInputType()(*AnswerInputType) {
    return m.answerInputType
}
// GetAnswerOptions gets the answerOptions property value. In case the answerInputType is radioButton, this will consists of a list of possible answer values.
func (m *BookingQuestionAnswer) GetAnswerOptions()([]string) {
    return m.answerOptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingQuestionAnswer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["answer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnswer(val)
        }
        return nil
    }
    res["answerInputType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnswerInputType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnswerInputType(val.(*AnswerInputType))
        }
        return nil
    }
    res["answerOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAnswerOptions(res)
        }
        return nil
    }
    res["isRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["question"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuestion(val)
        }
        return nil
    }
    res["questionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuestionId(val)
        }
        return nil
    }
    res["selectedOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelectedOptions(res)
        }
        return nil
    }
    return res
}
// GetIsRequired gets the isRequired property value. Indicates whether it is mandatory to answer the custom question.
func (m *BookingQuestionAnswer) GetIsRequired()(*bool) {
    return m.isRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BookingQuestionAnswer) GetOdataType()(*string) {
    return m.odataType
}
// GetQuestion gets the question property value. The question.
func (m *BookingQuestionAnswer) GetQuestion()(*string) {
    return m.question
}
// GetQuestionId gets the questionId property value. The ID of the custom question.
func (m *BookingQuestionAnswer) GetQuestionId()(*string) {
    return m.questionId
}
// GetSelectedOptions gets the selectedOptions property value. The answers selected by the user.
func (m *BookingQuestionAnswer) GetSelectedOptions()([]string) {
    return m.selectedOptions
}
// Serialize serializes information the current object
func (m *BookingQuestionAnswer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("answer", m.GetAnswer())
        if err != nil {
            return err
        }
    }
    if m.GetAnswerInputType() != nil {
        cast := (*m.GetAnswerInputType()).String()
        err := writer.WriteStringValue("answerInputType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnswerOptions() != nil {
        err := writer.WriteCollectionOfStringValues("answerOptions", m.GetAnswerOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("question", m.GetQuestion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("questionId", m.GetQuestionId())
        if err != nil {
            return err
        }
    }
    if m.GetSelectedOptions() != nil {
        err := writer.WriteCollectionOfStringValues("selectedOptions", m.GetSelectedOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingQuestionAnswer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnswer sets the answer property value. The answer given by the user in case the answerInputType is text.
func (m *BookingQuestionAnswer) SetAnswer(value *string)() {
    m.answer = value
}
// SetAnswerInputType sets the answerInputType property value. The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
func (m *BookingQuestionAnswer) SetAnswerInputType(value *AnswerInputType)() {
    m.answerInputType = value
}
// SetAnswerOptions sets the answerOptions property value. In case the answerInputType is radioButton, this will consists of a list of possible answer values.
func (m *BookingQuestionAnswer) SetAnswerOptions(value []string)() {
    m.answerOptions = value
}
// SetIsRequired sets the isRequired property value. Indicates whether it is mandatory to answer the custom question.
func (m *BookingQuestionAnswer) SetIsRequired(value *bool)() {
    m.isRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BookingQuestionAnswer) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuestion sets the question property value. The question.
func (m *BookingQuestionAnswer) SetQuestion(value *string)() {
    m.question = value
}
// SetQuestionId sets the questionId property value. The ID of the custom question.
func (m *BookingQuestionAnswer) SetQuestionId(value *string)() {
    m.questionId = value
}
// SetSelectedOptions sets the selectedOptions property value. The answers selected by the user.
func (m *BookingQuestionAnswer) SetSelectedOptions(value []string)() {
    m.selectedOptions = value
}
