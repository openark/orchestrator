package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Fido2AuthenticationMethodConfigurationCollectionResponse 
type Fido2AuthenticationMethodConfigurationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Fido2AuthenticationMethodConfigurationable
}
// NewFido2AuthenticationMethodConfigurationCollectionResponse instantiates a new Fido2AuthenticationMethodConfigurationCollectionResponse and sets the default values.
func NewFido2AuthenticationMethodConfigurationCollectionResponse()(*Fido2AuthenticationMethodConfigurationCollectionResponse) {
    m := &Fido2AuthenticationMethodConfigurationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateFido2AuthenticationMethodConfigurationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFido2AuthenticationMethodConfigurationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFido2AuthenticationMethodConfigurationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Fido2AuthenticationMethodConfigurationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFido2AuthenticationMethodConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fido2AuthenticationMethodConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(Fido2AuthenticationMethodConfigurationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Fido2AuthenticationMethodConfigurationCollectionResponse) GetValue()([]Fido2AuthenticationMethodConfigurationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Fido2AuthenticationMethodConfigurationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *Fido2AuthenticationMethodConfigurationCollectionResponse) SetValue(value []Fido2AuthenticationMethodConfigurationable)() {
    m.value = value
}
