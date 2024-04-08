package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SoftwareOathAuthenticationMethod 
type SoftwareOathAuthenticationMethod struct {
    AuthenticationMethod
    // The secret key of the method. Always returns null.
    secretKey *string
}
// NewSoftwareOathAuthenticationMethod instantiates a new SoftwareOathAuthenticationMethod and sets the default values.
func NewSoftwareOathAuthenticationMethod()(*SoftwareOathAuthenticationMethod) {
    m := &SoftwareOathAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    odataTypeValue := "#microsoft.graph.softwareOathAuthenticationMethod"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSoftwareOathAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSoftwareOathAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSoftwareOathAuthenticationMethod(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SoftwareOathAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["secretKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretKey(val)
        }
        return nil
    }
    return res
}
// GetSecretKey gets the secretKey property value. The secret key of the method. Always returns null.
func (m *SoftwareOathAuthenticationMethod) GetSecretKey()(*string) {
    return m.secretKey
}
// Serialize serializes information the current object
func (m *SoftwareOathAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("secretKey", m.GetSecretKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSecretKey sets the secretKey property value. The secret key of the method. Always returns null.
func (m *SoftwareOathAuthenticationMethod) SetSecretKey(value *string)() {
    m.secretKey = value
}
