package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodConfiguration 
type AuthenticationMethodConfiguration struct {
    Entity
    // The state of the policy. Possible values are: enabled, disabled.
    state *AuthenticationMethodState
}
// NewAuthenticationMethodConfiguration instantiates a new AuthenticationMethodConfiguration and sets the default values.
func NewAuthenticationMethodConfiguration()(*AuthenticationMethodConfiguration) {
    m := &AuthenticationMethodConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.emailAuthenticationMethodConfiguration":
                        return NewEmailAuthenticationMethodConfiguration(), nil
                    case "#microsoft.graph.fido2AuthenticationMethodConfiguration":
                        return NewFido2AuthenticationMethodConfiguration(), nil
                    case "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodConfiguration":
                        return NewMicrosoftAuthenticatorAuthenticationMethodConfiguration(), nil
                    case "#microsoft.graph.temporaryAccessPassAuthenticationMethodConfiguration":
                        return NewTemporaryAccessPassAuthenticationMethodConfiguration(), nil
                    case "#microsoft.graph.x509CertificateAuthenticationMethodConfiguration":
                        return NewX509CertificateAuthenticationMethodConfiguration(), nil
                }
            }
        }
    }
    return NewAuthenticationMethodConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AuthenticationMethodState))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state of the policy. Possible values are: enabled, disabled.
func (m *AuthenticationMethodConfiguration) GetState()(*AuthenticationMethodState) {
    return m.state
}
// Serialize serializes information the current object
func (m *AuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetState sets the state property value. The state of the policy. Possible values are: enabled, disabled.
func (m *AuthenticationMethodConfiguration) SetState(value *AuthenticationMethodState)() {
    m.state = value
}
