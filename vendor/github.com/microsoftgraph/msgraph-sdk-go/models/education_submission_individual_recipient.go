package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSubmissionIndividualRecipient 
type EducationSubmissionIndividualRecipient struct {
    EducationSubmissionRecipient
    // User ID of the user to whom the submission is assigned.
    userId *string
}
// NewEducationSubmissionIndividualRecipient instantiates a new EducationSubmissionIndividualRecipient and sets the default values.
func NewEducationSubmissionIndividualRecipient()(*EducationSubmissionIndividualRecipient) {
    m := &EducationSubmissionIndividualRecipient{
        EducationSubmissionRecipient: *NewEducationSubmissionRecipient(),
    }
    odataTypeValue := "#microsoft.graph.educationSubmissionIndividualRecipient"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationSubmissionIndividualRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSubmissionIndividualRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSubmissionIndividualRecipient(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSubmissionIndividualRecipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSubmissionRecipient.GetFieldDeserializers()
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetUserId gets the userId property value. User ID of the user to whom the submission is assigned.
func (m *EducationSubmissionIndividualRecipient) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *EducationSubmissionIndividualRecipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSubmissionRecipient.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserId sets the userId property value. User ID of the user to whom the submission is assigned.
func (m *EducationSubmissionIndividualRecipient) SetUserId(value *string)() {
    m.userId = value
}
