package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityProvider 
type IdentityProvider struct {
    Entity
    // The client ID for the application. This is the client ID obtained when registering the application with the identity provider. Required. Not nullable.
    clientId *string
    // The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.  Required. Not nullable.
    clientSecret *string
    // The display name of the identity provider. Not nullable.
    name *string
    // The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect. Not nullable.
    typeEscaped *string
}
// NewIdentityProvider instantiates a new identityProvider and sets the default values.
func NewIdentityProvider()(*IdentityProvider) {
    m := &IdentityProvider{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIdentityProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityProvider(), nil
}
// GetClientId gets the clientId property value. The client ID for the application. This is the client ID obtained when registering the application with the identity provider. Required. Not nullable.
func (m *IdentityProvider) GetClientId()(*string) {
    return m.clientId
}
// GetClientSecret gets the clientSecret property value. The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.  Required. Not nullable.
func (m *IdentityProvider) GetClientSecret()(*string) {
    return m.clientSecret
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The display name of the identity provider. Not nullable.
func (m *IdentityProvider) GetName()(*string) {
    return m.name
}
// GetType gets the type property value. The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect. Not nullable.
func (m *IdentityProvider) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *IdentityProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientId sets the clientId property value. The client ID for the application. This is the client ID obtained when registering the application with the identity provider. Required. Not nullable.
func (m *IdentityProvider) SetClientId(value *string)() {
    m.clientId = value
}
// SetClientSecret sets the clientSecret property value. The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.  Required. Not nullable.
func (m *IdentityProvider) SetClientSecret(value *string)() {
    m.clientSecret = value
}
// SetName sets the name property value. The display name of the identity provider. Not nullable.
func (m *IdentityProvider) SetName(value *string)() {
    m.name = value
}
// SetType sets the type property value. The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect. Not nullable.
func (m *IdentityProvider) SetType(value *string)() {
    m.typeEscaped = value
}
