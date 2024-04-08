package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintServiceEndpoint 
type PrintServiceEndpoint struct {
    Entity
    // A human-readable display name for the endpoint.
    displayName *string
    // The URI that can be used to access the service.
    uri *string
}
// NewPrintServiceEndpoint instantiates a new printServiceEndpoint and sets the default values.
func NewPrintServiceEndpoint()(*PrintServiceEndpoint) {
    m := &PrintServiceEndpoint{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintServiceEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintServiceEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintServiceEndpoint(), nil
}
// GetDisplayName gets the displayName property value. A human-readable display name for the endpoint.
func (m *PrintServiceEndpoint) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintServiceEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetUri gets the uri property value. The URI that can be used to access the service.
func (m *PrintServiceEndpoint) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *PrintServiceEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uri", m.GetUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. A human-readable display name for the endpoint.
func (m *PrintServiceEndpoint) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetUri sets the uri property value. The URI that can be used to access the service.
func (m *PrintServiceEndpoint) SetUri(value *string)() {
    m.uri = value
}
