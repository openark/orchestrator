package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityApiConnector 
type IdentityApiConnector struct {
    Entity
    // The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client certificate are supported.
    authenticationConfiguration ApiAuthenticationConfigurationBaseable
    // The name of the API connector.
    displayName *string
    // The URL of the API endpoint to call.
    targetUrl *string
}
// NewIdentityApiConnector instantiates a new IdentityApiConnector and sets the default values.
func NewIdentityApiConnector()(*IdentityApiConnector) {
    m := &IdentityApiConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIdentityApiConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityApiConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityApiConnector(), nil
}
// GetAuthenticationConfiguration gets the authenticationConfiguration property value. The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client certificate are supported.
func (m *IdentityApiConnector) GetAuthenticationConfiguration()(ApiAuthenticationConfigurationBaseable) {
    return m.authenticationConfiguration
}
// GetDisplayName gets the displayName property value. The name of the API connector.
func (m *IdentityApiConnector) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityApiConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApiAuthenticationConfigurationBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationConfiguration(val.(ApiAuthenticationConfigurationBaseable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["targetUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetUrl(val)
        }
        return nil
    }
    return res
}
// GetTargetUrl gets the targetUrl property value. The URL of the API endpoint to call.
func (m *IdentityApiConnector) GetTargetUrl()(*string) {
    return m.targetUrl
}
// Serialize serializes information the current object
func (m *IdentityApiConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authenticationConfiguration", m.GetAuthenticationConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetUrl", m.GetTargetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationConfiguration sets the authenticationConfiguration property value. The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client certificate are supported.
func (m *IdentityApiConnector) SetAuthenticationConfiguration(value ApiAuthenticationConfigurationBaseable)() {
    m.authenticationConfiguration = value
}
// SetDisplayName sets the displayName property value. The name of the API connector.
func (m *IdentityApiConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTargetUrl sets the targetUrl property value. The URL of the API endpoint to call.
func (m *IdentityApiConnector) SetTargetUrl(value *string)() {
    m.targetUrl = value
}
