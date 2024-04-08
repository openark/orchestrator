package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomQuestion represents a custom question of the business.
type BookingCustomQuestion struct {
    Entity
    // The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
    answerInputType *AnswerInputType
    // List of possible answer values.
    answerOptions []string
    // The question.
    displayName *string
}
// NewBookingCustomQuestion instantiates a new bookingCustomQuestion and sets the default values.
func NewBookingCustomQuestion()(*BookingCustomQuestion) {
    m := &BookingCustomQuestion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingCustomQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingCustomQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingCustomQuestion(), nil
}
// GetAnswerInputType gets the answerInputType property value. The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
func (m *BookingCustomQuestion) GetAnswerInputType()(*AnswerInputType) {
    return m.answerInputType
}
// GetAnswerOptions gets the answerOptions property value. List of possible answer values.
func (m *BookingCustomQuestion) GetAnswerOptions()([]string) {
    return m.answerOptions
}
// GetDisplayName gets the displayName property value. The question.
func (m *BookingCustomQuestion) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *BookingCustomQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnswerInputType() != nil {
        cast := (*m.GetAnswerInputType()).String()
        err = writer.WriteStringValue("answerInputType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnswerOptions() != nil {
        err = writer.WriteCollectionOfStringValues("answerOptions", m.GetAnswerOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnswerInputType sets the answerInputType property value. The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
func (m *BookingCustomQuestion) SetAnswerInputType(value *AnswerInputType)() {
    m.answerInputType = value
}
// SetAnswerOptions sets the answerOptions property value. List of possible answer values.
func (m *BookingCustomQuestion) SetAnswerOptions(value []string)() {
    m.answerOptions = value
}
// SetDisplayName sets the displayName property value. The question.
func (m *BookingCustomQuestion) SetDisplayName(value *string)() {
    m.displayName = value
}
