package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAnswerString 
type AccessPackageAnswerString struct {
    AccessPackageAnswer
    // The value stored on the requestor's user profile, if this answer is configured to be stored as a specific attribute.
    value *string
}
// NewAccessPackageAnswerString instantiates a new AccessPackageAnswerString and sets the default values.
func NewAccessPackageAnswerString()(*AccessPackageAnswerString) {
    m := &AccessPackageAnswerString{
        AccessPackageAnswer: *NewAccessPackageAnswer(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageAnswerString"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAccessPackageAnswerStringFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAnswerStringFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAnswerString(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAnswerString) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageAnswer.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value stored on the requestor's user profile, if this answer is configured to be stored as a specific attribute.
func (m *AccessPackageAnswerString) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *AccessPackageAnswerString) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessPackageAnswer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value stored on the requestor's user profile, if this answer is configured to be stored as a specific attribute.
func (m *AccessPackageAnswerString) SetValue(value *string)() {
    m.value = value
}
