package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Endpoint 
type Endpoint struct {
    DirectoryObject
    // The capability property
    capability *string
    // The providerId property
    providerId *string
    // The providerName property
    providerName *string
    // The providerResourceId property
    providerResourceId *string
    // The uri property
    uri *string
}
// NewEndpoint instantiates a new endpoint and sets the default values.
func NewEndpoint()(*Endpoint) {
    m := &Endpoint{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.endpoint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEndpoint(), nil
}
// GetCapability gets the capability property value. The capability property
func (m *Endpoint) GetCapability()(*string) {
    return m.capability
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Endpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["capability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapability(val)
        }
        return nil
    }
    res["providerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderId(val)
        }
        return nil
    }
    res["providerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderName(val)
        }
        return nil
    }
    res["providerResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderResourceId(val)
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
// GetProviderId gets the providerId property value. The providerId property
func (m *Endpoint) GetProviderId()(*string) {
    return m.providerId
}
// GetProviderName gets the providerName property value. The providerName property
func (m *Endpoint) GetProviderName()(*string) {
    return m.providerName
}
// GetProviderResourceId gets the providerResourceId property value. The providerResourceId property
func (m *Endpoint) GetProviderResourceId()(*string) {
    return m.providerResourceId
}
// GetUri gets the uri property value. The uri property
func (m *Endpoint) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *Endpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("capability", m.GetCapability())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerId", m.GetProviderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerName", m.GetProviderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerResourceId", m.GetProviderResourceId())
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
// SetCapability sets the capability property value. The capability property
func (m *Endpoint) SetCapability(value *string)() {
    m.capability = value
}
// SetProviderId sets the providerId property value. The providerId property
func (m *Endpoint) SetProviderId(value *string)() {
    m.providerId = value
}
// SetProviderName sets the providerName property value. The providerName property
func (m *Endpoint) SetProviderName(value *string)() {
    m.providerName = value
}
// SetProviderResourceId sets the providerResourceId property value. The providerResourceId property
func (m *Endpoint) SetProviderResourceId(value *string)() {
    m.providerResourceId = value
}
// SetUri sets the uri property value. The uri property
func (m *Endpoint) SetUri(value *string)() {
    m.uri = value
}
