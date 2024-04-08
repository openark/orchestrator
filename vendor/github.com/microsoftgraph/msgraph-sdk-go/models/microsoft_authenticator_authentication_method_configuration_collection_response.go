package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse 
type MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []MicrosoftAuthenticatorAuthenticationMethodConfigurationable
}
// NewMicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse instantiates a new MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse and sets the default values.
func NewMicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse()(*MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse) {
    m := &MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftAuthenticatorAuthenticationMethodConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftAuthenticatorAuthenticationMethodConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftAuthenticatorAuthenticationMethodConfigurationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse) GetValue()([]MicrosoftAuthenticatorAuthenticationMethodConfigurationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MicrosoftAuthenticatorAuthenticationMethodConfigurationCollectionResponse) SetValue(value []MicrosoftAuthenticatorAuthenticationMethodConfigurationable)() {
    m.value = value
}
