package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MultiValueLegacyExtendedProperty 
type MultiValueLegacyExtendedProperty struct {
    Entity
    // A collection of property values.
    value []string
}
// NewMultiValueLegacyExtendedProperty instantiates a new multiValueLegacyExtendedProperty and sets the default values.
func NewMultiValueLegacyExtendedProperty()(*MultiValueLegacyExtendedProperty) {
    m := &MultiValueLegacyExtendedProperty{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiValueLegacyExtendedProperty(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MultiValueLegacyExtendedProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. A collection of property values.
func (m *MultiValueLegacyExtendedProperty) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *MultiValueLegacyExtendedProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. A collection of property values.
func (m *MultiValueLegacyExtendedProperty) SetValue(value []string)() {
    m.value = value
}
