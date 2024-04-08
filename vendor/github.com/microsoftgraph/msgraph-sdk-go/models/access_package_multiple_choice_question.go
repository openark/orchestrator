package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageMultipleChoiceQuestion 
type AccessPackageMultipleChoiceQuestion struct {
    AccessPackageQuestion
    // List of answer choices.
    choices []AccessPackageAnswerChoiceable
    // Indicates whether requestor can select multiple choices as their answer.
    isMultipleSelectionAllowed *bool
}
// NewAccessPackageMultipleChoiceQuestion instantiates a new AccessPackageMultipleChoiceQuestion and sets the default values.
func NewAccessPackageMultipleChoiceQuestion()(*AccessPackageMultipleChoiceQuestion) {
    m := &AccessPackageMultipleChoiceQuestion{
        AccessPackageQuestion: *NewAccessPackageQuestion(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageMultipleChoiceQuestion"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessPackageMultipleChoiceQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageMultipleChoiceQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageMultipleChoiceQuestion(), nil
}
// GetChoices gets the choices property value. List of answer choices.
func (m *AccessPackageMultipleChoiceQuestion) GetChoices()([]AccessPackageAnswerChoiceable) {
    return m.choices
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageMultipleChoiceQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageQuestion.GetFieldDeserializers()
    res["choices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAnswerChoiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAnswerChoiceable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAnswerChoiceable)
            }
            m.SetChoices(res)
        }
        return nil
    }
    res["isMultipleSelectionAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMultipleSelectionAllowed(val)
        }
        return nil
    }
    return res
}
// GetIsMultipleSelectionAllowed gets the isMultipleSelectionAllowed property value. Indicates whether requestor can select multiple choices as their answer.
func (m *AccessPackageMultipleChoiceQuestion) GetIsMultipleSelectionAllowed()(*bool) {
    return m.isMultipleSelectionAllowed
}
// Serialize serializes information the current object
func (m *AccessPackageMultipleChoiceQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessPackageQuestion.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChoices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChoices()))
        for i, v := range m.GetChoices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("choices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMultipleSelectionAllowed", m.GetIsMultipleSelectionAllowed())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChoices sets the choices property value. List of answer choices.
func (m *AccessPackageMultipleChoiceQuestion) SetChoices(value []AccessPackageAnswerChoiceable)() {
    m.choices = value
}
// SetIsMultipleSelectionAllowed sets the isMultipleSelectionAllowed property value. Indicates whether requestor can select multiple choices as their answer.
func (m *AccessPackageMultipleChoiceQuestion) SetIsMultipleSelectionAllowed(value *bool)() {
    m.isMultipleSelectionAllowed = value
}
