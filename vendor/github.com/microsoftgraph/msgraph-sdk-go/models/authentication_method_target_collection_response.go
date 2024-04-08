package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodTargetCollectionResponse 
type AuthenticationMethodTargetCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []AuthenticationMethodTargetable
}
// NewAuthenticationMethodTargetCollectionResponse instantiates a new AuthenticationMethodTargetCollectionResponse and sets the default values.
func NewAuthenticationMethodTargetCollectionResponse()(*AuthenticationMethodTargetCollectionResponse) {
    m := &AuthenticationMethodTargetCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAuthenticationMethodTargetCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodTargetCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodTargetCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodTargetCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodTargetable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodTargetable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *AuthenticationMethodTargetCollectionResponse) GetValue()([]AuthenticationMethodTargetable) {
    return m.value
}
// Serialize serializes information the current object
func (m *AuthenticationMethodTargetCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AuthenticationMethodTargetCollectionResponse) SetValue(value []AuthenticationMethodTargetable)() {
    m.value = value
}
