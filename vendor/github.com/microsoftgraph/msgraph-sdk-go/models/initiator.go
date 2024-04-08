package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Initiator 
type Initiator struct {
    Identity
    // Type of initiator. Possible values are: user, application, system, unknownFutureValue.
    initiatorType *InitiatorType
}
// NewInitiator instantiates a new Initiator and sets the default values.
func NewInitiator()(*Initiator) {
    m := &Initiator{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.initiator"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInitiatorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInitiatorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInitiator(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Initiator) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["initiatorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInitiatorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatorType(val.(*InitiatorType))
        }
        return nil
    }
    return res
}
// GetInitiatorType gets the initiatorType property value. Type of initiator. Possible values are: user, application, system, unknownFutureValue.
func (m *Initiator) GetInitiatorType()(*InitiatorType) {
    return m.initiatorType
}
// Serialize serializes information the current object
func (m *Initiator) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInitiatorType() != nil {
        cast := (*m.GetInitiatorType()).String()
        err = writer.WriteStringValue("initiatorType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiatorType sets the initiatorType property value. Type of initiator. Possible values are: user, application, system, unknownFutureValue.
func (m *Initiator) SetInitiatorType(value *InitiatorType)() {
    m.initiatorType = value
}
